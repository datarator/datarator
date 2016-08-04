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
			inFile:  "./testresource/csv_1_in.json",
			outFile: "./testresource/csv_1_out",
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

func TestErr(t *testing.T) {
	var tests = []struct {
		inFile       string
		outErrCode   int
		outErrString string
	}{
		{
			inFile:       "./testresource/err_unsupported_type.json",
			outErrCode:   http.StatusBadRequest,
			outErrString: "Bad Request: Column: id has unsupported type: unsupported",
		},
		{
			inFile:       "./testresource/err_unsupported_template.json",
			outErrCode:   http.StatusBadRequest,
			outErrString: "Bad Request: Unsupported template: unsupported",
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
			Body().Equal(test.outErrString)
	}
}
