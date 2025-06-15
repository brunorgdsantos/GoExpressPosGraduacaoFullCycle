package matematica

type Carro struct { //A struc Carro esta sendo exportada pq esta com maiusculo
	marca string //Porém o atributo marca não esta sendo exportado
	Ano   string //Os atributos Ano e Preco estão sendo exportados pq estão com letras maiusculas
	Preco float64
}

func Soma[T int | float64](a, b T) T { //Em Go, quando uma func começa com letra maiuscula, significa que ela pode ser exportada e usada em outros arquivos.go
	return a + b
}

var A int = 10 //consigo exportar pq esta com letra maiúscula
var a int = 10 //consigo exportar pq esta com letra maiúscula
