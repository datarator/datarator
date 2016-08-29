package main

import "fmt"

const (
	errUnsupportedTemplate = "Unsupported template: %s"
)

type Schema struct {
	document   string
	emptyValue string
	count      int
	columns    []TypedColumn
	locale     string
}

type Template interface {
	Generate(chunk *Chunk) (string, error)
	ContentType() string
}

type TemplateFactory struct {
}

func (templateFactory TemplateFactory) CreateTemplate(id string, jSONSchema *JSONSchema) (Template, error) {
	var err error

	columns, err := createColumns(jSONSchema.Columns)
	if err != nil {
		return nil, err
	}

	schema := Schema{
		document:   id,
		emptyValue: jSONSchema.EmptyValue,
		count:      jSONSchema.Count,
		columns:    columns,
	}
	// "EmptyIndeces":     countEmptyIndeces(jSONColumn.EmptyPercent),
	// "Locale":      retrieveLocale(jSONColumn),

	var template Template

	switch jSONSchema.Template {
	case templateCSV:
		payload := TemplateCSVPayload{}
		err = loadPayload(jSONSchema.JSONPayload, &payload)
		template = TemplateCSV{
			schema:  schema,
			payload: payload,
		}
	case templateSQL:
		template = TemplateSQL{
			schema: schema,
		}
	case templateXML:
		payload := TemplateXMLPayload{}
		err = loadPayload(jSONSchema.JSONPayload, &payload)
		template = TemplateXML{
			schema:  schema,
			payload: payload,
		}
	default:
		err = fmt.Errorf(errUnsupportedTemplate, jSONSchema.Template)
	}

	if err != nil {
		return nil, err
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
