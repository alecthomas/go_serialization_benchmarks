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

Results with Go 1.5 on a 2.6 GHz Intel Core i7 (MacBook Pro 15"):

```
benchmark                                   iter           time/iter      bytes alloc               allocs
---------                                   ----           ---------      -----------               ------
BenchmarkMsgpMarshal-8                      5000000        243 ns/op      128 B/op        1 allocs/op
BenchmarkMsgpUnmarshal-8                    3000000        459 ns/op      112 B/op        3 allocs/op
BenchmarkVmihailencoMsgpackMarshal-8        1000000       2347 ns/op      400 B/op        6 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal-8      1000000       2196 ns/op      400 B/op       13 allocs/op
BenchmarkJsonMarshal-8                       500000       2817 ns/op      584 B/op        7 allocs/op
BenchmarkJsonUnmarshal-8                     300000       4681 ns/op      463 B/op        8 allocs/op
BenchmarkBsonMarshal-8                      1000000       1744 ns/op      416 B/op       10 allocs/op
BenchmarkBsonUnmarshal-8                    1000000       2228 ns/op      416 B/op       21 allocs/op
BenchmarkVitessBsonMarshal-8                1000000       1052 ns/op     1168 B/op        4 allocs/op
BenchmarkVitessBsonUnmarshal-8              2000000        836 ns/op      224 B/op        4 allocs/op
BenchmarkGobMarshal-8                       1000000       1108 ns/op       48 B/op        2 allocs/op
BenchmarkGobUnmarshal-8                     1000000       1261 ns/op      160 B/op        6 allocs/op
BenchmarkXdrMarshal-8                       1000000       2033 ns/op      592 B/op       21 allocs/op
BenchmarkXdrUnmarshal-8                     1000000       1592 ns/op      287 B/op       11 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal-8         500000       2830 ns/op     2721 B/op        8 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal-8       500000       3172 ns/op     3104 B/op       12 allocs/op
BenchmarkUgorjiCodecBincMarshal-8            500000       2855 ns/op     2753 B/op        8 allocs/op
BenchmarkUgorjiCodecBincUnmarshal-8          500000       3367 ns/op     3264 B/op       15 allocs/op
BenchmarkSerealMarshal-8                     300000       3839 ns/op      928 B/op       21 allocs/op
BenchmarkSerealUnmarshal-8                   500000       3730 ns/op     1152 B/op       34 allocs/op
BenchmarkBinaryMarshal-8                    1000000       1835 ns/op      352 B/op       16 allocs/op
BenchmarkBinaryUnmarshal-8                  1000000       2019 ns/op      448 B/op       22 allocs/op
BenchmarkFlatbuffersMarshal-8               5000000        387 ns/op        0 B/op        0 allocs/op
BenchmarkFlatBuffersUnmarshal-8             5000000        315 ns/op      112 B/op        3 allocs/op
BenchmarkCapNProtoMarshal-8                 3000000        515 ns/op       64 B/op        2 allocs/op
BenchmarkCapNProtoUnmarshal-8               3000000        476 ns/op      216 B/op        6 allocs/op
BenchmarkCapNProto2Marshal-8                1000000       1820 ns/op      608 B/op       15 allocs/op
BenchmarkCapNProto2Unmarshal-8              1000000       1740 ns/op      640 B/op       16 allocs/op
BenchmarkHproseMarshal-8                    1000000       1201 ns/op      489 B/op        8 allocs/op
BenchmarkHproseUnmarshal-8                  1000000       1317 ns/op      319 B/op       10 allocs/op
BenchmarkProtobufMarshal-8                  1000000       1175 ns/op      224 B/op        7 allocs/op
BenchmarkProtobufUnmarshal-8                2000000        998 ns/op      240 B/op       10 allocs/op
BenchmarkGoprotobufMarshal-8                2000000        697 ns/op      320 B/op        4 allocs/op
BenchmarkGoprotobufUnmarshal-8              2000000       1045 ns/op      432 B/op        9 allocs/op
BenchmarkGogoprotobufMarshal-8             10000000        198 ns/op       64 B/op        1 allocs/op
BenchmarkGogoprotobufUnmarshal-8            5000000        303 ns/op      112 B/op        3 allocs/op
```

**Note:** the gob results are not really representative of normal performance, as gob is designed for serializing streams or vectors of a single type, not individual values.


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
