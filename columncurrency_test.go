package main

import (
	"regexp"
	"testing"
)

func TestColumnCurrencyValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnCurrency{},
			outValue: "^[a-zA-Z ()]+$",
		},
		{
			inColumn: ColumnCurrencyCode{},
			outValue: "^[A-Z]{3}$",
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
