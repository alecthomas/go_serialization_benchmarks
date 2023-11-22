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
- [wellquite.org/bebop](https://fossil.wellquite.org/bebop) (generated code)

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

2023-11-09 Results with Go 1.21.4 darwin/arm64 on an `Apple M1 Max` processor.

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-10                      |    6962595 |    168 ns/op |        48 |        168 |   1.17 |        33420 |    1.00
Benchmark_Gotiny_Unmarshal-10                    |   13797159 |     88 ns/op |        48 |        112 |   1.22 |        66226 |    0.79
Benchmark_GotinyNoTime_Marshal-10                |    7157514 |    169 ns/op |        48 |        168 |   1.22 |        34356 |    1.01
Benchmark_GotinyNoTime_Unmarshal-10              |   14069850 |     84 ns/op |        47 |         96 |   1.19 |        67521 |    0.88
Benchmark_Msgp_Marshal-10                        |   19623474 |     61 ns/op |        97 |        128 |   1.20 |       190347 |    0.48
Benchmark_Msgp_Unmarshal-10                      |   11416347 |    106 ns/op |        97 |        112 |   1.22 |       110738 |    0.95
Benchmark_VmihailencoMsgpack_Marshal-10          |    3626785 |    330 ns/op |        92 |        264 |   1.20 |        33366 |    1.25
Benchmark_VmihailencoMsgpack_Unmarshal-10        |    2779572 |    422 ns/op |        92 |        160 |   1.17 |        25572 |    2.64
Benchmark_Json_Marshal-10                        |    2446729 |    486 ns/op |       143 |        194 |   1.19 |        35135 |    2.51
Benchmark_Json_Unmarshal-10                      |     841398 |   1336 ns/op |       143 |        351 |   1.12 |        12082 |    3.81
Benchmark_JsonIter_Marshal-10                    |    3517827 |    339 ns/op |       133 |        200 |   1.19 |        46892 |    1.70
Benchmark_JsonIter_Unmarshal-10                  |    2728002 |    432 ns/op |       133 |        152 |   1.18 |        36364 |    2.84
Benchmark_EasyJson_Marshal-10                    |    2795341 |    425 ns/op |       143 |        882 |   1.19 |        40141 |    0.48
Benchmark_EasyJson_Unmarshal-10                  |    2882943 |    396 ns/op |       143 |        112 |   1.14 |        41427 |    3.54
Benchmark_Bson_Marshal-10                        |    2903034 |    412 ns/op |       110 |        376 |   1.20 |        31933 |    1.10
Benchmark_Bson_Unmarshal-10                      |    1940450 |    618 ns/op |       110 |        224 |   1.20 |        21344 |    2.76
Benchmark_MongoBson_Marshal-10                   |    2002221 |    588 ns/op |       110 |        240 |   1.18 |        22024 |    2.45
Benchmark_MongoBson_Unmarshal-10                 |    1832785 |    670 ns/op |       110 |        408 |   1.23 |        20160 |    1.64
Benchmark_Gob_Marshal-10                         |     551840 |   2161 ns/op |       162 |       1744 |   1.19 |         8972 |    1.24
Benchmark_Gob_Unmarshal-10                       |     123770 |   9786 ns/op |       162 |       7720 |   1.21 |         2012 |    1.27
Benchmark_XDR_Marshal-10                         |    2269971 |    525 ns/op |        84 |        376 |   1.19 |        19067 |    1.40
Benchmark_XDR_Unmarshal-10                       |    2818861 |    424 ns/op |        84 |        216 |   1.20 |        23678 |    1.97
Benchmark_UgorjiCodecMsgpack_Marshal-10          |    2644554 |    468 ns/op |        91 |       1240 |   1.24 |        24065 |    0.38
Benchmark_UgorjiCodecMsgpack_Unmarshal-10        |    2664181 |    447 ns/op |        91 |        688 |   1.19 |        24244 |    0.65
Benchmark_UgorjiCodecBinc_Marshal-10             |    2405847 |    501 ns/op |        95 |       1256 |   1.21 |        22855 |    0.40
Benchmark_UgorjiCodecBinc_Unmarshal-10           |    2938502 |    409 ns/op |        95 |        688 |   1.20 |        27915 |    0.59
Benchmark_Sereal_Marshal-10                      |    1000000 |   1012 ns/op |       132 |        832 |   1.01 |        13200 |    1.22
Benchmark_Sereal_Unmarshal-10                    |     977356 |   1210 ns/op |       132 |        976 |   1.18 |        12901 |    1.24
Benchmark_Binary_Marshal-10                      |    1763468 |    683 ns/op |        61 |        360 |   1.21 |        10757 |    1.90
Benchmark_Binary_Unmarshal-10                    |    2047929 |    576 ns/op |        61 |        320 |   1.18 |        12492 |    1.80
Benchmark_FlatBuffers_Marshal-10                 |    3836520 |    313 ns/op |        95 |        376 |   1.20 |        36523 |    0.83
Benchmark_FlatBuffers_Unmarshal-10               |   12118488 |     98 ns/op |        95 |        112 |   1.19 |       115428 |    0.88
Benchmark_CapNProto_Marshal-10                   |    1582750 |    695 ns/op |        96 |       4392 |   1.10 |        15194 |    0.16
Benchmark_CapNProto_Unmarshal-10                 |    5813928 |    206 ns/op |        96 |        192 |   1.20 |        55813 |    1.08
Benchmark_CapNProto2_Marshal-10                  |    2330149 |    520 ns/op |        96 |       1452 |   1.21 |        22369 |    0.36
Benchmark_CapNProto2_Unmarshal-10                |    5733844 |    210 ns/op |        96 |        272 |   1.21 |        55044 |    0.77
Benchmark_Hprose_Marshal-10                      |    3745573 |    319 ns/op |        82 |        422 |   1.20 |        30803 |    0.76
Benchmark_Hprose_Unmarshal-10                    |    2646633 |    445 ns/op |        82 |        304 |   1.18 |        21773 |    1.47
Benchmark_Hprose2_Marshal-10                     |    6750758 |    177 ns/op |        82 |          0 |   1.20 |        55578 |    0.00
Benchmark_Hprose2_Unmarshal-10                   |    5225089 |    228 ns/op |        82 |        136 |   1.19 |        42986 |    1.68
Benchmark_Protobuf_Marshal-10                    |    4145830 |    289 ns/op |        52 |        144 |   1.20 |        21558 |    2.01
Benchmark_Protobuf_Unmarshal-10                  |    3617934 |    330 ns/op |        52 |        184 |   1.19 |        18813 |    1.79
Benchmark_Pulsar_Marshal-10                      |    5215272 |    233 ns/op |        51 |        304 |   1.22 |        26916 |    0.77
Benchmark_Pulsar_Unmarshal-10                    |    5162821 |    234 ns/op |        51 |        256 |   1.21 |        26640 |    0.91
Benchmark_Gogoprotobuf_Marshal-10                |   24577173 |     48 ns/op |        53 |         64 |   1.20 |       130259 |    0.76
Benchmark_Gogoprotobuf_Unmarshal-10              |   14057605 |     85 ns/op |        53 |         96 |   1.20 |        74505 |    0.89
Benchmark_Gogojsonpb_Marshal-10                  |     264712 |   4516 ns/op |       125 |       3095 |   1.20 |         3324 |    1.46
Benchmark_Gogojsonpb_Unmarshal-10                |     199486 |   6014 ns/op |       125 |       3379 |   1.20 |         2509 |    1.78
Benchmark_Colfer_Marshal-10                      |   26230510 |     45 ns/op |        51 |         64 |   1.18 |       134037 |    0.70
Benchmark_Colfer_Unmarshal-10                    |   16755222 |     72 ns/op |        51 |        112 |   1.21 |        85451 |    0.64
Benchmark_Gencode_Marshal-10                     |   19474090 |     59 ns/op |        53 |         80 |   1.16 |       103212 |    0.75
Benchmark_Gencode_Unmarshal-10                   |   17293776 |     69 ns/op |        53 |        112 |   1.20 |        91657 |    0.62
Benchmark_GencodeUnsafe_Marshal-10               |   33785923 |     35 ns/op |        46 |         48 |   1.21 |       155415 |    0.75
Benchmark_GencodeUnsafe_Unmarshal-10             |   20288104 |     59 ns/op |        46 |         96 |   1.20 |        93325 |    0.62
Benchmark_XDR2_Marshal-10                        |   18158763 |     66 ns/op |        60 |         64 |   1.20 |       108952 |    1.03
Benchmark_XDR2_Unmarshal-10                      |   26721618 |     45 ns/op |        60 |         32 |   1.20 |       160329 |    1.41
Benchmark_GoAvro_Marshal-10                      |    1418862 |    842 ns/op |        47 |        584 |   1.19 |         6668 |    1.44
Benchmark_GoAvro_Unmarshal-10                    |     579901 |   2111 ns/op |        47 |       2312 |   1.22 |         2725 |    0.91
Benchmark_GoAvro2Text_Marshal-10                 |     963158 |   1254 ns/op |       133 |       1320 |   1.21 |        12887 |    0.95
Benchmark_GoAvro2Text_Unmarshal-10               |    1000000 |   1145 ns/op |       133 |        736 |   1.15 |        13369 |    1.56
Benchmark_GoAvro2Binary_Marshal-10               |    3323559 |    361 ns/op |        47 |        464 |   1.20 |        15620 |    0.78
Benchmark_GoAvro2Binary_Unmarshal-10             |    3095734 |    388 ns/op |        47 |        544 |   1.20 |        14549 |    0.71
Benchmark_Ikea_Marshal-10                        |    2889531 |    414 ns/op |        55 |        184 |   1.20 |        15892 |    2.25
Benchmark_Ikea_Unmarshal-10                      |    3885115 |    309 ns/op |        55 |        160 |   1.20 |        21368 |    1.93
Benchmark_ShamatonMapMsgpack_Marshal-10          |    4632297 |    262 ns/op |        92 |        192 |   1.21 |        42617 |    1.36
Benchmark_ShamatonMapMsgpack_Unmarshal-10        |    4736762 |    252 ns/op |        92 |        168 |   1.20 |        43578 |    1.50
Benchmark_ShamatonArrayMsgpack_Marshal-10        |    5141290 |    233 ns/op |        50 |        160 |   1.20 |        25706 |    1.46
Benchmark_ShamatonArrayMsgpack_Unmarshal-10      |    5920248 |    202 ns/op |        50 |        168 |   1.20 |        29601 |    1.20
Benchmark_ShamatonMapMsgpackgen_Marshal-10       |   16191991 |     74 ns/op |        92 |         96 |   1.20 |       148966 |    0.77
Benchmark_ShamatonMapMsgpackgen_Unmarshal-10     |    8227104 |    145 ns/op |        92 |        112 |   1.20 |        75689 |    1.30
Benchmark_ShamatonArrayMsgpackgen_Marshal-10     |   20387590 |     58 ns/op |        50 |         64 |   1.20 |       101937 |    0.92
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-10   |   14803118 |     80 ns/op |        50 |        112 |   1.19 |        74015 |    0.72
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-10   |     593461 |   2002 ns/op |        55 |        440 |   1.19 |         3264 |    4.55
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-10 |     388472 |   3102 ns/op |        55 |       1184 |   1.21 |         2136 |    2.62
Benchmark_Bebop_200sc_Marshal-10                 |   27845754 |     42 ns/op |        55 |         64 |   1.19 |       153151 |    0.67
Benchmark_Bebop_200sc_Unmarshal-10               |   30457110 |     39 ns/op |        55 |         32 |   1.21 |       167514 |    1.24
Benchmark_Bebop_Wellquite_Marshal-10             |   27176578 |     43 ns/op |        55 |         64 |   1.19 |       149471 |    0.68
Benchmark_Bebop_Wellquite_Unmarshal-10           |   27232033 |     44 ns/op |        55 |         32 |   1.20 |       149776 |    1.38
Benchmark_FastJson_Marshal-10                    |    4538410 |    264 ns/op |       133 |        504 |   1.20 |        60723 |    0.52
Benchmark_FastJson_Unmarshal-10                  |    1820494 |    662 ns/op |       133 |       1704 |   1.21 |        24358 |    0.39
Benchmark_MUS_Marshal-10                         |   29512822 |     40 ns/op |        46 |         48 |   1.20 |       135758 |    0.85
Benchmark_MUS_Unmarshal-10                       |   23131359 |     52 ns/op |        46 |         32 |   1.21 |       106404 |    1.63
Benchmark_MUSUnsafe_Marshal-10                   |   31768863 |     37 ns/op |        49 |         64 |   1.20 |       155667 |    0.59
Benchmark_MUSUnsafe_Unmarshal-10                 |   53609522 |     22 ns/op |        49 |          0 |   1.22 |       262686 |    0.00


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MUSUnsafe_-10                 |   85378385 |     60 ns/op |        98 |         64 |   5.17 |       836708 |    0.95
Benchmark_Bebop_200sc_-10               |   58302864 |     82 ns/op |       110 |         96 |   4.79 |       641331 |    0.86
Benchmark_Bebop_Wellquite_-10           |   54408611 |     87 ns/op |       110 |         96 |   4.77 |       598494 |    0.91
Benchmark_MUS_-10                       |   52644181 |     92 ns/op |        92 |         80 |   4.90 |       484326 |    1.16
Benchmark_GencodeUnsafe_-10             |   54074027 |     94 ns/op |        92 |        144 |   5.13 |       497481 |    0.66
Benchmark_XDR2_-10                      |   44880381 |    111 ns/op |       120 |         96 |   4.99 |       538564 |    1.16
Benchmark_Colfer_-10                    |   42985732 |    117 ns/op |       102 |        176 |   5.04 |       438884 |    0.67
Benchmark_Gencode_-10                   |   36767866 |    129 ns/op |       106 |        192 |   4.75 |       389739 |    0.67
Benchmark_Gogoprotobuf_-10              |   38634778 |    134 ns/op |       106 |        160 |   5.19 |       409528 |    0.84
Benchmark_ShamatonArrayMsgpackgen_-10   |   35190708 |    139 ns/op |       100 |        176 |   4.91 |       351907 |    0.79
Benchmark_Msgp_-10                      |   31039821 |    168 ns/op |       194 |        240 |   5.22 |       602172 |    0.70
Benchmark_ShamatonMapMsgpackgen_-10     |   24419095 |    219 ns/op |       184 |        208 |   5.37 |       449311 |    1.06
Benchmark_GotinyNoTime_-10              |   21227364 |    254 ns/op |        95 |        264 |   5.39 |       203761 |    0.96
Benchmark_Gotiny_-10                    |   20759754 |    256 ns/op |        96 |        280 |   5.32 |       199293 |    0.92
Benchmark_Hprose2_-10                   |   11975847 |    406 ns/op |       164 |        136 |   4.86 |       197122 |    2.99
Benchmark_FlatBuffers_-10               |   15955008 |    412 ns/op |       190 |        488 |   6.58 |       303863 |    0.84
Benchmark_ShamatonArrayMsgpack_-10      |   11061538 |    435 ns/op |       100 |        328 |   4.82 |       110615 |    1.33
Benchmark_Pulsar_-10                    |   10378093 |    467 ns/op |       103 |        560 |   4.85 |       107112 |    0.83
Benchmark_ShamatonMapMsgpack_-10        |    9369059 |    514 ns/op |       184 |        360 |   4.82 |       172390 |    1.43
Benchmark_Protobuf_-10                  |    7763764 |    619 ns/op |       104 |        328 |   4.81 |        80743 |    1.89
Benchmark_Ikea_-10                      |    6774646 |    723 ns/op |       110 |        344 |   4.90 |        74521 |    2.10
Benchmark_CapNProto2_-10                |    8063993 |    731 ns/op |       192 |       1724 |   5.90 |       154828 |    0.42
Benchmark_GoAvro2Binary_-10             |    6419293 |    749 ns/op |        94 |       1008 |   4.81 |        60341 |    0.74
Benchmark_VmihailencoMsgpack_-10        |    6406357 |    752 ns/op |       184 |        424 |   4.82 |       117876 |    1.77
Benchmark_Hprose_-10                    |    6392206 |    764 ns/op |       164 |        726 |   4.89 |       105158 |    1.05
Benchmark_JsonIter_-10                  |    6245829 |    771 ns/op |       266 |        352 |   4.82 |       166513 |    2.19
Benchmark_EasyJson_-10                  |    5678284 |    821 ns/op |       287 |        994 |   4.66 |       163137 |    0.83
Benchmark_CapNProto_-10                 |    7396678 |    902 ns/op |       192 |       4584 |   6.68 |       142016 |    0.20
Benchmark_UgorjiCodecBinc_-10           |    5344349 |    910 ns/op |       190 |       1944 |   4.87 |       101542 |    0.47
Benchmark_UgorjiCodecMsgpack_-10        |    5308735 |    916 ns/op |       182 |       1928 |   4.87 |        96618 |    0.48
Benchmark_FastJson_-10                  |    6358904 |    926 ns/op |       267 |       2208 |   5.89 |       170164 |    0.42
Benchmark_XDR_-10                       |    5088832 |    950 ns/op |       168 |        592 |   4.84 |        85492 |    1.61
Benchmark_Bson_-10                      |    4843484 |   1030 ns/op |       220 |        600 |   4.99 |       106556 |    1.72
Benchmark_MongoBson_-10                 |    3835006 |   1258 ns/op |       220 |        648 |   4.83 |        84370 |    1.94
Benchmark_Binary_-10                    |    3811397 |   1260 ns/op |       122 |        680 |   4.80 |        46499 |    1.85
Benchmark_Json_-10                      |    3288127 |   1822 ns/op |       287 |        545 |   5.99 |        94435 |    3.34
Benchmark_Sereal_-10                    |    1977356 |   2222 ns/op |       264 |       1808 |   4.39 |        52202 |    1.23
Benchmark_GoAvro2Text_-10               |    1963158 |   2399 ns/op |       267 |       2056 |   4.71 |        52514 |    1.17
Benchmark_GoAvro_-10                    |    1998763 |   2953 ns/op |        94 |       2896 |   5.90 |        18788 |    1.02
Benchmark_SSZNoTimeNoStringNoFloatA_-10 |     981933 |   5104 ns/op |       110 |       1624 |   5.01 |        10801 |    3.14
Benchmark_Gogojsonpb_-10                |     464198 |  10530 ns/op |       251 |       6474 |   4.89 |        11669 |    1.63
Benchmark_Gob_-10                       |     675610 |  11947 ns/op |       325 |       9464 |   8.07 |        21970 |    1.26

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

Additionally, while not a correctness issue, FlatBuffers, ProtoBuffers, Cap'N'Proto, ikeapack and MUS
do not support time types directly. In the benchmarks an int64 value is used to hold a UnixNano timestamp.

Bebop (both libraries, by nature of the format) natively supports times rounded to 100ns ticks, and this is what is currently benchmarked (a unix nano timestamp is another valid approach).

MUSUnsafe results were obtained using the mus-go unsafe package. With this 
package, after decoding a byte slice into a string, any change to this slice 
will change the contents of the string. In such cases, the slice can be reused 
only after processing the received result.

1. **(major)** Goprotobuf has been disabled in the above, as it is no longer maintained and is incompatible with the latest changes to Google's Protobuf package. See discussions: [1](https://github.com/containerd/ttrpc/issues/62), [2](https://github.com/containerd/ttrpc/pull/99).

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
