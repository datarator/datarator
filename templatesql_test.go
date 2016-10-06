package main

import (
	"bytes"
	"testing"
)

func TestTemplateSQLGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateSQL
		inChunk    Chunk
		outValue   string
	}{
		{
			inTemplate: TemplateSQL{
				schema: Schema{
					count: 2,
					columns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "col1",
								},
							},
							payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									name: "col2",
								},
							},
							payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
					document: "foo",
				},
			},
			inChunk: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "INSERT INTO foo ( col1, col2 ) VALUES ( 'Hello', 'datarator' );\nINSERT INTO foo ( col1, col2 ) VALUES ( 'Hello', 'datarator' );\n",
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
