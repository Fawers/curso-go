package main

import (
	"fmt"
	"os"

	"github.com/Fawers/4monitor/daemon"
	"github.com/Fawers/4monitor/plugins"
)

func usage() {
	fmt.Printf(`Usage: %s CMD

Comandos:

	versao	- Mostra a versao do daemon
	start	- Inicializa o daemon
	plugins	- Descreve plugins instalados
`, os.Args[0])
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	d := daemon.New("v0.0.1-alpha")
	d.AddPlugin(plugins.NewCPU("Linux"))
	d.AddPlugin(plugins.NewMem(80))

	switch cmd {
	case "versao":
		d.Version()

	case "start":
		d.Start()

	case "plugins":
		d.DescribeAll()

	default:
		fmt.Printf("Comando inválido: `%s`\n", cmd)
	}
}
