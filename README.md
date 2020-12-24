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

2020-12-24 Results with Go 1.15.6 windows/amd64 on a 3.4 GHz Intel Core i7 16GB

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGotinyMarshal-8                 |    8333517 |    133 ns/op |    48 |   0 |   1.11 |   40000 |    0.00
BenchmarkGotinyUnmarshal-8               |    4545450 |    269 ns/op |    48 | 112 |   1.22 |   21818 |    2.40
BenchmarkGotinyNoTimeMarshal-8           |    9302505 |    133 ns/op |    48 |   0 |   1.24 |   44652 |    0.00
BenchmarkGotinyNoTimeUnmarshal-8         |    4852605 |    254 ns/op |    48 |  96 |   1.23 |   23292 |    2.65
BenchmarkMsgpMarshal-8                   |    6872631 |    178 ns/op |    97 | 128 |   1.22 |   66664 |    1.39
BenchmarkMsgpUnmarshal-8                 |    3732506 |    311 ns/op |    97 | 112 |   1.16 |   36205 |    2.78
BenchmarkVmihailencoMsgpackMarshal-8     |    1000000 |   1075 ns/op |   100 | 288 |   1.07 |   10000 |    3.73
BenchmarkVmihailencoMsgpackUnmarshal-8   |     727321 |   1477 ns/op |   100 | 160 |   1.07 |    7273 |    9.23
BenchmarkJsonMarshal-8                   |     280615 |   4168 ns/op |   150 | 944 |   1.17 |    4209 |    4.42
BenchmarkJsonUnmarshal-8                 |     292680 |   4165 ns/op |   150 | 344 |   1.22 |    4390 |   12.11
BenchmarkJsonIterMarshal-8               |     252634 |   4597 ns/op |   149 | 1113 |   1.16 |    3764 |    4.13
BenchmarkJsonIterUnmarshal-8             |     337996 |   3669 ns/op |   150 | 447 |   1.24 |    5069 |    8.21
BenchmarkEasyJsonMarshal-8               |     800031 |   1420 ns/op |   150 | 784 |   1.14 |   12000 |    1.81
BenchmarkEasyJsonUnmarshal-8             |     750205 |   1580 ns/op |   146 | 144 |   1.19 |   10952 |   10.97
BenchmarkBsonMarshal-8                   |    1000000 |   1298 ns/op |   110 | 392 |   1.30 |   11000 |    3.31
BenchmarkBsonUnmarshal-8                 |     749938 |   1753 ns/op |   110 | 232 |   1.31 |    8249 |    7.56
BenchmarkGobMarshal-8                    |    1512285 |    816 ns/op |    63 |  48 |   1.23 |    9618 |   17.00
BenchmarkGobUnmarshal-8                  |    1378516 |    876 ns/op |    63 | 112 |   1.21 |    8767 |    7.82
BenchmarkXDRMarshal-8                    |     685682 |   1793 ns/op |    92 | 456 |   1.23 |    6308 |    3.93
BenchmarkXDRUnmarshal-8                  |     749892 |   1594 ns/op |    92 | 240 |   1.20 |    6899 |    6.64
BenchmarkUgorjiCodecMsgpackMarshal-8     |    1000000 |   1299 ns/op |    91 | 1280 |   1.30 |    9100 |    1.01
BenchmarkUgorjiCodecMsgpackUnmarshal-8   |     960452 |   1212 ns/op |    91 | 480 |   1.16 |    8740 |    2.52
BenchmarkUgorjiCodecBincMarshal-8        |     888895 |   1400 ns/op |    95 | 1312 |   1.24 |    8444 |    1.07
BenchmarkUgorjiCodecBincUnmarshal-8      |     850628 |   1521 ns/op |    95 | 680 |   1.29 |    8080 |    2.24
BenchmarkSerealMarshal-8                 |     406792 |   3053 ns/op |   132 | 904 |   1.24 |    5369 |    3.38
BenchmarkSerealUnmarshal-8               |     347787 |   3473 ns/op |   132 | 1008 |   1.21 |    4590 |    3.45
BenchmarkBinaryMarshal-8                 |     857222 |   1428 ns/op |    61 | 320 |   1.22 |    5229 |    4.46
BenchmarkBinaryUnmarshal-8               |     799999 |   1511 ns/op |    61 | 320 |   1.21 |    4879 |    4.72
BenchmarkFlatBuffersMarshal-8            |    3947342 |    309 ns/op |    95 |   0 |   1.22 |   37578 |    0.00
BenchmarkFlatBuffersUnmarshal-8          |    4461014 |    271 ns/op |    95 | 112 |   1.21 |   42424 |    2.42
BenchmarkCapNProtoMarshal-8              |    3104388 |    392 ns/op |    96 |  56 |   1.22 |   29802 |    7.00
BenchmarkCapNProtoUnmarshal-8            |    2733318 |    448 ns/op |    96 | 200 |   1.22 |   26239 |    2.24
BenchmarkCapNProto2Marshal-8             |    1983369 |    613 ns/op |    96 | 244 |   1.22 |   19040 |    2.51
BenchmarkCapNProto2Unmarshal-8           |    1478056 |    843 ns/op |    96 | 320 |   1.25 |   14189 |    2.63
BenchmarkHproseMarshal-8                 |    1000000 |   1061 ns/op |    85 | 478 |   1.06 |    8530 |    2.22
BenchmarkHproseUnmarshal-8               |     923190 |   1317 ns/op |    85 | 320 |   1.22 |    7874 |    4.12
BenchmarkHprose2Marshal-8                |    2081517 |    613 ns/op |    85 |   0 |   1.28 |   17755 |    0.00
BenchmarkHprose2Unmarshal-8              |    1892756 |    619 ns/op |    85 | 144 |   1.17 |   16145 |    4.30
BenchmarkProtobufMarshal-8               |    1473292 |    769 ns/op |    52 | 152 |   1.13 |    7661 |    5.06
BenchmarkProtobufUnmarshal-8             |     800031 |   1262 ns/op |    52 | 280 |   1.01 |    4160 |    4.51
BenchmarkGoprotobufMarshal-8             |    1368230 |    881 ns/op |    53 |  96 |   1.21 |    7251 |    9.18
BenchmarkGoprotobufUnmarshal-8           |    1544403 |    783 ns/op |    53 | 184 |   1.21 |    8185 |    4.26
BenchmarkGogoprotobufMarshal-8           |    8391619 |    149 ns/op |    53 |  64 |   1.25 |   44475 |    2.33
BenchmarkGogoprotobufUnmarshal-8         |    5121476 |    241 ns/op |    53 |  96 |   1.23 |   27143 |    2.51
BenchmarkColferMarshal-8                 |    9486060 |    125 ns/op |    51 |  64 |   1.19 |   48473 |    1.95
BenchmarkColferUnmarshal-8               |    6137526 |    201 ns/op |    52 | 112 |   1.23 |   31915 |    1.79
BenchmarkGencodeMarshal-8                |    6722925 |    190 ns/op |    53 |  80 |   1.28 |   35631 |    2.38
BenchmarkGencodeUnmarshal-8              |    4771398 |    221 ns/op |    53 | 112 |   1.05 |   25288 |    1.97
BenchmarkGencodeUnsafeMarshal-8          |   12182888 |     99 ns/op |    46 |  48 |   1.22 |   56041 |    2.08
BenchmarkGencodeUnsafeUnmarshal-8        |    7145736 |    160 ns/op |    46 |  96 |   1.14 |   32870 |    1.67
BenchmarkXDR2Marshal-8                   |    7568751 |    162 ns/op |    60 |  64 |   1.23 |   45412 |    2.53
BenchmarkXDR2Unmarshal-8                 |    8695588 |    133 ns/op |    60 |  32 |   1.16 |   52173 |    4.16
BenchmarkGoAvroMarshal-8                 |     421068 |   2859 ns/op |    47 | 1008 |   1.20 |    1979 |    2.84
BenchmarkGoAvroUnmarshal-8               |     171081 |   7157 ns/op |    47 | 3328 |   1.22 |     804 |    2.15
BenchmarkGoAvro2TextMarshal-8            |     406749 |   2991 ns/op |   134 | 1320 |   1.22 |    5450 |    2.27
BenchmarkGoAvro2TextUnmarshal-8          |     406765 |   3014 ns/op |   134 | 799 |   1.23 |    5450 |    3.77
BenchmarkGoAvro2BinaryMarshal-8          |    1000000 |   1114 ns/op |    47 | 488 |   1.11 |    4700 |    2.28
BenchmarkGoAvro2BinaryUnmarshal-8        |    1000000 |   1128 ns/op |    47 | 560 |   1.13 |    4700 |    2.01
BenchmarkIkeaMarshal-8                   |    1637104 |    705 ns/op |    55 |  72 |   1.15 |    9004 |    9.79
BenchmarkIkeaUnmarshal-8                 |    1360539 |    866 ns/op |    55 | 160 |   1.18 |    7482 |    5.41
BenchmarkShamatonMapMsgpackMarshal-8     |    1438813 |    825 ns/op |    92 | 208 |   1.19 |   13237 |    3.97
BenchmarkShamatonMapMsgpackUnmarshal-8   |    1593631 |    777 ns/op |    92 | 144 |   1.24 |   14661 |    5.40
BenchmarkShamatonArrayMsgpackMarshal-8   |    1673511 |    714 ns/op |    50 | 176 |   1.19 |    8367 |    4.06
BenchmarkShamatonArrayMsgpackUnmarshal-8 |    2380966 |    519 ns/op |    50 | 144 |   1.24 |   11904 |    3.60
BenchmarkSSZNoTimeNoStringNoFloatAMarshal-8 |     266820 |   4700 ns/op |    55 | 440 |   1.25 |    1467 |   10.68
BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal-8 |     147238 |   7753 ns/op |    55 | 1392 |   1.14 |     809 |    5.57
BenchmarkBebopMarshal-8                  |    9575738 |    130 ns/op |    55 |  64 |   1.24 |   52666 |    2.03
BenchmarkBebopUnmarshal-8                |   11111234 |    107 ns/op |    55 |  32 |   1.19 |   61111 |    3.34


