package main

import "bytes"

const (
	templateSQL    = "sql"
	contentTypeSQL = "text/sql"
)

type TemplateSQL struct {
	schema     Schema
	linePrefix []byte
}

func (template TemplateSQL) Generate(chunk *Chunk) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.Write(template.getLinePrefix(chunk))
	for _, column := range template.schema.columns {
		val, err := column.Value(chunk)
		if err != nil {
			return nil, err
		}
		chunk.values[column.Column().name] = val
		buffer.WriteByte('\'')
		buffer.WriteString(val)
		buffer.WriteByte('\'')
		buffer.WriteByte(',')
	}
	buffer.Truncate(buffer.Len() - 1)
	buffer.WriteString(");")
	buffer.WriteByte('\n')
	return buffer.Bytes(), nil
}

func (template TemplateSQL) ContentType() string {
	return contentTypeSQL
}

func (template TemplateSQL) getLinePrefix(chunk *Chunk) []byte {
	if template.linePrefix == nil {
		var buffer bytes.Buffer
		buffer.WriteString("INSERT INTO ")
		buffer.WriteString(template.schema.document)
		buffer.WriteString(" (")
		for _, column := range template.schema.columns {
			buffer.WriteString(column.Column().name)
			buffer.WriteByte(',')
		}
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteString(") VALUES (")
		template.linePrefix = buffer.Bytes()
	}

	return template.linePrefix
}
