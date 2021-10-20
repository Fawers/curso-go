package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Execute o programa com um argumento")
		return
	}

	num, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randint := random.Intn(1000)
	num += randint

	if num%2 == 0 {
		fmt.Printf("%d é par!\n", num)
	} else {
		fmt.Println(num, "não é par!")
	}
}
