// https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide#how-to-build-a-console-menu-in-go
// https://karmi.github.io/gotalks/go-elasticsearch-files/saved_resource.html
// https://www.youtube.com/watch?v=3arH8SgCdIs
// https://x-team.com/blog/set-up-rabbitmq-with-docker-compose/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "api-go-elasticsearch/publisher/config"
	enderecoController "api-go-elasticsearch/publisher/controllers/endereco"

	"github.com/gorilla/mux"
)

func main() {
	config.Init()
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/", enderecoController.Get).Methods("GET")
	routes.HandleFunc("/", enderecoController.SendMessage).Methods("POST")
	port := os.Getenv("LOCAL_PORT")
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}
