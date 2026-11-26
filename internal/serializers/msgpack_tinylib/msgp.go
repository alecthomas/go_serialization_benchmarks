package msgpacktinylib

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type MsgpCodec struct{}

func (MsgpCodec) Marshal(a *A) ([]byte, error) {
	return a.MarshalMsg(nil)
}

func (MsgpCodec) Unmarshal(d []byte, a *A) error {
	_, err := a.UnmarshalMsg(d)
	return err
}

func (MsgpCodec) FromSmall(v *goserbench.SmallStruct, a *A) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (MsgpCodec) ToSmall(a *A, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = a.Siblings
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewMsgpCodec() goserbench.Codec[A] {
	return MsgpCodec{}
}
