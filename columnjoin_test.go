package main

import (
	"regexp"
	"testing"
)

func TestColumnJoinValue(t *testing.T) {
	var tests = []struct {
		inColumn TypedColumn
		outValue string
	}{
		{
			inColumn: ColumnJoin{
				TypedColumnBase: TypedColumnBase{
					column: Column{
						columns: []TypedColumn{
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
			outValue: "Hello datarator",
		},
		{
			inColumn: ColumnJoin{
				TypedColumnBase: TypedColumnBase{
					column: Column{
						columns: []TypedColumn{
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
			outValue: "Hellodatarator",
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
