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
- [github.com/nazarifard/fastape](https://github.com/nazarifard/fastape)

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
If speed matters, [fastape](https://github.com/nazarifard/fastape) and [MUS](https://github.com/mus-format/mus-go) and [BENC](https://github.com/deneonet/benc) are probably the best choice.

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

2024-06-11 Results with Go 1.22.4 linux/amd64 with processor `Intel(R) Core(TM) i7-3537U CPU @ 2.00GHz`

benchmark                                        | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
-------------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
BenchmarkSerializers/marshal/gotiny-4            |    2122971 |    552 ns/op |        47 |        168 |   1.17 |        10188 |    3.29
BenchmarkSerializers/unmarshal/gotiny-4          |    5377658 |    218 ns/op |        32 |          2 |   1.18 |        17208 |  109.25
BenchmarkSerializers/marshal/msgp-4              |    5777872 |    201 ns/op |        97 |        128 |   1.16 |        56045 |    1.57
BenchmarkSerializers/unmarshal/msgp-4            |    4178436 |    255 ns/op |        32 |          2 |   1.07 |        13370 |  127.85
BenchmarkSerializers/marshal/msgpack-4           |     985917 |   1057 ns/op |        92 |        264 |   1.04 |         9070 |    4.00
BenchmarkSerializers/unmarshal/msgpack-4         |     756810 |   1462 ns/op |        80 |          3 |   1.11 |         6054 |  487.33
BenchmarkSerializers/marshal/json-4              |     484378 |   2127 ns/op |       151 |        208 |   1.03 |         7352 |   10.23
BenchmarkSerializers/unmarshal/json-4            |     230168 |   4720 ns/op |       248 |          6 |   1.09 |         5708 |  786.67
BenchmarkSerializers/marshal/jsoniter-4          |    1002337 |   1206 ns/op |       141 |        200 |   1.21 |        14173 |    6.03
BenchmarkSerializers/unmarshal/jsoniter-4        |     709945 |   1762 ns/op |       136 |          5 |   1.25 |         9655 |  352.40
BenchmarkSerializers/marshal/easyjson-4          |     692306 |   1744 ns/op |       151 |        976 |   1.21 |        10502 |    1.79
BenchmarkSerializers/unmarshal/easyjson-4        |     835852 |   1374 ns/op |        32 |          2 |   1.15 |         2674 |  687.00
BenchmarkSerializers/marshal/bson-4              |     710941 |   1424 ns/op |       110 |        376 |   1.01 |         7820 |    3.79
BenchmarkSerializers/unmarshal/bson-4            |     469126 |   2285 ns/op |       144 |         18 |   1.07 |         6755 |  126.94
BenchmarkSerializers/marshal/mongobson-4         |     509228 |   2098 ns/op |       110 |        240 |   1.07 |         5601 |    8.74
BenchmarkSerializers/unmarshal/mongobson-4       |     517167 |   2373 ns/op |       328 |         14 |   1.23 |        16963 |  169.50
BenchmarkSerializers/marshal/gob-4               |     141312 |   7625 ns/op |       172 |       1744 |   1.08 |         2439 |    4.37
BenchmarkSerializers/unmarshal/gob-4             |      31680 |  38385 ns/op |      7656 |        203 |   1.22 |        24254 |  189.09
BenchmarkSerializers/marshal/davecgh/xdr-4       |     548881 |   1918 ns/op |        92 |        392 |   1.05 |         5049 |    4.89
BenchmarkSerializers/unmarshal/davecgh/xdr-4     |     722668 |   1607 ns/op |       152 |         10 |   1.16 |        10984 |  160.70
BenchmarkSerializers/marshal/ugorji/msgpack-4    |     713061 |   1734 ns/op |        91 |       1240 |   1.24 |         6488 |    1.40
BenchmarkSerializers/unmarshal/ugorji/msgpack-4  |     698718 |   1687 ns/op |       608 |          3 |   1.18 |        42482 |  562.33
BenchmarkSerializers/marshal/ugorji/binc-4       |     648315 |   1817 ns/op |        95 |       1256 |   1.18 |         6158 |    1.45
BenchmarkSerializers/unmarshal/ugorji/binc-4     |     594816 |   1930 ns/op |       832 |          6 |   1.15 |        49488 |  321.67
BenchmarkSerializers/marshal/sereal-4            |     263485 |   4206 ns/op |       142 |       1104 |   1.11 |         3741 |    3.81
BenchmarkSerializers/unmarshal/sereal-4          |     239680 |   4376 ns/op |       896 |         33 |   1.05 |        21475 |  132.61
BenchmarkSerializers/marshal/alecthomas/binary-4 |     493554 |   2336 ns/op |        61 |        360 |   1.15 |         3010 |    6.49
BenchmarkSerializers/unmarshal/alecthomas/binary-4 |     562761 |   2148 ns/op |       240 |         19 |   1.21 |        13506 |  113.05
BenchmarkSerializers/marshal/flatbuffers-4       |    1009303 |   1192 ns/op |        95 |        376 |   1.20 |         9604 |    3.17
BenchmarkSerializers/unmarshal/flatbuffers-4     |    4762442 |    244 ns/op |        32 |          2 |   1.17 |        15239 |  122.45
BenchmarkSerializers/marshal/capnproto-4         |     282118 |   4239 ns/op |        96 |       4392 |   1.20 |         2708 |    0.97
BenchmarkSerializers/unmarshal/capnproto-4       |    2070158 |    582 ns/op |       112 |          5 |   1.21 |        23185 |  116.58
BenchmarkSerializers/marshal/hprose-4            |     738270 |   1383 ns/op |        85 |        325 |   1.02 |         6295 |    4.26
BenchmarkSerializers/unmarshal/hprose-4          |     762679 |   1507 ns/op |       224 |          9 |   1.15 |        17084 |  167.44
BenchmarkSerializers/marshal/hprose2-4           |    1630219 |    725 ns/op |        85 |          0 |   1.18 |        13895 |    0.00
BenchmarkSerializers/unmarshal/hprose2-4         |    1751545 |    649 ns/op |        56 |          3 |   1.14 |         9808 |  216.53
BenchmarkSerializers/marshal/dedis/protobuf-4    |    1235511 |    933 ns/op |        52 |        144 |   1.15 |         6424 |    6.48
BenchmarkSerializers/unmarshal/dedis/protobuf-4  |     839454 |   1199 ns/op |       104 |          9 |   1.01 |         8730 |  133.22
BenchmarkSerializers/marshal/pulsar-4            |    1343660 |    904 ns/op |        51 |        304 |   1.22 |         6923 |    2.98
BenchmarkSerializers/unmarshal/pulsar-4          |    1527993 |    758 ns/op |       160 |          6 |   1.16 |        24447 |  126.43
BenchmarkSerializers/marshal/gogo/protobuf-4     |    6060012 |    195 ns/op |        53 |         64 |   1.18 |        32118 |    3.05
BenchmarkSerializers/unmarshal/gogo/protobuf-4   |    4428109 |    270 ns/op |        32 |          2 |   1.20 |        14169 |  135.20
BenchmarkSerializers/marshal/gogo/jsonpb-4       |      67770 |  16636 ns/op |       125 |       2755 |   1.13 |          852 |    6.04
BenchmarkSerializers/unmarshal/gogo/jsonpb-4     |      55772 |  21796 ns/op |      3159 |         55 |   1.22 |        17618 |  396.29
BenchmarkSerializers/marshal/colfer-4            |    5916736 |    188 ns/op |        51 |         64 |   1.11 |        30222 |    2.94
BenchmarkSerializers/unmarshal/colfer-4          |    5969661 |    201 ns/op |        32 |          2 |   1.20 |        19102 |  100.60
BenchmarkSerializers/marshal/gencode-4           |    7127251 |    164 ns/op |        53 |         16 |   1.17 |        37774 |   10.29
BenchmarkSerializers/unmarshal/gencode-4         |    5899357 |    196 ns/op |        32 |          2 |   1.16 |        18877 |   98.25
BenchmarkSerializers/marshal/gencode/unsafe-4    |   15469035 |     77 ns/op |        46 |          0 |   1.19 |        71157 |    0.00
BenchmarkSerializers/unmarshal/gencode/unsafe-4  |    6674700 |    174 ns/op |        32 |          2 |   1.17 |        21359 |   87.30
BenchmarkSerializers/marshal/calmh/xdr-4         |    4938007 |    245 ns/op |        60 |         64 |   1.21 |        29628 |    3.84
BenchmarkSerializers/unmarshal/calmh/xdr-4       |    5513344 |    205 ns/op |        32 |          2 |   1.13 |        17642 |  102.90
BenchmarkSerializers/marshal/goavro-4            |     390975 |   2849 ns/op |        47 |        584 |   1.11 |         1837 |    4.88
BenchmarkSerializers/unmarshal/goavro-4          |     148034 |   7339 ns/op |      2232 |         52 |   1.09 |        33041 |  141.13
BenchmarkSerializers/marshal/avro2/text-4        |     234855 |   4686 ns/op |       133 |       1320 |   1.10 |         3140 |    3.55
BenchmarkSerializers/unmarshal/avro2/text-4      |     284338 |   3892 ns/op |       656 |         30 |   1.11 |        18652 |  129.73
BenchmarkSerializers/marshal/avro2/binary-4      |     853518 |   1368 ns/op |        47 |        464 |   1.17 |         4011 |    2.95
BenchmarkSerializers/unmarshal/avro2/binary-4    |     698176 |   1442 ns/op |       464 |         10 |   1.01 |        32395 |  144.20
BenchmarkSerializers/marshal/ikea-4              |    1530195 |    784 ns/op |        55 |         72 |   1.20 |         8416 |   10.89
BenchmarkSerializers/unmarshal/ikea-4            |    1265690 |    960 ns/op |        96 |         10 |   1.22 |        12150 |   96.01
BenchmarkSerializers/marshal/shamaton/msgpack/map-4 |    1081304 |   1130 ns/op |        92 |        192 |   1.22 |         9947 |    5.89
BenchmarkSerializers/unmarshal/shamaton/msgpack/map-4 |    1343259 |    892 ns/op |        88 |          4 |   1.20 |        11820 |  223.18
BenchmarkSerializers/marshal/shamaton/msgpack/array-4 |    1269360 |    903 ns/op |        50 |        160 |   1.15 |         6346 |    5.65
BenchmarkSerializers/unmarshal/shamaton/msgpack/array-4 |    1623724 |    697 ns/op |        88 |          4 |   1.13 |        14288 |  174.28
BenchmarkSerializers/marshal/shamaton/msgpackgen/map-4 |    2844409 |    425 ns/op |        92 |        176 |   1.21 |        26168 |    2.42
BenchmarkSerializers/unmarshal/shamaton/msgpackgen/map-4 |    2052736 |    570 ns/op |       112 |          3 |   1.17 |        22990 |  190.13
BenchmarkSerializers/marshal/shamaton/msgpackgen/array-4 |    3426714 |    342 ns/op |        50 |        144 |   1.17 |        17133 |    2.38
BenchmarkSerializers/unmarshal/shamaton/msgpackgen/array-4 |    3412923 |    355 ns/op |       112 |          3 |   1.21 |        38224 |  118.40
BenchmarkSerializers/marshal/ssz-4               |     225304 |   4895 ns/op |        55 |        416 |   1.10 |         1239 |   11.77
BenchmarkSerializers/unmarshal/ssz-4             |     359684 |   3291 ns/op |       264 |         27 |   1.18 |         9495 |  121.89
BenchmarkSerializers/marshal/200sc/bebop-4       |   14792242 |     76 ns/op |        55 |          0 |   1.14 |        81357 |    0.00
BenchmarkSerializers/unmarshal/200sc/bebop-4     |    6694594 |    179 ns/op |        32 |          2 |   1.20 |        21422 |   89.70
BenchmarkSerializers/marshal/wellquite/bebop-4   |   14358042 |     84 ns/op |        55 |          0 |   1.21 |        78969 |    0.00
BenchmarkSerializers/unmarshal/wellquite/bebop-4 |    5627302 |    188 ns/op |        32 |          2 |   1.06 |        18007 |   94.25
BenchmarkSerializers/marshal/fastjson-4          |     506050 |   2413 ns/op |       133 |       1360 |   1.22 |         6770 |    1.77
BenchmarkSerializers/unmarshal/fastjson-4        |     354904 |   3452 ns/op |      1800 |         11 |   1.23 |        63882 |  313.82
BenchmarkSerializers/marshal/benc-4              |    7363042 |    145 ns/op |        51 |         64 |   1.07 |        37551 |    2.28
BenchmarkSerializers/unmarshal/benc-4            |    5789518 |    173 ns/op |        32 |          2 |   1.00 |        18526 |   86.60
BenchmarkSerializers/marshal/benc/usafe-4        |    8139350 |    148 ns/op |        51 |         64 |   1.20 |        41510 |    2.31
BenchmarkSerializers/unmarshal/benc/usafe-4      |   18498962 |     61 ns/op |         0 |          0 |   1.14 |            0 |    0.00
BenchmarkSerializers/marshal/mus-4               |    7359068 |    164 ns/op |        46 |         48 |   1.21 |        33851 |    3.43
BenchmarkSerializers/unmarshal/mus-4             |    5190254 |    226 ns/op |        32 |          2 |   1.17 |        16608 |  113.10
BenchmarkSerializers/marshal/mus/unsafe-4        |    7395588 |    158 ns/op |        49 |         64 |   1.17 |        36238 |    2.47
BenchmarkSerializers/unmarshal/mus/unsafe-4      |   12255289 |     94 ns/op |         0 |          0 |   1.16 |            0 |    0.00
BenchmarkSerializers/marshal/baseline-4          |   19410847 |     61 ns/op |        47 |          0 |   1.20 |        91230 |    0.00
BenchmarkSerializers/unmarshal/baseline-4        |   24976167 |     47 ns/op |         0 |          0 |   1.18 |            0 |    0.00
BenchmarkSerializers/marshal/fastape-4           |    6123601 |    195 ns/op |        55 |         64 |   1.20 |        33679 |    3.06
BenchmarkSerializers/unmarshal/fastape-4         |   13608326 |     85 ns/op |         0 |          0 |   1.16 |            0 |    0.00


Totals:


benchmark                               | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc
----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------
BenchmarkSerializers/unmarshal/baseline-4 |   24976167 |     47 ns/op |         0 |          0 |   1.18 |            0 | Zero-Allocs
BenchmarkSerializers/unmarshal/benc/usafe-4 |   18498962 |     61 ns/op |         0 |          0 |   1.14 |            0 | Zero-Allocs
BenchmarkSerializers/marshal/baseline-4 |   19410847 |     61 ns/op |        47 |          0 |   1.20 |        91230 | Zero-Allocs
BenchmarkSerializers/marshal/200sc/bebop-4 |   14792242 |     76 ns/op |        55 |          0 |   1.14 |        81357 | Zero-Allocs
BenchmarkSerializers/marshal/gencode/unsafe-4 |   15469035 |     77 ns/op |        46 |          0 |   1.19 |        71157 | Zero-Allocs
BenchmarkSerializers/marshal/wellquite/bebop-4 |   14358042 |     84 ns/op |        55 |          0 |   1.21 |        78969 | Zero-Allocs
BenchmarkSerializers/unmarshal/fastape-4 |   13608326 |     85 ns/op |         0 |          0 |   1.16 |            0 | Zero-Allocs
BenchmarkSerializers/unmarshal/mus/unsafe-4 |   12255289 |     94 ns/op |         0 |          0 |   1.16 |            0 | Zero-Allocs
BenchmarkSerializers/marshal/benc-4     |    7363042 |    145 ns/op |        51 |         64 |   1.07 |        37551 |    2.28
BenchmarkSerializers/marshal/benc/usafe-4 |    8139350 |    148 ns/op |        51 |         64 |   1.20 |        41510 |    2.31
BenchmarkSerializers/marshal/mus/unsafe-4 |    7395588 |    158 ns/op |        49 |         64 |   1.17 |        36238 |    2.47
BenchmarkSerializers/marshal/mus-4      |    7359068 |    164 ns/op |        46 |         48 |   1.21 |        33851 |    3.43
BenchmarkSerializers/marshal/gencode-4  |    7127251 |    164 ns/op |        53 |         16 |   1.17 |        37774 |   10.29
BenchmarkSerializers/unmarshal/benc-4   |    5789518 |    173 ns/op |        32 |          2 |   1.00 |        18526 |   86.60
BenchmarkSerializers/unmarshal/gencode/unsafe-4 |    6674700 |    174 ns/op |        32 |          2 |   1.17 |        21359 |   87.30
BenchmarkSerializers/unmarshal/200sc/bebop-4 |    6694594 |    179 ns/op |        32 |          2 |   1.20 |        21422 |   89.70
BenchmarkSerializers/marshal/colfer-4   |    5916736 |    188 ns/op |        51 |         64 |   1.11 |        30222 |    2.94
BenchmarkSerializers/unmarshal/wellquite/bebop-4 |    5627302 |    188 ns/op |        32 |          2 |   1.06 |        18007 |   94.25
BenchmarkSerializers/marshal/gogo/protobuf-4 |    6060012 |    195 ns/op |        53 |         64 |   1.18 |        32118 |    3.05
BenchmarkSerializers/marshal/fastape-4  |    6123601 |    195 ns/op |        55 |         64 |   1.20 |        33679 |    3.06
BenchmarkSerializers/unmarshal/gencode-4 |    5899357 |    196 ns/op |        32 |          2 |   1.16 |        18877 |   98.25
BenchmarkSerializers/marshal/msgp-4     |    5777872 |    201 ns/op |        97 |        128 |   1.16 |        56045 |    1.57
BenchmarkSerializers/unmarshal/colfer-4 |    5969661 |    201 ns/op |        32 |          2 |   1.20 |        19102 |  100.60
BenchmarkSerializers/unmarshal/calmh/xdr-4 |    5513344 |    205 ns/op |        32 |          2 |   1.13 |        17642 |  102.90
BenchmarkSerializers/unmarshal/gotiny-4 |    5377658 |    218 ns/op |        32 |          2 |   1.18 |        17208 |  109.25
BenchmarkSerializers/unmarshal/mus-4    |    5190254 |    226 ns/op |        32 |          2 |   1.17 |        16608 |  113.10
BenchmarkSerializers/unmarshal/flatbuffers-4 |    4762442 |    244 ns/op |        32 |          2 |   1.17 |        15239 |  122.45
BenchmarkSerializers/marshal/calmh/xdr-4 |    4938007 |    245 ns/op |        60 |         64 |   1.21 |        29628 |    3.84
BenchmarkSerializers/unmarshal/msgp-4   |    4178436 |    255 ns/op |        32 |          2 |   1.07 |        13370 |  127.85
BenchmarkSerializers/unmarshal/gogo/protobuf-4 |    4428109 |    270 ns/op |        32 |          2 |   1.20 |        14169 |  135.20
BenchmarkSerializers/marshal/shamaton/msgpackgen/array-4 |    3426714 |    342 ns/op |        50 |        144 |   1.17 |        17133 |    2.38
BenchmarkSerializers/unmarshal/shamaton/msgpackgen/array-4 |    3412923 |    355 ns/op |       112 |          3 |   1.21 |        38224 |  118.40
BenchmarkSerializers/marshal/shamaton/msgpackgen/map-4 |    2844409 |    425 ns/op |        92 |        176 |   1.21 |        26168 |    2.42
BenchmarkSerializers/marshal/gotiny-4   |    2122971 |    552 ns/op |        47 |        168 |   1.17 |        10188 |    3.29
BenchmarkSerializers/unmarshal/shamaton/msgpackgen/map-4 |    2052736 |    570 ns/op |       112 |          3 |   1.17 |        22990 |  190.13
BenchmarkSerializers/unmarshal/capnproto-4 |    2070158 |    582 ns/op |       112 |          5 |   1.21 |        23185 |  116.58
BenchmarkSerializers/unmarshal/hprose2-4 |    1751545 |    649 ns/op |        56 |          3 |   1.14 |         9808 |  216.53
BenchmarkSerializers/unmarshal/shamaton/msgpack/array-4 |    1623724 |    697 ns/op |        88 |          4 |   1.13 |        14288 |  174.28
BenchmarkSerializers/marshal/hprose2-4  |    1630219 |    725 ns/op |        85 |          0 |   1.18 |        13895 | Zero-Allocs
BenchmarkSerializers/unmarshal/pulsar-4 |    1527993 |    758 ns/op |       160 |          6 |   1.16 |        24447 |  126.43
BenchmarkSerializers/marshal/ikea-4     |    1530195 |    784 ns/op |        55 |         72 |   1.20 |         8416 |   10.89
BenchmarkSerializers/unmarshal/shamaton/msgpack/map-4 |    1343259 |    892 ns/op |        88 |          4 |   1.20 |        11820 |  223.18
BenchmarkSerializers/marshal/shamaton/msgpack/array-4 |    1269360 |    903 ns/op |        50 |        160 |   1.15 |         6346 |    5.65
BenchmarkSerializers/marshal/pulsar-4   |    1343660 |    904 ns/op |        51 |        304 |   1.22 |         6923 |    2.98
BenchmarkSerializers/marshal/dedis/protobuf-4 |    1235511 |    933 ns/op |        52 |        144 |   1.15 |         6424 |    6.48
BenchmarkSerializers/unmarshal/ikea-4   |    1265690 |    960 ns/op |        96 |         10 |   1.22 |        12150 |   96.01
BenchmarkSerializers/marshal/msgpack-4  |     985917 |   1057 ns/op |        92 |        264 |   1.04 |         9070 |    4.00
BenchmarkSerializers/marshal/shamaton/msgpack/map-4 |    1081304 |   1130 ns/op |        92 |        192 |   1.22 |         9947 |    5.89
BenchmarkSerializers/marshal/flatbuffers-4 |    1009303 |   1192 ns/op |        95 |        376 |   1.20 |         9604 |    3.17
BenchmarkSerializers/unmarshal/dedis/protobuf-4 |     839454 |   1199 ns/op |       104 |          9 |   1.01 |         8730 |  133.22
BenchmarkSerializers/marshal/jsoniter-4 |    1002337 |   1206 ns/op |       141 |        200 |   1.21 |        14173 |    6.03
BenchmarkSerializers/marshal/avro2/binary-4 |     853518 |   1368 ns/op |        47 |        464 |   1.17 |         4011 |    2.95
BenchmarkSerializers/unmarshal/easyjson-4 |     835852 |   1374 ns/op |        32 |          2 |   1.15 |         2674 |  687.00
BenchmarkSerializers/marshal/hprose-4   |     738270 |   1383 ns/op |        85 |        325 |   1.02 |         6295 |    4.26
BenchmarkSerializers/marshal/bson-4     |     710941 |   1424 ns/op |       110 |        376 |   1.01 |         7820 |    3.79
BenchmarkSerializers/unmarshal/avro2/binary-4 |     698176 |   1442 ns/op |       464 |         10 |   1.01 |        32395 |  144.20
BenchmarkSerializers/unmarshal/msgpack-4 |     756810 |   1462 ns/op |        80 |          3 |   1.11 |         6054 |  487.33
BenchmarkSerializers/unmarshal/hprose-4 |     762679 |   1507 ns/op |       224 |          9 |   1.15 |        17084 |  167.44
BenchmarkSerializers/unmarshal/davecgh/xdr-4 |     722668 |   1607 ns/op |       152 |         10 |   1.16 |        10984 |  160.70
BenchmarkSerializers/unmarshal/ugorji/msgpack-4 |     698718 |   1687 ns/op |       608 |          3 |   1.18 |        42482 |  562.33
BenchmarkSerializers/marshal/ugorji/msgpack-4 |     713061 |   1734 ns/op |        91 |       1240 |   1.24 |         6488 |    1.40
BenchmarkSerializers/marshal/easyjson-4 |     692306 |   1744 ns/op |       151 |        976 |   1.21 |        10502 |    1.79
BenchmarkSerializers/unmarshal/jsoniter-4 |     709945 |   1762 ns/op |       136 |          5 |   1.25 |         9655 |  352.40
BenchmarkSerializers/marshal/ugorji/binc-4 |     648315 |   1817 ns/op |        95 |       1256 |   1.18 |         6158 |    1.45
BenchmarkSerializers/marshal/davecgh/xdr-4 |     548881 |   1918 ns/op |        92 |        392 |   1.05 |         5049 |    4.89
BenchmarkSerializers/unmarshal/ugorji/binc-4 |     594816 |   1930 ns/op |       832 |          6 |   1.15 |        49488 |  321.67
BenchmarkSerializers/marshal/mongobson-4 |     509228 |   2098 ns/op |       110 |        240 |   1.07 |         5601 |    8.74
BenchmarkSerializers/marshal/json-4     |     484378 |   2127 ns/op |       151 |        208 |   1.03 |         7352 |   10.23
BenchmarkSerializers/unmarshal/alecthomas/binary-4 |     562761 |   2148 ns/op |       240 |         19 |   1.21 |        13506 |  113.05
BenchmarkSerializers/unmarshal/bson-4   |     469126 |   2285 ns/op |       144 |         18 |   1.07 |         6755 |  126.94
BenchmarkSerializers/marshal/alecthomas/binary-4 |     493554 |   2336 ns/op |        61 |        360 |   1.15 |         3010 |    6.49
BenchmarkSerializers/unmarshal/mongobson-4 |     517167 |   2373 ns/op |       328 |         14 |   1.23 |        16963 |  169.50
BenchmarkSerializers/marshal/fastjson-4 |     506050 |   2413 ns/op |       133 |       1360 |   1.22 |         6770 |    1.77
BenchmarkSerializers/marshal/goavro-4   |     390975 |   2849 ns/op |        47 |        584 |   1.11 |         1837 |    4.88
BenchmarkSerializers/unmarshal/ssz-4    |     359684 |   3291 ns/op |       264 |         27 |   1.18 |         9495 |  121.89
BenchmarkSerializers/unmarshal/fastjson-4 |     354904 |   3452 ns/op |      1800 |         11 |   1.23 |        63882 |  313.82
BenchmarkSerializers/unmarshal/avro2/text-4 |     284338 |   3892 ns/op |       656 |         30 |   1.11 |        18652 |  129.73
BenchmarkSerializers/marshal/sereal-4   |     263485 |   4206 ns/op |       142 |       1104 |   1.11 |         3741 |    3.81
BenchmarkSerializers/marshal/capnproto-4 |     282118 |   4239 ns/op |        96 |       4392 |   1.20 |         2708 |    0.97
BenchmarkSerializers/unmarshal/sereal-4 |     239680 |   4376 ns/op |       896 |         33 |   1.05 |        21475 |  132.61
BenchmarkSerializers/marshal/avro2/text-4 |     234855 |   4686 ns/op |       133 |       1320 |   1.10 |         3140 |    3.55
BenchmarkSerializers/unmarshal/json-4   |     230168 |   4720 ns/op |       248 |          6 |   1.09 |         5708 |  786.67
BenchmarkSerializers/marshal/ssz-4      |     225304 |   4895 ns/op |        55 |        416 |   1.10 |         1239 |   11.77
BenchmarkSerializers/unmarshal/goavro-4 |     148034 |   7339 ns/op |      2232 |         52 |   1.09 |        33041 |  141.13
BenchmarkSerializers/marshal/gob-4      |     141312 |   7625 ns/op |       172 |       1744 |   1.08 |         2439 |    4.37
BenchmarkSerializers/marshal/gogo/jsonpb-4 |      67770 |  16636 ns/op |       125 |       2755 |   1.13 |          852 |    6.04
BenchmarkSerializers/unmarshal/gogo/jsonpb-4 |      55772 |  21796 ns/op |      3159 |         55 |   1.22 |        17618 |  396.29
BenchmarkSerializers/unmarshal/gob-4    |      31680 |  38385 ns/op |      7656 |        203 |   1.22 |        24254 |  189.09

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
