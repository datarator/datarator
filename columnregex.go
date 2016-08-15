package main

import (
	"github.com/lucasjones/reggen"
)

const (
	COLUMN_REGEX = "regex"
)

type ColumnRegexPayload struct {
	Regex string
	Limit int // max number of times *,+,repeat should repeat
}

type ColumnRegex struct {
	column  Column
	Payload ColumnRegexPayload `json:"payload"`
}

func (column ColumnRegex) Value(context *Context) (string, error) {
	return reggen.Generate(column.Payload.Regex, column.Payload.Limit)
}

func (column ColumnRegex) Column() Column {
	return column.column
}
