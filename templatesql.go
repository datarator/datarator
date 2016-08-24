package main

import "bytes"

const (
	TEMPLATE_SQL     = "sql"
	CONTENT_TYPE_SQL = "application/octet-stream"
)

type TemplateSQL struct {
	Schema Schema
}

func (template TemplateSQL) Generate(context *Context) (string, error) {
	insertPrefix := template.getInsertPrefix(context)

	var buffer bytes.Buffer
	for context.setCurrentIndex(context.FromIndex); context.getCurrentIndex() < context.ToIndex; context.incrementCurrentIndex() {
		buffer.WriteString(insertPrefix)
		for _, typedColumn := range template.Schema.TypedColumns {
			val, err := typedColumn.Value(context)
			if err != nil {
				return "", err
			}
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

func (template TemplateSQL) ContentType(context *Context) string {
	return CONTENT_TYPE_SQL
}

func (template TemplateSQL) getInsertPrefix(context *Context) string {
	var buffer bytes.Buffer

	buffer.WriteString("INSERT INTO ")
	buffer.WriteString(template.Schema.Document)
	buffer.WriteString(" ( ")
	for _, typedColumn := range template.Schema.TypedColumns {
		buffer.WriteString(typedColumn.Column().Name)
		buffer.WriteString(", ")
	}
	buffer.Truncate(buffer.Len() - 2)
	buffer.WriteString(" ) VALUES (")

	return buffer.String()
}
