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

func (template TemplateCSV) Generate(chunk *Chunk) ([]byte, error) {
	var buffer bytes.Buffer

	if chunk.index == 0 {
		buffer.WriteString(template.getHeader(chunk))
	}

	for _, column := range template.schema.columns {
		val, err := generateValue(chunk, template.schema.emptyValue, column)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(val)
		buffer.WriteString(template.payload.Separator)
	}

	buffer.Truncate(buffer.Len() - len(template.payload.Separator))
	buffer.WriteByte('\n')
	return buffer.Bytes(), nil
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
