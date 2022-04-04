package enderecoController

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

func TestGenerateTransaction(t *testing.T) {
	request := http.Request{}
	new_body_content := `{
		"target_account": "*******-1",
		"source_account": "TESTFIELD",
		"target_mail": "*******@gmail.com",
		"source_mail": "*******@gmail.com",
		"value": 0.00
	}`
	request.Body = ioutil.NopCloser(strings.NewReader(new_body_content))
	body := getBody(&request)
	data, err := generateTransaction(body)
	if err != nil {
		t.Errorf("Not expect error on '%s'", err)
	}
	if data.SourceAccount != "TESTFIELD" {
		t.Errorf("Source account expected '%s' on '%s'", "TESTFIELD", data.SourceAccount)
	}
}

// func TestGet(t *testing.T) {
// 	t.Run("Testing getting data from elasticsearch", func(t *testing.T) {
// 		url := "localhost:8090" //os.Getenv("HOST") + os.Getenv("LOCAL_PUBLISHER_PORT")
// 		fmt.Println(url)
// 		resp, err := http.Get(url)
// 		//fmt.Println(resp)
// 		defer resp.Body.Close()

// 		if err != nil {
// 			t.Errorf("Expect error nil, recieved '%s'", err)
// 		}

// 		if http.StatusOK != resp.StatusCode {
// 			t.Errorf("Expect status '%d' on '%d'", http.StatusCreated, resp.StatusCode)
// 		}
// 	})
// }

// func TestPost(t *testing.T) {
// 	t.Run("Testing data sent to kafka", func(t *testing.T) {
// 		postBody, _ := json.Marshal(map[string]string{
// 			"target_account": "xxxxx-1",
// 			"source_account": "xxxxx-1",
// 			"target_mail":    "rodrigo.aze7@gmail.com",
// 			"source_mail":    "rodrigo.azevedo.fernandes1@gmail.com",
// 			"value":          "0.00",
// 		})
// 		responseBody := bytes.NewBuffer(postBody)
// 		url := os.Getenv("HOST") + os.Getenv("LOCAL_PUBLISHER_PORT")
// 		resp, err := http.Post(url, "application/json", responseBody)
// 		//defer resp.Body.Close()

// 		if err != nil {
// 			t.Errorf("Expect error nil, recieved '%s'", err)
// 		}

// 		if http.StatusCreated != resp.StatusCode {
// 			t.Errorf("Expect status '%d' on '%d'", http.StatusCreated, resp.StatusCode)
// 		}

// 	})
// }
