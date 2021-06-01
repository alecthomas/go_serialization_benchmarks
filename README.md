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

2021-6-2 Results with Go 1.16.2 darwin/amd64 on an Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz

benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                    |   13055208 |     94 ns/op |        48 |          0 |   1.23 |        62664 |    0.00
Benchmark_Gotiny_Unmarshal-8                  |    5709070 |    210 ns/op |        47 |        112 |   1.20 |        27397 |    1.88
Benchmark_GotinyNoTime_Marshal-8              |   12246804 |    110 ns/op |        47 |          0 |   1.35 |        58772 |    0.00
Benchmark_GotinyNoTime_Unmarshal-8            |    6063762 |    210 ns/op |        47 |         96 |   1.28 |        29099 |    2.19
Benchmark_Msgp_Marshal-8                      |    8012299 |    151 ns/op |        97 |        128 |   1.21 |        77719 |    1.18
Benchmark_Msgp_Unmarshal-8                    |    4554123 |    253 ns/op |        97 |        112 |   1.15 |        44174 |    2.26
Benchmark_VmihailencoMsgpack_Marshal-8        |    1438852 |    839 ns/op |       100 |        392 |   1.21 |        14388 |    2.14
Benchmark_VmihailencoMsgpack_Unmarshal-8      |     996763 |   1249 ns/op |       100 |        408 |   1.24 |         9967 |    3.06
Benchmark_Json_Marshal-8                      |     918471 |   1568 ns/op |       148 |        208 |   1.44 |        13648 |    7.54
Benchmark_Json_Unmarshal-8                    |     398766 |   3299 ns/op |       148 |        383 |   1.32 |         5929 |    8.61
Benchmark_JsonIter_Marshal-8                  |     769125 |   1482 ns/op |       138 |        248 |   1.14 |        10636 |    5.98
Benchmark_JsonIter_Unmarshal-8                |    1000000 |   1280 ns/op |       138 |        263 |   1.28 |        13820 |    4.87
Benchmark_EasyJson_Marshal-8                  |     957908 |   1253 ns/op |       148 |        896 |   1.20 |        14244 |    1.40
Benchmark_EasyJson_Unmarshal-8                |    1000000 |   1158 ns/op |       148 |        159 |   1.16 |        14860 |    7.28
Benchmark_Bson_Marshal-8                      |    1155300 |   1099 ns/op |       110 |        376 |   1.27 |        12708 |    2.92
Benchmark_Bson_Unmarshal-8                    |     901801 |   1379 ns/op |       110 |        224 |   1.24 |         9919 |    6.16
Benchmark_Gob_Marshal-8                       |    2049348 |    652 ns/op |        63 |         40 |   1.34 |        13035 |   16.30
Benchmark_Gob_Unmarshal-8                     |    1434219 |    751 ns/op |        63 |        112 |   1.08 |         9113 |    6.71
Benchmark_XDR_Marshal-8                       |     804596 |   1540 ns/op |        88 |        376 |   1.24 |         7080 |    4.10
Benchmark_XDR_Unmarshal-8                     |    1000000 |   1155 ns/op |        88 |        216 |   1.16 |         8800 |    5.35
Benchmark_UgorjiCodecMsgpack_Marshal-8        |    1304835 |    922 ns/op |        91 |       1304 |   1.20 |        11873 |    0.71
Benchmark_UgorjiCodecMsgpack_Unmarshal-8      |    1220308 |    973 ns/op |        91 |        496 |   1.19 |        11104 |    1.96
Benchmark_UgorjiCodecBinc_Marshal-8           |    1000000 |   1000 ns/op |        95 |       1320 |   1.00 |         9500 |    0.76
Benchmark_UgorjiCodecBinc_Unmarshal-8         |    1000000 |   1143 ns/op |        95 |        656 |   1.14 |         9500 |    1.74
Benchmark_Sereal_Marshal-8                    |     564523 |   2060 ns/op |       132 |        832 |   1.16 |         7451 |    2.48
Benchmark_Sereal_Unmarshal-8                  |     503326 |   2402 ns/op |       132 |        976 |   1.21 |         6643 |    2.46
Benchmark_Binary_Marshal-8                    |    1000000 |   1090 ns/op |        61 |        312 |   1.09 |         6100 |    3.49
Benchmark_Binary_Unmarshal-8                  |    1000000 |   1107 ns/op |        61 |        320 |   1.11 |         6100 |    3.46
Benchmark_FlatBuffers_Marshal-8               |    5481980 |    212 ns/op |        95 |          0 |   1.16 |        52166 |    0.00
Benchmark_FlatBuffers_Unmarshal-8             |    5858466 |    201 ns/op |        95 |        112 |   1.18 |        55819 |    1.80
Benchmark_CapNProto_Marshal-8                 |    3527721 |    343 ns/op |        96 |         56 |   1.21 |        33866 |    6.13
Benchmark_CapNProto_Unmarshal-8               |    3433887 |    350 ns/op |        96 |        192 |   1.20 |        32965 |    1.82
Benchmark_CapNProto2_Marshal-8                |    2280529 |    528 ns/op |        96 |        244 |   1.20 |        21893 |    2.16
Benchmark_CapNProto2_Unmarshal-8              |    1878963 |    640 ns/op |        96 |        304 |   1.20 |        18038 |    2.11
Benchmark_Hprose_Marshal-8                    |    1578699 |    726 ns/op |        82 |        354 |   1.15 |        12989 |    2.05
Benchmark_Hprose_Unmarshal-8                  |    1359528 |    878 ns/op |        82 |        303 |   1.19 |        11183 |    2.90
Benchmark_Hprose2_Marshal-8                   |    2846406 |    423 ns/op |        82 |          0 |   1.20 |        23411 |    0.00
Benchmark_Hprose2_Unmarshal-8                 |    2624582 |    454 ns/op |        82 |        136 |   1.19 |        21587 |    3.34
Benchmark_Protobuf_Marshal-8                  |    2041483 |    585 ns/op |        52 |        144 |   1.19 |        10615 |    4.06
Benchmark_Protobuf_Unmarshal-8                |    2076415 |    577 ns/op |        52 |        184 |   1.20 |        10797 |    3.14
Benchmark_Goprotobuf_Marshal-8                |    5145343 |    228 ns/op |        53 |         64 |   1.18 |        27270 |    3.57
Benchmark_Goprotobuf_Unmarshal-8              |    3470637 |    344 ns/op |        53 |        168 |   1.20 |        18394 |    2.05
Benchmark_Gogoprotobuf_Marshal-8              |   11631740 |    106 ns/op |        53 |         64 |   1.24 |        61648 |    1.66
Benchmark_Gogoprotobuf_Unmarshal-8            |    7103046 |    174 ns/op |        53 |         96 |   1.24 |        37646 |    1.82
Benchmark_Colfer_Marshal-8                    |   11822221 |    100 ns/op |        51 |         64 |   1.19 |        60411 |    1.57
Benchmark_Colfer_Unmarshal-8                  |    8095731 |    151 ns/op |        52 |        112 |   1.23 |        42097 |    1.36
Benchmark_Gencode_Marshal-8                   |    9361725 |    133 ns/op |        53 |         80 |   1.25 |        49617 |    1.67
Benchmark_Gencode_Unmarshal-8                 |    7527206 |    189 ns/op |        53 |        112 |   1.43 |        39894 |    1.70
Benchmark_GencodeUnsafe_Marshal-8             |   15368806 |     98 ns/op |        46 |         48 |   1.51 |        70696 |    2.05
Benchmark_GencodeUnsafe_Unmarshal-8           |    9819229 |    154 ns/op |        46 |         96 |   1.52 |        45168 |    1.61
Benchmark_XDR2_Marshal-8                      |    9061377 |    150 ns/op |        60 |         64 |   1.36 |        54368 |    2.35
Benchmark_XDR2_Unmarshal-8                    |   11242050 |    135 ns/op |        60 |         32 |   1.53 |        67452 |    4.24
Benchmark_GoAvro_Marshal-8                    |     553201 |   2644 ns/op |        47 |        872 |   1.46 |         2600 |    3.03
Benchmark_GoAvro_Unmarshal-8                  |     188358 |   6233 ns/op |        47 |       3024 |   1.17 |          885 |    2.06
Benchmark_GoAvro2Text_Marshal-8               |     504949 |   2319 ns/op |       133 |       1320 |   1.17 |         6751 |    1.76
Benchmark_GoAvro2Text_Unmarshal-8             |     558674 |   2324 ns/op |       133 |        783 |   1.30 |         7475 |    2.97
Benchmark_GoAvro2Binary_Marshal-8             |    1691079 |    697 ns/op |        47 |        464 |   1.18 |         7948 |    1.50
Benchmark_GoAvro2Binary_Unmarshal-8           |    1534810 |    774 ns/op |        47 |        544 |   1.19 |         7213 |    1.42
Benchmark_Ikea_Marshal-8                      |    2366568 |    509 ns/op |        55 |         72 |   1.21 |        13016 |    7.08
Benchmark_Ikea_Unmarshal-8                    |    1900034 |    643 ns/op |        55 |        160 |   1.22 |        10450 |    4.02
Benchmark_ShamatonMapMsgpack_Marshal-8        |    1572403 |    705 ns/op |        92 |        192 |   1.11 |        14466 |    3.67
Benchmark_ShamatonMapMsgpack_Unmarshal-8      |    1735891 |    666 ns/op |        92 |        136 |   1.16 |        15970 |    4.90
Benchmark_ShamatonArrayMsgpack_Marshal-8      |    2050185 |    702 ns/op |        50 |        160 |   1.44 |        10250 |    4.39
Benchmark_ShamatonArrayMsgpack_Unmarshal-8    |    2914256 |    408 ns/op |        50 |        136 |   1.19 |        14571 |    3.00
Benchmark_ShamatonMapMsgpackgen_Marshal-8     |    6816490 |    174 ns/op |        92 |         96 |   1.19 |        62711 |    1.82
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8   |    5988542 |    183 ns/op |        92 |         80 |   1.10 |        55094 |    2.30
Benchmark_ShamatonArrayMsgpackgen_Marshal-8   |    9056970 |    135 ns/op |        50 |         64 |   1.23 |        45284 |    2.12
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8 |    9500049 |    129 ns/op |        50 |         80 |   1.23 |        47500 |    1.62
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8 |     334351 |   3701 ns/op |        55 |        440 |   1.24 |         1838 |    8.41
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8 |     206965 |   6014 ns/op |        55 |       1184 |   1.24 |         1138 |    5.08
Benchmark_Mum_Marshal-8                       |   14044489 |     85 ns/op |        48 |          0 |   1.21 |        67413 |    0.00
Benchmark_Mum_Unmarshal-8                     |    6952032 |    172 ns/op |        48 |         80 |   1.20 |        33369 |    2.16
Benchmark_Bebop_Marshal-8                     |   11255248 |    105 ns/op |        55 |         64 |   1.19 |        61903 |    1.65
Benchmark_Bebop_Unmarshal-8                   |   12815527 |     84 ns/op |        55 |         32 |   1.09 |        70485 |    2.65
Benchmark_FastJson_Marshal-8                  |    2340780 |    505 ns/op |       133 |        504 |   1.18 |        31319 |    1.00
Benchmark_FastJson_Unmarshal-8                |     770043 |   1537 ns/op |       133 |       1704 |   1.18 |        10303 |    0.90


