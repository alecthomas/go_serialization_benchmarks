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

2021-06-21 Results with Go 1.17.8 Darwin/arm64 on an Apple M1 Max

benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-10                   |   18392354 |     64 ns/op |        48 |          0 |   1.18 |        88283 |    0.00
Benchmark_Gotiny_Unmarshal-10                 |   10847191 |    110 ns/op |        48 |        112 |   1.20 |        52066 |    0.98
Benchmark_GotinyNoTime_Marshal-10             |   18949262 |     62 ns/op |        48 |          0 |   1.18 |        90956 |    0.00
Benchmark_GotinyNoTime_Unmarshal-10           |   11873511 |    100 ns/op |        48 |         96 |   1.19 |        56992 |    1.05
Benchmark_Msgp_Marshal-10                     |   11368960 |    106 ns/op |        97 |        128 |   1.21 |       110278 |    0.83
Benchmark_Msgp_Unmarshal-10                   |    7026082 |    170 ns/op |        97 |        112 |   1.19 |        68152 |    1.52
Benchmark_VmihailencoMsgpack_Marshal-10       |    2125178 |    573 ns/op |       100 |        392 |   1.22 |        21251 |    1.46
Benchmark_VmihailencoMsgpack_Unmarshal-10     |    1348686 |    890 ns/op |       100 |        408 |   1.20 |        13486 |    2.18
Benchmark_Json_Marshal-10                     |    1440837 |    822 ns/op |       148 |        208 |   1.19 |        21425 |    3.96
Benchmark_Json_Unmarshal-10                   |     653754 |   1817 ns/op |       148 |        399 |   1.19 |         9721 |    4.55
Benchmark_JsonIter_Marshal-10                 |    1370590 |    865 ns/op |       138 |        248 |   1.19 |        18955 |    3.49
Benchmark_JsonIter_Unmarshal-10               |    1523732 |    799 ns/op |       138 |        263 |   1.22 |        21073 |    3.04
Benchmark_EasyJson_Marshal-10                 |    1613154 |    717 ns/op |       148 |        896 |   1.16 |        23987 |    0.80
Benchmark_EasyJson_Unmarshal-10               |    2010156 |    584 ns/op |       148 |        159 |   1.17 |        29891 |    3.67
Benchmark_Bson_Marshal-10                     |    1917372 |    621 ns/op |       110 |        376 |   1.19 |        21091 |    1.65
Benchmark_Bson_Unmarshal-10                   |    1250320 |    956 ns/op |       110 |        224 |   1.20 |        13753 |    4.27
Benchmark_MongoBson_Marshal-10                |    1263312 |    962 ns/op |       110 |        384 |   1.22 |        13896 |    2.51
Benchmark_MongoBson_Unmarshal-10              |    1000000 |   1002 ns/op |       110 |        408 |   1.00 |        11000 |    2.46
Benchmark_Gob_Marshal-10                      |    2750721 |    440 ns/op |        63 |         40 |   1.21 |        17494 |   11.01
Benchmark_Gob_Unmarshal-10                    |    2918254 |    410 ns/op |        63 |        112 |   1.20 |        18560 |    3.67
Benchmark_XDR_Marshal-10                      |    1000000 |   1002 ns/op |        88 |        376 |   1.00 |         8800 |    2.66
Benchmark_XDR_Unmarshal-10                    |    1432599 |    841 ns/op |        88 |        216 |   1.21 |        12606 |    3.90
Benchmark_UgorjiCodecMsgpack_Marshal-10       |    2054228 |    583 ns/op |        91 |       1304 |   1.20 |        18693 |    0.45
Benchmark_UgorjiCodecMsgpack_Unmarshal-10     |    1562437 |    762 ns/op |        91 |        496 |   1.19 |        14218 |    1.54
Benchmark_UgorjiCodecBinc_Marshal-10          |    1874546 |    638 ns/op |        95 |       1320 |   1.20 |        17808 |    0.48
Benchmark_UgorjiCodecBinc_Unmarshal-10        |    1316881 |    908 ns/op |        95 |        656 |   1.20 |        12510 |    1.38
Benchmark_Sereal_Marshal-10                   |     896787 |   1322 ns/op |       132 |        832 |   1.19 |        11837 |    1.59
Benchmark_Sereal_Unmarshal-10                 |     703737 |   1909 ns/op |       132 |        976 |   1.34 |         9289 |    1.96
Benchmark_Binary_Marshal-10                   |    1487862 |    764 ns/op |        61 |        312 |   1.14 |         9075 |    2.45
Benchmark_Binary_Unmarshal-10                 |    1370228 |    859 ns/op |        61 |        320 |   1.18 |         8358 |    2.69
Benchmark_FlatBuffers_Marshal-10              |    8232986 |    145 ns/op |        95 |          0 |   1.20 |        78427 |    0.00
Benchmark_FlatBuffers_Unmarshal-10            |    8787554 |    136 ns/op |        95 |        112 |   1.20 |        83490 |    1.22
Benchmark_CapNProto_Marshal-10                |    4019352 |    298 ns/op |        96 |         56 |   1.20 |        38585 |    5.33
Benchmark_CapNProto_Unmarshal-10              |    4814245 |    247 ns/op |        96 |        192 |   1.19 |        46216 |    1.29
Benchmark_CapNProto2_Marshal-10               |    3161616 |    368 ns/op |        96 |        244 |   1.17 |        30351 |    1.51
Benchmark_CapNProto2_Unmarshal-10             |    2804642 |    428 ns/op |        96 |        304 |   1.20 |        26924 |    1.41
Benchmark_Hprose_Marshal-10                   |    2958571 |    404 ns/op |        82 |        364 |   1.20 |        24337 |    1.11
Benchmark_Hprose_Unmarshal-10                 |    1900216 |    619 ns/op |        82 |        303 |   1.18 |        15634 |    2.04
Benchmark_Hprose2_Marshal-10                  |    4231543 |    284 ns/op |        82 |          0 |   1.20 |        34821 |    0.00
Benchmark_Hprose2_Unmarshal-10                |    4173704 |    287 ns/op |        82 |        136 |   1.20 |        34332 |    2.12
Benchmark_Protobuf_Marshal-10                 |    2852079 |    420 ns/op |        52 |        144 |   1.20 |        14830 |    2.92
Benchmark_Protobuf_Unmarshal-10               |    3327733 |    366 ns/op |        52 |        184 |   1.22 |        17304 |    1.99
Benchmark_Goprotobuf_Marshal-10               |    6831308 |    176 ns/op |        53 |         64 |   1.20 |        36205 |    2.75
Benchmark_Goprotobuf_Unmarshal-10             |    5746256 |    210 ns/op |        53 |        168 |   1.21 |        30455 |    1.25
Benchmark_Gogoprotobuf_Marshal-10             |   16528346 |     72 ns/op |        53 |         64 |   1.20 |        87600 |    1.13
Benchmark_Gogoprotobuf_Unmarshal-10           |   12764978 |     94 ns/op |        53 |         96 |   1.21 |        67654 |    0.99
Benchmark_Gogojsonpb_Marshal-10               |     203168 |   5914 ns/op |       125 |       3165 |   1.20 |         2551 |    1.87
Benchmark_Gogojsonpb_Unmarshal-10             |     142018 |   8337 ns/op |       126 |       3384 |   1.18 |         1790 |    2.46
Benchmark_Colfer_Marshal-10                   |   17383452 |     68 ns/op |        51 |         64 |   1.19 |        88864 |    1.07
Benchmark_Colfer_Unmarshal-10                 |   13542744 |     87 ns/op |        52 |        112 |   1.19 |        70422 |    0.79
Benchmark_Gencode_Marshal-10                  |   14423618 |     83 ns/op |        53 |         80 |   1.20 |        76445 |    1.04
Benchmark_Gencode_Unmarshal-10                |   13903652 |     86 ns/op |        53 |        112 |   1.20 |        73689 |    0.77
Benchmark_GencodeUnsafe_Marshal-10            |   22381569 |     54 ns/op |        46 |         48 |   1.21 |       102955 |    1.13
Benchmark_GencodeUnsafe_Unmarshal-10          |   17110041 |     68 ns/op |        46 |         96 |   1.18 |        78706 |    0.72
Benchmark_XDR2_Marshal-10                     |   12394992 |     95 ns/op |        60 |         64 |   1.18 |        74369 |    1.49
Benchmark_XDR2_Unmarshal-10                   |   17977414 |     66 ns/op |        60 |         32 |   1.20 |       107864 |    2.09
Benchmark_GoAvro_Marshal-10                   |     951213 |   1232 ns/op |        47 |        728 |   1.17 |         4470 |    1.69
Benchmark_GoAvro_Unmarshal-10                 |     436666 |   2736 ns/op |        47 |       2544 |   1.19 |         2052 |    1.08
Benchmark_GoAvro2Text_Marshal-10              |     911409 |   1324 ns/op |       133 |       1320 |   1.21 |        12194 |    1.00
Benchmark_GoAvro2Text_Unmarshal-10            |     913957 |   1324 ns/op |       133 |        783 |   1.21 |        12228 |    1.69
Benchmark_GoAvro2Binary_Marshal-10            |    2920245 |    407 ns/op |        47 |        464 |   1.19 |        13725 |    0.88
Benchmark_GoAvro2Binary_Unmarshal-10          |    2711745 |    442 ns/op |        47 |        544 |   1.20 |        12745 |    0.81
Benchmark_Ikea_Marshal-10                     |    3455250 |    346 ns/op |        55 |         72 |   1.20 |        19003 |    4.81
Benchmark_Ikea_Unmarshal-10                   |    2895451 |    414 ns/op |        55 |        160 |   1.20 |        15924 |    2.59
Benchmark_ShamatonMapMsgpack_Marshal-10       |    2671519 |    451 ns/op |        92 |        192 |   1.21 |        24577 |    2.35
Benchmark_ShamatonMapMsgpack_Unmarshal-10     |    3099800 |    386 ns/op |        92 |        136 |   1.20 |        28518 |    2.84
Benchmark_ShamatonArrayMsgpack_Marshal-10     |    2917430 |    412 ns/op |        50 |        160 |   1.20 |        14587 |    2.58
Benchmark_ShamatonArrayMsgpack_Unmarshal-10   |    4600650 |    260 ns/op |        50 |        136 |   1.20 |        23003 |    1.92
Benchmark_ShamatonMapMsgpackgen_Marshal-10    |    8721518 |    131 ns/op |        92 |         96 |   1.15 |        80237 |    1.37
Benchmark_ShamatonMapMsgpackgen_Unmarshal-10  |    7956811 |    143 ns/op |        92 |         80 |   1.14 |        73202 |    1.80
Benchmark_ShamatonArrayMsgpackgen_Marshal-10  |   10916773 |    106 ns/op |        50 |         64 |   1.16 |        54583 |    1.66
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-10 |   13831982 |     86 ns/op |        50 |         80 |   1.20 |        69159 |    1.08
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-10 |     385651 |   3180 ns/op |        55 |        440 |   1.23 |         2121 |    7.23
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-10 |     262132 |   4557 ns/op |        55 |       1184 |   1.19 |         1441 |    3.85
Benchmark_Mum_Marshal-10                      |   18999441 |     63 ns/op |        48 |          0 |   1.20 |        91197 |    0.00
Benchmark_Mum_Unmarshal-10                    |    9653524 |    125 ns/op |        48 |         80 |   1.21 |        46336 |    1.57
Benchmark_Bebop_Marshal-10                    |   16134940 |     73 ns/op |        55 |         64 |   1.19 |        88742 |    1.15
Benchmark_Bebop_Unmarshal-10                  |   21126357 |     56 ns/op |        55 |         32 |   1.20 |       116194 |    1.77
Benchmark_FastJson_Marshal-10                 |    3520270 |    339 ns/op |       133 |        504 |   1.20 |        47101 |    0.67
Benchmark_FastJson_Unmarshal-10               |    1443300 |    837 ns/op |       133 |       1704 |   1.21 |        19311 |    0.49
Benchmark_MusgoUnsafe_Marshal-10              |   21967660 |     53 ns/op |        48 |          0 |   1.17 |       106609 |    0.00
Benchmark_MusgoUnsafe_Unmarshal-10            |   17176309 |     70 ns/op |        48 |         64 |   1.21 |        83236 |    1.10
Benchmark_Musgo_Marshal-10                    |   22535546 |     53 ns/op |        48 |          0 |   1.21 |       109319 |    0.00
Benchmark_Musgo_Unmarshal-10                  |   12952696 |     90 ns/op |        48 |         96 |   1.17 |        62820 |    0.94


