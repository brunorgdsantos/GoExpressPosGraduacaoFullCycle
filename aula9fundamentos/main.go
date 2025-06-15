package main

import (
	"fmt"
)

func main() {
	valor := sum(1, 2, 5, 4)
	fmt.Printf("O Valor é: %d\n", valor)

	fmt.Println(sum(4, 5))

	funcaoAnonima := func() int { //No GO é possível ter funções anonimas, basta não definir um nome para ela.
		//É possível ter funções dentro de funções
		return sum(3, 9) * 2
	}()

	fmt.Printf("O valor é da Função Anônima: %d\n ", funcaoAnonima)
}

func sum(numeros ...int) int { //Isso é uma Função Variádica
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
