package daemon

import "fmt"

func (d *Daemon) Version() string {
	fmt.Println("Daemon version", d.version)
	return d.version
}
