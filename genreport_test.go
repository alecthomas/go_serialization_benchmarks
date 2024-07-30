// Generate the report by running `go test -tags genreport`.

//go:build genreport

package main

import (
	"fmt"
	"testing"
)

func TestGenerateReport(t *testing.T) {
	fmt.Printf(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

Validating = %t
Re-Generating Report = true

`, *validate)

	err := BenchAndReportSerializers(true, *validate)
	if err != nil {
		t.Fatal(err)
	}
}
