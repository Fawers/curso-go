package daemon

import (
	"context"
	"fmt"
	"log"

	"github.com/Fawers/4discovery/store"
	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
)

func (d *Daemon) initDocker() (err error) {
	d.dockerCtx = context.Background()

	d.dockerClient, err = docker.NewClientWithOpts(
		docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return
	}

	return nil
}

func (d *Daemon) checkDocker() bool {
	info, err := d.dockerClient.Ping(d.dockerCtx)
	if err != nil {
		log.Println("[daemon] error pinging docker client:", err)
		return false
	}

	log.Println("[daemon] successfully connected to docker client:", info.APIVersion)
	return true
}

func (d *Daemon) addService(eid string) {
	container, err := d.dockerClient.ContainerInspect(d.dockerCtx, eid)
	if err != nil {
		log.Println("unable to inspect container", eid)
		return
	}

	name, exists := container.Config.Labels["4DISCOVERY_NAME"]
	if !exists {
		log.Printf("[daemon] container %q does not contain 4DISCOVERY_NAME. ignoring.\n",
			container.ID)
	}

	svc := store.NewService(
		container.Name, eid, container.Config.Image, container.NetworkSettings.IPAddress,
		container.NetworkSettings.EndpointID, container.Config.Hostname,
		container.Config.Labels, nil,
	)

	log.Println("[daemon] adding service", svc.Name)
	d.store.Add(name, svc)
}

func (d *Daemon) remService(eid string) {
	container, err := d.dockerClient.ContainerInspect(d.dockerCtx, eid)
	if err != nil {
		fmt.Println("Nao foi possivel pegar info do container(containerInspect):", eid)
		return
	}

	name, exists := container.Config.Labels["4DISCOVERY_NAME"]
	if !exists {
		fmt.Printf("O container com o ID %q nao possui 4DISCOVERY_NAME, ignorando...\n",
			container.ID)
		return
	}

	fmt.Println("Removendo o servico:", name)
	d.store.Del(name)
}

func (d *Daemon) dockerListener(comm chan<- int) {
	cEvents, cErr := d.dockerClient.Events(d.dockerCtx, types.EventsOptions{})

	for {
		select {
		case event := <-cEvents:
			if event.Type == "container" {
				switch event.Status {
				case "create":
					d.addService(event.ID)

				case "die":
					d.remService(event.ID)
				}
			}

		case err := <-cErr:
			log.Println("[daemon] got error from docker daemon:", err)
			comm <- 1
		}
	}
}
