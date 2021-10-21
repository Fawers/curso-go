package main

import (
	"fmt"
	"os"
)

type usuario struct {
	login, senha string
}

type SenhaInvalida struct{}

func (e *SenhaInvalida) Error() string {
	return "Acesso Negado!"
}

func validarAcesso(u usuario) error {
	if u.login == "admin" && u.senha == "admin" {
		return nil
	}
	return &SenhaInvalida{}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Uso: %s <usuario> <senha>\n", os.Args[0])
		os.Exit(1)
	}

	login := os.Args[1]
	senha := os.Args[2]
	usuario := usuario{
		login: login,
		senha: senha,
	}

	if err := validarAcesso(usuario); err != nil {
		fmt.Println("Erro ao validar o acesso:", err)
		return
	}

	fmt.Println("Usuario logado com sucesso, ...")
}
