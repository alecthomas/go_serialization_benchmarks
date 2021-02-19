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
- [github.com/shamaton/msgpackgen](https://github.com/shamaton/msgpackgen) (generated code)

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

2021-2-19 Results with Go 1.15.7 darwin/amd64 on a 2.9 GHz Intel Core i5 8GB

benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
BenchmarkGotinyMarshal-4                      |   10749480 |    105 ns/op |        48 |          0 |   1.13 |        51597 |    0.00
BenchmarkGotinyUnmarshal-4                    |    5749803 |    214 ns/op |        48 |        112 |   1.23 |        27599 |    1.91
BenchmarkGotinyNoTimeMarshal-4                |   11001322 |     99 ns/op |        48 |          0 |   1.10 |        52806 |    0.00
BenchmarkGotinyNoTimeUnmarshal-4              |    6291321 |    193 ns/op |        48 |         96 |   1.21 |        30198 |    2.01
BenchmarkMsgpMarshal-4                        |    7611444 |    163 ns/op |        97 |        128 |   1.24 |        73831 |    1.27
BenchmarkMsgpUnmarshal-4                      |    4267449 |    288 ns/op |        97 |        112 |   1.23 |        41394 |    2.57
BenchmarkVmihailencoMsgpackMarshal-4          |    1243858 |    988 ns/op |       100 |        400 |   1.23 |        12438 |    2.47
BenchmarkVmihailencoMsgpackUnmarshal-4        |     894117 |   1594 ns/op |       100 |        416 |   1.43 |         8941 |    3.83
BenchmarkJsonMarshal-4                        |     656860 |   1797 ns/op |       149 |        208 |   1.18 |         9787 |    8.64
BenchmarkJsonUnmarshal-4                      |     369094 |   3526 ns/op |       149 |        391 |   1.30 |         5499 |    9.02
BenchmarkJsonIterMarshal-4                    |     643039 |   1798 ns/op |       138 |        248 |   1.16 |         8873 |    7.25
BenchmarkJsonIterUnmarshal-4                  |     812228 |   1613 ns/op |       138 |        263 |   1.31 |        11208 |    6.13
BenchmarkEasyJsonMarshal-4                    |     812439 |   1511 ns/op |       149 |        896 |   1.23 |        12105 |    1.69
BenchmarkEasyJsonUnmarshal-4                  |     767331 |   1589 ns/op |       149 |        287 |   1.22 |        11433 |    5.54
BenchmarkBsonMarshal-4                        |    1000000 |   1046 ns/op |       110 |        392 |   1.05 |        11000 |    2.67
BenchmarkBsonUnmarshal-4                      |     780426 |   1549 ns/op |       110 |        232 |   1.21 |         8584 |    6.68
BenchmarkGobMarshal-4                         |    1802167 |    669 ns/op |        63 |         48 |   1.21 |        11443 |   13.94
BenchmarkGobUnmarshal-4                       |    1649358 |    720 ns/op |        63 |        112 |   1.19 |        10489 |    6.43
BenchmarkXDRMarshal-4                         |     777394 |   1481 ns/op |        88 |        392 |   1.15 |         6841 |    3.78
BenchmarkXDRUnmarshal-4                       |     940797 |   1302 ns/op |        88 |        224 |   1.22 |         8279 |    5.81
BenchmarkUgorjiCodecMsgpackMarshal-4          |    1221223 |    975 ns/op |        91 |       1312 |   1.19 |        11113 |    0.74
BenchmarkUgorjiCodecMsgpackUnmarshal-4        |    1000000 |   1058 ns/op |        91 |        496 |   1.06 |         9100 |    2.13
BenchmarkUgorjiCodecBincMarshal-4             |    1000000 |   1116 ns/op |        95 |       1328 |   1.12 |         9500 |    0.84
BenchmarkUgorjiCodecBincUnmarshal-4           |     910153 |   1309 ns/op |        95 |        656 |   1.19 |         8646 |    2.00
BenchmarkSerealMarshal-4                      |     534934 |   2328 ns/op |       132 |        904 |   1.25 |         7061 |    2.58
BenchmarkSerealUnmarshal-4                    |     456093 |   2640 ns/op |       132 |       1008 |   1.20 |         6020 |    2.62
BenchmarkBinaryMarshal-4                      |    1000000 |   1230 ns/op |        61 |        320 |   1.23 |         6100 |    3.84
BenchmarkBinaryUnmarshal-4                    |     978750 |   1246 ns/op |        61 |        320 |   1.22 |         5970 |    3.89
BenchmarkFlatBuffersMarshal-4                 |    4363474 |    270 ns/op |        95 |          0 |   1.18 |        41496 |    0.00
BenchmarkFlatBuffersUnmarshal-4               |    5363686 |    223 ns/op |        95 |        112 |   1.20 |        51115 |    1.99
BenchmarkCapNProtoMarshal-4                   |    3202656 |    370 ns/op |        96 |         56 |   1.18 |        30745 |    6.61
BenchmarkCapNProtoUnmarshal-4                 |    3164347 |    379 ns/op |        96 |        200 |   1.20 |        30377 |    1.90
BenchmarkCapNProto2Marshal-4                  |    2154375 |    563 ns/op |        96 |        244 |   1.21 |        20682 |    2.31
BenchmarkCapNProto2Unmarshal-4                |    1609051 |    712 ns/op |        96 |        320 |   1.15 |        15446 |    2.23
BenchmarkHproseMarshal-4                      |    1323446 |    953 ns/op |        82 |        421 |   1.26 |        10891 |    2.26
BenchmarkHproseUnmarshal-4                    |    1000000 |   1126 ns/op |        82 |        320 |   1.13 |         8230 |    3.52
BenchmarkHprose2Marshal-4                     |    2053401 |    588 ns/op |        82 |          0 |   1.21 |        16899 |    0.00
BenchmarkHprose2Unmarshal-4                   |    1913242 |    630 ns/op |        82 |        144 |   1.21 |        15745 |    4.38
BenchmarkProtobufMarshal-4                    |    1761278 |    674 ns/op |        52 |        152 |   1.19 |         9158 |    4.43
BenchmarkProtobufUnmarshal-4                  |    1916198 |    627 ns/op |        52 |        192 |   1.20 |         9964 |    3.27
BenchmarkGoprotobufMarshal-4                  |    4823341 |    249 ns/op |        53 |         64 |   1.20 |        25563 |    3.89
BenchmarkGoprotobufUnmarshal-4                |    3088417 |    394 ns/op |        53 |        168 |   1.22 |        16368 |    2.35
BenchmarkGogoprotobufMarshal-4                |    9089631 |    131 ns/op |        53 |         64 |   1.19 |        48175 |    2.05
BenchmarkGogoprotobufUnmarshal-4              |    6390816 |    190 ns/op |        53 |         96 |   1.21 |        33871 |    1.98
BenchmarkColferMarshal-4                      |   10938900 |    108 ns/op |        51 |         64 |   1.18 |        55897 |    1.69
BenchmarkColferUnmarshal-4                    |    7260112 |    166 ns/op |        52 |        112 |   1.21 |        37752 |    1.48
BenchmarkGencodeMarshal-4                     |    6778201 |    174 ns/op |        53 |         80 |   1.18 |        35924 |    2.17
BenchmarkGencodeUnmarshal-4                   |    5876618 |    204 ns/op |        53 |        112 |   1.20 |        31146 |    1.82
BenchmarkGencodeUnsafeMarshal-4               |   13984578 |     87 ns/op |        46 |         48 |   1.22 |        64329 |    1.82
BenchmarkGencodeUnsafeUnmarshal-4             |    9094292 |    135 ns/op |        46 |         96 |   1.23 |        41833 |    1.41
BenchmarkXDR2Marshal-4                        |    8195052 |    151 ns/op |        60 |         64 |   1.24 |        49170 |    2.36
BenchmarkXDR2Unmarshal-4                      |   10386304 |    116 ns/op |        60 |         32 |   1.20 |        62317 |    3.62
BenchmarkGoAvroMarshal-4                      |     475334 |   2619 ns/op |        47 |       1008 |   1.24 |         2234 |    2.60
BenchmarkGoAvroUnmarshal-4                    |     196928 |   6017 ns/op |        47 |       3328 |   1.18 |          925 |    1.81
BenchmarkGoAvro2TextMarshal-4                 |     464100 |   2460 ns/op |       134 |       1320 |   1.14 |         6218 |    1.86
BenchmarkGoAvro2TextUnmarshal-4               |     463238 |   2487 ns/op |       134 |        799 |   1.15 |         6207 |    3.11
BenchmarkGoAvro2BinaryMarshal-4               |    1495969 |    807 ns/op |        47 |        488 |   1.21 |         7031 |    1.65
BenchmarkGoAvro2BinaryUnmarshal-4             |    1384635 |    888 ns/op |        47 |        560 |   1.23 |         6507 |    1.59
BenchmarkIkeaMarshal-4                        |    2066786 |    579 ns/op |        55 |         72 |   1.20 |        11367 |    8.04
BenchmarkIkeaUnmarshal-4                      |    1718088 |    712 ns/op |        55 |        160 |   1.22 |         9449 |    4.45
BenchmarkShamatonMapMsgpackMarshal-4          |    1752916 |    690 ns/op |        92 |        208 |   1.21 |        16126 |    3.32
BenchmarkShamatonMapMsgpackUnmarshal-4        |    1856962 |    644 ns/op |        92 |        144 |   1.20 |        17084 |    4.47
BenchmarkShamatonArrayMsgpackMarshal-4        |    1921989 |    623 ns/op |        50 |        176 |   1.20 |         9609 |    3.54
BenchmarkShamatonArrayMsgpackUnmarshal-4      |    2760435 |    440 ns/op |        50 |        144 |   1.21 |        13802 |    3.06
BenchmarkShamatonMapMsgpackgenMarshal-4       |    6075926 |    192 ns/op |        92 |         96 |   1.17 |        55898 |    2.00
BenchmarkShamatonMapMsgpackgenUnmarshal-4     |    5926278 |    203 ns/op |        92 |         80 |   1.20 |        54521 |    2.54
BenchmarkShamatonArrayMsgpackgenMarshal-4     |    7993288 |    147 ns/op |        50 |         64 |   1.18 |        39966 |    2.30
BenchmarkShamatonArrayMsgpackgenUnmarshal-4   |    8656678 |    141 ns/op |        50 |         80 |   1.22 |        43283 |    1.76
BenchmarkSSZNoTimeNoStringNoFloatAMarshal-4   |     302756 |   4054 ns/op |        55 |        440 |   1.23 |         1665 |    9.21
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal-4 |     184963 |   6556 ns/op |        55 |       1392 |   1.21 |         1017 |    4.71
BenchmarkMumMarshal-4                         |   12559153 |     96 ns/op |        48 |          0 |   1.22 |        60283 |    0.00
BenchmarkMumUnmarshal-4                       |    6098353 |    190 ns/op |        48 |         80 |   1.16 |        29272 |    2.38
BenchmarkBebopMarshal-4                       |   10607622 |    114 ns/op |        55 |         64 |   1.21 |        58341 |    1.78
BenchmarkBebopUnmarshal-4                     |   12939981 |     94 ns/op |        55 |         32 |   1.23 |        71169 |    2.97


Totals:


benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
BenchmarkBebop-4                         |   23547603 |    208 ns/op |       110 |         96 |   4.92 |       259023 |    2.18
BenchmarkGencodeUnsafe-4                 |   23078870 |    222 ns/op |        92 |        144 |   5.13 |       212325 |    1.54
BenchmarkXDR2-4                          |   18581356 |    267 ns/op |       120 |         96 |   4.96 |       222976 |    2.78
BenchmarkColfer-4                        |   18199012 |    274 ns/op |       103 |        176 |   4.99 |       187631 |    1.56
BenchmarkMum-4                           |   18657506 |    286 ns/op |        96 |         80 |   5.35 |       179112 |    3.58
BenchmarkShamatonArrayMsgpackgen-4       |   16649966 |    288 ns/op |       100 |        144 |   4.80 |       166499 |    2.00
BenchmarkGotinyNoTime-4                  |   17292643 |    292 ns/op |        96 |         96 |   5.06 |       166009 |    3.05
BenchmarkGotiny-4                        |   16499283 |    319 ns/op |        96 |        112 |   5.26 |       158393 |    2.85
BenchmarkGogoprotobuf-4                  |   15480447 |    321 ns/op |       106 |        160 |   4.97 |       164092 |    2.01
BenchmarkGencode-4                       |   12654819 |    378 ns/op |       106 |        192 |   4.78 |       134141 |    1.97
BenchmarkShamatonMapMsgpackgen-4         |   12002204 |    395 ns/op |       184 |        176 |   4.74 |       220840 |    2.24
BenchmarkMsgp-4                          |   11878893 |    451 ns/op |       194 |        240 |   5.36 |       230450 |    1.88
BenchmarkFlatBuffers-4                   |    9727160 |    493 ns/op |       190 |        112 |   4.80 |       185205 |    4.40
BenchmarkGoprotobuf-4                    |    7911758 |    643 ns/op |       106 |        232 |   5.09 |        83864 |    2.77
BenchmarkCapNProto-4                     |    6367003 |    749 ns/op |       192 |        256 |   4.77 |       122246 |    2.93
BenchmarkShamatonArrayMsgpack-4          |    4682424 |   1063 ns/op |       100 |        320 |   4.98 |        46824 |    3.32
BenchmarkHprose2-4                       |    3966643 |   1218 ns/op |       164 |        144 |   4.83 |        65290 |    8.46
BenchmarkCapNProto2-4                    |    3763426 |   1275 ns/op |       192 |        564 |   4.80 |        72257 |    2.26
BenchmarkIkea-4                          |    3784874 |   1291 ns/op |       110 |        232 |   4.89 |        41633 |    5.56
BenchmarkProtobuf-4                      |    3677476 |   1301 ns/op |       104 |        344 |   4.78 |        38245 |    3.78
BenchmarkShamatonMapMsgpack-4            |    3609878 |   1334 ns/op |       184 |        352 |   4.82 |        66421 |    3.79
BenchmarkGob-4                           |    3451525 |   1389 ns/op |       127 |        160 |   4.79 |        43868 |    8.68
BenchmarkGoAvro2Binary-4                 |    2880604 |   1695 ns/op |        94 |       1048 |   4.88 |        27077 |    1.62
BenchmarkUgorjiCodecMsgpack-4            |    2221223 |   2033 ns/op |       182 |       1808 |   4.52 |        40426 |    1.12
BenchmarkHprose-4                        |    2323446 |   2079 ns/op |       164 |        741 |   4.83 |        38243 |    2.81
BenchmarkUgorjiCodecBinc-4               |    1910153 |   2425 ns/op |       190 |       1984 |   4.63 |        36292 |    1.22
BenchmarkBinary-4                        |    1978750 |   2476 ns/op |       122 |        640 |   4.90 |        24140 |    3.87
BenchmarkVmihailencoMsgpack-4            |    2137975 |   2582 ns/op |       200 |        816 |   5.52 |        42759 |    3.16
BenchmarkBson-4                          |    1780426 |   2595 ns/op |       220 |        624 |   4.62 |        39169 |    4.16
BenchmarkXDR-4                           |    1718191 |   2783 ns/op |       176 |        616 |   4.78 |        30240 |    4.52
BenchmarkEasyJson-4                      |    1579770 |   3100 ns/op |       298 |       1183 |   4.90 |        47077 |    2.62
BenchmarkJsonIter-4                      |    1455267 |   3411 ns/op |       276 |        511 |   4.96 |        40165 |    6.68
BenchmarkGoAvro2Text-4                   |     927338 |   4947 ns/op |       268 |       2119 |   4.59 |        24852 |    2.33
BenchmarkSereal-4                        |     991027 |   4968 ns/op |       264 |       1912 |   4.92 |        26163 |    2.60
BenchmarkJson-4                          |    1025954 |   5323 ns/op |       298 |        599 |   5.46 |        30573 |    8.89
BenchmarkGoAvro-4                        |     672262 |   8636 ns/op |        94 |       4336 |   5.81 |         6319 |    1.99
BenchmarkSSZNoTimeNoStringNoFloatA-4     |     487719 |  10610 ns/op |       110 |       1832 |   5.17 |         5364 |    5.79




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
