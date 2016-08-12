package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_CREDIT_CARD_NUMBER = "credit_card.number"
	COLUMN_CREDIT_CARD_TYPE   = "credit_card.type"
)

type ColumnCreditCardNumberPayload struct {
	Type string
}

type ColumnCreditCardNumber struct {
	Column  Column
	Payload ColumnCreditCardNumberPayload `json:"payload"`
}

func (column ColumnCreditCardNumber) Value(context *Context) (string, error) {
	return fake.CreditCardNum(column.Payload.Type), nil
}

type ColumnCreditCardType struct {
	Column Column
}

func (column ColumnCreditCardType) Value(context *Context) (string, error) {
	return fake.CreditCardType(), nil
}
