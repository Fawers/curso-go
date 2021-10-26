package daemon

import (
	"context"
	"fmt"

	"github.com/Fawers/4discovery/store"
	docker "github.com/docker/docker/client"
)

type Daemon struct {
	HTTPEndpoint string
	dockerClient *docker.Client
	dockerCtx    context.Context
	store        store.ServiceStore
}

func NewDaemon(endpoint string) (d *Daemon, err error) {
	d = &Daemon{HTTPEndpoint: endpoint}

	err = d.initDocker()
	if err != nil {
		return nil, err
	}

	if !d.checkDocker() {
		return nil, fmt.Errorf("docker not responding")
	}

	d.store = store.NewStore()

	return
}

func (d *Daemon) Start() error {
	comm := make(chan int)

	go d.startHTTPServer(comm)
	go d.dockerListener(comm)

	if code := <-comm; code != 0 {
		return fmt.Errorf("[daemon] something errored")
	}

	return nil
}
