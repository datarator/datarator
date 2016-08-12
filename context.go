package main

type Column struct {
	Name         string
	TypedColumns []TypedColumn
	EmptyIndeces []int
	Locale       string
}

type TypedColumn interface {
	Value(context *Context) (string, error)
}

type Schema struct {
	Document     string
	EmptyValue   string
	Count        int
	TypedColumns []TypedColumn
	Locale       string
}

type Template interface {
	Generate(context *Context) (string, error)
}

type Context struct {
	FromIndex        int
	ToIndex          int
	CurrentRowIndex  int
	CurrentRowValues map[string]string // maps column names to generated values
}
