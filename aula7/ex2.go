package main

import "fmt"

type DivisionByZeroError struct {
	Dividend int
}

func (err *DivisionByZeroError) Error() string {
	return "division by zero"
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionByZeroError{a}
	}

	return a / b, nil
}

func main() {
	r, err := div(10, 0)

	if err != nil {
		fmt.Println(err.(*DivisionByZeroError))
		return
	}

	fmt.Println(r)
}
