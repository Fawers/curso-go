package main

import "fmt"

func main() {
	var numero int64 = 0xdeadbeef
	var gravidade float32 = 9.8
	var unsigned uint = 0xffffffffffffffff
	var truefalse bool = true && false

	fmt.Printf("Numero: %d\nG: %.2f\nUInt: %d\n", numero, gravidade, unsigned)
	fmt.Println(truefalse)
}
