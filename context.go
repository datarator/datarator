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
	FromIndex           int
	ToIndex             int
	CurrentIndex        []int             // supports nesting
	CurrentIndexValues  map[string]string // maps column index (nested column index separated by "|") to generated value
	CurrentNameValues   map[string]string // maps column names to generated values (nested column names are separated by "|")
	CurrentNestingLevel int               // how deep are we in nesting
}

func (context Context) getCurrentIndex() int {
	return context.CurrentIndex[context.CurrentNestingLevel]
}

func (context Context) setCurrentIndex(val int) {
	context.CurrentIndex[context.CurrentNestingLevel] = val
}

func (context Context) incrementCurrentIndex() int {
	context.CurrentIndex[context.CurrentNestingLevel] = context.CurrentIndex[context.CurrentNestingLevel] + 1
	return context.CurrentIndex[context.CurrentNestingLevel]
}
