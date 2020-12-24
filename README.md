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
- [github.com/200sc/bebop](https://github.com/200sc/bebop) (generated code)

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

2020-12-24 Results with Go 1.15.6 windows/amd64 on a 3.4 GHz Intel Core i7 16GB

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGotinyMarshal-8                 |    8633204 |    142 ns/op |    48 |   0 |   1.23 |   41439 |    0.00
BenchmarkGotinyUnmarshal-8               |    4225356 |    282 ns/op |    48 | 112 |   1.19 |   20281 |    2.52
BenchmarkGotinyNoTimeMarshal-8           |    8633067 |    143 ns/op |    48 |   0 |   1.23 |   41438 |    0.00
BenchmarkGotinyNoTimeUnmarshal-8         |    4135863 |    290 ns/op |    48 |  96 |   1.20 |   19852 |    3.02
BenchmarkMsgpMarshal-8                   |    6243512 |    196 ns/op |    97 | 128 |   1.22 |   60562 |    1.53
BenchmarkMsgpUnmarshal-8                 |    3463214 |    385 ns/op |    97 | 112 |   1.33 |   33593 |    3.44
BenchmarkVmihailencoMsgpackMarshal-8     |     923048 |   1427 ns/op |   100 | 400 |   1.32 |    9230 |    3.57
BenchmarkVmihailencoMsgpackUnmarshal-8   |     615320 |   2583 ns/op |   100 | 416 |   1.59 |    6153 |    6.21
BenchmarkJsonMarshal-8                   |     240012 |   5200 ns/op |   150 | 944 |   1.25 |    3600 |    5.51
BenchmarkJsonUnmarshal-8                 |     292686 |   4766 ns/op |   150 | 344 |   1.39 |    4390 |   13.85
BenchmarkJsonIterMarshal-8               |     206896 |   6107 ns/op |   150 | 1113 |   1.26 |    3103 |    5.49
BenchmarkJsonIterUnmarshal-8             |     299983 |   5474 ns/op |   150 | 447 |   1.64 |    4499 |   12.25
BenchmarkEasyJsonMarshal-8               |     558186 |   1836 ns/op |   150 | 784 |   1.02 |    8372 |    2.34
BenchmarkEasyJsonUnmarshal-8             |     727201 |   1671 ns/op |   150 | 160 |   1.22 |   10908 |   10.44
BenchmarkBsonMarshal-8                   |    1000000 |   1276 ns/op |   110 | 392 |   1.28 |   11000 |    3.26
BenchmarkBsonUnmarshal-8                 |     648606 |   1751 ns/op |   110 | 232 |   1.14 |    7134 |    7.55
BenchmarkGobMarshal-8                    |    1000000 |   1328 ns/op |    63 |  48 |   1.33 |    6360 |   27.67
BenchmarkGobUnmarshal-8                  |     923083 |   1110 ns/op |    63 | 112 |   1.02 |    5880 |    9.91
BenchmarkXDRMarshal-8                    |     615304 |   2030 ns/op |    92 | 456 |   1.25 |    5660 |    4.45
BenchmarkXDRUnmarshal-8                  |     666610 |   2300 ns/op |    92 | 240 |   1.53 |    6132 |    9.58
BenchmarkUgorjiCodecMsgpackMarshal-8     |     959968 |   1404 ns/op |    91 | 1312 |   1.35 |    8735 |    1.07
BenchmarkUgorjiCodecMsgpackUnmarshal-8   |     960084 |   1588 ns/op |    91 | 496 |   1.52 |    8736 |    3.20
BenchmarkUgorjiCodecBincMarshal-8        |     819354 |   1724 ns/op |    95 | 1328 |   1.41 |    7783 |    1.30
BenchmarkUgorjiCodecBincUnmarshal-8      |     666692 |   1773 ns/op |    95 | 656 |   1.18 |    6333 |    2.70
BenchmarkSerealMarshal-8                 |     387104 |   3702 ns/op |   132 | 904 |   1.43 |    5109 |    4.10
BenchmarkSerealUnmarshal-8               |     295628 |   3493 ns/op |   132 | 1008 |   1.03 |    3902 |    3.47
BenchmarkBinaryMarshal-8                 |     888888 |   1464 ns/op |    61 | 320 |   1.30 |    5422 |    4.58
BenchmarkBinaryUnmarshal-8               |     666525 |   2257 ns/op |    61 | 320 |   1.50 |    4065 |    7.05
BenchmarkFlatBuffersMarshal-8            |    2795731 |    495 ns/op |    95 |   0 |   1.38 |   26643 |    0.00
BenchmarkFlatBuffersUnmarshal-8          |    3399423 |    347 ns/op |    95 | 112 |   1.18 |   32362 |    3.10
BenchmarkCapNProtoMarshal-8              |    2142858 |    586 ns/op |    96 |  56 |   1.26 |   20571 |   10.46
BenchmarkCapNProtoUnmarshal-8            |    2318839 |    601 ns/op |    96 | 200 |   1.39 |   22260 |    3.00
BenchmarkCapNProto2Marshal-8             |    1595745 |    774 ns/op |    96 | 244 |   1.24 |   15319 |    3.17
BenchmarkCapNProto2Unmarshal-8           |    1463414 |    838 ns/op |    96 | 320 |   1.23 |   14048 |    2.62
BenchmarkHproseMarshal-8                 |     960214 |   1078 ns/op |    84 | 484 |   1.04 |    8065 |    2.23
BenchmarkHproseUnmarshal-8               |     999924 |   1286 ns/op |    85 | 319 |   1.29 |    8529 |    4.03
BenchmarkHprose2Marshal-8                |    2122011 |    583 ns/op |    84 |   0 |   1.24 |   17846 |    0.00
BenchmarkHprose2Unmarshal-8              |    2020216 |    606 ns/op |    85 | 144 |   1.22 |   17232 |    4.21
BenchmarkProtobufMarshal-8               |    1421798 |    838 ns/op |    52 | 152 |   1.19 |    7393 |    5.51
BenchmarkProtobufUnmarshal-8             |    1458079 |    959 ns/op |    52 | 192 |   1.40 |    7582 |    4.99
BenchmarkGoprotobufMarshal-8             |    3265320 |    377 ns/op |    53 |  64 |   1.23 |   17306 |    5.89
BenchmarkGoprotobufUnmarshal-8           |    2181825 |    564 ns/op |    53 | 168 |   1.23 |   11563 |    3.36
BenchmarkGogoprotobufMarshal-8           |    6722809 |    181 ns/op |    53 |  64 |   1.22 |   35630 |    2.83
BenchmarkGogoprotobufUnmarshal-8         |    4829002 |    240 ns/op |    53 |  96 |   1.16 |   25593 |    2.50
BenchmarkColferMarshal-8                 |    9266287 |    124 ns/op |    51 |  64 |   1.15 |   47350 |    1.94
BenchmarkColferUnmarshal-8               |    5607489 |    208 ns/op |    51 | 112 |   1.17 |   28598 |    1.86
BenchmarkGencodeMarshal-8                |    6349306 |    193 ns/op |    53 |  80 |   1.23 |   33651 |    2.41
BenchmarkGencodeUnmarshal-8              |    5483082 |    221 ns/op |    53 | 112 |   1.21 |   29060 |    1.97
BenchmarkGencodeUnsafeMarshal-8          |   10526334 |    113 ns/op |    46 |  48 |   1.19 |   48421 |    2.35
BenchmarkGencodeUnsafeUnmarshal-8        |    6837264 |    167 ns/op |    46 |  96 |   1.14 |   31451 |    1.74
BenchmarkXDR2Marshal-8                   |    7476676 |    178 ns/op |    60 |  64 |   1.33 |   44860 |    2.78
BenchmarkXDR2Unmarshal-8                 |    9061222 |    130 ns/op |    60 |  32 |   1.18 |   54367 |    4.06
BenchmarkGoAvroMarshal-8                 |     444445 |   2787 ns/op |    47 | 1008 |   1.24 |    2088 |    2.76
BenchmarkGoAvroUnmarshal-8               |     172666 |   8250 ns/op |    47 | 3328 |   1.42 |     811 |    2.48
BenchmarkGoAvro2TextMarshal-8            |     380944 |   4024 ns/op |   134 | 1320 |   1.53 |    5104 |    3.05
BenchmarkGoAvro2TextUnmarshal-8          |     375000 |   3003 ns/op |   134 | 799 |   1.13 |    5025 |    3.76
BenchmarkGoAvro2BinaryMarshal-8          |    1262494 |    993 ns/op |    47 | 488 |   1.25 |    5933 |    2.03
BenchmarkGoAvro2BinaryUnmarshal-8        |    1000000 |   1088 ns/op |    47 | 560 |   1.09 |    4700 |    1.94
BenchmarkIkeaMarshal-8                   |    1753094 |    679 ns/op |    55 |  72 |   1.19 |    9642 |    9.43
BenchmarkIkeaUnmarshal-8                 |    1346046 |    910 ns/op |    55 | 160 |   1.22 |    7403 |    5.69
BenchmarkShamatonMapMsgpackMarshal-8     |    1443176 |    842 ns/op |    92 | 208 |   1.22 |   13277 |    4.05
BenchmarkShamatonMapMsgpackUnmarshal-8   |    1756918 |    701 ns/op |    92 | 144 |   1.23 |   16163 |    4.87
BenchmarkShamatonArrayMsgpackMarshal-8   |    1619637 |    763 ns/op |    50 | 176 |   1.24 |    8098 |    4.34
BenchmarkShamatonArrayMsgpackUnmarshal-8 |    2585856 |    484 ns/op |    50 | 144 |   1.25 |   12929 |    3.36
BenchmarkSSZNoTimeNoStringNoFloatAMarshal-8 |     269680 |   4462 ns/op |    55 | 440 |   1.20 |    1483 |   10.14
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal-8 |     165517 |   8006 ns/op |    55 | 1392 |   1.33 |     910 |    5.75
BenchmarkBebopaMarshal-8                 |    9839056 |    122 ns/op |    55 |  64 |   1.20 |   54114 |    1.91
BenchmarkBebopaUnmarshal-8               |   11940262 |    102 ns/op |    55 |  32 |   1.22 |   65671 |    3.19


Totals:


benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkBebopa-8                        |   21779318 |    224 ns/op |   110 |  96 |   4.88 |  239572 |    2.33
BenchmarkGencodeUnsafe-8                 |   17363598 |    280 ns/op |    92 | 144 |   4.86 |  159745 |    1.94
BenchmarkXDR2-8                          |   16537898 |    308 ns/op |   120 |  96 |   5.09 |  198454 |    3.21
BenchmarkColfer-8                        |   14873776 |    332 ns/op |   102 | 176 |   4.94 |  151861 |    1.89
BenchmarkGencode-8                       |   11832388 |    414 ns/op |   106 | 192 |   4.90 |  125423 |    2.16
BenchmarkGogoprotobuf-8                  |   11551811 |    421 ns/op |   106 | 160 |   4.86 |  122449 |    2.63
BenchmarkGotiny-8                        |   12858560 |    424 ns/op |    96 | 112 |   5.45 |  123442 |    3.79
BenchmarkGotinyNoTime-8                  |   12768930 |    433 ns/op |    96 |  96 |   5.53 |  122581 |    4.51
BenchmarkMsgp-8                          |    9706726 |    581 ns/op |   194 | 240 |   5.64 |  188310 |    2.42
BenchmarkFlatBuffers-8                   |    6195154 |    842 ns/op |   190 | 112 |   5.22 |  118017 |    7.52
BenchmarkGoprotobuf-8                    |    5447145 |    941 ns/op |   106 | 232 |   5.13 |   57739 |    4.06
BenchmarkCapNProto-8                     |    4461697 |   1187 ns/op |   192 | 256 |   5.30 |   85664 |    4.64
BenchmarkHprose2-8                       |    4142227 |   1189 ns/op |   169 | 144 |   4.93 |   70169 |    8.26
BenchmarkShamatonArrayMsgpack-8          |    4205493 |   1247 ns/op |   100 | 320 |   5.24 |   42054 |    3.90
BenchmarkShamatonMapMsgpack-8            |    3200094 |   1543 ns/op |   184 | 352 |   4.94 |   58881 |    4.38
BenchmarkIkea-8                          |    3099140 |   1589 ns/op |   110 | 232 |   4.92 |   34090 |    6.85
BenchmarkCapNProto2-8                    |    3059159 |   1612 ns/op |   192 | 564 |   4.93 |   58735 |    2.86
BenchmarkProtobuf-8                      |    2879877 |   1797 ns/op |   104 | 344 |   5.18 |   29950 |    5.22
BenchmarkGoAvro2Binary-8                 |    2262494 |   2081 ns/op |    94 | 1048 |   4.71 |   21267 |    1.99
BenchmarkHprose-8                        |    1960138 |   2364 ns/op |   169 | 803 |   4.63 |   33185 |    2.94
BenchmarkGob-8                           |    1923083 |   2438 ns/op |   127 | 160 |   4.69 |   24480 |   15.24
BenchmarkUgorjiCodecMsgpack-8            |    1920052 |   2992 ns/op |   182 | 1808 |   5.74 |   34944 |    1.65
BenchmarkBson-8                          |    1648606 |   3027 ns/op |   220 | 624 |   4.99 |   36269 |    4.85
BenchmarkUgorjiCodecBinc-8               |    1486046 |   3497 ns/op |   190 | 1984 |   5.20 |   28234 |    1.76
BenchmarkEasyJson-8                      |    1285387 |   3507 ns/op |   300 | 944 |   4.51 |   38561 |    3.72
BenchmarkBinary-8                        |    1555413 |   3721 ns/op |   122 | 640 |   5.79 |   18976 |    5.81
BenchmarkVmihailencoMsgpack-8            |    1538368 |   4010 ns/op |   200 | 816 |   6.17 |   30767 |    4.91
BenchmarkXDR-8                           |    1281914 |   4330 ns/op |   184 | 696 |   5.55 |   23587 |    6.22
BenchmarkGoAvro2Text-8                   |     755944 |   7027 ns/op |   268 | 2119 |   5.31 |   20259 |    3.32
BenchmarkSereal-8                        |     682732 |   7195 ns/op |   264 | 1912 |   4.91 |   18024 |    3.76
BenchmarkJson-8                          |     532698 |   9966 ns/op |   300 | 1288 |   5.31 |   15980 |    7.74
BenchmarkGoAvro-8                        |     617111 |  11037 ns/op |    94 | 4336 |   6.81 |    5800 |    2.55
BenchmarkJsonIter-8                      |     506879 |  11581 ns/op |   300 | 1560 |   5.87 |   15206 |    7.42
BenchmarkSSZNoTimeNoStringNoFloatA-8     |     435197 |  12468 ns/op |   110 | 1832 |   5.43 |    4787 |    6.81



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
