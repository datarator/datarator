package main

import (
	"bytes"
)

const (
	columnJoin = "join"
)

type ColumnJoinPayload struct {
	Separator string
}

type ColumnJoin struct {
	TypedColumnBase
	payload ColumnJoinPayload
}

func (column ColumnJoin) Value(chunk *Chunk) (string, error) {
	var buffer bytes.Buffer

	for _, columns := range column.column.columns {
		nestedVal, err := columns.Value(chunk)
		if err != nil {
			return "", nil
		}
		buffer.WriteString(nestedVal)
		buffer.WriteString(column.payload.Separator)
	}

	return buffer.String(), nil
}
