package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contabilizarAcessos(acessos []string) (res map[string]int) {
	res = make(map[string]int)

	for _, acesso := range acessos {
		endereco := strings.SplitN(acesso, " ", 2)[0]
		res[endereco] += 1
	}

	return
}

func percentage(partial, total int) float32 {
	return 100 * float32(partial) / float32(total)
}

func main() {
	file, err := os.Open("apache_logs")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	requests := make([]string, 0, 16)

	for scanner.Scan() {
		requests = append(requests, scanner.Text())
	}

	accesses := contabilizarAcessos(requests)

	// fmt.Println(accesses)
	l := len(requests)

	for ad, reqs := range accesses {
		fmt.Printf("%15s => %3d (%.2f%%)\n", ad, reqs, percentage(reqs, l))
	}
}
