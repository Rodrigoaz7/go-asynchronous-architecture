package elasticsearch

import (
	"log"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func ConnectElasticSearch() (*elasticsearch.Client, error) {
	elasticsearchHost := os.Getenv("ELASTICSEARCH_HOST")
	elasticsearchPort := os.Getenv("ELASTICSEARCH_PORT")
	elasticsearchUrl := elasticsearchHost + elasticsearchPort
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticsearchUrl,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	log.Println("Elasticsearch connected successfull with version " + elasticsearch.Version)

	return es, nil
}
