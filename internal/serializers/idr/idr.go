package idr

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/chmike/ditp/idr/low"
)

type IDRSerializer struct {
}

func (s IDRSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	e := make([]byte, 0, 64)
	e = low.AppendTime(e, v.BirthDay)
	e = low.AppendFloat64(e, v.Money)
	e = low.AppendBool(e, v.Spouse)
	e = low.AppendVarInt(e, v.Siblings)
	e = low.AppendString(e, v.Name)
	e = low.AppendString(e, v.Phone)
	return e, nil
}

func (s IDRSerializer) Unmarshal(d []byte, o interface{}) (err error) {
	v := o.(*goserbench.SmallStruct)
	d, v.BirthDay = low.Time(d)
	d, v.Money = low.Float64(d)
	d, v.Spouse = low.Bool(d)
	d, v.Siblings = low.VarInt(d)
	d, v.Name = low.String(d, 1<<14)
	_, v.Phone = low.String(d, 20)
	return nil
}

func NewIDRSerializer() goserbench.Serializer {
	return IDRSerializer{}
}

type IDRSerializerReuse struct {
	e low.Encoder
}

func (s IDRSerializerReuse) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	e := low.Reset(s.e)
	e = low.AppendTime(e, v.BirthDay)
	e = low.AppendFloat64(e, v.Money)
	e = low.AppendBool(e, v.Spouse)
	e = low.AppendVarInt(e, v.Siblings)
	e = low.AppendString(e, v.Name)
	e = low.AppendString(e, v.Phone)
	return e, nil
}

func (s IDRSerializerReuse) Unmarshal(d []byte, o interface{}) (err error) {
	v := o.(*goserbench.SmallStruct)
	d, v.BirthDay = low.Time(d)
	d, v.Money = low.Float64(d)
	d, v.Spouse = low.Bool(d)
	d, v.Siblings = low.VarInt(d)
	d, v.Name = low.String(d, 1<<14)
	_, v.Phone = low.String(d, 20)
	return nil
}

func NewIDRSerializerReuse() goserbench.Serializer {
	return IDRSerializerReuse{e: make([]byte, 256)} // set initial buffer to 256 bytes
}
