package bencgen

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

//go:generate go run github.com/deneonet/benc/cmd/bencgen@v1.1.8 --in structdef.benc --lang go --out . --file structdef

type BENCGenSerializer struct {
	a SmallStruct
}

func (s *BENCGenSerializer) Marshal(o any) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	buf := make([]byte, a.Size())
	a.Marshal(buf)
	return buf, nil
}

func (s *BENCGenSerializer) Unmarshal(bs []byte, o any) error {
	a := &s.a
	if err := a.Unmarshal(bs); err != nil {
		return err
	}
	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return nil
}

func NewBENCGenSerializer() goserbench.Serializer {
	return &BENCGenSerializer{}
}

type BENCGenUnsafeSerializer struct {
	a SmallStructUnsafe
}

func (s *BENCGenUnsafeSerializer) Marshal(o any) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	buf := make([]byte, a.Size())
	a.Marshal(buf)
	return buf, nil
}

func (s *BENCGenUnsafeSerializer) Unmarshal(bs []byte, o any) error {
	a := &s.a
	if err := a.Unmarshal(bs); err != nil {
		return err
	}
	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return nil
}

func NewBENCGenUnsafeSerializer() goserbench.Serializer {
	return &BENCGenUnsafeSerializer{}
}
