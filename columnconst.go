package main

const (
	COLUMN_CONST = "const"
)

type ColumnOptionsConst struct {
	Value string
}

type ColumnConst struct {
	Column  Column
	Options ColumnOptionsConst `json:"options"`
}

func (columnConst ColumnConst) Value(context Context) (string, error) {
	return columnConst.Options.Value, nil
}
