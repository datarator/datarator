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
			inColumn: ColumnCreditCardNumber{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				column: Column{},
				Payload: ColumnCreditCardNumberPayload{
					Type: "amex",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				column: Column{},
				Payload: ColumnCreditCardNumberPayload{
					Type: "discover",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				column: Column{},
				Payload: ColumnCreditCardNumberPayload{
					Type: "mastercard",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardNumber{
				column: Column{},
				Payload: ColumnCreditCardNumberPayload{
					Type: "visa",
				},
			},
			inContext: Context{},
			outValue:  "^[0-9]{15,16}$",
		},
		{
			inColumn: ColumnCreditCardType{
				column: Column{},
			},
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
