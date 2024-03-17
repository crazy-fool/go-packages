package money

import (
	"github.com/shopspring/decimal"
)

const MoneyScale = 2

func NewMoneyFromString(str string, scale int32) decimal.Decimal {
	dml, err := decimal.NewFromString(str)
	if err != nil {
		dml = decimal.New(0, -scale)
	}

	return dml.Round(scale)
}

func NewMoneyFromInt(i int, scale int32) decimal.Decimal {
	return decimal.NewFromInt(int64(i)).Round(scale)
}

func FormatMoney(str string, scale int32) string {

	dml := NewMoneyFromString(str, scale)
	return dml.StringFixed(scale)
}
