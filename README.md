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

To update the table in the README:

```bash
./stats.sh
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

2019-08-28 Results with Go 1.12.6 darwin/amd64 on a 2.8 GHz Intel Core i7 16GB

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.time | tt.bytes     | time/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGotinyMarshal-8                 |   20000000 |     99 ns/op |     0 B/op |   0 allocs/op |   1.99 s |       0 KB |    0.00 ns/alloc
BenchmarkGotinyUnmarshal-8               |   10000000 |    192 ns/op |   112 B/op |   3 allocs/op |   1.92 s |  112000 KB |   64.00 ns/alloc
BenchmarkGotinyNoTimeMarshal-8           |   20000000 |     99 ns/op |     0 B/op |   0 allocs/op |   1.98 s |       0 KB |    0.00 ns/alloc
BenchmarkGotinyNoTimeUnmarshal-8         |   10000000 |    175 ns/op |    96 B/op |   3 allocs/op |   1.75 s |   96000 KB |   58.33 ns/alloc
BenchmarkMsgpMarshal-8                   |   10000000 |    140 ns/op |   128 B/op |   1 allocs/op |   1.40 s |  128000 KB |  140.00 ns/alloc
BenchmarkMsgpUnmarshal-8                 |    5000000 |    280 ns/op |   112 B/op |   3 allocs/op |   1.40 s |   56000 KB |   93.33 ns/alloc
BenchmarkVmihailencoMsgpackMarshal-8     |    2000000 |    986 ns/op |   400 B/op |   7 allocs/op |   1.97 s |   80000 KB |  140.86 ns/alloc
BenchmarkVmihailencoMsgpackUnmarshal-8   |    1000000 |   1271 ns/op |   416 B/op |  13 allocs/op |   1.27 s |   41600 KB |   97.77 ns/alloc
BenchmarkJsonMarshal-8                   |    1000000 |   1585 ns/op |   304 B/op |   4 allocs/op |   1.58 s |   30400 KB |  396.25 ns/alloc
BenchmarkJsonUnmarshal-8                 |     500000 |   3307 ns/op |   359 B/op |   7 allocs/op |   1.65 s |   17950 KB |  472.43 ns/alloc
BenchmarkJsonIterMarshal-8               |    1000000 |   1768 ns/op |   312 B/op |   5 allocs/op |   1.77 s |   31200 KB |  353.60 ns/alloc
BenchmarkJsonIterUnmarshal-8             |    1000000 |   1361 ns/op |   263 B/op |   6 allocs/op |   1.36 s |   26300 KB |  226.83 ns/alloc
BenchmarkEasyJsonMarshal-8               |    1000000 |   1125 ns/op |   784 B/op |   5 allocs/op |   1.12 s |   78400 KB |  225.00 ns/alloc
BenchmarkEasyJsonUnmarshal-8             |    1000000 |   1208 ns/op |   159 B/op |   4 allocs/op |   1.21 s |   15900 KB |  302.00 ns/alloc
BenchmarkBsonMarshal-8                   |    1000000 |   1005 ns/op |   392 B/op |  10 allocs/op |   1.00 s |   39200 KB |  100.50 ns/alloc
BenchmarkBsonUnmarshal-8                 |    1000000 |   1559 ns/op |   244 B/op |  19 allocs/op |   1.56 s |   24400 KB |   82.05 ns/alloc
BenchmarkGobMarshal-8                    |    2000000 |    744 ns/op |    48 B/op |   2 allocs/op |   1.49 s |    9600 KB |  372.00 ns/alloc
BenchmarkGobUnmarshal-8                  |    2000000 |    793 ns/op |   112 B/op |   3 allocs/op |   1.59 s |   22400 KB |  264.33 ns/alloc
BenchmarkXdrMarshal-8                    |    1000000 |   1332 ns/op |   392 B/op |  19 allocs/op |   1.33 s |   39200 KB |   70.11 ns/alloc
BenchmarkXdrUnmarshal-8                  |    1000000 |   1200 ns/op |   224 B/op |  11 allocs/op |   1.20 s |   22400 KB |  109.09 ns/alloc
BenchmarkUgorjiCodecMsgpackMarshal-8     |    2000000 |    922 ns/op |  1312 B/op |   3 allocs/op |   1.84 s |  262400 KB |  307.33 ns/alloc
BenchmarkUgorjiCodecMsgpackUnmarshal-8   |    1000000 |   1036 ns/op |   496 B/op |   4 allocs/op |   1.04 s |   49600 KB |  259.00 ns/alloc
BenchmarkUgorjiCodecBincMarshal-8        |    1000000 |   1019 ns/op |  1328 B/op |   4 allocs/op |   1.02 s |  132800 KB |  254.75 ns/alloc
BenchmarkUgorjiCodecBincUnmarshal-8      |    1000000 |   1179 ns/op |   640 B/op |   7 allocs/op |   1.18 s |   64000 KB |  168.43 ns/alloc
BenchmarkSerealMarshal-8                 |    1000000 |   2192 ns/op |   904 B/op |  20 allocs/op |   2.19 s |   90400 KB |  109.60 ns/alloc
BenchmarkSerealUnmarshal-8               |     500000 |   2634 ns/op |  1008 B/op |  34 allocs/op |   1.32 s |   50400 KB |   77.47 ns/alloc
BenchmarkBinaryMarshal-8                 |    1000000 |   1207 ns/op |   326 B/op |  21 allocs/op |   1.21 s |   32600 KB |   57.48 ns/alloc
BenchmarkBinaryUnmarshal-8               |    1000000 |   1218 ns/op |   320 B/op |  22 allocs/op |   1.22 s |   32000 KB |   55.36 ns/alloc
BenchmarkFlatBuffersMarshal-8            |    5000000 |    266 ns/op |     0 B/op |   0 allocs/op |   1.33 s |       0 KB |    0.00 ns/alloc
BenchmarkFlatBuffersUnmarshal-8          |   10000000 |    212 ns/op |   112 B/op |   3 allocs/op |   2.12 s |  112000 KB |   70.67 ns/alloc
BenchmarkCapNProtoMarshal-8              |    5000000 |    369 ns/op |    56 B/op |   2 allocs/op |   1.84 s |   28000 KB |  184.50 ns/alloc
BenchmarkCapNProtoUnmarshal-8            |    5000000 |    355 ns/op |   200 B/op |   6 allocs/op |   1.77 s |  100000 KB |   59.17 ns/alloc
BenchmarkCapNProto2Marshal-8             |    3000000 |    571 ns/op |   244 B/op |   3 allocs/op |   1.71 s |   73200 KB |  190.33 ns/alloc
BenchmarkCapNProto2Unmarshal-8           |    2000000 |    792 ns/op |   320 B/op |   6 allocs/op |   1.58 s |   64000 KB |  132.00 ns/alloc
BenchmarkHproseMarshal-8                 |    2000000 |    772 ns/op |   331 B/op |   8 allocs/op |   1.54 s |   66200 KB |   96.50 ns/alloc
BenchmarkHproseUnmarshal-8               |    2000000 |    959 ns/op |   319 B/op |  10 allocs/op |   1.92 s |   63800 KB |   95.90 ns/alloc
BenchmarkHprose2Marshal-8                |    3000000 |    475 ns/op |     0 B/op |   0 allocs/op |   1.43 s |       0 KB |    0.00 ns/alloc
BenchmarkHprose2Unmarshal-8              |    3000000 |    491 ns/op |   144 B/op |   4 allocs/op |   1.47 s |   43200 KB |  122.75 ns/alloc
BenchmarkProtobufMarshal-8               |    2000000 |    738 ns/op |   152 B/op |   7 allocs/op |   1.48 s |   30400 KB |  105.43 ns/alloc
BenchmarkProtobufUnmarshal-8             |    2000000 |    686 ns/op |   192 B/op |  10 allocs/op |   1.37 s |   38400 KB |   68.60 ns/alloc
BenchmarkGoprotobufMarshal-8             |    5000000 |    337 ns/op |    96 B/op |   2 allocs/op |   1.69 s |   48000 KB |  168.50 ns/alloc
BenchmarkGoprotobufUnmarshal-8           |    3000000 |    533 ns/op |   200 B/op |  10 allocs/op |   1.60 s |   60000 KB |   53.30 ns/alloc
BenchmarkGogoprotobufMarshal-8           |   10000000 |    132 ns/op |    64 B/op |   1 allocs/op |   1.32 s |   64000 KB |  132.00 ns/alloc
BenchmarkGogoprotobufUnmarshal-8         |   10000000 |    187 ns/op |    96 B/op |   3 allocs/op |   1.87 s |   96000 KB |   62.33 ns/alloc
BenchmarkColferMarshal-8                 |   20000000 |    108 ns/op |    64 B/op |   1 allocs/op |   2.16 s |  128000 KB |  108.00 ns/alloc
BenchmarkColferUnmarshal-8               |   10000000 |    156 ns/op |   112 B/op |   3 allocs/op |   1.56 s |  112000 KB |   52.00 ns/alloc
BenchmarkGencodeMarshal-8                |   10000000 |    134 ns/op |    80 B/op |   2 allocs/op |   1.34 s |   80000 KB |   67.00 ns/alloc
BenchmarkGencodeUnmarshal-8              |   10000000 |    160 ns/op |   112 B/op |   3 allocs/op |   1.60 s |  112000 KB |   53.33 ns/alloc
BenchmarkGencodeUnsafeMarshal-8          |   20000000 |     85 ns/op |    48 B/op |   1 allocs/op |   1.70 s |   96000 KB |   85.00 ns/alloc
BenchmarkGencodeUnsafeUnmarshal-8        |   10000000 |    123 ns/op |    96 B/op |   3 allocs/op |   1.23 s |   96000 KB |   41.00 ns/alloc
BenchmarkXDR2Marshal-8                   |   10000000 |    140 ns/op |    64 B/op |   1 allocs/op |   1.40 s |   64000 KB |  140.00 ns/alloc
BenchmarkXDR2Unmarshal-8                 |   20000000 |    111 ns/op |    32 B/op |   2 allocs/op |   2.22 s |   64000 KB |   55.50 ns/alloc
BenchmarkGoAvroMarshal-8                 |    1000000 |   2331 ns/op |  1030 B/op |  32 allocs/op |   2.33 s |  103000 KB |   72.84 ns/alloc
BenchmarkGoAvroUnmarshal-8               |     300000 |   5496 ns/op |  3372 B/op |  87 allocs/op |   1.65 s |  101160 KB |   63.17 ns/alloc
BenchmarkGoAvro2TextMarshal-8            |    1000000 |   2225 ns/op |  1326 B/op |  20 allocs/op |   2.23 s |  132600 KB |  111.25 ns/alloc
BenchmarkGoAvro2TextUnmarshal-8          |    1000000 |   2338 ns/op |   806 B/op |  33 allocs/op |   2.34 s |   80600 KB |   70.85 ns/alloc
BenchmarkGoAvro2BinaryMarshal-8          |    2000000 |    791 ns/op |   510 B/op |  11 allocs/op |   1.58 s |  102000 KB |   71.91 ns/alloc
BenchmarkGoAvro2BinaryUnmarshal-8        |    2000000 |    829 ns/op |   576 B/op |  13 allocs/op |   1.66 s |  115200 KB |   63.77 ns/alloc
BenchmarkIkeaMarshal-8                   |    3000000 |    537 ns/op |    72 B/op |   8 allocs/op |   1.61 s |   21600 KB |   67.12 ns/alloc
BenchmarkIkeaUnmarshal-8                 |    2000000 |    681 ns/op |   160 B/op |  11 allocs/op |   1.36 s |   32000 KB |   61.91 ns/alloc
BenchmarkShamatonMapMsgpackMarshal-8     |    2000000 |    680 ns/op |   208 B/op |   4 allocs/op |   1.36 s |   41600 KB |  170.00 ns/alloc
BenchmarkShamatonMapMsgpackUnmarshal-8   |    2000000 |    611 ns/op |   144 B/op |   3 allocs/op |   1.22 s |   28800 KB |  203.67 ns/alloc
BenchmarkShamatonArrayMsgpackMarshal-8   |    2000000 |    604 ns/op |   176 B/op |   4 allocs/op |   1.21 s |   35200 KB |  151.00 ns/alloc
BenchmarkShamatonArrayMsgpackUnmarshal-8 |    3000000 |    409 ns/op |   144 B/op |   3 allocs/op |   1.23 s |   43200 KB |  136.33 ns/alloc
BenchmarkSSZNoTimeNoStringNoFloatAMarshal-8 |     300000 |   4003 ns/op |   440 B/op |  71 allocs/op |   1.20 s |   13200 KB |   56.38 ns/alloc
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal-8 |     200000 |   6519 ns/op |  1392 B/op |  78 allocs/op |   1.30 s |   27840 KB |   83.58 ns/alloc


---
Totals:

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.time | tt.bytes     | time/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGencodeUnsafe-8                 |   30000000 |    208 ns/op |   144 B/op |   4 allocs/op |   6.24 s |  432000 KB |   52.00 ns/alloc
BenchmarkXDR2-8                          |   30000000 |    251 ns/op |    96 B/op |   3 allocs/op |   7.53 s |  288000 KB |   83.67 ns/alloc
BenchmarkColfer-8                        |   30000000 |    264 ns/op |   176 B/op |   4 allocs/op |   7.92 s |  528000 KB |   66.00 ns/alloc
BenchmarkGotinyNoTime-8                  |   30000000 |    274 ns/op |    96 B/op |   3 allocs/op |   8.22 s |  288000 KB |   91.37 ns/alloc
BenchmarkGotiny-8                        |   30000000 |    291 ns/op |   112 B/op |   3 allocs/op |   8.74 s |  336000 KB |   97.10 ns/alloc
BenchmarkGencode-8                       |   20000000 |    294 ns/op |   192 B/op |   5 allocs/op |   5.88 s |  384000 KB |   58.80 ns/alloc
BenchmarkGogoprotobuf-8                  |   20000000 |    319 ns/op |   160 B/op |   4 allocs/op |   6.38 s |  320000 KB |   79.75 ns/alloc
BenchmarkMsgp-8                          |   15000000 |    420 ns/op |   240 B/op |   4 allocs/op |   6.30 s |  360000 KB |  105.00 ns/alloc
BenchmarkFlatBuffers-8                   |   15000000 |    478 ns/op |   112 B/op |   3 allocs/op |   7.17 s |  168000 KB |  159.33 ns/alloc
BenchmarkCapNProto-8                     |   10000000 |    724 ns/op |   256 B/op |   8 allocs/op |   7.24 s |  256000 KB |   90.50 ns/alloc
BenchmarkGoprotobuf-8                    |    8000000 |    870 ns/op |   296 B/op |  12 allocs/op |   6.96 s |  236800 KB |   72.50 ns/alloc
BenchmarkHprose2-8                       |    6000000 |    966 ns/op |   144 B/op |   4 allocs/op |   5.80 s |   86400 KB |  241.50 ns/alloc
BenchmarkShamatonArrayMsgpack-8          |    5000000 |   1013 ns/op |   320 B/op |   7 allocs/op |   5.07 s |  160000 KB |  144.71 ns/alloc
BenchmarkIkea-8                          |    5000000 |   1218 ns/op |   232 B/op |  19 allocs/op |   6.09 s |  116000 KB |   64.11 ns/alloc
BenchmarkShamatonMapMsgpack-8            |    4000000 |   1291 ns/op |   352 B/op |   7 allocs/op |   5.16 s |  140800 KB |  184.43 ns/alloc
BenchmarkCapNProto2-8                    |    5000000 |   1363 ns/op |   564 B/op |   9 allocs/op |   6.82 s |  282000 KB |  151.44 ns/alloc
BenchmarkProtobuf-8                      |    4000000 |   1424 ns/op |   344 B/op |  17 allocs/op |   5.70 s |  137600 KB |   83.76 ns/alloc
BenchmarkGob-8                           |    4000000 |   1537 ns/op |   160 B/op |   5 allocs/op |   6.15 s |   64000 KB |  307.40 ns/alloc
BenchmarkGoAvro2Binary-8                 |    4000000 |   1620 ns/op |  1086 B/op |  24 allocs/op |   6.48 s |  434400 KB |   67.50 ns/alloc
BenchmarkHprose-8                        |    4000000 |   1731 ns/op |   650 B/op |  18 allocs/op |   6.92 s |  260000 KB |   96.17 ns/alloc
BenchmarkUgorjiCodecMsgpack-8            |    3000000 |   1958 ns/op |  1808 B/op |   7 allocs/op |   5.87 s |  542400 KB |  279.71 ns/alloc
BenchmarkUgorjiCodecBinc-8               |    2000000 |   2198 ns/op |  1968 B/op |  11 allocs/op |   4.40 s |  393600 KB |  199.82 ns/alloc
BenchmarkVmihailencoMsgpack-8            |    3000000 |   2257 ns/op |   816 B/op |  20 allocs/op |   6.77 s |  244800 KB |  112.85 ns/alloc
BenchmarkEasyJson-8                      |    2000000 |   2333 ns/op |   943 B/op |   9 allocs/op |   4.67 s |  188600 KB |  259.22 ns/alloc
BenchmarkBinary-8                        |    2000000 |   2425 ns/op |   646 B/op |  43 allocs/op |   4.85 s |  129200 KB |   56.40 ns/alloc
BenchmarkXdr-8                           |    2000000 |   2532 ns/op |   616 B/op |  30 allocs/op |   5.06 s |  123200 KB |   84.40 ns/alloc
BenchmarkBson-8                          |    2000000 |   2564 ns/op |   636 B/op |  29 allocs/op |   5.13 s |  127200 KB |   88.41 ns/alloc
BenchmarkJsonIter-8                      |    2000000 |   3129 ns/op |   575 B/op |  11 allocs/op |   6.26 s |  115000 KB |  284.45 ns/alloc
BenchmarkGoAvro2Text-8                   |    2000000 |   4563 ns/op |  2132 B/op |  53 allocs/op |   9.13 s |  426400 KB |   86.09 ns/alloc
BenchmarkSereal-8                        |    1500000 |   4826 ns/op |  1912 B/op |  54 allocs/op |   7.24 s |  286800 KB |   89.37 ns/alloc
BenchmarkJson-8                          |    1500000 |   4892 ns/op |   663 B/op |  11 allocs/op |   7.34 s |   99450 KB |  444.73 ns/alloc
BenchmarkGoAvro-8                        |    1300000 |   7827 ns/op |  4402 B/op | 119 allocs/op |  10.18 s |  572260 KB |   65.77 ns/alloc
BenchmarkSSZNoTimeNoStringNoFloatA-8     |     500000 |  10522 ns/op |  1832 B/op | 149 allocs/op |   5.26 s |   91600 KB |   70.62 ns/alloc



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
