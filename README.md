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

2022-09-01 Results with Go 1.16.5 linux/amd64 on an Intel(R) Core(TM) i7-3630QM

benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                    |     999067 |   1143 ns/op |        48 |        168 |   1.14 |         4795 |    6.80
Benchmark_Gotiny_Unmarshal-8                  |    1754476 |    681 ns/op |        48 |        112 |   1.20 |         8421 |    6.08
Benchmark_GotinyNoTime_Marshal-8              |    1000000 |   1156 ns/op |        47 |        168 |   1.16 |         4799 |    6.88
Benchmark_GotinyNoTime_Unmarshal-8            |    1874764 |    637 ns/op |        48 |         96 |   1.19 |         8998 |    6.64
Benchmark_Msgp_Marshal-8                      |    2603031 |    446 ns/op |        97 |        128 |   1.16 |        25249 |    3.49
Benchmark_Msgp_Unmarshal-8                    |    1363006 |    859 ns/op |        97 |        112 |   1.17 |        13221 |    7.67
Benchmark_VmihailencoMsgpack_Marshal-8        |     381148 |   3151 ns/op |       100 |        392 |   1.20 |         3811 |    8.04
Benchmark_VmihailencoMsgpack_Unmarshal-8      |     236878 |   4660 ns/op |       100 |        408 |   1.10 |         2368 |   11.42
Benchmark_Json_Marshal-8                      |     189709 |   6090 ns/op |       151 |        208 |   1.16 |         2875 |   29.28
Benchmark_Json_Unmarshal-8                    |      92833 |  12751 ns/op |       151 |        383 |   1.18 |         1408 |   33.29
Benchmark_JsonIter_Marshal-8                  |     183445 |   6367 ns/op |       141 |        248 |   1.17 |         2592 |   25.67
Benchmark_JsonIter_Unmarshal-8                |     223440 |   5280 ns/op |       141 |        264 |   1.18 |         3157 |   20.00
Benchmark_EasyJson_Marshal-8                  |     269104 |   4160 ns/op |       151 |        896 |   1.12 |         4079 |    4.64
Benchmark_EasyJson_Unmarshal-8                |     279536 |   4280 ns/op |       151 |        160 |   1.20 |         4240 |   26.75
Benchmark_Bson_Marshal-8                      |     402206 |   3858 ns/op |       110 |        376 |   1.55 |         4424 |   10.26
Benchmark_Bson_Unmarshal-8                    |     229580 |   4927 ns/op |       110 |        224 |   1.13 |         2525 |   22.00
Benchmark_MongoBson_Marshal-8                 |     216753 |   5081 ns/op |       110 |        384 |   1.10 |         2384 |   13.23
Benchmark_MongoBson_Unmarshal-8               |     181677 |   5901 ns/op |       110 |        408 |   1.07 |         1998 |   14.46
Benchmark_Gob_Marshal-8                       |      71692 |  16463 ns/op |       163 |       1616 |   1.18 |         1172 |   10.19
Benchmark_Gob_Unmarshal-8                     |      14772 |  84385 ns/op |       163 |       7688 |   1.25 |          241 |   10.98
Benchmark_XDR_Marshal-8                       |     243609 |   4925 ns/op |        92 |        440 |   1.20 |         2241 |   11.19
Benchmark_XDR_Unmarshal-8                     |     266436 |   4470 ns/op |        92 |        232 |   1.19 |         2451 |   19.27
Benchmark_UgorjiCodecMsgpack_Marshal-8        |     382062 |   3097 ns/op |        91 |       1304 |   1.18 |         3476 |    2.38
Benchmark_UgorjiCodecMsgpack_Unmarshal-8      |     326094 |   3561 ns/op |        91 |        496 |   1.16 |         2967 |    7.18
Benchmark_UgorjiCodecBinc_Marshal-8           |     365337 |   3294 ns/op |        95 |       1320 |   1.20 |         3470 |    2.50
Benchmark_UgorjiCodecBinc_Unmarshal-8         |     261735 |   4269 ns/op |        95 |        656 |   1.12 |         2486 |    6.51
Benchmark_Sereal_Marshal-8                    |     147440 |   7773 ns/op |       132 |        832 |   1.15 |         1946 |    9.34
Benchmark_Sereal_Unmarshal-8                  |     123325 |   9199 ns/op |       132 |        976 |   1.13 |         1627 |    9.43
Benchmark_Binary_Marshal-8                    |     325810 |   3790 ns/op |        61 |        312 |   1.23 |         1987 |   12.15
Benchmark_Binary_Unmarshal-8                  |     279754 |   4099 ns/op |        61 |        320 |   1.15 |         1706 |   12.81
Benchmark_FlatBuffers_Marshal-8               |     576027 |   2064 ns/op |        95 |        376 |   1.19 |         5481 |    5.49
Benchmark_FlatBuffers_Unmarshal-8             |    1704805 |    700 ns/op |        95 |        112 |   1.19 |        16250 |    6.26
Benchmark_CapNProto_Marshal-8                 |     266458 |   4089 ns/op |        96 |       4488 |   1.09 |         2557 |    0.91
Benchmark_CapNProto_Unmarshal-8               |    1000000 |   1181 ns/op |        96 |        192 |   1.18 |         9600 |    6.15
Benchmark_CapNProto2_Marshal-8                |     426711 |   2802 ns/op |        96 |       1452 |   1.20 |         4096 |    1.93
Benchmark_CapNProto2_Unmarshal-8              |     953179 |   1220 ns/op |        96 |        272 |   1.16 |         9150 |    4.49
Benchmark_Hprose_Marshal-8                    |     479809 |   2577 ns/op |        85 |        460 |   1.24 |         4089 |    5.60
Benchmark_Hprose_Unmarshal-8                  |     358312 |   3077 ns/op |        85 |        304 |   1.10 |         3054 |   10.12
Benchmark_Hprose2_Marshal-8                   |     792824 |   1491 ns/op |        85 |          0 |   1.18 |         6761 |    0.00
Benchmark_Hprose2_Unmarshal-8                 |     816315 |   1510 ns/op |        85 |        136 |   1.23 |         6961 |   11.10
Benchmark_Protobuf_Marshal-8                  |     580714 |   2094 ns/op |        52 |        144 |   1.22 |         3019 |   14.54
Benchmark_Protobuf_Unmarshal-8                |     601705 |   2030 ns/op |        52 |        184 |   1.22 |         3128 |   11.03
Benchmark_Goprotobuf_Marshal-8                |    1405010 |    854 ns/op |        53 |         64 |   1.20 |         7446 |   13.36
Benchmark_Goprotobuf_Unmarshal-8              |     973688 |   1255 ns/op |        53 |        168 |   1.22 |         5160 |    7.47
Benchmark_Gogoprotobuf_Marshal-8              |    3359550 |    354 ns/op |        53 |         64 |   1.19 |        17805 |    5.54
Benchmark_Gogoprotobuf_Unmarshal-8            |    1908633 |    619 ns/op |        53 |         96 |   1.18 |        10115 |    6.45
Benchmark_Gogojsonpb_Marshal-8                |      33103 |  35743 ns/op |       125 |       3107 |   1.18 |          414 |   11.50
Benchmark_Gogojsonpb_Unmarshal-8              |      24392 |  48703 ns/op |       126 |       3616 |   1.19 |          307 |   13.47
Benchmark_Colfer_Marshal-8                    |    3614686 |    334 ns/op |        51 |         64 |   1.21 |        18471 |    5.22
Benchmark_Colfer_Unmarshal-8                  |    2338790 |    513 ns/op |        52 |        112 |   1.20 |        12161 |    4.58
Benchmark_Gencode_Marshal-8                   |    2716658 |    440 ns/op |        53 |         80 |   1.20 |        14398 |    5.50
Benchmark_Gencode_Unmarshal-8                 |    2340530 |    507 ns/op |        53 |        112 |   1.19 |        12404 |    4.53
Benchmark_GencodeUnsafe_Marshal-8             |    4467661 |    267 ns/op |        46 |         48 |   1.20 |        20551 |    5.58
Benchmark_GencodeUnsafe_Unmarshal-8           |    2887224 |    414 ns/op |        46 |         96 |   1.20 |        13281 |    4.32
Benchmark_XDR2_Marshal-8                      |    2695273 |    431 ns/op |        60 |         64 |   1.16 |        16171 |    6.74
Benchmark_XDR2_Unmarshal-8                    |    3271778 |    364 ns/op |        60 |         32 |   1.19 |        19630 |   11.38
Benchmark_GoAvro_Marshal-8                    |     156092 |   7235 ns/op |        47 |        872 |   1.13 |          733 |    8.30
Benchmark_GoAvro_Unmarshal-8                  |      66531 |  17734 ns/op |        47 |       3024 |   1.18 |          312 |    5.86
Benchmark_GoAvro2Text_Marshal-8               |     151585 |   7631 ns/op |       133 |       1320 |   1.16 |         2028 |    5.78
Benchmark_GoAvro2Text_Unmarshal-8             |     148945 |   7776 ns/op |       133 |        784 |   1.16 |         1992 |    9.92
Benchmark_GoAvro2Binary_Marshal-8             |     499084 |   2442 ns/op |        47 |        464 |   1.22 |         2345 |    5.26
Benchmark_GoAvro2Binary_Unmarshal-8           |     435390 |   2776 ns/op |        47 |        544 |   1.21 |         2046 |    5.10
Benchmark_Ikea_Marshal-8                      |     396324 |   3075 ns/op |        55 |        184 |   1.22 |         2179 |   16.71
Benchmark_Ikea_Unmarshal-8                    |     530564 |   2299 ns/op |        55 |        160 |   1.22 |         2918 |   14.37
Benchmark_ShamatonMapMsgpack_Marshal-8        |     581019 |   2073 ns/op |        92 |        192 |   1.20 |         5345 |   10.80
Benchmark_ShamatonMapMsgpack_Unmarshal-8      |     606480 |   2007 ns/op |        92 |        136 |   1.22 |         5579 |   14.76
Benchmark_ShamatonArrayMsgpack_Marshal-8      |     654516 |   1876 ns/op |        50 |        160 |   1.23 |         3272 |   11.72
Benchmark_ShamatonArrayMsgpack_Unmarshal-8    |     864189 |   1395 ns/op |        50 |        136 |   1.21 |         4320 |   10.26
Benchmark_ShamatonMapMsgpackgen_Marshal-8     |    2120286 |    561 ns/op |        92 |         96 |   1.19 |        19506 |    5.85
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8   |    1858796 |    632 ns/op |        92 |         80 |   1.18 |        17100 |    7.91
Benchmark_ShamatonArrayMsgpackgen_Marshal-8   |    2739808 |    433 ns/op |        50 |         64 |   1.19 |        13699 |    6.78
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8 |    2829740 |    426 ns/op |        50 |         80 |   1.21 |        14148 |    5.33
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8 |      90843 |  13073 ns/op |        55 |        440 |   1.19 |          499 |   29.71
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8 |      55366 |  21635 ns/op |        55 |       1184 |   1.20 |          304 |   18.27
Benchmark_Mum_Marshal-8                       |    4099479 |    291 ns/op |        48 |          0 |   1.19 |        19677 |    0.00
Benchmark_Mum_Unmarshal-8                     |    2137136 |    563 ns/op |        48 |         80 |   1.20 |        10258 |    7.04
Benchmark_Bebop_Marshal-8                     |    3452628 |    342 ns/op |        55 |         64 |   1.18 |        18989 |    5.35
Benchmark_Bebop_Unmarshal-8                   |    4152378 |    286 ns/op |        55 |         32 |   1.19 |        22838 |    8.95
Benchmark_FastJson_Marshal-8                  |     712838 |   1667 ns/op |       133 |        504 |   1.19 |         9530 |    3.31
Benchmark_FastJson_Unmarshal-8                |     230169 |   4994 ns/op |       133 |       1704 |   1.15 |         3077 |    2.93
Benchmark_Musgo_Marshal-8                     |    4294477 |    280 ns/op |        46 |         48 |   1.21 |        19754 |    5.85
Benchmark_Musgo_Unmarshal-8                   |    2498404 |    480 ns/op |        46 |         96 |   1.20 |        11492 |    5.00
Benchmark_MusgoUnsafe_Marshal-8               |    4565900 |    261 ns/op |        46 |         48 |   1.19 |        21003 |    5.44
Benchmark_MusgoUnsafe_Unmarshal-8             |    4103084 |    286 ns/op |        46 |         64 |   1.17 |        18874 |    4.47


