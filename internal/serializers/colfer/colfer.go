package colfer

import "github.com/alecthomas/go_serialization_benchmarks/goserbench"

type ColferCodec struct{}

var _ goserbench.SerializerEnforcesTimezone = ColferCodec{}

func (ColferCodec) ForcesUTC() bool {
	return true
}

func (ColferCodec) Marshal(a *ColferA) ([]byte, error) {
	return a.MarshalBinary()
}

// Unmarshal relies on the harness zeroing a first, because Colfer does not
// reset fields that are absent from the input.
func (ColferCodec) Unmarshal(bs []byte, a *ColferA) error {
	return a.UnmarshalBinary(bs)
}

func (ColferCodec) FromSmall(v *goserbench.SmallStruct, a *ColferA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (ColferCodec) ToSmall(a *ColferA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewColferCodec() goserbench.Codec[ColferA] {
	return ColferCodec{}
}
