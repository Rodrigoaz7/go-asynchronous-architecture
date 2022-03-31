package enderecoController

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"api-go-elasticsearch/publisher/config"
	rabbitmq "api-go-elasticsearch/publisher/messages/rabbitmq"
	enderecoModel "api-go-elasticsearch/publisher/models/endereco"

	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
	"github.com/mitchellh/mapstructure"
)

const INDEX_NAME = "address"

func Get(w http.ResponseWriter, r *http.Request) {
	connection := getElasticSearchConnection()
	result, _ := connection.Search(connection.Search.WithIndex(INDEX_NAME), connection.Search.WithBody(nil))
	var queryResponse map[string]interface{}
	json.NewDecoder(result.Body).Decode(&queryResponse)

	var addressesResponse []enderecoModel.Message
	var addressReference enderecoModel.Message

	quantRegisters := queryResponse["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)

	if quantRegisters > 0 {
		for _, hit := range queryResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
			craft := hit.(map[string]interface{})["_source"]
			messageSoftwareInfo := craft.(map[string]interface{})["message_software"].(string)
			// this library cannot decode nested data, so was necessary to catch the message_software info above
			mapstructure.Decode(craft, &addressReference)
			addressReference.MessageSoftware = messageSoftwareInfo
			addressesResponse = append(addressesResponse, addressReference)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(addressesResponse); err != nil {
		panic(err)
	}

}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body := getBody(r)
	addressData, err := generateAddress(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	var output = []enderecoModel.Message{
		{
			Data:            addressData,
			MessageSoftware: "KAFKA",
		},
		{
			Data:            addressData,
			MessageSoftware: "RABBITMQ",
		},
	}

	dataToSend, _ := json.Marshal(output[1])
	err = rabbitmq.PublishMessage(dataToSend)

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

// Here we set the way error messages are displayed in the terminal.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
