package goserbench

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"go4.org/sort"
)

type (
	serializeBenchResault struct {
		marshalResult   testing.BenchmarkResult
		unmarshalResult testing.BenchmarkResult
		length          int
		serializer      Serializer
		merr            error
		unmerr          error
	}

	Serializer interface {
		Marshal(o interface{}) ([]byte, error)
		Unmarshal(d []byte, o interface{}) error
		String() string
		NeedGeneratingCode() bool
	}

	Object interface {
		AssertEqual(interface{}) (bool, error)
		Reset()
	}

	benchSerializer struct {
		s   Serializer
		new func() Object
	}
	BenchSerializer struct {
		ss     []benchSerializer
		result []*serializeBenchResault
		maxlen int
	}
)

func NewBenchSerializer() *BenchSerializer {
	return &BenchSerializer{}
}

func (bs *BenchSerializer) AddSerializer(s Serializer, new func() Object) {
	bs.ss = append(bs.ss, benchSerializer{
		s:   s,
		new: new,
	})

}

var cupnum = strconv.Itoa(runtime.NumCPU())

func (s *serializeBenchResault) String(l int) string {
	marshal := []string{fmt.Sprintf("%-*s\t", l, s.serializer.String()+"-Marshal-"+cupnum)}
	unmarshal := []string{fmt.Sprintf("%-*s\t", l, s.serializer.String()+"-Unmarshal-"+cupnum)}

	//marshal := []string{}
	//unmarshal := []string{}
	var length = s.length / s.marshalResult.N
	if s.merr != nil {
		marshal = append(marshal, s.merr.Error())
	} else {
		marshal = append(marshal, benchResultToString(&s.marshalResult, length, s.serializer.NeedGeneratingCode())...)
	}

	if s.unmerr != nil {
		unmarshal = append(unmarshal, s.unmerr.Error())
	} else {
		unmarshal = append(unmarshal, benchResultToString(&s.unmarshalResult, length, s.serializer.NeedGeneratingCode())...)
	}
	return strings.Join(marshal, "\t") + "\n" + strings.Join(unmarshal, "\t")
}

func benchResultToString(result *testing.BenchmarkResult, length int, generating bool) []string {
	var arr []string
	strs := strings.Split(result.String(), "\t")
	arr = append(arr, strs[1])
	if result.N != 0 {
		arr = append(arr, fmt.Sprintf("%5d byte", length))
	} else {
		arr = append(arr, fmt.Sprintf("%5d byte", 0))
	}
	memstrs := strings.Split(result.MemString(), "\t ")
	arr = append(arr, memstrs...)
	if generating {
		arr = append(arr, "    YES     ")
	} else {
		arr = append(arr, "    NO      ")
	}
	arr = append(arr, strs[0])
	return arr
}

func (bs *BenchSerializer) Bench(validate bool, match string) {
	reg := regexp.MustCompile(match)
	for _, s := range bs.ss {
		if reg.MatchString(s.s.String()) {
			maxlen := len(s.s.String() + "-Unmarshal-" + cupnum)
			if maxlen > bs.maxlen {
				bs.maxlen = maxlen
			}
		}
	}
	w := os.Stdout
	fmt.Fprintf(w, "goos: %s\n", runtime.GOOS)
	fmt.Fprintf(w, "goarch: %s\n", runtime.GOARCH)
	fmt.Fprintf(w, "%-*s\t", bs.maxlen, "name")
	fmt.Fprintln(w, "               time/iter \t  lenth/iter \t     bytes/iter       allocs/iter \t generate \t iter")
	fmt.Fprintf(w, "%-*s\t", bs.maxlen, "----")
	fmt.Fprintln(w, "               --------- \t  ---------- \t     ----------       ----------- \t -------- \t ----")
	for _, s := range bs.ss {
		if reg.MatchString(s.s.String()) {
			result := bench(s.s, s.new, validate)
			bs.result = append(bs.result, result)
			fmt.Fprintln(w, result.String(bs.maxlen))
		}
	}
	var result = bs.result
	sort.Slice(result, func(i, j int) bool {
		return bs.result[i].marshalResult.NsPerOp()+bs.result[i].unmarshalResult.NsPerOp() < bs.result[j].marshalResult.NsPerOp()+bs.result[j].unmarshalResult.NsPerOp()
	})

	fmt.Fprintln(w, "\nsort:")
	for _, r := range result {
		fmt.Fprintln(w, r.String(bs.maxlen))
	}

}
