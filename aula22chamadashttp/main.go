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

	res, err := io.ReadAll(req.Body) //Lendo o Body do retorno do response
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(res))
	}
	req.Body.Close()
}
