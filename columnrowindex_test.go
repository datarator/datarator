package main

import "testing"

func TestColumnRowIndexValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnRowIndex{},
			inContext: Context{
				CurrentIndex: []int{0},
			},
			outValue: "0",
		},
		{
			inColumn: ColumnRowIndex{},
			inContext: Context{
				CurrentIndex: []int{100},
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
