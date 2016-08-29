package main

import (
	"regexp"
	"testing"
)

func TestColumnAddressValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnAddressContinent{},
			outValue: "^[a-zA-Z ]+$",
		},
		{
			inColumn: ColumnAddressCountry{},
			outValue: "^[a-zA-Z ,]+$",
		},
		{
			inColumn: ColumnAddressCity{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressPhone{},
			outValue: "^[0-9-]+$",
		},
		{
			inColumn: ColumnAddressState{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressState{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnAddressStreet{},
			outValue: "^[- a-zA-Z.0-9]+$",
		},
		{
			inColumn: ColumnAddressZip{},
			outValue: "^[0-9]+$",
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
