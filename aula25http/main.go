package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCep)    //Definindo a rota
	http.ListenAndServe(":8080", nil) //Apenas com uma linha já é possível subir um servidor Http
}

func BuscaCep(w http.ResponseWriter, r *http.Request) { //r *http.Request chama as requests, w http.ResponseWriter chama as responses
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep") //Definindo um parametro como obritagorio na URL
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest) // Se nenhum parametro for passado, vai dar erroS
		return
	}

	w.Header().Set("Content-Type", "application/json") //Definindo como deve ser o retorno
	w.WriteHeader(http.StatusOK)                       //Definindo que o status deve ser 200
	w.Write([]byte("Teste Http"))                      //escrevendo na tela
}
