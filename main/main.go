package main

import (
	"flag"
	"os"

	bench "github.com/niubaoshu/go_serialization_benchmarks"
)

var flagSet = flag.NewFlagSet("my", 0)
var validate = flagSet.Bool("v", false, "validate")
var match = flagSet.String("m", ".*", "filter name")

func main() {
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return
	}
	bench.B.Bench(*validate, *match)
}
