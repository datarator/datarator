package main

import (
	"github.com/syscrusher/fake"
)

const (
	columnCreditCardNumber = "credit_card.number"
	columnCreditCardType   = "credit_card.type"
)

type ColumnCreditCardNumberPayload struct {
	Type string
}

type ColumnCreditCardNumber struct {
	TypedColumnBase
	payload ColumnCreditCardNumberPayload
}

func (column ColumnCreditCardNumber) Value(chunk *Chunk) (string, error) {
	return fake.CreditCardNum(column.payload.Type), nil
}

type ColumnCreditCardType struct {
	TypedColumnBase
}

func (column ColumnCreditCardType) Value(chunk *Chunk) (string, error) {
	return fake.CreditCardType(), nil
}
