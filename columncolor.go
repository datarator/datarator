package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_COLOR     = "color"
	COLUMN_COLOR_HEX = "color.hex"
)

type ColumnColorPayload struct {
	XmlType string `json:"xml"`
}

type ColumnColor struct {
	column  Column
	Payload ColumnColorPayload `json:"payload"`
}

func (column ColumnColor) Value(context *Context) (string, error) {
	return fake.Color(), nil
}

func (column ColumnColor) Column() Column {
	return column.column
}

func (column ColumnColor) XmlType() string {
	return column.Payload.XmlType
}

type ColumnColorHexPayload struct {
	XmlType string `json:"xml"`
	short   bool
}

type ColumnColorHex struct {
	column  Column
	Payload ColumnColorHexPayload `json:"payload"`
}

func (column ColumnColorHex) Value(context *Context) (string, error) {
	if column.Payload.short {
		return fake.HexColorShort(), nil
	}
	return fake.HexColor(), nil
}

func (column ColumnColorHex) Column() Column {
	return column.column
}

func (column ColumnColorHex) XmlType() string {
	return column.Payload.XmlType
}
