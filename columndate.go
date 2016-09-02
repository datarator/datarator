package main

import (
	"strconv"

	"github.com/syscrusher/fake"
)

const (
	columnDateDayOfWeek     = "date.day.of_week"
	columnDateDayOfWeekName = "date.day.of_week.name"
	columnDateDayOfMonth    = "date.day.of_month"
	columnDateMonth         = "date.month"
	columnDateMonthName     = "date.month.name"
	columnDateYear          = "date.year"
	columnDateOfBirth       = "date.of_birth"
)

type ColumnDateDayOfWeek struct {
	TypedColumnBase
}

func (column ColumnDateDayOfWeek) Value(chunk *Chunk) (string, error) {
	return strconv.Itoa(fake.WeekdayNum()), nil
}

type ColumnDateDayOfWeekNamePayload struct {
	Short bool
}

type ColumnDateDayOfWeekName struct {
	TypedColumnBase
	payload ColumnDateDayOfWeekNamePayload
}

func (column ColumnDateDayOfWeekName) Value(chunk *Chunk) (string, error) {
	if column.payload.Short {
		return fake.WeekDayShort(), nil
	}
	return fake.WeekDay(), nil
}

type ColumnDateDayOfMonth struct {
	TypedColumnBase
}

func (column ColumnDateDayOfMonth) Value(chunk *Chunk) (string, error) {
	return strconv.Itoa(fake.Day()), nil
}

type ColumnDateMonth struct {
	TypedColumnBase
}

func (column ColumnDateMonth) Value(chunk *Chunk) (string, error) {
	return strconv.Itoa(fake.MonthNum()), nil
}

type ColumnDateMonthNamePayload struct {
	Short bool
}

type ColumnDateMonthName struct {
	TypedColumnBase
	payload ColumnDateMonthNamePayload
}

func (column ColumnDateMonthName) Value(chunk *Chunk) (string, error) {
	if column.payload.Short {
		return fake.MonthShort(), nil
	}
	return fake.Month(), nil
}

type ColumnDateYear struct {
	TypedColumnBase
}

func (column ColumnDateYear) Value(chunk *Chunk) (string, error) {
	return strconv.Itoa(fake.Year(0, 2000)), nil
}

type ColumnDateOfBirthPayload struct {
	Age int
}

type ColumnDateOfBirth struct {
	TypedColumnBase
	payload ColumnDateOfBirthPayload
}

func (column ColumnDateOfBirth) Value(chunk *Chunk) (string, error) {
	if column.payload.Age > 0 {
		return fake.Birthdate(column.payload.Age).String(), nil
	}
	return fake.Birthdate(randInt(0, 120)).String(), nil
}
