package main

import (
	"bytes"
	"fmt"
)

const (
	templateXML         = "xml"
	payloadXMLAttribute = "attribute"
	payloadXMLCdata     = "cdata"
	payloadXMLComment   = "comment"
	payloadXMLElement   = "element"
	contentTypeXML      = "text/xml; charset=UTF-8"

	errUnsupportedXMLType = "Column: %s has unsupported XML type: %s"
)

type TemplateXMLPayload struct {
	PrettyPrint            bool `json:"pretty_print"`
	PrettyPrintTabs        bool `json:"pretty_print_tabs"`
	PrettyPrintSpacesCount int  `json:"pretty_print_spaces_count"`
}

type TemplateXML struct {
	schema  Schema
	payload TemplateXMLPayload
}

func (template TemplateXML) Generate(chunk *Chunk) (string, error) {
	var buffer bytes.Buffer
	for chunk.index = chunk.from; chunk.index < chunk.to; chunk.index++ {

		generated, err := template.generate(template.schema.columns, chunk)
		if err != nil {
			return "", err
		}

		buffer.WriteString(generated)
	}
	return buffer.String(), nil
}

func (template TemplateXML) ContentType() string {
	return contentTypeXML
}

func (template TemplateXML) generate(columns []TypedColumn, chunk *Chunk) (string, error) {
	var buffer bytes.Buffer
	if columns != nil {
		for _, column := range columns {
			xmlType := column.Payload().XmlType()
			if len(xmlType) == 0 {
				xmlType = payloadXMLElement
			}

			switch xmlType {
			case payloadXMLElement:
				buffer.WriteString(template.getIndent(chunk))
				buffer.WriteByte('<')
				buffer.WriteString(column.Column().name)

				nestedColumns := column.Column().columns
				if nestedColumns != nil {
					// iterate nested attributes only
					for _, nestedColumn := range nestedColumns {
						if nestedColumn.Payload().XmlType() == payloadXMLAttribute {
							val, err := nestedColumn.Value(chunk)
							if err != nil {
								return "", err
							}
							chunk.values[column.Column().name] = val

							buffer.WriteByte(' ')
							buffer.WriteString(nestedColumn.Column().name)
							buffer.WriteString("=\"")
							buffer.WriteString(val)
							buffer.WriteByte('"')
						}
					}
					buffer.WriteByte('>')

					// iterate nested nodes
					generated, err := template.generate(nestedColumns, &Chunk{
						from:   chunk.from,
						to:     chunk.to,
						values: make(map[string]string),
						parent: chunk,
					})
					if err != nil {
						return "", err
					}

					buffer.WriteString(generated)

					val, err := column.Value(chunk)
					if err != nil {
						return "", err
					}
					chunk.values[column.Column().name] = val

					buffer.WriteString(val)

					buffer.WriteString("</")
					buffer.WriteString(column.Column().name)
					buffer.WriteByte('>')
					buffer.WriteByte('\n')
				} else {
					buffer.WriteString("/>")
					buffer.WriteByte('\n')
				}

			case payloadXMLAttribute:
				// already covered in the default case => nothing to do here
			case payloadXMLCdata:
				val, err := column.Value(chunk)
				if err != nil {
					return "", err
				}
				chunk.values[column.Column().name] = val

				buffer.WriteString("CDATA[\n")
				buffer.WriteString(val)
				buffer.WriteString("\n]\n")
			case payloadXMLComment:
				val, err := column.Value(chunk)
				if err != nil {
					return "", err
				}
				chunk.values[column.Column().name] = val

				buffer.WriteString(template.getIndent(chunk))
				buffer.WriteString("<!--")
				buffer.WriteString(val)
				buffer.WriteString("-->")
			default:
				return "", fmt.Errorf(errUnsupportedXMLType, column.Column().name, xmlType)
			}
		}

	}
	return buffer.String(), nil
}

func (template TemplateXML) getIndent(chunk *Chunk) string {
	if !template.payload.PrettyPrint {
		return ""
	}

	var buffer bytes.Buffer
	if chunk.parent != nil {
		buffer.WriteByte('\n')
	}

	for currentChunk := chunk; currentChunk.parent != nil; currentChunk = currentChunk.parent {
		if !template.payload.PrettyPrintTabs {
			spacesCount := template.payload.PrettyPrintSpacesCount
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
