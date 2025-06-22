package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() { //Os contexts podem ser aplicados do lado do Serve quanto do lado do Client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil) //Se essa url demorar mais que 5 segundos para responder, esse contexto sera cancelado
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req) //Na linha 15 vc preparou a request, aqui vc esta executando ela
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body) //Vai mostrar no Browser
}
