package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_CURRENCY      = "currency"
	COLUMN_CURRENCY_CODE = "currency.code"
)

type ColumnCurrencyPayload struct {
	XmlType string `json:"xml"`
}

type ColumnCurrency struct {
	column  Column
	Payload ColumnCurrencyPayload `json:"payload"`
}

func (column ColumnCurrency) Value(context *Context) (string, error) {
	return fake.Currency(), nil
}

func (column ColumnCurrency) Column() Column {
	return column.column
}

func (column ColumnCurrency) XmlType() string {
	return column.Payload.XmlType
}

type ColumnCurrencyCodePayload struct {
	XmlType string `json:"xml"`
}

type ColumnCurrencyCode struct {
	column  Column
	Payload ColumnCurrencyCodePayload `json:"payload"`
}

func (column ColumnCurrencyCode) Value(context *Context) (string, error) {
	return fake.CurrencyCode(), nil
}

func (column ColumnCurrencyCode) Column() Column {
	return column.column
}

func (column ColumnCurrencyCode) XmlType() string {
	return column.Payload.XmlType
}
