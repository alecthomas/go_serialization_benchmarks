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
- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson)
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)


## Running the benchmarks

```bash
go get -u -t
go test -bench='.*' ./
```

Shameless plug: I use [pawk](https://github.com/alecthomas/pawk) to format the table:

```bash
go test -bench='.*' ./ | pawk '"%-40s %10s %10s %s" % f'
```

## Recommendation

If performance, correctness and interoperability are the most important
factors, it seems that
[github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack) is
currently the best choice.

If performance is the biggest factor, [github.com/davecgh/go-xdr](https://github.com/davecgh/go-xdr) is the
best choice.

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
BenchmarkUgorjiMsgpackMarshal                500000       3778 ns/op
BenchmarkUgorjiMsgpackUnmarshal              500000       3390 ns/op
BenchmarkVmihailencoMsgpackMarshal          1000000       1652 ns/op
BenchmarkVmihailencoMsgpackUnmarshal        1000000       1959 ns/op
BenchmarkJsonMarshal                         500000       3008 ns/op
BenchmarkJsonUnmarshal                       500000       4792 ns/op
BenchmarkBsonMarshal                        1000000       2118 ns/op
BenchmarkBsonUnmarshal                      1000000       2447 ns/op
BenchmarkVitessBsonMarshal                   200000       9219 ns/op
BenchmarkVitessBsonUnmarshal                 500000       3498 ns/op
BenchmarkGobMarshal                          200000       7051 ns/op
BenchmarkGobUnmarshal                         50000      52378 ns/op
BenchmarkXdrMarshal                         1000000       2777 ns/op
BenchmarkXdrUnmarshal                       1000000       1910 ns/op
BenchmarkUgorjiCodecMsgpackMarshal           500000       3831 ns/op
BenchmarkUgorjiCodecMsgpackUnmarshal         500000       3985 ns/op
BenchmarkUgorjiCodecBincMarshal              500000       5327 ns/op
BenchmarkUgorjiCodecBincUnmarshal            500000       5336 ns/op
BenchmarkSerealMarshal                       500000       4542 ns/op
BenchmarkSerealUnmarshal                     500000       4254 ns/op
BenchmarkBinaryMarshal                      1000000       2375 ns/op
BenchmarkBinaryUnmarshal                    1000000       2039 ns/op
```

**Note:** the gob results are not really representative of normal performance, as gob is designed for serializing streams or vectors of a single type, not individual values.


## Issues

The benchmarks can also be run with validation enabled.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(MAJOR)** Ugorji msgpack implementation drops the timezone frome `time.Time`.
2. **(minor)** BSON drops sub-second precision from `time.Time`.
3. **(minor)** Vitess BSON drops sub-second precision from `time.Time`.
4. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
BenchmarkUgorjiMsgpackMarshal     500000          3725 ns/op
BenchmarkUgorjiMsgpackUnmarshal --- FAIL: BenchmarkUgorjiMsgpackUnmarshal
    serialization_benchmarks_test.go:295: unmarshaled object differed:
        &{d8c7b339f1dd290e 2014-04-29 13:48:36.84484582 -0400 EDT 9894e18711 2 true 0.9013406798206226}
        &{d8c7b339f1dd290e 2014-04-29 17:48:36.84484582 +0000 UTC 9894e18711 2 true 0.9013406798206226}
BenchmarkVmihailencoMsgpackMarshal   1000000          1705 ns/op
BenchmarkVmihailencoMsgpackUnmarshal      500000          3238 ns/op
BenchmarkJsonMarshal      500000          3168 ns/op
BenchmarkJsonUnmarshal    500000          6242 ns/op
BenchmarkBsonMarshal     1000000          2072 ns/op
BenchmarkBsonUnmarshal  --- FAIL: BenchmarkBsonUnmarshal
    serialization_benchmarks_test.go:295: unmarshaled object differed:
        &{ab3ef8e215a0f8d9 2014-04-29 13:48:47.184790536 -0400 EDT affd248591 2 true 0.2546601960244086}
        &{ab3ef8e215a0f8d9 2014-04-29 13:48:47.184 -0400 EDT affd248591 2 true 0.2546601960244086}
BenchmarkVitessBsonMarshal    200000          8900 ns/op
BenchmarkVitessBsonUnmarshal    --- FAIL: BenchmarkVitessBsonUnmarshal
    serialization_benchmarks_test.go:295: unmarshaled object differed:
        &{fc885bace3d98d87 2014-04-29 13:48:49.067470923 -0400 EDT 0b02f4d722 2 false 0.7670348077372372}
        &{fc885bace3d98d87 2014-04-29 17:48:49.067 +0000 UTC 0b02f4d722 2 false 0.7670348077372372}
BenchmarkGobMarshal   200000          7054 ns/op
BenchmarkGobUnmarshal      50000         52048 ns/op
BenchmarkXdrMarshal  1000000          2657 ns/op
BenchmarkXdrUnmarshal    1000000          2953 ns/op
BenchmarkUgorjiCodecMsgpackMarshal    500000          3606 ns/op
BenchmarkUgorjiCodecMsgpackUnmarshal      500000          5073 ns/op
BenchmarkUgorjiCodecBincMarshal   500000          4984 ns/op
BenchmarkUgorjiCodecBincUnmarshal   --- FAIL: BenchmarkUgorjiCodecBincUnmarshal
    serialization_benchmarks_test.go:295: unmarshaled object differed:
        &{ec0d97b0656e8246 2014-04-29 13:49:06.446070434 -0400 EDT 7de36bc617 1 true 0.3675096311616762}
        &{ec0d97b0656e8246 2014-04-29 13:49:06.446070434 -0400 -0400 7de36bc617 1 true 0.3675096311616762}
BenchmarkSerealMarshal    500000          4278 ns/op
BenchmarkSerealUnmarshal      500000          5281 ns/op
BenchmarkBinaryMarshal   1000000          2207 ns/op
BenchmarkBinaryUnmarshal      500000          3188 ns/op
```

All other fields are correct however.
