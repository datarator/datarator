package main

import "strconv"

type ColumnOptionsRowIndex struct {
	// Increment int
	// Decrement int
}

type ColumnRowIndex struct {
	Column  Column
	Options ColumnOptionsRowIndex
}

func (columnRowIndex ColumnRowIndex) Value(context Context) (string, error) {
	//	return context.RowIndex + columnRowIndex.Options.Increment - columnRowIndex.Options.Decrement, nil
	return strconv.Itoa(context.CurrentRowIndex), nil
}
