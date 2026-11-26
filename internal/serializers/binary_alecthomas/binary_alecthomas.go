package binaryalecthomas

import (
	"github.com/alecthomas/binary"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o any) ([]byte, error) {
	return binary.Marshal(o)
}

func (b BinarySerializer) Unmarshal(d []byte, o any) error {
	return binary.Unmarshal(d, o)
}

func NewBinarySerializer() goserbench.Serializer {
	return BinarySerializer{}
}
