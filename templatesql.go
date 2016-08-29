package main

import "bytes"

const (
	templateSQL    = "sql"
	contentTypeSQL = "application/octet-stream"
)

type TemplateSQL struct {
	schema Schema
}

func (template TemplateSQL) Generate(chunk *Chunk) (string, error) {
	insertPrefix := template.getInsertPrefix(chunk)

	var buffer bytes.Buffer
	for chunk.index = chunk.from; chunk.index < chunk.to; chunk.index++ {
		buffer.WriteString(insertPrefix)
		for _, column := range template.schema.columns {
			val, err := column.Value(chunk)
			if err != nil {
				return "", err
			}
			chunk.values[column.Column().name] = val

			buffer.WriteString(" '")
			buffer.WriteString(val)
			buffer.WriteString("',")

		}
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteString(" );")
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}

func (template TemplateSQL) ContentType() string {
	return contentTypeSQL
}

func (template TemplateSQL) getInsertPrefix(chunk *Chunk) string {
	var buffer bytes.Buffer

	buffer.WriteString("INSERT INTO ")
	buffer.WriteString(template.schema.document)
	buffer.WriteString(" ( ")
	for _, column := range template.schema.columns {
		buffer.WriteString(column.Column().name)
		buffer.WriteString(", ")
	}
	buffer.Truncate(buffer.Len() - 2)
	buffer.WriteString(" ) VALUES (")

	return buffer.String()
}
