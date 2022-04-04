package enderecoController

import (
	"encoding/json"
	"os"

	model "go-asynchronous-architecture/consumer/models/pix"
	service "go-asynchronous-architecture/consumer/service"
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
