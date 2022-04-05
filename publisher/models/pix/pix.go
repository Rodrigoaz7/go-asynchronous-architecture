package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type PixTransaction struct {
	SourceAccount   string    `json:"source_account" mapstructure:"source_account" valid:"type(string), required"`
	TargetAccount   string    `json:"target_account" mapstructure:"target_account" valid:"type(string), required"`
	SourceMail      string    `json:"source_mail" mapstructure:"source_mail" valid:"email, required"`
	TargetMail      string    `json:"target_mail" mapstructure:"target_mail" valid:"email, required"`
	Value           float64   `json:"value" mapstructure:"value" valid:"float, required"`
	TransactionTime time.Time `json:"transaction_time" mapstructure:"transaction_time" valid:"optional"`
}

func (pix *PixTransaction) isValid() error {
	_, err := govalidator.ValidateStruct(pix)

	if err != nil {
		return err
	}

	if strings.EqualFold(pix.SourceAccount, pix.TargetAccount) || strings.EqualFold(pix.SourceMail, pix.TargetMail) {
		return errors.New("you can't send a pix to yourself")
	}

	if pix.Value <= 0 {
		return errors.New("a transaction should have a positive value")
	}

	return nil
}

func NewPixTransaction(data []byte) (*PixTransaction, error) {
	var newPixTransaction PixTransaction
	err := json.Unmarshal(data, &newPixTransaction)

	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	newPixTransaction.TransactionTime = currentTime

	errorValidation := newPixTransaction.isValid()

	if errorValidation != nil {
		return nil, errorValidation
	}

	return &newPixTransaction, nil
}

func NewPixTransactionFromElasticSearch(data interface{}) (*PixTransaction, error) {

	transactionTimeString := data.(map[string]interface{})["transaction_time"].(string)
	layout := "2006-01-02T15:04:05.0000000Z"
	transactionTime, err := time.Parse(layout, transactionTimeString)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transactionTime)

	newPixTransaction := PixTransaction{
		SourceAccount:   data.(map[string]interface{})["source_account"].(string),
		TargetAccount:   data.(map[string]interface{})["target_account"].(string),
		SourceMail:      data.(map[string]interface{})["source_mail"].(string),
		TargetMail:      data.(map[string]interface{})["target_mail"].(string),
		Value:           data.(map[string]interface{})["value"].(float64),
		TransactionTime: transactionTime,
	}

	errorValidation := newPixTransaction.isValid()

	if errorValidation != nil {
		return nil, errorValidation
	}

	return &newPixTransaction, nil
}
