package main

import (
	"fmt"
)

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaUsandoPonteiro(a, b *int) int {
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	fmt.Println(minhaVar1, minhaVar2)
	fmt.Println(soma(minhaVar1, minhaVar2))

	fmt.Println(somaUsandoPonteiro(&minhaVar1, &minhaVar2))
	fmt.Println(minhaVar1, minhaVar2)
}
