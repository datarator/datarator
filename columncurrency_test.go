package main

import (
	"regexp"
	"testing"
)

func TestColumnCurrencyValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn:  ColumnCurrency{},
			inContext: Context{},
			outValue:  "^[a-zA-Z ()]+$",
		},
		{
			inColumn:  ColumnCurrencyCode{},
			inContext: Context{},
			outValue:  "^[A-Z]{3}$",
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
