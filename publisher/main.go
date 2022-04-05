package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "go-asynchronous-architecture/publisher/config"
	controller "go-asynchronous-architecture/publisher/controllers/pix"

	"github.com/gorilla/mux"
)

func main() {
	config.Init()
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/", controller.Get).Methods("GET")
	routes.HandleFunc("/", controller.Post).Methods("POST")
	port := os.Getenv("LOCAL_PUBLISHER_PORT")
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}
