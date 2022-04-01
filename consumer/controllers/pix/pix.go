package enderecoController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	config "api-go-elasticsearch/consumer/elasticsearch"
	model "api-go-elasticsearch/consumer/models/pix"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
)

func PersistData(data []byte) error {
	var pixTransaction model.PixTransaction
	err := json.Unmarshal(data, &pixTransaction)
	if err != nil {
		return err
	}

	fmt.Println(pixTransaction)

	connection := getElasticSearchConnection()
	randomId := generateRandomDocumentId()
	index := os.Getenv("ELASTICSEARCH_INDEX_NAME")
	request := esapi.IndexRequest{Index: index, DocumentID: randomId, Body: strings.NewReader(generateDataToPersist(pixTransaction))}
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

func generateDataToPersist(pixTransaction model.PixTransaction) string {
	data, _ := json.Marshal(pixTransaction)
	return string(data)
}

func generateRandomDocumentId() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(999999999))
}
