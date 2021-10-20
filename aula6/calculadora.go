package main

import (
	"fmt"
	"os"

	"4linux.com.br/4calc/conv"
	"4linux.com.br/4calc/operations"
)

func main() {

	// <programa> add N N
	if len(os.Args) != 4 {
		fmt.Printf("usage: %s <add | sub | mul | div> <n1> <n2>\n", os.Args[0])
		os.Exit(-1)
	}

	operation := os.Args[1]
	opFunc, err := operations.GetOpFunc(operation)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	numbers, err := conv.Ints(os.Args[2:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if len(numbers) < 2 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("not enough numbers: %v (needs 2)", numbers))
	}

	a, b := numbers[0], numbers[1]
	result := opFunc(a, b)

	fmt.Printf("%s(%d, %d) = %d\n", operation, a, b, result)
}
