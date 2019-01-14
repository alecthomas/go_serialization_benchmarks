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

2019-01-14 Results with Go 1.11 on a single-core 2.1 GHz Intel Xeon Skylake based VPS:

```
benchmark                                      iter    time/iter   bytes/op     allocs/op  tt.time   tt.bytes       time/alloc
---------                                      ----    ---------   --------     ---------  -------   --------       ----------
BenchmarkGotinyMarshal                     10000000    148 ns/op     0 B/op   0 allocs/op   1.48 s       0 KB    0.00 ns/alloc
BenchmarkGotinyUnmarshal                    5000000    271 ns/op   112 B/op   3 allocs/op   1.35 s   56000 KB   90.33 ns/alloc
BenchmarkGotinyNoTimeMarshal               10000000    136 ns/op     0 B/op   0 allocs/op   1.36 s       0 KB    0.00 ns/alloc
BenchmarkGotinyNoTimeUnmarshal              5000000    260 ns/op    96 B/op   3 allocs/op   1.30 s   48000 KB   86.67 ns/alloc
BenchmarkMsgpMarshal                       10000000    219 ns/op   128 B/op   1 allocs/op   2.19 s  128000 KB  219.00 ns/alloc
BenchmarkMsgpUnmarshal                      5000000    394 ns/op   112 B/op   3 allocs/op   1.97 s   56000 KB  131.33 ns/alloc
BenchmarkVmihailencoMsgpackMarshal          1000000   2147 ns/op   368 B/op   6 allocs/op   2.15 s   36800 KB  357.83 ns/alloc
BenchmarkVmihailencoMsgpackUnmarshal         500000   2317 ns/op   384 B/op  13 allocs/op   1.16 s   19200 KB  178.23 ns/alloc
BenchmarkJsonMarshal                        1000000   2389 ns/op   304 B/op   4 allocs/op   2.39 s   30400 KB  597.25 ns/alloc
BenchmarkJsonUnmarshal                       300000   4841 ns/op   359 B/op   7 allocs/op   1.45 s   10770 KB  691.57 ns/alloc
BenchmarkJsonIterMarshal                    1000000   1571 ns/op   200 B/op   3 allocs/op   1.57 s   20000 KB  523.67 ns/alloc
BenchmarkJsonIterUnmarshal                  1000000   1814 ns/op   248 B/op   7 allocs/op   1.81 s   24800 KB  259.14 ns/alloc
BenchmarkEasyJsonMarshal                    1000000   1656 ns/op   784 B/op   5 allocs/op   1.66 s   78400 KB  331.20 ns/alloc
BenchmarkEasyJsonUnmarshal                  1000000   1623 ns/op   160 B/op   4 allocs/op   1.62 s   16000 KB  405.75 ns/alloc
BenchmarkBsonMarshal                        1000000   1585 ns/op   392 B/op  10 allocs/op   1.58 s   39200 KB  158.50 ns/alloc
BenchmarkBsonUnmarshal                      1000000   2116 ns/op   244 B/op  19 allocs/op   2.12 s   24400 KB  111.37 ns/alloc
BenchmarkGobMarshal                         1000000   1061 ns/op    48 B/op   2 allocs/op   1.06 s    4800 KB  530.50 ns/alloc
BenchmarkGobUnmarshal                       1000000   1114 ns/op   112 B/op   3 allocs/op   1.11 s   11200 KB  371.33 ns/alloc
BenchmarkXdrMarshal                         1000000   2172 ns/op   456 B/op  21 allocs/op   2.17 s   45600 KB  103.43 ns/alloc
BenchmarkXdrUnmarshal                       1000000   1826 ns/op   239 B/op  11 allocs/op   1.83 s   23900 KB  166.00 ns/alloc
BenchmarkUgorjiCodecMsgpackMarshal          1000000   1577 ns/op  1280 B/op   4 allocs/op   1.58 s  128000 KB  394.25 ns/alloc
BenchmarkUgorjiCodecMsgpackUnmarshal        1000000   1578 ns/op   464 B/op   5 allocs/op   1.58 s   46400 KB  315.60 ns/alloc
BenchmarkUgorjiCodecBincMarshal             1000000   1733 ns/op  1328 B/op   5 allocs/op   1.73 s  132800 KB  346.60 ns/alloc
BenchmarkUgorjiCodecBincUnmarshal           1000000   1803 ns/op   704 B/op   8 allocs/op   1.80 s   70400 KB  225.38 ns/alloc
BenchmarkSerealMarshal                       500000   3547 ns/op   904 B/op  20 allocs/op   1.77 s   45200 KB  177.35 ns/alloc
BenchmarkSerealUnmarshal                     500000   3740 ns/op  1008 B/op  34 allocs/op   1.87 s   50400 KB  110.00 ns/alloc
BenchmarkBinaryMarshal                      1000000   1716 ns/op   334 B/op  20 allocs/op   1.72 s   33400 KB   85.80 ns/alloc
BenchmarkBinaryUnmarshal                    1000000   1827 ns/op   336 B/op  22 allocs/op   1.83 s   33600 KB   83.05 ns/alloc
BenchmarkFlatBuffersMarshal                 5000000    360 ns/op     0 B/op   0 allocs/op   1.80 s       0 KB    0.00 ns/alloc
BenchmarkFlatBuffersUnmarshal               5000000    297 ns/op   112 B/op   3 allocs/op   1.49 s   56000 KB   99.00 ns/alloc
BenchmarkCapNProtoMarshal                   3000000    484 ns/op    56 B/op   2 allocs/op   1.45 s   16800 KB  242.00 ns/alloc
BenchmarkCapNProtoUnmarshal                 3000000    507 ns/op   200 B/op   6 allocs/op   1.52 s   60000 KB   84.50 ns/alloc
BenchmarkCapNProto2Marshal                  2000000    757 ns/op   244 B/op   3 allocs/op   1.51 s   48800 KB  252.33 ns/alloc
BenchmarkCapNProto2Unmarshal                1000000   1104 ns/op   320 B/op   6 allocs/op   1.10 s   32000 KB  184.00 ns/alloc
BenchmarkHproseMarshal                      1000000   1222 ns/op   460 B/op   8 allocs/op   1.22 s   46000 KB  152.75 ns/alloc
BenchmarkHproseUnmarshal                    1000000   1233 ns/op   319 B/op  10 allocs/op   1.23 s   31900 KB  123.30 ns/alloc
BenchmarkHprose2Marshal                     2000000    668 ns/op     0 B/op   0 allocs/op   1.34 s       0 KB    0.00 ns/alloc
BenchmarkHprose2Unmarshal                   2000000    658 ns/op   144 B/op   4 allocs/op   1.32 s   28800 KB  164.50 ns/alloc
BenchmarkProtobufMarshal                    1000000   1237 ns/op   200 B/op   7 allocs/op   1.24 s   20000 KB  176.71 ns/alloc
BenchmarkProtobufUnmarshal                  2000000    963 ns/op   192 B/op  10 allocs/op   1.93 s   38400 KB   96.30 ns/alloc
BenchmarkGoprotobufMarshal                  3000000    507 ns/op    96 B/op   2 allocs/op   1.52 s   28800 KB  253.50 ns/alloc
BenchmarkGoprotobufUnmarshal                2000000    799 ns/op   200 B/op  10 allocs/op   1.60 s   40000 KB   79.90 ns/alloc
BenchmarkGogoprotobufMarshal               10000000    211 ns/op    64 B/op   1 allocs/op   2.11 s   64000 KB  211.00 ns/alloc
BenchmarkGogoprotobufUnmarshal              5000000    268 ns/op    96 B/op   3 allocs/op   1.34 s   48000 KB   89.33 ns/alloc
BenchmarkColferMarshal                     10000000    170 ns/op    64 B/op   1 allocs/op   1.70 s   64000 KB  170.00 ns/alloc
BenchmarkColferUnmarshal                    5000000    234 ns/op   112 B/op   3 allocs/op   1.17 s   56000 KB   78.00 ns/alloc
BenchmarkGencodeMarshal                    10000000    221 ns/op    80 B/op   2 allocs/op   2.21 s   80000 KB  110.50 ns/alloc
BenchmarkGencodeUnmarshal                  10000000    230 ns/op   112 B/op   3 allocs/op   2.30 s  112000 KB   76.67 ns/alloc
BenchmarkGencodeUnsafeMarshal              10000000    138 ns/op    48 B/op   1 allocs/op   1.38 s   48000 KB  138.00 ns/alloc
BenchmarkGencodeUnsafeUnmarshal            10000000    186 ns/op    96 B/op   3 allocs/op   1.86 s   96000 KB   62.00 ns/alloc
BenchmarkXDR2Marshal                       10000000    208 ns/op    64 B/op   1 allocs/op   2.08 s   64000 KB  208.00 ns/alloc
BenchmarkXDR2Unmarshal                     10000000    190 ns/op    32 B/op   2 allocs/op   1.90 s   32000 KB   95.00 ns/alloc
BenchmarkGoAvroMarshal                       500000   3486 ns/op  1030 B/op  31 allocs/op   1.74 s   51500 KB  112.45 ns/alloc
BenchmarkGoAvroUnmarshal                     200000   8082 ns/op  3437 B/op  87 allocs/op   1.62 s   68740 KB   92.90 ns/alloc
BenchmarkGoAvro2TextMarshal                  500000   3700 ns/op  1326 B/op  20 allocs/op   1.85 s   66300 KB  185.00 ns/alloc
BenchmarkGoAvro2TextUnmarshal                500000   3306 ns/op   806 B/op  33 allocs/op   1.65 s   40300 KB  100.18 ns/alloc
BenchmarkGoAvro2BinaryMarshal               1000000   1129 ns/op   510 B/op  11 allocs/op   1.13 s   51000 KB  102.64 ns/alloc
BenchmarkGoAvro2BinaryUnmarshal             1000000   1183 ns/op   576 B/op  13 allocs/op   1.18 s   57600 KB   91.00 ns/alloc
BenchmarkIkeaMarshal                        2000000    738 ns/op    72 B/op   8 allocs/op   1.48 s   14400 KB   92.25 ns/alloc
BenchmarkIkeaUnmarshal                      1000000   1001 ns/op   160 B/op  11 allocs/op   1.00 s   16000 KB   91.00 ns/alloc
BenchmarkShamatonMapMsgpackMarshal          2000000    890 ns/op   208 B/op   4 allocs/op   1.78 s   41600 KB  222.50 ns/alloc
BenchmarkShamatonMapMsgpackUnmarshal        2000000    782 ns/op   144 B/op   3 allocs/op   1.56 s   28800 KB  260.67 ns/alloc
BenchmarkShamatonArrayMsgpackMarshal        2000000    804 ns/op   176 B/op   4 allocs/op   1.61 s   35200 KB  201.00 ns/alloc
BenchmarkShamatonArrayMsgpackUnmarshal      2000000    554 ns/op   144 B/op   3 allocs/op   1.11 s   28800 KB  184.67 ns/alloc
---
totals:
BenchmarkGencodeUnsafe                     20000000    324 ns/op   144 B/op   4 allocs/op   6.48 s  288000 KB   81.00 ns/alloc
BenchmarkGotinyNoTime                      15000000    396 ns/op    96 B/op   3 allocs/op   5.94 s  144000 KB  132.00 ns/alloc
BenchmarkXDR2                              20000000    398 ns/op    96 B/op   3 allocs/op   7.96 s  192000 KB  132.67 ns/alloc
BenchmarkColfer                            15000000    404 ns/op   176 B/op   4 allocs/op   6.06 s  264000 KB  101.00 ns/alloc
BenchmarkGotiny                            15000000    419 ns/op   112 B/op   3 allocs/op   6.29 s  168000 KB  139.67 ns/alloc
BenchmarkGencode                           20000000    451 ns/op   192 B/op   5 allocs/op   9.02 s  384000 KB   90.20 ns/alloc
BenchmarkGogoprotobuf                      15000000    479 ns/op   160 B/op   4 allocs/op   7.18 s  240000 KB  119.75 ns/alloc
BenchmarkMsgp                              15000000    613 ns/op   240 B/op   4 allocs/op   9.20 s  360000 KB  153.25 ns/alloc
BenchmarkFlatBuffers                       10000000    657 ns/op   112 B/op   3 allocs/op   6.57 s  112000 KB  219.00 ns/alloc
BenchmarkCapNProto                          6000000    991 ns/op   256 B/op   8 allocs/op   5.95 s  153600 KB  123.88 ns/alloc
BenchmarkGoprotobuf                         5000000   1306 ns/op   296 B/op  12 allocs/op   6.53 s  148000 KB  108.83 ns/alloc
BenchmarkHprose2                            4000000   1326 ns/op   144 B/op   4 allocs/op   5.30 s   57600 KB  331.50 ns/alloc
BenchmarkShamatonArrayMsgpack               4000000   1358 ns/op   320 B/op   7 allocs/op   5.43 s  128000 KB  194.00 ns/alloc
BenchmarkShamatonMapMsgpack                 4000000   1672 ns/op   352 B/op   7 allocs/op   6.69 s  140800 KB  238.86 ns/alloc
BenchmarkIkea                               3000000   1739 ns/op   232 B/op  19 allocs/op   5.22 s   69600 KB   91.53 ns/alloc
BenchmarkCapNProto2                         3000000   1861 ns/op   564 B/op   9 allocs/op   5.58 s  169200 KB  206.78 ns/alloc
BenchmarkGob                                2000000   2175 ns/op   160 B/op   5 allocs/op   4.35 s   32000 KB  435.00 ns/alloc
BenchmarkProtobuf                           3000000   2200 ns/op   392 B/op  17 allocs/op   6.60 s  117600 KB  129.41 ns/alloc
BenchmarkGoAvro2Binary                      2000000   2312 ns/op  1086 B/op  24 allocs/op   4.62 s  217200 KB   96.33 ns/alloc
BenchmarkHprose                             2000000   2455 ns/op   779 B/op  18 allocs/op   4.91 s  155800 KB  136.39 ns/alloc
BenchmarkUgorjiCodecMsgpack                 2000000   3155 ns/op  1744 B/op   9 allocs/op   6.31 s  348800 KB  350.56 ns/alloc
BenchmarkEasyJson                           2000000   3279 ns/op   944 B/op   9 allocs/op   6.56 s  188800 KB  364.33 ns/alloc
BenchmarkJsonIter                           2000000   3385 ns/op   448 B/op  10 allocs/op   6.77 s   89600 KB  338.50 ns/alloc
BenchmarkUgorjiCodecBinc                    2000000   3536 ns/op  2032 B/op  13 allocs/op   7.07 s  406400 KB  272.00 ns/alloc
BenchmarkBinary                             2000000   3543 ns/op   670 B/op  42 allocs/op   7.09 s  134000 KB   84.36 ns/alloc
BenchmarkBson                               2000000   3701 ns/op   636 B/op  29 allocs/op   7.40 s  127200 KB  127.62 ns/alloc
BenchmarkXdr                                2000000   3998 ns/op   695 B/op  32 allocs/op   8.00 s  139000 KB  124.94 ns/alloc
BenchmarkVmihailencoMsgpack                 1500000   4464 ns/op   752 B/op  19 allocs/op   6.70 s  112800 KB  234.95 ns/alloc
BenchmarkGoAvro2Text                        1000000   7006 ns/op  2132 B/op  53 allocs/op   7.01 s  213200 KB  132.19 ns/alloc
BenchmarkJson                               1300000   7230 ns/op   663 B/op  11 allocs/op   9.40 s   86190 KB  657.27 ns/alloc
BenchmarkSereal                             1000000   7287 ns/op  1912 B/op  54 allocs/op   7.29 s  191200 KB  134.94 ns/alloc
BenchmarkGoAvro                              700000  11568 ns/op  4467 B/op 118 allocs/op   8.10 s  312690 KB   98.03 ns/alloc
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
