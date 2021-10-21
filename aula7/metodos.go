package main

import (
	"fmt"
	"strings"
)

type Nome string

func (n Nome) primeiroNome() Nome {
	nomes := strings.Split(string(n), " ")

	if len(nomes) == 0 {
		return ""
	}

	return Nome(nomes[0])
}

func (n *Nome) caixaAlta() {
	*n = Nome(strings.ToUpper(string(*n)))
}

func main() {
	nomeComp := Nome("Fabricio Werneck")

	fmt.Println(nomeComp)
	fmt.Println(nomeComp.primeiroNome())
	nomeComp.caixaAlta()
	fmt.Println(nomeComp)
	fmt.Println(nomeComp.primeiroNome())
}
