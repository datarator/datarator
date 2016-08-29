package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/MStoykov/jsonutils"
	"github.com/kataras/iris"
	"github.com/xeipuuv/gojsonschema"
)

const (
	errPostDataEmpty       = "POST data empty"
	errPostJSONDataInvalid = "POST JSON data invalid:\n%s"
)

type RequestParser interface {
	parse(ctx *iris.Context, jSONSchema JSONSchema, id string) (Template, error)
}

type JSONColumn struct {
	Name         string
	Type         string
	EmptyPercent float32 `json:"emptyPercent"`
	Locale       string
	Columns      []JSONColumn
	JSONPayload  json.RawMessage `json:"payload"`
}

type JSONSchema struct {
	Template    string
	EmptyValue  string `json:"emptyValue"`
	Locale      string
	Count       int
	Columns     []JSONColumn
	JSONPayload json.RawMessage `json:"payload"`
}

type ValidatingJSONRequestParser struct {
}

func (p ValidatingJSONRequestParser) parse(ctx *iris.Context, jSONSchema *JSONSchema, id string) (Template, error) {
	// load JSON
	request := ctx.RequestCtx.Request.Body()
	if (len(request)) == 0 {
		return nil, errors.New(errPostDataEmpty)
	}

	// validate against JSON schema
	if err := validateRequest(request); err != nil {
		return nil, err
	}

	// unmarshall JSON
	if err := json.Unmarshal(request, jSONSchema); err != nil && err != io.EOF {
		return nil, err
	}
	template, err := TemplateFactory{}.CreateTemplate(id, jSONSchema)
	if err != nil {
		return nil, err
	}

	return template, nil
}

func validateRequest(request []byte) error {
	// TODO read once only
	schemaBytes, err := readFile("json_schema.json")
	if err != nil {
		return err
	}
	schemaLoader := gojsonschema.NewStringLoader(string(schemaBytes))
	documentLoader := gojsonschema.NewStringLoader(string(request))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		if serr, ok := err.(*json.SyntaxError); ok {
			return jsonutils.NewSyntaxError(serr, request, 5)
		}
		return err
	}
	if !result.Valid() {
		var buffer bytes.Buffer
		for _, desc := range result.Errors() {
			buffer.WriteString(fmt.Sprintf("- %s\n", desc))
		}
		return fmt.Errorf(errPostJSONDataInvalid, buffer.String())
	}

	return nil
}
