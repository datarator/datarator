package main

import "testing"

func TestColumnCopyValue(t *testing.T) {
	var tests = []struct {
		inColumn ColumnCopy
		inChunk  Chunk
		outValue string
	}{
		{
			inColumn: ColumnCopy{
				payload: ColumnCopyPayload{
					From: "foo",
				},
			},
			inChunk: Chunk{
				values: map[string]string{
					"foo": "bar",
				},
			},
			outValue: "bar",
		},
		{
			inColumn: ColumnCopy{
				payload: ColumnCopyPayload{
					From: "../foo",
				},
			},
			inChunk: Chunk{
				parent: &Chunk{
					values: map[string]string{
						"foo": "bar",
					},
				},
			},
			outValue: "bar",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(&test.inChunk)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
