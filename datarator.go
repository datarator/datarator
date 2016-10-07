package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"os"

	"github.com/kataras/iris"
)

const (
	errStaticDataNotFound = "File: %s was not found"

	// used for header: "Accept-Encoding" (to check if client supports compression)
	strGzip = "gzip"
)

func IrisAPI() *iris.Framework {

	api := iris.New()

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

		gzipAllowed := ctx.Request.Header.HasAcceptEncoding(strGzip)

		ctx.SetContentType(template.ContentType())

		ctx.StreamWriter(func(writer *bufio.Writer) {
			processor, err := ChunkProcessorFactory{}.CreateChunkProcessor()
			if err != nil {
				emmitError(http.StatusInternalServerError, err.Error(), ctx)
				return
			}

			doneChannel := make(chan struct{})
			defer close(doneChannel)

			outChannel := processor.Process(jSONSchema.Count, template, doneChannel)
			timeoutChannel := time.After(time.Duration(opts.Timeout) * time.Millisecond)

			var gzipWriter *gzip.Writer
			if gzipAllowed {
				gzipWriter = gzip.NewWriter(writer)
			}

			for {
				select {
				// timed out
				case <-timeoutChannel:
					// TODO: for processing timeout, OK result, without data, is that OK?
					// emmitError(http.StatusGatewayTimeout, "Too slow, time's up!", ctx)
					return
				case out, ok := <-outChannel:
					if out.err != nil {
						emmitError(http.StatusInternalServerError, out.err.Error(), ctx)
						return
					}

					// if channel closed
					if !ok {

						if gzipAllowed {
							if err := gzipWriter.Flush(); err != nil {
								emmitError(http.StatusInternalServerError, err.Error(), ctx)
								return
							}
							if err := gzipWriter.Close(); err != nil {
								emmitError(http.StatusInternalServerError, err.Error(), ctx)
								return
							}
						}

						writer.Flush()
						return
					}

					if gzipAllowed {
						if _, err := gzipWriter.Write(out.out); err != nil {
							emmitError(http.StatusInternalServerError, err.Error(), ctx)
							return
						}
					} else {
						if _, err := writer.Write(out.out); err != nil {
							emmitError(http.StatusInternalServerError, err.Error(), ctx)
							return
						}

					}
				}
			}
		})
	})

	return api
}

func main() {
	if exit := parseFlags(); exit {
		os.Exit(0)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Starting datarator (v. %s)...\n", version)
	IrisAPI().Listen(fmt.Sprintf(":%d", opts.Port))
}

func emmitError(errorCode int, errorString string, ctx *iris.Context) {
	ctx.EmitError(errorCode)
	ctx.Write(": " + errorString)
}
