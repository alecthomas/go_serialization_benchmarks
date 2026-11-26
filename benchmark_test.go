package main

import (
	"fmt"
	"testing"
)

func BenchmarkSerializers(b *testing.B) {
	fmt.Printf(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

Validating = %t

`, *validate)

	for i := range benchmarkCases {
		bc := benchmarkCases[i]
		b.Run("marshal/"+bc.Name, bc.Bench.Marshal)
		b.Run("unmarshal/"+bc.Name, func(b *testing.B) {
			bc.Bench.Unmarshal(b, *validate)
		})
	}
}
