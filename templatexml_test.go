package main

import "testing"

func TestTemplateXMLGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateXML
		inContext  Chunk
		outValue   string
	}{
		{
			inTemplate: TemplateXML{
				payload: TemplateXMLPayload{
					PrettyPrint: false,
				},
				schema: Schema{
					count: 2,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "Hello",
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
									name: "datarator",
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
			inContext: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "<Hello/>\n<datarator/>\n<Hello/>\n<datarator/>\n",
		},
		{
			inTemplate: TemplateXML{
				payload: TemplateXMLPayload{
					PrettyPrint: true,
				},
				schema: Schema{
					count: 2,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "Hello",
									columns: []TypedColumn{
										ColumnConst{
											TypedColumnBase: TypedColumnBase{
												column: Column{
													name: "Nested",
													columns: []TypedColumn{
														ColumnConst{
															TypedColumnBase: TypedColumnBase{
																column: Column{
																	name: "NestedAttr",
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
									name: "datarator",
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
			inContext: Chunk{
				to:     2,
				values: make(map[string]string),
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
