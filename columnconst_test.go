package main

import "testing"

var columnConstValueTests = []struct {
	inColumnConst ColumnConst
	inContext     Context
	outValue      string
}{
	{
		inColumnConst: ColumnConst{
			Options: ColumnOptionsConst{
				Value: "foo",
			},
			Column: Column{},
		},
		inContext: Context{},
		outValue:  "foo",
	},
}

func TestValue(t *testing.T) {
	for _, test := range columnConstValueTests {
		actual, _ := test.inColumnConst.Value(test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
