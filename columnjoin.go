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
	Column  Column
	Payload ColumnJoinPayload
}

func (column ColumnJoin) Value(context *Context) (string, error) {
	var buffer bytes.Buffer

	for _, nestedColumn := range column.Column.TypedColumns {
		nestedVal, err := nestedColumn.Value(context)
		if err != nil {
			return "", nil
		}
		buffer.WriteString(nestedVal)
		buffer.WriteString(column.Payload.Separator)
	}

	return buffer.String(), nil
}
