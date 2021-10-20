package main

import "fmt"

const (
	DOMINGO = 0
	SEGUNDA = 1
	TERCA   = 2
	QUARTA  = 3
	QUINTA  = 4
	SEXTA   = 5
	SABADO  = 6
)

func main() {
	diaDaSemana := QUARTA

	switch diaDaSemana {
	case DOMINGO:
		fmt.Println("Domingo")
	case SEGUNDA:
		fmt.Println("Segunda")
	case TERCA:
		fmt.Println("Terça")
	case QUARTA:
		fmt.Println("Quarta")
		fallthrough
	case QUINTA:
		fmt.Println("Quinta")
	case SEXTA:
		fmt.Println("Sexta")
	case SABADO:
		fmt.Println("Sabado")
	}

}
