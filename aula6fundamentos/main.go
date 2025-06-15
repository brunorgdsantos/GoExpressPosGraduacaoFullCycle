package main

import (
	"fmt"
)

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 100} /*Não é um Array, é um Slice. Para identificar
	basta ver o [] vazio. Se fosse Array teria um valor dentro de [].
	Slices podem ser dinamicos, Array são estaticos em relação ao tamanho.*/
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v \n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 120)                                   //Estou adicionando mais elemento ao Slice
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s) /*Veja que foi adicionado o 120 ao Slice
	Perceba que a capacidade foi dobrada 2x0 = 18. Ele pegou o array S e criou outro
	de mesmo tamanho.*/

	fmt.Printf("O tipo de s é: %T \n", s)
}
