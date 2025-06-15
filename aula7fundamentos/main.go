package main

import "fmt"

func main() {
	salarios := map[string]int{"Jose": 1000, "João": 2000, "Maria": 5000} //O Map é uma estrutura chave valor
	fmt.Println(salarios["Jose"])
	fmt.Println(salarios["Maria"])

	delete(salarios, "Maria") //Para deletar do Map salarios
	salarios["Bruno"] = 10000 //Para adicionar no Map salarios
	fmt.Println(salarios["Bruno"])

	sal := make(map[string]int) //Estou inicializando map chamado sal, porém sem nehum valor
	sal1 := map[string]int{}    //Mesma coisa da linha 14
	sal["Wallace"] = 4111
	sal1["WS"] = 1000
	fmt.Println(sal["Wallace"])
	fmt.Println(sal1["WS"])

	for nome, salario := range salarios { //Percorrendo todo o Map
		fmt.Println(nome, salario)
	}

	for _, salario := range salarios { //Vou exibir os salarios na tela sem exibir os nomes, para isso basta usar o _
		fmt.Println("O salario é:", salario)
	}
}
