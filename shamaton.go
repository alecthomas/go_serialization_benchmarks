package goserbench

import (
	shamaton "github.com/shamaton/msgpack/v2"
	shamatongen "github.com/shamaton/msgpackgen/msgpack"
)

type ShamatonMapMsgpackSerializer struct{}

func (m ShamatonMapMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamaton.MarshalAsMap(o)
}

func (m ShamatonMapMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamaton.UnmarshalAsMap(d, o)
}

func NewShamatonMapMsgpackSerializer() Serializer {
	return ShamatonMapMsgpackSerializer{}
}

type ShamatonArrayMsgpackSerializer struct{}

func (m ShamatonArrayMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamaton.MarshalAsArray(o)
}

func (m ShamatonArrayMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamaton.UnmarshalAsArray(d, o)
}

func NewShamatonArrayMsgPackSerializer() Serializer {
	RegisterGeneratedResolver()
	return ShamatonArrayMsgpackSerializer{}
}

type ShamatonMapMsgpackgenSerializer struct{}

func (m ShamatonMapMsgpackgenSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamatongen.MarshalAsMap(o)
}

func (m ShamatonMapMsgpackgenSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamatongen.UnmarshalAsMap(d, o)
}

func NewShamatonMapMsgPackgenSerializer() Serializer {
	RegisterGeneratedResolver()
	return ShamatonMapMsgpackgenSerializer{}
}

type ShamatonArrayMsgpackgenSerializer struct{}

func (m ShamatonArrayMsgpackgenSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamatongen.MarshalAsArray(o)

}

func (m ShamatonArrayMsgpackgenSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamatongen.UnmarshalAsArray(d, o)
}

func NewShamatonArrayMsgpackgenSerializer() Serializer {
	return ShamatonArrayMsgpackgenSerializer{}
}
