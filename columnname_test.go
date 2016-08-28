package main

import (
	"regexp"
	"testing"
)

func TestColumnNameValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn:  ColumnNameFirst{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameFirstMale{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameFirstFemale{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameLast{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameLastMale{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameLastFemale{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameFull{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameFullMale{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn:  ColumnNameFullFemale{},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
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
