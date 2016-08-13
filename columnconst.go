package main

const (
	COLUMN_CONST = "const"
)

type ColumnConstPayload struct {
	Value string
}

type ColumnConst struct {
	Column  Column
	Payload ColumnConstPayload `json:"payload"`
}

func (column ColumnConst) Value(context *Context) (string, error) {
	return column.Payload.Value, nil
}
