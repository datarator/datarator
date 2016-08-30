package main

import (
	"fmt"
	"strings"
)

const (
	columnCopy               = "copy"
	errColumnCopyInvalidFrom = "Invalid from value: %s (expected column names hierarchy separated by '/') for column: %s"
)

type ColumnCopyPayload struct {
	From string
}

type ColumnCopy struct {
	TypedColumnBase
	payload ColumnCopyPayload
}

func (column ColumnCopy) Value(chunk *Chunk) (string, error) {
	currentChunk := chunk
	names := strings.Split(column.payload.From, "/")
	for i := 0; i < len(names); i++ {
		if ".." == names[i] {
			currentChunk = currentChunk.parent
		} else if i == len(names)-1 {
			return currentChunk.values[names[i]], nil
		} else {
			return "", fmt.Errorf(errColumnCopyInvalidFrom, column.payload.From, column.Column().name)
		}
	}
	return "", fmt.Errorf(errColumnCopyInvalidFrom, column.payload.From, column.Column().name)
}
