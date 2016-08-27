package main

import (
	"bufio"
	"math/rand"
	"net/http"
	"time"

	"github.com/kataras/iris"
)

var (
	errStaticDataNotFound = "File: %s was not found"
)

func IrisAPI() *iris.Framework {

	api := iris.New()

	// TODO Gzip
	// iris.Config.Gzip = true

	// define the api
	api.Post("/api/schemas/:id", func(ctx *iris.Context) {

		id := ctx.Param("id")
		println("POST /api/schemas/" + id)

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

		ctx.SetContentType(template.ContentType())
		ctx.StreamWriter(func(writer *bufio.Writer) {
			// TODO Gzip
			// ctx.Response.WriteGzip()

			processor, err := ChunkProcessorFactory{}.CreateChunkProcessor()
			if err != nil {
				emmitError(http.StatusInternalServerError, err.Error(), ctx)
				return
			}

			doneChannel := make(chan struct{})
			defer close(doneChannel)

			outChannel := processor.Process(jSONSchema.Count, template, doneChannel)
			for out := range outChannel {
				if out.err != nil {
					emmitError(http.StatusInternalServerError, out.err.Error(), ctx)
					return
				}

				// ? TODO http://stackoverflow.com/questions/25171385/how-should-i-add-buffering-to-a-gzip-writer
				writer.WriteString(out.out)
			}
			// }
			writer.Flush()
		})

	})

	return api
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	IrisAPI().Listen(":9292")
}

func emmitError(errorCode int, errorString string, ctx *iris.Context) {
	ctx.EmitError(errorCode)
	ctx.Write(": " + errorString)
}
