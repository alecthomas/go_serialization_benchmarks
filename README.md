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

2023-02-06 Results with Go 1.20 linux/amd64 on an `AMD Ryzen 9 3900X 12-Core Processor`

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-24                      |    2404591 |    525 ns/op |        48 |        168 |   1.26 |        11542 |    3.12
Benchmark_Gotiny_Unmarshal-24                    |    4128519 |    280 ns/op |        48 |        112 |   1.16 |        19816 |    2.51
Benchmark_GotinyNoTime_Marshal-24                |    2301190 |    523 ns/op |        48 |        168 |   1.20 |        11045 |    3.12
Benchmark_GotinyNoTime_Unmarshal-24              |    4499371 |    257 ns/op |        48 |         96 |   1.16 |        21596 |    2.69
Benchmark_Msgp_Marshal-24                        |    6255897 |    190 ns/op |        97 |        128 |   1.19 |        60682 |    1.49
Benchmark_Msgp_Unmarshal-24                      |    3318363 |    364 ns/op |        97 |        112 |   1.21 |        32188 |    3.25
Benchmark_VmihailencoMsgpack_Marshal-24          |     962323 |   1155 ns/op |        92 |        264 |   1.11 |         8853 |    4.38
Benchmark_VmihailencoMsgpack_Unmarshal-24        |     667622 |   1586 ns/op |        92 |        160 |   1.06 |         6142 |    9.91
Benchmark_Json_Marshal-24                        |    1000000 |   1827 ns/op |       151 |        208 |   1.83 |        15160 |    8.78
Benchmark_Json_Unmarshal-24                      |     259165 |   4483 ns/op |       151 |        351 |   1.16 |         3928 |   12.77
Benchmark_JsonIter_Marshal-24                    |     905325 |   1242 ns/op |       141 |        200 |   1.12 |        12792 |    6.21
Benchmark_JsonIter_Unmarshal-24                  |     659034 |   1656 ns/op |       141 |        216 |   1.09 |         9312 |    7.67
Benchmark_EasyJson_Marshal-24                    |     897210 |   1395 ns/op |       151 |        896 |   1.25 |        13601 |    1.56
Benchmark_EasyJson_Unmarshal-24                  |     824758 |   1406 ns/op |       151 |        112 |   1.16 |        12503 |   12.55
Benchmark_Bson_Marshal-24                        |     793263 |   1464 ns/op |       110 |        376 |   1.16 |         8725 |    3.89
Benchmark_Bson_Unmarshal-24                      |     482554 |   2351 ns/op |       110 |        224 |   1.13 |         5308 |   10.50
Benchmark_MongoBson_Marshal-24                   |     527124 |   2136 ns/op |       110 |        240 |   1.13 |         5798 |    8.90
Benchmark_MongoBson_Unmarshal-24                 |     528327 |   2374 ns/op |       110 |        408 |   1.25 |         5811 |    5.82
Benchmark_Gob_Marshal-24                         |     169804 |   6793 ns/op |       163 |       1616 |   1.15 |         2777 |    4.20
Benchmark_Gob_Unmarshal-24                       |      33820 |  35276 ns/op |       163 |       7768 |   1.19 |          553 |    4.54
Benchmark_XDR_Marshal-24                         |     533103 |   2028 ns/op |        92 |        439 |   1.08 |         4904 |    4.62
Benchmark_XDR_Unmarshal-24                       |     669247 |   1691 ns/op |        92 |        231 |   1.13 |         6157 |    7.32
Benchmark_UgorjiCodecMsgpack_Marshal-24          |     930394 |   1193 ns/op |        91 |       1240 |   1.11 |         8466 |    0.96
Benchmark_UgorjiCodecMsgpack_Unmarshal-24        |     701311 |   1490 ns/op |        91 |        688 |   1.04 |         6381 |    2.17
Benchmark_UgorjiCodecBinc_Marshal-24             |     822146 |   1367 ns/op |        95 |       1256 |   1.12 |         7810 |    1.09
Benchmark_UgorjiCodecBinc_Unmarshal-24           |     805404 |   1362 ns/op |        95 |        688 |   1.10 |         7651 |    1.98
Benchmark_Sereal_Marshal-24                      |     399870 |   3612 ns/op |       132 |        832 |   1.44 |         5278 |    4.34
Benchmark_Sereal_Unmarshal-24                    |     298658 |   3769 ns/op |       132 |        976 |   1.13 |         3942 |    3.86
Benchmark_Binary_Marshal-24                      |     489596 |   2319 ns/op |        61 |        360 |   1.14 |         2986 |    6.44
Benchmark_Binary_Unmarshal-24                    |     567476 |   1965 ns/op |        61 |        320 |   1.12 |         3461 |    6.14
Benchmark_FlatBuffers_Marshal-24                 |    1000000 |   1024 ns/op |        95 |        376 |   1.02 |         9518 |    2.72
Benchmark_FlatBuffers_Unmarshal-24               |    4338973 |    273 ns/op |        95 |        112 |   1.19 |        41280 |    2.44
Benchmark_CapNProto_Marshal-24                   |     554866 |   2094 ns/op |        96 |       4392 |   1.16 |         5326 |    0.48
Benchmark_CapNProto_Unmarshal-24                 |    2061580 |    578 ns/op |        96 |        192 |   1.19 |        19791 |    3.01
Benchmark_CapNProto2_Marshal-24                  |     830979 |   1334 ns/op |        96 |       1452 |   1.11 |         7977 |    0.92
Benchmark_CapNProto2_Unmarshal-24                |    2015412 |    586 ns/op |        96 |        272 |   1.18 |        19347 |    2.16
Benchmark_Hprose_Marshal-24                      |    1531195 |    706 ns/op |        85 |        319 |   1.08 |        13054 |    2.21
Benchmark_Hprose_Unmarshal-24                    |     716538 |   1474 ns/op |        85 |        303 |   1.06 |         6108 |    4.86
Benchmark_Hprose2_Marshal-24                     |    3021306 |    368 ns/op |        85 |          0 |   1.11 |        25747 |    0.00
Benchmark_Hprose2_Unmarshal-24                   |    1433833 |    832 ns/op |        85 |        136 |   1.19 |        12221 |    6.12
Benchmark_Protobuf_Marshal-24                    |    1298545 |    935 ns/op |        52 |        144 |   1.21 |         6752 |    6.49
Benchmark_Protobuf_Unmarshal-24                  |     908842 |   1168 ns/op |        52 |        184 |   1.06 |         4725 |    6.35
Benchmark_Gogoprotobuf_Marshal-24                |    7548626 |    156 ns/op |        53 |         64 |   1.18 |        40007 |    2.45
Benchmark_Gogoprotobuf_Unmarshal-24              |    4389771 |    286 ns/op |        53 |         96 |   1.26 |        23265 |    2.99
Benchmark_Gogojsonpb_Marshal-24                  |      81925 |  15127 ns/op |       125 |       3103 |   1.24 |         1030 |    4.87
Benchmark_Gogojsonpb_Unmarshal-24                |      59716 |  19601 ns/op |       125 |       3372 |   1.17 |          749 |    5.81
Benchmark_Colfer_Marshal-24                      |    9303571 |    122 ns/op |        51 |         64 |   1.14 |        47513 |    1.92
Benchmark_Colfer_Unmarshal-24                    |    5323239 |    221 ns/op |        52 |        112 |   1.18 |        27680 |    1.98
Benchmark_Gencode_Marshal-24                     |    4799440 |    252 ns/op |        53 |         80 |   1.21 |        25437 |    3.16
Benchmark_Gencode_Unmarshal-24                   |    4207510 |    274 ns/op |        53 |        112 |   1.16 |        22299 |    2.45
Benchmark_GencodeUnsafe_Marshal-24               |   11602610 |     98 ns/op |        46 |         48 |   1.14 |        53372 |    2.05
Benchmark_GencodeUnsafe_Unmarshal-24             |    6123943 |    195 ns/op |        46 |         96 |   1.20 |        28170 |    2.04
Benchmark_XDR2_Marshal-24                        |    6591649 |    184 ns/op |        60 |         64 |   1.22 |        39549 |    2.89
Benchmark_XDR2_Unmarshal-24                      |    6934723 |    167 ns/op |        60 |         32 |   1.16 |        41608 |    5.22
Benchmark_GoAvro_Marshal-24                      |     369241 |   3029 ns/op |        47 |        728 |   1.12 |         1735 |    4.16
Benchmark_GoAvro_Unmarshal-24                    |     158101 |   7150 ns/op |        47 |       2544 |   1.13 |          743 |    2.81
Benchmark_GoAvro2Text_Marshal-24                 |     294999 |   3840 ns/op |       133 |       1320 |   1.13 |         3947 |    2.91
Benchmark_GoAvro2Text_Unmarshal-24               |     305624 |   3726 ns/op |       133 |        736 |   1.14 |         4089 |    5.06
Benchmark_GoAvro2Binary_Marshal-24               |     959138 |   1121 ns/op |        47 |        464 |   1.08 |         4507 |    2.42
Benchmark_GoAvro2Binary_Unmarshal-24             |     922796 |   1199 ns/op |        47 |        544 |   1.11 |         4337 |    2.20
Benchmark_Ikea_Marshal-24                        |     885784 |   1287 ns/op |        55 |        184 |   1.14 |         4871 |    6.99
Benchmark_Ikea_Unmarshal-24                      |    1000000 |   1014 ns/op |        55 |        160 |   1.01 |         5500 |    6.34
Benchmark_ShamatonMapMsgpack_Marshal-24          |    1433352 |    841 ns/op |        92 |        192 |   1.21 |        13186 |    4.38
Benchmark_ShamatonMapMsgpack_Unmarshal-24        |    1411447 |    902 ns/op |        92 |        168 |   1.27 |        12985 |    5.37
Benchmark_ShamatonArrayMsgpack_Marshal-24        |    1567261 |    755 ns/op |        50 |        160 |   1.18 |         7836 |    4.72
Benchmark_ShamatonArrayMsgpack_Unmarshal-24      |    1666802 |    719 ns/op |        50 |        168 |   1.20 |         8334 |    4.28
Benchmark_ShamatonMapMsgpackgen_Marshal-24       |    4747718 |    251 ns/op |        92 |         96 |   1.19 |        43679 |    2.62
Benchmark_ShamatonMapMsgpackgen_Unmarshal-24     |    2764378 |    435 ns/op |        92 |        112 |   1.20 |        25432 |    3.89
Benchmark_ShamatonArrayMsgpackgen_Marshal-24     |    6674094 |    179 ns/op |        50 |         64 |   1.20 |        33370 |    2.80
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-24   |    5670154 |    262 ns/op |        50 |        112 |   1.49 |        28350 |    2.35
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-24   |     188692 |   5692 ns/op |        55 |        440 |   1.07 |         1037 |   12.94
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-24 |     133633 |   9550 ns/op |        55 |       1184 |   1.28 |          734 |    8.07
Benchmark_Bebop_Marshal-24                       |    7810525 |    153 ns/op |        55 |         64 |   1.20 |        42957 |    2.40
Benchmark_Bebop_Unmarshal-24                     |    9193828 |    126 ns/op |        55 |         32 |   1.16 |        50566 |    3.94
Benchmark_FastJson_Marshal-24                    |    1528498 |    801 ns/op |       133 |        504 |   1.23 |        20451 |    1.59
Benchmark_FastJson_Unmarshal-24                  |     610503 |   1888 ns/op |       133 |       1704 |   1.15 |         8168 |    1.11
Benchmark_Musgo_Marshal-24                       |   10092196 |    122 ns/op |        46 |         48 |   1.24 |        46424 |    2.56
Benchmark_Musgo_Unmarshal-24                     |    5498397 |    215 ns/op |        46 |         96 |   1.18 |        25292 |    2.24
Benchmark_MusgoUnsafe_Marshal-24                 |   10406475 |    114 ns/op |        46 |         48 |   1.19 |        47869 |    2.38
Benchmark_MusgoUnsafe_Unmarshal-24               |   10281271 |    114 ns/op |        46 |         64 |   1.18 |        47293 |    1.80


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_MusgoUnsafe_-24               |   20687746 |    229 ns/op |        92 |        112 |   4.74 |       190327 |    2.05
Benchmark_Bebop_-24                     |   17004353 |    280 ns/op |       110 |         96 |   4.76 |       187047 |    2.92
Benchmark_GencodeUnsafe_-24             |   17726553 |    294 ns/op |        92 |        144 |   5.21 |       163084 |    2.04
Benchmark_Musgo_-24                     |   15590593 |    337 ns/op |        92 |        144 |   5.27 |       143433 |    2.35
Benchmark_Colfer_-24                    |   14626810 |    344 ns/op |       103 |        176 |   5.04 |       150758 |    1.96
Benchmark_XDR2_-24                      |   13526372 |    352 ns/op |       120 |         96 |   4.76 |       162316 |    3.67
Benchmark_ShamatonArrayMsgpackgen_-24   |   12344248 |    442 ns/op |       100 |        176 |   5.46 |       123442 |    2.51
Benchmark_Gogoprotobuf_-24              |   11938397 |    443 ns/op |       106 |        160 |   5.29 |       126547 |    2.77
Benchmark_Gencode_-24                   |    9006950 |    527 ns/op |       106 |        192 |   4.75 |        95473 |    2.75
Benchmark_Msgp_-24                      |    9574260 |    555 ns/op |       194 |        240 |   5.31 |       185740 |    2.31
Benchmark_ShamatonMapMsgpackgen_-24     |    7512096 |    686 ns/op |       184 |        208 |   5.16 |       138222 |    3.30
Benchmark_GotinyNoTime_-24              |    6800561 |    781 ns/op |        96 |        264 |   5.31 |        65285 |    2.96
Benchmark_Gotiny_-24                    |    6533110 |    805 ns/op |        96 |        280 |   5.27 |        62717 |    2.88
Benchmark_Hprose2_-24                   |    4455139 |   1201 ns/op |       170 |        136 |   5.35 |        75942 |    8.83
Benchmark_FlatBuffers_-24               |    5338973 |   1297 ns/op |       190 |        488 |   6.93 |       101611 |    2.66
Benchmark_ShamatonArrayMsgpack_-24      |    3234063 |   1474 ns/op |       100 |        328 |   4.77 |        32340 |    4.50
Benchmark_ShamatonMapMsgpack_-24        |    2844799 |   1743 ns/op |       184 |        360 |   4.96 |        52344 |    4.84
Benchmark_CapNProto2_-24                |    2846391 |   1920 ns/op |       192 |       1724 |   5.47 |        54650 |    1.11
Benchmark_Protobuf_-24                  |    2207387 |   2103 ns/op |       104 |        328 |   4.64 |        22956 |    6.41
Benchmark_Hprose_-24                    |    2247733 |   2180 ns/op |       170 |        622 |   4.90 |        38326 |    3.50
Benchmark_Ikea_-24                      |    1885784 |   2301 ns/op |       110 |        344 |   4.34 |        20743 |    6.69
Benchmark_GoAvro2Binary_-24             |    1881934 |   2320 ns/op |        94 |       1008 |   4.37 |        17690 |    2.30
Benchmark_CapNProto_-24                 |    2616446 |   2672 ns/op |       192 |       4584 |   6.99 |        50235 |    0.58
Benchmark_UgorjiCodecMsgpack_-24        |    1631705 |   2683 ns/op |       182 |       1928 |   4.38 |        29697 |    1.39
Benchmark_FastJson_-24                  |    2139001 |   2689 ns/op |       267 |       2208 |   5.75 |        57239 |    1.22
Benchmark_UgorjiCodecBinc_-24           |    1627550 |   2729 ns/op |       190 |       1944 |   4.44 |        30923 |    1.40
Benchmark_VmihailencoMsgpack_-24        |    1629945 |   2741 ns/op |       184 |        424 |   4.47 |        29990 |    6.46
Benchmark_EasyJson_-24                  |    1721968 |   2801 ns/op |       303 |       1008 |   4.82 |        52210 |    2.78
Benchmark_JsonIter_-24                  |    1564359 |   2898 ns/op |       282 |        416 |   4.53 |        44208 |    6.97
Benchmark_XDR_-24                       |    1202350 |   3719 ns/op |       184 |        670 |   4.47 |        22123 |    5.55
Benchmark_Bson_-24                      |    1275817 |   3815 ns/op |       220 |        600 |   4.87 |        28067 |    6.36
Benchmark_Binary_-24                    |    1057072 |   4284 ns/op |       122 |        680 |   4.53 |        12896 |    6.30
Benchmark_MongoBson_-24                 |    1055451 |   4510 ns/op |       220 |        648 |   4.76 |        23219 |    6.96
Benchmark_Json_-24                      |    1259165 |   6310 ns/op |       303 |        559 |   7.95 |        38177 |   11.29
Benchmark_Sereal_-24                    |     698528 |   7381 ns/op |       264 |       1808 |   5.16 |        18441 |    4.08
Benchmark_GoAvro2Text_-24               |     600623 |   7566 ns/op |       267 |       2056 |   4.54 |        16072 |    3.68
Benchmark_GoAvro_-24                    |     527342 |  10179 ns/op |        94 |       3272 |   5.37 |         4957 |    3.11
Benchmark_SSZNoTimeNoStringNoFloatA_-24 |     322325 |  15242 ns/op |       110 |       1624 |   4.91 |         3545 |    9.39
Benchmark_Gogojsonpb_-24                |     141641 |  34728 ns/op |       251 |       6475 |   4.92 |         3559 |    5.36
Benchmark_Gob_-24                       |     203624 |  42069 ns/op |       327 |       9384 |   8.57 |         6664 |    4.48

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
