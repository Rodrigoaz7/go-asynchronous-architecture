package endereco

import "time"

type PixTransaction struct {
	SourceAccount   string    `json:"source_account" mapstructure:"source_account"`
	TargetAccount   string    `json:"target_account" mapstructure:"target_account"`
	SourceMail      string    `json:"source_mail" mapstructure:"source_mail"`
	TargetMail      string    `json:"target_mail" mapstructure:"target_mail"`
	Value           float64   `json:"value" mapstructure:"value"`
	TransactionTime time.Time `json:"transaction_time" mapstructure:"transaction_time"`
}
