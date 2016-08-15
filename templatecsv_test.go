package main

import "testing"

func TestTemplateCSVGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateCSV
		inContext  Context
		outValue   string
	}{
		{
			inTemplate: TemplateCSV{
				Payload: TemplateCSVPayload{},
				Schema: Schema{
					Count: 2,
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
			inContext: Context{
				CurrentIndex: []int{0},
				ToIndex:      2,
			},
			outValue: "Hellodatarator\nHellodatarator\n",
		},
		{
			inTemplate: TemplateCSV{
				Payload: TemplateCSVPayload{
					Separator: ",",
				},
				Schema: Schema{
					Count: 2,
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
			inContext: Context{
				CurrentIndex: []int{0},
				ToIndex:      2,
			},
			outValue: "Hello,datarator\nHello,datarator\n",
		},
		{
			inTemplate: TemplateCSV{
				Payload: TemplateCSVPayload{
					Header:    true,
					Separator: ",",
				},
				Schema: Schema{
					Count: 2,
					TypedColumns: []TypedColumn{
						ColumnConst{
							column: Column{
								Name: "foo",
							},
							Payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							column: Column{
								Name: "bar",
							},
							Payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
				},
			},
			inContext: Context{
				CurrentIndex: []int{0},
				ToIndex:      2,
			},
			outValue: "foo,bar\nHello,datarator\nHello,datarator\n",
		},
	}

	for _, test := range tests {
		actual, _ := test.inTemplate.Generate(&test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
