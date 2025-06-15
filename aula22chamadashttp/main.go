package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://serverest.dev/produtos") //Realizando uma requisição
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() //Usado para atrasar a execução de uma linha. Sera sempre a ultima linha a ser executada no arquivo

	res, err := io.ReadAll(req.Body) //Lendo o Body do retorno do response
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(res))
	}
	req.Body.Close()
}
