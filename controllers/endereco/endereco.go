package enderecoController

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"

	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
	"github.com/mitchellh/mapstructure"
	"github.com/rodrigoaz7/api-go-elasticsearch/config"
	enderecoModel "github.com/rodrigoaz7/api-go-elasticsearch/models/endereco"
)

const INDEX_NAME = "address"

func Get(w http.ResponseWriter, r *http.Request) {
	connection := getElasticSearchConnection()
	result, _ := connection.Search(connection.Search.WithIndex(INDEX_NAME), connection.Search.WithBody(nil))
	var queryResponse map[string]interface{}
	json.NewDecoder(result.Body).Decode(&queryResponse)

	var addressesResponse []enderecoModel.Endereco
	var addressReference enderecoModel.Endereco

	for _, hit := range queryResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
		craft := hit.(map[string]interface{})["_source"]
		mapstructure.Decode(craft, &addressReference)
		addressesResponse = append(addressesResponse, addressReference)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(addressesResponse); err != nil {
		panic(err)
	}

}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body := getBody(r)
	newAddress, err := generateAddress(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	connection := getElasticSearchConnection()

	bodyOutput := generateOutput(&newAddress)
	randomId := generateRandomDocumentId()

	request := esapi.IndexRequest{Index: INDEX_NAME, DocumentID: randomId, Body: strings.NewReader(string(bodyOutput))}
	res, _ := request.Do(context.Background(), connection)
	defer res.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

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

func generateAddress(body []byte) (enderecoModel.Endereco, error) {
	var newAddress enderecoModel.Endereco
	err := json.Unmarshal(body, &newAddress)
	return newAddress, err
}

func getElasticSearchConnection() *elasticSearchInstance.Client {
	elasticInstance, err := config.ConnectElasticSearch()
	if err != nil {
		log.Fatalf("Error connecting to elasticsearch: %s", err)
		panic(err)
	}
	return elasticInstance
}

func generateOutput(addressModel *enderecoModel.Endereco) string {
	output, _ := json.Marshal(addressModel)
	return string(output)
}

func generateRandomDocumentId() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(999999999))
}
