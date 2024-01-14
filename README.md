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

```bash
go get -u -t
go test -bench='.*' ./
```

To update the table in the README:

```bash
./stats.sh
```

## Recommendation

If correctness and interoperability are the most
important factors, [protobuf](https://google.golang.org/protobuf) and [json](http://golang.org/pkg/encoding/json/) do the job good.
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

2023-11-16 Results with Go 1.21.4 linux/amd64 with processor `11th Gen Intel(R) Core(TM) i5-11300H @ 3.10GHz`

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_Gotiny_Marshal-8                       |    5175884 |    264 ns/op |        48 |        168 |   1.37 |        24844 |    1.58
Benchmark_Gotiny_Unmarshal-8                     |    7481815 |    145 ns/op |        47 |        112 |   1.09 |        35905 |    1.30
Benchmark_GotinyNoTime_Marshal-8                 |    4763344 |    268 ns/op |        47 |        168 |   1.28 |        22859 |    1.60
Benchmark_GotinyNoTime_Unmarshal-8               |    9295584 |    130 ns/op |        48 |         96 |   1.21 |        44618 |    1.36
Benchmark_Msgp_Marshal-8                         |   12086298 |    110 ns/op |        97 |        128 |   1.34 |       117237 |    0.86
Benchmark_Msgp_Unmarshal-8                       |    6761966 |    186 ns/op |        97 |        112 |   1.26 |        65591 |    1.67
Benchmark_VmihailencoMsgpack_Marshal-8           |    2206149 |    499 ns/op |        92 |        264 |   1.10 |        20296 |    1.89
Benchmark_VmihailencoMsgpack_Unmarshal-8         |    1907919 |    600 ns/op |        92 |        160 |   1.15 |        17552 |    3.75
Benchmark_Json_Marshal-8                         |    1505756 |    770 ns/op |       151 |        208 |   1.16 |        22842 |    3.70
Benchmark_Json_Unmarshal-8                       |     674160 |   1670 ns/op |       151 |        351 |   1.13 |        10220 |    4.76
Benchmark_JsonIter_Marshal-8                     |    2389435 |    519 ns/op |       141 |        200 |   1.24 |        33738 |    2.60
Benchmark_JsonIter_Unmarshal-8                   |    1742654 |    677 ns/op |       141 |        216 |   1.18 |        24623 |    3.14
Benchmark_EasyJson_Marshal-8                     |    1491632 |    845 ns/op |       151 |        896 |   1.26 |        22613 |    0.94
Benchmark_EasyJson_Unmarshal-8                   |    2055382 |    535 ns/op |       151 |        112 |   1.10 |        31159 |    4.78
Benchmark_Bson_Marshal-8                         |    1733908 |    675 ns/op |       110 |        376 |   1.17 |        19072 |    1.80
Benchmark_Bson_Unmarshal-8                       |    1340866 |    864 ns/op |       110 |        224 |   1.16 |        14749 |    3.86
Benchmark_MongoBson_Marshal-8                    |    1374638 |    923 ns/op |       110 |        240 |   1.27 |        15121 |    3.85
Benchmark_MongoBson_Unmarshal-8                  |    1208486 |    999 ns/op |       110 |        408 |   1.21 |        13293 |    2.45
Benchmark_Gob_Marshal-8                          |     452556 |   3222 ns/op |       162 |       1744 |   1.46 |         7354 |    1.85
Benchmark_Gob_Unmarshal-8                        |      77752 |  16301 ns/op |       162 |       7720 |   1.27 |         1264 |    2.11
Benchmark_XDR_Marshal-8                          |    1408309 |    913 ns/op |        92 |        440 |   1.29 |        12956 |    2.08
Benchmark_XDR_Unmarshal-8                        |    1724904 |    664 ns/op |        92 |        231 |   1.15 |        15869 |    2.88
Benchmark_UgorjiCodecMsgpack_Marshal-8           |    1722421 |    764 ns/op |        91 |       1240 |   1.32 |        15674 |    0.62
Benchmark_UgorjiCodecMsgpack_Unmarshal-8         |    1552963 |    844 ns/op |        91 |        688 |   1.31 |        14131 |    1.23
Benchmark_UgorjiCodecBinc_Marshal-8              |    1458865 |    872 ns/op |        95 |       1256 |   1.27 |        13859 |    0.69
Benchmark_UgorjiCodecBinc_Unmarshal-8            |    1684272 |    910 ns/op |        95 |        688 |   1.53 |        16000 |    1.32
Benchmark_Sereal_Marshal-8                       |     818120 |   1693 ns/op |       132 |        832 |   1.39 |        10799 |    2.03
Benchmark_Sereal_Unmarshal-8                     |     523800 |   1940 ns/op |       132 |        976 |   1.02 |         6914 |    1.99
Benchmark_Binary_Marshal-8                       |    1156994 |   1030 ns/op |        61 |        360 |   1.19 |         7057 |    2.86
Benchmark_Binary_Unmarshal-8                     |    1309048 |    883 ns/op |        61 |        320 |   1.16 |         7985 |    2.76
Benchmark_FlatBuffers_Marshal-8                  |    2113926 |    565 ns/op |        95 |        376 |   1.20 |        20141 |    1.51
Benchmark_FlatBuffers_Unmarshal-8                |    6082993 |    166 ns/op |        95 |        112 |   1.01 |        57946 |    1.48
Benchmark_CapNProto_Marshal-8                    |     980739 |   1092 ns/op |        96 |       4392 |   1.07 |         9415 |    0.25
Benchmark_CapNProto_Unmarshal-8                  |    3516151 |    320 ns/op |        96 |        192 |   1.13 |        33755 |    1.67
Benchmark_CapNProto2_Marshal-8                   |    1310677 |   1129 ns/op |        96 |       1452 |   1.48 |        12582 |    0.78
Benchmark_CapNProto2_Unmarshal-8                 |    3776460 |    410 ns/op |        96 |        272 |   1.55 |        36254 |    1.51
Benchmark_Hprose_Marshal-8                       |    2180806 |    499 ns/op |        85 |        390 |   1.09 |        18589 |    1.28
Benchmark_Hprose_Unmarshal-8                     |    1823271 |    746 ns/op |        85 |        303 |   1.36 |        15554 |    2.46
Benchmark_Hprose2_Marshal-8                      |    4629981 |    246 ns/op |        85 |          0 |   1.14 |        39475 |    0.00
Benchmark_Hprose2_Unmarshal-8                    |    3439027 |    346 ns/op |        85 |        136 |   1.19 |        29328 |    2.55
Benchmark_Protobuf_Marshal-8                     |    2780270 |    417 ns/op |        52 |        144 |   1.16 |        14457 |    2.90
Benchmark_Protobuf_Unmarshal-8                   |    2198923 |    503 ns/op |        52 |        184 |   1.11 |        11434 |    2.74
Benchmark_Pulsar_Marshal-8                       |    3066033 |    438 ns/op |        51 |        304 |   1.35 |        15820 |    1.44
Benchmark_Pulsar_Unmarshal-8                     |    3018387 |    428 ns/op |        51 |        256 |   1.29 |        15547 |    1.67
Benchmark_Gogoprotobuf_Marshal-8                 |   11825176 |     86 ns/op |        53 |         64 |   1.02 |        62673 |    1.35
Benchmark_Gogoprotobuf_Unmarshal-8               |    7308174 |    140 ns/op |        53 |         96 |   1.02 |        38733 |    1.46
Benchmark_Gogojsonpb_Marshal-8                   |     180114 |   7772 ns/op |       125 |       3095 |   1.40 |         2264 |    2.51
Benchmark_Gogojsonpb_Unmarshal-8                 |     123753 |   8544 ns/op |       125 |       3375 |   1.06 |         1554 |    2.53
Benchmark_Colfer_Marshal-8                       |   15535346 |     84 ns/op |        51 |         64 |   1.31 |        79339 |    1.32
Benchmark_Colfer_Unmarshal-8                     |   11936880 |    122 ns/op |        52 |        112 |   1.46 |        62071 |    1.09
Benchmark_Gencode_Marshal-8                      |   12112273 |    107 ns/op |        53 |         80 |   1.30 |        64195 |    1.34
Benchmark_Gencode_Unmarshal-8                    |    8852018 |    142 ns/op |        53 |        112 |   1.26 |        46915 |    1.27
Benchmark_GencodeUnsafe_Marshal-8                |   21924096 |     62 ns/op |        46 |         48 |   1.37 |       100850 |    1.30
Benchmark_GencodeUnsafe_Unmarshal-8              |   10802174 |    110 ns/op |        46 |         96 |   1.19 |        49690 |    1.15
Benchmark_XDR2_Marshal-8                         |   12868144 |    110 ns/op |        60 |         64 |   1.42 |        77208 |    1.72
Benchmark_XDR2_Unmarshal-8                       |   17968026 |     66 ns/op |        60 |         32 |   1.19 |       107808 |    2.07
Benchmark_GoAvro_Marshal-8                       |    1000000 |   1246 ns/op |        47 |        584 |   1.25 |         4700 |    2.13
Benchmark_GoAvro_Unmarshal-8                     |     406524 |   3571 ns/op |        47 |       2312 |   1.45 |         1910 |    1.54
Benchmark_GoAvro2Text_Marshal-8                  |     570015 |   2303 ns/op |       133 |       1320 |   1.31 |         7626 |    1.74
Benchmark_GoAvro2Text_Unmarshal-8                |     846238 |   1588 ns/op |       133 |        736 |   1.34 |        11322 |    2.16
Benchmark_GoAvro2Binary_Marshal-8                |    1801123 |    672 ns/op |        47 |        464 |   1.21 |         8465 |    1.45
Benchmark_GoAvro2Binary_Unmarshal-8              |    1789248 |    589 ns/op |        47 |        544 |   1.05 |         8409 |    1.08
Benchmark_Ikea_Marshal-8                         |    2166458 |    531 ns/op |        55 |        184 |   1.15 |        11915 |    2.89
Benchmark_Ikea_Unmarshal-8                       |    2764864 |    429 ns/op |        55 |        160 |   1.19 |        15206 |    2.68
Benchmark_ShamatonMapMsgpack_Marshal-8           |    2728044 |    462 ns/op |        92 |        192 |   1.26 |        25098 |    2.41
Benchmark_ShamatonMapMsgpack_Unmarshal-8         |    3132741 |    374 ns/op |        92 |        168 |   1.17 |        28821 |    2.23
Benchmark_ShamatonArrayMsgpack_Marshal-8         |    3066362 |    374 ns/op |        50 |        160 |   1.15 |        15331 |    2.34
Benchmark_ShamatonArrayMsgpack_Unmarshal-8       |    3704914 |    351 ns/op |        50 |        168 |   1.30 |        18524 |    2.09
Benchmark_ShamatonMapMsgpackgen_Marshal-8        |    9664315 |    132 ns/op |        92 |         96 |   1.28 |        88911 |    1.38
Benchmark_ShamatonMapMsgpackgen_Unmarshal-8      |    6034456 |    224 ns/op |        92 |        112 |   1.35 |        55516 |    2.00
Benchmark_ShamatonArrayMsgpackgen_Marshal-8      |   10569614 |    100 ns/op |        50 |         64 |   1.06 |        52848 |    1.57
Benchmark_ShamatonArrayMsgpackgen_Unmarshal-8    |    9751200 |    146 ns/op |        50 |        112 |   1.43 |        48756 |    1.31
Benchmark_SSZNoTimeNoStringNoFloatA_Marshal-8    |     495163 |   2378 ns/op |        55 |        440 |   1.18 |         2723 |    5.40
Benchmark_SSZNoTimeNoStringNoFloatA_Unmarshal-8  |     278251 |   3984 ns/op |        55 |       1184 |   1.11 |         1530 |    3.36
Benchmark_Bebop_200sc_Marshal-8                  |   14961906 |     81 ns/op |        55 |         64 |   1.21 |        82290 |    1.27
Benchmark_Bebop_200sc_Unmarshal-8                |   20842504 |     61 ns/op |        55 |         32 |   1.27 |       114633 |    1.91
Benchmark_Bebop_Wellquite_Marshal-8              |   17579114 |     72 ns/op |        55 |         64 |   1.28 |        96685 |    1.13
Benchmark_Bebop_Wellquite_Unmarshal-8            |   20299090 |     66 ns/op |        55 |         32 |   1.34 |       111644 |    2.06
Benchmark_FastJson_Marshal-8                     |    3281961 |    382 ns/op |       133 |        504 |   1.26 |        43912 |    0.76
Benchmark_FastJson_Unmarshal-8                   |    1261986 |    919 ns/op |       133 |       1704 |   1.16 |        16885 |    0.54
Benchmark_BENC_Marshal-8                         |   21525518 |     65 ns/op |        51 |         64 |   1.41 |       109780 |    1.02
Benchmark_BENC_Unmarshal-8                       |   23009800 |     60 ns/op |        51 |         32 |   1.39 |       117349 |    1.89
Benchmark_BENCUnsafe_Marshal-8                   |   20615397 |     52 ns/op |        51 |         64 |   1.08 |       105138 |    0.82
Benchmark_BENCUnsafe_Unmarshal-8                 |   54961459 |     21 ns/op |        51 |          0 |   1.21 |       280303 |    0.00
Benchmark_BENCUnsafePre_Marshal-8                |   50289554 |     24 ns/op |        51 |          0 |   1.22 |       256476 |    0.00
Benchmark_BENCUnsafePre_Unmarshal-8              |   59493313 |     19 ns/op |        51 |          0 |   1.18 |       303415 |    0.00
Benchmark_MUS_Marshal-8                          |   18902049 |     64 ns/op |        46 |         48 |   1.22 |        86949 |    1.35
Benchmark_MUS_Unmarshal-8                        |   15403171 |     76 ns/op |        46 |         32 |   1.18 |        70854 |    2.39
Benchmark_MUSUnsafe_Marshal-8                    |   20805338 |     65 ns/op |        49 |         64 |   1.35 |       101946 |    1.02
Benchmark_MUSUnsafe_Unmarshal-8                  |   35980262 |     29 ns/op |        49 |          0 |   1.06 |       176303 |    0.00


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
Benchmark_BENCUnsafePre_-8              |  109782867 |     44 ns/op |       102 |          0 |   4.84 |      1119785 | Zero-Allocs
Benchmark_BENCUnsafe_-8                 |   75576856 |     74 ns/op |       102 |         64 |   5.64 |       770883 |    1.17
Benchmark_MUSUnsafe_-8                  |   56785600 |     94 ns/op |        98 |         64 |   5.37 |       556498 |    1.48
Benchmark_BENC_-8                       |   44535318 |    125 ns/op |       102 |         96 |   5.61 |       454260 |    1.31
Benchmark_Bebop_Wellquite_-8            |   37878204 |    138 ns/op |       110 |         96 |   5.25 |       416660 |    1.44
Benchmark_MUS_-8                        |   34305220 |    140 ns/op |        92 |         80 |   4.84 |       315608 |    1.76
Benchmark_Bebop_200sc_-8                |   35804410 |    142 ns/op |       110 |         96 |   5.09 |       393848 |    1.48
Benchmark_GencodeUnsafe_-8              |   32726270 |    172 ns/op |        92 |        144 |   5.65 |       301081 |    1.20
Benchmark_XDR2_-8                       |   30836170 |    176 ns/op |       120 |         96 |   5.44 |       370034 |    1.84
Benchmark_Colfer_-8                     |   27472226 |    206 ns/op |       103 |        176 |   5.68 |       283156 |    1.18
Benchmark_Gogoprotobuf_-8               |   19133350 |    226 ns/op |       106 |        160 |   4.33 |       202813 |    1.42
Benchmark_ShamatonArrayMsgpackgen_-8    |   20320814 |    247 ns/op |       100 |        176 |   5.03 |       203208 |    1.41
Benchmark_Gencode_-8                    |   20964291 |    249 ns/op |       106 |        192 |   5.23 |       222221 |    1.30
Benchmark_Msgp_-8                       |   18848264 |    297 ns/op |       194 |        240 |   5.61 |       365656 |    1.24
Benchmark_ShamatonMapMsgpackgen_-8      |   15698771 |    356 ns/op |       184 |        208 |   5.60 |       288857 |    1.71
Benchmark_GotinyNoTime_-8               |   14058928 |    398 ns/op |        95 |        264 |   5.61 |       134951 |    1.51
Benchmark_Gotiny_-8                     |   12657699 |    410 ns/op |        95 |        280 |   5.19 |       121501 |    1.47
Benchmark_Hprose2_-8                    |    8069008 |    592 ns/op |       170 |        136 |   4.78 |       137608 |    4.36
Benchmark_ShamatonArrayMsgpack_-8       |    6771276 |    725 ns/op |       100 |        328 |   4.91 |        67712 |    2.21
Benchmark_FlatBuffers_-8                |    8196919 |    732 ns/op |       190 |        488 |   6.00 |       156184 |    1.50
Benchmark_ShamatonMapMsgpack_-8         |    5860785 |    836 ns/op |       184 |        360 |   4.90 |       107838 |    2.32
Benchmark_Pulsar_-8                     |    6084420 |    867 ns/op |       103 |        560 |   5.28 |        62736 |    1.55
Benchmark_Protobuf_-8                   |    4979193 |    920 ns/op |       104 |        328 |   4.59 |        51783 |    2.81
Benchmark_Ikea_-8                       |    4931322 |    961 ns/op |       110 |        344 |   4.74 |        54244 |    2.79
Benchmark_VmihailencoMsgpack_-8         |    4114068 |   1099 ns/op |       184 |        424 |   4.52 |        75698 |    2.59
Benchmark_JsonIter_-8                   |    4132089 |   1196 ns/op |       282 |        416 |   4.94 |       116731 |    2.88
Benchmark_Hprose_-8                     |    4004077 |   1245 ns/op |       170 |        693 |   4.99 |        68289 |    1.80
Benchmark_GoAvro2Binary_-8              |    3590371 |   1261 ns/op |        94 |       1008 |   4.53 |        33749 |    1.25
Benchmark_FastJson_-8                   |    4543947 |   1302 ns/op |       267 |       2208 |   5.92 |       121596 |    0.59
Benchmark_EasyJson_-8                   |    3547014 |   1380 ns/op |       303 |       1008 |   4.90 |       107545 |    1.37
Benchmark_CapNProto_-8                  |    4496890 |   1412 ns/op |       192 |       4584 |   6.35 |        86340 |    0.31
Benchmark_Bson_-8                       |    3074774 |   1539 ns/op |       220 |        600 |   4.73 |        67645 |    2.57
Benchmark_CapNProto2_-8                 |    5087137 |   1539 ns/op |       192 |       1724 |   7.83 |        97673 |    0.89
Benchmark_XDR_-8                        |    3133213 |   1578 ns/op |       184 |        671 |   4.95 |        57651 |    2.35
Benchmark_UgorjiCodecMsgpack_-8         |    3275384 |   1609 ns/op |       182 |       1928 |   5.27 |        59611 |    0.83
Benchmark_UgorjiCodecBinc_-8            |    3143137 |   1783 ns/op |       190 |       1944 |   5.60 |        59719 |    0.92
Benchmark_Binary_-8                     |    2466042 |   1913 ns/op |       122 |        680 |   4.72 |        30085 |    2.81
Benchmark_MongoBson_-8                  |    2583124 |   1923 ns/op |       220 |        648 |   4.97 |        56828 |    2.97
Benchmark_Json_-8                       |    2179916 |   2440 ns/op |       303 |        559 |   5.32 |        66116 |    4.37
Benchmark_Sereal_-8                     |    1341920 |   3633 ns/op |       264 |       1808 |   4.88 |        35426 |    2.01
Benchmark_GoAvro2Text_-8                |    1416253 |   3891 ns/op |       267 |       2056 |   5.51 |        37898 |    1.89
Benchmark_GoAvro_-8                     |    1406524 |   4817 ns/op |        94 |       2896 |   6.78 |        13221 |    1.66
Benchmark_SSZNoTimeNoStringNoFloatA_-8  |     773414 |   6362 ns/op |       110 |       1624 |   4.92 |         8507 |    3.92
Benchmark_Gogojsonpb_-8                 |     303867 |  16316 ns/op |       251 |       6470 |   4.96 |         7636 |    2.52
Benchmark_Gob_-8                        |     530308 |  19523 ns/op |       325 |       9464 |  10.35 |        17240 |    2.06

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

Additionally, while not a correctness issue, FlatBuffers, ProtoBuffers, Cap'N'Proto, ikeapack and MUS
do not support time types directly. In the benchmarks an int64 value is used to hold a UnixNano timestamp.

Bebop (both libraries, by nature of the format) natively supports times rounded to 100ns ticks, and this is what is currently benchmarked (a unix nano timestamp is another valid approach).

MUSUnsafe results were obtained using the mus-go unsafe package. With this 
package, after decoding a byte slice into a string, any change to this slice 
will change the contents of the string. In such cases, the slice can be reused 
only after processing the received result.

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
