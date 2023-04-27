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
- [google.golang.org/protobuf](https://google.golang.org/protobuf) (generated code)
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
- [github.com/mus-format/mus-go](https://github.com/mus-format/mus-go)

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

2023-04-27 Results with Go 1.20.3 linux/amd64 on an `AMD Ryzen 7 PRO 5850U` processor.

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-16                      |    2509231 |    653 ns/op |        48 |        168 |   1.64 |        12044 |    3.89
Benchmark_Gotiny_Unmarshal-16                    |    3494467 |    357 ns/op |        47 |        112 |   1.25 |        16769 |    3.19
Benchmark_GotinyNoTime_Marshal-16                |    1665592 |    715 ns/op |        48 |        168 |   1.19 |         7994 |    4.26
Benchmark_GotinyNoTime_Unmarshal-16              |    3763646 |    345 ns/op |        48 |         96 |   1.30 |        18065 |    3.60
Benchmark_Msgp_Marshal-16                        |    4259458 |    279 ns/op |        97 |        128 |   1.19 |        41316 |    2.19
Benchmark_Msgp_Unmarshal-16                      |    3301718 |    429 ns/op |        97 |        112 |   1.42 |        32026 |    3.84
Benchmark_VmihailencoMsgpack_Marshal-16          |     870376 |   1231 ns/op |        92 |        264 |   1.07 |         8007 |    4.66
Benchmark_VmihailencoMsgpack_Unmarshal-16        |     739293 |   1470 ns/op |        92 |        160 |   1.09 |         6801 |    9.19
Benchmark_Json_Marshal-16                        |    1069760 |   1537 ns/op |       151 |        208 |   1.64 |        16217 |    7.39
Benchmark_Json_Unmarshal-16                      |     257702 |   3933 ns/op |       151 |        351 |   1.01 |         3906 |   11.21
Benchmark_JsonIter_Marshal-16                    |    1000000 |   1015 ns/op |       141 |        200 |   1.01 |        14130 |    5.08
Benchmark_JsonIter_Unmarshal-16                  |     897181 |   1292 ns/op |       141 |        216 |   1.16 |        12677 |    5.98
Benchmark_EasyJson_Marshal-16                    |    1000000 |   1343 ns/op |       151 |        896 |   1.34 |        15170 |    1.50
Benchmark_EasyJson_Unmarshal-16                  |    1041213 |   1298 ns/op |       151 |        112 |   1.35 |        15795 |   11.59
Benchmark_Bson_Marshal-16                        |     722877 |   1491 ns/op |       110 |        376 |   1.08 |         7951 |    3.97
Benchmark_Bson_Unmarshal-16                      |     748554 |   1855 ns/op |       110 |        224 |   1.39 |         8234 |    8.28
Benchmark_MongoBson_Marshal-16                   |     764475 |   1499 ns/op |       110 |        240 |   1.15 |         8409 |    6.25
Benchmark_MongoBson_Unmarshal-16                 |     483150 |   2166 ns/op |       110 |        408 |   1.05 |         5314 |    5.31
Benchmark_Gob_Marshal-16                         |     201495 |   6363 ns/op |       163 |       1616 |   1.28 |         3298 |    3.94
Benchmark_Gob_Unmarshal-16                       |      37329 |  33507 ns/op |       163 |       7768 |   1.25 |          610 |    4.31
Benchmark_XDR_Marshal-16                         |     593994 |   1987 ns/op |        92 |        439 |   1.18 |         5464 |    4.53
Benchmark_XDR_Unmarshal-16                       |     769069 |   1573 ns/op |        92 |        232 |   1.21 |         7075 |    6.78
Benchmark_UgorjiCodecMsgpack_Marshal-16          |    1000000 |   1538 ns/op |        91 |       1240 |   1.54 |         9100 |    1.24
Benchmark_UgorjiCodecMsgpack_Unmarshal-16        |     529971 |   2164 ns/op |        91 |        688 |   1.15 |         4822 |    3.15
Benchmark_UgorjiCodecBinc_Marshal-16             |     555745 |   2007 ns/op |        95 |       1256 |   1.12 |         5279 |    1.60
Benchmark_UgorjiCodecBinc_Unmarshal-16           |     569472 |   2001 ns/op |        95 |        688 |   1.14 |         5409 |    2.91
Benchmark_Sereal_Marshal-16                      |     295790 |   4727 ns/op |       132 |        832 |   1.40 |         3904 |    5.68
Benchmark_Sereal_Unmarshal-16                    |     201170 |   5099 ns/op |       132 |        976 |   1.03 |         2655 |    5.22
Benchmark_Binary_Marshal-16                      |     502834 |   2571 ns/op |        61 |        360 |   1.29 |         3067 |    7.14
Benchmark_Binary_Unmarshal-16                    |     401703 |   2501 ns/op |        61 |        320 |   1.00 |         2450 |    7.82
Benchmark_FlatBuffers_Marshal-16                 |    1000000 |   1317 ns/op |        95 |        376 |   1.32 |         9511 |    3.50
Benchmark_FlatBuffers_Unmarshal-16               |    2784686 |    401 ns/op |        95 |        112 |   1.12 |        26512 |    3.59
Benchmark_CapNProto_Marshal-16                   |     430296 |   3012 ns/op |        96 |       4392 |   1.30 |         4130 |    0.69
Benchmark_CapNProto_Unmarshal-16                 |    1210761 |    896 ns/op |        96 |        192 |   1.09 |        11623 |    4.67
Benchmark_CapNProto2_Marshal-16                  |     480901 |   2187 ns/op |        96 |       1452 |   1.05 |         4616 |    1.51
Benchmark_CapNProto2_Unmarshal-16                |    1365840 |    770 ns/op |        96 |        272 |   1.05 |        13112 |    2.83
Benchmark_Hprose_Marshal-16                      |    2070277 |    522 ns/op |        85 |        403 |   1.08 |        17655 |    1.30
Benchmark_Hprose_Unmarshal-16                    |     748009 |   1667 ns/op |        85 |        303 |   1.25 |         6378 |    5.50
Benchmark_Hprose2_Marshal-16                     |    4282826 |    253 ns/op |        85 |          0 |   1.09 |        36519 |    0.00
Benchmark_Hprose2_Unmarshal-16                   |    1670266 |    752 ns/op |        85 |        136 |   1.26 |        14240 |    5.53
Benchmark_Protobuf_Marshal-16                    |    1000000 |   1025 ns/op |        52 |        144 |   1.02 |         5200 |    7.12
Benchmark_Protobuf_Unmarshal-16                  |    1189854 |   1439 ns/op |        52 |        184 |   1.71 |         6187 |    7.82
Benchmark_Pulsar_Marshal-16                      |     974084 |   1166 ns/op |        51 |        304 |   1.14 |         5033 |    3.84
Benchmark_Pulsar_Unmarshal-16                    |    1195524 |   1136 ns/op |        51 |        256 |   1.36 |         6166 |    4.44
Benchmark_Gogoprotobuf_Marshal-16                |    5046446 |    239 ns/op |        53 |         64 |   1.21 |        26746 |    3.74
Benchmark_Gogoprotobuf_Unmarshal-16              |    3007797 |    352 ns/op |        53 |         96 |   1.06 |        15941 |    3.67
Benchmark_Gogojsonpb_Marshal-16                  |      71272 |  15293 ns/op |       126 |       3121 |   1.09 |          899 |    4.90
Benchmark_Gogojsonpb_Unmarshal-16                |      54970 |  22405 ns/op |       125 |       3378 |   1.23 |          691 |    6.63
Benchmark_Colfer_Marshal-16                      |    9342464 |    157 ns/op |        51 |         64 |   1.47 |        47739 |    2.45
Benchmark_Colfer_Unmarshal-16                    |    3477394 |    292 ns/op |        49 |        112 |   1.02 |        17039 |    2.61
Benchmark_Gencode_Marshal-16                     |    6034485 |    204 ns/op |        53 |         80 |   1.23 |        31982 |    2.55
Benchmark_Gencode_Unmarshal-16                   |    5354996 |    322 ns/op |        53 |        112 |   1.73 |        28381 |    2.88
Benchmark_GencodeUnsafe_Marshal-16               |    6939962 |    165 ns/op |        46 |         48 |   1.15 |        31923 |    3.45
Benchmark_GencodeUnsafe_Unmarshal-16             |    5896057 |    223 ns/op |        46 |         96 |   1.32 |        27121 |    2.32
Benchmark_XDR2_Marshal-16                        |    4457760 |    259 ns/op |        60 |         64 |   1.15 |        26746 |    4.05
Benchmark_XDR2_Unmarshal-16                      |    5318547 |    206 ns/op |        60 |         32 |   1.10 |        31911 |    6.46
Benchmark_GoAvro_Marshal-16                      |     265651 |   4044 ns/op |        47 |        728 |   1.07 |         1248 |    5.55
Benchmark_GoAvro_Unmarshal-16                    |     131302 |   9988 ns/op |        47 |       2544 |   1.31 |          617 |    3.93
Benchmark_GoAvro2Text_Marshal-16                 |     198501 |   5475 ns/op |       133 |       1320 |   1.09 |         2653 |    4.15
Benchmark_GoAvro2Text_Unmarshal-16               |     285451 |   4918 ns/op |       133 |        736 |   1.40 |         3819 |    6.68
Benchmark_GoAvro2Binary_Marshal-16               |     775006 |   1516 ns/op |        47 |        464 |   1.17 |         3642 |    3.27
Benchmark_GoAvro2Binary_Unmarshal-16             |     946983 |   1169 ns/op |        47 |        544 |   1.11 |         4450 |    2.15
Benchmark_Ikea_Marshal-16                        |    1000000 |   1125 ns/op |        55 |        184 |   1.12 |         5500 |    6.11
Benchmark_Ikea_Unmarshal-16                      |    1438875 |    933 ns/op |        55 |        160 |   1.34 |         7913 |    5.83
Benchmark_ShamatonMapMsgpack_Marshal-16          |    1482009 |    828 ns/op |        92 |        192 |   1.23 |        13634 |    4.32
Benchmark_ShamatonMapMsgpack_Unmarshal-16        |    1399692 |    812 ns/op |        92 |        168 |   1.14 |        12877 |    4.84
Benchmark_ShamatonArrayMsgpack_Marshal-16        |    1529468 |    719 ns/op |        50 |        160 |   1.10 |         7647 |    4.50
Benchmark_ShamatonArrayMsgpack_Unmarshal-16      |    1764112 |    675 ns/op |        50 |        168 |   1.19 |         8820 |    4.02
Benchmark_ShamatonMapMsgpackgen_Marshal-16       |    4684315 |    245 ns/op |        92 |         96 |   1.15 |        43095 |    2.56
Benchmark_ShamatonMapMsgpackgen_Unmarshal-16     |    2481304 |    486 ns/op |        92 |        112 |   1.21 |        22827 |    4.34
Benchmark_ShamatonArrayMsgpackgen_Marshal-16     |    5872029 |    179 ns/op |        50 |         64 |   1.06 |        29360 |    2.81
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-16   |    4249507 |    265 ns/op |        50 |        112 |   1.13 |        21247 |    2.37
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-16   |     390163 |   4420 ns/op |        55 |        440 |   1.72 |         2145 |   10.05
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-16 |     141002 |   7252 ns/op |        55 |       1184 |   1.02 |          775 |    6.12
Benchmark_Bebop_Marshal-16                       |    7579790 |    153 ns/op |        55 |         64 |   1.16 |        41688 |    2.40
Benchmark_Bebop_Unmarshal-16                     |    9232676 |    127 ns/op |        55 |         32 |   1.18 |        50779 |    4.00
Benchmark_FastJson_Marshal-16                    |    1000000 |   1216 ns/op |       133 |        504 |   1.22 |        13369 |    2.41
Benchmark_FastJson_Unmarshal-16                  |     734396 |   1536 ns/op |       133 |       1704 |   1.13 |         9818 |    0.90
Benchmark_MUS_Marshal-16                         |    8805204 |    125 ns/op |        46 |         48 |   1.10 |        40503 |    2.61
Benchmark_MUS_Unmarshal-16                       |    7109659 |    181 ns/op |        46 |         32 |   1.29 |        32704 |    5.68
Benchmark_MUSUnsafe_Marshal-16                   |   12780822 |    111 ns/op |        49 |         64 |   1.43 |        62626 |    1.74
Benchmark_MUSUnsafe_Unmarshal-16                 |   37978525 |     30 ns/op |        49 |          0 |   1.16 |       186094 |    0.00


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MUSUnsafe_-16                 |   50759347 |    142 ns/op |        98 |         64 |   7.21 |       497441 |    2.22
Benchmark_Bebop_-16                     |   16812466 |    281 ns/op |       110 |         96 |   4.73 |       184937 |    2.93
Benchmark_MUS_-16                       |   15914863 |    307 ns/op |        92 |         80 |   4.89 |       146416 |    3.84
Benchmark_GencodeUnsafe_-16             |   12836019 |    388 ns/op |        92 |        144 |   4.99 |       118091 |    2.70
Benchmark_ShamatonArrayMsgpackgen_-16   |   10121536 |    445 ns/op |       100 |        176 |   4.51 |       101215 |    2.53
Benchmark_Colfer_-16                    |   12819858 |    449 ns/op |       100 |        176 |   5.77 |       128326 |    2.56
Benchmark_XDR2_-16                      |    9776307 |    465 ns/op |       120 |         96 |   4.55 |       117315 |    4.85
Benchmark_Gencode_-16                   |   11389481 |    527 ns/op |       106 |        192 |   6.00 |       120728 |    2.75
Benchmark_Gogoprotobuf_-16              |    8054243 |    592 ns/op |       106 |        160 |   4.77 |        85374 |    3.70
Benchmark_Msgp_-16                      |    7561176 |    709 ns/op |       194 |        240 |   5.36 |       146686 |    2.96
Benchmark_ShamatonMapMsgpackgen_-16     |    7165619 |    731 ns/op |       184 |        208 |   5.24 |       131847 |    3.52
Benchmark_Hprose2_-16                   |    5953092 |   1006 ns/op |       170 |        136 |   5.99 |       101518 |    7.40
Benchmark_Gotiny_-16                    |    6003698 |   1010 ns/op |        95 |        280 |   6.07 |        57629 |    3.61
Benchmark_GotinyNoTime_-16              |    5429238 |   1060 ns/op |        96 |        264 |   5.76 |        52120 |    4.02
Benchmark_ShamatonArrayMsgpack_-16      |    3293580 |   1395 ns/op |       100 |        328 |   4.59 |        32935 |    4.25
Benchmark_ShamatonMapMsgpack_-16        |    2881701 |   1641 ns/op |       184 |        360 |   4.73 |        53023 |    4.56
Benchmark_FlatBuffers_-16               |    3784686 |   1718 ns/op |       190 |        488 |   6.50 |        72030 |    3.52
Benchmark_Ikea_-16                      |    2438875 |   2058 ns/op |       110 |        344 |   5.02 |        26827 |    5.98
Benchmark_Hprose_-16                    |    2818286 |   2189 ns/op |       170 |        706 |   6.17 |        48065 |    3.10
Benchmark_Pulsar_-16                    |    2169608 |   2302 ns/op |       103 |        560 |   4.99 |        22401 |    4.11
Benchmark_JsonIter_-16                  |    1897181 |   2307 ns/op |       282 |        416 |   4.38 |        53614 |    5.55
Benchmark_Protobuf_-16                  |    2189854 |   2464 ns/op |       104 |        328 |   5.40 |        22774 |    7.51
Benchmark_EasyJson_-16                  |    2041213 |   2641 ns/op |       303 |       1008 |   5.39 |        61930 |    2.62
Benchmark_GoAvro2Binary_-16             |    1721989 |   2685 ns/op |        94 |       1008 |   4.62 |        16186 |    2.66
Benchmark_VmihailencoMsgpack_-16        |    1609669 |   2701 ns/op |       184 |        424 |   4.35 |        29617 |    6.37
Benchmark_FastJson_-16                  |    1734396 |   2752 ns/op |       267 |       2208 |   4.77 |        46377 |    1.25
Benchmark_CapNProto2_-16                |    1846741 |   2957 ns/op |       192 |       1724 |   5.46 |        35457 |    1.72
Benchmark_Bson_-16                      |    1471431 |   3346 ns/op |       220 |        600 |   4.92 |        32371 |    5.58
Benchmark_XDR_-16                       |    1363063 |   3560 ns/op |       184 |        671 |   4.85 |        25080 |    5.31
Benchmark_MongoBson_-16                 |    1247625 |   3665 ns/op |       220 |        648 |   4.57 |        27447 |    5.66
Benchmark_UgorjiCodecMsgpack_-16        |    1529971 |   3702 ns/op |       182 |       1928 |   5.66 |        27845 |    1.92
Benchmark_CapNProto_-16                 |    1641057 |   3908 ns/op |       192 |       4584 |   6.41 |        31508 |    0.85
Benchmark_UgorjiCodecBinc_-16           |    1125217 |   4008 ns/op |       190 |       1944 |   4.51 |        21379 |    2.06
Benchmark_Binary_-16                    |     904537 |   5072 ns/op |       122 |        680 |   4.59 |        11035 |    7.46
Benchmark_Json_-16                      |    1327462 |   5470 ns/op |       303 |        559 |   7.26 |        40248 |    9.79
Benchmark_Sereal_-16                    |     496960 |   9826 ns/op |       264 |       1808 |   4.88 |        13119 |    5.43
Benchmark_GoAvro2Text_-16               |     483952 |  10393 ns/op |       267 |       2056 |   5.03 |        12945 |    5.05
Benchmark_SSZNoTimeNoStringNoFloatA_-16 |     531165 |  11672 ns/op |       110 |       1624 |   6.20 |         5842 |    7.19
Benchmark_GoAvro_-16                    |     396953 |  14032 ns/op |        94 |       3272 |   5.57 |         3731 |    4.29
Benchmark_Gogojsonpb_-16                |     126242 |  37698 ns/op |       252 |       6499 |   4.76 |         3181 |    5.80
Benchmark_Gob_-16                       |     238824 |  39870 ns/op |       327 |       9384 |   9.52 |         7816 |    4.25

## Issues


The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
2. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

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

3. **(major)** Goprotobuf has been disabled in the above, as it is no longer maintained and is incompatible with the latest changes to Google's Protobuf package. See discussions: [1](https://github.com/containerd/ttrpc/issues/62), [2](https://github.com/containerd/ttrpc/pull/99).

```
panic: protobuf tag not enough fields in ProtoBufA.state:

goroutine 148 [running]:
github.com/gogo/protobuf/proto.(*unmarshalInfo).computeUnmarshalInfo(0xc0000d9c20)
	/home/ken/src/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:341 +0x135f
github.com/gogo/protobuf/proto.(*unmarshalInfo).unmarshal(0xc0000d9c20, {0xcff820?}, {0xc000122fc0, 0x35, 0x35})
	/home/ken/src/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:138 +0x67
github.com/gogo/protobuf/proto.(*InternalMessageInfo).Unmarshal(0xc00013a8c0?, {0xe6d5f0, 0xc000114060}, {0xc000122fc0?, 0x35?, 0x35?})
	/home/ken/src/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:63 +0xd0
github.com/gogo/protobuf/proto.(*Buffer).Unmarshal(0xc00009fe28, {0xe6d5f0, 0xc000114060})
	/home/ken/src/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:424 +0x153
github.com/gogo/protobuf/proto.Unmarshal({0xc000122fc0, 0x35, 0x35}, {0xe6d5f0, 0xc000114060})
	/home/ken/src/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:342 +0xe6
github.com/alecthomas/go_serialization_benchmarks.Benchmark_Goprotobuf_Unmarshal(0xc0000d5680)
	/home/ken/src/go/src/github.com/alecthomas/go_serialization_benchmarks/serialization_benchmarks_test.go:853 +0x26a
testing.(*B).runN(0xc0000d5680, 0x1)
	/usr/local/go/src/testing/benchmark.go:193 +0x102
testing.(*B).run1.func1()
	/usr/local/go/src/testing/benchmark.go:233 +0x59
created by testing.(*B).run1
	/usr/local/go/src/testing/benchmark.go:226 +0x9c
exit status 2
FAIL	github.com/alecthomas/go_serialization_benchmarks	61.159s
```
