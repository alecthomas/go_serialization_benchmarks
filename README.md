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
- [gopkg.in/vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp) *(code generator for msgpack)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://github.com/gogo/protobuf) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)
- [github.com/google/flatbuffers](https://github.com/google/flatbuffers)
- [github.com/hprose/hprose-go/io](https://github.com/hprose/hprose-go)
- [github.com/glycerine/go-capnproto](https://github.com/glycerine/go-capnproto)
- [zombiezen.com/go/capnproto2](https://godoc.org/zombiezen.com/go/capnproto2)
- [github.com/andyleap/gencode](https://github.com/andyleap/gencode)
- [github.com/pascaldekloe/colfer](https://github.com/pascaldekloe/colfer)
- [github.com/linkedin/goavro](https://github.com/linkedin/goavro)
- [github.com/ikkerens/ikeapack](https://github.com/ikkerens/ikeapack)
- [github.com/niubaoshu/gotiny](https://github.com/niubaoshu/gotiny)

## Running the benchmarks

```bash
go get -u -t
make install
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

2018-08-19 Results with Go 1.10 on a 2.5 GHz Intel Core i7 MacBook Pro (Retina, 15-inch, Mid 2015):

```
goos: darwin
goarch: amd64
name                       	               time/iter 	  lenth/iter 	     bytes/iter       allocs/iter 	 generate 	 iter
----                       	               --------- 	  ---------- 	     ----------       ----------- 	 -------- 	 ----
gotiny-Marshal-8           		       142 ns/op	   52 byte	       0 B/op	      0 allocs/op	    NO      	10000000
gotiny-Unmarshal-8         		       194 ns/op	   52 byte	      32 B/op	      2 allocs/op	    NO      	10000000
gotiny_notime-Marshal-8    		       142 ns/op	   48 byte	       0 B/op	      0 allocs/op	    NO      	10000000
gotiny_notime-Unmarshal-8  		       185 ns/op	   48 byte	      32 B/op	      2 allocs/op	    NO      	10000000
msgp-Marshal-8             		       183 ns/op	  106 byte	     128 B/op	      1 allocs/op	    YES     	10000000
msgp-Unmarshal-8           		       297 ns/op	  106 byte	      32 B/op	      2 allocs/op	    YES     	 5000000
msgpack-Marshal-8          		      1980 ns/op	   93 byte	     368 B/op	      6 allocs/op	    NO      	 1000000
msgpack-Unmarshal-8        		      2036 ns/op	   93 byte	     304 B/op	     12 allocs/op	    NO      	 1000000
json-Marshal-8             		      2006 ns/op	  150 byte	     304 B/op	      4 allocs/op	    NO      	 1000000
json-Unmarshal-8           		      4115 ns/op	  150 byte	     279 B/op	      6 allocs/op	    NO      	  300000
jsoniter-Marshal-8         		      1300 ns/op	  139 byte	     200 B/op	      3 allocs/op	    NO      	 1000000
jsoniter-Unmarshal-8       		      1479 ns/op	  139 byte	     167 B/op	      6 allocs/op	    NO      	 1000000
easyJson-Marshal-8         		      1433 ns/op	  150 byte	     896 B/op	      6 allocs/op	    YES     	 1000000
easyJson-Unmarshal-8       		      1339 ns/op	  150 byte	     207 B/op	      4 allocs/op	    YES     	 1000000
bson-Marshal-8             		      1233 ns/op	  111 byte	     392 B/op	     10 allocs/op	    NO      	 1000000
bson-Unmarshal-8           		      1778 ns/op	  111 byte	     164 B/op	     18 allocs/op	    NO      	 1000000
gob-Marshal-8              		       937 ns/op	   95 byte	      48 B/op	      2 allocs/op	    NO      	 2000000
gob-Unmarshal-8            		       907 ns/op	   95 byte	      32 B/op	      2 allocs/op	    NO      	 2000000
xdr-Marshal-8              		      1713 ns/op	   88 byte	     392 B/op	     19 allocs/op	    NO      	 1000000
xdr-Unmarshal-8            		      1481 ns/op	   88 byte	     144 B/op	     10 allocs/op	    NO      	 1000000
ugorjicodec-Marshal-8      		      1230 ns/op	   95 byte	     577 B/op	      7 allocs/op	    NO      	 1000000
ugorjicodec-Unmarshal-8    		      1509 ns/op	   95 byte	     609 B/op	      8 allocs/op	    NO      	 1000000
sereal-Marshal-8           		      2748 ns/op	  134 byte	     904 B/op	     20 allocs/op	    NO      	  500000
sereal-Unmarshal-8         		      2850 ns/op	  134 byte	     928 B/op	     33 allocs/op	    NO      	  500000
alecthomas-Marshal-8       		      1463 ns/op	   61 byte	     334 B/op	     20 allocs/op	    NO      	 1000000
alecthomas-Unmarshal-8     		      1591 ns/op	   61 byte	     256 B/op	     21 allocs/op	    NO      	 1000000
FlatBuffer-Marshal-8       		       376 ns/op	  127 byte	       0 B/op	      0 allocs/op	    YES     	 3000000
FlatBuffer-Unmarshal-8     		       237 ns/op	  127 byte	      32 B/op	      2 allocs/op	    YES     	10000000
CapNProto-Marshal-8        		       557 ns/op	  128 byte	      56 B/op	      2 allocs/op	    YES     	 3000000
CapNProto-Unmarshal-8      		       517 ns/op	  128 byte	     120 B/op	      5 allocs/op	    YES     	 3000000
CapNProto2-Marshal-8       		       746 ns/op	  144 byte	     244 B/op	      3 allocs/op	    YES     	 2000000
CapNProto2-Unmarshal-8     		       891 ns/op	  144 byte	     240 B/op	      5 allocs/op	    YES     	 2000000
Hprose-Marshal-8           		      1011 ns/op	   83 byte	     329 B/op	      8 allocs/op	    NO      	 1000000
Hprose-Unmarshal-8         		      1033 ns/op	   83 byte	     239 B/op	      9 allocs/op	    NO      	 1000000
Hprose2-Marshal-8          		       653 ns/op	  123 byte	       0 B/op	      0 allocs/op	    NO      	 2000000
Hprose2-Unmarshal-8        		       535 ns/op	  123 byte	      64 B/op	      3 allocs/op	    NO      	 3000000
DeDiS-protobuf-Marshal-8   		      1018 ns/op	   52 byte	     200 B/op	      7 allocs/op	    NO      	 1000000
DeDiS-protobuf-Unmarshal-8 		       791 ns/op	   52 byte	     112 B/op	      9 allocs/op	    NO      	 2000000
golang-protobuf-Marshal-8  		       339 ns/op	   63 byte	      64 B/op	      1 allocs/op	    YES     	 5000000
golang-protobuf-Unmarshal-8		       424 ns/op	   63 byte	      88 B/op	      8 allocs/op	    YES     	 3000000
gogo-protobuf-Marshal-8    		       216 ns/op	   58 byte	      64 B/op	      1 allocs/op	    YES     	10000000
gogo-protobuf-Unmarshal-8  		       210 ns/op	   58 byte	      32 B/op	      2 allocs/op	    YES     	10000000
colfer-Marshal-8           		       151 ns/op	   56 byte	      64 B/op	      1 allocs/op	    YES     	10000000
colfer-Unmarshal-8         		       149 ns/op	   56 byte	      32 B/op	      2 allocs/op	    YES     	10000000
gencode-Marshal-8          		       202 ns/op	   58 byte	      80 B/op	      2 allocs/op	    YES     	10000000
gencode-Unmarshal-8        		       157 ns/op	   58 byte	      32 B/op	      2 allocs/op	    YES     	10000000
unsafe-gencode-Marshal-8   		       129 ns/op	   50 byte	      48 B/op	      1 allocs/op	    YES     	10000000
unsafe-gencode-Unmarshal-8 		       129 ns/op	   50 byte	      32 B/op	      2 allocs/op	    YES     	10000000
XDR-Marshal-8              		       183 ns/op	   70 byte	      64 B/op	      1 allocs/op	    YES     	10000000
XDR-Unmarshal-8            		       165 ns/op	   70 byte	      32 B/op	      2 allocs/op	    YES     	10000000
GoAvro-Marshal-8           		      2595 ns/op	   47 byte	    1030 B/op	     31 allocs/op	    NO      	  500000
GoAvro-Unmarshal-8         		      6166 ns/op	   47 byte	    3357 B/op	     86 allocs/op	    NO      	  200000
GoAvro2Text-Marshal-8      		      2805 ns/op	  136 byte	    1326 B/op	     20 allocs/op	    NO      	  500000
GoAvro2Text-Unmarshal-8    		      2537 ns/op	  136 byte	     726 B/op	     32 allocs/op	    NO      	  500000
GoAvro2Binary-Marshal-8    		       915 ns/op	   70 byte	     510 B/op	     11 allocs/op	    NO      	 2000000
GoAvro2Binary-Unmarshal-8  		       921 ns/op	   70 byte	     496 B/op	     12 allocs/op	    NO      	 2000000
Ikeapack-Marshal-8         		       735 ns/op	   88 byte	     192 B/op	      9 allocs/op	    NO      	 2000000
Ikeapack-Unmarshal-8       		       951 ns/op	   88 byte	     224 B/op	     11 allocs/op	    NO      	 2000000

sort:
unsafe-gencode-Marshal-8   		       129 ns/op	   50 byte	      48 B/op	      1 allocs/op	    YES     	10000000
unsafe-gencode-Unmarshal-8 		       129 ns/op	   50 byte	      32 B/op	      2 allocs/op	    YES     	10000000
colfer-Marshal-8           		       151 ns/op	   56 byte	      64 B/op	      1 allocs/op	    YES     	10000000
colfer-Unmarshal-8         		       149 ns/op	   56 byte	      32 B/op	      2 allocs/op	    YES     	10000000
gotiny_notime-Marshal-8    		       142 ns/op	   48 byte	       0 B/op	      0 allocs/op	    NO      	10000000
gotiny_notime-Unmarshal-8  		       185 ns/op	   48 byte	      32 B/op	      2 allocs/op	    NO      	10000000
gotiny-Marshal-8           		       142 ns/op	   52 byte	       0 B/op	      0 allocs/op	    NO      	10000000
gotiny-Unmarshal-8         		       194 ns/op	   52 byte	      32 B/op	      2 allocs/op	    NO      	10000000
XDR-Marshal-8              		       183 ns/op	   70 byte	      64 B/op	      1 allocs/op	    YES     	10000000
XDR-Unmarshal-8            		       165 ns/op	   70 byte	      32 B/op	      2 allocs/op	    YES     	10000000
gencode-Marshal-8          		       202 ns/op	   58 byte	      80 B/op	      2 allocs/op	    YES     	10000000
gencode-Unmarshal-8        		       157 ns/op	   58 byte	      32 B/op	      2 allocs/op	    YES     	10000000
gogo-protobuf-Marshal-8    		       216 ns/op	   58 byte	      64 B/op	      1 allocs/op	    YES     	10000000
gogo-protobuf-Unmarshal-8  		       210 ns/op	   58 byte	      32 B/op	      2 allocs/op	    YES     	10000000
msgp-Marshal-8             		       183 ns/op	  106 byte	     128 B/op	      1 allocs/op	    YES     	10000000
msgp-Unmarshal-8           		       297 ns/op	  106 byte	      32 B/op	      2 allocs/op	    YES     	 5000000
FlatBuffer-Marshal-8       		       376 ns/op	  127 byte	       0 B/op	      0 allocs/op	    YES     	 3000000
FlatBuffer-Unmarshal-8     		       237 ns/op	  127 byte	      32 B/op	      2 allocs/op	    YES     	10000000
golang-protobuf-Marshal-8  		       339 ns/op	   63 byte	      64 B/op	      1 allocs/op	    YES     	 5000000
golang-protobuf-Unmarshal-8		       424 ns/op	   63 byte	      88 B/op	      8 allocs/op	    YES     	 3000000
CapNProto-Marshal-8        		       557 ns/op	  128 byte	      56 B/op	      2 allocs/op	    YES     	 3000000
CapNProto-Unmarshal-8      		       517 ns/op	  128 byte	     120 B/op	      5 allocs/op	    YES     	 3000000
Hprose2-Marshal-8          		       653 ns/op	  123 byte	       0 B/op	      0 allocs/op	    NO      	 2000000
Hprose2-Unmarshal-8        		       535 ns/op	  123 byte	      64 B/op	      3 allocs/op	    NO      	 3000000
CapNProto2-Marshal-8       		       746 ns/op	  144 byte	     244 B/op	      3 allocs/op	    YES     	 2000000
CapNProto2-Unmarshal-8     		       891 ns/op	  144 byte	     240 B/op	      5 allocs/op	    YES     	 2000000
Ikeapack-Marshal-8         		       735 ns/op	   88 byte	     192 B/op	      9 allocs/op	    NO      	 2000000
Ikeapack-Unmarshal-8       		       951 ns/op	   88 byte	     224 B/op	     11 allocs/op	    NO      	 2000000
DeDiS-protobuf-Marshal-8   		      1018 ns/op	   52 byte	     200 B/op	      7 allocs/op	    NO      	 1000000
DeDiS-protobuf-Unmarshal-8 		       791 ns/op	   52 byte	     112 B/op	      9 allocs/op	    NO      	 2000000
GoAvro2Binary-Marshal-8    		       915 ns/op	   70 byte	     510 B/op	     11 allocs/op	    NO      	 2000000
GoAvro2Binary-Unmarshal-8  		       921 ns/op	   70 byte	     496 B/op	     12 allocs/op	    NO      	 2000000
gob-Marshal-8              		       937 ns/op	   95 byte	      48 B/op	      2 allocs/op	    NO      	 2000000
gob-Unmarshal-8            		       907 ns/op	   95 byte	      32 B/op	      2 allocs/op	    NO      	 2000000
Hprose-Marshal-8           		      1011 ns/op	   83 byte	     329 B/op	      8 allocs/op	    NO      	 1000000
Hprose-Unmarshal-8         		      1033 ns/op	   83 byte	     239 B/op	      9 allocs/op	    NO      	 1000000
ugorjicodec-Marshal-8      		      1230 ns/op	   95 byte	     577 B/op	      7 allocs/op	    NO      	 1000000
ugorjicodec-Unmarshal-8    		      1509 ns/op	   95 byte	     609 B/op	      8 allocs/op	    NO      	 1000000
easyJson-Marshal-8         		      1433 ns/op	  150 byte	     896 B/op	      6 allocs/op	    YES     	 1000000
easyJson-Unmarshal-8       		      1339 ns/op	  150 byte	     207 B/op	      4 allocs/op	    YES     	 1000000
jsoniter-Marshal-8         		      1300 ns/op	  139 byte	     200 B/op	      3 allocs/op	    NO      	 1000000
jsoniter-Unmarshal-8       		      1479 ns/op	  139 byte	     167 B/op	      6 allocs/op	    NO      	 1000000
bson-Marshal-8             		      1233 ns/op	  111 byte	     392 B/op	     10 allocs/op	    NO      	 1000000
bson-Unmarshal-8           		      1778 ns/op	  111 byte	     164 B/op	     18 allocs/op	    NO      	 1000000
alecthomas-Marshal-8       		      1463 ns/op	   61 byte	     334 B/op	     20 allocs/op	    NO      	 1000000
alecthomas-Unmarshal-8     		      1591 ns/op	   61 byte	     256 B/op	     21 allocs/op	    NO      	 1000000
xdr-Marshal-8              		      1713 ns/op	   88 byte	     392 B/op	     19 allocs/op	    NO      	 1000000
xdr-Unmarshal-8            		      1481 ns/op	   88 byte	     144 B/op	     10 allocs/op	    NO      	 1000000
msgpack-Marshal-8          		      1980 ns/op	   93 byte	     368 B/op	      6 allocs/op	    NO      	 1000000
msgpack-Unmarshal-8        		      2036 ns/op	   93 byte	     304 B/op	     12 allocs/op	    NO      	 1000000
GoAvro2Text-Marshal-8      		      2805 ns/op	  136 byte	    1326 B/op	     20 allocs/op	    NO      	  500000
GoAvro2Text-Unmarshal-8    		      2537 ns/op	  136 byte	     726 B/op	     32 allocs/op	    NO      	  500000
sereal-Marshal-8           		      2748 ns/op	  134 byte	     904 B/op	     20 allocs/op	    NO      	  500000
sereal-Unmarshal-8         		      2850 ns/op	  134 byte	     928 B/op	     33 allocs/op	    NO      	  500000
json-Marshal-8             		      2006 ns/op	  150 byte	     304 B/op	      4 allocs/op	    NO      	 1000000
json-Unmarshal-8           		      4115 ns/op	  150 byte	     279 B/op	      6 allocs/op	    NO      	  300000
GoAvro-Marshal-8           		      2595 ns/op	   47 byte	    1030 B/op	     31 allocs/op	    NO      	  500000
GoAvro-Unmarshal-8         		      6166 ns/op	   47 byte	    3357 B/op	     86 allocs/op	    NO      	  200000
```

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
