package bebop200sc

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type Bebop200ScCodec struct {
	buf []byte
}

func (s *Bebop200ScCodec) Marshal(a *BebopBuf200sc) ([]byte, error) {
	if s.buf == nil {
		return a.MarshalBebop(), nil
	}
	n := a.MarshalBebopTo(s.buf)
	return s.buf[:n], nil
}

func (s *Bebop200ScCodec) Unmarshal(bs []byte, a *BebopBuf200sc) error {
	return a.UnmarshalBebop(bs)
}

func (s *Bebop200ScCodec) FromSmall(v *goserbench.SmallStruct, a *BebopBuf200sc) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (s *Bebop200ScCodec) ToSmall(a *BebopBuf200sc, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func (s *Bebop200ScCodec) TimePrecision() time.Duration {
	return 100 * time.Nanosecond
}

var _ goserbench.SerializerEnforcesTimezone = (*Bebop200ScCodec)(nil)

func (s *Bebop200ScCodec) ForcesUTC() bool {
	return true
}

func NewBebop200ScCodec() goserbench.Codec[BebopBuf200sc] {
	return &Bebop200ScCodec{}
}

func NewBebop200ScReuseCodec() goserbench.Codec[BebopBuf200sc] {
	return &Bebop200ScCodec{buf: make([]byte, 1024)}
}
