package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCep)    //Definindo a rota
	http.ListenAndServe(":8080", nil) //Apenas com uma linha já é possível subir um servidor Http
}

func BuscaCep(w http.ResponseWriter, r *http.Request) { //r *http.Request chama as requests, w http.ResponseWriter chama as responses
	w.Write([]byte("Teste Http"))
}