Totals:


benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MusgoUnsafe_-8                 |    8668984 |    547 ns/op |        92 |        112 |   4.74 |        79754 |    4.89
Benchmark_Bebop_-8                       |    7605006 |    628 ns/op |       110 |         96 |   4.78 |        83655 |    6.55
Benchmark_GencodeUnsafe_-8               |    7354885 |    682 ns/op |        92 |        144 |   5.02 |        67664 |    4.74
Benchmark_Musgo_-8                       |    6792881 |    760 ns/op |        92 |        144 |   5.17 |        62494 |    5.28
Benchmark_XDR2_-8                        |    5967051 |    795 ns/op |       120 |         96 |   4.75 |        71604 |    8.28
Benchmark_Colfer_-8                      |    5953476 |    847 ns/op |       103 |        176 |   5.04 |        61380 |    4.81
Benchmark_Mum_-8                         |    6236615 |    854 ns/op |        96 |         80 |   5.33 |        59871 |   10.68
Benchmark_ShamatonArrayMsgpackgen_-8     |    5569548 |    859 ns/op |       100 |        144 |   4.79 |        55695 |    5.97
Benchmark_Gencode_-8                     |    5057188 |    947 ns/op |       106 |        192 |   4.79 |        53606 |    4.93
Benchmark_Gogoprotobuf_-8                |    5268183 |    974 ns/op |       106 |        160 |   5.13 |        55842 |    6.09
Benchmark_ShamatonMapMsgpackgen_-8       |    3979082 |   1194 ns/op |       184 |        176 |   4.75 |        73215 |    6.79
Benchmark_Msgp_-8                        |    3966037 |   1305 ns/op |       194 |        240 |   5.18 |        76941 |    5.44
Benchmark_GotinyNoTime_-8                |    2874764 |   1793 ns/op |        95 |        264 |   5.16 |        27594 |    6.79
Benchmark_Gotiny_-8                      |    2753543 |   1824 ns/op |        96 |        280 |   5.02 |        26434 |    6.52
Benchmark_Goprotobuf_-8                  |    2378698 |   2109 ns/op |       106 |        232 |   5.02 |        25214 |    9.09
Benchmark_FlatBuffers_-8                 |    2280832 |   2764 ns/op |       190 |        488 |   6.31 |        43445 |    5.67
Benchmark_Hprose2_-8                     |    1609139 |   3001 ns/op |       170 |        136 |   4.83 |        27447 |   22.07
Benchmark_ShamatonArrayMsgpack_-8        |    1518705 |   3271 ns/op |       100 |        296 |   4.97 |        15187 |   11.05
Benchmark_CapNProto2_-8                  |    1379890 |   4022 ns/op |       192 |       1724 |   5.55 |        26493 |    2.33
Benchmark_ShamatonMapMsgpack_-8          |    1187499 |   4080 ns/op |       184 |        328 |   4.84 |        21849 |   12.44
Benchmark_Protobuf_-8                    |    1182419 |   4124 ns/op |       104 |        328 |   4.88 |        12297 |   12.57
Benchmark_GoAvro2Binary_-8               |     934474 |   5218 ns/op |        94 |       1008 |   4.88 |         8784 |    5.18
Benchmark_CapNProto_-8                   |    1266458 |   5270 ns/op |       192 |       4680 |   6.67 |        24315 |    1.13
Benchmark_Ikea_-8                        |     926888 |   5374 ns/op |       110 |        344 |   4.98 |        10195 |   15.62
Benchmark_Hprose_-8                      |     838121 |   5654 ns/op |       170 |        764 |   4.74 |        14289 |    7.40
Benchmark_UgorjiCodecMsgpack_-8          |     708156 |   6658 ns/op |       182 |       1800 |   4.71 |        12888 |    3.70
Benchmark_FastJson_-8                    |     943007 |   6661 ns/op |       267 |       2208 |   6.28 |        25216 |    3.02
Benchmark_UgorjiCodecBinc_-8             |     627072 |   7563 ns/op |       190 |       1976 |   4.74 |        11914 |    3.83
Benchmark_VmihailencoMsgpack_-8          |     618026 |   7811 ns/op |       200 |        800 |   4.83 |        12360 |    9.76
Benchmark_Binary_-8                      |     605564 |   7889 ns/op |       122 |        632 |   4.78 |         7387 |   12.48
Benchmark_EasyJson_-8                    |     548640 |   8440 ns/op |       303 |       1056 |   4.63 |        16640 |    7.99
Benchmark_Bson_-8                        |     631786 |   8785 ns/op |       220 |        600 |   5.55 |        13899 |   14.64
Benchmark_XDR_-8                         |     510045 |   9395 ns/op |       184 |        672 |   4.79 |         9384 |   13.98
Benchmark_MongoBson_-8                   |     398430 |  10982 ns/op |       220 |        792 |   4.38 |         8765 |   13.87
Benchmark_JsonIter_-8                    |     406885 |  11647 ns/op |       282 |        512 |   4.74 |        11498 |   22.75
Benchmark_GoAvro2Text_-8                 |     300530 |  15407 ns/op |       267 |       2104 |   4.63 |         8042 |    7.32
Benchmark_Sereal_-8                      |     270765 |  16972 ns/op |       264 |       1808 |   4.60 |         7148 |    9.39
Benchmark_Json_-8                        |     282542 |  18841 ns/op |       303 |        591 |   5.32 |         8569 |   31.88
Benchmark_GoAvro_-8                      |     222623 |  24969 ns/op |        94 |       3896 |   5.56 |         2092 |    6.41
Benchmark_SSZNoTimeNoStringNoFloatA_-8   |     146209 |  34708 ns/op |       110 |       1624 |   5.07 |         1608 |   21.37
Benchmark_Gogojsonpb_-8                  |      57495 |  84446 ns/op |       251 |       6723 |   4.86 |         1444 |   12.56
Benchmark_Gob_-8                         |      86464 | 100848 ns/op |       327 |       9304 |   8.72 |         2829 |   10.84


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
