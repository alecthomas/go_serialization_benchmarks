package goserbench

import (
	reflect "reflect"

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

func NewGotinySerializer() Serializer {
	ot := reflect.TypeOf(A{})
	return GotinySerializer{
		dec: gotiny.NewDecoderWithType(ot),
	}
}
