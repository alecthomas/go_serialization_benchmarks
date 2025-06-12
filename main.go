package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"testing"
)

var (
	validate  = flag.Bool("validate", false, "to validate the correctness of the serializers")
	genReport = flag.Bool("genreport", false, "to re-generate the report")
	nameFlag  = flag.String("name", "", "restrict benchmarks to run only if they match this regexp")
)

func main() {
	flag.Parse()

	fmt.Printf(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

Validating = %t
Re-Generating Report = %t

`, *validate, *genReport)

	var nameRe *regexp.Regexp
	if *nameFlag != "" {
		var err error
		nameRe, err = regexp.Compile(*nameFlag)
		if err != nil {
			fmt.Printf("Error compiling -name regexp: %v\n", err)
			os.Exit(1)
		}
	}

	// Manually call testing.Init(), otherwise validation errors may cause
	// internal panics.
	testing.Init()

	err := BenchAndReportSerializers(*genReport, *validate, nameRe)
	if err != nil {
		panic(err)
	}
}
