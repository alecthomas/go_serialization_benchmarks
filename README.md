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

2018-07-20 Results with Go 1.10.3 on a 2.5 GHz Intel Core i7 MacBook Pro (Retina, 15-inch, Mid 2015):

```
benchmark                                      iter    time/iter   bytes/op     allocs/op  tt.time   tt.bytes       time/alloc
---------                                      ----    ---------   --------     ---------  -------   --------       ----------
BenchmarkGotinyMarshal-8                    5000000    271 ns/op    16 B/op   1 allocs/op   1.35 s    8000 KB  271.00 ns/alloc
BenchmarkGotinyUnmarshal-8                  5000000    313 ns/op   112 B/op   3 allocs/op   1.56 s   56000 KB  104.33 ns/alloc
BenchmarkGotinyNoTimeMarshal-8             10000000    157 ns/op     0 B/op   0 allocs/op   1.57 s       0 KB    0.00 ns/alloc
BenchmarkGotinyNoTimeUnmarshal-8           10000000    224 ns/op    96 B/op   3 allocs/op   2.24 s   96000 KB   74.67 ns/alloc
BenchmarkMsgpMarshal-8                     10000000    170 ns/op   128 B/op   1 allocs/op   1.70 s  128000 KB  170.00 ns/alloc
BenchmarkMsgpUnmarshal-8                    5000000    322 ns/op   112 B/op   3 allocs/op   1.61 s   56000 KB  107.33 ns/alloc
BenchmarkVmihailencoMsgpackMarshal-8        1000000   1742 ns/op   368 B/op   6 allocs/op   1.74 s   36800 KB  290.33 ns/alloc
BenchmarkVmihailencoMsgpackUnmarshal-8      1000000   1901 ns/op   383 B/op  13 allocs/op   1.90 s   38300 KB  146.23 ns/alloc
BenchmarkJsonMarshal-8                      1000000   2145 ns/op   528 B/op   5 allocs/op   2.15 s   52800 KB  429.00 ns/alloc
BenchmarkJsonUnmarshal-8                     300000   4125 ns/op   495 B/op   8 allocs/op   1.24 s   14850 KB  515.62 ns/alloc
BenchmarkJsonIterMarshal-8                  1000000   1252 ns/op   200 B/op   3 allocs/op   1.25 s   20000 KB  417.33 ns/alloc
BenchmarkJsonIterUnmarshal-8                1000000   1511 ns/op   248 B/op   7 allocs/op   1.51 s   24800 KB  215.86 ns/alloc
BenchmarkEasyJsonMarshal-8                  1000000   1280 ns/op   784 B/op   5 allocs/op   1.28 s   78400 KB  256.00 ns/alloc
BenchmarkEasyJsonUnmarshal-8                1000000   1306 ns/op   160 B/op   4 allocs/op   1.31 s   16000 KB  326.50 ns/alloc
BenchmarkBsonMarshal-8                      1000000   1347 ns/op   392 B/op  10 allocs/op   1.35 s   39200 KB  134.70 ns/alloc
BenchmarkBsonUnmarshal-8                    1000000   1814 ns/op   244 B/op  19 allocs/op   1.81 s   24400 KB   95.47 ns/alloc
BenchmarkGobMarshal-8                       2000000    891 ns/op    48 B/op   2 allocs/op   1.78 s    9600 KB  445.50 ns/alloc
BenchmarkGobUnmarshal-8                     2000000    899 ns/op   112 B/op   3 allocs/op   1.80 s   22400 KB  299.67 ns/alloc
BenchmarkXdrMarshal-8                       1000000   1584 ns/op   456 B/op  21 allocs/op   1.58 s   45600 KB   75.43 ns/alloc
BenchmarkXdrUnmarshal-8                     1000000   1373 ns/op   239 B/op  11 allocs/op   1.37 s   23900 KB  124.82 ns/alloc
BenchmarkUgorjiCodecMsgpackMarshal-8        1000000   1031 ns/op   545 B/op   6 allocs/op   1.03 s   54500 KB  171.83 ns/alloc
BenchmarkUgorjiCodecMsgpackUnmarshal-8      1000000   1265 ns/op   465 B/op   6 allocs/op   1.26 s   46500 KB  210.83 ns/alloc
BenchmarkUgorjiCodecBincMarshal-8           1000000   1124 ns/op   577 B/op   7 allocs/op   1.12 s   57700 KB  160.57 ns/alloc
BenchmarkUgorjiCodecBincUnmarshal-8         1000000   1425 ns/op   689 B/op   9 allocs/op   1.43 s   68900 KB  158.33 ns/alloc
BenchmarkSerealMarshal-8                     500000   2532 ns/op   911 B/op  21 allocs/op   1.27 s   45550 KB  120.57 ns/alloc
BenchmarkSerealUnmarshal-8                   500000   2751 ns/op  1008 B/op  34 allocs/op   1.38 s   50400 KB   80.91 ns/alloc
BenchmarkBinaryMarshal-8                    1000000   1377 ns/op   334 B/op  20 allocs/op   1.38 s   33400 KB   68.85 ns/alloc
BenchmarkBinaryUnmarshal-8                  1000000   1484 ns/op   336 B/op  22 allocs/op   1.48 s   33600 KB   67.45 ns/alloc
BenchmarkFlatBuffersMarshal-8               5000000    317 ns/op     0 B/op   0 allocs/op   1.58 s       0 KB    0.00 ns/alloc
BenchmarkFlatBuffersUnmarshal-8             5000000    242 ns/op   112 B/op   3 allocs/op   1.21 s   56000 KB   80.67 ns/alloc
BenchmarkCapNProtoMarshal-8                 3000000    464 ns/op    56 B/op   2 allocs/op   1.39 s   16800 KB  232.00 ns/alloc
BenchmarkCapNProtoUnmarshal-8               3000000    406 ns/op   200 B/op   6 allocs/op   1.22 s   60000 KB   67.67 ns/alloc
BenchmarkCapNProto2Marshal-8                2000000    682 ns/op   244 B/op   3 allocs/op   1.36 s   48800 KB  227.33 ns/alloc
BenchmarkCapNProto2Unmarshal-8              2000000    936 ns/op   320 B/op   6 allocs/op   1.87 s   64000 KB  156.00 ns/alloc
BenchmarkHproseMarshal-8                    2000000    891 ns/op   473 B/op   8 allocs/op   1.78 s   94600 KB  111.38 ns/alloc
BenchmarkHproseUnmarshal-8                  1000000   1105 ns/op   319 B/op  10 allocs/op   1.10 s   31900 KB  110.50 ns/alloc
BenchmarkHprose2Marshal-8                   2000000    586 ns/op     0 B/op   0 allocs/op   1.17 s       0 KB    0.00 ns/alloc
BenchmarkHprose2Unmarshal-8                 3000000    546 ns/op   144 B/op   4 allocs/op   1.64 s   43200 KB  136.50 ns/alloc
BenchmarkProtobufMarshal-8                  2000000    937 ns/op   200 B/op   7 allocs/op   1.87 s   40000 KB  133.86 ns/alloc
BenchmarkProtobufUnmarshal-8                2000000    833 ns/op   192 B/op  10 allocs/op   1.67 s   38400 KB   83.30 ns/alloc
BenchmarkGoprotobufMarshal-8                3000000    411 ns/op    96 B/op   2 allocs/op   1.23 s   28800 KB  205.50 ns/alloc
BenchmarkGoprotobufUnmarshal-8              2000000    616 ns/op   200 B/op  10 allocs/op   1.23 s   40000 KB   61.60 ns/alloc
BenchmarkGogoprotobufMarshal-8             10000000    153 ns/op    64 B/op   1 allocs/op   1.53 s   64000 KB  153.00 ns/alloc
BenchmarkGogoprotobufUnmarshal-8           10000000    219 ns/op    96 B/op   3 allocs/op   2.19 s   96000 KB   73.00 ns/alloc
BenchmarkColferMarshal-8                   10000000    128 ns/op    64 B/op   1 allocs/op   1.28 s   64000 KB  128.00 ns/alloc
BenchmarkColferUnmarshal-8                 10000000    173 ns/op   112 B/op   3 allocs/op   1.73 s  112000 KB   57.67 ns/alloc
BenchmarkGencodeMarshal-8                  10000000    162 ns/op    80 B/op   2 allocs/op   1.62 s   80000 KB   81.00 ns/alloc
BenchmarkGencodeUnmarshal-8                10000000    179 ns/op   112 B/op   3 allocs/op   1.79 s  112000 KB   59.67 ns/alloc
BenchmarkGencodeUnsafeMarshal-8            20000000    104 ns/op    48 B/op   1 allocs/op   2.08 s   96000 KB  104.00 ns/alloc
BenchmarkGencodeUnsafeUnmarshal-8          10000000    137 ns/op    96 B/op   3 allocs/op   1.37 s   96000 KB   45.67 ns/alloc
BenchmarkXDR2Marshal-8                     10000000    158 ns/op    64 B/op   1 allocs/op   1.58 s   64000 KB  158.00 ns/alloc
BenchmarkXDR2Unmarshal-8                   10000000    139 ns/op    32 B/op   2 allocs/op   1.39 s   32000 KB   69.50 ns/alloc
BenchmarkGoAvroMarshal-8                     500000   2402 ns/op  1030 B/op  31 allocs/op   1.20 s   51500 KB   77.48 ns/alloc
BenchmarkGoAvroUnmarshal-8                   200000   5768 ns/op  3437 B/op  87 allocs/op   1.15 s   68740 KB   66.30 ns/alloc
BenchmarkGoAvro2TextMarshal-8                500000   2733 ns/op  1326 B/op  20 allocs/op   1.37 s   66300 KB  136.65 ns/alloc
BenchmarkGoAvro2TextUnmarshal-8              500000   2602 ns/op   807 B/op  34 allocs/op   1.30 s   40350 KB   76.53 ns/alloc
BenchmarkGoAvro2BinaryMarshal-8             2000000    844 ns/op   510 B/op  11 allocs/op   1.69 s  102000 KB   76.73 ns/alloc
BenchmarkGoAvro2BinaryUnmarshal-8           2000000    902 ns/op   576 B/op  13 allocs/op   1.80 s  115200 KB   69.38 ns/alloc
BenchmarkIkeaMarshal-8                      2000000    618 ns/op    72 B/op   8 allocs/op   1.24 s   14400 KB   77.25 ns/alloc
BenchmarkIkeaUnmarshal-8                    2000000    795 ns/op   160 B/op  11 allocs/op   1.59 s   32000 KB   72.27 ns/alloc
---
totals:
BenchmarkGencodeUnsafe-8                   30000000    241 ns/op   144 B/op   4 allocs/op   7.23 s  432000 KB   60.25 ns/alloc
BenchmarkXDR2-8                            20000000    297 ns/op    96 B/op   3 allocs/op   5.94 s  192000 KB   99.00 ns/alloc
BenchmarkColfer-8                          20000000    301 ns/op   176 B/op   4 allocs/op   6.02 s  352000 KB   75.25 ns/alloc
BenchmarkGencode-8                         20000000    341 ns/op   192 B/op   5 allocs/op   6.82 s  384000 KB   68.20 ns/alloc
BenchmarkGogoprotobuf-8                    20000000    372 ns/op   160 B/op   4 allocs/op   7.44 s  320000 KB   93.00 ns/alloc
BenchmarkGotinyNoTime-8                    20000000    381 ns/op    96 B/op   3 allocs/op   7.62 s  192000 KB  127.00 ns/alloc
BenchmarkMsgp-8                            15000000    492 ns/op   240 B/op   4 allocs/op   7.38 s  360000 KB  123.00 ns/alloc
BenchmarkFlatBuffers-8                     10000000    559 ns/op   112 B/op   3 allocs/op   5.59 s  112000 KB  186.33 ns/alloc
BenchmarkGotiny-8                          10000000    584 ns/op   128 B/op   4 allocs/op   5.84 s  128000 KB  146.00 ns/alloc
BenchmarkCapNProto-8                        6000000    870 ns/op   256 B/op   8 allocs/op   5.22 s  153600 KB  108.75 ns/alloc
BenchmarkGoprotobuf-8                       5000000   1027 ns/op   296 B/op  12 allocs/op   5.13 s  148000 KB   85.58 ns/alloc
BenchmarkHprose2-8                          5000000   1132 ns/op   144 B/op   4 allocs/op   5.66 s   72000 KB  283.00 ns/alloc
BenchmarkIkea-8                             4000000   1413 ns/op   232 B/op  19 allocs/op   5.65 s   92800 KB   74.37 ns/alloc
BenchmarkCapNProto2-8                       4000000   1618 ns/op   564 B/op   9 allocs/op   6.47 s  225600 KB  179.78 ns/alloc
BenchmarkGoAvro2Binary-8                    4000000   1746 ns/op  1086 B/op  24 allocs/op   6.98 s  434400 KB   72.75 ns/alloc
BenchmarkProtobuf-8                         4000000   1770 ns/op   392 B/op  17 allocs/op   7.08 s  156800 KB  104.12 ns/alloc
BenchmarkGob-8                              4000000   1790 ns/op   160 B/op   5 allocs/op   7.16 s   64000 KB  358.00 ns/alloc
BenchmarkHprose-8                           3000000   1996 ns/op   792 B/op  18 allocs/op   5.99 s  237600 KB  110.89 ns/alloc
BenchmarkUgorjiCodecMsgpack-8               2000000   2296 ns/op  1010 B/op  12 allocs/op   4.59 s  202000 KB  191.33 ns/alloc
BenchmarkUgorjiCodecBinc-8                  2000000   2549 ns/op  1266 B/op  16 allocs/op   5.10 s  253200 KB  159.31 ns/alloc
BenchmarkEasyJson-8                         2000000   2586 ns/op   944 B/op   9 allocs/op   5.17 s  188800 KB  287.33 ns/alloc
BenchmarkJsonIter-8                         2000000   2763 ns/op   448 B/op  10 allocs/op   5.53 s   89600 KB  276.30 ns/alloc
BenchmarkBinary-8                           2000000   2861 ns/op   670 B/op  42 allocs/op   5.72 s  134000 KB   68.12 ns/alloc
BenchmarkXdr-8                              2000000   2957 ns/op   695 B/op  32 allocs/op   5.91 s  139000 KB   92.41 ns/alloc
BenchmarkBson-8                             2000000   3161 ns/op   636 B/op  29 allocs/op   6.32 s  127200 KB  109.00 ns/alloc
BenchmarkVmihailencoMsgpack-8               2000000   3643 ns/op   751 B/op  19 allocs/op   7.29 s  150200 KB  191.74 ns/alloc
BenchmarkSereal-8                           1000000   5283 ns/op  1919 B/op  55 allocs/op   5.28 s  191900 KB   96.05 ns/alloc
BenchmarkGoAvro2Text-8                      1000000   5335 ns/op  2133 B/op  54 allocs/op   5.33 s  213300 KB   98.80 ns/alloc
BenchmarkJson-8                             1300000   6270 ns/op  1023 B/op  13 allocs/op   8.15 s  132990 KB  482.31 ns/alloc
BenchmarkGoAvro-8                            700000   8170 ns/op  4467 B/op 118 allocs/op   5.72 s  312690 KB   69.24 ns/alloc
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
