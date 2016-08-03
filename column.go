package main

import (
	"encoding/json"
	"io"
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
	var errOpts error

	switch jSONColumn.Type {
	case COLUMN_CONST:
		options := ColumnOptionsConst{}
		errOpts = loadOptions(jSONColumn.JSONOptions, &options)
		typedColumn = ColumnConst{
			Options: options,
			Column:  column,
		}
	case COLUMN_NAME_FIRST:
		typedColumn = ColumnNameFirst{
			Column: column,
		}
	case COLUMN_NAME_LAST:
		typedColumn = ColumnNameLast{
			Column: column,
		}
	case COLUMN_JOIN:
		options := ColumnOptionsJoin{}
		errOpts = loadOptions(jSONColumn.JSONOptions, &options)
		typedColumn = ColumnJoin{
			Options: options,
			Column:  column,
		}
	case COLUMN_ROW_INDEX:
		options := ColumnOptionsRowIndex{}
		errOpts = loadOptions(jSONColumn.JSONOptions, &options)
		typedColumn = ColumnRowIndex{
			Options: options,
			Column:  column,
		}
	default:
		return nil, nil // errors.New("Invalid schema Type: %v", schemaType)
	}

	if errOpts != nil {
		return nil, errOpts
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
			return nil, nil
		}
		nestedColumns = append(nestedColumns, nestedColumn)
	}
	return nestedColumns, nil
}

func loadOptions(jSONOptions json.RawMessage, options interface{}) error {
	if len(jSONOptions) > 0 {
		err := json.Unmarshal(jSONOptions, &options)
		if err != nil && err != io.EOF {
			return err
		}
	}
	return nil
}
