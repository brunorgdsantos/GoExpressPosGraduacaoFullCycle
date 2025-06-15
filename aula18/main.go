package main

import (
	"fmt"
	"github.com/brunorgdsantos/GoExpressPosGraduacaoFullCycle/aula18/matematica"
)

func main() {
	sm := matematica.Soma(1, 2)
	fmt.Printf("O resultado é: %v\n", sm)

	fmt.Println(matematica.A) //Consigo acessar pq A esta com maiusculo
	//fmt.Println(matematica.a) //Não consigo acessar pq a esta minusculo

	teste := matematica.Carro{
		Ano:   "2025",
		Preco: 125.5,
	}
	fmt.Printf("O resultado é: %v e %v\n", teste.Ano, teste.Preco)
}
