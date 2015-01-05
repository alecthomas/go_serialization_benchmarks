# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

## Tested serialization methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)
- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson) *(using the bsongen code generator)*
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/philhofer/msgp](https://github.com/philhofer/msgp) *(code generator for msgpack)*
- [code.google.com/p/goprotobuf/proto](https://code.google.com/p/goprotobuf/proto) (generated code)
- [github.com/gogo/protobuf](https://gogo.github.io/) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)


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

If performance, correctness and interoperability are the most important
factors, [github.com/philhofer/msgp](https://github.com/philhofer/msgp) is currently
the best choice. It does require a pre-processing step (eg. via Go 1.4's "go
generate" command).

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

Results with Go 1.3 on an Intel i7-3930K:

```
benchmark                                  iter        time/iter   bytes alloc          allocs
---------                                  ----        ---------   -----------          ------
BenchmarkVmihailencoMsgpackMarshal      1000000    2005.00 ns/op      413 B/op     6 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal    1000000    2543.00 ns/op      421 B/op    10 allocs/op
BenchmarkJsonMarshal                     500000    4374.00 ns/op      590 B/op     7 allocs/op
BenchmarkJsonUnmarshal                   500000    7055.00 ns/op      468 B/op     7 allocs/op
BenchmarkBsonMarshal                    1000000    2516.00 ns/op      489 B/op    13 allocs/op
BenchmarkBsonUnmarshal                  1000000    2784.00 ns/op      282 B/op    10 allocs/op
BenchmarkVitessBsonMarshal              1000000    1770.00 ns/op     1169 B/op     4 allocs/op
BenchmarkVitessBsonUnmarshal            2000000     930.00 ns/op      227 B/op     4 allocs/op
BenchmarkGobMarshal                      200000    8277.00 ns/op     1672 B/op    25 allocs/op
BenchmarkGobUnmarshal                     50000   59502.00 ns/op    19060 B/op   365 allocs/op
BenchmarkXdrMarshal                      500000    3296.00 ns/op      520 B/op    15 allocs/op
BenchmarkXdrUnmarshal                   1000000    2439.00 ns/op      274 B/op     9 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal       500000    4321.00 ns/op     1527 B/op    23 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal     500000    4037.00 ns/op     1159 B/op    23 allocs/op
BenchmarkUgorjiCodecBincMarshal          500000    4403.00 ns/op     1544 B/op    23 allocs/op
BenchmarkUgorjiCodecBincUnmarshal        500000    4257.00 ns/op     1303 B/op    25 allocs/op
BenchmarkSerealMarshal                   500000    5733.00 ns/op     1372 B/op    24 allocs/op
BenchmarkSerealUnmarshal                 500000    4970.00 ns/op      998 B/op    24 allocs/op
BenchmarkBinaryMarshal                  1000000    2384.00 ns/op      405 B/op    12 allocs/op
BenchmarkBinaryUnmarshal                1000000    2768.00 ns/op      433 B/op    17 allocs/op
BenchmarkMsgpMarshal                    5000000     466.00 ns/op      144 B/op     1 allocs/op
BenchmarkMsgpUnmarshal                  5000000     601.00 ns/op      113 B/op     3 allocs/op
BenchmarkGoprotobufMarshal              2000000     988.00 ns/op      314 B/op     3 allocs/op
BenchmarkGoprotobufUnmarshal            1000000    1376.00 ns/op      440 B/op     9 allocs/op
BenchmarkGogoprotobufMarshal           10000000     250.00 ns/op       64 B/op     1 allocs/op
BenchmarkGogoprotobufUnmarshal          5000000     349.00 ns/op      113 B/op     3 allocs/op
BenchmarkProtobufMarshal                1000000    1399.00 ns/op      213 B/op     6 allocs/op
BenchmarkProtobufUnmarshal              1000000    1371.00 ns/op      193 B/op     7 allocs/op
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
BenchmarkVmihailencoMsgpackMarshal   1000000          1770 ns/op         408 B/op          8 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal      500000          3073 ns/op         512 B/op         15 allocs/op
BenchmarkJsonMarshal      500000          3292 ns/op         584 B/op          7 allocs/op
BenchmarkJsonUnmarshal    200000          6077 ns/op         543 B/op         10 allocs/op
BenchmarkBsonMarshal     1000000          2302 ns/op         504 B/op         19 allocs/op
BenchmarkBsonUnmarshal  --- FAIL: BenchmarkBsonUnmarshal
    serialization_benchmarks_test.go:299: unmarshaled object differed:
        &{36f9dfc8311e1efd 2014-12-20 08:28:56.697060419 +1100 AEDT bce6c01982 0 true 0.9078469433531421}
        &{36f9dfc8311e1efd 2014-12-20 08:28:56.697 +1100 AEDT bce6c01982 0 true 0.9078469433531421}
BenchmarkVitessBsonMarshal   1000000          1644 ns/op        1168 B/op          4 allocs/op
BenchmarkVitessBsonUnmarshal    --- FAIL: BenchmarkVitessBsonUnmarshal
    serialization_benchmarks_test.go:299: unmarshaled object differed:
        &{d22635208bfe823f 2014-12-20 08:28:58.370539555 +1100 AEDT 6c6e6d4ef8 2 true 0.2026063151373092}
        &{d22635208bfe823f 2014-12-19 21:28:58.37 +0000 UTC 6c6e6d4ef8 2 true 0.2026063151373092}
BenchmarkGobMarshal   200000          9770 ns/op        1688 B/op         35 allocs/op
BenchmarkGobUnmarshal      30000         47695 ns/op       17590 B/op        379 allocs/op
BenchmarkXdrMarshal   500000          2559 ns/op         519 B/op         24 allocs/op
BenchmarkXdrUnmarshal     500000          3016 ns/op         352 B/op         14 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal    500000          3662 ns/op        1464 B/op         24 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal      300000          4533 ns/op        1232 B/op         29 allocs/op
BenchmarkUgorjiCodecBincMarshal   500000          3745 ns/op        1480 B/op         24 allocs/op
BenchmarkUgorjiCodecBincUnmarshal   --- FAIL: BenchmarkUgorjiCodecBincUnmarshal
    serialization_benchmarks_test.go:299: unmarshaled object differed:
        &{94fd6da7835a0346 2014-12-20 08:29:10.50608751 +1100 AEDT 7749978493 0 true 0.09633537464986415}
        &{94fd6da7835a0346 2014-12-20 08:29:10.50608751 +1100 +1100 7749978493 0 true 0.09633537464986415}
BenchmarkSerealMarshal    300000          4572 ns/op        1360 B/op         26 allocs/op
BenchmarkSerealUnmarshal      300000          4950 ns/op        1068 B/op         39 allocs/op
BenchmarkBinaryMarshal   1000000          2168 ns/op         408 B/op         19 allocs/op
BenchmarkBinaryUnmarshal      500000          3249 ns/op         512 B/op         26 allocs/op
BenchmarkGoprotobufMarshal   2000000           928 ns/op         312 B/op          4 allocs/op
BenchmarkGoprotobufUnmarshal     1000000          1201 ns/op         432 B/op          9 allocs/op
BenchmarkProtobufMarshal     1000000          1342 ns/op         232 B/op          9 allocs/op
BenchmarkProtobufUnmarshal   1000000          2128 ns/op         288 B/op         12 allocs/op
```

All other fields are correct however.
