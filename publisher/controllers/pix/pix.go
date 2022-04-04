package enderecoController

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	connectionElasticSearch "api-go-elasticsearch/publisher/elasticsearch"
	kafka "api-go-elasticsearch/publisher/messages"
	model "api-go-elasticsearch/publisher/models/pix"
	"api-go-elasticsearch/publisher/service"

	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
)

func Get(w http.ResponseWriter, r *http.Request) {
	index := os.Getenv("ELASTICSEARCH_INDEX_NAME")
	dataResponse := service.FindAll(index, r.URL.Query())
	isExistsAnyRegisters := len(dataResponse) > 0

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if isExistsAnyRegisters {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(dataResponse); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body := getBody(r)
	pixData, err := generateTransaction(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	dataToSend, _ := json.Marshal(pixData)
	err = kafka.SendMessage(dataToSend)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}

func getBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	return body
}

func generateTransaction(body []byte) (model.PixTransaction, error) {
	var newPixTransaction model.PixTransaction
	err := json.Unmarshal(body, &newPixTransaction)
	currentTime := time.Now()
	newPixTransaction.TransactionTime = currentTime.Format("2006-01-02 15:04:05")
	return newPixTransaction, err
}

func getElasticSearchConnection() *elasticSearchInstance.Client {
	elasticInstance, err := connectionElasticSearch.ConnectElasticSearch()
	if err != nil {
		log.Fatalf("Error connecting to elasticsearch: %s", err)
		panic(err)
	}
	return elasticInstance
}