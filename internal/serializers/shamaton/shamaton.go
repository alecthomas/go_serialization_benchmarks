package shamaton

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	shamaton "github.com/shamaton/msgpack/v2"
	shamatongen "github.com/shamaton/msgpackgen/msgpack"
)

type ShamatonMapMsgpackCodec struct{ goserbench.NativeSmallStruct }

func (ShamatonMapMsgpackCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return shamaton.MarshalAsMap(o)
}

func (ShamatonMapMsgpackCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return shamaton.UnmarshalAsMap(d, o)
}

func NewShamatonMapMsgpackCodec() goserbench.Codec[goserbench.SmallStruct] {
	return ShamatonMapMsgpackCodec{}
}

type ShamatonArrayMsgpackCodec struct{ goserbench.NativeSmallStruct }

func (ShamatonArrayMsgpackCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return shamaton.MarshalAsArray(o)
}

func (ShamatonArrayMsgpackCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return shamaton.UnmarshalAsArray(d, o)
}

func NewShamatonArrayMsgPackCodec() goserbench.Codec[goserbench.SmallStruct] {
	RegisterGeneratedResolver()
	return ShamatonArrayMsgpackCodec{}
}

// shamatonGenConversions converts between A and goserbench.SmallStruct for
// both generated shamaton codecs.
type shamatonGenConversions struct{}

func (shamatonGenConversions) FromSmall(v *goserbench.SmallStruct, a *A) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (shamatonGenConversions) ToSmall(a *A, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = a.Siblings
	v.Spouse = a.Spouse
	v.Money = a.Money
}

type ShamatonMapMsgpackgenCodec struct{ shamatonGenConversions }

func (ShamatonMapMsgpackgenCodec) Marshal(a *A) ([]byte, error) {
	return shamatongen.MarshalAsMap(a)
}

func (ShamatonMapMsgpackgenCodec) Unmarshal(d []byte, a *A) error {
	return shamatongen.UnmarshalAsMap(d, a)
}

func NewShamatonMapMsgPackgenCodec() goserbench.Codec[A] {
	RegisterGeneratedResolver()
	return ShamatonMapMsgpackgenCodec{}
}

type ShamatonArrayMsgpackgenCodec struct{ shamatonGenConversions }

func (ShamatonArrayMsgpackgenCodec) Marshal(a *A) ([]byte, error) {
	return shamatongen.MarshalAsArray(a)
}

func (ShamatonArrayMsgpackgenCodec) Unmarshal(d []byte, a *A) error {
	return shamatongen.UnmarshalAsArray(d, a)
}

func NewShamatonArrayMsgpackgenCodec() goserbench.Codec[A] {
	return ShamatonArrayMsgpackgenCodec{}
}
