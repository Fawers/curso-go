package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/raffs/go-labs/fs"
)

type Request struct {
	Method      string
	Path        string
	HTTPVersion string
	Headers     map[string]string
}

func (r *Request) String() string {
	return fmt.Sprintf("%#v", r)
}

func newRequest() *Request {
	r := new(Request)
	r.Headers = make(map[string]string)

	return r
}

func parseLines(lines []string) []*Request {
	var partialLine uint // will be reset for each request
	var currentRequest *Request
	requests := make([]*Request, 0, 1)

	for lineNum, lineText := range lines {
		fmt.Printf("Processando linha #%d: %q\n", lineNum, lineText)

		if lineText == "" {
			partialLine = 0
			continue
		}

		if partialLine == 0 {
			currentRequest = newRequest()
			requests = append(requests, currentRequest)
			pieces := strings.Split(lineText, " ")

			switch len(pieces) {
			case 3:
				currentRequest.HTTPVersion = pieces[2]
				fallthrough

			case 2:
				currentRequest.Method = pieces[0]
				currentRequest.Path = pieces[1]
			}

		} else {
			headerInfo := strings.Split(lineText, ":")
			currentRequest.Headers[headerInfo[0]] = strings.TrimSpace(headerInfo[1])
		}

		partialLine += 1
	}

	return requests
}

func main() {
	// linha0 := "GET /"
	// linha1 := "Host: 4linux.com.br"

	fileName := "request.txt"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	lines, err := fs.LerArquivo(fileName)
	if err != nil {
		fmt.Println("Erro ao tentar ler o arquivo: ", err.Error())
		os.Exit(1)
	}

	// 1. Identificar qual linha temos que processar de acordo o delimitador:
	//      Primeira linha: GET /
	//      Outras linhas: host: 4linux ...
	//
	// 2. Parsar o arquivo `request.txt` por linha de comando
	//
	// 3. (Desafio) Varios request HTTP, no mesmo payload
	//    GET /
	//    Host: 4linux
	//    ...
	//    <linha vazia>
	//    GET /
	//    Host: 4linux.com.br
	//    <linha vazia>

	requests := parseLines(lines)
	fmt.Println("\n\nRequests parseados:")
	fmt.Println(requests)
}
