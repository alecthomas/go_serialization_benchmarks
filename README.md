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

2021-06-21 Results with Go 1.16.5 linux/amd64 on an Intel(R) Core(TM) i7-3630QM CPU @ 2.40GHz

benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                    |    8161915 |    144 ns/op |        48 |          0 |   1.18 |        39177 |    0.00
Benchmark_Gotiny_Unmarshal-8                  |    3062521 |    396 ns/op |        48 |        112 |   1.21 |        14700 |    3.54
Benchmark_GotinyNoTime_Marshal-8              |    8340646 |    134 ns/op |        48 |          0 |   1.12 |        40035 |    0.00
Benchmark_GotinyNoTime_Unmarshal-8            |    3323733 |    361 ns/op |        48 |         96 |   1.20 |        15953 |    3.77
Benchmark_Msgp_Marshal-8                      |    3723291 |    289 ns/op |        97 |        128 |   1.08 |        36115 |    2.26
Benchmark_Msgp_Unmarshal-8                    |    2651005 |    467 ns/op |        97 |        112 |   1.24 |        25714 |    4.17
Benchmark_VmihailencoMsgpack_Marshal-8        |     856572 |   1605 ns/op |       100 |        392 |   1.37 |         8565 |    4.09
Benchmark_VmihailencoMsgpack_Unmarshal-8      |     461299 |   2329 ns/op |       100 |        408 |   1.07 |         4612 |    5.71
Benchmark_Json_Marshal-8                      |     501478 |   2538 ns/op |       151 |        208 |   1.27 |         7602 |   12.20
Benchmark_Json_Unmarshal-8                    |     226456 |   5023 ns/op |       151 |        383 |   1.14 |         3435 |   13.11
Benchmark_JsonIter_Marshal-8                  |     440287 |   2611 ns/op |       141 |        248 |   1.15 |         6221 |   10.53
Benchmark_JsonIter_Unmarshal-8                |     582590 |   2232 ns/op |       141 |        264 |   1.30 |         8231 |    8.45
Benchmark_EasyJson_Marshal-8                  |     520345 |   2520 ns/op |       151 |        896 |   1.31 |         7893 |    2.81
Benchmark_EasyJson_Unmarshal-8                |     666243 |   1827 ns/op |       151 |        160 |   1.22 |        10100 |   11.42
Benchmark_Bson_Marshal-8                      |     723712 |   1603 ns/op |       110 |        376 |   1.16 |         7960 |    4.26
Benchmark_Bson_Unmarshal-8                    |     564585 |   2151 ns/op |       110 |        224 |   1.21 |         6210 |    9.60
Benchmark_Gob_Marshal-8                       |    1320562 |    882 ns/op |        63 |         40 |   1.16 |         8394 |   22.05
Benchmark_Gob_Unmarshal-8                     |    1000000 |   1041 ns/op |        63 |        112 |   1.04 |         6357 |    9.29
Benchmark_XDR_Marshal-8                       |     592130 |   2350 ns/op |        92 |        440 |   1.39 |         5447 |    5.34
Benchmark_XDR_Unmarshal-8                     |     579439 |   2160 ns/op |        92 |        232 |   1.25 |         5330 |    9.31
Benchmark_UgorjiCodecMsgpack_Marshal-8        |     827307 |   2137 ns/op |        91 |       1304 |   1.77 |         7528 |    1.64
Benchmark_UgorjiCodecMsgpack_Unmarshal-8      |     566781 |   2076 ns/op |        91 |        496 |   1.18 |         5157 |    4.19
Benchmark_UgorjiCodecBinc_Marshal-8           |     547809 |   2256 ns/op |        95 |       1320 |   1.24 |         5204 |    1.71
Benchmark_UgorjiCodecBinc_Unmarshal-8         |     438355 |   2323 ns/op |        95 |        656 |   1.02 |         4164 |    3.54
Benchmark_Sereal_Marshal-8                    |     287544 |   4017 ns/op |       132 |        832 |   1.16 |         3795 |    4.83
Benchmark_Sereal_Unmarshal-8                  |     269700 |   4571 ns/op |       132 |        976 |   1.23 |         3560 |    4.68
Benchmark_Binary_Marshal-8                    |     645805 |   1804 ns/op |        61 |        312 |   1.17 |         3939 |    5.78
Benchmark_Binary_Unmarshal-8                  |     681712 |   1824 ns/op |        61 |        320 |   1.24 |         4158 |    5.70
Benchmark_FlatBuffers_Marshal-8               |    4016948 |    283 ns/op |        95 |          0 |   1.14 |        38233 |    0.00
Benchmark_FlatBuffers_Unmarshal-8             |    2870288 |    390 ns/op |        95 |        112 |   1.12 |        27293 |    3.48
Benchmark_CapNProto_Marshal-8                 |    2463114 |    486 ns/op |        96 |         56 |   1.20 |        23645 |    8.69
Benchmark_CapNProto_Unmarshal-8               |    1854186 |    659 ns/op |        96 |        192 |   1.22 |        17800 |    3.44
Benchmark_CapNProto2_Marshal-8                |    1304802 |    927 ns/op |        96 |        244 |   1.21 |        12526 |    3.80
Benchmark_CapNProto2_Unmarshal-8              |    1000000 |   1156 ns/op |        96 |        304 |   1.16 |         9600 |    3.80
Benchmark_Hprose_Marshal-8                    |    1278418 |    914 ns/op |        85 |        376 |   1.17 |        10893 |    2.43
Benchmark_Hprose_Unmarshal-8                  |     887418 |   1539 ns/op |        85 |        304 |   1.37 |         7567 |    5.06
Benchmark_Hprose2_Marshal-8                   |    2158150 |    569 ns/op |        85 |          0 |   1.23 |        18396 |    0.00
Benchmark_Hprose2_Unmarshal-8                 |    1697797 |    741 ns/op |        85 |        136 |   1.26 |        14473 |    5.45
Benchmark_Protobuf_Marshal-8                  |    1253037 |    942 ns/op |        52 |        144 |   1.18 |         6515 |    6.55
Benchmark_Protobuf_Unmarshal-8                |    1000000 |   1037 ns/op |        52 |        184 |   1.04 |         5200 |    5.64
Benchmark_Goprotobuf_Marshal-8                |    3247056 |    378 ns/op |        53 |         64 |   1.23 |        17209 |    5.92
Benchmark_Goprotobuf_Unmarshal-8              |    1839267 |    651 ns/op |        53 |        168 |   1.20 |         9748 |    3.88
Benchmark_Gogoprotobuf_Marshal-8              |    5886194 |    204 ns/op |        53 |         64 |   1.20 |        31196 |    3.20
Benchmark_Gogoprotobuf_Unmarshal-8            |    3464098 |    345 ns/op |        53 |         96 |   1.20 |        18359 |    3.60
Benchmark_Colfer_Marshal-8                    |    6263342 |    199 ns/op |        51 |         64 |   1.25 |        32024 |    3.11
Benchmark_Colfer_Unmarshal-8                  |    3480002 |    332 ns/op |        52 |        112 |   1.16 |        18096 |    2.97
Benchmark_Gencode_Marshal-8                   |    4923776 |    237 ns/op |        53 |         80 |   1.17 |        26096 |    2.97
Benchmark_Gencode_Unmarshal-8                 |    3545628 |    320 ns/op |        53 |        112 |   1.14 |        18791 |    2.87
Benchmark_GencodeUnsafe_Marshal-8             |    7635088 |    151 ns/op |        46 |         48 |   1.15 |        35121 |    3.15
Benchmark_GencodeUnsafe_Unmarshal-8           |    4488044 |    261 ns/op |        46 |         96 |   1.17 |        20645 |    2.72
Benchmark_XDR2_Marshal-8                      |    5162656 |    244 ns/op |        60 |         64 |   1.26 |        30975 |    3.82
Benchmark_XDR2_Unmarshal-8                    |    6587552 |    177 ns/op |        60 |         32 |   1.17 |        39525 |    5.53
Benchmark_GoAvro_Marshal-8                    |     307700 |   3622 ns/op |        47 |        872 |   1.11 |         1446 |    4.15
Benchmark_GoAvro_Unmarshal-8                  |     119338 |   9757 ns/op |        47 |       3024 |   1.16 |          560 |    3.23
Benchmark_GoAvro2Text_Marshal-8               |     292531 |   4478 ns/op |       133 |       1320 |   1.31 |         3911 |    3.39
Benchmark_GoAvro2Text_Unmarshal-8             |     293199 |   3840 ns/op |       133 |        784 |   1.13 |         3920 |    4.90
Benchmark_GoAvro2Binary_Marshal-8             |     922191 |   1449 ns/op |        47 |        464 |   1.34 |         4334 |    3.12
Benchmark_GoAvro2Binary_Unmarshal-8           |     702349 |   1572 ns/op |        47 |        544 |   1.10 |         3301 |    2.89
Benchmark_Ikea_Marshal-8                      |    1524121 |    786 ns/op |        55 |         72 |   1.20 |         8382 |   10.92
Benchmark_Ikea_Unmarshal-8                    |    1000000 |   1069 ns/op |        55 |        160 |   1.07 |         5500 |    6.68
Benchmark_ShamatonMapMsgpack_Marshal-8        |    1000000 |   1010 ns/op |        92 |        192 |   1.01 |         9200 |    5.26
Benchmark_ShamatonMapMsgpack_Unmarshal-8      |    1313325 |    909 ns/op |        92 |        136 |   1.19 |        12082 |    6.68
Benchmark_ShamatonArrayMsgpack_Marshal-8      |    1387516 |    880 ns/op |        50 |        160 |   1.22 |         6937 |    5.50
Benchmark_ShamatonArrayMsgpack_Unmarshal-8    |    1780369 |    674 ns/op |        50 |        136 |   1.20 |         8901 |    4.96
Benchmark_ShamatonMapMsgpackgen_Marshal-8     |    3733912 |    324 ns/op |        92 |         96 |   1.21 |        34351 |    3.38
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8   |    3799119 |    326 ns/op |        92 |         80 |   1.24 |        34951 |    4.08
Benchmark_ShamatonArrayMsgpackgen_Marshal-8   |    5116324 |    235 ns/op |        50 |         64 |   1.20 |        25581 |    3.67
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8 |    4815538 |    257 ns/op |        50 |         80 |   1.24 |        24077 |    3.21
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8 |     223473 |   5345 ns/op |        55 |        440 |   1.19 |         1229 |   12.15
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8 |     124766 |   9385 ns/op |        55 |       1184 |   1.17 |          686 |    7.93
Benchmark_Mum_Marshal-8                       |   10593858 |    110 ns/op |        48 |          0 |   1.17 |        50850 |    0.00
Benchmark_Mum_Unmarshal-8                     |    3715300 |    313 ns/op |        48 |         80 |   1.16 |        17833 |    3.92
Benchmark_Bebop_Marshal-8                     |    5987127 |    201 ns/op |        55 |         64 |   1.21 |        32929 |    3.15
Benchmark_Bebop_Unmarshal-8                   |    9378084 |    148 ns/op |        55 |         32 |   1.40 |        51579 |    4.65
Benchmark_FastJson_Marshal-8                  |    1185556 |   1049 ns/op |       133 |        504 |   1.24 |        15862 |    2.08
Benchmark_FastJson_Unmarshal-8                |     403963 |   3135 ns/op |       133 |       1704 |   1.27 |         5405 |    1.84
Benchmark_MusgoUnsafe_Marshal-8               |   12572750 |     90 ns/op |        48 |          0 |   1.13 |        60965 |    0.00
Benchmark_MusgoUnsafe_Unmarshal-8             |    5594883 |    232 ns/op |        48 |         64 |   1.30 |        27118 |    3.63
Benchmark_Musgo_Marshal-8                     |   12882543 |     86 ns/op |        48 |          0 |   1.12 |        62428 |    0.00
Benchmark_Musgo_Unmarshal-8                   |    3381966 |    343 ns/op |        48 |         96 |   1.16 |        16412 |    3.58


