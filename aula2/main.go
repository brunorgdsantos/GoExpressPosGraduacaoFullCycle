package main

const a = "Hello, World!" //Valores declarados como const não podem ser alterados!

var ( //A Declaração dessas variaveis tera escopo global dentro do package
	b bool           //Por padrão bool é sempre false, se não for definida
	c int            //Se não for denifido, int sera igual a 0
	d string         //Se for string não definida, sera atribuindo vazio, ficara em branco
	e float64        //Sera atribuido valores em 00000...
	g bool    = true //Nesse caso estou declarando e atribuindo uma variavel, assim como nas linhas 11 e 12 também
	h string  = "Declarando var string"
	i float64 = 1.2
)

func main() {
	println(a)

	b = true
	println(b)
	println(c)
	println(d)
	println(e)

	var f string /*É possível declarar variaveis dentro da propria func, porém variaveis
	declaradas dentro de uma func tem escopo local apenas. Ou seja, só da pra acessar
	apenas dentro da própria func*/
	println(f)

	println(g)
	println(h)
	println(i)

	/* b = 10 -> Go é fortemente tipada, ou seja, caso vc atribua uma var de um tipo, não
	pode inserir outros valores nela. Nesse caso b é bool e não int. Se uma var nasceu com
	um determinado tipo, deve ficar nele para sempre*/

	/*var j string = "teste" -> Se vc declarar uma var de escopo local (dentro de func)
	e não utiliza-la, o GO ira reclamar. Variaveis de escopo local devem ser utilizadas.
	Var() da linha 5 é escopo global, estão tudo bem!*/

	k := "testes" //É possível criar uma variavel e ja atribuir, o GO entendera de qual tipo ela é
	println(k)
}
