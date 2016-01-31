# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

## Tested serialization methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [gopkg.in/vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson) *(using the bsongen code generator)*
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp) *(code generator for msgpack)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://gogo.github.io/) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)
- [github.com/google/flatbuffers](https://github.com/google/flatbuffers)
- [github.com/hprose/hprose-go/io](https://github.com/hprose/hprose-go)
- [github.com/glycerine/go-capnproto](https://github.com/glycerine/go-capnproto)
- [zombiezen.com/go/capnproto2](https://godoc.org/zombiezen.com/go/capnproto2)
- [github.com/andyleap/gencode](https://github.com/andyleap/gencode)

## Running the benchmarks

```bash
go get -u -t
go test -bench='.*' ./
```

Shameless plug: I use [pawk](https://github.com/alecthomas/pawk) to format the table:

```bash
go test -bench='.*' ./ | pawk -F'\t' '"%-40s %10s %10s %s %s" % f'
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

Results with Go 1.5.1 on a 3.1 GHz Intel Core i7 (MacBook Pro 13"):

```
benchmark                                   iter           time/iter      bytes alloc               allocs
---------                                   ----           ---------      -----------               ------
BenchmarkMsgpMarshal-4                	 5000000	       299 ns/op	     128 B/op	       1 allocs/op
BenchmarkMsgpUnmarshal-4              	 3000000	       532 ns/op	     112 B/op	       3 allocs/op
BenchmarkVmihailencoMsgpackMarshal-4  	  500000	      2597 ns/op	     400 B/op	       6 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal-4	  500000	      2381 ns/op	     400 B/op	      13 allocs/op
BenchmarkJsonMarshal-4                	  500000	      3116 ns/op	     584 B/op	       7 allocs/op
BenchmarkJsonUnmarshal-4              	  300000	      5344 ns/op	     463 B/op	       8 allocs/op
BenchmarkBsonMarshal-4                	 1000000	      2078 ns/op	     416 B/op	      10 allocs/op
BenchmarkBsonUnmarshal-4              	  500000	      2890 ns/op	     416 B/op	      21 allocs/op
BenchmarkVitessBsonMarshal-4          	 1000000	      1485 ns/op	    1168 B/op	       4 allocs/op
BenchmarkVitessBsonUnmarshal-4        	 1000000	      1014 ns/op	     224 B/op	       4 allocs/op
BenchmarkGobMarshal-4                 	 1000000	      1292 ns/op	      48 B/op	       2 allocs/op
BenchmarkGobUnmarshal-4               	 1000000	      1589 ns/op	     160 B/op	       6 allocs/op
BenchmarkXdrMarshal-4                 	  500000	      2589 ns/op	     591 B/op	      20 allocs/op
BenchmarkXdrUnmarshal-4               	 1000000	      1954 ns/op	     287 B/op	      11 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal-4  	  500000	      4036 ns/op	    2720 B/op	       8 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal-4	  300000	      4749 ns/op	    3104 B/op	      12 allocs/op
BenchmarkUgorjiCodecBincMarshal-4     	  300000	      4732 ns/op	    2752 B/op	       8 allocs/op
BenchmarkUgorjiCodecBincUnmarshal-4   	  300000	      4918 ns/op	    3264 B/op	      15 allocs/op
BenchmarkSerealMarshal-4              	  300000	      5415 ns/op	     928 B/op	      21 allocs/op
BenchmarkSerealUnmarshal-4            	  300000	      4894 ns/op	    1152 B/op	      34 allocs/op
BenchmarkBinaryMarshal-4              	 1000000	      2203 ns/op	     352 B/op	      16 allocs/op
BenchmarkBinaryUnmarshal-4            	 1000000	      2347 ns/op	     448 B/op	      22 allocs/op
BenchmarkFlatbuffersMarshal-4         	 3000000	       475 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlatBuffersUnmarshal-4       	 3000000	       386 ns/op	     112 B/op	       3 allocs/op
BenchmarkCapNProtoMarshal-4           	 2000000	       627 ns/op	      64 B/op	       2 allocs/op
BenchmarkCapNProtoUnmarshal-4         	10000000	       229 ns/op	       0 B/op	       0 allocs/op
BenchmarkCapNProto2Marshal-4          	 1000000	      2334 ns/op	     608 B/op	      15 allocs/op
BenchmarkCapNProto2Unmarshal-4        	 1000000	      2191 ns/op	     640 B/op	      16 allocs/op
BenchmarkHproseMarshal-4              	 1000000	      1514 ns/op	     495 B/op	       8 allocs/op
BenchmarkHproseUnmarshal-4            	 1000000	      1625 ns/op	     319 B/op	      10 allocs/op
BenchmarkProtobufMarshal-4            	 1000000	      1506 ns/op	     224 B/op	       7 allocs/op
BenchmarkProtobufUnmarshal-4          	 1000000	      1286 ns/op	     240 B/op	      10 allocs/op
BenchmarkGoprotobufMarshal-4          	 2000000	       963 ns/op	     320 B/op	       4 allocs/op
BenchmarkGoprotobufUnmarshal-4        	 1000000	      1236 ns/op	     432 B/op	       9 allocs/op
BenchmarkGogoprotobufMarshal-4        	 5000000	       249 ns/op	      64 B/op	       1 allocs/op
BenchmarkGogoprotobufUnmarshal-4      	 5000000	       378 ns/op	     112 B/op	       3 allocs/op
BenchmarkGencodeMarshal-4             	 5000000	       257 ns/op	      80 B/op	       2 allocs/op
BenchmarkGencodeUnmarshal-4           	 5000000	       319 ns/op	     112 B/op	       3 allocs/op
```

## Issues


The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
2. **(minor)** Vitess BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
--- FAIL: BenchmarkBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20b999e3621bd773 2016-01-19 14:05:02.469416459 -0800 PST f017c8e9de 4 true 0.20887343719329818}
        &{20b999e3621bd773 2016-01-19 14:05:02.469 -0800 PST f017c8e9de 4 true 0.20887343719329818}
--- FAIL: BenchmarkVitessBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{64204cc8fe408156 2016-01-19 14:05:04.068418034 -0800 PST 95d04e2d49 3 false 0.011931519836294776}
        &{64204cc8fe408156 2016-01-19 22:05:04.068 +0000 UTC 95d04e2d49 3 false 0.011931519836294776}
--- FAIL: BenchmarkUgorjiCodecBincUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 PST 71f3bf4233 0 false 0.8712180830484527}
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 -0800 71f3bf4233 0 false 0.8712180830484527}
```

All other fields are correct however.

Additionally, while not a correctness issue, FlatBuffers, ProtoBuffers and Cap'N'Proto do not
support time types directly. In the benchmarks an int64 value is used to hold a UnixNano timestamp.