Totals:


benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MusgoUnsafe_-8                 |   18167633 |    322 ns/op |        96 |         64 |   5.86 |       176153 |    5.04
Benchmark_Bebop_-8                       |   15365211 |    350 ns/op |       110 |         96 |   5.38 |       169017 |    3.65
Benchmark_GencodeUnsafe_-8               |   12123132 |    412 ns/op |        92 |        144 |   5.00 |       111532 |    2.86
Benchmark_XDR2_-8                        |   11750208 |    421 ns/op |       120 |         96 |   4.95 |       141002 |    4.39
Benchmark_Mum_-8                         |   14309158 |    424 ns/op |        96 |         80 |   6.07 |       137367 |    5.30
Benchmark_Musgo_-8                       |   16264509 |    430 ns/op |        96 |         96 |   7.01 |       157749 |    4.49
Benchmark_ShamatonArrayMsgpackgen_-8     |    9931862 |    492 ns/op |       100 |        144 |   4.89 |        99318 |    3.42
Benchmark_GotinyNoTime_-8                |   11664379 |    496 ns/op |        96 |         96 |   5.79 |       111978 |    5.17
Benchmark_Colfer_-8                      |    9743344 |    531 ns/op |       103 |        176 |   5.18 |       100483 |    3.02
Benchmark_Gotiny_-8                      |   11224436 |    540 ns/op |        96 |        112 |   6.07 |       107754 |    4.83
Benchmark_Gogoprotobuf_-8                |    9350292 |    549 ns/op |       106 |        160 |   5.14 |        99113 |    3.44
Benchmark_Gencode_-8                     |    8469404 |    558 ns/op |       106 |        192 |   4.73 |        89775 |    2.91
Benchmark_ShamatonMapMsgpackgen_-8       |    7533031 |    650 ns/op |       184 |        176 |   4.90 |       138607 |    3.70
Benchmark_FlatBuffers_-8                 |    6887236 |    674 ns/op |       190 |        112 |   4.64 |       131043 |    6.02
Benchmark_Msgp_-8                        |    6374296 |    757 ns/op |       194 |        240 |   4.83 |       123661 |    3.15
Benchmark_Goprotobuf_-8                  |    5086323 |   1029 ns/op |       106 |        232 |   5.24 |        53915 |    4.44
Benchmark_CapNProto_-8                   |    4317300 |   1146 ns/op |       192 |        248 |   4.95 |        82892 |    4.62
Benchmark_Hprose2_-8                     |    3855947 |   1311 ns/op |       170 |        136 |   5.06 |        65740 |    9.64
Benchmark_ShamatonArrayMsgpack_-8        |    3167885 |   1554 ns/op |       100 |        296 |   4.92 |        31678 |    5.25
Benchmark_Ikea_-8                        |    2524121 |   1855 ns/op |       110 |        232 |   4.68 |        27765 |    8.00
Benchmark_ShamatonMapMsgpack_-8          |    2313325 |   1919 ns/op |       184 |        328 |   4.44 |        42565 |    5.85
Benchmark_Gob_-8                         |    2320562 |   1923 ns/op |       127 |        152 |   4.46 |        29503 |   12.65
Benchmark_Protobuf_-8                    |    2253037 |   1979 ns/op |       104 |        328 |   4.46 |        23431 |    6.04
Benchmark_CapNProto2_-8                  |    2304802 |   2083 ns/op |       192 |        548 |   4.80 |        44252 |    3.80
Benchmark_Hprose_-8                      |    2165836 |   2453 ns/op |       170 |        680 |   5.31 |        36925 |    3.61
Benchmark_GoAvro2Binary_-8               |    1624540 |   3021 ns/op |        94 |       1008 |   4.91 |        15270 |    3.00
Benchmark_Binary_-8                      |    1327517 |   3628 ns/op |       122 |        632 |   4.82 |        16195 |    5.74
Benchmark_Bson_-8                        |    1288297 |   3754 ns/op |       220 |        600 |   4.84 |        28342 |    6.26
Benchmark_VmihailencoMsgpack_-8          |    1317871 |   3934 ns/op |       200 |        800 |   5.18 |        26357 |    4.92
Benchmark_FastJson_-8                    |    1589519 |   4184 ns/op |       267 |       2208 |   6.65 |        42535 |    1.89
Benchmark_UgorjiCodecMsgpack_-8          |    1394088 |   4213 ns/op |       182 |       1800 |   5.87 |        25372 |    2.34
Benchmark_EasyJson_-8                    |    1186588 |   4347 ns/op |       303 |       1056 |   5.16 |        35989 |    4.12
Benchmark_XDR_-8                         |    1171569 |   4510 ns/op |       184 |        672 |   5.28 |        21556 |    6.71
Benchmark_UgorjiCodecBinc_-8             |     986164 |   4579 ns/op |       190 |       1976 |   4.52 |        18737 |    2.32
Benchmark_JsonIter_-8                    |    1022877 |   4843 ns/op |       282 |        512 |   4.95 |        28906 |    9.46
Benchmark_Json_-8                        |     727934 |   7561 ns/op |       303 |        591 |   5.50 |        22078 |   12.79
Benchmark_GoAvro2Text_-8                 |     585730 |   8318 ns/op |       267 |       2104 |   4.87 |        15662 |    3.95
Benchmark_Sereal_-8                      |     557244 |   8588 ns/op |       264 |       1808 |   4.79 |        14711 |    4.75
Benchmark_GoAvro_-8                      |     427038 |  13379 ns/op |        94 |       3896 |   5.71 |         4014 |    3.43
Benchmark_SSZNoTimeNoStringNoFloatA_-8   |     348239 |  14730 ns/op |       110 |       1624 |   5.13 |         3830 |    9.07



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
