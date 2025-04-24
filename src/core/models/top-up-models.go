package models

import "github.com/shopspring/decimal"

type TopUpModel struct {
	UserId  int64
	Amount  decimal.Decimal
	TopUpId string
	Cur     CurrencyEnum
}
