package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"` //Voce pode usar tags para chamar os atributos
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 100,
		Saldo:  200,
	}

	res, err := json.Marshal(conta) //Marshal transforma tudo que for passado em () em Json
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) //O Encoder o valor ja é retornado automaticamente, no Marshal não
	if err != nil {
		fmt.Println(err)
	}
	//encoder.Encode(conta)
	//fmt.Println(err)

	//A linhas acima transformam do Struc para Json, vamo ver como transformar do Json para Struct abaixo
	jsonPuro := []byte(`{"Numero":2, "Saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaX.Saldo, contaX.Numero)

	jsonPuroTest := []byte(`{"n":2, "s":200}`) //Usando as tags definidas no Struct. Porém quando são definidas, vc é obrigado a usar
	var contaXteste Conta
	err = json.Unmarshal(jsonPuroTest, &contaXteste)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaXteste.Saldo, contaXteste.Numero)
}
