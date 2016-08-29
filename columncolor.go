package main

import (
	"github.com/syscrusher/fake"
)

const (
	columnColor    = "color"
	columnColorHex = "color.hex"
)

type ColumnColor struct {
	TypedColumnBase
}

func (column ColumnColor) Value(chunk *Chunk) (string, error) {
	return fake.Color(), nil
}

type ColumnColorHexPayload struct {
	Short bool
}

type ColumnColorHex struct {
	TypedColumnBase
	payload ColumnColorHexPayload
}

func (column ColumnColorHex) Value(chunk *Chunk) (string, error) {
	if column.payload.Short {
		return fake.HexColorShort(), nil
	}
	return fake.HexColor(), nil
}
