package bebop200sc

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type Bebop200ScSerializer struct {
	a   BebopBuf200sc
	buf []byte
}

func (s *Bebop200ScSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	n := a.MarshalBebopTo(s.buf)
	return s.buf[:n], nil
}

func (s *Bebop200ScSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	err = a.UnmarshalBebop(bs)
	if err != nil {
		return
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}

func (s *Bebop200ScSerializer) TimePrecision() time.Duration {
	return 100 * time.Nanosecond
}

func (s *Bebop200ScSerializer) ForceUTC() bool {
	return true
}

func NewBebop200ScSerializer() goserbench.Serializer {
	return &Bebop200ScSerializer{buf: make([]byte, 1024)}
}
