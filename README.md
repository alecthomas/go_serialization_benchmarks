# Benchmarks of Go serialization methods

[![Gitter chat](https://badges.gitter.im/alecthomas.png)](https://gitter.im/alecthomas/Lobby)

This is a test suite for benchmarking various Go serialization methods.

# Current Serialization Results

https://alecthomas.github.io/go_serialization_benchmarks

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

To update the benchmark report with using a longer run benchmark (to get more
accurate results):

```bash
go test -tags genreport -run TestGenerateReport -benchtime 10s -timeout 1h
```

## Recommendation

If correctness and interoperability are the most
important factors [JSON](http://golang.org/pkg/encoding/json/) or [Protobuf](https://google.golang.org/protobuf) are your best options.

But as always, make your own choice based on your requirements.

## Adding New Serializers

Review the following instructions _before_ opening the PR to add a new
serializer:

- Create all the required serializer code in `internal/<short serializer name>`.
- Add an entry to the serializer in [benchmarks.go](benchmarks.go).
- If the serializer supports both reusing/not reusing its marshalling buffer:
  - Add both a `serializer` and `serializer/reuse` entries, each one
    respectively reusing/not reusing the resulting marshalling buffer. Set the
    `BufferReuseMarshal` flag accordingly.
- If the serializer supports both safe and unsafe string unmarshalling:
  - Add both a `serializer` and `serializer/unsafe` entries, each one
    respectively unmarshalling into safe and unsafe strings. Set the
    `UnsafeStringUnmarshal` flag accordingly.
- If the serializer supports both marshalling buffer reuse and unsafe string
  unmarshalling, merge both options into a single `serializer/unsafe_reuse`
  entry (check the baseline serializer for an example).
- Regenerate the report by running:

```
go test -tags genreport -run TestGenerateReport
```

- **Include the updated report data in your PR**

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

