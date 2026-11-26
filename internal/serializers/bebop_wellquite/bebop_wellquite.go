package bebopwellquite

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type BebopWellquiteCodec struct {
	buf []byte
}

func (s *BebopWellquiteCodec) Marshal(a *BebopBufWellquite) ([]byte, error) {
	return a.MarshalBebop(s.buf)
}

func (s *BebopWellquiteCodec) Unmarshal(bs []byte, a *BebopBufWellquite) error {
	_, err := a.UnmarshalBebop(bs)
	return err
}

func (s *BebopWellquiteCodec) FromSmall(v *goserbench.SmallStruct, a *BebopBufWellquite) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (s *BebopWellquiteCodec) ToSmall(a *BebopBufWellquite, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func (s *BebopWellquiteCodec) TimePrecision() time.Duration {
	return 100 * time.Nanosecond
}

func NewBebopWellquiteCodec() goserbench.Codec[BebopBufWellquite] {
	return &BebopWellquiteCodec{}
}

func NewBebopWellquiteReuseCodec() goserbench.Codec[BebopBufWellquite] {
	return &BebopWellquiteCodec{
		buf: make([]byte, 1024),
	}
}
