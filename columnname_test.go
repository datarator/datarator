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
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFirstMale{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFirstFemale{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLast{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLastMale{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameLastFemale{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFull{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFullMale{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
		{
			inColumn: ColumnNameFullFemale{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z.]+$",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(test.inContext)
		matched, _ := regexp.MatchString(test.outValue, actual)
		if !matched {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
