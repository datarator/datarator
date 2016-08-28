package main

import (
	"github.com/lucasjones/reggen"
)

const (
	columnRegex = "regex"
)

type ColumnRegexPayload struct {
	Regex string
	Limit int // max number of times *,+,repeat should repeat
}

type ColumnRegex struct {
	TypedColumnBase
	payload ColumnRegexPayload
}

func (column ColumnRegex) Value(context *Context) (string, error) {
	return reggen.Generate(column.payload.Regex, column.payload.Limit)
}
