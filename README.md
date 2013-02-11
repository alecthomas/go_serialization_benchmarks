# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.
To run:

```
go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson code.google.com/p/vitess/go/bson github.com/vmihailenco/msgpack
go test -bench='.*' ./
```

Shameless plug: I use [pawk](https://github.com/alecthomas/pawk) to format the table:

`go test -bench='.*' ./ | pawk '"%-40s %10s %10s %s" % f'`)

No testing for serialization correctness is performed.

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
- `labix.org/v2/mgo/bson`
- `code.google.com/p/vitess/go/bson`
- `github.com/vmihailenco/msgpack`


## Caveats

Note that mgo's bson serializer seems to drop some precision with `time.Date` structs. Uncommenting validation code in the benchmark fails with:

```
serialization_benchmarks_test.go:156: 	unmarshaled object differed:
&{d8c7b339f1dd290e 2013-01-18 11:39:39.64337014 -0500 EST 9894e18711 2 true 0.9013406798206226}
&{d8c7b339f1dd290e 2013-01-18 11:39:39.643 -0500 EST 9894e18711 2 true 0.9013406798206226}
```

All other fields are correct however.


## Results

Results on my late 2012 MacBook Air 11" are:

```
BenchmarkUgorjiMsgpackMarshal                500000       5809 ns/op
BenchmarkUgorjiMsgpackUnmarshal              200000      11027 ns/op
BenchmarkVmihailencoMsgpackMarshal           500000       3383 ns/op
BenchmarkVmihailencoMsgpackUnmarshal         500000       7054 ns/op
BenchmarkJsonMarshal                         200000       9762 ns/op
BenchmarkJsonUnmarshal                       200000      10897 ns/op
BenchmarkBsonMarshal                         500000       4437 ns/op
BenchmarkBsonUnmarshal                       500000       4477 ns/op
BenchmarkVitessBsonMarshal                   200000       8564 ns/op
BenchmarkVitessBsonUnmarshal                 500000       6553 ns/op
BenchmarkGobMarshal                          200000      12531 ns/op
BenchmarkGobUnmarshal                         20000      99961 ns/op
```

**Note:** the gob results are not really representative of normal performance, as gob is designed for serializing streams or vectors of a single type, not individual values.
