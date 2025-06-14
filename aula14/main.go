package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo float64
}

func NewConta() *Conta { //Ponteiros são muito usadas quando vc quer garantir que ao mudar o valor de uma variável, ela seja alterar em TODOS os lugares
	return &Conta{saldo: 0}
}

func (c Cliente) andou() {
	c.nome = "Albert Einstein"
	fmt.Printf("O cliente %v andou!\n", c.nome)
}

func (d Conta) simular(valor float64) float64 {
	d.saldo += valor
	return d.saldo
}

func main() {
	teste := Cliente{
		nome: "Bill",
	}

	fmt.Printf("O valor da Struct com nome %v\n", Cliente{nome: "Bill"})
	fmt.Printf("O valor da Struct com nome %v\n", teste.nome)
	teste.andou()

	conta := Conta{saldo: 100}
	conta.simular(2000)
	fmt.Println(conta.saldo)
	fmt.Println(conta.simular(0))
}
