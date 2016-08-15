package main

import (
	"encoding/json"
	"fmt"
	"io"
)

var (
	errUnsupportedType = "Column: %s has unsupported type: %s"
)

type ColumnFactory struct {
}

func (columnFactory ColumnFactory) CreateColumn(jSONColumn JSONColumn) (TypedColumn, error) {

	nestedColums, err := createColumns(jSONColumn.Columns)
	if err != nil {
		return nil, err
	}

	column := Column{
		Name:         jSONColumn.Name,
		TypedColumns: nestedColums,
	}
	// "EmptyIndeces":     countEmptyIndeces(jSONColumn.EmptyPercent),
	// "Locale":      retrieveLocale(jSONColumn),

	var typedColumn TypedColumn
	var errPayload error

	switch jSONColumn.Type {
	case COLUMN_CONST:
		payload := ColumnConstPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnConst{
			column:  column,
			Payload: payload,
		}
	case COLUMN_CREDIT_CARD_NUMBER:
		payload := ColumnCreditCardNumberPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnCreditCardNumber{
			column:  column,
			Payload: payload,
		}
	case COLUMN_CREDIT_CARD_TYPE:
		typedColumn = ColumnCreditCardType{
			column: column,
		}
	case COLUMN_JOIN:
		payload := ColumnJoinPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnJoin{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FIRST:
		typedColumn = ColumnNameFirst{
			column: column,
		}
	case COLUMN_NAME_FIRST_FEMALE:
		typedColumn = ColumnNameFirstFemale{
			column: column,
		}
	case COLUMN_NAME_FIRST_MALE:
		typedColumn = ColumnNameFirstMale{
			column: column,
		}
	case COLUMN_NAME_FULL:
		typedColumn = ColumnNameFull{
			column: column,
		}
	case COLUMN_NAME_FULL_FEMALE:
		typedColumn = ColumnNameFullFemale{
			column: column,
		}
	case COLUMN_NAME_FULL_MALE:
		typedColumn = ColumnNameFullMale{
			column: column,
		}
	case COLUMN_NAME_LAST:
		typedColumn = ColumnNameLast{
			column: column,
		}
	case COLUMN_NAME_LAST_FEMALE:
		typedColumn = ColumnNameLastFemale{
			column: column,
		}
	case COLUMN_NAME_LAST_MALE:
		typedColumn = ColumnNameLastMale{
			column: column,
		}
	case COLUMN_REGEX:
		payload := ColumnRegexPayload{
			Limit: 10,
		}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnRegex{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ROW_INDEX:
		// payload := ColumnRowIndexPayload{}
		// errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnRowIndex{
			column: column,
			// Payload: payload,
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
