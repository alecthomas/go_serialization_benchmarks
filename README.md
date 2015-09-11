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
- [github.com/philhofer/msgp](https://github.com/philhofer/msgp) *(code generator for msgpack)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://gogo.github.io/) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)
- [github.com/google/flatbuffers](https://github.com/google/flatbuffers)


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

Results with Go 1.3 on an Intel i7-3930K:

```
benchmark                                  iter        time/iter   bytes alloc          allocs
---------                                  ----        ---------   -----------          ------
BenchmarkVmihailencoMsgpackMarshal          1000000       2376 ns/op      352 B/op        5 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal        1000000       2380 ns/op      347 B/op       13 allocs/op
BenchmarkJsonMarshal                         500000       3501 ns/op      584 B/op        7 allocs/op
BenchmarkJsonUnmarshal                       300000       5553 ns/op      447 B/op        8 allocs/op
BenchmarkBsonMarshal                         500000       2347 ns/op      504 B/op       19 allocs/op
BenchmarkBsonUnmarshal                       500000       2596 ns/op      296 B/op       22 allocs/op
BenchmarkVitessBsonMarshal                  1000000       1735 ns/op     1168 B/op        4 allocs/op
BenchmarkVitessBsonUnmarshal                1000000       1036 ns/op      224 B/op        4 allocs/op
BenchmarkGobMarshal                          200000      10662 ns/op     1688 B/op       35 allocs/op
BenchmarkGobUnmarshal                         30000      49289 ns/op    17493 B/op      377 allocs/op
BenchmarkXdrMarshal                          500000       2821 ns/op      520 B/op       24 allocs/op
BenchmarkXdrUnmarshal                       1000000       2163 ns/op      271 B/op       12 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal           500000       4096 ns/op     1905 B/op       10 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal         500000       3770 ns/op     1840 B/op       14 allocs/op
BenchmarkUgorjiCodecBincMarshal              300000       4066 ns/op     1938 B/op       10 allocs/op
BenchmarkUgorjiCodecBincUnmarshal            500000       4004 ns/op     2000 B/op       17 allocs/op
BenchmarkSerealMarshal                       300000       4857 ns/op     1360 B/op       26 allocs/op
BenchmarkSerealUnmarshal                     500000       4287 ns/op      972 B/op       37 allocs/op
BenchmarkBinaryMarshal                      1000000       2440 ns/op      408 B/op       19 allocs/op
BenchmarkBinaryUnmarshal                    1000000       2339 ns/op      416 B/op       24 allocs/op
BenchmarkMsgpMarshal                        3000000        447 ns/op      144 B/op        1 allocs/op
BenchmarkMsgpUnmarshal                      3000000        584 ns/op      112 B/op        3 allocs/op
BenchmarkGoprotobufMarshal                  2000000        971 ns/op      312 B/op        4 allocs/op
BenchmarkGoprotobufUnmarshal                1000000       1220 ns/op      432 B/op        9 allocs/op
BenchmarkGogoprotobufMarshal               10000000        229 ns/op       64 B/op        1 allocs/op
BenchmarkGogoprotobufUnmarshal              5000000        323 ns/op      112 B/op        3 allocs/op
BenchmarkFlatbuffersMarshal                 3000000        523 ns/op        0 B/op        0 allocs/op
BenchmarkFlatBuffersUnmarshal              30000000       53.4 ns/op        0 B/op        0 allocs/op
BenchmarkProtobufMarshal                    1000000       1452 ns/op      232 B/op        9 allocs/op
BenchmarkProtobufUnmarshal                  1000000       1111 ns/op      192 B/op       10 allocs/op
```

**Note:** the gob results are not really representative of normal performance, as gob is designed for serializing streams or vectors of a single type, not individual values.


## Issues

The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

For compatibility we round the birthdays on milliseconds.

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
2. **(minor)** Vitess BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.
