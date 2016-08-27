package main

import (
	"strconv"

	"github.com/syscrusher/fake"
)

const (
	COLUMN_DATE_DAY_OF_WEEK      = "date.day.of_week"
	COLUMN_DATE_DAY_OF_WEEK_NAME = "date.day.of_week.name"
	COLUMN_DATE_DAY_OF_MONTH     = "date.day.of_month"
	COLUMN_DATE_MONTH            = "date.month"
	COLUMN_DATE_MONTH_NAME       = "date.month.name"
	COLUMN_DATE_YEAR             = "date.year"
	COLUMN_DATE_OF_BIRTH         = "date.of_birth"
)

type ColumnDateDayOfWeekPayload struct {
	XmlType string `json:"xml"`
}

type ColumnDateDayOfWeek struct {
	column  Column
	Payload ColumnDateDayOfWeekPayload `json:"payload"`
}

func (column ColumnDateDayOfWeek) Value(context *Context) (string, error) {
	return strconv.Itoa(fake.WeekdayNum()), nil
}

func (column ColumnDateDayOfWeek) Column() Column {
	return column.column
}

func (column ColumnDateDayOfWeek) XmlType() string {
	return column.Payload.XmlType
}

type ColumnDateDayOfWeekNamePayload struct {
	XmlType string `json:"xml"`
	short   bool
}

type ColumnDateDayOfWeekName struct {
	column  Column
	Payload ColumnDateDayOfWeekNamePayload `json:"payload"`
}

func (column ColumnDateDayOfWeekName) Value(context *Context) (string, error) {
	if column.Payload.short {
		return fake.WeekDayShort(), nil
	}
	return fake.WeekDay(), nil
}

func (column ColumnDateDayOfWeekName) Column() Column {
	return column.column
}

func (column ColumnDateDayOfWeekName) XmlType() string {
	return column.Payload.XmlType
}

type ColumnDateDayOfMonthPayload struct {
	XmlType string `json:"xml"`
}

type ColumnDateDayOfMonth struct {
	column  Column
	Payload ColumnDateDayOfMonthPayload `json:"payload"`
}

func (column ColumnDateDayOfMonth) Value(context *Context) (string, error) {
	return strconv.Itoa(fake.Day()), nil
}

func (column ColumnDateDayOfMonth) Column() Column {
	return column.column
}

func (column ColumnDateDayOfMonth) XmlType() string {
	return column.Payload.XmlType
}

type ColumnDateMonthPayload struct {
	XmlType string `json:"xml"`
}

type ColumnDateMonth struct {
	column  Column
	Payload ColumnDateMonthPayload `json:"payload"`
}

func (column ColumnDateMonth) Value(context *Context) (string, error) {
	return strconv.Itoa(fake.MonthNum()), nil
}

func (column ColumnDateMonth) Column() Column {
	return column.column
}

func (column ColumnDateMonth) XmlType() string {
	return column.Payload.XmlType
}

type ColumnDateMonthNamePayload struct {
	XmlType string `json:"xml"`
	short   bool
}

type ColumnDateMonthName struct {
	column  Column
	Payload ColumnDateMonthNamePayload `json:"payload"`
}

func (column ColumnDateMonthName) Value(context *Context) (string, error) {
	if column.Payload.short {
		return fake.MonthShort(), nil
	}
	return fake.Month(), nil
}

func (column ColumnDateMonthName) Column() Column {
	return column.column
}

func (column ColumnDateMonthName) XmlType() string {
	return column.Payload.XmlType
}

type ColumnDateYearPayload struct {
	XmlType string `json:"xml"`
}

type ColumnDateYear struct {
	column  Column
	Payload ColumnDateYearPayload `json:"payload"`
}

func (column ColumnDateYear) Value(context *Context) (string, error) {
	min := randInt(0, 2000)
	return strconv.Itoa(fake.Year(min, randInt(min, 2000))), nil
}

func (column ColumnDateYear) Column() Column {
	return column.column
}

func (column ColumnDateYear) XmlType() string {
	return column.Payload.XmlType
}

type ColumnDateOfBirthPayload struct {
	XmlType string `json:"xml"`
	age     int
}

type ColumnDateOfBirth struct {
	column  Column
	Payload ColumnDateOfBirthPayload `json:"payload"`
}

func (column ColumnDateOfBirth) Value(context *Context) (string, error) {
	if column.Payload.age > 0 {
		return fake.Birthdate(column.Payload.age).String(), nil
	}
	return fake.Birthdate(randInt(0, 2000)).String(), nil
}

func (column ColumnDateOfBirth) Column() Column {
	return column.column
}

func (column ColumnDateOfBirth) XmlType() string {
	return column.Payload.XmlType
}
