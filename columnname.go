package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_NAME_FIRST = "name.first"
)

type ColumnNameFirst struct {
	Column Column
}

func (columnNameFirst ColumnNameFirst) Value(context Context) (string, error) {
	return fake.FirstName(), nil
}
