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

2018-08-19 Results with Go 1.10 on a 2.5 GHz Intel Core i7 MacBook Pro (Retina, 15-inch, Mid 2015):

```
benchmark                                      iter    time/iter   bytes/op     allocs/op  tt.time   tt.bytes       time/alloc
---------                                      ----    ---------   --------     ---------  -------   --------       ----------
BenchmarkGotinyMarshal-8                   10000000    132 ns/op     0 B/op   0 allocs/op   1.32 s       0 KB    0.00 ns/alloc
BenchmarkGotinyUnmarshal-8                 10000000    216 ns/op   112 B/op   3 allocs/op   2.16 s  112000 KB   72.00 ns/alloc
BenchmarkGotinyNoTimeMarshal-8             10000000    124 ns/op     0 B/op   0 allocs/op   1.24 s       0 KB    0.00 ns/alloc
BenchmarkGotinyNoTimeUnmarshal-8           10000000    203 ns/op    96 B/op   3 allocs/op   2.03 s   96000 KB   67.67 ns/alloc
BenchmarkMsgpMarshal-8                     10000000    171 ns/op   128 B/op   1 allocs/op   1.71 s  128000 KB  171.00 ns/alloc
BenchmarkMsgpUnmarshal-8                    5000000    326 ns/op   112 B/op   3 allocs/op   1.63 s   56000 KB  108.67 ns/alloc
BenchmarkVmihailencoMsgpackMarshal-8        1000000   1706 ns/op   368 B/op   6 allocs/op   1.71 s   36800 KB  284.33 ns/alloc
BenchmarkVmihailencoMsgpackUnmarshal-8      1000000   1863 ns/op   384 B/op  13 allocs/op   1.86 s   38400 KB  143.31 ns/alloc
BenchmarkJsonMarshal-8                      1000000   2165 ns/op   528 B/op   5 allocs/op   2.17 s   52800 KB  433.00 ns/alloc
BenchmarkJsonUnmarshal-8                     300000   4134 ns/op   495 B/op   8 allocs/op   1.24 s   14850 KB  516.75 ns/alloc
BenchmarkJsonIterMarshal-8                  1000000   1259 ns/op   200 B/op   3 allocs/op   1.26 s   20000 KB  419.67 ns/alloc
BenchmarkJsonIterUnmarshal-8                1000000   1518 ns/op   248 B/op   7 allocs/op   1.52 s   24800 KB  216.86 ns/alloc
BenchmarkEasyJsonMarshal-8                  1000000   1304 ns/op   784 B/op   5 allocs/op   1.30 s   78400 KB  260.80 ns/alloc
BenchmarkEasyJsonUnmarshal-8                1000000   1411 ns/op   160 B/op   4 allocs/op   1.41 s   16000 KB  352.75 ns/alloc
BenchmarkBsonMarshal-8                      1000000   1373 ns/op   392 B/op  10 allocs/op   1.37 s   39200 KB  137.30 ns/alloc
BenchmarkBsonUnmarshal-8                    1000000   1838 ns/op   244 B/op  19 allocs/op   1.84 s   24400 KB   96.74 ns/alloc
BenchmarkGobMarshal-8                       2000000    942 ns/op    48 B/op   2 allocs/op   1.88 s    9600 KB  471.00 ns/alloc
BenchmarkGobUnmarshal-8                     2000000    945 ns/op   112 B/op   3 allocs/op   1.89 s   22400 KB  315.00 ns/alloc
BenchmarkXdrMarshal-8                       1000000   1660 ns/op   456 B/op  21 allocs/op   1.66 s   45600 KB   79.05 ns/alloc
BenchmarkXdrUnmarshal-8                     1000000   1451 ns/op   240 B/op  11 allocs/op   1.45 s   24000 KB  131.91 ns/alloc
BenchmarkUgorjiCodecMsgpackMarshal-8        1000000   1072 ns/op   545 B/op   6 allocs/op   1.07 s   54500 KB  178.67 ns/alloc
BenchmarkUgorjiCodecMsgpackUnmarshal-8      1000000   1276 ns/op   465 B/op   6 allocs/op   1.28 s   46500 KB  212.67 ns/alloc
BenchmarkUgorjiCodecBincMarshal-8           1000000   1154 ns/op   577 B/op   7 allocs/op   1.15 s   57700 KB  164.86 ns/alloc
BenchmarkUgorjiCodecBincUnmarshal-8         1000000   1449 ns/op   689 B/op   9 allocs/op   1.45 s   68900 KB  161.00 ns/alloc
BenchmarkSerealMarshal-8                     500000   2544 ns/op   912 B/op  21 allocs/op   1.27 s   45600 KB  121.14 ns/alloc
BenchmarkSerealUnmarshal-8                   500000   2776 ns/op  1008 B/op  34 allocs/op   1.39 s   50400 KB   81.65 ns/alloc
BenchmarkBinaryMarshal-8                    1000000   1403 ns/op   334 B/op  20 allocs/op   1.40 s   33400 KB   70.15 ns/alloc
BenchmarkBinaryUnmarshal-8                  1000000   1507 ns/op   335 B/op  22 allocs/op   1.51 s   33500 KB   68.50 ns/alloc
BenchmarkFlatBuffersMarshal-8               5000000    319 ns/op     0 B/op   0 allocs/op   1.59 s       0 KB    0.00 ns/alloc
BenchmarkFlatBuffersUnmarshal-8             5000000    248 ns/op   112 B/op   3 allocs/op   1.24 s   56000 KB   82.67 ns/alloc
BenchmarkCapNProtoMarshal-8                 3000000    475 ns/op    56 B/op   2 allocs/op   1.43 s   16800 KB  237.50 ns/alloc
BenchmarkCapNProtoUnmarshal-8               3000000    413 ns/op   200 B/op   6 allocs/op   1.24 s   60000 KB   68.83 ns/alloc
BenchmarkCapNProto2Marshal-8                2000000    696 ns/op   244 B/op   3 allocs/op   1.39 s   48800 KB  232.00 ns/alloc
BenchmarkCapNProto2Unmarshal-8              2000000   1016 ns/op   320 B/op   6 allocs/op   2.03 s   64000 KB  169.33 ns/alloc
BenchmarkHproseMarshal-8                    2000000    897 ns/op   473 B/op   8 allocs/op   1.79 s   94600 KB  112.12 ns/alloc
BenchmarkHproseUnmarshal-8                  1000000   1083 ns/op   320 B/op  10 allocs/op   1.08 s   32000 KB  108.30 ns/alloc
BenchmarkHprose2Marshal-8                   2000000    664 ns/op     0 B/op   0 allocs/op   1.33 s       0 KB    0.00 ns/alloc
BenchmarkHprose2Unmarshal-8                 3000000    536 ns/op   144 B/op   4 allocs/op   1.61 s   43200 KB  134.00 ns/alloc
BenchmarkProtobufMarshal-8                  2000000    954 ns/op   200 B/op   7 allocs/op   1.91 s   40000 KB  136.29 ns/alloc
BenchmarkProtobufUnmarshal-8                2000000    835 ns/op   192 B/op  10 allocs/op   1.67 s   38400 KB   83.50 ns/alloc
BenchmarkGoprotobufMarshal-8                3000000    411 ns/op    96 B/op   2 allocs/op   1.23 s   28800 KB  205.50 ns/alloc
BenchmarkGoprotobufUnmarshal-8              2000000    624 ns/op   200 B/op  10 allocs/op   1.25 s   40000 KB   62.40 ns/alloc
BenchmarkGogoprotobufMarshal-8             10000000    153 ns/op    64 B/op   1 allocs/op   1.53 s   64000 KB  153.00 ns/alloc
BenchmarkGogoprotobufUnmarshal-8           10000000    225 ns/op    96 B/op   3 allocs/op   2.25 s   96000 KB   75.00 ns/alloc
BenchmarkColferMarshal-8                   10000000    129 ns/op    64 B/op   1 allocs/op   1.29 s   64000 KB  129.00 ns/alloc
BenchmarkColferUnmarshal-8                 10000000    178 ns/op   112 B/op   3 allocs/op   1.78 s  112000 KB   59.33 ns/alloc
BenchmarkGencodeMarshal-8                  10000000    168 ns/op    80 B/op   2 allocs/op   1.68 s   80000 KB   84.00 ns/alloc
BenchmarkGencodeUnmarshal-8                10000000    181 ns/op   112 B/op   3 allocs/op   1.81 s  112000 KB   60.33 ns/alloc
BenchmarkGencodeUnsafeMarshal-8            20000000    105 ns/op    48 B/op   1 allocs/op   2.10 s   96000 KB  105.00 ns/alloc
BenchmarkGencodeUnsafeUnmarshal-8          10000000    140 ns/op    96 B/op   3 allocs/op   1.40 s   96000 KB   46.67 ns/alloc
BenchmarkXDR2Marshal-8                     10000000    159 ns/op    64 B/op   1 allocs/op   1.59 s   64000 KB  159.00 ns/alloc
BenchmarkXDR2Unmarshal-8                   10000000    140 ns/op    32 B/op   2 allocs/op   1.40 s   32000 KB   70.00 ns/alloc
BenchmarkGoAvroMarshal-8                     500000   2419 ns/op  1030 B/op  31 allocs/op   1.21 s   51500 KB   78.03 ns/alloc
BenchmarkGoAvroUnmarshal-8                   200000   5810 ns/op  3436 B/op  87 allocs/op   1.16 s   68720 KB   66.78 ns/alloc
BenchmarkGoAvro2TextMarshal-8                500000   2740 ns/op  1326 B/op  20 allocs/op   1.37 s   66300 KB  137.00 ns/alloc
BenchmarkGoAvro2TextUnmarshal-8              500000   2650 ns/op   807 B/op  34 allocs/op   1.32 s   40350 KB   77.94 ns/alloc
BenchmarkGoAvro2BinaryMarshal-8             2000000    859 ns/op   510 B/op  11 allocs/op   1.72 s  102000 KB   78.09 ns/alloc
BenchmarkGoAvro2BinaryUnmarshal-8           2000000    914 ns/op   576 B/op  13 allocs/op   1.83 s  115200 KB   70.31 ns/alloc
BenchmarkIkeaMarshal-8                      2000000    614 ns/op    72 B/op   8 allocs/op   1.23 s   14400 KB   76.75 ns/alloc
BenchmarkIkeaUnmarshal-8                    2000000    782 ns/op   160 B/op  11 allocs/op   1.56 s   32000 KB   71.09 ns/alloc
---
totals:
BenchmarkGencodeUnsafe-8                   30000000    245 ns/op   144 B/op   4 allocs/op   7.35 s  432000 KB   61.25 ns/alloc
BenchmarkXDR2-8                            20000000    299 ns/op    96 B/op   3 allocs/op   5.98 s  192000 KB   99.67 ns/alloc
BenchmarkColfer-8                          20000000    307 ns/op   176 B/op   4 allocs/op   6.14 s  352000 KB   76.75 ns/alloc
BenchmarkGotinyNoTime-8                    20000000    327 ns/op    96 B/op   3 allocs/op   6.54 s  192000 KB  109.00 ns/alloc
BenchmarkGotiny-8                          20000000    348 ns/op   112 B/op   3 allocs/op   6.96 s  224000 KB  116.00 ns/alloc
BenchmarkGencode-8                         20000000    349 ns/op   192 B/op   5 allocs/op   6.98 s  384000 KB   69.80 ns/alloc
BenchmarkGogoprotobuf-8                    20000000    378 ns/op   160 B/op   4 allocs/op   7.56 s  320000 KB   94.50 ns/alloc
BenchmarkMsgp-8                            15000000    497 ns/op   240 B/op   4 allocs/op   7.46 s  360000 KB  124.25 ns/alloc
BenchmarkFlatBuffers-8                     10000000    567 ns/op   112 B/op   3 allocs/op   5.67 s  112000 KB  189.00 ns/alloc
BenchmarkCapNProto-8                        6000000    888 ns/op   256 B/op   8 allocs/op   5.33 s  153600 KB  111.00 ns/alloc
BenchmarkGoprotobuf-8                       5000000   1035 ns/op   296 B/op  12 allocs/op   5.17 s  148000 KB   86.25 ns/alloc
BenchmarkHprose2-8                          5000000   1200 ns/op   144 B/op   4 allocs/op   6.00 s   72000 KB  300.00 ns/alloc
BenchmarkIkea-8                             4000000   1396 ns/op   232 B/op  19 allocs/op   5.58 s   92800 KB   73.47 ns/alloc
BenchmarkCapNProto2-8                       4000000   1712 ns/op   564 B/op   9 allocs/op   6.85 s  225600 KB  190.22 ns/alloc
BenchmarkGoAvro2Binary-8                    4000000   1773 ns/op  1086 B/op  24 allocs/op   7.09 s  434400 KB   73.88 ns/alloc
BenchmarkProtobuf-8                         4000000   1789 ns/op   392 B/op  17 allocs/op   7.16 s  156800 KB  105.24 ns/alloc
BenchmarkGob-8                              4000000   1887 ns/op   160 B/op   5 allocs/op   7.55 s   64000 KB  377.40 ns/alloc
BenchmarkHprose-8                           3000000   1980 ns/op   793 B/op  18 allocs/op   5.94 s  237900 KB  110.00 ns/alloc
BenchmarkUgorjiCodecMsgpack-8               2000000   2348 ns/op  1010 B/op  12 allocs/op   4.70 s  202000 KB  195.67 ns/alloc
BenchmarkUgorjiCodecBinc-8                  2000000   2603 ns/op  1266 B/op  16 allocs/op   5.21 s  253200 KB  162.69 ns/alloc
BenchmarkEasyJson-8                         2000000   2715 ns/op   944 B/op   9 allocs/op   5.43 s  188800 KB  301.67 ns/alloc
BenchmarkJsonIter-8                         2000000   2777 ns/op   448 B/op  10 allocs/op   5.55 s   89600 KB  277.70 ns/alloc
BenchmarkBinary-8                           2000000   2910 ns/op   669 B/op  42 allocs/op   5.82 s  133800 KB   69.29 ns/alloc
BenchmarkXdr-8                              2000000   3111 ns/op   696 B/op  32 allocs/op   6.22 s  139200 KB   97.22 ns/alloc
BenchmarkBson-8                             2000000   3211 ns/op   636 B/op  29 allocs/op   6.42 s  127200 KB  110.72 ns/alloc
BenchmarkVmihailencoMsgpack-8               2000000   3569 ns/op   752 B/op  19 allocs/op   7.14 s  150400 KB  187.84 ns/alloc
BenchmarkSereal-8                           1000000   5320 ns/op  1920 B/op  55 allocs/op   5.32 s  192000 KB   96.73 ns/alloc
BenchmarkGoAvro2Text-8                      1000000   5390 ns/op  2133 B/op  54 allocs/op   5.39 s  213300 KB   99.81 ns/alloc
BenchmarkJson-8                             1300000   6299 ns/op  1023 B/op  13 allocs/op   8.19 s  132990 KB  484.54 ns/alloc
BenchmarkGoAvro-8                            700000   8229 ns/op  4466 B/op 118 allocs/op   5.76 s  312620 KB   69.74 ns/alloc
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
