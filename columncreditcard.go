package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_CREDIT_CARD_NUMBER = "credit_card.number"
	COLUMN_CREDIT_CARD_TYPE   = "credit_card.type"
)

type ColumnCreditCardNumberPayload struct {
	XmlType string `json:"xml"`
	Type    string
}

type ColumnCreditCardNumber struct {
	column  Column
	Payload ColumnCreditCardNumberPayload `json:"payload"`
}

func (column ColumnCreditCardNumber) Value(context *Context) (string, error) {
	return fake.CreditCardNum(column.Payload.Type), nil
}

func (column ColumnCreditCardNumber) Column() Column {
	return column.column
}

func (column ColumnCreditCardNumber) XmlType() string {
	return column.Payload.XmlType
}

type ColumnCreditCardTypePayload struct {
	XmlType string `json:"xml"`
}

type ColumnCreditCardType struct {
	column  Column
	Payload ColumnCreditCardTypePayload `json:"payload"`
}

func (column ColumnCreditCardType) Value(context *Context) (string, error) {
	return fake.CreditCardType(), nil
}

func (column ColumnCreditCardType) Column() Column {
	return column.column
}

func (column ColumnCreditCardType) XmlType() string {
	return column.Payload.XmlType
}
