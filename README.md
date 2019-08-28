# Benchmarks of Go serialization methods

[![Gitter chat](https://badges.gitter.im/alecthomas.png)](https://gitter.im/alecthomas/Lobby)

This is a test suite for benchmarking various Go serialization methods.

## Tested serialization methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/json-iterator/go](https://github.com/json-iterator/go)
- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [gopkg.in/vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp) *(code generator for msgpack)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://github.com/gogo/protobuf) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)
- [github.com/google/flatbuffers](https://github.com/google/flatbuffers)
- [github.com/hprose/hprose-go/io](https://github.com/hprose/hprose-go)
- [github.com/glycerine/go-capnproto](https://github.com/glycerine/go-capnproto)
- [zombiezen.com/go/capnproto2](https://godoc.org/zombiezen.com/go/capnproto2)
- [github.com/andyleap/gencode](https://github.com/andyleap/gencode)
- [github.com/pascaldekloe/colfer](https://github.com/pascaldekloe/colfer)
- [github.com/linkedin/goavro](https://github.com/linkedin/goavro)
- [github.com/ikkerens/ikeapack](https://github.com/ikkerens/ikeapack)
- [github.com/niubaoshu/gotiny](https://github.com/niubaoshu/gotiny)
- [github.com/prysmaticlabs/go-ssz](https://github.com/prysmaticlabs/go-ssz)

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

2019-07-30 Results with Go 1.12.7 linux/amd64 on a google cloud n1-standard-1 (1 vCPU, 3.75 GB memory):

```
benchmark                                           iter               time/iter             bytes/op            allocs/op ---------                                       --------               ---------             --------          ----------- BenchmarkGotinyMarshal                          10000000               128 ns/op               0 B/op          0 allocs/op
BenchmarkGotinyUnmarshal                         5000000               271 ns/op             112 B/op          3 allocs/op
BenchmarkGotinyNoTimeMarshal                    10000000               127 ns/op               0 B/op          0 allocs/op
BenchmarkGotinyNoTimeUnmarshal                   5000000               250 ns/op              96 B/op          3 allocs/op
BenchmarkMsgpMarshal                            10000000               217 ns/op             128 B/op          1 allocs/op
BenchmarkMsgpUnmarshal                           5000000               381 ns/op             112 B/op          3 allocs/op
BenchmarkVmihailencoMsgpackMarshal               1000000              1956 ns/op             368 B/op          7 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal             1000000              2170 ns/op             384 B/op         13 allocs/op
BenchmarkJsonMarshal                             1000000              2106 ns/op             303 B/op          4 allocs/op
BenchmarkJsonUnmarshal                            300000              4378 ns/op             343 B/op          7 allocs/op
BenchmarkJsonIterMarshal                         1000000              2323 ns/op             296 B/op          5 allocs/op
BenchmarkJsonIterUnmarshal                       1000000              1755 ns/op             184 B/op          6 allocs/op
BenchmarkEasyJsonMarshal                         1000000              1615 ns/op             783 B/op          5 allocs/op
BenchmarkEasyJsonUnmarshal                       1000000              1529 ns/op             144 B/op          4 allocs/op
BenchmarkBsonMarshal                             1000000              1540 ns/op             392 B/op         10 allocs/op
BenchmarkBsonUnmarshal                           1000000              2101 ns/op             244 B/op         19 allocs/op
BenchmarkGobMarshal                              1000000              1026 ns/op              48 B/op          2 allocs/op
BenchmarkGobUnmarshal                            1000000              1073 ns/op             112 B/op          3 allocs/op
BenchmarkXdrMarshal                              1000000              2086 ns/op             392 B/op         19 allocs/op
BenchmarkXdrUnmarshal                            1000000              1572 ns/op             224 B/op         11 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal               1000000              1546 ns/op            1312 B/op          3 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal             1000000              1439 ns/op             496 B/op          4 allocs/op
BenchmarkUgorjiCodecBincMarshal                  1000000              1667 ns/op            1328 B/op          4 allocs/op
BenchmarkUgorjiCodecBincUnmarshal                1000000              1452 ns/op             496 B/op          4 allocs/op
BenchmarkSerealMarshal                            500000              3269 ns/op             904 B/op         20 allocs/op
BenchmarkSerealUnmarshal                          500000              3589 ns/op            1008 B/op         34 allocs/op
BenchmarkBinaryMarshal                           1000000              1609 ns/op             326 B/op         21 allocs/op
BenchmarkBinaryUnmarshal                         1000000              1632 ns/op             320 B/op         22 allocs/op
BenchmarkFlatBuffersMarshal                      5000000               321 ns/op               0 B/op          0 allocs/op
BenchmarkFlatBuffersUnmarshal                    5000000               295 ns/op             112 B/op          3 allocs/op
BenchmarkCapNProtoMarshal                        3000000               445 ns/op              56 B/op          2 allocs/op
BenchmarkCapNProtoUnmarshal                      3000000               503 ns/op             200 B/op          6 allocs/op
BenchmarkCapNProto2Marshal                       2000000               734 ns/op             244 B/op          3 allocs/op
BenchmarkCapNProto2Unmarshal                     1000000              1101 ns/op             320 B/op          6 allocs/op
BenchmarkHproseMarshal                           1000000              1150 ns/op             473 B/op          8 allocs/op
BenchmarkHproseUnmarshal                         1000000              1325 ns/op             319 B/op         10 allocs/op
BenchmarkHprose2Marshal                          3000000               592 ns/op               0 B/op          0 allocs/op
BenchmarkHprose2Unmarshal                        2000000               665 ns/op             144 B/op          4 allocs/op
BenchmarkProtobufMarshal                         1000000              1011 ns/op             152 B/op          7 allocs/op
BenchmarkProtobufUnmarshal                       2000000               932 ns/op             192 B/op         10 allocs/op
BenchmarkGoprotobufMarshal                       3000000               450 ns/op              96 B/op          2 allocs/op
BenchmarkGoprotobufUnmarshal                     2000000               776 ns/op             200 B/op         10 allocs/op
BenchmarkGogoprotobufMarshal                    10000000               200 ns/op              64 B/op          1 allocs/op
BenchmarkGogoprotobufUnmarshal                   5000000               269 ns/op              96 B/op          3 allocs/op
BenchmarkColferMarshal                          10000000               160 ns/op              64 B/op          1 allocs/op
BenchmarkColferUnmarshal                        10000000               229 ns/op             112 B/op          3 allocs/op
BenchmarkGencodeMarshal                         10000000               216 ns/op              80 B/op          2 allocs/op
BenchmarkGencodeUnmarshal                       10000000               229 ns/op             112 B/op          3 allocs/op
BenchmarkGencodeUnsafeMarshal                   10000000               132 ns/op              48 B/op          1 allocs/op
BenchmarkGencodeUnsafeUnmarshal                 10000000               183 ns/op              96 B/op          3 allocs/op
BenchmarkXDR2Marshal                            10000000               205 ns/op              64 B/op          1 allocs/op
BenchmarkXDR2Unmarshal                          10000000               179 ns/op              32 B/op          2 allocs/op
BenchmarkGoAvroMarshal                            500000              3433 ns/op            1030 B/op         32 allocs/op
BenchmarkGoAvroUnmarshal                          200000              7986 ns/op            3372 B/op         87 allocs/op
BenchmarkGoAvro2TextMarshal                       500000              3438 ns/op            1326 B/op         20 allocs/op
BenchmarkGoAvro2TextUnmarshal                     500000              3346 ns/op             806 B/op         33 allocs/op
BenchmarkGoAvro2BinaryMarshal                    1000000              1148 ns/op             510 B/op         11 allocs/op
BenchmarkGoAvro2BinaryUnmarshal                  1000000              1204 ns/op             576 B/op         13 allocs/op
BenchmarkIkeaMarshal                             2000000               709 ns/op              72 B/op          8 allocs/op
BenchmarkIkeaUnmarshal                           2000000               926 ns/op             160 B/op         11 allocs/op
BenchmarkShamatonMapMsgpackMarshal               2000000               911 ns/op             208 B/op          4 allocs/op
BenchmarkShamatonMapMsgpackUnmarshal             2000000               817 ns/op             144 B/op          3 allocs/op
BenchmarkShamatonArrayMsgpackMarshal             2000000               790 ns/op             176 B/op          4 allocs/op
BenchmarkShamatonArrayMsgpackUnmarshal           3000000               559 ns/op             144 B/op          3 allocs/op
BenchmarkSSZNoTimeNoStringNoFloatAMarshal         500000              2932 ns/op            1216 B/op         18 allocs/op
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal      1000000              1048 ns/op             216 B/op          9 allocs/op
```

## Issues


The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
--- FAIL: BenchmarkBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20b999e3621bd773 2016-01-19 14:05:02.469416459 -0800 PST f017c8e9de 4 true 0.20887343719329818}
        &{20b999e3621bd773 2016-01-19 14:05:02.469 -0800 PST f017c8e9de 4 true 0.20887343719329818}
--- FAIL: BenchmarkUgorjiCodecBincUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 PST 71f3bf4233 0 false 0.8712180830484527}
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 -0800 71f3bf4233 0 false 0.8712180830484527}
```

All other fields are correct however.

Additionally, while not a correctness issue, FlatBuffers, ProtoBuffers, Cap'N'Proto and ikeapack do not
support time types directly. In the benchmarks an int64 value is used to hold a UnixNano timestamp.
