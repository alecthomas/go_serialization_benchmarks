package gotiny

import (
	reflect "reflect"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/cybriq/gotiny"
)

type GotinySerializer struct {
	dec *gotiny.Decoder
}

func (g GotinySerializer) Marshal(o interface{}) ([]byte, error) {
	return gotiny.Marshal(o), nil
}

func (g GotinySerializer) Unmarshal(d []byte, o interface{}) error {
	g.dec.Decode(d, o)
	return nil
}

func NewGotinySerializer() goserbench.Serializer {
	ot := reflect.TypeOf(goserbench.SmallStruct{})
	return GotinySerializer{
		dec: gotiny.NewDecoderWithType(ot),
	}
}
