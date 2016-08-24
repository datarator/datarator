package main

import "bytes"

const (
	TEMPLATE_CSV     = "csv"
	CONTENT_TYPE_CSV = "text/csv; charset=UTF-8"
)

type TemplateCSVPayload struct {
	Header    bool
	Separator string
}

type TemplateCSV struct {
	Schema  Schema
	Payload TemplateCSVPayload `json:"payload"`
}

func (template TemplateCSV) Generate(context *Context) (string, error) {
	var buffer bytes.Buffer

	buffer.WriteString(template.getHeader(context))

	for context.setCurrentIndex(context.FromIndex); context.getCurrentIndex() < context.ToIndex; context.incrementCurrentIndex() {
		for _, typedColumn := range template.Schema.TypedColumns {
			val, err := typedColumn.Value(context)
			if err != nil {
				return "", err
			}

			buffer.WriteString(val)
			buffer.WriteString(template.Payload.Separator)
		}
		buffer.Truncate(buffer.Len() - len(template.Payload.Separator))
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}

func (template TemplateCSV) ContentType() string {
	return CONTENT_TYPE_CSV
}

func (template TemplateCSV) getHeader(context *Context) string {
	if !template.Payload.Header || context.getCurrentIndex() != 0 {
		return ""
	}

	var buffer bytes.Buffer
	for _, typedColumn := range template.Schema.TypedColumns {
		buffer.WriteString(typedColumn.Column().Name)
		buffer.WriteString(template.Payload.Separator)
	}
	buffer.Truncate(buffer.Len() - len(template.Payload.Separator))
	buffer.WriteByte('\n')

	return buffer.String()
}
