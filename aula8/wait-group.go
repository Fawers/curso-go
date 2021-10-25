package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(sleepTime time.Duration, id int) {
	fmt.Printf("Worker %d sleeping for %s\n", id, sleepTime)
	time.Sleep(sleepTime)
	fmt.Printf("Worker %d done sleeping\n", id)
}

func main() {
	var wg sync.WaitGroup
	workers := 4
	rand.Seed(time.Now().UnixMilli())

	fmt.Println("MAIN")
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		id := i
		go func() {
			worker(time.Duration(rand.Intn(10000))*time.Millisecond, id)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("MAIN wait done")
}
