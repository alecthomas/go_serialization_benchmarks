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

2020-12-30 Results with Go 1.15.6 windows/amd64 on a 3.4 GHz Intel Core i7 16GB

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGotinyMarshal-8                 |    8823405 |    130 ns/op |    48 |   0 |   1.15 |   42352 |    0.00
BenchmarkGotinyUnmarshal-8               |    4536901 |    267 ns/op |    48 | 112 |   1.21 |   21777 |    2.38
BenchmarkGotinyNoTimeMarshal-8           |    9448945 |    136 ns/op |    48 |   0 |   1.29 |   45354 |    0.00
BenchmarkGotinyNoTimeUnmarshal-8         |    4878072 |    241 ns/op |    48 |  96 |   1.18 |   23414 |    2.51
BenchmarkMsgpMarshal-8                   |    6896487 |    174 ns/op |    97 | 128 |   1.20 |   66895 |    1.36
BenchmarkMsgpUnmarshal-8                 |    3833878 |    314 ns/op |    97 | 112 |   1.20 |   37188 |    2.80
BenchmarkVmihailencoMsgpackMarshal-8     |    1000000 |   1116 ns/op |   100 | 400 |   1.12 |   10000 |    2.79
BenchmarkVmihailencoMsgpackUnmarshal-8   |     727237 |   1722 ns/op |   100 | 416 |   1.25 |    7272 |    4.14
BenchmarkJsonMarshal-8                   |     555409 |   2212 ns/op |   150 | 208 |   1.23 |    8331 |   10.63
BenchmarkJsonUnmarshal-8                 |     262119 |   4574 ns/op |   149 | 391 |   1.20 |    3905 |   11.70
BenchmarkJsonIterMarshal-8               |     521743 |   2392 ns/op |   139 | 248 |   1.25 |    7252 |    9.65
BenchmarkJsonIterUnmarshal-8             |     648538 |   1874 ns/op |   139 | 264 |   1.22 |    9014 |    7.10
BenchmarkEasyJsonMarshal-8               |     705790 |   1789 ns/op |   149 | 895 |   1.26 |   10516 |    2.00
BenchmarkEasyJsonUnmarshal-8             |     705720 |   1724 ns/op |   150 | 288 |   1.22 |   10585 |    5.99
BenchmarkBsonMarshal-8                   |    1000000 |   1171 ns/op |   110 | 392 |   1.17 |   11000 |    2.99
BenchmarkBsonUnmarshal-8                 |     749938 |   1694 ns/op |   110 | 232 |   1.27 |    8249 |    7.30
BenchmarkGobMarshal-8                    |    1414255 |    834 ns/op |    63 |  48 |   1.18 |    8994 |   17.38
BenchmarkGobUnmarshal-8                  |    1323007 |    900 ns/op |    63 | 112 |   1.19 |    8414 |    8.04
BenchmarkXDRMarshal-8                    |     705927 |   1836 ns/op |    92 | 456 |   1.30 |    6494 |    4.03
BenchmarkXDRUnmarshal-8                  |     489799 |   2255 ns/op |    90 | 235 |   1.10 |    4452 |    9.60
BenchmarkUgorjiCodecMsgpackMarshal-8     |     959930 |   1422 ns/op |    91 | 1312 |   1.37 |    8735 |    1.08
BenchmarkUgorjiCodecMsgpackUnmarshal-8   |     857185 |   1398 ns/op |    91 | 496 |   1.20 |    7800 |    2.82
BenchmarkUgorjiCodecBincMarshal-8        |     857001 |   1468 ns/op |    95 | 1328 |   1.26 |    8141 |    1.11
BenchmarkUgorjiCodecBincUnmarshal-8      |     705873 |   1587 ns/op |    95 | 656 |   1.12 |    6705 |    2.42
BenchmarkSerealMarshal-8                 |     420970 |   2936 ns/op |   132 | 904 |   1.24 |    5556 |    3.25
BenchmarkSerealUnmarshal-8               |     363639 |   3377 ns/op |   132 | 1008 |   1.23 |    4800 |    3.35
BenchmarkBinaryMarshal-8                 |     922891 |   1364 ns/op |    61 | 320 |   1.26 |    5629 |    4.26
BenchmarkBinaryUnmarshal-8               |     799935 |   1511 ns/op |    61 | 320 |   1.21 |    4879 |    4.72
BenchmarkFlatBuffersMarshal-8            |    4116678 |    298 ns/op |    95 |   0 |   1.23 |   39190 |    0.00
BenchmarkFlatBuffersUnmarshal-8          |    4444442 |    265 ns/op |    95 | 112 |   1.18 |   42311 |    2.37
BenchmarkCapNProtoMarshal-8              |    3183031 |    386 ns/op |    96 |  56 |   1.23 |   30557 |    6.89
BenchmarkCapNProtoUnmarshal-8            |    2724205 |    443 ns/op |    96 | 200 |   1.21 |   26152 |    2.21
BenchmarkCapNProto2Marshal-8             |    2035632 |    586 ns/op |    96 | 244 |   1.19 |   19542 |    2.40
BenchmarkCapNProto2Unmarshal-8           |    1560622 |    778 ns/op |    96 | 320 |   1.21 |   14981 |    2.43
BenchmarkHproseMarshal-8                 |    1313616 |    915 ns/op |    85 | 402 |   1.20 |   11205 |    2.28
BenchmarkHproseUnmarshal-8               |     960014 |   1195 ns/op |    85 | 319 |   1.15 |    8188 |    3.75
BenchmarkHprose2Marshal-8                |    2063629 |    604 ns/op |    85 |   0 |   1.25 |   17602 |    0.00
BenchmarkHprose2Unmarshal-8              |    2015106 |    609 ns/op |    85 | 144 |   1.23 |   17188 |    4.23
BenchmarkProtobufMarshal-8               |    1537672 |    801 ns/op |    52 | 152 |   1.23 |    7995 |    5.27
BenchmarkProtobufUnmarshal-8             |    1533541 |    790 ns/op |    52 | 192 |   1.21 |    7974 |    4.11
BenchmarkGoprotobufMarshal-8             |    3545023 |    317 ns/op |    53 |  64 |   1.12 |   18788 |    4.95
BenchmarkGoprotobufUnmarshal-8           |    2456142 |    481 ns/op |    53 | 168 |   1.18 |   13017 |    2.86
BenchmarkGogoprotobufMarshal-8           |    8163325 |    147 ns/op |    53 |  64 |   1.20 |   43265 |    2.30
BenchmarkGogoprotobufUnmarshal-8         |    5172433 |    230 ns/op |    53 |  96 |   1.19 |   27413 |    2.40
BenchmarkColferMarshal-8                 |    9599860 |    124 ns/op |    51 |  64 |   1.19 |   49055 |    1.94
BenchmarkColferUnmarshal-8               |    6220048 |    197 ns/op |    50 | 112 |   1.23 |   31100 |    1.76
BenchmarkGencodeMarshal-8                |    6557677 |    186 ns/op |    53 |  80 |   1.22 |   34755 |    2.33
BenchmarkGencodeUnmarshal-8              |    5417618 |    222 ns/op |    53 | 112 |   1.20 |   28713 |    1.98
BenchmarkGencodeUnsafeMarshal-8          |   11650496 |     98 ns/op |    46 |  48 |   1.15 |   53592 |    2.05
BenchmarkGencodeUnsafeUnmarshal-8        |    7637104 |    161 ns/op |    46 |  96 |   1.23 |   35130 |    1.68
BenchmarkXDR2Marshal-8                   |    7619090 |    159 ns/op |    60 |  64 |   1.21 |   45714 |    2.48
BenchmarkXDR2Unmarshal-8                 |    9031648 |    131 ns/op |    60 |  32 |   1.18 |   54189 |    4.09
BenchmarkGoAvroMarshal-8                 |     436340 |   2828 ns/op |    47 | 1008 |   1.23 |    2050 |    2.81
BenchmarkGoAvroUnmarshal-8               |     179101 |   6962 ns/op |    47 | 3328 |   1.25 |     841 |    2.09
BenchmarkGoAvro2TextMarshal-8            |     399990 |   2975 ns/op |   134 | 1320 |   1.19 |    5359 |    2.25
BenchmarkGoAvro2TextUnmarshal-8          |     413822 |   2826 ns/op |   134 | 799 |   1.17 |    5545 |    3.54
BenchmarkGoAvro2BinaryMarshal-8          |    1245332 |    950 ns/op |    47 | 488 |   1.18 |    5853 |    1.95
BenchmarkGoAvro2BinaryUnmarshal-8        |    1000000 |   1092 ns/op |    47 | 560 |   1.09 |    4700 |    1.95
BenchmarkIkeaMarshal-8                   |    1772526 |    670 ns/op |    55 |  72 |   1.19 |    9748 |    9.31
BenchmarkIkeaUnmarshal-8                 |    1468875 |    871 ns/op |    55 | 160 |   1.28 |    8078 |    5.44
BenchmarkShamatonMapMsgpackMarshal-8     |    1434548 |    819 ns/op |    92 | 208 |   1.17 |   13197 |    3.94
BenchmarkShamatonMapMsgpackUnmarshal-8   |    1675980 |    738 ns/op |    92 | 144 |   1.24 |   15419 |    5.12
BenchmarkShamatonArrayMsgpackMarshal-8   |    1523809 |    758 ns/op |    50 | 176 |   1.16 |    7619 |    4.31
BenchmarkShamatonArrayMsgpackUnmarshal-8 |    2542365 |    483 ns/op |    50 | 144 |   1.23 |   12711 |    3.35
BenchmarkSSZNoTimeNoStringNoFloatAMarshal-8 |     260883 |   4922 ns/op |    55 | 440 |   1.28 |    1434 |   11.19
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal-8 |     157892 |   7694 ns/op |    55 | 1392 |   1.21 |     868 |    5.53
BenchmarkMumMarshal-8                    |   12371503 |     97 ns/op |    48 |   0 |   1.21 |   59383 |    0.00
BenchmarkMumUnmarshal-8                  |    5517202 |    216 ns/op |    48 |  80 |   1.19 |   26482 |    2.70
BenchmarkBebopMarshal-8                  |    9795806 |    124 ns/op |    55 |  64 |   1.21 |   53876 |    1.94
BenchmarkBebopUnmarshal-8                |   11537041 |    104 ns/op |    55 |  32 |   1.20 |   63453 |    3.25

