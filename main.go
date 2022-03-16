package main

import (
	"fmt"
	"log"
	"net/http"

	enderecoController "github.com/rodrigoaz7/api-go-elasticsearch/controllers/endereco"

	"github.com/gorilla/mux"
)

func main() {
	rotas := mux.NewRouter().StrictSlash(true)
	rotas.HandleFunc("/", enderecoController.Get).Methods("GET")
	rotas.HandleFunc("/", enderecoController.Create).Methods("POST")
	var port = ":8090"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))
}
