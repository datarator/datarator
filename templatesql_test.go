package main

import "testing"

func TestTemplateSQLGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateSQL
		inContext  Chunk
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
			inContext: Chunk{
				to:     2,
				values: make(map[string]string),
			},
			outValue: "INSERT INTO foo ( col1, col2 ) VALUES ( 'Hello', 'datarator' );\nINSERT INTO foo ( col1, col2 ) VALUES ( 'Hello', 'datarator' );\n",
		},
	}

	for _, test := range tests {
		actual, _ := test.inTemplate.Generate(&test.inContext)
		if actual != test.outValue {
			t.Fatalf("Expected: %v\nActual: %v", test.outValue, actual)
		}
	}
}
