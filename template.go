package main

import "fmt"

var (
	errUnsupportedTemplate = "Unsupported template: %s"
)

type TemplateFactory struct {
}

func (templateFactory TemplateFactory) CreateTemplate(id string, jSONSchema *JSONSchema) (Template, error) {

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
	case TEMPLATE_CSV:
		payload := TemplateCSVPayload{}
		errOpts = loadPayload(jSONSchema.JSONPayload, &payload)
		template = TemplateCSV{
			Schema:  schema,
			Payload: payload,
		}
	case TEMPLATE_XML:
		payload := TemplateXMLPayload{}
		errOpts = loadPayload(jSONSchema.JSONPayload, &payload)
		template = TemplateXML{
			Schema:  schema,
			Payload: payload,
		}
	default:
		return nil, fmt.Errorf(errUnsupportedTemplate, jSONSchema.Template)
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
