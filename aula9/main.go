package main

import (
	"fmt"
)

func main() {
	valor := sum(1, 2, 5, 4)
	fmt.Printf("O Valor é: %d\n", valor)

	fmt.Println(sum(4, 5))
}

func sum(numeros ...int) int { //Isso é uma Função Variádica
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
