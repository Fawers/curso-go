package plugins

import (
	"errors"
	"fmt"

	"github.com/raffs/go-labs/sistema"
)

type CPU struct {
	OS string
}

func NewCPU(os string) *CPU {
	return &CPU{os}
}

func (cpu *CPU) Init() {
	fmt.Println("Inicializando CPU", cpu.OS)
}

func (cpu *CPU) Collect() (err error) {
	usoCpu := sistema.CpuUso()

	if usoCpu < 0 {
		err = errors.New("falha ao coletar cpu")
	} else {
		fmt.Printf("CPU: %.2f%%\n", usoCpu)
	}

	return
}

func (cpu *CPU) Describe() string {
	return "CPU: Uso de CPU de " + cpu.OS
}