Totals:


benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkBebop-8                         |   20686972 |    237 ns/op |   110 |  96 |   4.90 |  227556 |    2.47
BenchmarkGencodeUnsafe-8                 |   19328624 |    259 ns/op |    92 | 144 |   5.02 |  177823 |    1.80
BenchmarkXDR2-8                          |   16264339 |    295 ns/op |   120 |  96 |   4.80 |  195172 |    3.07
BenchmarkColfer-8                        |   15623586 |    326 ns/op |   103 | 176 |   5.09 |  161079 |    1.85
BenchmarkGotinyNoTime-8                  |   14155110 |    387 ns/op |    96 |  96 |   5.48 |  135889 |    4.03
BenchmarkGogoprotobuf-8                  |   13513095 |    390 ns/op |   106 | 160 |   5.27 |  143238 |    2.44
BenchmarkGotiny-8                        |   12878967 |    402 ns/op |    96 | 112 |   5.18 |  123638 |    3.59
BenchmarkGencode-8                       |   11494323 |    411 ns/op |   106 | 192 |   4.72 |  121839 |    2.14
BenchmarkMsgp-8                          |   10605137 |    489 ns/op |   194 | 240 |   5.19 |  205739 |    2.04
BenchmarkFlatBuffers-8                   |    8408356 |    580 ns/op |   190 | 112 |   4.88 |  160011 |    5.18
BenchmarkCapNProto-8                     |    5837706 |    840 ns/op |   192 | 256 |   4.90 |  112083 |    3.28
BenchmarkHprose2-8                       |    3974273 |   1232 ns/op |   170 | 144 |   4.90 |   67801 |    8.56
BenchmarkShamatonArrayMsgpack-8          |    4054477 |   1233 ns/op |   100 | 320 |   5.00 |   40544 |    3.85
BenchmarkCapNProto2-8                    |    3461425 |   1456 ns/op |   192 | 564 |   5.04 |   66459 |    2.58
BenchmarkIkea-8                          |    2997643 |   1571 ns/op |   110 | 232 |   4.71 |   32974 |    6.77
BenchmarkShamatonMapMsgpack-8            |    3032444 |   1602 ns/op |   184 | 352 |   4.86 |   55796 |    4.55
BenchmarkGoprotobuf-8                    |    2912633 |   1664 ns/op |   106 | 280 |   4.85 |   30873 |    5.94
BenchmarkGob-8                           |    2890801 |   1692 ns/op |   127 | 160 |   4.89 |   36770 |   10.57
BenchmarkProtobuf-8                      |    2273323 |   2031 ns/op |   104 | 432 |   4.62 |   23642 |    4.70
BenchmarkGoAvro2Binary-8                 |    2000000 |   2242 ns/op |    94 | 1048 |   4.48 |   18800 |    2.14
BenchmarkHprose-8                        |    1923190 |   2378 ns/op |   170 | 798 |   4.57 |   32809 |    2.98
BenchmarkUgorjiCodecMsgpack-8            |    1960452 |   2511 ns/op |   182 | 1760 |   4.92 |   35680 |    1.43
BenchmarkVmihailencoMsgpack-8            |    1727321 |   2552 ns/op |   200 | 448 |   4.41 |   34546 |    5.70
BenchmarkUgorjiCodecBinc-8               |    1739523 |   2921 ns/op |   190 | 1992 |   5.08 |   33050 |    1.47
BenchmarkBinary-8                        |    1657221 |   2939 ns/op |   122 | 640 |   4.87 |   20218 |    4.59
BenchmarkEasyJson-8                      |    1550236 |   3000 ns/op |   296 | 928 |   4.65 |   45886 |    3.23
BenchmarkBson-8                          |    1749938 |   3051 ns/op |   220 | 624 |   5.34 |   38498 |    4.89
BenchmarkXDR-8                           |    1435574 |   3387 ns/op |   184 | 696 |   4.86 |   26414 |    4.87
BenchmarkGoAvro2Text-8                   |     813514 |   6005 ns/op |   268 | 2119 |   4.89 |   21802 |    2.83
BenchmarkSereal-8                        |     754579 |   6526 ns/op |   264 | 1912 |   4.92 |   19920 |    3.41
BenchmarkJsonIter-8                      |     590630 |   8266 ns/op |   299 | 1560 |   4.88 |   17659 |    5.30
BenchmarkJson-8                          |     573295 |   8333 ns/op |   300 | 1288 |   4.78 |   17198 |    6.47
BenchmarkGoAvro-8                        |     592149 |  10016 ns/op |    94 | 4336 |   5.93 |    5566 |    2.31
BenchmarkSSZNoTimeNoStringNoFloatA-8     |     414058 |  12453 ns/op |   110 | 1832 |   5.16 |    4554 |    6.80



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
