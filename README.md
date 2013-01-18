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

## Results

Results on my late 2012 MacBook Air 11" are:

```
BenchmarkMsgpackMarshal	  	500000	      5153 ns/op
BenchmarkMsgpackUnmarshal 	200000	     10342 ns/op
BenchmarkJsonMarshal	  	200000	      8321 ns/op
BenchmarkJsonUnmarshal	  	200000	     10411 ns/op
BenchmarkXmlMarshal	  		200000	      7666 ns/op
BenchmarkXmlUnmarshal	   	50000	     34113 ns/op
BenchmarkBsonMarshal	  	500000	      3757 ns/op
BenchmarkBsonUnmarshal	  	500000	      4083 ns/op
```