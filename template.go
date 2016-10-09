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
	Generate(chunk *Chunk) ([]byte, error)
	ContentType() string
}

type TemplateFactory struct {
}

func (templateFactory TemplateFactory) CreateTemplate(id string, jSONSchema *JSONSchema) (Template, error) {
	var err error

	columns, err := createColumns(jSONSchema.Columns, jSONSchema.Count)
	if err != nil {
		return nil, err
	}

	schema := Schema{
		document:   id,
		emptyValue: jSONSchema.EmptyValue,
		count:      jSONSchema.Count,
		columns:    columns,
	}

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

func generateValue(chunk *Chunk, emptyValue string, column TypedColumn) (string, error) {
	var val string
	emptyIndeces := column.Payload().EmptyIndeces()
	if emptyIndeces[chunk.index] {
		val = emptyValue
		chunk.values[column.Column().name] = val
		return val, nil
	} else {
		var err error
		val, err = column.Value(chunk)
		if err != nil {
			return "", err
		}
	}
	chunk.values[column.Column().name] = val
	return val, nil
}
