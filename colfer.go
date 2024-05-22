package goserbench

import "github.com/alecthomas/go_serialization_benchmarks/goserbench"

type ColferSerializer struct {
	a ColferA
}

func (s *ColferSerializer) ForceUTC() bool {
	return true
}

func (s *ColferSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return a.MarshalBinary()
}

func (s *ColferSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a

	// Colfer requires manually claring the fields to their default value.
	*a = ColferA{}

	err = a.UnmarshalBinary(bs)
	if err != nil {
		return
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}

func NewColferSerializer() Serializer {
	return &ColferSerializer{}
}
