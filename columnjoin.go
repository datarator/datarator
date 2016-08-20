package main

import (
	"bytes"
)

const (
	COLUMN_JOIN = "join"
)

type ColumnJoinPayload struct {
	XmlType   string `json:"xml"`
	Separator string
}

type ColumnJoin struct {
	column  Column
	Payload ColumnJoinPayload
}

func (column ColumnJoin) Value(context *Context) (string, error) {
	var buffer bytes.Buffer

	for _, nestedColumn := range column.column.TypedColumns {
		nestedVal, err := nestedColumn.Value(context)
		if err != nil {
			return "", nil
		}
		buffer.WriteString(nestedVal)
		buffer.WriteString(column.Payload.Separator)
	}

	return buffer.String(), nil
}

func (column ColumnJoin) Column() Column {
	return column.column
}

func (column ColumnJoin) XmlType() string {
	return column.Payload.XmlType
}
