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
				TypedColumnBase: TypedColumnBase{
					column: Column{
						TypedColumns: []TypedColumn{
							ColumnConst{
								payload: ColumnConstPayload{
									Value: "Hello",
								},
							},
							ColumnConst{
								payload: ColumnConstPayload{
									Value: "datarator",
								},
							},
						},
					},
				},
				payload: ColumnJoinPayload{
					Separator: " ",
				},
			},
			inContext: Context{},
			outValue:  "Hello datarator",
		},
		{
			inColumn: ColumnJoin{
				TypedColumnBase: TypedColumnBase{
					column: Column{
						TypedColumns: []TypedColumn{
							ColumnConst{
								payload: ColumnConstPayload{
									Value: "Hello",
								},
							},
							ColumnConst{
								payload: ColumnConstPayload{
									Value: "datarator",
								},
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
