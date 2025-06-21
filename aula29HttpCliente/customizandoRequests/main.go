package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}                       //Client que vai fazer a requisição
	req, err := http.NewRequest("GET", "http://google.com", nil) //Dados da requisição
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req) //Passando os dados da request par ao Client
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
