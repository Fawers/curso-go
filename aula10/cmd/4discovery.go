package main

import (
	"fmt"
	"os"

	"github.com/Fawers/4discovery/daemon"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s URL\n", os.Args[0])
		os.Exit(-1)
	}

	url := os.Args[1]

	daemon, err := daemon.NewDaemon(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = daemon.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
