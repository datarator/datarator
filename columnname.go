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

type ColumnNameFirstPayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameFirst struct {
	column  Column
	Payload ColumnNameFirstPayload `json:"payload"`
}

func (column ColumnNameFirst) Value(context *Context) (string, error) {
	return fake.FirstName(), nil
}

func (column ColumnNameFirst) Column() Column {
	return column.column
}

func (column ColumnNameFirst) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameFirstFemalePayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameFirstFemale struct {
	column  Column
	Payload ColumnNameFirstFemalePayload `json:"payload"`
}

func (column ColumnNameFirstFemale) Value(context *Context) (string, error) {
	return fake.FemaleFirstName(), nil
}

func (column ColumnNameFirstFemale) Column() Column {
	return column.column
}

func (column ColumnNameFirstFemale) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameFirstMalePayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameFirstMale struct {
	column  Column
	Payload ColumnNameFirstMalePayload `json:"payload"`
}

func (column ColumnNameFirstMale) Value(context *Context) (string, error) {
	return fake.MaleFirstName(), nil
}

func (column ColumnNameFirstMale) Column() Column {
	return column.column
}

func (column ColumnNameFirstMale) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameFullPayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameFull struct {
	column  Column
	Payload ColumnNameFullPayload `json:"payload"`
}

func (column ColumnNameFull) Value(context *Context) (string, error) {
	return fake.FullName(), nil
}

func (column ColumnNameFull) Column() Column {
	return column.column
}

func (column ColumnNameFull) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameFullFemalePayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameFullFemale struct {
	column  Column
	Payload ColumnNameFullFemalePayload `json:"payload"`
}

func (column ColumnNameFullFemale) Value(context *Context) (string, error) {
	return fake.FemaleFullName(), nil
}

func (column ColumnNameFullFemale) Column() Column {
	return column.column
}

func (column ColumnNameFullFemale) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameFullMalePayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameFullMale struct {
	column  Column
	Payload ColumnNameFullMalePayload `json:"payload"`
}

func (column ColumnNameFullMale) Value(context *Context) (string, error) {
	return fake.MaleFullName(), nil
}

func (column ColumnNameFullMale) Column() Column {
	return column.column
}

func (column ColumnNameFullMale) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameLastPayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameLast struct {
	column  Column
	Payload ColumnNameLastPayload `json:"payload"`
}

func (column ColumnNameLast) Value(context *Context) (string, error) {
	return fake.LastName(), nil
}

func (column ColumnNameLast) Column() Column {
	return column.column
}

func (column ColumnNameLast) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameLastFemalePayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameLastFemale struct {
	column  Column
	Payload ColumnNameLastFemalePayload `json:"payload"`
}

func (column ColumnNameLastFemale) Value(context *Context) (string, error) {
	return fake.FemaleLastName(), nil
}
func (column ColumnNameLastFemale) Column() Column {
	return column.column
}

func (column ColumnNameLastFemale) XmlType() string {
	return column.Payload.XmlType
}

type ColumnNameLastMalePayload struct {
	XmlType string `json:"xml"`
}

type ColumnNameLastMale struct {
	column  Column
	Payload ColumnNameLastMalePayload `json:"payload"`
}

func (column ColumnNameLastMale) Value(context *Context) (string, error) {
	return fake.MaleLastName(), nil
}

func (column ColumnNameLastMale) Column() Column {
	return column.column
}

func (column ColumnNameLastMale) XmlType() string {
	return column.Payload.XmlType
}