Totals:


benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------

BenchmarkBebop-8                         |   21332847 |    228 ns/op |   110 |  96 |   4.86 |  234661 |    2.38
BenchmarkGencodeUnsafe-8                 |   19287600 |    259 ns/op |    92 | 144 |   5.01 |  177445 |    1.80
BenchmarkXDR2-8                          |   16650738 |    290 ns/op |   120 |  96 |   4.83 |  199808 |    3.02
BenchmarkMum-8                           |   17888705 |    313 ns/op |    96 |  80 |   5.61 |  171731 |    3.92
BenchmarkColfer-8                        |   15819908 |    321 ns/op |   101 | 176 |   5.08 |  159939 |    1.82
BenchmarkGogoprotobuf-8                  |   13335758 |    377 ns/op |   106 | 160 |   5.03 |  141359 |    2.36
BenchmarkGotiny-8                        |   13360306 |    397 ns/op |    96 | 112 |   5.30 |  128258 |    3.54
BenchmarkGencode-8                       |   11975295 |    408 ns/op |   106 | 192 |   4.89 |  126938 |    2.12
BenchmarkMsgp-8                          |   10730365 |    488 ns/op |   194 | 240 |   5.24 |  208169 |    2.03
BenchmarkFlatBuffers-8                   |    8561120 |    563 ns/op |   190 | 112 |   4.82 |  163003 |    5.03
BenchmarkGoprotobuf-8                    |    6001165 |    798 ns/op |   106 | 232 |   4.79 |   63612 |    3.44
BenchmarkCapNProto-8                     |    5907236 |    829 ns/op |   192 | 256 |   4.90 |  113418 |    3.24
BenchmarkHprose2-8                       |    4078735 |   1213 ns/op |   170 | 144 |   4.95 |   69583 |    8.42
BenchmarkShamatonArrayMsgpack-8          |    4066174 |   1241 ns/op |   100 | 320 |   5.05 |   40661 |    3.88
BenchmarkCapNProto2-8                    |    3596254 |   1364 ns/op |   192 | 564 |   4.91 |   69048 |    2.42
BenchmarkIkea-8                          |    3241401 |   1541 ns/op |   110 | 232 |   4.99 |   35655 |    6.64
BenchmarkShamatonMapMsgpack-8            |    3110528 |   1557 ns/op |   184 | 352 |   4.84 |   57233 |    4.42
BenchmarkProtobuf-8                      |    3071213 |   1591 ns/op |   104 | 344 |   4.89 |   31940 |    4.62
BenchmarkGob-8                           |    2737262 |   1734 ns/op |   127 | 160 |   4.75 |   34817 |   10.84
BenchmarkGoAvro2Binary-8                 |    2245332 |   2042 ns/op |    94 | 1048 |   4.58 |   21106 |    1.95
BenchmarkHprose-8                        |    2273630 |   2110 ns/op |   170 | 721 |   4.80 |   38788 |    2.93
BenchmarkUgorjiCodecMsgpack-8            |    1817115 |   2820 ns/op |   182 | 1808 |   5.12 |   33071 |    1.56
BenchmarkVmihailencoMsgpack-8            |    1727237 |   2838 ns/op |   200 | 816 |   4.90 |   34544 |    3.48
BenchmarkBson-8                          |    1749938 |   2865 ns/op |   220 | 624 |   5.01 |   38498 |    4.59
BenchmarkBinary-8                        |    1722826 |   2875 ns/op |   122 | 640 |   4.95 |   21018 |    4.49
BenchmarkUgorjiCodecBinc-8               |    1562874 |   3055 ns/op |   190 | 1984 |   4.77 |   29694 |    1.54
BenchmarkEasyJson-8                      |    1411510 |   3513 ns/op |   299 | 1183 |   4.96 |   42204 |    2.97
BenchmarkXDR-8                           |    1195726 |   4091 ns/op |   182 | 691 |   4.89 |   21869 |    5.92
BenchmarkJsonIter-8                      |    1170281 |   4266 ns/op |   278 | 512 |   4.99 |   32533 |    8.33
BenchmarkGoAvro2Text-8                   |     813812 |   5801 ns/op |   268 | 2119 |   4.72 |   21810 |    2.74
BenchmarkSereal-8                        |     784609 |   6313 ns/op |   264 | 1912 |   4.95 |   20713 |    3.30
BenchmarkJson-8                          |     817528 |   6786 ns/op |   299 | 599 |   5.55 |   24444 |   11.33
BenchmarkGoAvro-8                        |     615441 |   9790 ns/op |    94 | 4336 |   6.03 |    5785 |    2.26
BenchmarkSSZNoTimeNoStringNoFloatA-8     |     418775 |  12616 ns/op |   110 | 1832 |   5.28 |    4606 |    6.89



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
