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
	http.HandleFunc("/", BuscaCep)    //Definindo a rota
	http.ListenAndServe(":8080", nil) //Apenas com uma linha já é possível subir um servidor Http
}

func BuscaCep(w http.ResponseWriter, r *http.Request) { //r *http.Request chama as requests, w http.ResponseWriter chama as responses
	if r.URL.Path != "/" { //Se for passando algo na url diferente de "/" retornara erro
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep") //Definindo um parametro como obrigatorio na URL
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest) // Se nenhum parametro for passado, vai dar erro
		return
	}

	dataCep := ConsultaCepViaCepResolucaoPessoal(cepParam)

	w.Header().Set("Content-Type", "application/json") //Definindo como deve ser o retorno
	//w.WriteHeader(http.StatusOK)                       //Definindo que o status deve ser 200
	//w.Write([]byte("TESTE"))                      //Escrevendo na tela

	jsonData, err := json.Marshal(dataCep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

func ConsultaCepViaCepResolucaoPessoal(cep string) ViaCep {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil { //Tratando erro
		fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v", err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v", err)
	}

	var data ViaCep
	err = json.Unmarshal(res, &data) //Transformando de Json para Struct
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parce da resposta: %v", err)
	}
	return data
}
