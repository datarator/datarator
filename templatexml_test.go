package main

import "testing"

func TestTemplateXMLGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateXML
		inContext  Context
		outValue   string
	}{
		{
			inTemplate: TemplateXML{
				Payload: TemplateXMLPayload{
					PrettyPrint: false,
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
			outValue: "<Hello/>\n<datarator/>\n<Hello/>\n<datarator/>\n",
		},
		{
			inTemplate: TemplateXML{
				Payload: TemplateXMLPayload{
					PrettyPrint: true,
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
			outValue: "<Hello/>\n<datarator/>\n<Hello/>\n<datarator/>\n",
		},
	}

	for _, test := range tests {
		actual, _ := test.inTemplate.Generate(&test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
