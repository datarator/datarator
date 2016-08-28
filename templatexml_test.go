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
							TypedColumnBase: TypedColumnBase{
								column: Column{
									Name: "Hello",
								},
								payload: TypedColumnBasePayload{},
							},
							payload: ColumnConstPayload{
								Value: "",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									Name: "datarator",
								},
								payload: TypedColumnBasePayload{},
							},
							payload: ColumnConstPayload{
								Value: "",
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
							TypedColumnBase: TypedColumnBase{
								column: Column{
									Name: "Hello",
									TypedColumns: []TypedColumn{
										ColumnConst{
											TypedColumnBase: TypedColumnBase{
												column: Column{
													Name: "Nested",
													TypedColumns: []TypedColumn{
														ColumnConst{
															TypedColumnBase: TypedColumnBase{
																column: Column{
																	Name: "NestedAttr",
																},
																payload: TypedColumnBasePayload{
																	Xml: "attribute",
																},
															},
															payload: ColumnConstPayload{
																Value: "NestedAttrVal",
															},
														},
													},
												},
												payload: TypedColumnBasePayload{},
											},
											payload: ColumnConstPayload{
												Value: "Nestedval",
											},
										},
									},
								},
								payload: TypedColumnBasePayload{},
							},
							payload: ColumnConstPayload{
								Value: "",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									Name: "datarator",
								},
								payload: TypedColumnBasePayload{},
							},
							payload: ColumnConstPayload{
								Value: "",
							},
						},
					},
				},
			},
			inContext: Context{
				CurrentIndex: []int{0},
				ToIndex:      2,
			},
			outValue: "<Hello>\n    <Nested NestedAttr=\"NestedAttrVal\">Nestedval</Nested>\n</Hello>\n<datarator/>\n<Hello>\n    <Nested NestedAttr=\"NestedAttrVal\">Nestedval</Nested>\n</Hello>\n<datarator/>\n",
		},
	}

	for _, test := range tests {
		actual, _ := test.inTemplate.Generate(&test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
