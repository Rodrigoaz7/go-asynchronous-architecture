package mail

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetFormattedData(t *testing.T) {
	stringMock := `{
		"target_account": "TESTFIELD",
		"source_account": "TESTFIELD",
		"target_mail": "rodrigo.aze7@gmail.com",
		"source_mail": "rodrigo.aze7@gmail.com",
		"value": 0.00
	}`

	reader := strings.NewReader(stringMock)
	data, _ := ioutil.ReadAll(reader)

	model := getFormattedData(data)
	if model.SourceMail != "rodrigo.aze7@gmail.com" {
		t.Errorf("Expect source email '%s' on '%s'", "rodrigo.aze7@gmail.com", model.SourceMail)
	}
}
