package main

import (
	"fmt"
	"time"
)

type statusUsuario struct {
	habilitado  bool
	ultimoLogin time.Time
	processos   uint
}

type usuario struct {
	nome, email string
	idade       uint
	status      statusUsuario
}

func main() {
	var u1 usuario

	u1.nome = "Fabricio"
	u1.email = "fabricio.werneck@4linux.com.br"
	u1.idade = 29
	u1.status.ultimoLogin = time.Now()

	u2 := usuario{nome: "Thiago", email: "thiago.viana@4linux.com.br", idade: 27}

	fmt.Println("U1: ", u1)
	fmt.Println("U2: ", u2)
}
