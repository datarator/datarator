package main

import "strconv"

const (
	COLUMN_ROW_INDEX = "row_index"
)

type ColumnRowIndex struct {
	Column Column
}

func (column ColumnRowIndex) Value(context *Context) (string, error) {
	return strconv.Itoa(context.getCurrentIndex()), nil
}
