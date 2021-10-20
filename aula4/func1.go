package main

import "fmt"

func divmod(a, b int) (r struct{ Div, Mod int }) {
	r.Div = a / b
	r.Mod = a % b
	return
}

func main() {
	minutos := 180

	hormin := divmod(minutos, 60)
	fmt.Printf("%v\n", hormin)
}
