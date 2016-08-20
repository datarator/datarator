package main

import (
	"bytes"
	"fmt"
)

const (
	TEMPLATE_XML          = "xml"
	PAYLOAD_XML_ATTRIBUTE = "attribute"
	PAYLOAD_XML_CDATA     = "cdata"
	PAYLOAD_XML_COMMENT   = "comment"
	PAYLOAD_XML_ELEMENT   = "element"
)

var (
	errUnsupportedXmlType = "Column: %s has unsupported XML type: %s"
)

type TemplateXMLPayload struct {
	PrettyPrint            bool `json:"pretty_print"`
	PrettyPrintTabs        bool `json:"pretty_print_tabs"`
	PrettyPrintSpacesCount int  `json:"pretty_print_spaces_count"`
}

type TemplateXML struct {
	Schema  Schema
	Payload TemplateXMLPayload `json:"payload"`
}

func (template TemplateXML) Generate(context *Context) (string, error) {
	var buffer bytes.Buffer
	for context.setCurrentIndex(context.FromIndex); context.getCurrentIndex() < context.ToIndex; context.incrementCurrentIndex() {

		generated, err := template.generate(template.Schema.TypedColumns, context)
		if err != nil {
			return "", err
		}

		buffer.WriteString(generated)
	}
	return buffer.String(), nil
}

func (template TemplateXML) generate(typedColumns []TypedColumn, context *Context) (string, error) {
	var buffer bytes.Buffer
	if typedColumns != nil {
		for _, typedColumn := range typedColumns {
			xmlType := typedColumn.XmlType()
			if len(xmlType) == 0 {
				xmlType = PAYLOAD_XML_ELEMENT
			}

			switch xmlType {
			case PAYLOAD_XML_ELEMENT:
				buffer.WriteString(template.getIndent(context))
				buffer.WriteByte('<')
				buffer.WriteString(typedColumn.Column().Name)

				nestedColumns := typedColumn.Column().TypedColumns
				if nestedColumns != nil {
					// iterate nested attributes only
					for _, typedColumnNested := range nestedColumns {
						if typedColumnNested.XmlType() == PAYLOAD_XML_ATTRIBUTE {
							val, err := typedColumnNested.Value(context)
							if err != nil {
								return "", err
							}
							buffer.WriteByte(' ')
							buffer.WriteString(typedColumnNested.Column().Name)
							buffer.WriteString("=\"")
							buffer.WriteString(val)
							buffer.WriteByte('"')
						}
					}
					buffer.WriteByte('>')

					if err := context.nest(); err != nil {
						return "", err
					}

					// iterate nested nodes
					generated, err := template.generate(nestedColumns, context)
					if err != nil {
						return "", err
					}

					if err := context.unnest(); err != nil {
						return "", err
					}

					buffer.WriteString(generated)

					val, err := typedColumn.Value(context)
					if err != nil {
						return "", err
					}
					buffer.WriteString(val)

					buffer.WriteString("</")
					buffer.WriteString(typedColumn.Column().Name)
					buffer.WriteByte('>')
					buffer.WriteByte('\n')
				} else {
					buffer.WriteString("/>")
					buffer.WriteByte('\n')
				}

			case PAYLOAD_XML_ATTRIBUTE:
				// already covered in the default case => nothing to do here
			case PAYLOAD_XML_CDATA:
				val, err := typedColumn.Value(context)
				if err != nil {
					return "", err
				}

				buffer.WriteString("CDATA[\n")
				buffer.WriteString(val)
				buffer.WriteString("\n]\n")
			case PAYLOAD_XML_COMMENT:
				val, err := typedColumn.Value(context)
				if err != nil {
					return "", err
				}

				buffer.WriteString(template.getIndent(context))
				buffer.WriteString("<!--")
				buffer.WriteString(val)
				buffer.WriteString("-->")
			default:
				return "", fmt.Errorf(errUnsupportedXmlType, typedColumn.Column().Name, xmlType)
			}
		}

	}
	return buffer.String(), nil
}

func (template TemplateXML) getIndent(context *Context) string {
	if !template.Payload.PrettyPrint {
		return ""
	}

	var buffer bytes.Buffer
	if context.CurrentNestingLevel > 0 {
		buffer.WriteByte('\n')
	}
	for i := 0; i < context.CurrentNestingLevel; i++ {
		if !template.Payload.PrettyPrintTabs {
			spacesCount := template.Payload.PrettyPrintSpacesCount
			if spacesCount == 0 {
				spacesCount = 4
			}
			for j := 0; j < spacesCount; j++ {
				buffer.WriteByte(' ')
			}
		} else {
			buffer.WriteByte('\t')
		}
	}
	return buffer.String()
}
