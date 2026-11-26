package jsoniter

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	jsoniter "github.com/json-iterator/go"
)

var (
	jsoniterFast = jsoniter.ConfigFastest
)

type JsonIterCodec struct{ goserbench.NativeSmallStruct }

func (j JsonIterCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return jsoniterFast.Marshal(o)
}

func (j JsonIterCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return jsoniterFast.Unmarshal(d, o)
}

func (j JsonIterCodec) ReduceFloat64Precision() uint {
	return 6
}

func NewJSONIterCodec() goserbench.Codec[goserbench.SmallStruct] {
	return JsonIterCodec{}
}
