# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

## Tested serialization methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [gopkg.in/vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson) *(using the bsongen code generator)*
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp) *(code generator for msgpack)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://gogo.github.io/) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)
- [github.com/google/flatbuffers](https://github.com/google/flatbuffers)
- [github.com/hprose/hprose-go/io](https://github.com/hprose/hprose-go)
- [github.com/glycerine/go-capnproto](https://github.com/glycerine/go-capnproto)
- [zombiezen.com/go/capnproto2](https://godoc.org/zombiezen.com/go/capnproto2)
- [github.com/andyleap/gencode](https://github.com/andyleap/gencode)

## Running the benchmarks

```bash
go get -u -t
go test -bench='.*' ./
```

Shameless plug: I use [pawk](https://github.com/alecthomas/pawk) to format the table:

```bash
go test -bench='.*' ./ | pawk -F'\t' '"%-40s %10s %10s %s %s" % f'
```

## Recommendation

If performance, correctness and interoperability are the most
important factors, [gogoprotobuf](https://gogo.github.io/) is
currently the best choice. It does require a pre-processing step (eg.
via Go 1.4's "go generate" command).

But as always, make your own choice based on your requirements.

## Data

The data being serialized is the following structure with randomly generated values:

```go
type A struct {
    Name     string
    BirthDay time.Time
    Phone    string
    Siblings int
    Spouse   bool
    Money    float64
}
```


## Results

Results with Go 1.5.3 on a 2.2 GHz Intel Core i7 (MacBook Pro 15"):

```
benchmark                                   iter           time/iter      bytes alloc               allocs
---------                                   ----           ---------      -----------               ------
BenchmarkMsgpMarshal-8                   5000000           277 ns/op         128 B/op          1 allocs/op
BenchmarkMsgpUnmarshal-8                 3000000           514 ns/op         112 B/op          3 allocs/op
BenchmarkVmihailencoMsgpackMarshal-8     1000000          1760 ns/op         352 B/op          5 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal-8    500000          2478 ns/op         400 B/op         13 allocs/op
BenchmarkJsonMarshal-8                    500000          3063 ns/op         536 B/op          7 allocs/op
BenchmarkJsonUnmarshal-8                  300000          4876 ns/op         447 B/op          8 allocs/op
BenchmarkBsonMarshal-8                   1000000          1970 ns/op         416 B/op         10 allocs/op
BenchmarkBsonUnmarshal-8                  500000          2608 ns/op         416 B/op         21 allocs/op
BenchmarkVitessBsonMarshal-8             1000000          1220 ns/op        1168 B/op          4 allocs/op
BenchmarkVitessBsonUnmarshal-8           2000000           946 ns/op         224 B/op          4 allocs/op
BenchmarkGobMarshal-8                    1000000          1254 ns/op          48 B/op          2 allocs/op
BenchmarkGobUnmarshal-8                  1000000          1450 ns/op         160 B/op          6 allocs/op
BenchmarkXdrMarshal-8                    1000000          2256 ns/op         527 B/op         19 allocs/op
BenchmarkXdrUnmarshal-8                  1000000          1798 ns/op         272 B/op         11 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal-8      500000          3235 ns/op        2721 B/op          8 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal-8    500000          3600 ns/op        3104 B/op         12 allocs/op
BenchmarkUgorjiCodecBincMarshal-8         500000          3274 ns/op        2753 B/op          8 allocs/op
BenchmarkUgorjiCodecBincUnmarshal-8       500000          3650 ns/op        3120 B/op         12 allocs/op
BenchmarkSerealMarshal-8                  300000          4564 ns/op         928 B/op         21 allocs/op
BenchmarkSerealUnmarshal-8                300000          4550 ns/op        1152 B/op         34 allocs/op
BenchmarkBinaryMarshal-8                 1000000          1967 ns/op         352 B/op         16 allocs/op
BenchmarkBinaryUnmarshal-8               1000000          2109 ns/op         448 B/op         22 allocs/op
BenchmarkFlatbuffersMarshal-8            3000000           411 ns/op           0 B/op          0 allocs/op
BenchmarkFlatBuffersUnmarshal-8          5000000           366 ns/op         112 B/op          3 allocs/op
BenchmarkCapNProtoMarshal-8              2000000           594 ns/op          64 B/op          2 allocs/op
BenchmarkCapNProtoUnmarshal-8            3000000           535 ns/op         216 B/op          6 allocs/op
BenchmarkCapNProto2Marshal-8             1000000          2054 ns/op         608 B/op         15 allocs/op
BenchmarkCapNProto2Unmarshal-8           1000000          1943 ns/op         640 B/op         16 allocs/op
BenchmarkHproseMarshal-8                 1000000          1308 ns/op         489 B/op          8 allocs/op
BenchmarkHproseUnmarshal-8               1000000          1506 ns/op         320 B/op         10 allocs/op
BenchmarkProtobufMarshal-8               1000000          1361 ns/op         224 B/op          7 allocs/op
BenchmarkProtobufUnmarshal-8             1000000          1072 ns/op         240 B/op         10 allocs/op
BenchmarkGoprotobufMarshal-8             2000000           764 ns/op         320 B/op          4 allocs/op
BenchmarkGoprotobufUnmarshal-8           1000000          1044 ns/op         432 B/op          9 allocs/op
BenchmarkGogoprotobufMarshal-8          10000000           214 ns/op          64 B/op          1 allocs/op
BenchmarkGogoprotobufUnmarshal-8         5000000           306 ns/op          96 B/op          3 allocs/op
BenchmarkGencodeMarshal-8               10000000           217 ns/op          80 B/op          2 allocs/op
BenchmarkGencodeUnmarshal-8              5000000           266 ns/op         112 B/op          3 allocs/op
```

## Issues


The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
2. **(minor)** Vitess BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
--- FAIL: BenchmarkBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20b999e3621bd773 2016-01-19 14:05:02.469416459 -0800 PST f017c8e9de 4 true 0.20887343719329818}
        &{20b999e3621bd773 2016-01-19 14:05:02.469 -0800 PST f017c8e9de 4 true 0.20887343719329818}
--- FAIL: BenchmarkVitessBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{64204cc8fe408156 2016-01-19 14:05:04.068418034 -0800 PST 95d04e2d49 3 false 0.011931519836294776}
        &{64204cc8fe408156 2016-01-19 22:05:04.068 +0000 UTC 95d04e2d49 3 false 0.011931519836294776}
--- FAIL: BenchmarkUgorjiCodecBincUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 PST 71f3bf4233 0 false 0.8712180830484527}
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 -0800 71f3bf4233 0 false 0.8712180830484527}
```

All other fields are correct however.

Additionally, while not a correctness issue, FlatBuffers, ProtoBuffers and Cap'N'Proto do not
support time types directly. In the benchmarks an int64 value is used to hold a UnixNano timestamp.
