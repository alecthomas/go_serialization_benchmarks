package easyjson

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	easyjson "github.com/mailru/easyjson"
)

type EasyJSONSerializer struct {
	a A
}

func (m EasyJSONSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	return easyjson.Marshal(a)
}

func (m EasyJSONSerializer) Unmarshal(d []byte, o interface{}) error {
	a := &m.a
	err := easyjson.Unmarshal(d, a)
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

func NewEasyJSONSerializer() goserbench.Serializer {
	return EasyJSONSerializer{}
}
