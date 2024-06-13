# Benchmarks of Go serialization methods

[![Gitter chat](https://badges.gitter.im/alecthomas.png)](https://gitter.im/alecthomas/Lobby)

This is a test suite for benchmarking various Go serialization methods.

# Current Serialization Results

https://htmlpreview.github.io/?https://github.com/alecthomas/go_serialization_benchmarks/blob/master/report/index.html


## Running the benchmarks

```bash
go test -bench=.
```

To validate the correctness of the serializers:

```bash
VALIDATE=1 go test -bench=. -benchtime=1ms
```

To update the benchmark report:

```bash
go test -tags genreport -run TestGenerateReport
```

## Recommendation

If correctness and interoperability are the most
important factors [JSON](http://golang.org/pkg/encoding/json/) or [Protobuf](https://google.golang.org/protobuf) are your best options.

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
}
```

