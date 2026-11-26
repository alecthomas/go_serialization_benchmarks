package gotiny

import (
	reflect "reflect"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/cybriq/gotiny"
)

type GotinyCodec struct {
	goserbench.NativeSmallStruct

	dec *gotiny.Decoder
}

func (g GotinyCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return gotiny.Marshal(o), nil
}

func (g GotinyCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	g.dec.Decode(d, o)
	return nil
}

func NewGotinyCodec() goserbench.Codec[goserbench.SmallStruct] {
	ot := reflect.TypeOf(goserbench.SmallStruct{})
	return GotinyCodec{
		dec: gotiny.NewDecoderWithType(ot),
	}
}
