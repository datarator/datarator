package main

import (
	"bytes"
	"testing"
)

func TestTemplateXMLGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateXML
		inChunk    Chunk
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
			inChunk: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "<Hello/>\n<datarator/>\n<Hello/>\n<datarator/>\n",
		},
		{
			inTemplate: TemplateXML{
				payload: TemplateXMLPayload{
					PrettyPrint: false,
				},
				schema: Schema{
					count: 1,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "Hello",
									columns: []TypedColumn{
										ColumnConst{
											TypedColumnBase: TypedColumnBase{
												column: Column{
													name: "datarator",
												},
												payload: TypedColumnBasePayload{
													Xml: "cdata",
												},
											},
											payload: ColumnConstPayload{
												Value: "datarator",
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
					},
				},
			},
			inChunk: Chunk{
				to:     1,
				values: make(map[string]string),
			},
			outValue: "<Hello>\n<![CDATA[\ndatarator\n]]>\n</Hello>\n",
		},
		{
			inTemplate: TemplateXML{
				payload: TemplateXMLPayload{
					PrettyPrint: false,
				},
				schema: Schema{
					count: 1,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "Hello",
									columns: []TypedColumn{
										ColumnConst{
											TypedColumnBase: TypedColumnBase{
												column: Column{
													name: "datarator",
												},
												payload: TypedColumnBasePayload{
													Xml: "comment",
												},
											},
											payload: ColumnConstPayload{
												Value: "datarator",
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
					},
				},
			},
			inChunk: Chunk{
				to:     1,
				values: make(map[string]string),
			},
			outValue: "<Hello>\n<!--datarator-->\n</Hello>\n",
		},
		{
			inTemplate: TemplateXML{
				payload: TemplateXMLPayload{
					PrettyPrint: false,
				},
				schema: Schema{
					count: 1,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "Hello",
									columns: []TypedColumn{
										ColumnConst{
											TypedColumnBase: TypedColumnBase{
												column: Column{
													name: "datarator",
												},
												payload: TypedColumnBasePayload{
													Xml: "value",
												},
											},
											payload: ColumnConstPayload{
												Value: "datarator",
											},
										},
									},
								},
								payload: TypedColumnBasePayload{},
							},
							payload: ColumnConstPayload{
								Value: "ignored",
							},
						},
					},
				},
			},
			inChunk: Chunk{
				to:     1,
				values: make(map[string]string),
			},
			outValue: "<Hello>datarator</Hello>\n",
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
														ColumnConst{
															TypedColumnBase: TypedColumnBase{
																column: Column{
																	name: "NestedVal",
																},
																payload: TypedColumnBasePayload{
																	Xml: "value",
																},
															},
															payload: ColumnConstPayload{
																Value: "NestedVal",
															},
														},
													},
												},
												payload: TypedColumnBasePayload{},
											},
											payload: ColumnConstPayload{
												Value: "ignored",
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
			inChunk: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "<Hello>\n    <Nested NestedAttr=\"NestedAttrVal\">NestedVal</Nested>\n</Hello>\n<datarator/>\n<Hello>\n    <Nested NestedAttr=\"NestedAttrVal\">NestedVal</Nested>\n</Hello>\n<datarator/>\n",
		},
	}

	for _, test := range tests {
		var actual bytes.Buffer
		for test.inChunk.index = test.inChunk.from; test.inChunk.index < test.inChunk.to; test.inChunk.index++ {
			bytes, _ := test.inTemplate.Generate(&test.inChunk)
			actual.Write(bytes)
		}
		if actual.String() != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
