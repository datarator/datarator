package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
)

const (
	errUnsupportedType = "Column: %s has unsupported type: %s"
)

type Column struct {
	name         string
	columns      []TypedColumn
	emptyIndeces []int
	// locale       string
}

type TypedColumn interface {
	Value(chunk *Chunk) (string, error)
	Column() Column
	Payload() TypedColumnPayload
}

type TypedColumnBase struct {
	column  Column
	payload TypedColumnPayload
}

func (column TypedColumnBase) Column() Column {
	return column.column
}

func (column TypedColumnBase) Payload() TypedColumnPayload {
	return column.payload
}

type TypedColumnPayload interface {
	XmlType() string
}

type TypedColumnBasePayload struct {
	Xml string `json:"xml"`
}

func (payload TypedColumnBasePayload) XmlType() string {
	return payload.Xml
}

type ColumnFactory struct {
}

func (columnFactory ColumnFactory) CreateColumn(jSONColumn JSONColumn) (TypedColumn, error) {

	nestedColums, err := createColumns(jSONColumn.Columns)
	if err != nil {
		return nil, err
	}

	column := Column{
		name:    jSONColumn.Name,
		columns: nestedColums,
	}
	// "EmptyIndeces":     countEmptyIndeces(jSONColumn.EmptyPercent),
	// "Locale":      retrieveLocale(jSONColumn),

	var typedColumn TypedColumn
	var errPayload error

	basePayload := TypedColumnBasePayload{}
	errBasePayload := loadPayload(jSONColumn.JSONPayload, &basePayload)
	if errBasePayload != nil {
		return nil, errBasePayload
	}
	typedColumnBase := TypedColumnBase{
		column:  column,
		payload: basePayload,
	}

	switch jSONColumn.Type {
	case columnAddressContinent:
		typedColumn = ColumnAddressContinent{
			TypedColumnBase: typedColumnBase,
		}
	case columnAddressCountry:
		typedColumn = ColumnAddressCountry{
			TypedColumnBase: typedColumnBase,
		}
	case columnAddressCity:
		typedColumn = ColumnAddressCity{
			TypedColumnBase: typedColumnBase,
		}
	case columnAddressPhone:
		typedColumn = ColumnAddressPhone{
			TypedColumnBase: typedColumnBase,
		}
	case columnAddressState:
		typedColumn = ColumnAddressState{
			TypedColumnBase: typedColumnBase,
		}
	case columnAddressStreet:
		typedColumn = ColumnAddressStreet{
			TypedColumnBase: typedColumnBase,
		}
	case columnAddressZip:
		typedColumn = ColumnAddressZip{
			TypedColumnBase: typedColumnBase,
		}
	case columnColor:
		typedColumn = ColumnColor{
			TypedColumnBase: typedColumnBase,
		}
	case columnColorHex:
		payload := ColumnColorHexPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnColorHex{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnConst:
		payload := ColumnConstPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnConst{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnCopy:
		payload := ColumnCopyPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnCopy{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnCreditCardNumber:
		payload := ColumnCreditCardNumberPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnCreditCardNumber{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnCreditCardType:
		typedColumn = ColumnCreditCardType{
			TypedColumnBase: typedColumnBase,
		}
	case columnCurrency:
		typedColumn = ColumnCurrency{
			TypedColumnBase: typedColumnBase,
		}
	case columnCurrencyCode:
		typedColumn = ColumnCurrencyCode{
			TypedColumnBase: typedColumnBase,
		}

	case columnDateDayOfWeek:
		typedColumn = ColumnDateDayOfWeek{
			TypedColumnBase: typedColumnBase,
		}
	case columnDateDayOfWeekName:
		payload := ColumnDateDayOfWeekNamePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnDateDayOfWeekName{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnDateDayOfMonth:
		typedColumn = ColumnDateDayOfMonth{
			TypedColumnBase: typedColumnBase,
		}
	case columnDateMonth:
		typedColumn = ColumnDateMonth{
			TypedColumnBase: typedColumnBase,
		}
	case columnDateMonthName:
		payload := ColumnDateMonthNamePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnDateMonthName{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnDateYear:
		typedColumn = ColumnDateYear{
			TypedColumnBase: typedColumnBase,
		}
	case columnDateOfBirth:
		payload := ColumnDateOfBirthPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnDateOfBirth{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnJoin:
		payload := ColumnJoinPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnJoin{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnNameFirst:
		typedColumn = ColumnNameFirst{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameFirstFemale:
		typedColumn = ColumnNameFirstFemale{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameFirstMale:
		typedColumn = ColumnNameFirstMale{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameFull:
		typedColumn = ColumnNameFull{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameFullFemale:
		typedColumn = ColumnNameFullFemale{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameFullMale:
		typedColumn = ColumnNameFullMale{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameLast:
		typedColumn = ColumnNameLast{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameLastFemale:
		typedColumn = ColumnNameLastFemale{
			TypedColumnBase: typedColumnBase,
		}
	case columnNameLastMale:
		typedColumn = ColumnNameLastMale{
			TypedColumnBase: typedColumnBase,
		}
	case columnRegex:
		payload := ColumnRegexPayload{
			Limit: 10,
		}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnRegex{
			TypedColumnBase: typedColumnBase,
			payload:         payload,
		}
	case columnRowIndex:
		typedColumn = ColumnRowIndex{
			TypedColumnBase: typedColumnBase,
		}
	default:
		return nil, fmt.Errorf(errUnsupportedType, jSONColumn.Name, jSONColumn.Type)
	}

	if errPayload != nil {
		return nil, errPayload
	}

	return typedColumn, nil
}

// func countEmptyIndeces(EmptyPercent float32) ([]int, error) {
//     // TODO
//     return []int {1}, nil
// }

// func retrieveLocale(jSONColumn JSONColumn) (string, error) {
//     // TODO traverse all the way up to root to retrieve the locale
//     return "en", nil
// }

func createColumns(columns []JSONColumn) ([]TypedColumn, error) {
	columnFactory := ColumnFactory{}
	nestedColumns := []TypedColumn{}

	for _, nestedJSONColumn := range columns {
		nestedColumn, err := columnFactory.CreateColumn(nestedJSONColumn)
		if err != nil {
			return nil, err
		}
		nestedColumns = append(nestedColumns, nestedColumn)
	}
	return nestedColumns, nil
}

func loadPayload(jSONPayload json.RawMessage, payload interface{}) error {
	if len(jSONPayload) > 0 {
		err := json.Unmarshal(jSONPayload, &payload)
		if err != nil && err != io.EOF {
			return err
		}
	}
	return nil
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
