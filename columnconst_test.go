package main

import "testing"

func TestColumnConstValue(t *testing.T) {
	var tests = []struct {
		inColumn ColumnConst
		outValue string
	}{
		{
			inColumn: ColumnConst{
				payload: ColumnConstPayload{
					Value: "foo",
				},
			},
			outValue: "foo",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(nil)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
