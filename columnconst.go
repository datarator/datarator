package main

const (
	COLUMN_CONST = "const"
)

type ColumnConstPayload struct {
	Value string
}

type ColumnConst struct {
	column  Column
	Payload ColumnConstPayload `json:"payload"`
}

func (column ColumnConst) Value(context *Context) (string, error) {
	return column.Payload.Value, nil
}

func (column ColumnConst) Column() Column {
	return column.column
}
