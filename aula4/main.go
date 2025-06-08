package main

import "fmt" //Biblioteca para formatos

type ID float64

var (
	a float64 = 1.2
	b string  = "teste"
	c ID      = 1.5
)

func main() {
	fmt.Printf("O tipo é %T \n", a) //%T retorna o tipo da variavel, se é int/string/float...
	fmt.Printf("O tipo é %T \n", b)
	fmt.Printf("O tipo é %T \n", c) //Nesse caso, retornara o tipo main.ID
}
