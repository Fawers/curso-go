package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada1 string = "  go  "
	var entrada2 string = "  GO  "
	var entrada3 string = "---Go----"

	fmt.Println("Entrada1 = ", strings.ToUpper(strings.TrimSpace(entrada1)) == "GO")
	fmt.Println("Entrada2 = ", strings.ToLower(strings.TrimSpace(entrada2)) == "go")
	fmt.Println("Entrada3 = ", strings.ToLower(strings.TrimFunc(entrada3, func(r rune) bool {
		return r == '-'
	})) == "go")
}
