package main

import (
	"regexp"
	"testing"
)

func TestColumnDateValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnDateDayOfWeek{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[1-7]$",
		},
		{
			inColumn: ColumnDateDayOfWeekName{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[A-Za-z]+$",
		},
		{
			inColumn: ColumnDateDayOfMonth{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9]+$",
		},
		{
			inColumn: ColumnDateMonth{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9]+$",
		},
		{
			inColumn: ColumnDateMonthName{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[A-Za-z]+$",
		},
		{
			inColumn: ColumnDateYear{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9]+$",
		},
		{
			inColumn: ColumnDateOfBirth{
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[-+: 0-9A-Z]+$",
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
