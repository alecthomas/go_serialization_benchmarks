package idr

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/chmike/ditp/idr/low"
)

type IDRCodec struct {
	goserbench.NativeSmallStruct

}

func (s IDRCodec) Marshal(v *goserbench.SmallStruct) ([]byte, error) {
	e := make([]byte, 0, 64)
	e = low.PutTime(e, v.BirthDay)
	e = low.PutFloat64(e, v.Money)
	e = low.PutBool(e, v.Spouse)
	e = low.PutVarInt(e, v.Siblings)
	e = low.PutString(e, v.Name)
	e = low.PutString(e, v.Phone)
	return e, nil
}

func (s IDRCodec) Unmarshal(d []byte, v *goserbench.SmallStruct) (err error) {
	d, v.BirthDay = low.Time(d)
	d, v.Money = low.Float64(d)
	d, v.Spouse = low.Bool(d)
	d, v.Siblings = low.VarInt(d)
	d, v.Name = low.String(d, 1<<14)
	_, v.Phone = low.String(d, 20)
	return nil
}

func NewIDRCodec() goserbench.Codec[goserbench.SmallStruct] {
	return IDRCodec{}
}

type IDRReuseCodec struct {
	goserbench.NativeSmallStruct

	e low.Encoder
}

func (s IDRReuseCodec) Marshal(v *goserbench.SmallStruct) ([]byte, error) {
	e := low.Reset(s.e)
	e = low.PutTime(e, v.BirthDay)
	e = low.PutFloat64(e, v.Money)
	e = low.PutBool(e, v.Spouse)
	e = low.PutVarInt(e, v.Siblings)
	e = low.PutString(e, v.Name)
	e = low.PutString(e, v.Phone)
	return e, nil
}

func (s IDRReuseCodec) Unmarshal(d []byte, v *goserbench.SmallStruct) (err error) {
	d, v.BirthDay = low.Time(d)
	d, v.Money = low.Float64(d)
	d, v.Spouse = low.Bool(d)
	d, v.Siblings = low.VarInt(d)
	d, v.Name = low.String(d, 1<<14)
	_, v.Phone = low.String(d, 20)
	return nil
}

func NewIDRReuseCodec() goserbench.Codec[goserbench.SmallStruct] {
	return IDRReuseCodec{e: make([]byte, 256)} // set initial buffer to 256 bytes
}
