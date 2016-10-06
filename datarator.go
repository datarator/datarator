package main

import (
	"bufio"
	"bytes"
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

					// channel closed
					if !ok {
						writer.Flush()
						return
					}

					// TODO optimize performance
					if gzipAllowed {
						gzipped, err := pack(out.out)
						if err != nil {
							// gzipping failed => go on in a non-gzipped way
						} else {
							out.out = gzipped.String()
						}
					}
					writer.WriteString(out.out)
				}
			}
		})
	})

	api.Get("/", func(ctx *iris.Context) {
		ctx.WriteString("Datarator at your service")
	})

	return api
}

func pack(str string) (bytes.Buffer, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		return b, err
	}
	if err := gz.Flush(); err != nil {
		return b, err
	}
	if err := gz.Close(); err != nil {
		return b, err
	}
	return b, nil
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
