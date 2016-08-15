package main

import "strconv"

const (
	COLUMN_ROW_INDEX = "row_index"
)

type ColumnRowIndex struct {
	column Column
}

func (column ColumnRowIndex) Value(context *Context) (string, error) {
	return strconv.Itoa(context.getCurrentIndex()), nil
}

func (column ColumnRowIndex) Column() Column {
	return column.column
}
