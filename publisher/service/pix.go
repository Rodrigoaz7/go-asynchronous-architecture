package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	connectionElasticSearch "go-asynchronous-architecture/publisher/elasticsearch"
	model "go-asynchronous-architecture/publisher/models/pix"

	elasticSearchInstance "github.com/elastic/go-elasticsearch/v7"
)

func FindAll(index string, params url.Values) []model.PixTransaction {
	connection := getElasticSearchConnection()

	body := generateSearchBody(params)
	result, _ := connection.Search(
		connection.Search.WithIndex(index),
		connection.Search.WithBody(&body),
		connection.Search.WithTrackTotalHits(true),
		connection.Search.WithPretty(),
	)
	defer result.Body.Close()

	var queryResponse map[string]interface{}
	err := json.NewDecoder(result.Body).Decode(&queryResponse)
	if err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil
	}

	var dataResponse []model.PixTransaction

	isExistsIndex := queryResponse["hits"] != nil
	if !isExistsIndex {
		return []model.PixTransaction{}
	}

	quantRegisters := queryResponse["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	isExistsAnyRegisters := quantRegisters > 0
	if isExistsAnyRegisters {
		for _, hit := range queryResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
			item := hit.(map[string]interface{})["_source"]
			pixReference, err := model.NewPixTransactionFromElasticSearch(item)
			if err == nil {
				fmt.Println(*pixReference)
				dataResponse = append(dataResponse, *pixReference)
			}
		}
	}

	return dataResponse
}

func generateSearchBody(params url.Values) bytes.Buffer {
	var buf bytes.Buffer
	if len(params) > 0 {
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match_phrase_prefix": map[string]interface{}{
					"source_mail": params.Get("email"),
				},
			},
		}
		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query: %s", err)
		}
	}
	return buf
}

func getElasticSearchConnection() *elasticSearchInstance.Client {
	elasticInstance, err := connectionElasticSearch.ConnectElasticSearch()
	if err != nil {
		log.Fatalf("Error connecting to elasticsearch: %s", err)
		panic(err)
	}
	return elasticInstance
}
