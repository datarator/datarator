package main

import "strconv"

const (
	COLUMN_ROW_INDEX = "row_index"
)

type ColumnRowIndexPayload struct {
	XmlType string `json:"xml"`
}

type ColumnRowIndex struct {
	column  Column
	Payload ColumnRowIndexPayload `json:"payload"`
}

func (column ColumnRowIndex) Value(context *Context) (string, error) {
	return strconv.Itoa(context.getCurrentIndex()), nil
}

func (column ColumnRowIndex) Column() Column {
	return column.column
}

func (column ColumnRowIndex) XmlType() string {
	return column.Payload.XmlType
}
