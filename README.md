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

2022-08-16 Results with Go 1.16.5 Darwin/arm64 on an Apple M1

benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                    |    5157031 |    216 ns/op |        48 |        168 |   1.12 |        24753 |    1.29
Benchmark_Gotiny_Unmarshal-8                  |   10985799 |    109 ns/op |        48 |        112 |   1.20 |        52731 |    0.97
Benchmark_GotinyNoTime_Marshal-8              |    5549127 |    214 ns/op |        48 |        168 |   1.19 |        26635 |    1.27
Benchmark_GotinyNoTime_Unmarshal-8            |   12082882 |     99 ns/op |        48 |         96 |   1.20 |        57997 |    1.04
Benchmark_Msgp_Marshal-8                      |   11179303 |    107 ns/op |        97 |        128 |   1.20 |       108439 |    0.84
Benchmark_Msgp_Unmarshal-8                    |    7097386 |    169 ns/op |        97 |        112 |   1.20 |        68844 |    1.51
Benchmark_VmihailencoMsgpack_Marshal-8        |    2139926 |    560 ns/op |       100 |        392 |   1.20 |        21399 |    1.43
Benchmark_VmihailencoMsgpack_Unmarshal-8      |    1371309 |    876 ns/op |       100 |        408 |   1.20 |        13713 |    2.15
Benchmark_Json_Marshal-8                      |    1383134 |    868 ns/op |       148 |        208 |   1.20 |        20553 |    4.17
Benchmark_Json_Unmarshal-8                    |     676347 |   1764 ns/op |       148 |        383 |   1.19 |        10057 |    4.61
Benchmark_JsonIter_Marshal-8                  |    1309188 |    915 ns/op |       138 |        248 |   1.20 |        18106 |    3.69
Benchmark_JsonIter_Unmarshal-8                |    1531534 |    783 ns/op |       138 |        263 |   1.20 |        21181 |    2.98
Benchmark_EasyJson_Marshal-8                  |    1702153 |    700 ns/op |       148 |        896 |   1.19 |        25311 |    0.78
Benchmark_EasyJson_Unmarshal-8                |    2073801 |    581 ns/op |       148 |        159 |   1.21 |        30816 |    3.66
Benchmark_Bson_Marshal-8                      |    1967242 |    610 ns/op |       110 |        376 |   1.20 |        21639 |    1.62
Benchmark_Bson_Unmarshal-8                    |    1257063 |    955 ns/op |       110 |        224 |   1.20 |        13827 |    4.26
Benchmark_MongoBson_Marshal-8                 |    1270983 |    946 ns/op |       110 |        384 |   1.20 |        13980 |    2.46
Benchmark_MongoBson_Unmarshal-8               |    1219966 |    984 ns/op |       110 |        408 |   1.20 |        13419 |    2.41
Benchmark_Gob_Marshal-8                       |     379551 |   3164 ns/op |       163 |       1616 |   1.20 |         6209 |    1.96
Benchmark_Gob_Unmarshal-8                     |      75757 |  15827 ns/op |       163 |       7688 |   1.20 |         1240 |    2.06
Benchmark_XDR_Marshal-8                       |    1270831 |    945 ns/op |        88 |        376 |   1.20 |        11183 |    2.51
Benchmark_XDR_Unmarshal-8                     |    1469666 |    816 ns/op |        88 |        216 |   1.20 |        12933 |    3.78
Benchmark_UgorjiCodecMsgpack_Marshal-8        |    2148814 |    558 ns/op |        91 |       1304 |   1.20 |        19554 |    0.43
Benchmark_UgorjiCodecMsgpack_Unmarshal-8      |    1590700 |    754 ns/op |        91 |        496 |   1.20 |        14475 |    1.52
Benchmark_UgorjiCodecBinc_Marshal-8           |    1992244 |    602 ns/op |        95 |       1320 |   1.20 |        18926 |    0.46
Benchmark_UgorjiCodecBinc_Unmarshal-8         |    1358884 |    882 ns/op |        95 |        656 |   1.20 |        12909 |    1.35
Benchmark_Sereal_Marshal-8                    |     935347 |   1272 ns/op |       132 |        832 |   1.19 |        12346 |    1.53
Benchmark_Sereal_Unmarshal-8                  |     796164 |   1497 ns/op |       132 |        976 |   1.19 |        10509 |    1.53
Benchmark_Binary_Marshal-8                    |    1652059 |    728 ns/op |        61 |        312 |   1.20 |        10077 |    2.33
Benchmark_Binary_Unmarshal-8                  |    1427588 |    842 ns/op |        61 |        320 |   1.20 |         8708 |    2.63
Benchmark_FlatBuffers_Marshal-8               |    3138784 |    381 ns/op |        95 |        376 |   1.20 |        29884 |    1.01
Benchmark_FlatBuffers_Unmarshal-8             |    8881845 |    135 ns/op |        95 |        112 |   1.20 |        84564 |    1.21
Benchmark_CapNProto_Marshal-8                 |    1485355 |    811 ns/op |        96 |       4488 |   1.21 |        14259 |    0.18
Benchmark_CapNProto_Unmarshal-8               |    4934485 |    243 ns/op |        96 |        192 |   1.20 |        47371 |    1.27
Benchmark_CapNProto2_Marshal-8                |    2063122 |    579 ns/op |        96 |       1452 |   1.20 |        19805 |    0.40
Benchmark_CapNProto2_Unmarshal-8              |    4443772 |    270 ns/op |        96 |        272 |   1.20 |        42660 |    0.99
Benchmark_Hprose_Marshal-8                    |    2847484 |    419 ns/op |        82 |        379 |   1.19 |        23429 |    1.11
Benchmark_Hprose_Unmarshal-8                  |    1947406 |    615 ns/op |        82 |        304 |   1.20 |        16023 |    2.02
Benchmark_Hprose2_Marshal-8                   |    3723235 |    325 ns/op |        82 |          0 |   1.21 |        30631 |    0.00
Benchmark_Hprose2_Unmarshal-8                 |    4194978 |    286 ns/op |        82 |        136 |   1.20 |        34528 |    2.11
Benchmark_Protobuf_Marshal-8                  |    2804445 |    427 ns/op |        52 |        144 |   1.20 |        14583 |    2.97
Benchmark_Protobuf_Unmarshal-8                |    3319700 |    361 ns/op |        52 |        184 |   1.20 |        17262 |    1.96
Benchmark_Goprotobuf_Marshal-8                |    6751698 |    176 ns/op |        53 |         64 |   1.19 |        35783 |    2.76
Benchmark_Goprotobuf_Unmarshal-8              |    5818316 |    207 ns/op |        53 |        168 |   1.20 |        30837 |    1.23
Benchmark_Gogoprotobuf_Marshal-8              |   16825938 |     71 ns/op |        53 |         64 |   1.21 |        89177 |    1.12
Benchmark_Gogoprotobuf_Unmarshal-8            |   12882285 |     93 ns/op |        53 |         96 |   1.20 |        68276 |    0.97
Benchmark_Gogojsonpb_Marshal-8                |     213758 |   5622 ns/op |       125 |       3117 |   1.20 |         2682 |    1.80
Benchmark_Gogojsonpb_Unmarshal-8              |     151483 |   7923 ns/op |       125 |       3610 |   1.20 |         1902 |    2.19
Benchmark_Colfer_Marshal-8                    |   18066087 |     66 ns/op |        51 |         64 |   1.20 |        92426 |    1.03
Benchmark_Colfer_Unmarshal-8                  |   14401274 |     83 ns/op |        51 |        112 |   1.21 |        73446 |    0.75
Benchmark_Gencode_Marshal-8                   |   14489538 |     82 ns/op |        53 |         80 |   1.20 |        76794 |    1.03
Benchmark_Gencode_Unmarshal-8                 |   14329048 |     83 ns/op |        53 |        112 |   1.20 |        75943 |    0.75
Benchmark_GencodeUnsafe_Marshal-8             |   22314405 |     52 ns/op |        46 |         48 |   1.18 |       102646 |    1.10
Benchmark_GencodeUnsafe_Unmarshal-8           |   18527250 |     65 ns/op |        46 |         96 |   1.21 |        85225 |    0.68
Benchmark_XDR2_Marshal-8                      |   12781372 |     93 ns/op |        60 |         64 |   1.20 |        76688 |    1.47
Benchmark_XDR2_Unmarshal-8                    |   17961876 |     66 ns/op |        60 |         32 |   1.19 |       107771 |    2.07
Benchmark_GoAvro_Marshal-8                    |     902552 |   1330 ns/op |        47 |        872 |   1.20 |         4241 |    1.53
Benchmark_GoAvro_Unmarshal-8                  |     381249 |   3192 ns/op |        47 |       3024 |   1.22 |         1791 |    1.06
Benchmark_GoAvro2Text_Marshal-8               |     959385 |   1259 ns/op |       133 |       1320 |   1.21 |        12836 |    0.95
Benchmark_GoAvro2Text_Unmarshal-8             |     934458 |   1282 ns/op |       133 |        783 |   1.20 |        12503 |    1.64
Benchmark_GoAvro2Binary_Marshal-8             |    3020731 |    396 ns/op |        47 |        464 |   1.20 |        14197 |    0.85
Benchmark_GoAvro2Binary_Unmarshal-8           |    2812490 |    426 ns/op |        47 |        544 |   1.20 |        13218 |    0.78
Benchmark_Ikea_Marshal-8                      |    2206000 |    544 ns/op |        55 |        184 |   1.20 |        12133 |    2.96
Benchmark_Ikea_Unmarshal-8                    |    2875797 |    416 ns/op |        55 |        160 |   1.20 |        15816 |    2.60
Benchmark_ShamatonMapMsgpack_Marshal-8        |    2647074 |    452 ns/op |        92 |        192 |   1.20 |        24353 |    2.35
Benchmark_ShamatonMapMsgpack_Unmarshal-8      |    3114541 |    386 ns/op |        92 |        136 |   1.20 |        28653 |    2.84
Benchmark_ShamatonArrayMsgpack_Marshal-8      |    2907676 |    412 ns/op |        50 |        160 |   1.20 |        14538 |    2.58
Benchmark_ShamatonArrayMsgpack_Unmarshal-8    |    4661485 |    257 ns/op |        50 |        136 |   1.20 |        23307 |    1.89
Benchmark_ShamatonMapMsgpackgen_Marshal-8     |    9296109 |    128 ns/op |        92 |         96 |   1.20 |        85524 |    1.34
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8   |    8447602 |    142 ns/op |        92 |         80 |   1.20 |        77717 |    1.78
Benchmark_ShamatonArrayMsgpackgen_Marshal-8   |   11615052 |    103 ns/op |        50 |         64 |   1.20 |        58075 |    1.62
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8 |   14123607 |     85 ns/op |        50 |         80 |   1.20 |        70618 |    1.06
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8 |     390501 |   3039 ns/op |        55 |        440 |   1.19 |         2147 |    6.91
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8 |     268053 |   4483 ns/op |        55 |       1184 |   1.20 |         1474 |    3.79
Benchmark_Mum_Marshal-8                       |   18664054 |     63 ns/op |        48 |          0 |   1.19 |        89587 |    0.00
Benchmark_Mum_Unmarshal-8                     |    9962642 |    120 ns/op |        48 |         80 |   1.20 |        47820 |    1.51
Benchmark_Bebop_Marshal-8                     |   16628799 |     71 ns/op |        55 |         64 |   1.19 |        91458 |    1.12
Benchmark_Bebop_Unmarshal-8                   |   21506419 |     55 ns/op |        55 |         32 |   1.19 |       118285 |    1.73
Benchmark_FastJson_Marshal-8                  |    3820375 |    312 ns/op |       133 |        504 |   1.19 |        51116 |    0.62
Benchmark_FastJson_Unmarshal-8                |    1539577 |    779 ns/op |       133 |       1704 |   1.20 |        20599 |    0.46
Benchmark_MusgoUnsafe_Marshal-8               |   14030799 |     84 ns/op |        48 |         55 |   1.19 |        68035 |    1.54
Benchmark_MusgoUnsafe_Unmarshal-8             |   16834632 |     71 ns/op |        48 |         64 |   1.20 |        81647 |    1.11
Benchmark_Musgo_Marshal-8                     |   14076384 |     84 ns/op |        48 |         55 |   1.19 |        68242 |    1.54
Benchmark_Musgo_Unmarshal-8                   |   13497546 |     89 ns/op |        48 |         96 |   1.20 |        65490 |    0.93


