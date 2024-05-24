package binaryalecthomas

import (
	"github.com/alecthomas/binary"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o interface{}) ([]byte, error) {
	return binary.Marshal(o)
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func NewBinarySerializer() goserbench.Serializer {
	return BinarySerializer{}
}
