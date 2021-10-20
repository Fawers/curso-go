package main

import (
	"fmt"
)

func calcularMedia(nums []int) (m float64, s int) {
	for _, valor := range nums {
		s += valor
	}

	m = float64(s) / float64(len(nums))

	return
}

func main() {
	media, soma := calcularMedia([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Printf("Media=%.2f | Soma=%d\n", media, soma)
}