Totals:


benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Bebop_-8                       |   24070775 |    190 ns/op |       110 |         96 |   4.58 |       264778 |    1.98
Benchmark_Colfer_-8                      |   19917952 |    252 ns/op |       103 |        176 |   5.03 |       205354 |    1.43
Benchmark_GencodeUnsafe_-8               |   25188035 |    253 ns/op |        92 |        144 |   6.37 |       231729 |    1.76
Benchmark_Mum_-8                         |   20996521 |    258 ns/op |        96 |         80 |   5.43 |       201566 |    3.23
Benchmark_ShamatonArrayMsgpackgen_-8     |   18557019 |    265 ns/op |       100 |        144 |   4.92 |       185570 |    1.84
Benchmark_Gogoprotobuf_-8                |   18734786 |    281 ns/op |       106 |        160 |   5.27 |       198588 |    1.76
Benchmark_XDR2_-8                        |   20303427 |    286 ns/op |       120 |         96 |   5.81 |       243641 |    2.98
Benchmark_Gotiny_-8                      |   18764278 |    304 ns/op |        95 |        112 |   5.72 |       180118 |    2.72
Benchmark_GotinyNoTime_-8                |   18310566 |    321 ns/op |        95 |         96 |   5.88 |       175744 |    3.34
Benchmark_Gencode_-8                     |   16888931 |    323 ns/op |       106 |        192 |   5.46 |       179022 |    1.68
Benchmark_ShamatonMapMsgpackgen_-8       |   12805032 |    358 ns/op |       184 |        176 |   4.59 |       235612 |    2.04
Benchmark_Msgp_-8                        |   12566422 |    404 ns/op |       194 |        240 |   5.08 |       243788 |    1.69
Benchmark_FlatBuffers_-8                 |   11340446 |    413 ns/op |       190 |        112 |   4.69 |       215967 |    3.70
Benchmark_Goprotobuf_-8                  |    8615980 |    573 ns/op |       106 |        232 |   4.94 |        91329 |    2.47
Benchmark_CapNProto_-8                   |    6961608 |    693 ns/op |       192 |        248 |   4.83 |       133662 |    2.80
Benchmark_Hprose2_-8                     |    5470988 |    877 ns/op |       164 |        136 |   4.80 |        89997 |    6.45
Benchmark_ShamatonArrayMsgpack_-8        |    4964441 |   1110 ns/op |       100 |        296 |   5.51 |        49644 |    3.75
Benchmark_Ikea_-8                        |    4266602 |   1153 ns/op |       110 |        232 |   4.92 |        46932 |    4.97
Benchmark_Protobuf_-8                    |    4117898 |   1162 ns/op |       104 |        328 |   4.79 |        42826 |    3.54
Benchmark_CapNProto2_-8                  |    4159492 |   1168 ns/op |       192 |        548 |   4.86 |        79862 |    2.13
Benchmark_ShamatonMapMsgpack_-8          |    3308294 |   1371 ns/op |       184 |        328 |   4.54 |        60872 |    4.18
Benchmark_Gob_-8                         |    3483567 |   1403 ns/op |       127 |        152 |   4.89 |        44293 |    9.24
Benchmark_GoAvro2Binary_-8               |    3225889 |   1472 ns/op |        94 |       1008 |   4.75 |        30323 |    1.46
Benchmark_Hprose_-8                      |    2938227 |   1605 ns/op |       164 |        657 |   4.72 |        48345 |    2.44
Benchmark_UgorjiCodecMsgpack_-8          |    2525143 |   1896 ns/op |       182 |       1800 |   4.79 |        45957 |    1.05
Benchmark_FastJson_-8                    |    3110823 |   2042 ns/op |       267 |       2208 |   6.35 |        83245 |    0.93
Benchmark_VmihailencoMsgpack_-8          |    2435615 |   2088 ns/op |       200 |        800 |   5.09 |        48712 |    2.61
Benchmark_UgorjiCodecBinc_-8             |    2000000 |   2143 ns/op |       190 |       1976 |   4.29 |        38000 |    1.08
Benchmark_Binary_-8                      |    2000000 |   2197 ns/op |       122 |        632 |   4.39 |        24400 |    3.48
Benchmark_EasyJson_-8                    |    1957908 |   2411 ns/op |       297 |       1055 |   4.72 |        58208 |    2.29
Benchmark_Bson_-8                        |    2057101 |   2478 ns/op |       220 |        600 |   5.10 |        45256 |    4.13
Benchmark_XDR_-8                         |    1804596 |   2695 ns/op |       176 |        592 |   4.86 |        31760 |    4.55
Benchmark_JsonIter_-8                    |    1769125 |   2762 ns/op |       276 |        511 |   4.89 |        48916 |    5.41
Benchmark_Sereal_-8                      |    1067849 |   4462 ns/op |       264 |       1808 |   4.76 |        28191 |    2.47
Benchmark_GoAvro2Text_-8                 |    1063623 |   4643 ns/op |       267 |       2103 |   4.94 |        28451 |    2.21
Benchmark_Json_-8                        |    1317237 |   4867 ns/op |       297 |        591 |   6.41 |        39161 |    8.24
Benchmark_GoAvro_-8                      |     741559 |   8877 ns/op |        94 |       3896 |   6.58 |         6970 |    2.28
Benchmark_SSZNoTimeNoStringNoFloatA_-8   |     541316 |   9715 ns/op |       110 |       1624 |   5.26 |         5954 |    5.98



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
