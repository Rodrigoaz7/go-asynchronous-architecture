package enderecoController

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"api-go-elasticsearch/consumer/config"
	enderecoModel "api-go-elasticsearch/consumer/models/endereco"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
)

const INDEX_NAME = "address"

func Create(body []byte) error {
	var data enderecoModel.Message
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	connection := getElasticSearchConnection()
	randomId := generateRandomDocumentId()
	request := esapi.IndexRequest{Index: INDEX_NAME, DocumentID: randomId, Body: strings.NewReader(generateDataToPersist(data))}
	res, err := request.Do(context.Background(), connection)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func getElasticSearchConnection() *elasticSearchInstance.Client {
	elasticInstance, err := config.ConnectElasticSearch()
	if err != nil {
		log.Fatalf("Error connecting to elasticsearch: %s", err)
		panic(err)
	}
	return elasticInstance
}

func generateDataToPersist(addressModel enderecoModel.Message) string {
	data, _ := json.Marshal(addressModel)
	return string(data)
}

func generateRandomDocumentId() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(999999999))
}
