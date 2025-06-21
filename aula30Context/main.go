package main

import (
	"context"
	"fmt"
	"time"
)

func main() { //Sempre que vc quiser cancelar processos, pode utilizar o Context
	ctx := context.Background()                             //Contexto inicial e rodando em background
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //Se demorar menos de 3 segundos, executa ctx, se mais executa o cancel

	defer cancel() //É uma boa pratica finalizar com defer cancel

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select { //Fica esperando algo para executar uma ação
	case <-ctx.Done(): //Ação finalizada. Se o timeout >3 segundos, sera cancelado. Essa linha vai ser executada
		fmt.Println("Hotel Booking cancelled. Timeout reached!")
		return
	case <-time.After(5 * time.Second): //Ação após time excedido. Por padrão timeSecond é 1 segundo
		fmt.Println("Hotel Booked")
	}
}
