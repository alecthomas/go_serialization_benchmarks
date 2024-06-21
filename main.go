package main

import (
	"flag"
	"fmt"
)

var (
	validate  = flag.Bool("validate", false, "to validate the correctness of the serializers")
	genReport = flag.Bool("genreport", false, "to re-generate the report")
)

func main() {
	flag.Parse()

	fmt.Printf(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

Validating = %t
Re-Generating Report = %t

`, *validate, *genReport)

	err := BenchAndReportSerializers(*genReport, *validate)
	if err != nil {
		panic(err)
	}
}
