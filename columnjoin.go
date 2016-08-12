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

func (columnJoin ColumnJoin) Value(context *Context) (string, error) {
	var buffer bytes.Buffer

	for _, column := range columnJoin.Column.TypedColumns {
		nestedVal, err := column.Value(context)
		if err != nil {
			return "", nil
		}
		buffer.WriteString(nestedVal)
		buffer.WriteString(columnJoin.Payload.Separator)
	}

	return buffer.String(), nil
}
