package stdlib

import (
	"encoding/json"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type JsonCodec struct{ goserbench.NativeSmallStruct }

func (j JsonCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return json.Marshal(o)
}

func (j JsonCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return json.Unmarshal(d, o)
}

func NewJSONCodec() goserbench.Codec[goserbench.SmallStruct] {
	return JsonCodec{}
}
