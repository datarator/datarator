package main

import (
	"net/http"

	"github.com/kataras/iris"
)

var (
	errStaticDataNotFound = "File: %s was not found"
)

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

		parser := ValidatingJSONRequestParser{}
		template, err := parser.parse(ctx, &jSONSchema, id)
		if err != nil {
			emmitError(http.StatusBadRequest, err.Error(), ctx)
			return
		}

		// TODO slice to chunks + in separate go routines
		context := Context{
			FromIndex:    0,
			ToIndex:      jSONSchema.Count,
			CurrentIndex: []int{0},
		}

		out, err := template.Generate(&context)
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
