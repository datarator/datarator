package main

import (
	"regexp"
	"testing"
)

func TestColumnDateValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnDateDayOfWeek{},
			outValue: "^[1-7]$",
		},
		{
			inColumn: ColumnDateDayOfWeekName{},
			outValue: "^[A-Za-z]+$",
		},
		{
			inColumn: ColumnDateDayOfMonth{},
			outValue: "^[0-9]+$",
		},
		{
			inColumn: ColumnDateMonth{},
			outValue: "^[0-9]+$",
		},
		{
			inColumn: ColumnDateMonthName{},
			outValue: "^[A-Za-z]+$",
		},
		{
			inColumn: ColumnDateYear{},
			outValue: "^[0-9]+$",
		},
		{
			inColumn: ColumnDateOfBirth{},
			outValue: "^[-+: 0-9A-Z]+$",
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
