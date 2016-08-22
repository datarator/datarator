package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func irisTester(t *testing.T) *httpexpect.Expect {
	api := IrisAPI()

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
		inFile  string
		outFile string
	}{
		{
			inFile:  "./testresource/csv_const_in.json",
			outFile: "./testresource/csv_const_out",
		},
		{
			inFile:  "./testresource/csv_join_in.json",
			outFile: "./testresource/csv_join_out",
		},
		{
			inFile:  "./testresource/csv_regex_in.json",
			outFile: "./testresource/csv_regex_out",
		},
		{
			inFile:  "./testresource/xml_flat_in.json",
			outFile: "./testresource/xml_flat_out",
		},
		{
			inFile:  "./testresource/xml_misc_xml_payload_in.json",
			outFile: "./testresource/xml_misc_xml_payload_out",
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

		irisTest.POST("/api/schemas/foo").WithBytes(in).
			Expect().
			Status(http.StatusOK).
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
		outErrCode        int
		outErrStringRegex string
	}{
		{
			inFile:            "./testresource/err_unsupported_type.json",
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "columns[.]0[.]type: columns[.]0[.]type must be one of the following:",
		},
		{
			inFile:            "./testresource/err_unsupported_template.json",
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "template: template must be one of the following: \"csv\"",
		},
		{
			inFile:            "./testresource/err_invalid_json.json",
			outErrCode:        http.StatusBadRequest,
			outErrStringRegex: "on line 4 column 2 got :invalid character '}' looking for beginning of object key string",
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
			Status(test.outErrCode).
			Body().Match(test.outErrStringRegex).NotEmpty()
	}
}
