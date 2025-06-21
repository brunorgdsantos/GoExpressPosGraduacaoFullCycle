package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second} //Tempo limite de espera de resposta
	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() //Sera executado por ultimo

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
