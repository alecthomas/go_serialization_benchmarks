package gencode

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type GencodeCodec struct{}

func (GencodeCodec) Marshal(a *GencodeA) ([]byte, error) {
	return a.Marshal(nil)
}

func (GencodeCodec) Unmarshal(bs []byte, a *GencodeA) error {
	_, err := a.Unmarshal(bs)
	return err
}

func (GencodeCodec) FromSmall(v *goserbench.SmallStruct, a *GencodeA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (GencodeCodec) ToSmall(a *GencodeA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewGencodeCodec() goserbench.Codec[GencodeA] {
	return GencodeCodec{}
}

type GencodeUnsafeCodec struct {
	buf []byte
}

func (s *GencodeUnsafeCodec) Marshal(a *GencodeUnsafeA) ([]byte, error) {
	return a.Marshal(s.buf)
}

func (s *GencodeUnsafeCodec) Unmarshal(bs []byte, a *GencodeUnsafeA) error {
	_, err := a.Unmarshal(bs)
	return err
}

func (s *GencodeUnsafeCodec) FromSmall(v *goserbench.SmallStruct, a *GencodeUnsafeA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (s *GencodeUnsafeCodec) ToSmall(a *GencodeUnsafeA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewGencodeUnsafeCodec() goserbench.Codec[GencodeUnsafeA] {
	return &GencodeUnsafeCodec{buf: make([]byte, 0, 1024)}
}
