package main

import (
	"encoding/json"
	"io"
	"net/http"
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

	dataCep, error := ProcurandoCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") //Definindo como deve ser o retorno
	w.WriteHeader(http.StatusOK)                       //Definindo que o status deve ser 200
	json.NewEncoder(w).Encode(dataCep)                 //Essa linha resume tudo feito entre as linhas 51 a 56

	/*resul, err := json.Marshal(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resul)*/
}

func ProcurandoCep(cep string) (*ViaCep, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil { //Tratando erro
		return nil, err
	}

	defer resp.Body.Close() //Irá executar por último

	body, err := io.ReadAll(resp.Body) //Lendo o Body
	if err != nil {
		return nil, err
	}

	var data ViaCep
	err = json.Unmarshal(body, &data) //Transformando de Json para Struct
	if err != nil {
		return nil, err
	}
	return &data, nil
}
