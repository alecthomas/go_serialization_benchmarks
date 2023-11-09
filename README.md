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

2023-11-09 Results with Go 1.20.8 linux/amd64 on an `AMD Ryzen 7 5700G` processor.

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-16                      |    5893668 |    202 ns/op |        48 |        168 |   1.19 |        28289 |    1.20
Benchmark_Gotiny_Unmarshal-16                    |   11002984 |    107 ns/op |        48 |        112 |   1.18 |        52814 |    0.96
Benchmark_GotinyNoTime_Marshal-16                |    5981194 |    200 ns/op |        48 |        168 |   1.20 |        28709 |    1.19
Benchmark_GotinyNoTime_Unmarshal-16              |   11647933 |    101 ns/op |        47 |         96 |   1.18 |        55898 |    1.05
Benchmark_Msgp_Marshal-16                        |   16109664 |     74 ns/op |        97 |        128 |   1.21 |       156263 |    0.59
Benchmark_Msgp_Unmarshal-16                      |    9277227 |    128 ns/op |        97 |        112 |   1.19 |        89989 |    1.15
Benchmark_VmihailencoMsgpack_Marshal-16          |    2942233 |    406 ns/op |        92 |        264 |   1.20 |        27068 |    1.54
Benchmark_VmihailencoMsgpack_Unmarshal-16        |    2204208 |    544 ns/op |        92 |        160 |   1.20 |        20278 |    3.40
Benchmark_Json_Marshal-16                        |    1771520 |    678 ns/op |       146 |        207 |   1.20 |        25988 |    3.28
Benchmark_Json_Unmarshal-16                      |     713296 |   1582 ns/op |       146 |        351 |   1.13 |        10464 |    4.51
Benchmark_JsonIter_Marshal-16                    |    2865139 |    421 ns/op |       136 |        200 |   1.21 |        39051 |    2.11
Benchmark_JsonIter_Unmarshal-16                  |    2241669 |    537 ns/op |       136 |        152 |   1.20 |        30553 |    3.53
Benchmark_EasyJson_Marshal-16                    |    2291847 |    523 ns/op |       146 |        895 |   1.20 |        33621 |    0.59
Benchmark_EasyJson_Unmarshal-16                  |    2524495 |    474 ns/op |       146 |        112 |   1.20 |        37034 |    4.23
Benchmark_Bson_Marshal-16                        |    2266069 |    525 ns/op |       110 |        376 |   1.19 |        24926 |    1.40
Benchmark_Bson_Unmarshal-16                      |    1485040 |    806 ns/op |       110 |        224 |   1.20 |        16335 |    3.60
Benchmark_MongoBson_Marshal-16                   |    1573221 |    763 ns/op |       110 |        240 |   1.20 |        17305 |    3.18
Benchmark_MongoBson_Unmarshal-16                 |    1425972 |    839 ns/op |       110 |        408 |   1.20 |        15685 |    2.06
Benchmark_Gob_Marshal-16                         |     471745 |   2455 ns/op |       163 |       1616 |   1.16 |         7717 |    1.52
Benchmark_Gob_Unmarshal-16                       |      96074 |  12607 ns/op |       163 |       7768 |   1.21 |         1571 |    1.62
Benchmark_XDR_Marshal-16                         |    1821038 |    653 ns/op |        87 |        376 |   1.19 |        16019 |    1.74
Benchmark_XDR_Unmarshal-16                       |    2168800 |    555 ns/op |        87 |        216 |   1.20 |        19076 |    2.57
Benchmark_UgorjiCodecMsgpack_Marshal-16          |    2454757 |    489 ns/op |        91 |       1240 |   1.20 |        22338 |    0.39
Benchmark_UgorjiCodecMsgpack_Unmarshal-16        |    2158896 |    554 ns/op |        91 |        688 |   1.20 |        19645 |    0.81
Benchmark_UgorjiCodecBinc_Marshal-16             |    2212545 |    541 ns/op |        95 |       1256 |   1.20 |        21019 |    0.43
Benchmark_UgorjiCodecBinc_Unmarshal-16           |    2398892 |    502 ns/op |        95 |        688 |   1.21 |        22789 |    0.73
Benchmark_Sereal_Marshal-16                      |     881689 |   1223 ns/op |       132 |        832 |   1.08 |        11638 |    1.47
Benchmark_Sereal_Unmarshal-16                    |     760290 |   1458 ns/op |       132 |        976 |   1.11 |        10035 |    1.49
Benchmark_Binary_Marshal-16                      |    1521852 |    789 ns/op |        61 |        360 |   1.20 |         9283 |    2.19
Benchmark_Binary_Unmarshal-16                    |    1814486 |    663 ns/op |        61 |        320 |   1.20 |        11068 |    2.07
Benchmark_FlatBuffers_Marshal-16                 |    3043201 |    393 ns/op |        95 |        376 |   1.20 |        28965 |    1.05
Benchmark_FlatBuffers_Unmarshal-16               |   10666461 |    112 ns/op |        95 |        112 |   1.20 |       101448 |    1.01
Benchmark_CapNProto_Marshal-16                   |    1428817 |    836 ns/op |        96 |       4392 |   1.20 |        13716 |    0.19
Benchmark_CapNProto_Unmarshal-16                 |    4897686 |    243 ns/op |        96 |        192 |   1.19 |        47017 |    1.27
Benchmark_CapNProto2_Marshal-16                  |    2003498 |    599 ns/op |        96 |       1452 |   1.20 |        19233 |    0.41
Benchmark_CapNProto2_Unmarshal-16                |    4822640 |    245 ns/op |        96 |        272 |   1.18 |        46297 |    0.90
Benchmark_Hprose_Marshal-16                      |    2800008 |    410 ns/op |        85 |        335 |   1.15 |        23878 |    1.23
Benchmark_Hprose_Unmarshal-16                    |    2336965 |    512 ns/op |        85 |        304 |   1.20 |        19922 |    1.69
Benchmark_Hprose2_Marshal-16                     |    5420389 |    221 ns/op |        85 |          0 |   1.20 |        46208 |    0.00
Benchmark_Hprose2_Unmarshal-16                   |    4422158 |    270 ns/op |        85 |        136 |   1.20 |        37725 |    1.99
Benchmark_Protobuf_Marshal-16                    |    3635408 |    329 ns/op |        52 |        144 |   1.20 |        18904 |    2.29
Benchmark_Protobuf_Unmarshal-16                  |    2963258 |    405 ns/op |        52 |        184 |   1.20 |        15408 |    2.20
Benchmark_Pulsar_Marshal-16                      |    4111294 |    291 ns/op |        51 |        304 |   1.20 |        21177 |    0.96
Benchmark_Pulsar_Unmarshal-16                    |    4258818 |    282 ns/op |        51 |        256 |   1.20 |        21966 |    1.10
Benchmark_Gogoprotobuf_Marshal-16                |   18167898 |     64 ns/op |        53 |         64 |   1.18 |        96289 |    1.01
Benchmark_Gogoprotobuf_Unmarshal-16              |   11376072 |    105 ns/op |        53 |         96 |   1.20 |        60293 |    1.09
Benchmark_Gogojsonpb_Marshal-16                  |     208504 |   5608 ns/op |       125 |       3103 |   1.17 |         2622 |    1.81
Benchmark_Gogojsonpb_Unmarshal-16                |     159320 |   7437 ns/op |       125 |       3371 |   1.18 |         1996 |    2.21
Benchmark_Colfer_Marshal-16                      |   21178388 |     56 ns/op |        51 |         64 |   1.19 |       108200 |    0.88
Benchmark_Colfer_Unmarshal-16                    |   13176285 |     89 ns/op |        51 |        112 |   1.17 |        67199 |    0.80
Benchmark_Gencode_Marshal-16                     |   16241857 |     72 ns/op |        53 |         80 |   1.19 |        86081 |    0.91
Benchmark_Gencode_Unmarshal-16                   |   13085042 |     89 ns/op |        53 |        112 |   1.18 |        69350 |    0.80
Benchmark_GencodeUnsafe_Marshal-16               |   26336108 |     44 ns/op |        46 |         48 |   1.16 |       121146 |    0.92
Benchmark_GencodeUnsafe_Unmarshal-16             |   15414268 |     76 ns/op |        46 |         96 |   1.17 |        70905 |    0.79
Benchmark_XDR2_Marshal-16                        |   13755733 |     85 ns/op |        60 |         64 |   1.18 |        82534 |    1.33
Benchmark_XDR2_Unmarshal-16                      |   20688716 |     56 ns/op |        60 |         32 |   1.17 |       124132 |    1.77
Benchmark_GoAvro_Marshal-16                      |     998076 |   1097 ns/op |        47 |        728 |   1.09 |         4690 |    1.51
Benchmark_GoAvro_Unmarshal-16                    |     434211 |   2682 ns/op |        47 |       2544 |   1.16 |         2040 |    1.05
Benchmark_GoAvro2Text_Marshal-16                 |     729447 |   1469 ns/op |       133 |       1320 |   1.07 |         9760 |    1.11
Benchmark_GoAvro2Text_Unmarshal-16               |     842248 |   1357 ns/op |       133 |        736 |   1.14 |        11269 |    1.84
Benchmark_GoAvro2Binary_Marshal-16               |    2727043 |    435 ns/op |        47 |        464 |   1.19 |        12817 |    0.94
Benchmark_GoAvro2Binary_Unmarshal-16             |    2510193 |    478 ns/op |        47 |        544 |   1.20 |        11797 |    0.88
Benchmark_Ikea_Marshal-16                        |    2711467 |    441 ns/op |        55 |        184 |   1.20 |        14913 |    2.40
Benchmark_Ikea_Unmarshal-16                      |    3208425 |    372 ns/op |        55 |        160 |   1.20 |        17646 |    2.33
Benchmark_ShamatonMapMsgpack_Marshal-16          |    3776151 |    317 ns/op |        92 |        192 |   1.20 |        34740 |    1.65
Benchmark_ShamatonMapMsgpack_Unmarshal-16        |    3857624 |    308 ns/op |        92 |        168 |   1.19 |        35490 |    1.84
Benchmark_ShamatonArrayMsgpack_Marshal-16        |    4236583 |    281 ns/op |        50 |        160 |   1.19 |        21182 |    1.76
Benchmark_ShamatonArrayMsgpack_Unmarshal-16      |    4675030 |    254 ns/op |        50 |        168 |   1.19 |        23375 |    1.52
Benchmark_ShamatonMapMsgpackgen_Marshal-16       |   12103406 |     98 ns/op |        92 |         96 |   1.19 |       111351 |    1.02
Benchmark_ShamatonMapMsgpackgen_Unmarshal-16     |    7575436 |    158 ns/op |        92 |        112 |   1.20 |        69694 |    1.41
Benchmark_ShamatonArrayMsgpackgen_Marshal-16     |   15862881 |     74 ns/op |        50 |         64 |   1.18 |        79314 |    1.16
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-16   |   11567385 |    102 ns/op |        50 |        112 |   1.19 |        57836 |    0.92
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-16   |     526494 |   2151 ns/op |        55 |        440 |   1.13 |         2895 |    4.89
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-16 |     344326 |   3414 ns/op |        55 |       1184 |   1.18 |         1893 |    2.88
Benchmark_Bebop_Marshal-16                       |   19384552 |     61 ns/op |        55 |         64 |   1.20 |       106615 |    0.96
Benchmark_Bebop_Unmarshal-16                     |   25699130 |     46 ns/op |        55 |         32 |   1.20 |       141345 |    1.45
Benchmark_Bebop_Wellquite_Marshal-16             |   22055068 |     54 ns/op |        55 |         64 |   1.20 |       121302 |    0.85
Benchmark_Bebop_Wellquite_Unmarshal-16           |   22203843 |     53 ns/op |        55 |         32 |   1.18 |       122121 |    1.66
Benchmark_FastJson_Marshal-16                    |    3591696 |    330 ns/op |       133 |        504 |   1.19 |        48020 |    0.66
Benchmark_FastJson_Unmarshal-16                  |    1559835 |    773 ns/op |       133 |       1704 |   1.21 |        20870 |    0.45
Benchmark_MUS_Marshal-16                         |   23202069 |     49 ns/op |        46 |         48 |   1.16 |       106729 |    1.04
Benchmark_MUS_Unmarshal-16                       |   17132107 |     69 ns/op |        46 |         32 |   1.19 |        78807 |    2.17
Benchmark_MUSUnsafe_Marshal-16                   |   24492463 |     47 ns/op |        49 |         64 |   1.17 |       120013 |    0.75
Benchmark_MUSUnsafe_Unmarshal-16                 |   45169773 |     26 ns/op |        49 |          0 |   1.20 |       221331 |    0.00


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MUSUnsafe_-16                 |   69662236 |     74 ns/op |        98 |         64 |   5.18 |       682689 |    1.16
Benchmark_Bebop_Wellquite_-16           |   44258911 |    107 ns/op |       110 |         96 |   4.75 |       486848 |    1.12
Benchmark_Bebop_-16                     |   45083682 |    108 ns/op |       110 |         96 |   4.88 |       495920 |    1.13
Benchmark_MUS_-16                       |   40334176 |    119 ns/op |        92 |         80 |   4.81 |       371074 |    1.49
Benchmark_GencodeUnsafe_-16             |   41750376 |    120 ns/op |        92 |        144 |   5.02 |       384103 |    0.84
Benchmark_XDR2_-16                      |   34444449 |    142 ns/op |       120 |         96 |   4.90 |       413333 |    1.48
Benchmark_Colfer_-16                    |   34354673 |    145 ns/op |       102 |        176 |   4.99 |       350726 |    0.83
Benchmark_Gencode_-16                   |   29326899 |    162 ns/op |       106 |        192 |   4.78 |       310865 |    0.85
Benchmark_Gogoprotobuf_-16              |   29543970 |    169 ns/op |       106 |        160 |   5.02 |       313166 |    1.06
Benchmark_ShamatonArrayMsgpackgen_-16   |   27430266 |    177 ns/op |       100 |        176 |   4.86 |       274302 |    1.01
Benchmark_Msgp_-16                      |   25386891 |    203 ns/op |       194 |        240 |   5.17 |       492505 |    0.85
Benchmark_ShamatonMapMsgpackgen_-16     |   19678842 |    256 ns/op |       184 |        208 |   5.04 |       362090 |    1.23
Benchmark_GotinyNoTime_-16              |   17629127 |    301 ns/op |        95 |        264 |   5.32 |       169221 |    1.14
Benchmark_Gotiny_-16                    |   16896652 |    309 ns/op |        96 |        280 |   5.23 |       162207 |    1.11
Benchmark_Hprose2_-16                   |    9842547 |    491 ns/op |       170 |        136 |   4.84 |       167874 |    3.61
Benchmark_FlatBuffers_-16               |   13709662 |    506 ns/op |       190 |        488 |   6.94 |       260881 |    1.04
Benchmark_ShamatonArrayMsgpack_-16      |    8911613 |    536 ns/op |       100 |        328 |   4.78 |        89116 |    1.64
Benchmark_Pulsar_-16                    |    8370112 |    573 ns/op |       103 |        560 |   4.80 |        86287 |    1.02
Benchmark_ShamatonMapMsgpack_-16        |    7633775 |    625 ns/op |       184 |        360 |   4.78 |       140461 |    1.74
Benchmark_Protobuf_-16                  |    6598666 |    735 ns/op |       104 |        328 |   4.85 |        68626 |    2.24
Benchmark_Ikea_-16                      |    5919892 |    814 ns/op |       110 |        344 |   4.82 |        65118 |    2.37
Benchmark_CapNProto2_-16                |    6826138 |    844 ns/op |       192 |       1724 |   5.77 |       131061 |    0.49
Benchmark_GoAvro2Binary_-16             |    5237236 |    913 ns/op |        94 |       1008 |   4.79 |        49230 |    0.91
Benchmark_Hprose_-16                    |    5136973 |    923 ns/op |       170 |        639 |   4.74 |        87600 |    1.44
Benchmark_VmihailencoMsgpack_-16        |    5146441 |    951 ns/op |       184 |        424 |   4.90 |        94694 |    2.24
Benchmark_JsonIter_-16                  |    5106808 |    959 ns/op |       272 |        352 |   4.90 |       139211 |    2.72
Benchmark_EasyJson_-16                  |    4816342 |    998 ns/op |       293 |       1007 |   4.81 |       141311 |    0.99
Benchmark_UgorjiCodecMsgpack_-16        |    4613653 |   1044 ns/op |       182 |       1928 |   4.82 |        83968 |    0.54
Benchmark_UgorjiCodecBinc_-16           |    4611437 |   1044 ns/op |       190 |       1944 |   4.82 |        87617 |    0.54
Benchmark_CapNProto_-16                 |    6326503 |   1080 ns/op |       192 |       4584 |   6.84 |       121468 |    0.24
Benchmark_FastJson_-16                  |    5151531 |   1103 ns/op |       267 |       2208 |   5.69 |       137803 |    0.50
Benchmark_XDR_-16                       |    3989838 |   1208 ns/op |       175 |        592 |   4.82 |        70193 |    2.04
Benchmark_Bson_-16                      |    3751109 |   1332 ns/op |       220 |        600 |   5.00 |        82524 |    2.22
Benchmark_Binary_-16                    |    3336338 |   1452 ns/op |       122 |        680 |   4.85 |        40703 |    2.14
Benchmark_MongoBson_-16                 |    2999193 |   1602 ns/op |       220 |        648 |   4.81 |        65982 |    2.47
Benchmark_Json_-16                      |    2484816 |   2260 ns/op |       293 |        558 |   5.62 |        72904 |    4.05
Benchmark_Sereal_-16                    |    1641979 |   2681 ns/op |       264 |       1808 |   4.40 |        43348 |    1.48
Benchmark_GoAvro2Text_-16               |    1571695 |   2826 ns/op |       267 |       2056 |   4.44 |        42058 |    1.37
Benchmark_GoAvro_-16                    |    1432287 |   3779 ns/op |        94 |       3272 |   5.41 |        13463 |    1.15
Benchmark_SSZNoTimeNoStringNoFloatA_-16 |     870820 |   5565 ns/op |       110 |       1624 |   4.85 |         9579 |    3.43
Benchmark_Gogojsonpb_-16                |     367824 |  13045 ns/op |       251 |       6474 |   4.80 |         9236 |    2.01
Benchmark_Gob_-16                       |     567819 |  15062 ns/op |       327 |       9384 |   8.55 |        18579 |    1.61

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
