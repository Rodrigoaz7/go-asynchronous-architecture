package controllers

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetBody(t *testing.T) {
	request := http.Request{}
	new_body_content := "New content."
	request.Body = ioutil.NopCloser(strings.NewReader(new_body_content))
	body := getBody(&request)
	if len(body) == 0 {
		t.Errorf("Expect len > 0 on '%d'", len(body))
	}
}

// func TestGenerateTransaction(t *testing.T) {
// 	request := http.Request{}
// 	new_body_content := `{
// 		"target_account": "*******-1",
// 		"source_account": "TESTFIELD",
// 		"target_mail": "*******@gmail.com",
// 		"source_mail": "*******@gmail.com",
// 		"value": 0.00
// 	}`
// 	request.Body = ioutil.NopCloser(strings.NewReader(new_body_content))
// 	body := getBody(&request)
// 	data, err := generateTransaction(body)
// 	if err != nil {
// 		t.Errorf("Not expect error on '%s'", err)
// 	}
// 	if data.SourceAccount != "TESTFIELD" {
// 		t.Errorf("Source account expected '%s' on '%s'", "TESTFIELD", data.SourceAccount)
// 	}
// }
