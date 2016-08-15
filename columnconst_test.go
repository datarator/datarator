package main

import "testing"

func TestColumnConstValue(t *testing.T) {
	var tests = []struct {
		inColumn  ColumnConst
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnConst{
				Payload: ColumnConstPayload{
					Value: "foo",
				},
				column: Column{},
			},
			inContext: Context{},
			outValue:  "foo",
		},
	}

	for _, test := range tests {
		actual, _ := test.inColumn.Value(&test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
