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
			inColumn:  ColumnAddressContinent{},
			inContext: Context{},
			outValue:  "^[a-zA-Z ]+$",
		},
		{
			inColumn:  ColumnAddressCountry{},
			inContext: Context{},
			outValue:  "^[a-zA-Z ,]+$",
		},
		{
			inColumn:  ColumnAddressCity{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnAddressPhone{},
			inContext: Context{},
			outValue:  "^[0-9-]+$",
		},
		{
			inColumn:  ColumnAddressState{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnAddressState{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnAddressStreet{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnAddressZip{},
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
