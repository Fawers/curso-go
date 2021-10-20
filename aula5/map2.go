package main

import "fmt"

type Usuario map[string]string
type Usuarios map[string]Usuario

func (u *Usuarios) pegar(usuario string) *Usuario {
	if u_, ok := (*u)[usuario]; ok {
		return &u_
	}

	return nil
}

func main() {
	usuarios := Usuarios{
		"alice": {
			"email":    "alice@4linux.com.br",
			"telefone": "123456789",
		},
	}

	fmt.Printf("%v\n", usuarios)
	a := usuarios.pegar("alice")
	fmt.Println(a)
	(*a)["telefone"] = "no"
	fmt.Println(usuarios.pegar("bob"))
	fmt.Printf("%v\n", usuarios)
}
