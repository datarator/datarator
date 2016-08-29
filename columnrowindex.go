package main

import "strconv"

const (
	columnRowIndex = "row_index"
)

type ColumnRowIndex struct {
	TypedColumnBase
}

func (column ColumnRowIndex) Value(chunk *Chunk) (string, error) {
	return strconv.Itoa(chunk.index), nil
}
