package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TopUpModel struct {
	UserId  int64
	Amount  decimal.Decimal
	TopUpId uuid.UUID
	Cur     CurrencyEnum
}
