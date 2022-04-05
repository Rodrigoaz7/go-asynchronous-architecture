package models

import (
	"encoding/json"
	"errors"
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
