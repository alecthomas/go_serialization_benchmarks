package stdlib

import (
	"encoding/json"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type JsonSerializer struct{}

func (j JsonSerializer) Marshal(o any) ([]byte, error) {
	return json.Marshal(o)
}

func (j JsonSerializer) Unmarshal(d []byte, o any) error {
	return json.Unmarshal(d, o)
}

func NewJSONSerializer() goserbench.Serializer {
	return JsonSerializer{}
}
