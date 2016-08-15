package main

import "bytes"

const (
	TEMPLATE_CSV = "csv"
)

type TemplateCSVPayload struct {
	Header    bool
	Separator string
}

type TemplateCSV struct {
	Schema  Schema
	Payload TemplateCSVPayload `json:"payload"`
}

func (templateCSV TemplateCSV) Generate(context *Context) (string, error) {
	var buffer bytes.Buffer
	for context.setCurrentIndex(context.FromIndex); context.getCurrentIndex() < context.ToIndex; context.incrementCurrentIndex() {
		for _, typedColumn := range templateCSV.Schema.TypedColumns {
			val, err := typedColumn.Value(context)
			if err != nil {
				return "", err
			}

			buffer.WriteString(val)
			buffer.WriteString(templateCSV.Payload.Separator)
		}
		buffer.Truncate(buffer.Len() - len(templateCSV.Payload.Separator))
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}
