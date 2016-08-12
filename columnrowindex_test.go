package main

import "testing"

func TestColumnRowIndexValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnRowIndex{
				// Payload: ColumnRowIndexPayload{},
				Column: Column{},
			},
			inContext: Context{
				CurrentRowIndex: 0,
			},
			outValue: "0",
		},
		{
			inColumn: ColumnRowIndex{
				// Payload: ColumnRowIndexPayload{},
				Column: Column{},
			},
			inContext: Context{
				CurrentRowIndex: 100,
			},
			outValue: "100",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(&test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
