package main

import (
	"regexp"
	"testing"
)

func TestColumnJoinValue(t *testing.T) {
	var tests = []struct {
		inColumn  TypedColumn
		inContext Context
		outValue  string
	}{
		{
			inColumn: ColumnJoin{
				column: Column{
					TypedColumns: []TypedColumn{
						ColumnConst{
							Payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							Payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
				},
				Payload: ColumnJoinPayload{
					Separator: " ",
				},
			},
			inContext: Context{},
			outValue:  "Hello datarator",
		},
		{
			inColumn: ColumnJoin{
				column: Column{
					TypedColumns: []TypedColumn{
						ColumnConst{
							Payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							Payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
				},
			},
			inContext: Context{},
			outValue:  "Hellodatarator",
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
