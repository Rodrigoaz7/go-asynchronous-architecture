package enderecoController

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	enderecoModel "github.com/rodrigoaz7/api-go-elasticsearch/models/endereco"
)

var saida = []enderecoModel.Endereco{
	{Rua: "Rua das Flores", Cep: "58255-009", Bairro: "Jardim Botânico",
		Cidade: "São José das Letras", Estado: "ES", Pais: "Brasil"},
	{Rua: "Rua das Espadas", Cep: "45444-009", Bairro: "Bezerrão",
		Cidade: "Pelotas", Estado: "RS", Pais: "Brasil"},
}

func Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(saida)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var novoEndereco enderecoModel.Endereco

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &novoEndereco); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, &novoEndereco)

	saida = append(saida, novoEndereco)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(saida); err != nil {
		panic(err)
	}
}
