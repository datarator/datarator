package main

import (
	"regexp"
	"testing"
)

func TestColumnColorValue(t *testing.T) {
	var tests = []struct {
		inColumn  ColumnColor
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnColor{
				Payload: ColumnColorPayload{},
				column:  Column{},
			},
			inContext: Context{},
			outValue:  "^[A-Z][a-z]+$",
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

func TestColumnColorHexValue(t *testing.T) {
	var tests = []struct {
		inColumn  ColumnColorHex
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnColorHex{
				Payload: ColumnColorHexPayload{},
				column:  Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9a-z]{6}$",
		},
		{
			inColumn: ColumnColorHex{
				Payload: ColumnColorHexPayload{
					short: true,
				},
				column: Column{},
			},
			inContext: Context{},
			outValue:  "^[0-9a-z]{3}$",
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
