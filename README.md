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
- [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo) (generated code)

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

2023-03-19 Results with Go 1.20.2 darwin/arm64 on an `Apple M1 Pro` processor.

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                       |    6910878 |    164 ns/op |        48 |        168 |   1.14 |        33172 |    0.98
Benchmark_Gotiny_Unmarshal-8                     |   13162723 |     90 ns/op |        48 |        112 |   1.19 |        63181 |    0.81
Benchmark_GotinyNoTime_Marshal-8                 |    7306672 |    163 ns/op |        48 |        168 |   1.20 |        35072 |    0.97
Benchmark_GotinyNoTime_Unmarshal-8               |   14211855 |     84 ns/op |        48 |         96 |   1.20 |        68216 |    0.88
Benchmark_Msgp_Marshal-8                         |   18400051 |     66 ns/op |        97 |        128 |   1.23 |       178480 |    0.52
Benchmark_Msgp_Unmarshal-8                       |   11166733 |    107 ns/op |        97 |        112 |   1.20 |       108317 |    0.96
Benchmark_VmihailencoMsgpack_Marshal-8           |    3738108 |    322 ns/op |        92 |        264 |   1.20 |        34390 |    1.22
Benchmark_VmihailencoMsgpack_Unmarshal-8         |    2877516 |    416 ns/op |        92 |        160 |   1.20 |        26473 |    2.60
Benchmark_Json_Marshal-8                         |    2333619 |    513 ns/op |       148 |        208 |   1.20 |        34700 |    2.47
Benchmark_Json_Unmarshal-8                       |     876066 |   1339 ns/op |       148 |        351 |   1.17 |        13027 |    3.81
Benchmark_JsonIter_Marshal-8                     |    3414224 |    350 ns/op |       138 |        200 |   1.20 |        47218 |    1.75
Benchmark_JsonIter_Unmarshal-8                   |    2518658 |    476 ns/op |       138 |        215 |   1.20 |        34833 |    2.21
Benchmark_EasyJson_Marshal-8                     |    2640439 |    451 ns/op |       148 |        896 |   1.19 |        39263 |    0.50
Benchmark_EasyJson_Unmarshal-8                   |    2858536 |    418 ns/op |       148 |        112 |   1.20 |        42506 |    3.74
Benchmark_Bson_Marshal-8                         |    2860502 |    418 ns/op |       110 |        376 |   1.20 |        31465 |    1.11
Benchmark_Bson_Unmarshal-8                       |    1930910 |    623 ns/op |       110 |        224 |   1.20 |        21240 |    2.78
Benchmark_MongoBson_Marshal-8                    |    2019808 |    594 ns/op |       110 |        240 |   1.20 |        22217 |    2.48
Benchmark_MongoBson_Unmarshal-8                  |    1813014 |    660 ns/op |       110 |        408 |   1.20 |        19943 |    1.62
Benchmark_Gob_Marshal-8                          |     554912 |   2129 ns/op |       163 |       1616 |   1.18 |         9078 |    1.32
Benchmark_Gob_Unmarshal-8                        |     114760 |  10493 ns/op |       163 |       7768 |   1.20 |         1877 |    1.35
Benchmark_XDR_Marshal-8                          |    2199034 |    545 ns/op |        88 |        376 |   1.20 |        19351 |    1.45
Benchmark_XDR_Unmarshal-8                        |    2682135 |    447 ns/op |        88 |        216 |   1.20 |        23602 |    2.07
Benchmark_UgorjiCodecMsgpack_Marshal-8           |    2700848 |    438 ns/op |        91 |       1240 |   1.18 |        24577 |    0.35
Benchmark_UgorjiCodecMsgpack_Unmarshal-8         |    2670764 |    451 ns/op |        91 |        688 |   1.20 |        24303 |    0.66
Benchmark_UgorjiCodecBinc_Marshal-8              |    2528095 |    474 ns/op |        95 |       1256 |   1.20 |        24016 |    0.38
Benchmark_UgorjiCodecBinc_Unmarshal-8            |    2868908 |    419 ns/op |        95 |        688 |   1.20 |        27254 |    0.61
Benchmark_Sereal_Marshal-8                       |    1000000 |   1024 ns/op |       132 |        832 |   1.02 |        13200 |    1.23
Benchmark_Sereal_Unmarshal-8                     |     969272 |   1234 ns/op |       132 |        976 |   1.20 |        12794 |    1.26
Benchmark_Binary_Marshal-8                       |    1726345 |    694 ns/op |        61 |        360 |   1.20 |        10530 |    1.93
Benchmark_Binary_Unmarshal-8                     |    2007010 |    598 ns/op |        61 |        320 |   1.20 |        12242 |    1.87
Benchmark_FlatBuffers_Marshal-8                  |    3699286 |    324 ns/op |        95 |        376 |   1.20 |        35150 |    0.86
Benchmark_FlatBuffers_Unmarshal-8                |   11826961 |    100 ns/op |        95 |        112 |   1.19 |       112639 |    0.90
Benchmark_CapNProto_Marshal-8                    |    1537599 |    746 ns/op |        96 |       4392 |   1.15 |        14760 |    0.17
Benchmark_CapNProto_Unmarshal-8                  |    5760372 |    207 ns/op |        96 |        192 |   1.20 |        55299 |    1.08
Benchmark_CapNProto2_Marshal-8                   |    2257982 |    534 ns/op |        96 |       1452 |   1.21 |        21676 |    0.37
Benchmark_CapNProto2_Unmarshal-8                 |    5674141 |    211 ns/op |        96 |        272 |   1.20 |        54471 |    0.78
Benchmark_Hprose_Marshal-8                       |    3417669 |    353 ns/op |        82 |        450 |   1.21 |        28127 |    0.79
Benchmark_Hprose_Unmarshal-8                     |    2459268 |    485 ns/op |        82 |        304 |   1.19 |        20220 |    1.60
Benchmark_Hprose2_Marshal-8                      |    5984535 |    200 ns/op |        82 |          0 |   1.20 |        49216 |    0.00
Benchmark_Hprose2_Unmarshal-8                    |    4719901 |    254 ns/op |        82 |        136 |   1.20 |        38830 |    1.87
Benchmark_Protobuf_Marshal-8                     |    4109466 |    291 ns/op |        52 |        144 |   1.20 |        21369 |    2.03
Benchmark_Protobuf_Unmarshal-8                   |    3563036 |    333 ns/op |        52 |        184 |   1.19 |        18527 |    1.81
Benchmark_Pulsar_Marshal-8                       |    7813304 |    153 ns/op |        51 |         96 |   1.20 |        40316 |    1.60
Benchmark_Pulsar_Unmarshal-8                     |    5099353 |    237 ns/op |        51 |        256 |   1.21 |        26307 |    0.93
Benchmark_Gogoprotobuf_Marshal-8                 |   23081472 |     51 ns/op |        53 |         64 |   1.20 |       122331 |    0.81
Benchmark_Gogoprotobuf_Unmarshal-8               |   13722277 |     87 ns/op |        53 |         96 |   1.20 |        72728 |    0.91
Benchmark_Gogojsonpb_Marshal-8                   |     261174 |   4574 ns/op |       125 |       3103 |   1.19 |         3288 |    1.47
Benchmark_Gogojsonpb_Unmarshal-8                 |     195382 |   6093 ns/op |       125 |       3379 |   1.19 |         2457 |    1.80
Benchmark_Colfer_Marshal-8                       |   25008987 |     47 ns/op |        51 |         64 |   1.19 |       127770 |    0.74
Benchmark_Colfer_Unmarshal-8                     |   15736904 |     75 ns/op |        51 |        112 |   1.18 |        80258 |    0.67
Benchmark_Gencode_Marshal-8                      |   16971225 |     69 ns/op |        53 |         80 |   1.18 |        89947 |    0.87
Benchmark_Gencode_Unmarshal-8                    |   14103494 |     84 ns/op |        53 |        112 |   1.20 |        74748 |    0.76
Benchmark_GencodeUnsafe_Marshal-8                |   32042224 |     37 ns/op |        46 |         48 |   1.20 |       147394 |    0.78
Benchmark_GencodeUnsafe_Unmarshal-8              |   19686882 |     60 ns/op |        46 |         96 |   1.20 |        90559 |    0.63
Benchmark_XDR2_Marshal-8                         |   17436548 |     68 ns/op |        60 |         64 |   1.19 |       104619 |    1.07
Benchmark_XDR2_Unmarshal-8                       |   25500716 |     46 ns/op |        60 |         32 |   1.20 |       153004 |    1.47
Benchmark_GoAvro_Marshal-8                       |    1208142 |    998 ns/op |        47 |        728 |   1.21 |         5678 |    1.37
Benchmark_GoAvro_Unmarshal-8                     |     489298 |   2435 ns/op |        47 |       2544 |   1.19 |         2299 |    0.96
Benchmark_GoAvro2Text_Marshal-8                  |     922174 |   1279 ns/op |       133 |       1320 |   1.18 |        12338 |    0.97
Benchmark_GoAvro2Text_Unmarshal-8                |    1000000 |   1156 ns/op |       133 |        736 |   1.16 |        13380 |    1.57
Benchmark_GoAvro2Binary_Marshal-8                |    3268442 |    367 ns/op |        47 |        464 |   1.20 |        15361 |    0.79
Benchmark_GoAvro2Binary_Unmarshal-8              |    3011289 |    398 ns/op |        47 |        544 |   1.20 |        14153 |    0.73
Benchmark_Ikea_Marshal-8                         |    2916994 |    411 ns/op |        55 |        184 |   1.20 |        16043 |    2.24
Benchmark_Ikea_Unmarshal-8                       |    3703088 |    323 ns/op |        55 |        160 |   1.20 |        20366 |    2.02
Benchmark_ShamatonMapMsgpack_Marshal-8           |    4550641 |    263 ns/op |        92 |        192 |   1.20 |        41865 |    1.37
Benchmark_ShamatonMapMsgpack_Unmarshal-8         |    4591184 |    261 ns/op |        92 |        168 |   1.20 |        42238 |    1.56
Benchmark_ShamatonArrayMsgpack_Marshal-8         |    5046403 |    236 ns/op |        50 |        160 |   1.19 |        25232 |    1.48
Benchmark_ShamatonArrayMsgpack_Unmarshal-8       |    5755782 |    206 ns/op |        50 |        168 |   1.19 |        28778 |    1.23
Benchmark_ShamatonMapMsgpackgen_Marshal-8        |   15044930 |     79 ns/op |        92 |         96 |   1.19 |       138413 |    0.83
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8      |    7666126 |    157 ns/op |        92 |        112 |   1.20 |        70528 |    1.40
Benchmark_ShamatonArrayMsgpackgen_Marshal-8      |   18860732 |     63 ns/op |        50 |         64 |   1.20 |        94303 |    0.99
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8    |   14339086 |     84 ns/op |        50 |        112 |   1.20 |        71695 |    0.75
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8    |     597661 |   1944 ns/op |        55 |        440 |   1.16 |         3287 |    4.42
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8  |     389184 |   3084 ns/op |        55 |       1184 |   1.20 |         2140 |    2.60
Benchmark_Bebop_Marshal-8                        |   23227604 |     51 ns/op |        55 |         64 |   1.19 |       127751 |    0.80
Benchmark_Bebop_Unmarshal-8                      |   30521892 |     38 ns/op |        55 |         32 |   1.18 |       167870 |    1.21
Benchmark_FastJson_Marshal-8                     |    4377237 |    276 ns/op |       133 |        504 |   1.21 |        58567 |    0.55
Benchmark_FastJson_Unmarshal-8                   |    1757622 |    686 ns/op |       133 |       1704 |   1.21 |        23516 |    0.40
Benchmark_Musgo_Marshal-8                        |   26404720 |     44 ns/op |        46 |         48 |   1.17 |       121461 |    0.92
Benchmark_Musgo_Unmarshal-8                      |   17563813 |     67 ns/op |        46 |         96 |   1.19 |        80793 |    0.71
Benchmark_MusgoUnsafe_Marshal-8                  |   29941893 |     40 ns/op |        46 |         48 |   1.20 |       137732 |    0.83
Benchmark_MusgoUnsafe_Unmarshal-8                |   26668123 |     44 ns/op |        46 |         64 |   1.19 |       122673 |    0.70


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MusgoUnsafe_-8                |   56610016 |     84 ns/op |        92 |        112 |   4.80 |       520812 |    0.76
Benchmark_Bebop_-8                      |   53749496 |     89 ns/op |       110 |         96 |   4.82 |       591244 |    0.93
Benchmark_GencodeUnsafe_-8              |   51729106 |     98 ns/op |        92 |        144 |   5.09 |       475907 |    0.68
Benchmark_Musgo_-8                      |   43968533 |    112 ns/op |        92 |        144 |   4.93 |       404510 |    0.78
Benchmark_XDR2_-8                       |   42937264 |    115 ns/op |       120 |         96 |   4.95 |       515247 |    1.20
Benchmark_Colfer_-8                     |   40745891 |    122 ns/op |       102 |        176 |   5.00 |       415974 |    0.70
Benchmark_Gogoprotobuf_-8               |   36803749 |    138 ns/op |       106 |        160 |   5.11 |       390119 |    0.87
Benchmark_ShamatonArrayMsgpackgen_-8    |   33199818 |    147 ns/op |       100 |        176 |   4.90 |       331998 |    0.84
Benchmark_Gencode_-8                    |   31074719 |    154 ns/op |       106 |        192 |   4.80 |       329392 |    0.81
Benchmark_Msgp_-8                       |   29566784 |    173 ns/op |       194 |        240 |   5.14 |       573595 |    0.72
Benchmark_ShamatonMapMsgpackgen_-8      |   22711056 |    236 ns/op |       184 |        208 |   5.37 |       417883 |    1.14
Benchmark_GotinyNoTime_-8               |   21518527 |    247 ns/op |        96 |        264 |   5.33 |       206577 |    0.94
Benchmark_Gotiny_-8                     |   20073601 |    255 ns/op |        96 |        280 |   5.13 |       192706 |    0.91
Benchmark_Pulsar_-8                     |   12912657 |    391 ns/op |       103 |        352 |   5.05 |       133245 |    1.11
Benchmark_FlatBuffers_-8                |   15526247 |    425 ns/op |       190 |        488 |   6.60 |       295402 |    0.87
Benchmark_ShamatonArrayMsgpack_-8       |   10802185 |    443 ns/op |       100 |        328 |   4.79 |       108021 |    1.35
Benchmark_Hprose2_-8                    |   10704436 |    454 ns/op |       164 |        136 |   4.87 |       176098 |    3.34
Benchmark_ShamatonMapMsgpack_-8         |    9141825 |    525 ns/op |       184 |        360 |   4.80 |       168209 |    1.46
Benchmark_Protobuf_-8                   |    7672502 |    625 ns/op |       104 |        328 |   4.80 |        79794 |    1.91
Benchmark_Ikea_-8                       |    6620082 |    734 ns/op |       110 |        344 |   4.86 |        72820 |    2.14
Benchmark_VmihailencoMsgpack_-8         |    6615624 |    738 ns/op |       184 |        424 |   4.89 |       121727 |    1.74
Benchmark_CapNProto2_-8                 |    7932123 |    745 ns/op |       192 |       1724 |   5.91 |       152296 |    0.43
Benchmark_GoAvro2Binary_-8              |    6279731 |    766 ns/op |        94 |       1008 |   4.81 |        59029 |    0.76
Benchmark_JsonIter_-8                   |    5932882 |    826 ns/op |       276 |        415 |   4.91 |       164103 |    1.99
Benchmark_Hprose_-8                     |    5876937 |    839 ns/op |       164 |        754 |   4.93 |        96687 |    1.11
Benchmark_EasyJson_-8                   |    5498975 |    870 ns/op |       297 |       1008 |   4.79 |       163539 |    0.86
Benchmark_UgorjiCodecMsgpack_-8         |    5371612 |    889 ns/op |       182 |       1928 |   4.78 |        97763 |    0.46
Benchmark_UgorjiCodecBinc_-8            |    5397003 |    894 ns/op |       190 |       1944 |   4.83 |       102543 |    0.46
Benchmark_CapNProto_-8                  |    7297971 |    954 ns/op |       192 |       4584 |   6.96 |       140121 |    0.21
Benchmark_FastJson_-8                   |    6134859 |    962 ns/op |       267 |       2208 |   5.91 |       164168 |    0.44
Benchmark_XDR_-8                        |    4881169 |    992 ns/op |       176 |        592 |   4.85 |        85908 |    1.68
Benchmark_Bson_-8                       |    4791412 |   1042 ns/op |       220 |        600 |   5.00 |       105411 |    1.74
Benchmark_MongoBson_-8                  |    3832822 |   1254 ns/op |       220 |        648 |   4.81 |        84322 |    1.94
Benchmark_Binary_-8                     |    3733355 |   1293 ns/op |       122 |        680 |   4.83 |        45546 |    1.90
Benchmark_Json_-8                       |    3209685 |   1852 ns/op |       297 |        559 |   5.95 |        95456 |    3.31
Benchmark_Sereal_-8                     |    1969272 |   2258 ns/op |       264 |       1808 |   4.45 |        51988 |    1.25
Benchmark_GoAvro2Text_-8                |    1922174 |   2435 ns/op |       267 |       2056 |   4.68 |        51437 |    1.18
Benchmark_GoAvro_-8                     |    1697440 |   3433 ns/op |        94 |       3272 |   5.83 |        15955 |    1.05
Benchmark_SSZNoTimeNoStringNoFloatA_-8  |     986845 |   5028 ns/op |       110 |       1624 |   4.96 |        10855 |    3.10
Benchmark_Gogojsonpb_-8                 |     456556 |  10667 ns/op |       251 |       6482 |   4.87 |        11491 |    1.65
Benchmark_Gob_-8                        |     669672 |  12622 ns/op |       327 |       9384 |   8.45 |        21911 |    1.35

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
