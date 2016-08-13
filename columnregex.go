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
	Column  Column
	Payload ColumnRegexPayload `json:"payload"`
}

func (columnConst ColumnRegex) Value(context *Context) (string, error) {
	return reggen.Generate(columnConst.Payload.Regex, columnConst.Payload.Limit)
}
