package msgpacktinylib

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type MsgpSerializer struct {
	a A
}

func (m MsgpSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	return a.MarshalMsg(nil)
}

func (m MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	a := &m.a
	_, err := a.UnmarshalMsg(d)
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
	return err
}

func NewMsgpSerializer() goserbench.Serializer {
	return MsgpSerializer{}
}
