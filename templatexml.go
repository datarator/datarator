package main

import "bytes"

const (
	TEMPLATE_XML = "xml"
)

type TemplateXMLPayload struct {
	PrettyPrint            bool `json:"pretty_print"`
	PrettyPrintSpacesUsed  bool `json:"pretty_print_spaces_used"`
	PrettyPrintSpacesCount int  `json:"pretty_print_spaces_count"`
}

type TemplateXML struct {
	Schema  Schema
	Payload TemplateXMLPayload `json:"payload"`
}

func (template TemplateXML) Generate(context *Context) (string, error) {
	var buffer bytes.Buffer
	for context.setCurrentIndex(context.FromIndex); context.getCurrentIndex() < context.ToIndex; context.incrementCurrentIndex() {
		for _, typedColumn := range template.Schema.TypedColumns {
			val, err := typedColumn.Value(context)
			if err != nil {
				return "", err
			}

			if template.Payload.PrettyPrint {
				buffer.WriteString(template.getIndent(context))
			}

			buffer.WriteByte('<')
			buffer.WriteString(val)
			buffer.WriteString("/>")
			buffer.WriteByte('\n')
		}
	}
	return buffer.String(), nil
}

func (template TemplateXML) getIndent(context *Context) string {
	var buffer bytes.Buffer
	for i := 0; i < context.CurrentNestingLevel; i++ {
		if template.Payload.PrettyPrintSpacesUsed {
			for j := 0; j < template.Payload.PrettyPrintSpacesCount; j++ {
				buffer.WriteByte(' ')
			}
		} else {
			buffer.WriteByte('\t')
		}
	}
	return buffer.String()
}
