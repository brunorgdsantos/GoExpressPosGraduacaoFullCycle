package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("Request Iniciada!")
	defer fmt.Println("Request Finalizada!")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request Processada com Sucesso!")                       //Esse log imprime no stdout/command line
		w.Write([]byte("Request processada com sucesso - Mensagem Server!")) //Esse mensagem imprime no LocalHost/Browser
	case <-ctx.Done():
		log.Println("Request cancelada pelo Client!")                             //Imprime no command line
		http.Error(w, "Request cancelada pelo Client", http.StatusRequestTimeout) //Imprime no Browser
	}
}
