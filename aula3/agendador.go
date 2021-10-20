package main56789

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/raffs/go-labs/agendador"
)

func Main() {
	fmt.Println("Inicializando o agendador de tarefa")

	confFile := "agendador.conf"

	if len(os.Args) >= 2 {
		confFile = os.Args[1]
	}

	tarefas, _ := agendador.LerConfigAgendador(confFile)

	for {
		agora := time.Now()
		segundoAtual := agora.Second()
		minutoAtual := agora.Minute()
		horaAtual := agora.Hour()

		for _, tarefa := range tarefas {
			if tarefa.Minuto == minutoAtual {
				fmt.Printf("[Minuto=%d] Executando comando %s\n", minutoAtual, tarefa.Command)
				exec.Command(tarefa.Command).Start()
			}

			if tarefa.Segundo == segundoAtual {
				fmt.Printf("[Segundo=%d] Executando comando %s\n", segundoAtual, tarefa.Command)
				exec.Command(tarefa.Command).Start()
			}

			if tarefa.Hora == horaAtual {
				fmt.Printf("[hora=%d] Executando comando %s\n", horaAtual, tarefa.Command)
				exec.Command(tarefa.Command).Start()
			}
		}

		time.Sleep(200)
	}
}
