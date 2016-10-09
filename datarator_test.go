package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

const defaultTimeout = 3000

func irisTester(t *testing.T) *httpexpect.Expect {
	api := IrisAPI()

	opts.Timeout = defaultTimeout
	opts.ChunkSize = 1000

	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(api.ListenVirtual().Handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func TestCsv(t *testing.T) {
	// if testing.Short() {
	//     t.Skip("skipping test in short mode.")
	// }

	var tests = []struct {
		inFile         string
		headers        map[string]string
		outFile        string
		outContentType string
	}{
		{
			inFile:         "./testresource/csv_const_in.json",
			outFile:        "./testresource/csv_const_out",
			headers:        make(map[string]string),
			outContentType: "text/csv",
		},
		{
			inFile:         "./testresource/csv_const_empty_in.json",
			outFile:        "./testresource/csv_const_empty_out",
			headers:        make(map[string]string),
			outContentType: "text/csv",
		},
		{
			inFile:         "./testresource/csv_const_in.json",
			outFile:        "./testresource/csv_const_out.gz",
			headers:        map[string]string{"Accept-Encoding": "gzip,deflate"},
			outContentType: "text/csv",
		},
		{
			inFile:         "./testresource/csv_join_in.json",
			outFile:        "./testresource/csv_join_out",
			headers:        make(map[string]string),
			outContentType: "text/csv",
		},
		{
			inFile:         "./testresource/csv_regex_in.json",
			outFile:        "./testresource/csv_regex_out",
			headers:        make(map[string]string),
			outContentType: "text/csv",
		},
		{
			inFile:  "./testresource/sql_const_in.json",
			outFile: "./testresource/sql_const_out", headers: make(map[string]string),
			outContentType: "text/sql",
		},
		{
			inFile:         "./testresource/xml_flat_in.json",
			outFile:        "./testresource/xml_flat_out",
			headers:        make(map[string]string),
			outContentType: "text/xml",
		},
		{
			inFile:         "./testresource/xml_misc_xml_payload_in.json",
			outFile:        "./testresource/xml_misc_xml_payload_out",
			headers:        make(map[string]string),
			outContentType: "text/xml",
		},
	}

	for _, test := range tests {
		in, errIn := ioutil.ReadFile(test.inFile)
		if errIn != nil {
			t.Fatalf("Failed reading file %v: %v", test.inFile, errIn.Error())
		}

		out, errOut := ioutil.ReadFile(test.outFile)
		if errOut != nil {
			t.Fatalf("Failed reading file %v: %v", test.inFile, errOut.Error())
		}

		irisTest := irisTester(t)

		irisTest.POST("/api/schemas/foo").
			WithHeaders(test.headers).WithBytes(in).
			Expect().
			Status(http.StatusOK).
			ContentType(test.outContentType).
			Body().Equal(string(out))
	}
}

func TestAllColumns(t *testing.T) {
	// if testing.Short() {
	//     t.Skip("skipping test in short mode.")
	// }

	var tests = []struct {
		inFile string
	}{
		{
			inFile: "./testresource/all_columns_in.json",
		},
	}

	for _, test := range tests {
		in, errIn := ioutil.ReadFile(test.inFile)
		if errIn != nil {
			t.Fatalf("Failed reading file %v: %v", test.inFile, errIn.Error())
		}

		irisTest := irisTester(t)

		irisTest.POST("/api/schemas/foo").WithBytes(in).
			Expect().
			Status(http.StatusOK).
			Body().Contains(",")
	}
}

func TestErr(t *testing.T) {
	var tests = []struct {
		inFile            string
		timeout           int
		outErrCode        int
		outErrStringRegex string
	}{
		{
			inFile:            "./testresource/err_unsupported_type.json",
			timeout:           defaultTimeout,
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "columns[.]0[.]type: columns[.]0[.]type must be one of the following:",
		},
		{
			inFile:            "./testresource/err_unsupported_template.json",
			timeout:           defaultTimeout,
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "template: template must be one of the following: \"csv\"",
		},
		{
			inFile:            "./testresource/err_invalid_json.json",
			timeout:           defaultTimeout,
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "on line 4 column 2 got :invalid character '}' looking for beginning of object key string",
		},
		{
			inFile:            "./testresource/err_invalid_empty_percent_negative.json",
			timeout:           0,
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "Must be greater than or equal to 0",
		},
		{
			inFile:            "./testresource/err_invalid_empty_percent_over_100.json",
			timeout:           0,
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "Must be less than or equal to 100",
		},
		{
			inFile:            "./testresource/err_timeout.json",
			timeout:           0,
			outErrCode:        http.StatusOK,
			outErrStringRegex: "",
		},
	}

	for _, test := range tests {
		in, errIn := ioutil.ReadFile(test.inFile)
		if errIn != nil {
			t.Fatalf("Failed reading file %v: %v", test.inFile, errIn.Error())
		}

		irisTest := irisTester(t)
		opts.Timeout = test.timeout

		irisTest.POST("/api/schemas/foo").WithBytes(in).
			Expect().
			Status(test.outErrCode).
			Body().Match(test.outErrStringRegex).NotEmpty()
	}
}
