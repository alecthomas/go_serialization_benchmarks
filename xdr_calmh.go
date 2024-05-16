package goserbench

import (
	"math"
	"time"
)

type XDRCalmhSerializer struct {
	a XDRA
}

func (s *XDRCalmhSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
	return a.MarshalXDR()
}

func (s *XDRCalmhSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	err = a.UnmarshalXDR(bs)
	if err != nil {
		return
	}

	v := o.(*A)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
	return
}

func NewXDRCalmhSerializer() Serializer {
	return &XDRCalmhSerializer{}
}
