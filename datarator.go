package main

import (
	"encoding/json"
	"net/http"

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
			emmitError(http.StatusBadRequest, err.Error(), ctx)
			return
		}

		template, err := TemplateFactory{}.CreateTemplate(id, jSONSchema)
		if err != nil {
			emmitError(http.StatusBadRequest, err.Error(), ctx)
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
			emmitError(http.StatusInternalServerError, err.Error(), ctx)
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

func emmitError(errorCode int, errorString string, ctx *iris.Context) {
	ctx.EmitError(errorCode)
	ctx.Write(": " + errorString)
}
