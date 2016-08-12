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

func (columnConst ColumnConst) Value(context *Context) (string, error) {
	return columnConst.Payload.Value, nil
}
