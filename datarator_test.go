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
		// BaseURL: "http://localhost:2345",
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
