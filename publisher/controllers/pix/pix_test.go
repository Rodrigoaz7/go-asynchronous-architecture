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
