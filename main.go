package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "github.com/rodrigoaz7/api-go-elasticsearch/config"
	enderecoController "github.com/rodrigoaz7/api-go-elasticsearch/controllers/endereco"

	"github.com/gorilla/mux"
)

func main() {
	config.Init()
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/", enderecoController.Get).Methods("GET")
	routes.HandleFunc("/", enderecoController.Create).Methods("POST")
	port := os.Getenv("LOCAL_PORT")
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}
