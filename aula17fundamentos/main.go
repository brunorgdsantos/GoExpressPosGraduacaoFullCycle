package main

import (
	"fmt"
)

type IntOuFloat interface {
	int | float64
}

func SomaInteiro(m map[string]int) int { //No GO map[string]int, significa uma chave no formato string e um valor no formato value, semelhante a estrutura do Json
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaFloat(m map[string]float64) float64 {
	var somaft float64
	for _, v := range m {
		somaft += v
	}
	return somaft
}

func unindoSomaIntFloat[T int | float64](m map[string]T) T { // [T int|float64] -> Isso é um Generics, estou dizendo que UnindoSomaIntFloat pode ser do tipo int OU float
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func unindoSomaIntFloatOutraFormaDeFazer[T IntOuFloat](m map[string]T) T { //Nesse caso estou usando uma Constraint definida em type IntOuFloat interface
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T IntOuFloat](a T, b T) bool { //Fazendo comparações usando Generics
	if a == b {
		return true
	}
	return false
}

func ComparaOutraForma[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{
		"Wesley": 100,
		"João":   200,
		"José":   300,
	}

	fmt.Println(SomaInteiro(m))

	f := map[string]float64{
		"Wesley Float": 100.10,
		"João Float":   200.20,
		"José Float":   300.30,
	}

	fmt.Println(somaFloat(f))

	fmt.Println(unindoSomaIntFloat(m)) //Usando Generics. Passando um map de int. Perceba que usamos uma unica função Generica que aceita valores tanto de Int quanto Float
	fmt.Println(unindoSomaIntFloat(f)) //Usando Generics. Passando um map de float64

	fmt.Println(unindoSomaIntFloatOutraFormaDeFazer(m)) //Usando Constraits
	fmt.Println(unindoSomaIntFloatOutraFormaDeFazer(f)) //Usando Constraits

	fmt.Println(Compara(10, 10.0))           //Usando a funcão Compara
	fmt.Println(ComparaOutraForma(10, 10.0)) //Usando a funcão comparable
}
