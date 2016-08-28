package main

import "errors"

var (
	errNestingIndexCantBeDecremented = "Nesting index: '0' can't be decremented!"
)

type Schema struct {
	Document     string
	EmptyValue   string
	Count        int
	TypedColumns []TypedColumn
	Locale       string
}

type Context struct {
	FromIndex           int
	ToIndex             int
	CurrentIndex        []int             // supports nesting
	CurrentIndexValues  map[string]string // maps column index (nested column index separated by "|") to generated value
	CurrentNameValues   map[string]string // maps column names to generated values (nested column names are separated by "|")
	CurrentNestingLevel int               // how deep are we in nesting
}

func (context *Context) getCurrentIndex() int {
	return context.CurrentIndex[context.CurrentNestingLevel]
}

func (context *Context) setCurrentIndex(val int) {
	context.CurrentIndex[context.CurrentNestingLevel] = val
}

func (context *Context) incrementCurrentIndex() int {
	context.CurrentIndex[context.CurrentNestingLevel]++
	return context.CurrentIndex[context.CurrentNestingLevel]
}

func (context *Context) nest() error {
	context.CurrentNestingLevel++
	context.CurrentIndex = append(context.CurrentIndex, 0)
	return nil
}

func (context *Context) unnest() error {
	if len(context.CurrentIndex) == 0 {
		return errors.New(errNestingIndexCantBeDecremented)
	}

	context.CurrentIndex = context.CurrentIndex[:context.CurrentNestingLevel]
	context.CurrentNestingLevel--
	return nil
}
