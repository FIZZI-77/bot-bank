package models

import "github.com/shopspring/decimal"

type Balance map[CurrencyEnum]decimal.Decimal
