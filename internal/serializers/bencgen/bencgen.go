package bencgen

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

//go:generate go run github.com/deneonet/benc/cmd/bencgen@v1.1.8 --in structdef.benc --lang go --out . --file structdef

type BENCGenCodec struct{}

func (BENCGenCodec) Marshal(a *SmallStruct) ([]byte, error) {
	buf := make([]byte, a.Size())
	a.Marshal(buf)
	return buf, nil
}

func (BENCGenCodec) Unmarshal(bs []byte, a *SmallStruct) error {
	return a.Unmarshal(bs)
}

func (BENCGenCodec) FromSmall(v *goserbench.SmallStruct, a *SmallStruct) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (BENCGenCodec) ToSmall(a *SmallStruct, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewBENCGenCodec() goserbench.Codec[SmallStruct] {
	return BENCGenCodec{}
}

type BENCGenUnsafeCodec struct{}

func (BENCGenUnsafeCodec) Marshal(a *SmallStructUnsafe) ([]byte, error) {
	buf := make([]byte, a.Size())
	a.Marshal(buf)
	return buf, nil
}

func (BENCGenUnsafeCodec) Unmarshal(bs []byte, a *SmallStructUnsafe) error {
	return a.Unmarshal(bs)
}

func (BENCGenUnsafeCodec) FromSmall(v *goserbench.SmallStruct, a *SmallStructUnsafe) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (BENCGenUnsafeCodec) ToSmall(a *SmallStructUnsafe, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewBENCGenUnsafeCodec() goserbench.Codec[SmallStructUnsafe] {
	return BENCGenUnsafeCodec{}
}
