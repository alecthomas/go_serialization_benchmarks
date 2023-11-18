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
- [github.com/deneonet/benc](https://github.com/deneonet/benc)

## Running the benchmarks

To make it fair I use [perflock](https://github.com/aclements/perflock) to lock the CPU usage at 70%

```bash
go get -u -t
perflock -governor 70% go test -bench='.*' ./ -count=1 > "results.txt"
```

To update the table in the README:

```bash
./stats.sh
```

## Recommendation

If correctness and interoperability are the most
important factors, [gob](http://golang.org/pkg/encoding/gob/) and [json](http://golang.org/pkg/encoding/json/) do the job good.
If speed matters, [MUS](https://github.com/mus-format/mus-go) and [BENC](https://github.com/deneonet/benc) are probably the best choice.

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

Command used:
```bash
perflock -governor 70% go test -bench='.*' ./ -count=1 > "results.txt"
```
2023-11-16 Results with Go 1.21.4 linux/amd64 with processor `11th Gen Intel(R) Core(TM) i5-11300H @ 3.10GHz`

benchmark                                                       | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                                      |    3714318 |    296 ns/op |        48 |        168 |   1.10 |        17828 |    1.76
Benchmark_Gotiny_Unmarshal-8                                    |    7337814 |    162 ns/op |        47 |        112 |   1.19 |        35214 |    1.45
Benchmark_GotinyNoTime_Marshal-8                                |    3779130 |    372 ns/op |        47 |        168 |   1.41 |        18136 |    2.21
Benchmark_GotinyNoTime_Unmarshal-8                              |    7829782 |    147 ns/op |        47 |         96 |   1.16 |        37575 |    1.54
Benchmark_Msgp_Marshal-8                                        |   10978256 |    105 ns/op |        97 |        128 |   1.16 |       106489 |    0.82
Benchmark_Msgp_Unmarshal-8                                      |    6372080 |    198 ns/op |        97 |        112 |   1.26 |        61809 |    1.77
Benchmark_VmihailencoMsgpack_Marshal-8                          |    1918615 |    626 ns/op |        92 |        264 |   1.20 |        17651 |    2.37
Benchmark_VmihailencoMsgpack_Unmarshal-8                        |    1406083 |    922 ns/op |        92 |        160 |   1.30 |        12935 |    5.77
Benchmark_Json_Marshal-8                                        |    1000000 |   1038 ns/op |       151 |        208 |   1.04 |        15160 |    4.99
Benchmark_Json_Unmarshal-8                                      |     471050 |   2366 ns/op |       151 |        351 |   1.11 |         7145 |    6.74
Benchmark_JsonIter_Marshal-8                                    |    1656364 |    693 ns/op |       141 |        200 |   1.15 |        23404 |    3.47
Benchmark_JsonIter_Unmarshal-8                                  |    1275422 |    889 ns/op |       141 |        216 |   1.13 |        18021 |    4.12
Benchmark_EasyJson_Marshal-8                                    |    1501216 |    756 ns/op |       151 |        896 |   1.14 |        22773 |    0.84
Benchmark_EasyJson_Unmarshal-8                                  |    1629051 |    719 ns/op |       151 |        112 |   1.17 |        24712 |    6.42
Benchmark_Bson_Marshal-8                                        |    1432402 |    759 ns/op |       110 |        376 |   1.09 |        15756 |    2.02
Benchmark_Bson_Unmarshal-8                                      |    1000000 |   1163 ns/op |       110 |        224 |   1.16 |        11000 |    5.19
Benchmark_MongoBson_Marshal-8                                   |    1000000 |   1121 ns/op |       110 |        240 |   1.12 |        11000 |    4.67
Benchmark_MongoBson_Unmarshal-8                                 |     936487 |   1191 ns/op |       110 |        408 |   1.12 |        10301 |    2.92
Benchmark_Gob_Marshal-8                                         |     332130 |   3677 ns/op |       162 |       1744 |   1.22 |         5397 |    2.11
Benchmark_Gob_Unmarshal-8                                       |      66332 |  17594 ns/op |       162 |       7720 |   1.17 |         1078 |    2.28
Benchmark_XDR_Marshal-8                                         |    1000000 |   1103 ns/op |        91 |        439 |   1.10 |         9199 |    2.51
Benchmark_XDR_Unmarshal-8                                       |    1373052 |    833 ns/op |        92 |        232 |   1.14 |        12632 |    3.59
Benchmark_UgorjiCodecMsgpack_Marshal-8                          |    1640578 |    693 ns/op |        91 |       1240 |   1.14 |        14929 |    0.56
Benchmark_UgorjiCodecMsgpack_Unmarshal-8                        |    1498993 |    756 ns/op |        91 |        688 |   1.13 |        13640 |    1.10
Benchmark_UgorjiCodecBinc_Marshal-8                             |    1472635 |    760 ns/op |        95 |       1256 |   1.12 |        13990 |    0.61
Benchmark_UgorjiCodecBinc_Unmarshal-8                           |    1527250 |    720 ns/op |        95 |        688 |   1.10 |        14508 |    1.05
Benchmark_Sereal_Marshal-8                                      |     659978 |   1787 ns/op |       132 |        832 |   1.18 |         8711 |    2.15
Benchmark_Sereal_Unmarshal-8                                    |     563316 |   2230 ns/op |       132 |        976 |   1.26 |         7435 |    2.28
Benchmark_Binary_Marshal-8                                      |     896958 |   1237 ns/op |        61 |        360 |   1.11 |         5471 |    3.44
Benchmark_Binary_Unmarshal-8                                    |    1000000 |   1085 ns/op |        61 |        320 |   1.08 |         6100 |    3.39
Benchmark_FlatBuffers_Marshal-8                                 |    2095118 |    529 ns/op |        95 |        376 |   1.11 |        19951 |    1.41
Benchmark_FlatBuffers_Unmarshal-8                               |    7173885 |    168 ns/op |        95 |        112 |   1.21 |        68209 |    1.51
Benchmark_CapNProto_Marshal-8                                   |     862030 |   1414 ns/op |        96 |       4392 |   1.22 |         8275 |    0.32
Benchmark_CapNProto_Unmarshal-8                                 |    3344185 |    339 ns/op |        96 |        192 |   1.14 |        32104 |    1.77
Benchmark_CapNProto2_Marshal-8                                  |    1396081 |    823 ns/op |        96 |       1452 |   1.15 |        13402 |    0.57
Benchmark_CapNProto2_Unmarshal-8                                |    3327934 |    342 ns/op |        96 |        272 |   1.14 |        31948 |    1.26
Benchmark_Hprose_Marshal-8                                      |    1664455 |    728 ns/op |        85 |        466 |   1.21 |        14189 |    1.56
Benchmark_Hprose_Unmarshal-8                                    |    1401349 |    975 ns/op |        85 |        304 |   1.37 |        11949 |    3.21
Benchmark_Hprose2_Marshal-8                                     |    2377542 |    431 ns/op |        85 |          0 |   1.03 |        20268 |    0.00
Benchmark_Hprose2_Unmarshal-8                                   |    2290743 |    519 ns/op |        85 |        136 |   1.19 |        19528 |    3.82
Benchmark_Protobuf_Marshal-8                                    |    1668072 |    653 ns/op |        52 |        144 |   1.09 |         8673 |    4.54
Benchmark_Protobuf_Unmarshal-8                                  |    1751280 |    847 ns/op |        52 |        184 |   1.48 |         9106 |    4.60
Benchmark_Pulsar_Marshal-8                                      |    2235410 |    478 ns/op |        51 |        304 |   1.07 |        11521 |    1.57
Benchmark_Pulsar_Unmarshal-8                                    |    2334020 |    503 ns/op |        51 |        256 |   1.17 |        12041 |    1.97
Benchmark_Gogoprotobuf_Marshal-8                                |   12454290 |    112 ns/op |        53 |         64 |   1.40 |        66007 |    1.76
Benchmark_Gogoprotobuf_Unmarshal-8                              |    6689876 |    189 ns/op |        53 |         96 |   1.26 |        35456 |    1.97
Benchmark_Gogojsonpb_Marshal-8                                  |     134907 |   9816 ns/op |       125 |       3083 |   1.32 |         1690 |    3.18
Benchmark_Gogojsonpb_Unmarshal-8                                |      80768 |  14565 ns/op |       125 |       3376 |   1.18 |         1014 |    4.31
Benchmark_Colfer_Marshal-8                                      |    8374770 |    122 ns/op |        51 |         64 |   1.03 |        42820 |    1.92
Benchmark_Colfer_Unmarshal-8                                    |    6232246 |    182 ns/op |        52 |        112 |   1.13 |        32407 |    1.62
Benchmark_Gencode_Marshal-8                                     |   10375894 |    114 ns/op |        53 |         80 |   1.19 |        54992 |    1.43
Benchmark_Gencode_Unmarshal-8                                   |    8154621 |    150 ns/op |        53 |        112 |   1.23 |        43219 |    1.34
Benchmark_GencodeUnsafe_Marshal-8                               |   16499416 |     71 ns/op |        46 |         48 |   1.18 |        75897 |    1.48
Benchmark_GencodeUnsafe_Unmarshal-8                             |    9791605 |    148 ns/op |        46 |         96 |   1.46 |        45041 |    1.55
Benchmark_XDR2_Marshal-8                                        |    8810739 |    148 ns/op |        60 |         64 |   1.31 |        52864 |    2.32
Benchmark_XDR2_Unmarshal-8                                      |   11508285 |    106 ns/op |        60 |         32 |   1.23 |        69049 |    3.34
Benchmark_GoAvro_Marshal-8                                      |     690482 |   1800 ns/op |        47 |        584 |   1.24 |         3245 |    3.08
Benchmark_GoAvro_Unmarshal-8                                    |     280483 |   4021 ns/op |        47 |       2312 |   1.13 |         1318 |    1.74
Benchmark_GoAvro2Text_Marshal-8                                 |     425636 |   2387 ns/op |       133 |       1320 |   1.02 |         5695 |    1.81
Benchmark_GoAvro2Text_Unmarshal-8                               |     573084 |   2127 ns/op |       133 |        736 |   1.22 |         7667 |    2.89
Benchmark_GoAvro2Binary_Marshal-8                               |    1691883 |    808 ns/op |        47 |        464 |   1.37 |         7951 |    1.74
Benchmark_GoAvro2Binary_Unmarshal-8                             |    1000000 |   1046 ns/op |        47 |        544 |   1.05 |         4700 |    1.92
Benchmark_Ikea_Marshal-8                                        |    1419968 |    914 ns/op |        55 |        184 |   1.30 |         7809 |    4.97
Benchmark_Ikea_Unmarshal-8                                      |    2206431 |    548 ns/op |        55 |        160 |   1.21 |        12135 |    3.43
Benchmark_ShamatonMapMsgpack_Marshal-8                          |    2266911 |    491 ns/op |        92 |        192 |   1.11 |        20855 |    2.56
Benchmark_ShamatonMapMsgpack_Unmarshal-8                        |    2309575 |    490 ns/op |        92 |        168 |   1.13 |        21248 |    2.92
Benchmark_ShamatonArrayMsgpack_Marshal-8                        |    2313030 |    503 ns/op |        50 |        160 |   1.16 |        11565 |    3.15
Benchmark_ShamatonArrayMsgpack_Unmarshal-8                      |    2080389 |    492 ns/op |        50 |        168 |   1.02 |        10401 |    2.93
Benchmark_ShamatonMapMsgpackgen_Marshal-8                       |    7936972 |    147 ns/op |        92 |         96 |   1.17 |        73020 |    1.53
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8                     |    4416087 |    261 ns/op |        92 |        112 |   1.15 |        40628 |    2.33
Benchmark_ShamatonArrayMsgpackgen_Marshal-8                     |   10875586 |    123 ns/op |        50 |         64 |   1.34 |        54377 |    1.92
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8                   |    6395007 |    197 ns/op |        50 |        112 |   1.26 |        31975 |    1.76
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8                   |     277348 |   3892 ns/op |        55 |        440 |   1.08 |         1525 |    8.85
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8                 |     174462 |   6584 ns/op |        55 |       1184 |   1.15 |          959 |    5.56
Benchmark_Bebop_200sc_Marshal-8                                 |   10749061 |    104 ns/op |        55 |         64 |   1.12 |        59119 |    1.62
Benchmark_Bebop_200sc_Unmarshal-8                               |   14900529 |     76 ns/op |        55 |         32 |   1.14 |        81952 |    2.39
Benchmark_Bebop_Wellquite_Marshal-8                             |   14748912 |     85 ns/op |        55 |         64 |   1.27 |        81119 |    1.34
Benchmark_Bebop_Wellquite_Unmarshal-8                           |   15139518 |     82 ns/op |        55 |         32 |   1.26 |        83267 |    2.59
Benchmark_FastJson_Marshal-8                                    |    2375751 |    532 ns/op |       133 |        504 |   1.27 |        31787 |    1.06
Benchmark_FastJson_Unmarshal-8                                  |     781520 |   1994 ns/op |       133 |       1704 |   1.56 |        10456 |    1.17
Benchmark_BENC_Marshal-8                                        |   16171626 |     80 ns/op |        51 |         64 |   1.31 |        82475 |    1.26
Benchmark_BENC_Unmarshal-8                                      |   15027098 |     86 ns/op |        51 |         32 |   1.37 |        76638 |    2.84
Benchmark_BENC_UnsafeStringConvertion_Marshal-8                 |   17900857 |     68 ns/op |        51 |         64 |   1.22 |        91294 |    1.07
Benchmark_BENC_UnsafeStringConvertion_Unmarshal-8               |   40609983 |     29 ns/op |        51 |          0 |   1.19 |       207110 |    0.00
Benchmark_BENC_UnsafeStringConvertion_PreAllocation_Marshal-8   |   35734764 |     34 ns/op |        51 |          0 |   1.22 |       182247 |    0.00
Benchmark_BENC_UnsafeStringConvertion_PreAllocation_Unmarshal-8 |   42486434 |     30 ns/op |        51 |          0 |   1.29 |       216680 |    0.00
Benchmark_MUS_Marshal-8                                         |   12888994 |     92 ns/op |        46 |         48 |   1.20 |        59289 |    1.94
Benchmark_MUS_Unmarshal-8                                       |   10089038 |    122 ns/op |        46 |         32 |   1.24 |        46409 |    3.83
Benchmark_MUSUnsafe_Marshal-8                                   |   13599637 |     87 ns/op |        49 |         64 |   1.19 |        66638 |    1.36
Benchmark_MUSUnsafe_Unmarshal-8                                 |   30552170 |     45 ns/op |        49 |          0 |   1.39 |       149705 |    0.00


Totals:


benchmark                                              | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_BENC_UnsafeStringConvertion_PreAllocation_-8 |   78221198 |     64 ns/op |       102 |          0 |   5.05 |       797856 | Zero-Allocs
Benchmark_BENC_UnsafeStringConvertion_-8               |   58510840 |     97 ns/op |       102 |         64 |   5.70 |       596810 |    1.52
Benchmark_MUSUnsafe_-8                                 |   44151807 |    132 ns/op |        98 |         64 |   5.86 |       432687 |    2.07
Benchmark_BENC_-8                                      |   31198724 |    167 ns/op |       102 |         96 |   5.35 |       318226 |    1.79
Benchmark_Bebop_Wellquite_-8                           |   29888430 |    168 ns/op |       110 |         96 |   5.04 |       328772 |    1.76
Benchmark_Bebop_200sc_-8                               |   25649590 |    180 ns/op |       110 |         96 |   4.63 |       282145 |    1.88
Benchmark_MUS_-8                                       |   22978032 |    215 ns/op |        92 |         80 |   4.95 |       211397 |    2.69
Benchmark_GencodeUnsafe_-8                             |   26291021 |    220 ns/op |        92 |        144 |   5.79 |       241877 |    1.53
Benchmark_XDR2_-8                                      |   20319024 |    255 ns/op |       120 |         96 |   5.19 |       243828 |    2.66
Benchmark_Gencode_-8                                   |   18530515 |    264 ns/op |       106 |        192 |   4.91 |       196423 |    1.38
Benchmark_Gogoprotobuf_-8                              |   19144166 |    301 ns/op |       106 |        160 |   5.77 |       202928 |    1.88
Benchmark_Msgp_-8                                      |   17350336 |    304 ns/op |       194 |        240 |   5.27 |       336596 |    1.27
Benchmark_Colfer_-8                                    |   14607016 |    304 ns/op |       103 |        176 |   4.45 |       150642 |    1.73
Benchmark_ShamatonArrayMsgpackgen_-8                   |   17270593 |    320 ns/op |       100 |        176 |   5.54 |       172705 |    1.82
Benchmark_ShamatonMapMsgpackgen_-8                     |   12353059 |    408 ns/op |       184 |        208 |   5.04 |       227296 |    1.96
Benchmark_Gotiny_-8                                    |   11052132 |    458 ns/op |        95 |        280 |   5.07 |       106089 |    1.64
Benchmark_GotinyNoTime_-8                              |   11608912 |    519 ns/op |        95 |        264 |   6.03 |       111422 |    1.97
Benchmark_FlatBuffers_-8                               |    9269003 |    698 ns/op |       190 |        488 |   6.47 |       176398 |    1.43
Benchmark_Hprose2_-8                                   |    4668285 |    951 ns/op |       170 |        136 |   4.44 |        79594 |    6.99
Benchmark_Pulsar_-8                                    |    4569430 |    981 ns/op |       103 |        560 |   4.49 |        47124 |    1.75
Benchmark_ShamatonMapMsgpack_-8                        |    4576486 |    982 ns/op |       184 |        360 |   4.50 |        84207 |    2.73
Benchmark_ShamatonArrayMsgpack_-8                      |    4393419 |    996 ns/op |       100 |        328 |   4.38 |        43934 |    3.04
Benchmark_CapNProto2_-8                                |    4724015 |   1165 ns/op |       192 |       1724 |   5.51 |        90701 |    0.68
Benchmark_UgorjiCodecMsgpack_-8                        |    3139571 |   1450 ns/op |       182 |       1928 |   4.55 |        57140 |    0.75
Benchmark_Ikea_-8                                      |    3626399 |   1462 ns/op |       110 |        344 |   5.30 |        39890 |    4.25
Benchmark_EasyJson_-8                                  |    3130267 |   1475 ns/op |       303 |       1008 |   4.62 |        94972 |    1.46
Benchmark_UgorjiCodecBinc_-8                           |    2999885 |   1480 ns/op |       190 |       1944 |   4.44 |        56997 |    0.76
Benchmark_Protobuf_-8                                  |    3419352 |   1500 ns/op |       104 |        328 |   5.13 |        35561 |    4.57
Benchmark_VmihailencoMsgpack_-8                        |    3324698 |   1548 ns/op |       184 |        424 |   5.15 |        61174 |    3.65
Benchmark_JsonIter_-8                                  |    2931786 |   1583 ns/op |       282 |        416 |   4.64 |        82852 |    3.81
Benchmark_Hprose_-8                                    |    3065804 |   1704 ns/op |       170 |        770 |   5.22 |        52278 |    2.21
Benchmark_CapNProto_-8                                 |    4206215 |   1753 ns/op |       192 |       4584 |   7.38 |        80759 |    0.38
Benchmark_GoAvro2Binary_-8                             |    2691883 |   1854 ns/op |        94 |       1008 |   4.99 |        25303 |    1.84
Benchmark_Bson_-8                                      |    2432402 |   1922 ns/op |       220 |        600 |   4.68 |        53512 |    3.20
Benchmark_XDR_-8                                       |    2373052 |   1936 ns/op |       183 |        671 |   4.60 |        43661 |    2.89
Benchmark_MongoBson_-8                                 |    1936487 |   2312 ns/op |       220 |        648 |   4.48 |        42602 |    3.57
Benchmark_Binary_-8                                    |    1896958 |   2322 ns/op |       122 |        680 |   4.40 |        23142 |    3.41
Benchmark_FastJson_-8                                  |    3157271 |   2526 ns/op |       267 |       2208 |   7.98 |        84488 |    1.14
Benchmark_Json_-8                                      |    1471050 |   3404 ns/op |       303 |        559 |   5.01 |        44616 |    6.09
Benchmark_Sereal_-8                                    |    1223294 |   4017 ns/op |       264 |       1808 |   4.91 |        32294 |    2.22
Benchmark_GoAvro2Text_-8                               |     998720 |   4514 ns/op |       267 |       2056 |   4.51 |        26725 |    2.20
Benchmark_GoAvro_-8                                    |     970965 |   5821 ns/op |        94 |       2896 |   5.65 |         9127 |    2.01
Benchmark_SSZNoTimeNoStringNoFloatA_-8                 |     451810 |  10476 ns/op |       110 |       1624 |   4.73 |         4969 |    6.45
Benchmark_Gob_-8                                       |     398462 |  21271 ns/op |       325 |       9464 |   8.48 |        12953 |    2.25
Benchmark_Gogojsonpb_-8                                |     215675 |  24381 ns/op |       250 |       6459 |   5.26 |         5411 |    3.77

## Issues


The benchmarks can also be run with validation enabled.

(without perflock, because it's only validation)
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

Bebop (both libraries, by nature of the format) natively supports times rounded to 100ns ticks, and this is what is currently benchmarked (a unix nano timestamp is another valid approach).

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
