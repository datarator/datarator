package main

import (
	"regexp"
	"testing"
)

func TestColumnCreditCardTypeValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn:  ColumnCreditCardNumber{},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "amex",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "discover",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "mastercard",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				payload: ColumnCreditCardNumberPayload{
					Type: "visa",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn:  ColumnCreditCardType{},
			inContext: Context{},
			outValue:  "^(VISA|MasterCard|American Express|Discover)$",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(&test.inContext)
		matched, _ := regexp.MatchString(test.outValue, actual)
		if !matched {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
