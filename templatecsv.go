package main

import "bytes"

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
	for context.CurrentRowIndex = context.FromIndex; context.CurrentRowIndex < context.ToIndex; context.CurrentRowIndex++ {
		for _, typedColumn := range templateCSV.Schema.TypedColumns {
			val, err := typedColumn.Value(context)
			if err != nil {
				return "", err
			}

			buffer.WriteString(val)
			buffer.WriteString(templateCSV.Payload.Separator)
		}
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}
