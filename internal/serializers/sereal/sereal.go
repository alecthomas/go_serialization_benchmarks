package sereal

import (
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o any) ([]byte, error) {
	return sereal.Marshal(o)
}

func (s SerealSerializer) Unmarshal(d []byte, o any) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func NewSerealSerializer() goserbench.Serializer {
	return SerealSerializer{}
}
