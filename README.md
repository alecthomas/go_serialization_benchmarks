# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

## Tested serialization methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go-msgpack](https://github.com/ugorji/go-msgpack)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)
- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson) *(using the bsongen code generator)*
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/philhofer/msgp](https://github.com/philhofer/msgp) *(code generator for msgpack)*


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

If performance, correctness and interoperability are the most important
factors, it seems that [github.com/philhofer/msgp](https://github.com/philhofer/msgp) or
[github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack) are
currently the best choice.

If performance is the biggest factor,
[github.com/youtube/vitess/tree/master/go/bson](//github.com/youtube/vitess/tree/master/go/bson)
is the best choice.

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
	Tags     map[string]string
	Aliases  []string
}
```


## Results

Results on my late 2013 MacBook Pro 15" are:

```
BenchmarkUgorjiMsgpackMarshal                500000       3716 ns/op     1328 B/op       21 allocs/op
BenchmarkUgorjiMsgpackUnmarshal              500000       3518 ns/op      644 B/op       20 allocs/op

BenchmarkVmihailencoMsgpackMarshal          1000000       1692 ns/op      412 B/op        6 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal        1000000       1999 ns/op      421 B/op       10 allocs/op

BenchmarkJsonMarshal                        1000000       2918 ns/op      590 B/op        7 allocs/op
BenchmarkJsonUnmarshal                       500000       4968 ns/op      468 B/op        7 allocs/op

BenchmarkBsonMarshal                        1000000       1981 ns/op      488 B/op       13 allocs/op
BenchmarkBsonUnmarshal                      1000000       2199 ns/op      281 B/op       10 allocs/op

BenchmarkVitessBsonMarshal                  1000000       1412 ns/op     1169 B/op        4 allocs/op
BenchmarkVitessBsonUnmarshal                2000000        803 ns/op      227 B/op        4 allocs/op

BenchmarkGobMarshal                          500000       6398 ns/op     1661 B/op       25 allocs/op
BenchmarkGobUnmarshal                         50000      46302 ns/op    19024 B/op      365 allocs/op

BenchmarkXdrMarshal                         1000000       2518 ns/op      520 B/op       15 allocs/op
BenchmarkXdrUnmarshal                       1000000       1857 ns/op      274 B/op        9 allocs/op

BenchmarkUgorjiCodecMsgpackMarshal           500000       3597 ns/op     1427 B/op       22 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal         500000       3613 ns/op     1145 B/op       29 allocs/op

BenchmarkUgorjiCodecBincMarshal              500000       4872 ns/op     2142 B/op       24 allocs/op
BenchmarkUgorjiCodecBincUnmarshal            500000       4795 ns/op     2064 B/op       34 allocs/op

BenchmarkSerealMarshal                       500000       3928 ns/op     1277 B/op       21 allocs/op
BenchmarkSerealUnmarshal                     500000       4319 ns/op      696 B/op       30 allocs/op

BenchmarkBinaryMarshal                      1000000       2102 ns/op      479 B/op       15 allocs/op
BenchmarkBinaryUnmarshal                    1000000       2232 ns/op      432 B/op       17 allocs/op

BenchmarkMsgpMarshal                        2000000        950 ns/op      322 B/op        4 allocs/op
BenchmarkMsgpUnmarshal                      1000000       1136 ns/op      178 B/op        5 allocs/op
```

**Note:** the gob results are not really representative of normal performance, as gob is designed for serializing streams or vectors of a single type, not individual values.


## Issues

The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(MAJOR)** Ugorji msgpack implementation drops the timezone frome `time.Time`.
2. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Vitess BSON drops sub-microsecond precision from `time.Time`.
4. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
BenchmarkUgorjiMsgpackMarshal     500000          3678 ns/op        1327 B/op         21 allocs/op
BenchmarkUgorjiMsgpackUnmarshal --- FAIL: BenchmarkUgorjiMsgpackUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{d8c7b339f1dd290e 2014-09-26 14:46:03.137970198 +1000 AEST 9894e18711 2 true 0.9013406798206226}
        &{d8c7b339f1dd290e 2014-09-26 04:46:03.137970198 +0000 UTC 9894e18711 2 true 0.9013406798206226}
BenchmarkVmihailencoMsgpackMarshal   1000000          1665 ns/op         412 B/op          6 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal     1000000          3001 ns/op         517 B/op         12 allocs/op
BenchmarkJsonMarshal     1000000          2843 ns/op         590 B/op          7 allocs/op
BenchmarkJsonUnmarshal    500000          5782 ns/op         564 B/op          9 allocs/op
BenchmarkBsonMarshal     1000000          1925 ns/op         488 B/op         13 allocs/op
BenchmarkBsonUnmarshal  --- FAIL: BenchmarkBsonUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{3b86c4a97a5aa287 2014-09-26 14:46:15.684430354 +1000 AEST a3ff184699 4 true 0.5503346859316104}
        &{3b86c4a97a5aa287 2014-09-26 14:46:15.684 +1000 AEST a3ff184699 4 true 0.5503346859316104}
BenchmarkVitessBsonMarshal   1000000          1417 ns/op        1169 B/op          4 allocs/op
BenchmarkVitessBsonUnmarshal    --- FAIL: BenchmarkVitessBsonUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{825f2ed8bc78185b 2014-09-26 14:46:17.126931876 +1000 AEST 929f58adf2 4 true 0.19285299476299536}
        &{825f2ed8bc78185b 2014-09-26 04:46:17.126 +0000 UTC 929f58adf2 4 true 0.19285299476299536}
BenchmarkGobMarshal   500000          6554 ns/op        1661 B/op         25 allocs/op
BenchmarkGobUnmarshal      50000         47651 ns/op       19169 B/op        367 allocs/op
BenchmarkXdrMarshal  1000000          2517 ns/op         519 B/op         15 allocs/op
BenchmarkXdrUnmarshal    1000000          2994 ns/op         369 B/op         11 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal    500000          3572 ns/op        1428 B/op         22 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal      500000          4847 ns/op        1244 B/op         31 allocs/op
BenchmarkUgorjiCodecBincMarshal   500000          4898 ns/op        2141 B/op         24 allocs/op
BenchmarkUgorjiCodecBincUnmarshal   --- FAIL: BenchmarkUgorjiCodecBincUnmarshal
    serialization_benchmarks_test.go:301: unmarshaled object differed:
        &{8ca5570b13d51126 2014-09-26 14:46:35.800449873 +1000 AEST 89522df312 2 false 0.6136756208926619}
        &{8ca5570b13d51126 2014-09-26 14:46:35.800449873 +1000 +1000 89522df312 2 false 0.6136756208926619}
BenchmarkSerealMarshal    500000          3966 ns/op        1277 B/op         21 allocs/op
BenchmarkSerealUnmarshal      500000          5543 ns/op         814 B/op         32 allocs/op
BenchmarkBinaryMarshal   1000000          2146 ns/op         479 B/op         15 allocs/op
BenchmarkBinaryUnmarshal      500000          3263 ns/op         530 B/op         19 allocs/op
BenchmarkMsgpMarshal     2000000           922 ns/op         322 B/op          4 allocs/op
BenchmarkMsgpUnmarshal   1000000          2266 ns/op         277 B/op          7 allocs/op
```

All other fields are correct however.
