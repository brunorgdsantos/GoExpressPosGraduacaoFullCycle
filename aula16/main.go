package main

import "fmt"

func main() {
	var minhaVar interface{} = "BS"

	print(minhaVar)
	fmt.Println(minhaVar.(string))

	resultado, ok := minhaVar.(int) //Se conversão der certo ok = true
	fmt.Printf("O valor de resultado é: %v e de OK é: %v\n", resultado, ok)

}
