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
			inColumn: ColumnNameFirst{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFirstMale{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFirstFemale{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLast{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLastMale{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLastFemale{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFull{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFullMale{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFullFemale{
				column: Column{},
			},
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
