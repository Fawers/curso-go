package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 5

	v := <-c

	fmt.Println(v)
}
