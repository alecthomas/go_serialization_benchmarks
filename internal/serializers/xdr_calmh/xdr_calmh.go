package xdrcalmh

import (
	"math"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type XDRCalmhCodec struct{}

func (XDRCalmhCodec) Marshal(a *XDRA) ([]byte, error) {
	return a.MarshalXDR()
}

func (XDRCalmhCodec) Unmarshal(bs []byte, a *XDRA) error {
	return a.UnmarshalXDR(bs)
}

func (XDRCalmhCodec) FromSmall(v *goserbench.SmallStruct, a *XDRA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
}

func (XDRCalmhCodec) ToSmall(a *XDRA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
}

func NewXDRCalmhCodec() goserbench.Codec[XDRA] {
	return XDRCalmhCodec{}
}
