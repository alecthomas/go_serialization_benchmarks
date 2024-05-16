package goserbench

import (
	"time"

	pproto "google.golang.org/protobuf/proto"
)

type PulsarSerializer struct {
	a PulsarBufA
}

func (s *PulsarSerializer) ForceUTC() bool {
	return true
}

func (s *PulsarSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return pproto.Marshal(a)
}

func (s *PulsarSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a

	// Pulsar requires manually claring the fields to their default value.
	// *a = PulsarA{}

	err = pproto.Unmarshal(bs, a)
	if err != nil {
		return
	}

	v := o.(*A)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}

func NewPulsarSerializer() *PulsarSerializer {
	return &PulsarSerializer{}
}
