package daemon

import (
	"fmt"

	"github.com/Fawers/4monitor/api"
)

type Daemon struct {
	version string
	Plugins []api.Plugin
}

func New(version string) (d Daemon) {
	d.version = version
	d.Plugins = make([]api.Plugin, 0)
	return
}

func (d *Daemon) AddPlugin(p api.Plugin) {
	d.Plugins = append(d.Plugins, p)
}

func (d *Daemon) DescribeAll() {
	for _, p := range d.Plugins {
		fmt.Println(p.Describe())
	}
}
