package main

import "fmt"

func main() { //A unica estrutura de repetição no GO é o for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numeros := []string{"um", "dois", "tres"} //Como no [] esta vazio, não temos um Array, mas sim um Slice. Slice são dinamicos quanto ao tamanho, Arrays são estáticos.
	for k, v := range numeros {
		fmt.Println(k, v)
	}

	for _, v := range numeros { //O underscore chama-se Blank Identify
		fmt.Println(v)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Loop Infinito!")
	}
}
