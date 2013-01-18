# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.
To run:

```
    go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson
    go test -bench='.*' ./

```