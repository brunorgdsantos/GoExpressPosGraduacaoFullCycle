package main

import (
	"errors"
	"fmt"
)

func main() {
	soma := sum(3, 5)
	fmt.Println(soma)

	soma2, verdadeirofalso := sum2(51, 4)
	fmt.Println(soma2, verdadeirofalso)

	soma3, erro := sum3(60, 4)
	fmt.Println(soma3, erro)

	if erro != nil {
		fmt.Println(erro)
	}
}

func sum(a, b int) int { //O que esta depois do () é o retorno, nesse caso sera um int
	return a + b
}

func sum2(a, b int) (int, bool) { //funções em GO podem retornar mais de um tipo, nesse caso sera um int e um bool
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

func sum3(a, b int) (int, error) { //Normalmente se usa o ultimo parametro de uma func para tratar as exceções pois em GO não ha try/catch
	if a+b >= 50 {
		return 0, errors.New("A Soma é Maior que 50")
	}
	return a + b, nil //Nil é o mesmo que nulo
}
