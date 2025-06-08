package main

type ID int
type id string

var (
	a ID = 1 /*Perceba que ID é um typo, estão pode ser usado em outras variaveis.
	Nesse caso, f sera do tipo int*/
	b id = "testes"
)

func main() {
	println(a)
	println(b)
}
