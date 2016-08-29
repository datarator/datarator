package main

import "bytes"

const (
	templateCSV    = "csv"
	contentTypeCSV = "text/csv; charset=UTF-8"
)

type TemplateCSVPayload struct {
	Header    bool
	Separator string
}

type TemplateCSV struct {
	schema  Schema
	payload TemplateCSVPayload
}

func (template TemplateCSV) Generate(chunk *Chunk) (string, error) {
	var buffer bytes.Buffer

	buffer.WriteString(template.getHeader(chunk))

	for chunk.index = chunk.from; chunk.index < chunk.to; chunk.index++ {
		for _, column := range template.schema.columns {
			val, err := column.Value(chunk)
			if err != nil {
				return "", err
			}
			chunk.values[column.Column().name] = val

			buffer.WriteString(val)
			buffer.WriteString(template.payload.Separator)
		}
		buffer.Truncate(buffer.Len() - len(template.payload.Separator))
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}

func (template TemplateCSV) ContentType() string {
	return contentTypeCSV
}

func (template TemplateCSV) getHeader(chunk *Chunk) string {
	if !template.payload.Header || chunk.index != 0 {
		return ""
	}

	var buffer bytes.Buffer
	for _, column := range template.schema.columns {
		buffer.WriteString(column.Column().name)
		buffer.WriteString(template.payload.Separator)
	}
	buffer.Truncate(buffer.Len() - len(template.payload.Separator))
	buffer.WriteByte('\n')

	return buffer.String()
}
