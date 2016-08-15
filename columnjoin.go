package main

import (
	"bytes"
)

const (
	COLUMN_JOIN = "join"
)

type ColumnJoinPayload struct {
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
