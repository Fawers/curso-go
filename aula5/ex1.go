package main

import (
	"fmt"
	"strings"
)

func extrairEndereco(linha string) string {
	return strings.SplitN(linha, " ", 2)[0]
}

func main() {
	linhas := []string{
		"8.8.8.8 [01010101] GET / ...",
		"8.8.8.4 [01010101] GET / ...",
		"8.8.8.8 [01010101] GET / ...",
		"8.8.8.8 [01010101] GET / ...",
		"8.8.8.4 [01010101] GET / ...",
		"8.8.8.8 [01010101] GET / ...",
	}

	for _, linha := range linhas {
		endereco := extrairEndereco(linha)
		fmt.Println("Extraindo Endereco", endereco)
	}
}
