package operations

import "fmt"

type OpFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func GetOpFunc(op string) (OpFunc, error) {
	switch op {
	case "add":
		return add, nil

	case "sub":
		return sub, nil

	case "mul":
		return mul, nil

	case "div":
		return div, nil

	default:
		return nil, fmt.Errorf("invalid operation: `%s`", op)
	}
}
