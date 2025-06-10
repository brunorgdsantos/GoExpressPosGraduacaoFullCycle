package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro int
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct { //Struct são estruturas de valores conjuntos
	nome     string
	idade    int
	valor    float64
	ativo    bool
	Endereco          //Posso passar uma Struct dentro de uma Struct
	end      Endereco //Posso definir um nome de tipo Endereço
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

	teste.Logradouro = 13 //Observe que estou entrando no Struct Endereco dentro do Struct Cliente
	teste.Numero = 10
	teste.Cidade = "João Pessoa"
	teste.Estado = "PB"

	teste.Endereco.Logradouro = 13 //É o mesmo da linha 34
	teste.Endereco.Numero = 10
	teste.Endereco.Cidade = "João Pessoa"
	teste.Endereco.Estado = "PB"

	teste.end.Logradouro = 13 //Estou chamando a var end do tipo Endereço
	teste.end.Numero = 10
	teste.end.Cidade = "João Pessoa"
	teste.end.Estado = "PB"
}
