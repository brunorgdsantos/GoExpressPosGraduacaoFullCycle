package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("arquivo.txt") //Criando arquivo
	if err != nil {
		panic(err)
	}

	//tamanho, err :=file.Write([]bytes("Testes Manipulação Arquivos")) //Se vc não sabe o tipo da informação que ira escrever (string/int/float...), coloque como Bytes
	tamanho, err := file.WriteString("Testes Manipulação Arquivos") //Escrevendo no arquivo
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("O tamanho do arquivo em bytes é: %d\n", tamanho)
	}

	arquivo, err := os.ReadFile("arquivo.txt") //Lendo o arquivo
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(arquivo))
	}

	arq, err := os.Open("arquivo.txt") //Se vc tiver um arquivo muito grande, pode lê-lo por partes.
	if err != nil {
		panic(err)
	} else {
		reader := bufio.NewReader(arq) //Coloca o arquivo em um buffer e lê aos poucos
		buffer := make([]byte, 15)
		for {
			n, err := reader.Read(buffer)
			if err != nil {
				break
			}
			fmt.Println(string(buffer[:n]))
		}
	}

	err = os.Remove("arquivo.txt") //Usando para remover arquivos
	if err != nil {
		panic(err)
	}

	file.Close()
}
