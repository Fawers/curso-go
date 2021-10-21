package daemon

import (
	"fmt"
	"time"
)

func (d *Daemon) initPlugins() {
	for _, plugin := range d.Plugins {
		plugin.Init()
	}
}

func (d *Daemon) Start() {
	fmt.Println("Inicializando plugins...")
	d.initPlugins()

	for {
		for _, plugin := range d.Plugins {
			if err := plugin.Collect(); err != nil {
				fmt.Println("Erro ao coletar plugin:", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
