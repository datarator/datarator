package main

import "testing"

func TestTemplateSQLGenerate(t *testing.T) {
	var tests = []struct {
		inTemplate TemplateSQL
		inContext  Context
		outValue   string
	}{
		{
			inTemplate: TemplateSQL{
				Schema: Schema{
					Count: 2,
					TypedColumns: []TypedColumn{
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									Name: "col1",
								},
							},
							payload: ColumnConstPayload{
								Value: "Hello",
							},
						},
						ColumnConst{
							TypedColumnBase: TypedColumnBase{
								column: Column{
									Name: "col2",
								},
							},
							payload: ColumnConstPayload{
								Value: "datarator",
							},
						},
					},
					Document: "foo",
				},
			},
			inContext: Context{
				CurrentIndex: []int{0},
				ToIndex:      2,
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
