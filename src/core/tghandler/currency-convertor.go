package tghandler

import (
	"fmt"
	"tg_transaction/src/core/models"
)

var stringToCurrency = map[string]models.CurrencyEnum{
	"RUB": models.RUB,
	"USD": models.USD,
	"EUR": models.EUR,
}

func ParseCurrencyToEnum(s string) (models.CurrencyEnum, error) {
	if val, ok := stringToCurrency[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("invalid currency: %s", s)
}

var currencyToString = map[models.CurrencyEnum]string{
	models.RUB: "RUB",
	models.USD: "USD",
	models.EUR: "EUR",
}
