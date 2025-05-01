package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionModel struct {
	UserId        uuid.UUID
	Amount        decimal.Decimal
	Recipient     int64
	TransactionId uuid.UUID
	Cur           CurrencyEnum
}
