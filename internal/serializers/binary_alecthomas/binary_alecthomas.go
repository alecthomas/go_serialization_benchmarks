package binaryalecthomas

import (
	"github.com/alecthomas/binary"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type BinaryCodec struct{ goserbench.NativeSmallStruct }

func (b BinaryCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return binary.Marshal(o)
}

func (b BinaryCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return binary.Unmarshal(d, o)
}

func NewBinaryCodec() goserbench.Codec[goserbench.SmallStruct] {
	return BinaryCodec{}
}
