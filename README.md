# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.
To run:

```
go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson code.google.com/p/vitess/go/bson
go test -bench='.*' ./
```

No testing for serialization correctness is performed.

## Serialization methods

Currently tested are:

- `encoding/json`
- `encoding/xml`
- `github.com/ugorji/go-msgpack`
- `labix.org/v2/mgo/bson`
- `code.google.com/p/vitess/go/bson`


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
BenchmarkMsgpackMarshal	  		500000	      5190 ns/op
BenchmarkMsgpackUnmarshal	  	200000	     10695 ns/op
BenchmarkJsonMarshal	  		200000	      8365 ns/op
BenchmarkJsonUnmarshal	  		200000	     10472 ns/op
BenchmarkXmlMarshal	  			200000	      8007 ns/op
BenchmarkXmlUnmarshal	   		50000	     34064 ns/op
BenchmarkBsonMarshal	  		500000	      3734 ns/op
BenchmarkBsonUnmarshal	  		500000	      3934 ns/op
BenchmarkVitessBsonMarshal	  	200000	      8507 ns/op
BenchmarkVitessBsonUnmarshal	500000	      5627 ns/op
```