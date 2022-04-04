package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"text/template"

	model "api-go-elasticsearch/mailService/models/pix"
)

func SendMail(message []byte) {

	messageData := getFormattedData(message)
	// Sender data.
	from := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		messageData.SourceMail,
		messageData.TargetMail,
	}

	// smtp server configuration.
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	workDir, _ := os.Getwd()
	t, _ := template.ParseFiles(workDir + "/mail/template.html")

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: PIX TRANSACTION \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		SourceEmail string
		TargetEmail string
		Date        string
		Value       float64
	}{
		SourceEmail: messageData.SourceMail,
		TargetEmail: messageData.TargetMail,
		Date:        messageData.TransactionTime,
		Value:       messageData.Value,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

func getFormattedData(data []byte) model.PixTransaction {
	var modelData model.PixTransaction
	json.Unmarshal(data, &modelData)
	return modelData
}
