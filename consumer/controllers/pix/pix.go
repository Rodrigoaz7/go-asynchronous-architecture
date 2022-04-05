package enderecoController

import (
	"os"

	model "go-asynchronous-architecture/consumer/models/pix"
	service "go-asynchronous-architecture/consumer/service"
)

func PersistData(data []byte) error {
	pixTransaction, err := model.NewPixTransaction(data)

	if err != nil {
		return err
	}

	index := os.Getenv("ELASTICSEARCH_INDEX_NAME")
	ok := service.Create(*pixTransaction, index)
	if ok != nil {
		return ok
	}

	return nil
}
