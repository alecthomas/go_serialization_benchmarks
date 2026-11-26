package sereal

import (
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type SerealCodec struct{ goserbench.NativeSmallStruct }

func (s SerealCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return sereal.Marshal(o)
}

func (s SerealCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func NewSerealCodec() goserbench.Codec[goserbench.SmallStruct] {
	return SerealCodec{}
}
