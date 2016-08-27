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
	case COLUMN_ADDRESS_CONTINENT:
		payload := ColumnAddressContinentPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressContinent{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ADDRESS_COUNTRY:
		payload := ColumnAddressCountryPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressCountry{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ADDRESS_CITY:
		payload := ColumnAddressCityPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressCity{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ADDRESS_PHONE:
		payload := ColumnAddressPhonePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressPhone{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ADDRESS_STATE:
		payload := ColumnAddressStatePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressState{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ADDRESS_STREET:
		payload := ColumnAddressStreetPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressStreet{
			column:  column,
			Payload: payload,
		}
	case COLUMN_ADDRESS_ZIP:
		payload := ColumnAddressZipPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnAddressZip{
			column:  column,
			Payload: payload,
		}
	case COLUMN_COLOR:
		payload := ColumnColorPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnColor{
			column:  column,
			Payload: payload,
		}
	case COLUMN_COLOR_HEX:
		payload := ColumnColorHexPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnColorHex{
			column:  column,
			Payload: payload,
		}
	case COLUMN_CONST:
		payload := ColumnConstPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnConst{
			column:  column,
			Payload: payload,
		}
	case COLUMN_CURRENCY:
		payload := ColumnCurrencyPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnCurrency{
			column:  column,
			Payload: payload,
		}
	case COLUMN_CURRENCY_CODE:
		payload := ColumnCurrencyCodePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnCurrencyCode{
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
		payload := ColumnCreditCardTypePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnCreditCardType{
			column:  column,
			Payload: payload,
		}
	case COLUMN_JOIN:
		payload := ColumnJoinPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnJoin{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FIRST:
		payload := ColumnNameFirstPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameFirst{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FIRST_FEMALE:
		payload := ColumnNameFirstFemalePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameFirstFemale{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FIRST_MALE:
		payload := ColumnNameFirstMalePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameFirstMale{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FULL:
		payload := ColumnNameFullPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameFull{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FULL_FEMALE:
		payload := ColumnNameFullFemalePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameFullFemale{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_FULL_MALE:
		payload := ColumnNameFullMalePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameFullMale{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_LAST:
		payload := ColumnNameLastPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameLast{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_LAST_FEMALE:
		payload := ColumnNameLastFemalePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameLastFemale{
			column:  column,
			Payload: payload,
		}
	case COLUMN_NAME_LAST_MALE:
		payload := ColumnNameLastMalePayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnNameLastMale{
			column:  column,
			Payload: payload,
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
		payload := ColumnRowIndexPayload{}
		errPayload = loadPayload(jSONColumn.JSONPayload, &payload)
		typedColumn = ColumnRowIndex{
			column:  column,
			Payload: payload,
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
