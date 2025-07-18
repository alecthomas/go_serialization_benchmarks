package goccyjson

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	json "github.com/goccy/go-json"
)

type GoccyJsonSerializer struct{}

func (m GoccyJsonSerializer) TimePrecision() time.Duration {
	return time.Millisecond
}

func (m GoccyJsonSerializer) Marshal(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func (m GoccyJsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func NewGoccyJsonSerializer() goserbench.Serializer {
	return GoccyJsonSerializer{}
}
