package main

import "fmt"

type Test interface {
	//Uma Interface vazia significa que ela implementa em tudo, não há limite
}

func main() {
	var x interface{} = 10            //Interface vazia
	var y interface{} = "Hello World" //Interface vazia

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é: %T e o valor é %v\n", t, t)
}
