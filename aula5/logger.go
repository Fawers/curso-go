package main

import "fmt"

const (
	versao string = "0.0.1"
)

// Precisa: [INFO] texto ....
func log(formato string, opts ...interface{}) {
	fmt.Printf(formato, opts...)
}

func info(formato string, opts ...interface{}) {
	log(fmt.Sprintf("[INFO] %s\n", formato), opts...)
}

func warn(formato string, opts ...interface{}) {
	log(fmt.Sprintf("[WARN] %s\n", formato), opts...)
}

func main() {
	info("Inicializando o servidor de paginas web (%s)...", versao)
	warn("Não conectado à internet.")
}
