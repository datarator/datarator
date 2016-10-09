package main

import (
	"bytes"
	"testing"
)

func TestTemplateCSVGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateCSV
		inChunk    Chunk
		outValue   string
	}{
		{
			inTemplate: TemplateCSV{
				payload: TemplateCSVPayload{},
				schema: Schema{
					count: 2,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								payload: TypedColumnBasePayload{
									emptyIndeces: make(map[int]bool),
								},
							},
							payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								payload: TypedColumnBasePayload{
									emptyIndeces: make(map[int]bool),
								},
							},
							payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
				},
			},
			inChunk: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "Hellodatarator\nHellodatarator\n",
		},
		{
			inTemplate: TemplateCSV{
				payload: TemplateCSVPayload{
					Separator: ",",
				},
				schema: Schema{
					count: 2,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								payload: TypedColumnBasePayload{
									emptyIndeces: make(map[int]bool),
								},
							},
							payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								payload: TypedColumnBasePayload{
									emptyIndeces: make(map[int]bool),
								},
							},
							payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
				},
			},
			inChunk: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "Hello,datarator\nHello,datarator\n",
		},
		{
			inTemplate: TemplateCSV{
				payload: TemplateCSVPayload{
					Header:    true,
					Separator: ",",
				},
				schema: Schema{
					count: 2,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "foo",
								},
								payload: TypedColumnBasePayload{
									emptyIndeces: make(map[int]bool),
								},
							},
							payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "bar",
								}, payload: TypedColumnBasePayload{
									emptyIndeces: make(map[int]bool),
								},
							},
							payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
				},
			},
			inChunk: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "foo,bar\nHello,datarator\nHello,datarator\n",
		},
	}

	for _, test := range tests {
		var actual bytes.Buffer
		for test.inChunk.index = test.inChunk.from; test.inChunk.index < test.inChunk.to; test.inChunk.index++ {
			bytes, _ := test.inTemplate.Generate(&test.inChunk)
			actual.Write(bytes)
		}
		if actual.String() != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual.String())
		}
	}
}
