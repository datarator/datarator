package main

import (
	"github.com/syscrusher/fake"
)

const (
	columnNameFirst       = "name.first"
	columnNameFirstFemale = "name.first.female"
	columnNameFirstMale   = "name.first.male"
	columnNameFull        = "name.full"
	columnNameFullFemale  = "name.full.female"
	columnNameFullMale    = "name.full.male"
	columnNameLast        = "name.last"
	columnNameLastFemale  = "name.last.female"
	columnNameLastMale    = "name.last.male"
)

type ColumnNameFirst struct {
	TypedColumnBase
}

func (column ColumnNameFirst) Value(context *Context) (string, error) {
	return fake.FirstName(), nil
}

type ColumnNameFirstFemale struct {
	TypedColumnBase
}

func (column ColumnNameFirstFemale) Value(context *Context) (string, error) {
	return fake.FemaleFirstName(), nil
}

type ColumnNameFirstMale struct {
	TypedColumnBase
}

func (column ColumnNameFirstMale) Value(context *Context) (string, error) {
	return fake.MaleFirstName(), nil
}

type ColumnNameFull struct {
	TypedColumnBase
}

func (column ColumnNameFull) Value(context *Context) (string, error) {
	return fake.FullName(), nil
}

type ColumnNameFullFemale struct {
	TypedColumnBase
}

func (column ColumnNameFullFemale) Value(context *Context) (string, error) {
	return fake.FemaleFullName(), nil
}

type ColumnNameFullMale struct {
	TypedColumnBase
}

func (column ColumnNameFullMale) Value(context *Context) (string, error) {
	return fake.MaleFullName(), nil
}

type ColumnNameLast struct {
	TypedColumnBase
}

func (column ColumnNameLast) Value(context *Context) (string, error) {
	return fake.LastName(), nil
}

type ColumnNameLastFemale struct {
	TypedColumnBase
}

func (column ColumnNameLastFemale) Value(context *Context) (string, error) {
	return fake.FemaleLastName(), nil
}

type ColumnNameLastMale struct {
	TypedColumnBase
}

func (column ColumnNameLastMale) Value(context *Context) (string, error) {
	return fake.MaleLastName(), nil
}
