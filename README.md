# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.
To run:

```
go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson code.google.com/p/vitess/go/bson github.com/vmihailenco/msgpack
go test -bench='.*' ./
```

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
BenchmarkMsgpackMarshal   			200000             10578 ns/op
BenchmarkMsgpackUnmarshal         	100000             24522 ns/op
BenchmarkJsonMarshal      			100000             19423 ns/op
BenchmarkJsonUnmarshal    			100000             24113 ns/op
BenchmarkBsonMarshal      			200000             10657 ns/op
BenchmarkBsonUnmarshal    			100000             16348 ns/op
BenchmarkVitessBsonMarshal        	100000             14922 ns/op
BenchmarkVitessBsonUnmarshal      	100000             16769 ns/op
```
