package shamaton

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
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

func NewShamatonMapMsgpackSerializer() goserbench.Serializer {
	return ShamatonMapMsgpackSerializer{}
}

type ShamatonArrayMsgpackSerializer struct{}

func (m ShamatonArrayMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamaton.MarshalAsArray(o)
}

func (m ShamatonArrayMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamaton.UnmarshalAsArray(d, o)
}

func NewShamatonArrayMsgPackSerializer() goserbench.Serializer {
	RegisterGeneratedResolver()
	return ShamatonArrayMsgpackSerializer{}
}

type ShamatonMapMsgpackgenSerializer struct {
	a A
}

func (m ShamatonMapMsgpackgenSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	return shamatongen.MarshalAsMap(a)
}

func (m ShamatonMapMsgpackgenSerializer) Unmarshal(d []byte, o interface{}) error {
	a := &m.a
	err := shamatongen.UnmarshalAsMap(d, a)
	if err != nil {
		return err
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return nil
}

func NewShamatonMapMsgPackgenSerializer() goserbench.Serializer {
	RegisterGeneratedResolver()
	return ShamatonMapMsgpackgenSerializer{}
}

type ShamatonArrayMsgpackgenSerializer struct {
	a A
}

func (m ShamatonArrayMsgpackgenSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	return shamatongen.MarshalAsArray(a)

}

func (m ShamatonArrayMsgpackgenSerializer) Unmarshal(d []byte, o interface{}) error {
	a := &m.a
	err := shamatongen.UnmarshalAsArray(d, a)
	if err != nil {
		return err
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return nil
}

func NewShamatonArrayMsgpackgenSerializer() goserbench.Serializer {
	return ShamatonArrayMsgpackgenSerializer{}
}
