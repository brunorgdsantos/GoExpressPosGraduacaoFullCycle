package main

import "fmt"

var ()

func main() {
	var meuArray [3]int
	fmt.Println(meuArray)

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10
	fmt.Println(meuArray)
	fmt.Println(len(meuArray))             //Retorna o tamanho do array
	fmt.Println(len(meuArray) - 1)         //Retorna len(3) - 1
	fmt.Println(meuArray[len(meuArray)-1]) //Retorna o elemente da posição 2
	fmt.Println(meuArray[2])

	for i := 0; i < len(meuArray); i++ {
		fmt.Println(meuArray[i])
	}

	for i, v := range meuArray {
		fmt.Printf("O valor do index:%d é %d \n", i, v)
	}
}
