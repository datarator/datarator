package main

import "bytes"

type TemplateOptionsCSV struct {
	Header    bool
	Separator string
}

type TemplateCSV struct {
	Schema  Schema
	Options TemplateOptionsCSV `json:"options"`
}

func (templateCSV TemplateCSV) Generate(context Context) (string, error) {
	var buffer bytes.Buffer
	for context.CurrentRowIndex = context.FromIndex; context.CurrentRowIndex < context.ToIndex; context.CurrentRowIndex++ {
		for _, typedColumn := range templateCSV.Schema.TypedColumns {
			val, err := typedColumn.Value(context)
			if err != nil {
				return "", err
			}

			buffer.WriteString(val)
			buffer.WriteString(templateCSV.Options.Separator)
		}
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}
