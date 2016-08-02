package main

import (
	"bytes"
)

type ColumnOptionsJoin struct {
	Separator string
}

type ColumnJoin struct {
	Column  Column
	Options ColumnOptionsJoin
}

func (columnJoin ColumnJoin) Value(context Context) (string, error) {
	var buffer bytes.Buffer

	for _, column := range columnJoin.Column.TypedColumns {
		nestedVal, err := column.Value(context)
		if err != nil {
			return "", nil
		}
		buffer.WriteString(nestedVal)
		buffer.WriteString(columnJoin.Options.Separator)
	}

	return buffer.String(), nil
}
