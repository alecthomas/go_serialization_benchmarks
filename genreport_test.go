// Generate the report by running `go test -tags genreport`.

//go:build genreport

package main

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestGenerateReport(t *testing.T) {
	fmt.Printf(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

Validating = %t
Re-Generating Report = true

`, *validate)

	var nameRe *regexp.Regexp
	if *nameFlag != "" {
		var err error
		nameRe, err = regexp.Compile(*nameFlag)
		if err != nil {
			fmt.Printf("Error compiling -name regexp: %v\n", err)
			os.Exit(1)
		}
	}

	err := BenchAndReportSerializers(true, *validate, nameRe)
	if err != nil {
		t.Fatal(err)
	}
}
