package plugins

import (
	"fmt"

	"github.com/raffs/go-labs/sistema"
)

type Memory struct {
	critical float64
}

func NewMem(critical float64) *Memory {
	return &Memory{critical}
}

func (m *Memory) Init() {
	fmt.Println("Inicializando memória")
}

func (m *Memory) Collect() error {
	usoMem := sistema.GetInfoMemoriaDefault().UsoPerct()

	if usoMem > m.critical {
		return fmt.Errorf(
			"MEM: Uso acima do crítico (%.2f%%): %.2f%%",
			m.critical, usoMem)
	}

	fmt.Printf("MEM: %.2f%%\n", usoMem)
	return nil
}

func (m *Memory) Describe() string {
	return fmt.Sprintf("MEM: Uso de memória com crítico = %.2f%%",
		m.critical)
}
