package main

import (
	"regexp"
	"testing"
)

func TestColumnNameFirstValue(t *testing.T) {
	var tests = []struct {
		inColumn  ColumnNameFirst
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnNameFirst{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z]+$",
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

func TestColumnNameLastValue(t *testing.T) {
	var tests = []struct {
		inColumn  ColumnNameLast
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnNameLast{
				Column: Column{},
			},
			inContext: Context{},
			outValue:  "^[- a-zA-Z]+$",
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
