package protobufdedis

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"go.dedis.ch/protobuf"
)

type ProtobufSerializer struct{}

func (m ProtobufSerializer) Marshal(o interface{}) ([]byte, error) {
	return protobuf.Encode(o)
}

func (m ProtobufSerializer) Unmarshal(d []byte, o interface{}) error {
	return protobuf.Decode(d, o)
}

func NewProtobufSerializer() goserbench.Serializer {
	return ProtobufSerializer{}
}
