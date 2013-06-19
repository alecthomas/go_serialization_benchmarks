# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.
To run:

```
go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson code.google.com/p/vitess/go/bson github.com/vmihailenco/msgpack github.com/davecgh/go-xdr/xdr
go test -bench='.*' ./
```

Shameless plug: I use [pawk](https://github.com/alecthomas/pawk) to format the table:

```
go test -bench='.*' ./ | pawk '"%-40s %10s %10s %s" % f'
```

## Data

The data being serialized is the following structure with randomly generated values:

```
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

## Tested serialization methods

Currently tested are:

- `encoding/json`
- `github.com/ugorji/go-msgpack`
- `github.com/vmihailenco/msgpack`
- `labix.org/v2/mgo/bson`
- `code.google.com/p/vitess/go/bson`
- `github.com/davecgh/go-xdr/xdr`

## Results

Results on my late 2012 MacBook Air 11" are:

```
BenchmarkUgorjiMsgpackMarshal                500000       5827 ns/op
BenchmarkUgorjiMsgpackUnmarshal              200000      12132 ns/op
BenchmarkVmihailencoMsgpackMarshal           500000       3567 ns/op
BenchmarkVmihailencoMsgpackUnmarshal         200000       8110 ns/op
BenchmarkJsonMarshal                         200000      10610 ns/op
BenchmarkJsonUnmarshal                       200000       8335 ns/op
BenchmarkBsonMarshal                         500000       3262 ns/op
BenchmarkBsonUnmarshal                       500000       3880 ns/op
BenchmarkVitessBsonMarshal                   200000       8544 ns/op
BenchmarkVitessBsonUnmarshal                 500000       6566 ns/op
BenchmarkGobMarshal                          200000      11953 ns/op
BenchmarkGobUnmarshal                         20000      79726 ns/op
BenchmarkXdrMarshal                         1000000       3191 ns/op
BenchmarkXdrUnmarshal                       1000000       2347 ns/op
```

**Note:** the gob results are not really representative of normal performance, as gob is designed for serializing streams or vectors of a single type, not individual values.


## Issues

The benchmarks can also be run with validation enabled.

```
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(MAJOR)** Ugorji msgpack implementation drops the timezone frome `time.Time`.
2. **(minor)** BSON drops sub-second precision from `time.Time`.
3. **(minor)** Vitess BSON drops sub-second precision from `time.Time`.
4. **(MAJOR)** XDR does not support `time.Time` at all.

```
BenchmarkUgorjiMsgpackMarshal     500000          6096 ns/op
BenchmarkUgorjiMsgpackUnmarshal --- FAIL: BenchmarkUgorjiMsgpackUnmarshal
    serialization_benchmarks_test.go:231: unmarshaled object differed:
        &{d8c7b339f1dd290e 2013-06-18 20:41:02.839730293 -0400 EDT 9894e18711 2 true 0.9013406798206226}
        &{d8c7b339f1dd290e 2013-06-19 00:41:02.839730293 +0000 UTC 9894e18711 2 true 0.9013406798206226}
BenchmarkVmihailencoMsgpackMarshal    500000          3670 ns/op
BenchmarkVmihailencoMsgpackUnmarshal      200000          9790 ns/op
BenchmarkJsonMarshal      200000         12237 ns/op
BenchmarkJsonUnmarshal    200000         10718 ns/op
BenchmarkBsonMarshal      500000          3615 ns/op
BenchmarkBsonUnmarshal  --- FAIL: BenchmarkBsonUnmarshal
    serialization_benchmarks_test.go:231: unmarshaled object differed:
        &{07fc0beeca61082c 2013-06-18 20:41:13.613889418 -0400 EDT 617b961719 1 true 0.7691677810589427}
        &{07fc0beeca61082c 2013-06-18 20:41:13.613 -0400 EDT 617b961719 1 true 0.7691677810589427}
BenchmarkVitessBsonMarshal    200000         10272 ns/op
BenchmarkVitessBsonUnmarshal    --- FAIL: BenchmarkVitessBsonUnmarshal
    serialization_benchmarks_test.go:231: unmarshaled object differed:
        &{0a03e905be639edf 2013-06-18 20:41:15.817873329 -0400 EDT 7198c6e01e 1 true 0.9318296699538678}
        &{0a03e905be639edf 2013-06-19 00:41:15.817 +0000 UTC 7198c6e01e 1 true 0.9318296699538678}
BenchmarkGobMarshal   100000         13581 ns/op
BenchmarkGobUnmarshal      20000         89988 ns/op
BenchmarkXdrMarshal  1000000          3155 ns/op
BenchmarkXdrUnmarshal   --- FAIL: BenchmarkXdrUnmarshal
    serialization_benchmarks_test.go:231: unmarshaled object differed:
        &{aecffef8776fe751 2013-06-18 20:41:23.301829709 -0400 EDT 8a3636d0e0 3 true 0.27165374829545236}
        &{aecffef8776fe751 0001-01-01 00:00:00 +0000 UTC 8a3636d0e0 3 true 0.27165374829545236}
```

All other fields are correct however.
