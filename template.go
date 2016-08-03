package main

const (
	CSV = "csv"
	SQL = "sql"
)

type TemplateFactory struct {
}

func (templateFactory TemplateFactory) CreateTemplate(id string, jSONSchema JSONSchema) (Template, error) {

	nestedColums, err := createColumns(jSONSchema.Columns)
	if err != nil {
		return nil, err
	}

	schema := Schema{
		Document:     id,
		EmptyValue:   jSONSchema.EmptyValue,
		Count:        jSONSchema.Count,
		TypedColumns: nestedColums,
	}
	// "EmptyIndeces":     countEmptyIndeces(jSONColumn.EmptyPercent),
	// "Locale":      retrieveLocale(jSONColumn),

	var template Template
	var errOpts error

	switch jSONSchema.Template {
	case CSV:
		options := TemplateOptionsCSV{}
		errOpts = loadOptions(jSONSchema.JSONOptions, &options)
		template = TemplateCSV{
			Schema:  schema,
			Options: options,
		}
	// case JOIN:
	// 	options := ColumnOptionsJoin{}
	// 	err := json.Unmarshal(jSONSchema.JSONOptions, &options)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	typedColumn = ColumnJoin{
	// 		"Options": options,
	// 	}
	default:
		return nil, nil // errors.New("Invalid schema Type: %v", schemaType)
	}

	if errOpts != nil {
		return nil, errOpts
	}

	return template, nil
}

// func countEmptyIndeces(EmptyPercent float32) ([]int, error) {
//     // TODO
//     return []int {1}, nil
// }

// func retrieveLocale(jSONColumn JSONColumn) (string, error) {
//     // TODO traverse all the way up to root to retrieve the locale
//     return "en", nil
// }
