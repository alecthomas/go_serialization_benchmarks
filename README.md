# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

To run:

```bash
go get -u \
    github.com/ugorji/go-msgpack \
    labix.org/v2/mgo/bson \
    code.google.com/p/vitess/go/bson \
    github.com/vmihailenco/msgpack \
    github.com/davecgh/go-xdr/xdr \
    github.com/ugorji/go/codec
go test -bench='.*' ./
```

Shameless plug: I use [pawk](https://github.com/alecthomas/pawk) to format the table:

```bash
go test -bench='.*' ./ | pawk '"%-40s %10s %10s %s" % f'
```

## Recommendation

If performance and correctness are the most important factor, it seems that https://github.com/vmihailenco/msgpack is currently the best choice.

But as always, make your own choice based on your own requirements.

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

## Tested serialization methods

Currently tested are:

- `encoding/json`
- `github.com/ugorji/go-msgpack`
- `github.com/vmihailenco/msgpack`
- `labix.org/v2/mgo/bson`
- `code.google.com/p/vitess/go/bson`
- `github.com/davecgh/go-xdr/xdr`
- `github.com/ugorji/go/codec`

## Results

Results on my late 2012 MacBook Air 11" are:

```
BenchmarkUgorjiMsgpackMarshal                500000       4882 ns/op
BenchmarkUgorjiMsgpackUnmarshal              500000       4581 ns/op
BenchmarkVmihailencoMsgpackMarshal          1000000       2256 ns/op
BenchmarkVmihailencoMsgpackUnmarshal        1000000       2759 ns/op
BenchmarkJsonMarshal                         500000       4608 ns/op
BenchmarkJsonUnmarshal                       500000       7794 ns/op
BenchmarkBsonMarshal                        1000000       2640 ns/op
BenchmarkBsonUnmarshal                       500000       3114 ns/op
BenchmarkVitessBsonMarshal                   200000      10754 ns/op
BenchmarkVitessBsonUnmarshal                 500000       4169 ns/op
BenchmarkGobMarshal                          200000      11297 ns/op
BenchmarkGobUnmarshal                         20000      76537 ns/op
BenchmarkXdrMarshal                         1000000       2839 ns/op
BenchmarkXdrUnmarshal                       1000000       2068 ns/op
BenchmarkUgorjiCodecMsgpackMarshal           500000       5685 ns/op
BenchmarkUgorjiCodecMsgpackUnmarshal         500000       5491 ns/op
BenchmarkUgorjiCodecBincMarshal              500000       7927 ns/op
BenchmarkUgorjiCodecBincUnmarshal            200000       7624 ns/op
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
4. **(MAJOR)** XDR does not support `time.Time` at all.
5. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
BenchmarkUgorjiMsgpackMarshal     500000          4919 ns/op
BenchmarkUgorjiMsgpackUnmarshal --- FAIL: BenchmarkUgorjiMsgpackUnmarshal
    serialization_benchmarks_test.go:261: unmarshaled object differed:
        &{d8c7b339f1dd290e 2013-12-14 10:25:28.377646823 -0500 EST 9894e18711 2 true 0.9013406798206226}
        &{d8c7b339f1dd290e 2013-12-14 15:25:28.377646823 +0000 UTC 9894e18711 2 true 0.9013406798206226}
BenchmarkVmihailencoMsgpackMarshal   1000000          2137 ns/op
BenchmarkVmihailencoMsgpackUnmarshal      500000          4633 ns/op
BenchmarkJsonMarshal      500000          4702 ns/op
BenchmarkJsonUnmarshal    200000          9538 ns/op
BenchmarkBsonMarshal     1000000          2627 ns/op
BenchmarkBsonUnmarshal  --- FAIL: BenchmarkBsonUnmarshal
    serialization_benchmarks_test.go:261: unmarshaled object differed:
        &{63acca08873729ff 2013-12-14 10:25:40.050202084 -0500 EST ef8cbd5b6d 4 true 0.4770114308682012}
        &{63acca08873729ff 2013-12-14 10:25:40.05 -0500 EST ef8cbd5b6d 4 true 0.4770114308682012}
BenchmarkVitessBsonMarshal    200000         10774 ns/op
BenchmarkVitessBsonUnmarshal    --- FAIL: BenchmarkVitessBsonUnmarshal
    serialization_benchmarks_test.go:261: unmarshaled object differed:
        &{00e3e3b19864179d 2013-12-14 10:25:42.332382225 -0500 EST c36680fc0a 0 false 0.7249153738920362}
        &{00e3e3b19864179d 2013-12-14 15:25:42.332 +0000 UTC c36680fc0a 0 false 0.7249153738920362}
BenchmarkGobMarshal   200000          9776 ns/op
BenchmarkGobUnmarshal      50000         71520 ns/op
BenchmarkXdrMarshal  1000000          2548 ns/op
BenchmarkXdrUnmarshal   --- FAIL: BenchmarkXdrUnmarshal
    serialization_benchmarks_test.go:261: unmarshaled object differed:
        &{9a5709db87d05551 2013-12-14 10:25:51.335978215 -0500 EST 812e29e037 4 false 0.05093211967185806}
        &{9a5709db87d05551 0001-01-01 00:00:00 +0000 UTC 812e29e037 4 false 0.05093211967185806}
BenchmarkUgorjiCodecMsgpackMarshal    500000          4841 ns/op
BenchmarkUgorjiCodecMsgpackUnmarshal      500000          7104 ns/op
BenchmarkUgorjiCodecBincMarshal   500000          6717 ns/op
BenchmarkUgorjiCodecBincUnmarshal   --- FAIL: BenchmarkUgorjiCodecBincUnmarshal
    serialization_benchmarks_test.go:261: unmarshaled object differed:
        &{c9ad6275a298bd49 2013-12-14 10:26:00.92156572 -0500 EST a34a9bfe58 4 false 0.09946586035661706}
        &{c9ad6275a298bd49 2013-12-14 10:26:00.92156572 -0500 -0500 a34a9bfe58 4 false 0.09946586035661706}```

All other fields are correct however.
