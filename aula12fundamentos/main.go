package main

import "fmt"

func main() { //PONTEIRO = endereço de memória
	a := 10
	fmt.Println(a)  //Ira retornar o valor de a
	fmt.Println(&a) //Ira retornar o endereço de memoria de a que esta armazenando o valor 10

	var ponteiros *int = &a //O *int mostra que é um ponteiro. Um ponteiro sempre é um endereço de memória
	fmt.Println(ponteiros)

	*ponteiros = 20 //O ponteiros tinha o msm endereço de memoria de a, portanto toda alteração reflete em a
	fmt.Println(a)

	b := &a
	fmt.Println(b)  //Vai imprimir o endereço de memoria
	fmt.Println(*b) //Vai imprimir o valor de a atribuido b
}
