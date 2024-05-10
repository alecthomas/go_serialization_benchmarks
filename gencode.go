package goserbench

import "time"

type GencodeSerializer struct {
	buf []byte
	a   GencodeA
}

func (s *GencodeSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return a.Marshal(s.buf)
}

func (s *GencodeSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	_, err = a.Unmarshal(bs)
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

func newGencodeSerializer() *GencodeSerializer {
	return &GencodeSerializer{buf: make([]byte, 0, 1024)}
}

type GencodeUnsafeSerializer struct {
	buf []byte
	a   GencodeUnsafeA
}

func (s *GencodeUnsafeSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return a.Marshal(s.buf)
}

func (s *GencodeUnsafeSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	_, err = a.Unmarshal(bs)
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

func newGencodeUnsafeSerializer() *GencodeUnsafeSerializer {
	return &GencodeUnsafeSerializer{buf: make([]byte, 0, 1024)}
}
