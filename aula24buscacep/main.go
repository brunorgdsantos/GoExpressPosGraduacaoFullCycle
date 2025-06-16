package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] { //Lendo a url que for digitada no terminal
		req, err := http.Get(url)
		if err != nil { //Tratando erro
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v", err)
		}

		defer req.Body.Close() //Sera executada por ultimo

		res, err := io.ReadAll(req.Body) //Lendo o Body da requisição
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v", err)
		}

		var data ViaCep

		err = json.Unmarshal(res, &data) //Transformando de Json para Struct
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parce da resposta: %v", err)
		}

		fmt.Println(data) //Imprimindo todos os dados da Struc

		file, err := os.Create("cidade.txt") //Criando arquivo para salvar as informações
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v", err)
		}

		defer file.Close()

		//Salvando as informações no arquivo criado
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
	}
}
