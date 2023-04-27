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
Benchmark_Gotiny_Marshal-16                      |    2649405 |    500 ns/op |        48 |        168 |   1.33 |        12717 |    2.98
Benchmark_Gotiny_Unmarshal-16                    |    4602387 |    267 ns/op |        47 |        112 |   1.23 |        22086 |    2.39
Benchmark_GotinyNoTime_Marshal-16                |    2324325 |    532 ns/op |        47 |        168 |   1.24 |        11154 |    3.17
Benchmark_GotinyNoTime_Unmarshal-16              |    4299528 |    271 ns/op |        47 |         96 |   1.17 |        20633 |    2.82
Benchmark_Msgp_Marshal-16                        |    6384872 |    179 ns/op |        97 |        128 |   1.15 |        61933 |    1.40
Benchmark_Msgp_Unmarshal-16                      |    3399418 |    317 ns/op |        97 |        112 |   1.08 |        32974 |    2.84
Benchmark_VmihailencoMsgpack_Marshal-16          |     962656 |   1185 ns/op |        92 |        264 |   1.14 |         8856 |    4.49
Benchmark_VmihailencoMsgpack_Unmarshal-16        |    1262151 |   1031 ns/op |        92 |        160 |   1.30 |        11611 |    6.44
Benchmark_Json_Marshal-16                        |     825260 |   1765 ns/op |       151 |        208 |   1.46 |        12519 |    8.49
Benchmark_Json_Unmarshal-16                      |     406338 |   3296 ns/op |       151 |        351 |   1.34 |         6164 |    9.39
Benchmark_JsonIter_Marshal-16                    |    1000000 |   1402 ns/op |       141 |        200 |   1.40 |        14120 |    7.01
Benchmark_JsonIter_Unmarshal-16                  |     870942 |   1586 ns/op |       141 |        216 |   1.38 |        12306 |    7.34
Benchmark_EasyJson_Marshal-16                    |     677398 |   2023 ns/op |       151 |        896 |   1.37 |        10269 |    2.26
Benchmark_EasyJson_Unmarshal-16                  |    1193428 |   1035 ns/op |       151 |        112 |   1.24 |        18092 |    9.24
Benchmark_Bson_Marshal-16                        |     736555 |   1926 ns/op |       110 |        376 |   1.42 |         8102 |    5.12
Benchmark_Bson_Unmarshal-16                      |     810910 |   1475 ns/op |       110 |        224 |   1.20 |         8920 |    6.58
Benchmark_MongoBson_Marshal-16                   |     698608 |   1966 ns/op |       110 |        240 |   1.37 |         7684 |    8.19
Benchmark_MongoBson_Unmarshal-16                 |     549484 |   2015 ns/op |       110 |        408 |   1.11 |         6044 |    4.94
Benchmark_Gob_Marshal-16                         |     167848 |   6530 ns/op |       163 |       1616 |   1.10 |         2745 |    4.04
Benchmark_Gob_Unmarshal-16                       |      43647 |  35315 ns/op |       163 |       7768 |   1.54 |          714 |    4.55
Benchmark_XDR_Marshal-16                         |     584995 |   1932 ns/op |        91 |        439 |   1.13 |         5381 |    4.40
Benchmark_XDR_Unmarshal-16                       |     961742 |   1567 ns/op |        92 |        231 |   1.51 |         8848 |    6.78
Benchmark_UgorjiCodecMsgpack_Marshal-16          |    1000000 |   1583 ns/op |        91 |       1240 |   1.58 |         9100 |    1.28
Benchmark_UgorjiCodecMsgpack_Unmarshal-16        |     596980 |   2041 ns/op |        91 |        688 |   1.22 |         5432 |    2.97
Benchmark_UgorjiCodecBinc_Marshal-16             |     644101 |   1953 ns/op |        95 |       1256 |   1.26 |         6118 |    1.55
Benchmark_UgorjiCodecBinc_Unmarshal-16           |     864344 |   1309 ns/op |        95 |        688 |   1.13 |         8211 |    1.90
Benchmark_Sereal_Marshal-16                      |     495840 |   3072 ns/op |       132 |        832 |   1.52 |         6545 |    3.69
Benchmark_Sereal_Unmarshal-16                    |     285642 |   3996 ns/op |       132 |        976 |   1.14 |         3770 |    4.09
Benchmark_Binary_Marshal-16                      |     628082 |   1931 ns/op |        61 |        360 |   1.21 |         3831 |    5.36
Benchmark_Binary_Unmarshal-16                    |     755530 |   1770 ns/op |        61 |        320 |   1.34 |         4608 |    5.53
Benchmark_FlatBuffers_Marshal-16                 |    1000000 |   1009 ns/op |        95 |        376 |   1.01 |         9512 |    2.68
Benchmark_FlatBuffers_Unmarshal-16               |    4943922 |    299 ns/op |        95 |        112 |   1.48 |        47021 |    2.68
Benchmark_CapNProto_Marshal-16                   |     720890 |   3100 ns/op |        96 |       4392 |   2.23 |         6920 |    0.71
Benchmark_CapNProto_Unmarshal-16                 |    1215064 |    930 ns/op |        96 |        192 |   1.13 |        11664 |    4.84
Benchmark_CapNProto2_Marshal-16                  |     481316 |   2147 ns/op |        96 |       1452 |   1.03 |         4620 |    1.48
Benchmark_CapNProto2_Unmarshal-16                |    1991179 |    694 ns/op |        96 |        272 |   1.38 |        19115 |    2.55
Benchmark_Hprose_Marshal-16                      |    1761524 |    601 ns/op |        85 |        448 |   1.06 |        15015 |    1.34
Benchmark_Hprose_Unmarshal-16                    |    1000000 |   1603 ns/op |        85 |        304 |   1.60 |         8525 |    5.27
Benchmark_Hprose2_Marshal-16                     |    4057759 |    270 ns/op |        85 |          0 |   1.10 |        34612 |    0.00
Benchmark_Hprose2_Unmarshal-16                   |    1785067 |    692 ns/op |        85 |        136 |   1.24 |        15217 |    5.09
Benchmark_Protobuf_Marshal-16                    |    1720783 |    787 ns/op |        52 |        144 |   1.35 |         8948 |    5.47
Benchmark_Protobuf_Unmarshal-16                  |     939722 |   1209 ns/op |        52 |        184 |   1.14 |         4886 |    6.57
Benchmark_Pulsar_Marshal-16                      |    1332962 |    846 ns/op |        51 |        304 |   1.13 |         6878 |    2.78
Benchmark_Pulsar_Unmarshal-16                    |    1746132 |    724 ns/op |        51 |        256 |   1.26 |         9013 |    2.83
Benchmark_Gogoprotobuf_Marshal-16                |    8852894 |    169 ns/op |        53 |         64 |   1.50 |        46920 |    2.65
Benchmark_Gogoprotobuf_Unmarshal-16              |    4653692 |    247 ns/op |        53 |         96 |   1.15 |        24664 |    2.57
Benchmark_Gogojsonpb_Marshal-16                  |      88912 |  13517 ns/op |       125 |       3096 |   1.20 |         1117 |    4.37
Benchmark_Gogojsonpb_Unmarshal-16                |      76432 |  20184 ns/op |       125 |       3381 |   1.54 |          962 |    5.97
Benchmark_Colfer_Marshal-16                      |    9113589 |    132 ns/op |        51 |         64 |   1.21 |        46616 |    2.07
Benchmark_Colfer_Unmarshal-16                    |    6470188 |    212 ns/op |        51 |        112 |   1.37 |        32997 |    1.89
Benchmark_Gencode_Marshal-16                     |    5744428 |    192 ns/op |        53 |         80 |   1.10 |        30445 |    2.40
Benchmark_Gencode_Unmarshal-16                   |    6495428 |    202 ns/op |        53 |        112 |   1.31 |        34425 |    1.81
Benchmark_GencodeUnsafe_Marshal-16               |   10857216 |    118 ns/op |        46 |         48 |   1.29 |        49943 |    2.47
Benchmark_GencodeUnsafe_Unmarshal-16             |    6532112 |    175 ns/op |        46 |         96 |   1.15 |        30047 |    1.83
Benchmark_XDR2_Marshal-16                        |    4959430 |    241 ns/op |        60 |         64 |   1.20 |        29756 |    3.77
Benchmark_XDR2_Unmarshal-16                      |    7523264 |    142 ns/op |        60 |         32 |   1.07 |        45139 |    4.44
Benchmark_GoAvro_Marshal-16                      |     639590 |   2620 ns/op |        47 |        728 |   1.68 |         3006 |    3.60
Benchmark_GoAvro_Unmarshal-16                    |     192807 |   5900 ns/op |        47 |       2544 |   1.14 |          906 |    2.32
Benchmark_GoAvro2Text_Marshal-16                 |     354640 |   3321 ns/op |       133 |       1320 |   1.18 |         4745 |    2.52
Benchmark_GoAvro2Text_Unmarshal-16               |     350242 |   3778 ns/op |       133 |        736 |   1.32 |         4686 |    5.13
Benchmark_GoAvro2Binary_Marshal-16               |    1105582 |   1251 ns/op |        47 |        464 |   1.38 |         5196 |    2.70
Benchmark_GoAvro2Binary_Unmarshal-16             |     784065 |   1785 ns/op |        47 |        544 |   1.40 |         3685 |    3.28
Benchmark_Ikea_Marshal-16                        |    1063552 |   1178 ns/op |        55 |        184 |   1.25 |         5849 |    6.40
Benchmark_Ikea_Unmarshal-16                      |    1000000 |   1046 ns/op |        55 |        160 |   1.05 |         5500 |    6.54
Benchmark_ShamatonMapMsgpack_Marshal-16          |     912945 |   1170 ns/op |        92 |        192 |   1.07 |         8399 |    6.09
Benchmark_ShamatonMapMsgpack_Unmarshal-16        |    1120543 |   1076 ns/op |        92 |        168 |   1.21 |        10308 |    6.40
Benchmark_ShamatonArrayMsgpack_Marshal-16        |    1744819 |    645 ns/op |        50 |        160 |   1.13 |         8724 |    4.04
Benchmark_ShamatonArrayMsgpack_Unmarshal-16      |    2083960 |    622 ns/op |        50 |        168 |   1.30 |        10419 |    3.71
Benchmark_ShamatonMapMsgpackgen_Marshal-16       |    6039444 |    230 ns/op |        92 |         96 |   1.39 |        55562 |    2.40
Benchmark_ShamatonMapMsgpackgen_Unmarshal-16     |    2745090 |    426 ns/op |        92 |        112 |   1.17 |        25254 |    3.81
Benchmark_ShamatonArrayMsgpackgen_Marshal-16     |    7060170 |    177 ns/op |        50 |         64 |   1.25 |        35300 |    2.77
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-16   |    5035602 |    263 ns/op |        50 |        112 |   1.33 |        25178 |    2.35
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-16   |     235539 |   4604 ns/op |        55 |        440 |   1.08 |         1295 |   10.46
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-16 |     114822 |  10033 ns/op |        55 |       1184 |   1.15 |          631 |    8.47
Benchmark_Bebop_Marshal-16                       |    6879166 |    146 ns/op |        55 |         64 |   1.01 |        37835 |    2.29
Benchmark_Bebop_Unmarshal-16                     |   10850620 |    117 ns/op |        55 |         32 |   1.27 |        59678 |    3.66
Benchmark_FastJson_Marshal-16                    |    1287648 |   1119 ns/op |       133 |        504 |   1.44 |        17228 |    2.22
Benchmark_FastJson_Unmarshal-16                  |     411573 |   2536 ns/op |       133 |       1704 |   1.04 |         5506 |    1.49


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Bebop_-16                     |   17729786 |    263 ns/op |       110 |         96 |   4.67 |       195027 |    2.74
Benchmark_GencodeUnsafe_-16             |   17389328 |    294 ns/op |        92 |        144 |   5.11 |       159981 |    2.04
Benchmark_Colfer_-16                    |   15583777 |    344 ns/op |       102 |        176 |   5.37 |       159188 |    1.96
Benchmark_XDR2_-16                      |   12482694 |    383 ns/op |       120 |         96 |   4.79 |       149792 |    3.99
Benchmark_Gencode_-16                   |   12239856 |    394 ns/op |       106 |        192 |   4.83 |       129742 |    2.06
Benchmark_Gogoprotobuf_-16              |   13506586 |    416 ns/op |       106 |        160 |   5.62 |       143169 |    2.60
Benchmark_ShamatonArrayMsgpackgen_-16   |   12095772 |    440 ns/op |       100 |        176 |   5.33 |       120957 |    2.50
Benchmark_Msgp_-16                      |    9784290 |    497 ns/op |       194 |        240 |   4.87 |       189815 |    2.07
Benchmark_ShamatonMapMsgpackgen_-16     |    8784534 |    656 ns/op |       184 |        208 |   5.77 |       161635 |    3.16
Benchmark_Gotiny_-16                    |    7251792 |    768 ns/op |        95 |        280 |   5.57 |        69609 |    2.74
Benchmark_GotinyNoTime_-16              |    6623853 |    803 ns/op |        95 |        264 |   5.32 |        63575 |    3.04
Benchmark_Hprose2_-16                   |    5842826 |    963 ns/op |       170 |        136 |   5.63 |        99649 |    7.09
Benchmark_ShamatonArrayMsgpack_-16      |    3828779 |   1268 ns/op |       100 |        328 |   4.86 |        38287 |    3.87
Benchmark_FlatBuffers_-16               |    5943922 |   1308 ns/op |       190 |        488 |   7.78 |       113071 |    2.68
Benchmark_Pulsar_-16                    |    3079094 |   1570 ns/op |       103 |        560 |   4.83 |        31782 |    2.80
Benchmark_Protobuf_-16                  |    2660505 |   1996 ns/op |       104 |        328 |   5.31 |        27669 |    6.09
Benchmark_Hprose_-16                    |    2761524 |   2204 ns/op |       170 |        752 |   6.09 |        47081 |    2.93
Benchmark_VmihailencoMsgpack_-16        |    2224807 |   2216 ns/op |       184 |        424 |   4.93 |        40936 |    5.23
Benchmark_Ikea_-16                      |    2063552 |   2224 ns/op |       110 |        344 |   4.59 |        22699 |    6.47
Benchmark_ShamatonMapMsgpack_-16        |    2033488 |   2246 ns/op |       184 |        360 |   4.57 |        37416 |    6.24
Benchmark_CapNProto2_-16                |    2472495 |   2841 ns/op |       192 |       1724 |   7.02 |        47471 |    1.65
Benchmark_JsonIter_-16                  |    1870942 |   2988 ns/op |       282 |        416 |   5.59 |        52854 |    7.18
Benchmark_GoAvro2Binary_-16             |    1889647 |   3036 ns/op |        94 |       1008 |   5.74 |        17762 |    3.01
Benchmark_EasyJson_-16                  |    1870826 |   3058 ns/op |       303 |       1008 |   5.72 |        56723 |    3.03
Benchmark_UgorjiCodecBinc_-16           |    1508445 |   3262 ns/op |       190 |       1944 |   4.92 |        28660 |    1.68
Benchmark_Bson_-16                      |    1547465 |   3401 ns/op |       220 |        600 |   5.26 |        34044 |    5.67
Benchmark_XDR_-16                       |    1546737 |   3499 ns/op |       183 |        670 |   5.41 |        28458 |    5.22
Benchmark_UgorjiCodecMsgpack_-16        |    1596980 |   3624 ns/op |       182 |       1928 |   5.79 |        29065 |    1.88
Benchmark_FastJson_-16                  |    1699221 |   3655 ns/op |       267 |       2208 |   6.21 |        45471 |    1.66
Benchmark_Binary_-16                    |    1383612 |   3701 ns/op |       122 |        680 |   5.12 |        16880 |    5.44
Benchmark_MongoBson_-16                 |    1248092 |   3981 ns/op |       220 |        648 |   4.97 |        27458 |    6.14
Benchmark_CapNProto_-16                 |    1935954 |   4030 ns/op |       192 |       4584 |   7.80 |        37170 |    0.88
Benchmark_Json_-16                      |    1231598 |   5061 ns/op |       303 |        559 |   6.23 |        37366 |    9.05
Benchmark_Sereal_-16                    |     781482 |   7068 ns/op |       264 |       1808 |   5.52 |        20631 |    3.91
Benchmark_GoAvro2Text_-16               |     704882 |   7099 ns/op |       267 |       2056 |   5.00 |        18862 |    3.45
Benchmark_GoAvro_-16                    |     832397 |   8520 ns/op |        94 |       3272 |   7.09 |         7824 |    2.60
Benchmark_SSZNoTimeNoStringNoFloatA_-16 |     350361 |  14637 ns/op |       110 |       1624 |   5.13 |         3853 |    9.01
Benchmark_Gogojsonpb_-16                |     165344 |  33701 ns/op |       251 |       6477 |   5.57 |         4160 |    5.20
Benchmark_Gob_-16                       |     211495 |  41845 ns/op |       327 |       9384 |   8.85 |         6920 |    4.46

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
