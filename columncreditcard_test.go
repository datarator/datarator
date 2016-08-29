package main

import (
	"regexp"
	"testing"
)

func TestColumnCreditCardTypeValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnCreditCardNumber{},
			outValue: "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "amex",
				},
			},
			outValue: "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "discover",
				},
			},
			outValue: "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "mastercard",
				},
			},
			outValue: "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "visa",
				},
			},
			outValue: "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardType{},
			outValue: "^(VISA|MasterCard|American Express|Discover)$",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(nil)
		matched, _ := regexp.MatchString(test.outValue, actual)
		if !matched {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
