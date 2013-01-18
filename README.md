# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.
To run:

```
    go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson
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