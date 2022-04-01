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

	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
	"github.com/mitchellh/mapstructure"
)

func Get(w http.ResponseWriter, r *http.Request) {
	connection := getElasticSearchConnection()
	index := os.Getenv("ELASTICSEARCH_INDEX_NAME")
	result, _ := connection.Search(connection.Search.WithIndex(index), connection.Search.WithBody(nil))
	var queryResponse map[string]interface{}
	json.NewDecoder(result.Body).Decode(&queryResponse)

	var dataResponse []model.PixTransaction
	var pixReference model.PixTransaction

	quantRegisters := queryResponse["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	isExistsAnyRegisters := quantRegisters > 0
	if isExistsAnyRegisters {
		for _, hit := range queryResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
			item := hit.(map[string]interface{})["_source"]
			mapstructure.Decode(item, &pixReference)
			dataResponse = append(dataResponse, pixReference)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if isExistsAnyRegisters {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
	if err := json.NewEncoder(w).Encode(dataResponse); err != nil {
		panic(err)
	}

}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body := getBody(r)
	pixData, err := generateTransaction(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
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
