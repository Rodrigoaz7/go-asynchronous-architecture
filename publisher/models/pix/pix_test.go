package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPixTransaction(t *testing.T) {
	content := `{
		"source_account": "******-1",
		"target_account": "******-2",
		"target_mail": "myemail@gmail.com",
		"source_mail": "myemail2@gmail.com",
		"value": 10
	}`
	dataAsBytes := []byte(content)
	model, err := NewPixTransaction(dataAsBytes)

	require.Nil(t, err)
	require.Equal(t, "******-1", model.SourceAccount)
	require.Equal(t, "myemail@gmail.com", model.TargetMail)
	require.Equal(t, float64(10), model.Value)
}
