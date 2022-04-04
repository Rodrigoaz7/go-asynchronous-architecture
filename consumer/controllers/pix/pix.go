package enderecoController

import (
	"encoding/json"
	"os"

	model "api-go-elasticsearch/consumer/models/pix"
	service "api-go-elasticsearch/consumer/service"
)

func PersistData(data []byte) error {
	var pixTransaction model.PixTransaction
	err := json.Unmarshal(data, &pixTransaction)
	if err != nil {
		return err
	}

	index := os.Getenv("ELASTICSEARCH_INDEX_NAME")
	ok := service.Create(pixTransaction, index)
	if ok != nil {
		return ok
	}

	return nil
}
