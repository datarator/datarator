package main

import (
	"regexp"
	"testing"
)

func TestColumnNameValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnNameFirst{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFirstMale{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFirstFemale{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLast{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLastMale{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLastFemale{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFull{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFullMale{},
			outValue: "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFullFemale{},
			outValue: "^[- a-zA-Z.]+$",
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
