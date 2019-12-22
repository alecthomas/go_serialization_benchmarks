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
- [github.com/vmihailenco/msgpack/v4](https://github.com/vmihailenco/msgpack)
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp) *(code generator for msgpack)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://github.com/gogo/protobuf) (generated code, optimized version of `goprotobuf`)
- [go.dedis.ch/protobuf](https://go.dedis.ch/protobuf) (reflection based)
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

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGotinyMarshal-8                 |   20000000 |     99 ns/op |     0 |   0 |   1.99 |       0 |    0.00
BenchmarkGotinyUnmarshal-8               |   10000000 |    192 ns/op |   112 |   3 |   1.92 |  112000 |   64.00
BenchmarkGotinyNoTimeMarshal-8           |   20000000 |     99 ns/op |     0 |   0 |   1.98 |       0 |    0.00
BenchmarkGotinyNoTimeUnmarshal-8         |   10000000 |    175 ns/op |    96 |   3 |   1.75 |   96000 |   58.33
BenchmarkMsgpMarshal-8                   |   10000000 |    140 ns/op |   128 |   1 |   1.40 |  128000 |  140.00
BenchmarkMsgpUnmarshal-8                 |    5000000 |    280 ns/op |   112 |   3 |   1.40 |   56000 |   93.33
BenchmarkVmihailencoMsgpackMarshal-8     |    2000000 |    986 ns/op |   400 |   7 |   1.97 |   80000 |  140.86
BenchmarkVmihailencoMsgpackUnmarshal-8   |    1000000 |   1271 ns/op |   416 |  13 |   1.27 |   41600 |   97.77
BenchmarkJsonMarshal-8                   |    1000000 |   1585 ns/op |   304 |   4 |   1.58 |   30400 |  396.25
BenchmarkJsonUnmarshal-8                 |     500000 |   3307 ns/op |   359 |   7 |   1.65 |   17950 |  472.43
BenchmarkJsonIterMarshal-8               |    1000000 |   1768 ns/op |   312 |   5 |   1.77 |   31200 |  353.60
BenchmarkJsonIterUnmarshal-8             |    1000000 |   1361 ns/op |   263 |   6 |   1.36 |   26300 |  226.83
BenchmarkEasyJsonMarshal-8               |    1000000 |   1125 ns/op |   784 |   5 |   1.12 |   78400 |  225.00
BenchmarkEasyJsonUnmarshal-8             |    1000000 |   1208 ns/op |   159 |   4 |   1.21 |   15900 |  302.00
BenchmarkBsonMarshal-8                   |    1000000 |   1005 ns/op |   392 |  10 |   1.00 |   39200 |  100.50
BenchmarkBsonUnmarshal-8                 |    1000000 |   1559 ns/op |   244 |  19 |   1.56 |   24400 |   82.05
BenchmarkGobMarshal-8                    |    2000000 |    744 ns/op |    48 |   2 |   1.49 |    9600 |  372.00
BenchmarkGobUnmarshal-8                  |    2000000 |    793 ns/op |   112 |   3 |   1.59 |   22400 |  264.33
BenchmarkXdrMarshal-8                    |    1000000 |   1332 ns/op |   392 |  19 |   1.33 |   39200 |   70.11
BenchmarkXdrUnmarshal-8                  |    1000000 |   1200 ns/op |   224 |  11 |   1.20 |   22400 |  109.09
BenchmarkUgorjiCodecMsgpackMarshal-8     |    2000000 |    922 ns/op |  1312 |   3 |   1.84 |  262400 |  307.33
BenchmarkUgorjiCodecMsgpackUnmarshal-8   |    1000000 |   1036 ns/op |   496 |   4 |   1.04 |   49600 |  259.00
BenchmarkUgorjiCodecBincMarshal-8        |    1000000 |   1019 ns/op |  1328 |   4 |   1.02 |  132800 |  254.75
BenchmarkUgorjiCodecBincUnmarshal-8      |    1000000 |   1179 ns/op |   640 |   7 |   1.18 |   64000 |  168.43
BenchmarkSerealMarshal-8                 |    1000000 |   2192 ns/op |   904 |  20 |   2.19 |   90400 |  109.60
BenchmarkSerealUnmarshal-8               |     500000 |   2634 ns/op |  1008 |  34 |   1.32 |   50400 |   77.47
BenchmarkBinaryMarshal-8                 |    1000000 |   1207 ns/op |   326 |  21 |   1.21 |   32600 |   57.48
BenchmarkBinaryUnmarshal-8               |    1000000 |   1218 ns/op |   320 |  22 |   1.22 |   32000 |   55.36
BenchmarkFlatBuffersMarshal-8            |    5000000 |    266 ns/op |     0 |   0 |   1.33 |       0 |    0.00
BenchmarkFlatBuffersUnmarshal-8          |   10000000 |    212 ns/op |   112 |   3 |   2.12 |  112000 |   70.67
BenchmarkCapNProtoMarshal-8              |    5000000 |    369 ns/op |    56 |   2 |   1.84 |   28000 |  184.50
BenchmarkCapNProtoUnmarshal-8            |    5000000 |    355 ns/op |   200 |   6 |   1.77 |  100000 |   59.17
BenchmarkCapNProto2Marshal-8             |    3000000 |    571 ns/op |   244 |   3 |   1.71 |   73200 |  190.33
BenchmarkCapNProto2Unmarshal-8           |    2000000 |    792 ns/op |   320 |   6 |   1.58 |   64000 |  132.00
BenchmarkHproseMarshal-8                 |    2000000 |    772 ns/op |   331 |   8 |   1.54 |   66200 |   96.50
BenchmarkHproseUnmarshal-8               |    2000000 |    959 ns/op |   319 |  10 |   1.92 |   63800 |   95.90
BenchmarkHprose2Marshal-8                |    3000000 |    475 ns/op |     0 |   0 |   1.43 |       0 |    0.00
BenchmarkHprose2Unmarshal-8              |    3000000 |    491 ns/op |   144 |   4 |   1.47 |   43200 |  122.75
BenchmarkProtobufMarshal-8               |    2000000 |    738 ns/op |   152 |   7 |   1.48 |   30400 |  105.43
BenchmarkProtobufUnmarshal-8             |    2000000 |    686 ns/op |   192 |  10 |   1.37 |   38400 |   68.60
BenchmarkGoprotobufMarshal-8             |    5000000 |    337 ns/op |    96 |   2 |   1.69 |   48000 |  168.50
BenchmarkGoprotobufUnmarshal-8           |    3000000 |    533 ns/op |   200 |  10 |   1.60 |   60000 |   53.30
BenchmarkGogoprotobufMarshal-8           |   10000000 |    132 ns/op |    64 |   1 |   1.32 |   64000 |  132.00
BenchmarkGogoprotobufUnmarshal-8         |   10000000 |    187 ns/op |    96 |   3 |   1.87 |   96000 |   62.33
BenchmarkColferMarshal-8                 |   20000000 |    108 ns/op |    64 |   1 |   2.16 |  128000 |  108.00
BenchmarkColferUnmarshal-8               |   10000000 |    156 ns/op |   112 |   3 |   1.56 |  112000 |   52.00
BenchmarkGencodeMarshal-8                |   10000000 |    134 ns/op |    80 |   2 |   1.34 |   80000 |   67.00
BenchmarkGencodeUnmarshal-8              |   10000000 |    160 ns/op |   112 |   3 |   1.60 |  112000 |   53.33
BenchmarkGencodeUnsafeMarshal-8          |   20000000 |     85 ns/op |    48 |   1 |   1.70 |   96000 |   85.00
BenchmarkGencodeUnsafeUnmarshal-8        |   10000000 |    123 ns/op |    96 |   3 |   1.23 |   96000 |   41.00
BenchmarkXDR2Marshal-8                   |   10000000 |    140 ns/op |    64 |   1 |   1.40 |   64000 |  140.00
BenchmarkXDR2Unmarshal-8                 |   20000000 |    111 ns/op |    32 |   2 |   2.22 |   64000 |   55.50
BenchmarkGoAvroMarshal-8                 |    1000000 |   2331 ns/op |  1030 |  32 |   2.33 |  103000 |   72.84
BenchmarkGoAvroUnmarshal-8               |     300000 |   5496 ns/op |  3372 |  87 |   1.65 |  101160 |   63.17
BenchmarkGoAvro2TextMarshal-8            |    1000000 |   2225 ns/op |  1326 |  20 |   2.23 |  132600 |  111.25
BenchmarkGoAvro2TextUnmarshal-8          |    1000000 |   2338 ns/op |   806 |  33 |   2.34 |   80600 |   70.85
BenchmarkGoAvro2BinaryMarshal-8          |    2000000 |    791 ns/op |   510 |  11 |   1.58 |  102000 |   71.91
BenchmarkGoAvro2BinaryUnmarshal-8        |    2000000 |    829 ns/op |   576 |  13 |   1.66 |  115200 |   63.77
BenchmarkIkeaMarshal-8                   |    3000000 |    537 ns/op |    72 |   8 |   1.61 |   21600 |   67.12
BenchmarkIkeaUnmarshal-8                 |    2000000 |    681 ns/op |   160 |  11 |   1.36 |   32000 |   61.91
BenchmarkShamatonMapMsgpackMarshal-8     |    2000000 |    680 ns/op |   208 |   4 |   1.36 |   41600 |  170.00
BenchmarkShamatonMapMsgpackUnmarshal-8   |    2000000 |    611 ns/op |   144 |   3 |   1.22 |   28800 |  203.67
BenchmarkShamatonArrayMsgpackMarshal-8   |    2000000 |    604 ns/op |   176 |   4 |   1.21 |   35200 |  151.00
BenchmarkShamatonArrayMsgpackUnmarshal-8 |    3000000 |    409 ns/op |   144 |   3 |   1.23 |   43200 |  136.33
BenchmarkSSZNoTimeNoStringNoFloatAMarshal-8 |     300000 |   4003 ns/op |   440 |  71 |   1.20 |   13200 |   56.38
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal-8 |     200000 |   6519 ns/op |  1392 |  78 |   1.30 |   27840 |   83.58


Totals:


benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGencodeUnsafe-8                 |   30000000 |    208 ns/op |   144 |   4 |   6.24 |  432000 |   52.00
BenchmarkXDR2-8                          |   30000000 |    251 ns/op |    96 |   3 |   7.53 |  288000 |   83.67
BenchmarkColfer-8                        |   30000000 |    264 ns/op |   176 |   4 |   7.92 |  528000 |   66.00
BenchmarkGotinyNoTime-8                  |   30000000 |    274 ns/op |    96 |   3 |   8.22 |  288000 |   91.37
BenchmarkGotiny-8                        |   30000000 |    291 ns/op |   112 |   3 |   8.74 |  336000 |   97.10
BenchmarkGencode-8                       |   20000000 |    294 ns/op |   192 |   5 |   5.88 |  384000 |   58.80
BenchmarkGogoprotobuf-8                  |   20000000 |    319 ns/op |   160 |   4 |   6.38 |  320000 |   79.75
BenchmarkMsgp-8                          |   15000000 |    420 ns/op |   240 |   4 |   6.30 |  360000 |  105.00
BenchmarkFlatBuffers-8                   |   15000000 |    478 ns/op |   112 |   3 |   7.17 |  168000 |  159.33
BenchmarkCapNProto-8                     |   10000000 |    724 ns/op |   256 |   8 |   7.24 |  256000 |   90.50
BenchmarkGoprotobuf-8                    |    8000000 |    870 ns/op |   296 |  12 |   6.96 |  236800 |   72.50
BenchmarkHprose2-8                       |    6000000 |    966 ns/op |   144 |   4 |   5.80 |   86400 |  241.50
BenchmarkShamatonArrayMsgpack-8          |    5000000 |   1013 ns/op |   320 |   7 |   5.07 |  160000 |  144.71
BenchmarkIkea-8                          |    5000000 |   1218 ns/op |   232 |  19 |   6.09 |  116000 |   64.11
BenchmarkShamatonMapMsgpack-8            |    4000000 |   1291 ns/op |   352 |   7 |   5.16 |  140800 |  184.43
BenchmarkCapNProto2-8                    |    5000000 |   1363 ns/op |   564 |   9 |   6.82 |  282000 |  151.44
BenchmarkProtobuf-8                      |    4000000 |   1424 ns/op |   344 |  17 |   5.70 |  137600 |   83.76
BenchmarkGob-8                           |    4000000 |   1537 ns/op |   160 |   5 |   6.15 |   64000 |  307.40
BenchmarkGoAvro2Binary-8                 |    4000000 |   1620 ns/op |  1086 |  24 |   6.48 |  434400 |   67.50
BenchmarkHprose-8                        |    4000000 |   1731 ns/op |   650 |  18 |   6.92 |  260000 |   96.17
BenchmarkUgorjiCodecMsgpack-8            |    3000000 |   1958 ns/op |  1808 |   7 |   5.87 |  542400 |  279.71
BenchmarkUgorjiCodecBinc-8               |    2000000 |   2198 ns/op |  1968 |  11 |   4.40 |  393600 |  199.82
BenchmarkVmihailencoMsgpack-8            |    3000000 |   2257 ns/op |   816 |  20 |   6.77 |  244800 |  112.85
BenchmarkEasyJson-8                      |    2000000 |   2333 ns/op |   943 |   9 |   4.67 |  188600 |  259.22
BenchmarkBinary-8                        |    2000000 |   2425 ns/op |   646 |  43 |   4.85 |  129200 |   56.40
BenchmarkXdr-8                           |    2000000 |   2532 ns/op |   616 |  30 |   5.06 |  123200 |   84.40
BenchmarkBson-8                          |    2000000 |   2564 ns/op |   636 |  29 |   5.13 |  127200 |   88.41
BenchmarkJsonIter-8                      |    2000000 |   3129 ns/op |   575 |  11 |   6.26 |  115000 |  284.45
BenchmarkGoAvro2Text-8                   |    2000000 |   4563 ns/op |  2132 |  53 |   9.13 |  426400 |   86.09
BenchmarkSereal-8                        |    1500000 |   4826 ns/op |  1912 |  54 |   7.24 |  286800 |   89.37
BenchmarkJson-8                          |    1500000 |   4892 ns/op |   663 |  11 |   7.34 |   99450 |  444.73
BenchmarkGoAvro-8                        |    1300000 |   7827 ns/op |  4402 | 119 |  10.18 |  572260 |   65.77
BenchmarkSSZNoTimeNoStringNoFloatA-8     |     500000 |  10522 ns/op |  1832 | 149 |   5.26 |   91600 |   70.62




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