Totals:


benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_GencodeUnsafe_-10              |   39491610 |    122 ns/op |        92 |        144 |   4.86 |       363322 |    0.85
Benchmark_MusgoUnsafe_-10                |   39143969 |    123 ns/op |        96 |         64 |   4.84 |       379657 |    1.93
Benchmark_Bebop_-10                      |   37261297 |    130 ns/op |       110 |         96 |   4.86 |       409874 |    1.36
Benchmark_Musgo_-10                      |   35488242 |    143 ns/op |        97 |         96 |   5.11 |       344271 |    1.50
Benchmark_Colfer_-10                     |   30926196 |    156 ns/op |       103 |        176 |   4.84 |       318910 |    0.89
Benchmark_XDR2_-10                       |   30372406 |    162 ns/op |       120 |         96 |   4.93 |       364468 |    1.69
Benchmark_GotinyNoTime_-10               |   30822773 |    162 ns/op |        96 |         96 |   5.02 |       295898 |    1.70
Benchmark_Gogoprotobuf_-10               |   29293324 |    167 ns/op |       106 |        160 |   4.90 |       310509 |    1.04
Benchmark_Gencode_-10                    |   28327270 |    169 ns/op |       106 |        192 |   4.80 |       300269 |    0.88
Benchmark_Gotiny_-10                     |   29239545 |    174 ns/op |        96 |        112 |   5.10 |       280699 |    1.56
Benchmark_Mum_-10                        |   28652965 |    188 ns/op |        96 |         80 |   5.41 |       275068 |    2.36
Benchmark_ShamatonArrayMsgpackgen_-10    |   24748755 |    192 ns/op |       100 |        144 |   4.77 |       247487 |    1.34
Benchmark_ShamatonMapMsgpackgen_-10      |   16678329 |    275 ns/op |       184 |        176 |   4.60 |       306881 |    1.57
Benchmark_Msgp_-10                       |   18395042 |    276 ns/op |       194 |        240 |   5.08 |       356863 |    1.15
Benchmark_FlatBuffers_-10                |   17020540 |    281 ns/op |       190 |        112 |   4.80 |       323849 |    2.52
Benchmark_Goprotobuf_-10                 |   12577564 |    386 ns/op |       106 |        232 |   4.86 |       133322 |    1.67
Benchmark_CapNProto_-10                  |    8833597 |    545 ns/op |       192 |        248 |   4.82 |       169605 |    2.20
Benchmark_Hprose2_-10                    |    8405247 |    572 ns/op |       164 |        136 |   4.81 |       138308 |    4.21
Benchmark_ShamatonArrayMsgpack_-10       |    7518080 |    672 ns/op |       100 |        296 |   5.06 |        75180 |    2.27
Benchmark_Ikea_-10                       |    6350701 |    761 ns/op |       110 |        232 |   4.83 |        69857 |    3.28
Benchmark_Protobuf_-10                   |    6179812 |    787 ns/op |       104 |        328 |   4.87 |        64270 |    2.40
Benchmark_CapNProto2_-10                 |    5966258 |    796 ns/op |       192 |        548 |   4.75 |       114552 |    1.45
Benchmark_ShamatonMapMsgpack_-10         |    5771319 |    837 ns/op |       184 |        328 |   4.83 |       106192 |    2.55
Benchmark_GoAvro2Binary_-10              |    5631990 |    850 ns/op |        94 |       1008 |   4.79 |        52940 |    0.84
Benchmark_Gob_-10                        |    5668975 |    851 ns/op |       127 |        152 |   4.83 |        72109 |    5.60
Benchmark_Hprose_-10                     |    4858787 |   1023 ns/op |       164 |        667 |   4.97 |        79946 |    1.53
Benchmark_FastJson_-10                   |    4963570 |   1177 ns/op |       267 |       2208 |   5.84 |       132825 |    0.53
Benchmark_EasyJson_-10                   |    3623310 |   1301 ns/op |       297 |       1055 |   4.71 |       107757 |    1.23
Benchmark_UgorjiCodecMsgpack_-10         |    3616665 |   1346 ns/op |       182 |       1800 |   4.87 |        65823 |    0.75
Benchmark_VmihailencoMsgpack_-10         |    3473864 |   1464 ns/op |       200 |        800 |   5.09 |        69477 |    1.83
Benchmark_UgorjiCodecBinc_-10            |    3191427 |   1546 ns/op |       190 |       1976 |   4.93 |        60637 |    0.78
Benchmark_Bson_-10                       |    3167692 |   1578 ns/op |       220 |        600 |   5.00 |        69689 |    2.63
Benchmark_Binary_-10                     |    2858090 |   1623 ns/op |       122 |        632 |   4.64 |        34868 |    2.57
Benchmark_JsonIter_-10                   |    2894322 |   1665 ns/op |       276 |        511 |   4.82 |        80056 |    3.26
Benchmark_XDR_-10                        |    2432599 |   1843 ns/op |       176 |        592 |   4.48 |        42813 |    3.11
Benchmark_MongoBson_-10                  |    2263312 |   1964 ns/op |       220 |        792 |   4.45 |        49792 |    2.48
Benchmark_Json_-10                       |    2094591 |   2639 ns/op |       297 |        607 |   5.53 |        62293 |    4.35
Benchmark_GoAvro2Text_-10                |    1825366 |   2648 ns/op |       267 |       2103 |   4.83 |        48846 |    1.26
Benchmark_Sereal_-10                     |    1600524 |   3231 ns/op |       264 |       1808 |   5.17 |        42253 |    1.79
Benchmark_GoAvro_-10                     |    1387879 |   3968 ns/op |        94 |       3272 |   5.51 |        13046 |    1.21
Benchmark_SSZNoTimeNoStringNoFloatA_-10  |     647783 |   7737 ns/op |       110 |       1624 |   5.01 |         7125 |    4.76
Benchmark_Gogojsonpb_-10                 |     345186 |  14251 ns/op |       251 |       6549 |   4.92 |         8688 |    2.18



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
