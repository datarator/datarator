package main

import (
	"regexp"
	"testing"
)

func TestColumnColorValue(t *testing.T) {
	var tests = []struct {
		inColumn ColumnColor
		outValue string
	}{
		{
			inColumn: ColumnColor{},
			outValue: "^[A-Z][a-z]+$",
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

func TestColumnColorHexValue(t *testing.T) {
	var tests = []struct {
		inColumn ColumnColorHex
		outValue string
	}{
		{
			inColumn: ColumnColorHex{},
			outValue: "^[0-9a-z]{6}$",
		},
		{
			inColumn: ColumnColorHex{
				payload: ColumnColorHexPayload{
					Short: true,
				},
			},
			outValue: "^[0-9a-z]{3}$",
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
