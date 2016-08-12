package main

import "strconv"

const (
	COLUMN_ROW_INDEX = "row_index"
)

// type ColumnRowIndexPayload struct {
// 	// Increment int
// 	// Decrement int
// }

type ColumnRowIndex struct {
	Column Column
	// Payload ColumnRowIndexPayload
}

func (columnRowIndex ColumnRowIndex) Value(context *Context) (string, error) {
	//	return context.RowIndex + columnRowIndex.Payload.Increment - columnRowIndex.Payload.Decrement, nil
	return strconv.Itoa(context.CurrentRowIndex), nil
}
