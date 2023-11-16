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

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                       |    4270570 |    277 ns/op |        48 |        168 |   1.18 |        20498 |    1.65
Benchmark_Gotiny_Unmarshal-8                     |    8058099 |    177 ns/op |        47 |        112 |   1.43 |        38670 |    1.58
Benchmark_GotinyNoTime_Marshal-8                 |    3805028 |    286 ns/op |        48 |        168 |   1.09 |        18264 |    1.70
Benchmark_GotinyNoTime_Unmarshal-8               |    7809414 |    140 ns/op |        48 |         96 |   1.10 |        37485 |    1.47
Benchmark_Msgp_Marshal-8                         |   11670178 |    101 ns/op |        97 |        128 |   1.18 |       113200 |    0.79
Benchmark_Msgp_Unmarshal-8                       |    5751547 |    190 ns/op |        97 |        112 |   1.10 |        55790 |    1.70
Benchmark_VmihailencoMsgpack_Marshal-8           |    1836144 |    571 ns/op |        92 |        264 |   1.05 |        16892 |    2.16
Benchmark_VmihailencoMsgpack_Unmarshal-8         |    1499811 |    787 ns/op |        92 |        160 |   1.18 |        13798 |    4.92
Benchmark_Json_Marshal-8                         |    1000000 |   1073 ns/op |       151 |        208 |   1.07 |        15170 |    5.16
Benchmark_Json_Unmarshal-8                       |     528459 |   2234 ns/op |       151 |        351 |   1.18 |         8016 |    6.36
Benchmark_JsonIter_Marshal-8                     |    1916157 |    571 ns/op |       141 |        200 |   1.10 |        27075 |    2.86
Benchmark_JsonIter_Unmarshal-8                   |    1408056 |    844 ns/op |       141 |        216 |   1.19 |        19895 |    3.91
Benchmark_EasyJson_Marshal-8                     |    1483146 |    754 ns/op |       151 |        896 |   1.12 |        22499 |    0.84
Benchmark_EasyJson_Unmarshal-8                   |    1660232 |    705 ns/op |       151 |        112 |   1.17 |        25185 |    6.30
Benchmark_Bson_Marshal-8                         |    1501688 |    725 ns/op |       110 |        376 |   1.09 |        16518 |    1.93
Benchmark_Bson_Unmarshal-8                       |     986688 |   1155 ns/op |       110 |        224 |   1.14 |        10853 |    5.16
Benchmark_MongoBson_Marshal-8                    |    1000000 |   1083 ns/op |       110 |        240 |   1.08 |        11000 |    4.51
Benchmark_MongoBson_Unmarshal-8                  |     981248 |   1197 ns/op |       110 |        408 |   1.17 |        10793 |    2.93
Benchmark_Gob_Marshal-8                          |     334194 |   3626 ns/op |       162 |       1744 |   1.21 |         5433 |    2.08
Benchmark_Gob_Unmarshal-8                        |      69415 |  17147 ns/op |       162 |       7720 |   1.19 |         1128 |    2.22
Benchmark_XDR_Marshal-8                          |    1000000 |   1005 ns/op |        91 |        439 |   1.00 |         9199 |    2.29
Benchmark_XDR_Unmarshal-8                        |    1437243 |    854 ns/op |        92 |        231 |   1.23 |        13222 |    3.70
Benchmark_UgorjiCodecMsgpack_Marshal-8           |    1458154 |    699 ns/op |        91 |       1240 |   1.02 |        13269 |    0.56
Benchmark_UgorjiCodecMsgpack_Unmarshal-8         |    1583008 |    746 ns/op |        91 |        688 |   1.18 |        14405 |    1.09
Benchmark_UgorjiCodecBinc_Marshal-8              |    1541138 |    723 ns/op |        95 |       1256 |   1.11 |        14640 |    0.58
Benchmark_UgorjiCodecBinc_Unmarshal-8            |    1692334 |    709 ns/op |        95 |        688 |   1.20 |        16077 |    1.03
Benchmark_Sereal_Marshal-8                       |     649387 |   1816 ns/op |       132 |        832 |   1.18 |         8571 |    2.18
Benchmark_Sereal_Unmarshal-8                     |     559626 |   2127 ns/op |       132 |        976 |   1.19 |         7387 |    2.18
Benchmark_Binary_Marshal-8                       |     896160 |   1230 ns/op |        61 |        360 |   1.10 |         5466 |    3.42
Benchmark_Binary_Unmarshal-8                     |    1000000 |   1060 ns/op |        61 |        320 |   1.06 |         6100 |    3.31
Benchmark_FlatBuffers_Marshal-8                  |    2075660 |    531 ns/op |        95 |        376 |   1.10 |        19735 |    1.41
Benchmark_FlatBuffers_Unmarshal-8                |    7277006 |    161 ns/op |        95 |        112 |   1.18 |        69357 |    1.44
Benchmark_CapNProto_Marshal-8                    |     878935 |   1499 ns/op |        96 |       4392 |   1.32 |         8437 |    0.34
Benchmark_CapNProto_Unmarshal-8                  |    3542133 |    333 ns/op |        96 |        192 |   1.18 |        34004 |    1.74
Benchmark_CapNProto2_Marshal-8                   |    1442034 |    816 ns/op |        96 |       1452 |   1.18 |        13843 |    0.56
Benchmark_CapNProto2_Unmarshal-8                 |    3247956 |    336 ns/op |        96 |        272 |   1.09 |        31180 |    1.24
Benchmark_Hprose_Marshal-8                       |    1687570 |    721 ns/op |        85 |        462 |   1.22 |        14396 |    1.56
Benchmark_Hprose_Unmarshal-8                     |    1439083 |    787 ns/op |        85 |        304 |   1.13 |        12269 |    2.59
Benchmark_Hprose2_Marshal-8                      |    3498620 |    345 ns/op |        85 |          0 |   1.21 |        29832 |    0.00
Benchmark_Hprose2_Unmarshal-8                    |    2788557 |    397 ns/op |        85 |        136 |   1.11 |        23778 |    2.93
Benchmark_Protobuf_Marshal-8                     |    2458848 |    477 ns/op |        52 |        144 |   1.17 |        12786 |    3.31
Benchmark_Protobuf_Unmarshal-8                   |    1764528 |    690 ns/op |        52 |        184 |   1.22 |         9175 |    3.75
Benchmark_Pulsar_Marshal-8                       |    2533876 |    444 ns/op |        51 |        304 |   1.13 |        13077 |    1.46
Benchmark_Pulsar_Unmarshal-8                     |    2581611 |    418 ns/op |        51 |        256 |   1.08 |        13328 |    1.63
Benchmark_Gogoprotobuf_Marshal-8                 |   13846714 |     85 ns/op |        53 |         64 |   1.19 |        73387 |    1.34
Benchmark_Gogoprotobuf_Unmarshal-8               |    7278008 |    169 ns/op |        53 |         96 |   1.23 |        38573 |    1.77
Benchmark_Gogojsonpb_Marshal-8                   |     155838 |   7659 ns/op |       124 |       3065 |   1.19 |         1946 |    2.50
Benchmark_Gogojsonpb_Unmarshal-8                 |     117559 |  10185 ns/op |       125 |       3373 |   1.20 |         1474 |    3.02
Benchmark_Colfer_Marshal-8                       |   15176950 |     86 ns/op |        51 |         64 |   1.31 |        77523 |    1.35
Benchmark_Colfer_Unmarshal-8                     |    9196304 |    148 ns/op |        51 |        112 |   1.36 |        46901 |    1.32
Benchmark_Gencode_Marshal-8                      |   11268625 |    104 ns/op |        53 |         80 |   1.18 |        59723 |    1.31
Benchmark_Gencode_Unmarshal-8                    |    9081552 |    124 ns/op |        53 |        112 |   1.13 |        48132 |    1.11
Benchmark_GencodeUnsafe_Marshal-8                |   19534160 |     61 ns/op |        46 |         48 |   1.21 |        89857 |    1.29
Benchmark_GencodeUnsafe_Unmarshal-8              |   11635015 |    104 ns/op |        46 |         96 |   1.21 |        53521 |    1.08
Benchmark_XDR2_Marshal-8                         |   10992777 |    106 ns/op |        60 |         64 |   1.17 |        65956 |    1.66
Benchmark_XDR2_Unmarshal-8                       |   15285810 |     78 ns/op |        60 |         32 |   1.20 |        91714 |    2.46
Benchmark_GoAvro_Marshal-8                       |     862538 |   1395 ns/op |        47 |        584 |   1.20 |         4053 |    2.39
Benchmark_GoAvro_Unmarshal-8                     |     342262 |   3596 ns/op |        47 |       2312 |   1.23 |         1608 |    1.56
Benchmark_GoAvro2Text_Marshal-8                  |     562929 |   2081 ns/op |       133 |       1320 |   1.17 |         7526 |    1.58
Benchmark_GoAvro2Text_Unmarshal-8                |     637239 |   1942 ns/op |       133 |        736 |   1.24 |         8526 |    2.64
Benchmark_GoAvro2Binary_Marshal-8                |    1879074 |    620 ns/op |        47 |        464 |   1.17 |         8831 |    1.34
Benchmark_GoAvro2Binary_Unmarshal-8              |    1693023 |    655 ns/op |        47 |        544 |   1.11 |         7957 |    1.20
Benchmark_Ikea_Marshal-8                         |    1931623 |    608 ns/op |        55 |        184 |   1.17 |        10623 |    3.31
Benchmark_Ikea_Unmarshal-8                       |    2193319 |    531 ns/op |        55 |        160 |   1.17 |        12063 |    3.32
Benchmark_ShamatonMapMsgpack_Marshal-8           |    2245851 |    496 ns/op |        92 |        192 |   1.11 |        20661 |    2.58
Benchmark_ShamatonMapMsgpack_Unmarshal-8         |    2511224 |    420 ns/op |        92 |        168 |   1.05 |        23103 |    2.50
Benchmark_ShamatonArrayMsgpack_Marshal-8         |    2803377 |    404 ns/op |        50 |        160 |   1.13 |        14016 |    2.53
Benchmark_ShamatonArrayMsgpack_Unmarshal-8       |    3239434 |    362 ns/op |        50 |        168 |   1.17 |        16197 |    2.16
Benchmark_ShamatonMapMsgpackgen_Marshal-8        |    8974251 |    130 ns/op |        92 |         96 |   1.17 |        82563 |    1.36
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8      |    4937666 |    239 ns/op |        92 |        112 |   1.18 |        45426 |    2.13
Benchmark_ShamatonArrayMsgpackgen_Marshal-8      |   11963331 |     98 ns/op |        50 |         64 |   1.18 |        59816 |    1.55
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8    |    8094484 |    153 ns/op |        50 |        112 |   1.24 |        40472 |    1.37
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8    |     386017 |   3026 ns/op |        55 |        440 |   1.17 |         2123 |    6.88
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8  |     237327 |   5098 ns/op |        55 |       1184 |   1.21 |         1305 |    4.31
Benchmark_Bebop_200sc_Marshal-8                  |   16402246 |     73 ns/op |        55 |         64 |   1.20 |        90212 |    1.14
Benchmark_Bebop_200sc_Unmarshal-8                |   17687097 |     70 ns/op |        55 |         32 |   1.24 |        97279 |    2.19
Benchmark_Bebop_Wellquite_Marshal-8              |   16185852 |     74 ns/op |        55 |         64 |   1.20 |        89022 |    1.16
Benchmark_Bebop_Wellquite_Unmarshal-8            |   16249633 |     76 ns/op |        55 |         32 |   1.24 |        89372 |    2.38
Benchmark_FastJson_Marshal-8                     |    2836292 |    409 ns/op |       133 |        504 |   1.16 |        37949 |    0.81
Benchmark_FastJson_Unmarshal-8                   |     880944 |   1146 ns/op |       133 |       1704 |   1.01 |        11787 |    0.67
Benchmark_BENC_Marshal-8                         |   21569998 |     55 ns/op |        51 |         64 |   1.19 |       110006 |    0.87
Benchmark_BENC_Unmarshal-8                       |   18378385 |     67 ns/op |        51 |         32 |   1.24 |        93729 |    2.12
Benchmark_MUS_Marshal-8                          |   18301550 |     65 ns/op |        46 |         48 |   1.20 |        84187 |    1.37
Benchmark_MUS_Unmarshal-8                        |   10975488 |     98 ns/op |        46 |         32 |   1.08 |        50487 |    3.07
Benchmark_MUSUnsafe_Marshal-8                    |   19911859 |     61 ns/op |        49 |         64 |   1.22 |        97568 |    0.96
Benchmark_MUSUnsafe_Unmarshal-8                  |   32676490 |     37 ns/op |        49 |          0 |   1.23 |       160114 |    0.00


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MUSUnsafe_-8                  |   52588349 |     98 ns/op |        98 |         64 |   5.20 |       515365 |    1.54
Benchmark_BENC_-8                       |   39948383 |    123 ns/op |       102 |         96 |   4.92 |       407473 |    1.28
Benchmark_Bebop_200sc_-8                |   34089343 |    143 ns/op |       110 |         96 |   4.89 |       374982 |    1.49
Benchmark_Bebop_Wellquite_-8            |   32435485 |    150 ns/op |       110 |         96 |   4.88 |       356790 |    1.57
Benchmark_MUS_-8                        |   29277038 |    163 ns/op |        92 |         80 |   4.79 |       269348 |    2.05
Benchmark_GencodeUnsafe_-8              |   31169175 |    165 ns/op |        92 |        144 |   5.17 |       286756 |    1.15
Benchmark_XDR2_-8                       |   26278587 |    185 ns/op |       120 |         96 |   4.86 |       315343 |    1.93
Benchmark_Gencode_-8                    |   20350177 |    229 ns/op |       106 |        192 |   4.67 |       215711 |    1.19
Benchmark_Colfer_-8                     |   24373254 |    234 ns/op |       102 |        176 |   5.73 |       248802 |    1.33
Benchmark_ShamatonArrayMsgpackgen_-8    |   20057815 |    251 ns/op |       100 |        176 |   5.05 |       200578 |    1.43
Benchmark_Gogoprotobuf_-8               |   21124722 |    255 ns/op |       106 |        160 |   5.40 |       223922 |    1.60
Benchmark_Msgp_-8                       |   17421725 |    291 ns/op |       194 |        240 |   5.09 |       337981 |    1.22
Benchmark_ShamatonMapMsgpackgen_-8      |   13911917 |    369 ns/op |       184 |        208 |   5.14 |       255979 |    1.78
Benchmark_GotinyNoTime_-8               |   11614442 |    426 ns/op |        96 |        264 |   4.96 |       111498 |    1.62
Benchmark_Gotiny_-8                     |   12328669 |    454 ns/op |        95 |        280 |   5.60 |       118342 |    1.62
Benchmark_FlatBuffers_-8                |    9352666 |    693 ns/op |       190 |        488 |   6.48 |       178065 |    1.42
Benchmark_Hprose2_-8                    |    6287177 |    743 ns/op |       170 |        136 |   4.67 |       107221 |    5.46
Benchmark_ShamatonArrayMsgpack_-8       |    6042811 |    766 ns/op |       100 |        328 |   4.63 |        60428 |    2.34
Benchmark_Pulsar_-8                     |    5115487 |    862 ns/op |       103 |        560 |   4.41 |        52812 |    1.54
Benchmark_ShamatonMapMsgpack_-8         |    4757075 |    916 ns/op |       184 |        360 |   4.36 |        87530 |    2.55
Benchmark_Ikea_-8                       |    4124942 |   1139 ns/op |       110 |        344 |   4.70 |        45374 |    3.31
Benchmark_CapNProto2_-8                 |    4689990 |   1152 ns/op |       192 |       1724 |   5.40 |        90047 |    0.67
Benchmark_Protobuf_-8                   |    4223376 |   1168 ns/op |       104 |        328 |   4.93 |        43923 |    3.56
Benchmark_GoAvro2Binary_-8              |    3572097 |   1275 ns/op |        94 |       1008 |   4.56 |        33577 |    1.27
Benchmark_VmihailencoMsgpack_-8         |    3335955 |   1358 ns/op |       184 |        424 |   4.53 |        61381 |    3.20
Benchmark_JsonIter_-8                   |    3324213 |   1416 ns/op |       282 |        416 |   4.71 |        93942 |    3.40
Benchmark_UgorjiCodecBinc_-8            |    3233472 |   1432 ns/op |       190 |       1944 |   4.63 |        61435 |    0.74
Benchmark_UgorjiCodecMsgpack_-8         |    3041162 |   1446 ns/op |       182 |       1928 |   4.40 |        55349 |    0.75
Benchmark_EasyJson_-8                   |    3143378 |   1460 ns/op |       303 |       1008 |   4.59 |        95370 |    1.45
Benchmark_Hprose_-8                     |    3126653 |   1509 ns/op |       170 |        766 |   4.72 |        53331 |    1.97
Benchmark_FastJson_-8                   |    3717236 |   1555 ns/op |       267 |       2208 |   5.78 |        99473 |    0.70
Benchmark_CapNProto_-8                  |    4421068 |   1832 ns/op |       192 |       4584 |   8.10 |        84884 |    0.40
Benchmark_XDR_-8                        |    2437243 |   1859 ns/op |       183 |        670 |   4.53 |        44842 |    2.78
Benchmark_Bson_-8                       |    2488376 |   1880 ns/op |       220 |        600 |   4.68 |        54744 |    3.13
Benchmark_MongoBson_-8                  |    1981248 |   2280 ns/op |       220 |        648 |   4.52 |        43587 |    3.52
Benchmark_Binary_-8                     |    1896160 |   2290 ns/op |       122 |        680 |   4.34 |        23133 |    3.37
Benchmark_Json_-8                       |    1528459 |   3307 ns/op |       303 |        559 |   5.05 |        46373 |    5.92
Benchmark_Sereal_-8                     |    1209013 |   3943 ns/op |       264 |       1808 |   4.77 |        31917 |    2.18
Benchmark_GoAvro2Text_-8                |    1200168 |   4023 ns/op |       267 |       2056 |   4.83 |        32104 |    1.96
Benchmark_GoAvro_-8                     |    1204800 |   4991 ns/op |        94 |       2896 |   6.01 |        11325 |    1.72
Benchmark_SSZNoTimeNoStringNoFloatA_-8  |     623344 |   8124 ns/op |       110 |       1624 |   5.06 |         6856 |    5.00
Benchmark_Gogojsonpb_-8                 |     273397 |  17844 ns/op |       250 |       6438 |   4.88 |         6843 |    2.77
Benchmark_Gob_-8                        |     403609 |  20773 ns/op |       325 |       9464 |   8.38 |        13125 |    2.19

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
