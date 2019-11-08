package number

import (
	"github.com/shopspring/decimal"
)

func Decimal(v string) decimal.Decimal {
	d, _ := decimal.NewFromString(v)
	return d
}
