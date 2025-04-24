package repository

import (
	"fmt"
	"tgtransaction/src/core/models"
)

var currencyToString = map[models.CurrencyEnum]string{
	models.RUB: "RUB",
	models.USD: "USD",
	models.EUR: "EUR",
}

func FromCurrencyEnum(e models.CurrencyEnum) string {
	if val, ok := currencyToString[e]; ok {
		return val
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	panic(fmt.Sprintf("unknown currency enum: %d", e))

}

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
