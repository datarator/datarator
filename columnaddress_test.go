package main

import (
	"regexp"
	"testing"
)

func TestColumnAddressValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnAddressContinent{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[a-zA-Z ]+$",
		},
		{
			inColumn: ColumnAddressCountry{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[a-zA-Z ,]+$",
		},
		{
			inColumn: ColumnAddressCity{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressPhone{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9-]+$",
		},
		{
			inColumn: ColumnAddressState{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressState{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressStreet{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressZip{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9]+$",
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
