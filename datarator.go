package main

import (
	"encoding/json"
	"errors"

	"github.com/kataras/iris"
)

type SchemaAPI struct {
	*iris.Context
}

type JSONColumn struct {
	Name         string
	Type         string
	EmptyPercent float32 `json:"emptyPercent"`
	Locale       string
	Columns      []JSONColumn    `json:"columns"`
	JSONOptions  json.RawMessage `json:"options"`
}

type JSONSchema struct {
	Template    string
	EmptyValue  string `json:"emptyValue"`
	Locale      string
	Count       int
	Columns     []JSONColumn    `json:"columns"`
	JSONOptions json.RawMessage `json:"options"`
}

// errors

var (
	errFoo = errors.New("foo: %s")
	// errFlashNotFound    = errors.New("Unable to get flash message. Trace: Cookie does not exists")
	// errSessionNil       = errors.New("Unable to set session, Config().Session.Provider is nil, please refer to the docs!")
	// errNoForm           = errors.New("Request has no any valid form")
	// errWriteJSON        = errors.New("Before JSON be written to the body, JSON Encoder returned an error. Trace: %s")
	// errRenderMarshalled = errors.New("Before +type Rendering, MarshalIndent returned an error. Trace: %s")
	// errReadBody         = errors.New("While trying to read %s from the request body. Trace %s")
	// errServeContent     = errors.New("While trying to serve content to the client. Trace %s")
)

// POST /schemas/:param1
func (schemaAPI SchemaAPI) PostBy(id string) {
	// defaults
	jSONSchema := JSONSchema{
		Template:   "csv",
		EmptyValue: "",
		Count:      10,
	}

	// parse input
	if err := schemaAPI.ReadJSON(&jSONSchema); err != nil {
		schemaAPI.EmitError(iris.StatusInternalServerError)
		schemaAPI.Write(err.Error())
		return
	}

	template, err := TemplateFactory{}.CreateTemplate(id, jSONSchema)
	if err != nil {
		schemaAPI.EmitError(iris.StatusInternalServerError)
		schemaAPI.Write(err.Error())
		return
	}

	// TODO slice to chunks + in separate go routines
	context := Context{
		CurrentRowIndex: 0,
		FromIndex:       0,
		ToIndex:         jSONSchema.Count,
	}

	out, err := template.Generate(context)
	if err != nil {
		schemaAPI.EmitError(iris.StatusInternalServerError)
		schemaAPI.Write(err.Error())
		return
	}

	// name := u.FormValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	// println(string(name))
	println("POST /api/schemas/" + id)

	schemaAPI.Text(iris.StatusOK, out)
}

func main() {
	iris.API("/api/schemas", SchemaAPI{})
	iris.Listen(":9292")
}

// func main() {

// 	iris.Post("/api/schemas", func(ctx *iris.Context) {
// 		// https://kataras.gitbooks.io/iris/content/render_rest.html
// 		ctx.Text(iris.StatusOK, "Plain text here")

// 		// new(ParserJSON).Parse(in)

// 		// rowindex generate
// 		//
// 	})

// 	iris.Listen(":9292")
// }
