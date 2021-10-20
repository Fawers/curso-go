package main

import (
	"fmt"
)

func processarLinha(linha string, delim rune) (pos int) {
	for i, char := range linha {
		if char == delim {
			pos = i + 1
			break
		}
	}

	return
}

func main() {
	linhas := []struct {
		Text  string
		Delim rune
	}{
		{"GET /", ' '},
		{"Host: 4linux.com.br", ':'},
	}

	for i, linha := range linhas {
		fmt.Printf("Processando linha %d: %s\n", i+1, linha.Text)

		pos := processarLinha(linha.Text, linha.Delim)

		if pos == 0 {
			fmt.Println("Delimitador não encontrado")
		} else {
			fmt.Printf("`%c` encontrado na posição %d\n", linha.Delim, pos)
		}
	}
}
