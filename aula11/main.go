package main

import (
	"fmt"
)

type Cliente struct { //Struct são estruturas de valores conjuntos
	nome  string
	idade int
	valor float64
	ativo bool
}

func main() {
	teste := Cliente{
		nome:  "Bruno",
		idade: 42,
		valor: 42.0,
	}

	fmt.Printf("Nome: %s, idade: %d, valor: %.2f\n", teste.nome, teste.idade, teste.valor)

	teste.ativo = true //Mudando o valor da struct
	fmt.Println(teste.ativo)
}
