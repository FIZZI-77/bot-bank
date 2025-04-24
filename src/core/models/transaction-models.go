package models

import "github.com/shopspring/decimal"

type TransactionModel struct {
	UserId        string
	Amount        decimal.Decimal
	Recipient     int64
	TransactionId string
	Cur           CurrencyEnum
}
