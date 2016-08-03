package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_NAME_FIRST = "name.first"
	COLUMN_NAME_LAST  = "name.last"
)

type ColumnNameFirst struct {
	Column Column
}

func (columnNameFirst ColumnNameFirst) Value(context Context) (string, error) {
	return fake.FirstName(), nil
}

type ColumnNameLast struct {
	Column Column
}

func (columnNameLast ColumnNameLast) Value(context Context) (string, error) {
	return fake.LastName(), nil
}
