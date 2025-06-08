package main /*Com exceção do pagote main, todos os outros arquivos segue
o mesmo nome de package igual ao diretorio em que eles estão incluidos.
Se não for o nome do diretorio, o nome deve ser maim!*/

import "fmt"

func main() {
	fmt.Printf("Hello World\n")

	println(a) /*Como o arquivo teste.go esta no mesmo pacote/diretorio que main.go,
	tudo que esta no teste.go pode ser usando no main.go*/
}
