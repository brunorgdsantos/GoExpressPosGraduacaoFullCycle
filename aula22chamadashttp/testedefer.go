package main

import "fmt"

func main() {
	teste()
}

func teste() { //Perceba o comportamento do Defer
	fmt.Println("Primeira Linha!")
	fmt.Println("Segunda Linha!")
	fmt.Println("Terceira Linha!")

	fmt.Println("------------------")

	fmt.Println("Primeira Linha!")
	defer fmt.Println("Segunda Linha!")
	fmt.Println("Terceira Linha!")
}