Totals:


benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_GencodeUnsafe_-8               |   40841655 |    118 ns/op |        92 |        144 |   4.83 |       375743 |    0.82
Benchmark_Bebop_-8                       |   38135218 |    127 ns/op |       110 |         96 |   4.85 |       419487 |    1.32
Benchmark_Colfer_-8                      |   32467361 |    149 ns/op |       102 |        176 |   4.87 |       331686 |    0.85
Benchmark_MusgoUnsafe_-8                 |   30865431 |    155 ns/op |        96 |        119 |   4.81 |       299363 |    1.31
Benchmark_XDR2_-8                        |   30743248 |    160 ns/op |       120 |         96 |   4.92 |       368918 |    1.67
Benchmark_Gogoprotobuf_-8                |   29708223 |    164 ns/op |       106 |        160 |   4.90 |       314907 |    1.03
Benchmark_Gencode_-8                     |   28818586 |    166 ns/op |       106 |        192 |   4.80 |       305477 |    0.87
Benchmark_Musgo_-8                       |   27573930 |    173 ns/op |        97 |        151 |   4.79 |       267467 |    1.15
Benchmark_Mum_-8                         |   28626696 |    184 ns/op |        96 |         80 |   5.28 |       274816 |    2.30
Benchmark_ShamatonArrayMsgpackgen_-8     |   25738659 |    188 ns/op |       100 |        144 |   4.85 |       257386 |    1.31
Benchmark_ShamatonMapMsgpackgen_-8       |   17743711 |    271 ns/op |       184 |        176 |   4.81 |       326484 |    1.54
Benchmark_Msgp_-8                        |   18276689 |    276 ns/op |       194 |        240 |   5.05 |       354567 |    1.15
Benchmark_GotinyNoTime_-8                |   17632009 |    313 ns/op |        96 |        264 |   5.53 |       169267 |    1.19
Benchmark_Gotiny_-8                      |   16142830 |    325 ns/op |        96 |        280 |   5.25 |       154971 |    1.16
Benchmark_Goprotobuf_-8                  |   12570014 |    383 ns/op |       106 |        232 |   4.82 |       133242 |    1.65
Benchmark_FlatBuffers_-8                 |   12020629 |    517 ns/op |       190 |        488 |   6.22 |       228896 |    1.06
Benchmark_Hprose2_-8                     |    7918213 |    611 ns/op |       164 |        136 |   4.84 |       130317 |    4.50
Benchmark_ShamatonArrayMsgpack_-8        |    7569161 |    670 ns/op |       100 |        296 |   5.07 |        75691 |    2.26
Benchmark_Protobuf_-8                    |    6124145 |    789 ns/op |       104 |        328 |   4.83 |        63691 |    2.41
Benchmark_GoAvro2Binary_-8               |    5833221 |    822 ns/op |        94 |       1008 |   4.80 |        54832 |    0.82
Benchmark_ShamatonMapMsgpack_-8          |    5761615 |    838 ns/op |       184 |        328 |   4.83 |       106013 |    2.55
Benchmark_CapNProto2_-8                  |    6506894 |    849 ns/op |       192 |       1724 |   5.53 |       124932 |    0.49
Benchmark_Ikea_-8                        |    5081797 |    961 ns/op |       110 |        344 |   4.88 |        55899 |    2.79
Benchmark_Hprose_-8                      |    4794890 |   1035 ns/op |       164 |        683 |   4.96 |        78904 |    1.52
Benchmark_CapNProto_-8                   |    6419840 |   1055 ns/op |       192 |       4680 |   6.77 |       123260 |    0.23
Benchmark_FastJson_-8                    |    5359952 |   1091 ns/op |       267 |       2208 |   5.85 |       143432 |    0.49
Benchmark_EasyJson_-8                    |    3775954 |   1282 ns/op |       297 |       1055 |   4.84 |       112259 |    1.22
Benchmark_UgorjiCodecMsgpack_-8          |    3739514 |   1312 ns/op |       182 |       1800 |   4.91 |        68059 |    0.73
Benchmark_VmihailencoMsgpack_-8          |    3511235 |   1436 ns/op |       200 |        800 |   5.04 |        70224 |    1.80
Benchmark_UgorjiCodecBinc_-8             |    3351128 |   1485 ns/op |       190 |       1976 |   4.98 |        63671 |    0.75
Benchmark_Bson_-8                        |    3224305 |   1566 ns/op |       220 |        600 |   5.05 |        70934 |    2.61
Benchmark_Binary_-8                      |    3079647 |   1570 ns/op |       122 |        632 |   4.84 |        37571 |    2.49
Benchmark_JsonIter_-8                    |    2840722 |   1698 ns/op |       276 |        511 |   4.82 |        78574 |    3.32
Benchmark_XDR_-8                         |    2740497 |   1762 ns/op |       176 |        592 |   4.83 |        48232 |    2.98
Benchmark_MongoBson_-8                   |    2490949 |   1930 ns/op |       220 |        792 |   4.81 |        54800 |    2.44
Benchmark_GoAvro2Text_-8                 |    1893843 |   2541 ns/op |       267 |       2103 |   4.81 |        50679 |    1.21
Benchmark_Json_-8                        |    2059481 |   2632 ns/op |       297 |        591 |   5.42 |        61228 |    4.45
Benchmark_Sereal_-8                      |    1731511 |   2769 ns/op |       264 |       1808 |   4.79 |        45711 |    1.53
Benchmark_GoAvro_-8                      |    1283801 |   4522 ns/op |        94 |       3896 |   5.81 |        12067 |    1.16
Benchmark_SSZNoTimeNoStringNoFloatA_-8   |     658554 |   7522 ns/op |       110 |       1624 |   4.95 |         7244 |    4.63
Benchmark_Gogojsonpb_-8                  |     365241 |  13545 ns/op |       251 |       6727 |   4.95 |         9171 |    2.01
Benchmark_Gob_-8                         |     455308 |  18991 ns/op |       327 |       9304 |   8.65 |        14902 |    2.04


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
