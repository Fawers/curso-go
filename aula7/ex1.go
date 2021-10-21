package main

import (
	"fmt"
	"time"
)

type statusUsuario struct {
	logado      bool
	ultimoLogin time.Time
	processos   uint
}

type usuario struct {
	nome, email string
	idade       uint
	statusUsuario
}

// Metodo: login()
// Metodo: logout() // ele retorna `false` se o usuario ja estava deslogado

func (u *usuario) login() {
	if !u.logado {
		u.logado = true
		u.ultimoLogin = time.Now()
	}
}

func (u *usuario) logout() bool {
	if u.logado {
		u.logado = false
		return true
	}

	return false
}

func main() {
	var usuarioLogado usuario = usuario{
		nome:  "Alice",
		email: "alice@4linux.com.br",
		idade: 25,
	}

	fmt.Println("prélogin: ", usuarioLogado)
	usuarioLogado.login()
	fmt.Println("Usuario: ", usuarioLogado)
	fmt.Println("Logout: ", usuarioLogado.logout())
	fmt.Println("Usuario: ", usuarioLogado)
	fmt.Println("Logout: ", usuarioLogado.logout())
	fmt.Println("Usuario: ", usuarioLogado)
}
