package main

import "testing"

func TestColumnRowIndexValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		inChunk  Chunk
		outValue string
	}{
		{
			inColumn: ColumnRowIndex{},
			inChunk:  Chunk{},
			outValue: "0",
		},
		{
			inColumn: ColumnRowIndex{},
			inChunk: Chunk{
				index: 100,
			},
			outValue: "100",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(&test.inChunk)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
