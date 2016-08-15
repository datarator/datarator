package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_NAME_FIRST        = "name.first"
	COLUMN_NAME_FIRST_FEMALE = "name.first.female"
	COLUMN_NAME_FIRST_MALE   = "name.first.male"
	COLUMN_NAME_FULL         = "name.full"
	COLUMN_NAME_FULL_FEMALE  = "name.full.female"
	COLUMN_NAME_FULL_MALE    = "name.full.male"
	COLUMN_NAME_LAST         = "name.last"
	COLUMN_NAME_LAST_FEMALE  = "name.last.female"
	COLUMN_NAME_LAST_MALE    = "name.last.male"
)

type ColumnNameFirst struct {
	column Column
}

func (column ColumnNameFirst) Value(context *Context) (string, error) {
	return fake.FirstName(), nil
}

func (column ColumnNameFirst) Column() Column {
	return column.column
}

type ColumnNameFirstFemale struct {
	column Column
}

func (column ColumnNameFirstFemale) Value(context *Context) (string, error) {
	return fake.FemaleFirstName(), nil
}

func (column ColumnNameFirstFemale) Column() Column {
	return column.column
}

type ColumnNameFirstMale struct {
	column Column
}

func (column ColumnNameFirstMale) Value(context *Context) (string, error) {
	return fake.MaleFirstName(), nil
}

func (column ColumnNameFirstMale) Column() Column {
	return column.column
}

type ColumnNameFull struct {
	column Column
}

func (column ColumnNameFull) Value(context *Context) (string, error) {
	return fake.FullName(), nil
}

func (column ColumnNameFull) Column() Column {
	return column.column
}

type ColumnNameFullFemale struct {
	column Column
}

func (column ColumnNameFullFemale) Value(context *Context) (string, error) {
	return fake.FemaleFullName(), nil
}

func (column ColumnNameFullFemale) Column() Column {
	return column.column
}

type ColumnNameFullMale struct {
	column Column
}

func (column ColumnNameFullMale) Value(context *Context) (string, error) {
	return fake.MaleFullName(), nil
}

func (column ColumnNameFullMale) Column() Column {
	return column.column
}

type ColumnNameLast struct {
	column Column
}

func (column ColumnNameLast) Value(context *Context) (string, error) {
	return fake.LastName(), nil
}

func (column ColumnNameLast) Column() Column {
	return column.column
}

type ColumnNameLastFemale struct {
	column Column
}

func (column ColumnNameLastFemale) Value(context *Context) (string, error) {
	return fake.FemaleLastName(), nil
}
func (column ColumnNameLastFemale) Column() Column {
	return column.column
}

type ColumnNameLastMale struct {
	column Column
}

func (column ColumnNameLastMale) Value(context *Context) (string, error) {
	return fake.MaleLastName(), nil
}

func (column ColumnNameLastMale) Column() Column {
	return column.column
}
