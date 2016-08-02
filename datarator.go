package main

import (
	"encoding/json"

	"github.com/kataras/iris"
)

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
// var (
// 	errFoo = errors.New("foo: %s")
// 	// errFlashNotFound    = errors.New("Unable to get flash message. Trace: Cookie does not exists")
// 	// errSessionNil       = errors.New("Unable to set session, Config().Session.Provider is nil, please refer to the docs!")
// 	// errNoForm           = errors.New("Request has no any valid form")
// 	// errWriteJSON        = errors.New("Before JSON be written to the body, JSON Encoder returned an error. Trace: %s")
// 	// errRenderMarshalled = errors.New("Before +type Rendering, MarshalIndent returned an error. Trace: %s")
// 	// errReadBody         = errors.New("While trying to read %s from the request body. Trace %s")
// 	// errServeContent     = errors.New("While trying to serve content to the client. Trace %s")
// )

func IrisAPI() *iris.Framework {
	api := iris.New()
	// define the api
	api.Post("/api/schemas/:id", func(ctx *iris.Context) {

		id := ctx.Param("id")

		// defaults
		jSONSchema := JSONSchema{
			Template:   "csv",
			EmptyValue: "",
			Count:      10,
		}

		// parse input
		if err := ctx.ReadJSON(&jSONSchema); err != nil {
			ctx.EmitError(iris.StatusInternalServerError)
			ctx.Write(err.Error())
			return
		}

		template, err := TemplateFactory{}.CreateTemplate(id, jSONSchema)
		if err != nil {
			ctx.EmitError(iris.StatusInternalServerError)
			ctx.Write(err.Error())
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
			ctx.EmitError(iris.StatusInternalServerError)
			ctx.Write(err.Error())
			return
		}

		println("POST /api/schemas/" + id)

		ctx.Text(iris.StatusOK, out)
	})

	return api
}

func main() {
	IrisAPI().Listen(":9292")
}
