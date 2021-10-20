package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func computarVarios(ordenado []int) (r struct {
	Maior, Menor   int
	Media, Mediana float32
}) {
	length := len(ordenado)
	soma := somar(ordenado)

	r.Menor = ordenado[0]
	r.Maior = ordenado[length-1]
	r.Media = float32(soma) / float32(length)

	if length%2 == 0 {
		meio := length/2 - 1
		r.Mediana = float32(somar(ordenado[meio:meio+2])) / 2
	} else {
		r.Mediana = float32(ordenado[length/2])
	}

	return
}

func somar(numeros []int) (soma int) {
	for _, v := range numeros {
		soma += v
	}

	return
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Uso: %s <numero1> [numero2] ....\n", os.Args[0])
		os.Exit(1)
	}

	var numeros []int
	for _, valorStr := range os.Args[1:] {
		num, err := strconv.Atoi(valorStr)
		if err != nil {
			fmt.Printf("WARING: O valor '%s' na pode ser convertido para numero\n", valorStr)
			continue
		}

		numeros = append(numeros, num)
	}

	// fmt.Println(numeros)
	sort.Ints(numeros)
	// fmt.Println(numeros)

	// 1. Computar o menor
	// 2. Computar a media
	// 3. Computar o mediano // [1 *2* 3] (se a qtde for impar, pegar o valor do meio
	//                          [1 *2 3* 4] (se a qtde for par, media dos dois valores do meio)

	valores := computarVarios(numeros)
	fmt.Printf("%+v\n", valores)
}
