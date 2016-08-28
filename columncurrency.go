package main

import (
	"github.com/syscrusher/fake"
)

const (
	columnCurrency     = "currency"
	columnCurrencyCode = "currency.code"
)

type ColumnCurrency struct {
	TypedColumnBase
}

func (column ColumnCurrency) Value(context *Context) (string, error) {
	return fake.Currency(), nil
}

type ColumnCurrencyCode struct {
	TypedColumnBase
}

func (column ColumnCurrencyCode) Value(context *Context) (string, error) {
	return fake.CurrencyCode(), nil
}
