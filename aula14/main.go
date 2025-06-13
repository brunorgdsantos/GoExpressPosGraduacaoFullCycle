package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Cliente Teste"
	fmt.Printf("O cliente %v andou!", c.nome)
}

func main() {
	teste := Cliente{
		nome: "Bill",
	}

	teste.andou()
}
