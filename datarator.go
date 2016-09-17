package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/kataras/iris"
)

const (
	errStaticDataNotFound = "File: %s was not found"
)

func IrisAPI() *iris.Framework {

	api := iris.New()

	// TODO Gzip
	// iris.Config.Gzip = true

	// define the api
	api.Post("/api/schemas/:id", func(ctx *iris.Context) {

		// TODO implement timeout
		// time.After(time.Duration(*timeoutFlag) * time.Millisecond)

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
			writer.Flush()
		})

		// TODO implement timeout
		// select {
		// 	case <-ch:
		// 		// a read from ch has occurred
		// 	case <-timeout:
		// 		// the read from ch has timed out
		// }
	})

	api.Get("/", func(ctx *iris.Context) {
		ctx.WriteString("Datarator at your service")
	})

	return api
}

func main() {
	parseFlags()
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Starting datarator (v. %s)...\n", version)
	IrisAPI().Listen(fmt.Sprintf(":%d", opts.Port))
}

func emmitError(errorCode int, errorString string, ctx *iris.Context) {
	ctx.EmitError(errorCode)
	ctx.Write(": " + errorString)
}
