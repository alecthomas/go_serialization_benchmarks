package sereal

import (
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o interface{}) ([]byte, error) {
	return sereal.Marshal(o)
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func NewSerealSerializer() goserbench.Serializer {
	return SerealSerializer{}
}
