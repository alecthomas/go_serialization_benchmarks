package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

func BenchmarkSerializers(b *testing.B) {
	fmt.Printf(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

Validating = %t

`, *validate)

	for i := range benchmarkCases {
		bc := benchmarkCases[i]
		b.Run("marshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchMarshalSmallStruct(b, bc.New())
		})
		b.Run("unmarshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchUnmarshalSmallStruct(b, bc.New(), *validate)
		})
	}
}
