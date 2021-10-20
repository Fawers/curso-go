package main

import (
	"fmt"

	"github.com/raffs/go-labs/sistema"
)

func main() {
	var msg string

	usoCpu := sistema.GetInfoCpuDefault()
	usoMem := sistema.GetInfoMemoriaDefault().UsoPerct()

	if usoCpu <= 60 {
		msg = "normal"
	} else if usoCpu < 85 {
		msg = "atenção"
	} else {
		msg = "alerta de uso"
	}

	fmt.Printf("CPU: %.2f%% (%s)\n", usoCpu, msg)

	if usoMem <= 60 {
		msg = "normal"
	} else if usoMem < 85 {
		msg = "atenção"
	} else {
		msg = "alerta de uso"
	}

	fmt.Printf("Mem: %.2f%% (%s)\n", usoMem, msg)
}
