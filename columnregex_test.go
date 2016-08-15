package main

import (
	"regexp"
	"testing"
)

func TestColumnRegexValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnRegex{
				column: Column{},
				Payload: ColumnRegexPayload{
					Regex: "foo",
					Limit: 10,
				},
			},
			inContext: Context{},
			outValue:  "foo",
		},
		{
			inColumn: ColumnRegex{
				column: Column{},
				Payload: ColumnRegexPayload{
					Regex: "f{1,1}",
					Limit: 10,
				},
			},
			inContext: Context{},
			outValue:  "f",
		},
		{
			inColumn: ColumnRegex{
				column: Column{},
				Payload: ColumnRegexPayload{
					Regex: "f{1,1}",
					Limit: 10,
				},
			},
			inContext: Context{},
			outValue:  "f",
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
