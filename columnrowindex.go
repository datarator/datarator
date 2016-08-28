package main

import "strconv"

const (
	columnRowIndex = "row_index"
)

type ColumnRowIndex struct {
	TypedColumnBase
}

func (column ColumnRowIndex) Value(context *Context) (string, error) {
	return strconv.Itoa(context.getCurrentIndex()), nil
}
