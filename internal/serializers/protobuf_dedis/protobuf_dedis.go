package protobufdedis

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"go.dedis.ch/protobuf"
)

type ProtobufCodec struct{ goserbench.NativeSmallStruct }

func (m ProtobufCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return protobuf.Encode(o)
}

func (m ProtobufCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return protobuf.Decode(d, o)
}

func NewProtobufCodec() goserbench.Codec[goserbench.SmallStruct] {
	return ProtobufCodec{}
}
