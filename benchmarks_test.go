package goserbench

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

var (
	validate = os.Getenv("VALIDATE") != ""
)

func TestMessage(t *testing.T) {
	fmt.Print(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

`)
}

func BenchmarkSerializers(b *testing.B) {
	for i := range benchmarkCases {
		bc := benchmarkCases[i]
		if ALG_NAME != "any" && bc.Name != ALG_NAME {
			continue
		}
		b.Run("marshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchMarshalSmallStruct(b, bc.New())
		})
		b.Run("unmarshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchUnmarshalSmallStruct(b, bc.New(), validate)
		})
	}
}
