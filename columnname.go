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
	Column Column
}

func (column ColumnNameFirst) Value(context *Context) (string, error) {
	return fake.FirstName(), nil
}

type ColumnNameFirstFemale struct {
	Column Column
}

func (column ColumnNameFirstFemale) Value(context *Context) (string, error) {
	return fake.FemaleFirstName(), nil
}

type ColumnNameFirstMale struct {
	Column Column
}

func (column ColumnNameFirstMale) Value(context *Context) (string, error) {
	return fake.MaleFirstName(), nil
}

type ColumnNameFull struct {
	Column Column
}

func (column ColumnNameFull) Value(context *Context) (string, error) {
	return fake.FullName(), nil
}

type ColumnNameFullFemale struct {
	Column Column
}

func (column ColumnNameFullFemale) Value(context *Context) (string, error) {
	return fake.FemaleFullName(), nil
}

type ColumnNameFullMale struct {
	Column Column
}

func (column ColumnNameFullMale) Value(context *Context) (string, error) {
	return fake.MaleFullName(), nil
}

type ColumnNameLast struct {
	Column Column
}

func (column ColumnNameLast) Value(context *Context) (string, error) {
	return fake.LastName(), nil
}

type ColumnNameLastFemale struct {
	Column Column
}

func (column ColumnNameLastFemale) Value(context *Context) (string, error) {
	return fake.FemaleLastName(), nil
}

type ColumnNameLastMale struct {
	Column Column
}

func (column ColumnNameLastMale) Value(context *Context) (string, error) {
	return fake.MaleLastName(), nil
}
