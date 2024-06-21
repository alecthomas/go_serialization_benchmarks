// This file is used to generate the report and is ignored by default on tests.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type reportLine struct {
	Name                  string `json:"name"`
	MarshalIterCount      int    `json:"marshal_iter_count"`
	UnmarshalIterCount    int    `json:"unmarshal_iter_count"`
	TotalIterCount        int    `json:"total_iter_count"`
	UnsafeStringUnmarshal bool   `json:"unsafe_string_unmarshal"`
	BufferReuseMarshal    bool   `json:"buffer_reuse_marshal"`
	MarshalNsOp           int64  `json:"marshal_ns_op"`
	UnmarshalNsOp         int64  `json:"unmarshal_ns_op"`
	TotalNsOp             int64  `json:"total_ns_op"`
	SerializationSize     int64  `json:"serialization_size"`
	MarshalAllocBytes     int64  `json:"marshal_alloc_bytes"`
	UnmarshalAllocBytes   int64  `json:"unmarshal_alloc_bytes"`
	TotalAllocBytes       int64  `json:"total_alloc_bytes"`
	MarshalAllocs         int64  `json:"marshal_allocs"`
	UnmarshalAllocs       int64  `json:"unmarshal_allocs"`
	TotalAllocs           int64  `json:"total_allocs"`
	TimeSupport           string `json:"time_support"`
	APIKind               string `json:"api_kind"`
	URL                   string `json:"url"`
	Notes                 string `json:"notes"`
}

func BenchAndReportSerializers(generateReport bool, validate bool) error {
	data := make([]reportLine, len(benchmarkCases))
	for i, bench := range benchmarkCases {
		marshalRes := testing.Benchmark(func(b *testing.B) {
			goserbench.BenchMarshalSmallStruct(b, bench.New())
		})
		fmt.Printf("%10s -   Marshal - %s %s\n", bench.Name, marshalRes.String(),
			marshalRes.MemString())
		unmarshalRes := testing.Benchmark(func(b *testing.B) {
			goserbench.BenchUnmarshalSmallStruct(b, bench.New(), validate)
		})
		fmt.Printf("%10s - Unmarshal - %s %s\n", bench.Name, unmarshalRes.String(),
			unmarshalRes.MemString())

		data[i] = reportLine{
			Name:                  bench.Name,
			MarshalIterCount:      marshalRes.N,
			UnmarshalIterCount:    unmarshalRes.N,
			TotalIterCount:        marshalRes.N + unmarshalRes.N,
			MarshalNsOp:           marshalRes.NsPerOp(),
			UnmarshalNsOp:         unmarshalRes.NsPerOp(),
			TotalNsOp:             marshalRes.NsPerOp() + unmarshalRes.NsPerOp(),
			UnsafeStringUnmarshal: bench.UnsafeStringUnmarshal,
			BufferReuseMarshal:    bench.BufferReuseMarshal,
			SerializationSize:     int64(marshalRes.Extra["B/serial"]),
			MarshalAllocBytes:     marshalRes.AllocedBytesPerOp(),
			UnmarshalAllocBytes:   unmarshalRes.AllocedBytesPerOp(),
			TotalAllocBytes: marshalRes.AllocedBytesPerOp() +
				unmarshalRes.AllocedBytesPerOp(),
			MarshalAllocs:   marshalRes.AllocsPerOp(),
			UnmarshalAllocs: unmarshalRes.AllocsPerOp(),
			TotalAllocs:     marshalRes.AllocsPerOp() + unmarshalRes.AllocsPerOp(),
			TimeSupport:     string(bench.TimeSupport),
			APIKind:         string(bench.APIKind),
			URL:             bench.URL,
			Notes:           strings.Join(bench.Notes, "\n"),
		}
	}

	if !generateReport {
		return nil
	}

	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	f, err := os.Create("report/data.js")
	if err != nil {
		return err
	}
	if _, err = f.Write([]byte("var data = ")); err != nil {
		return err
	}
	if _, err = f.Write(bytes); err != nil {
		return err
	}
	if _, err = f.Write([]byte(";")); err != nil {
		return err
	}

	fmt.Printf("\nSaved report to report/data.js !\n\n")
	return f.Close()
}
