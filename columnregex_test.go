package main

import (
	"regexp"
	"testing"
)

func TestColumnRegexValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnRegex{
				payload: ColumnRegexPayload{
					Regex: "foo",
					Limit: 10,
				},
			},
			outValue: "foo",
		},
		{
			inColumn: ColumnRegex{
				payload: ColumnRegexPayload{
					Regex: "f{1,1}",
					Limit: 10,
				},
			},
			outValue: "f",
		},
		{
			inColumn: ColumnRegex{
				payload: ColumnRegexPayload{
					Regex: "f{1,1}",
					Limit: 10,
				},
			},
			outValue: "f",
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
