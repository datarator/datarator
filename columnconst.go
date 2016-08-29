package main

const (
	columnConst = "const"
)

type ColumnConstPayload struct {
	Value string
}

type ColumnConst struct {
	TypedColumnBase
	payload ColumnConstPayload
}

func (column ColumnConst) Value(chunk *Chunk) (string, error) {
	return column.payload.Value, nil
}
