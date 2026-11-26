package easyjson

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	easyjson "github.com/mailru/easyjson"
)

type EasyJSONCodec struct{}

func (EasyJSONCodec) Marshal(a *A) ([]byte, error) {
	return easyjson.Marshal(a)
}

func (EasyJSONCodec) Unmarshal(d []byte, a *A) error {
	return easyjson.Unmarshal(d, a)
}

func (EasyJSONCodec) FromSmall(v *goserbench.SmallStruct, a *A) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (EasyJSONCodec) ToSmall(a *A, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = a.Siblings
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewEasyJSONCodec() goserbench.Codec[A] {
	return EasyJSONCodec{}
}
