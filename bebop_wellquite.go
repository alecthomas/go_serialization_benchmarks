package goserbench

import "time"

type BebopWellquiteSerializer struct {
	a   BebopBufWellquite
	buf []byte
}

func (s *BebopWellquiteSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return a.MarshalBebop(s.buf)
}

func (s *BebopWellquiteSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	_, err = a.UnmarshalBebop(bs)
	if err != nil {
		return
	}

	v := o.(*A)
	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}

func (s *BebopWellquiteSerializer) TimePrecision() time.Duration {
	return 100 * time.Nanosecond
}

func NewBebopWellquiteSerializer() Serializer {
	return &BebopWellquiteSerializer{
		buf: make([]byte, 1024),
	}
}
