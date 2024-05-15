package goserbench

import (
	"bytes"
	"math"
	"time"

	ikea "github.com/ikkerens/ikeapack"
)

type IkeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64 // NOTE: Ike does not support float64 - it needs to be converted to an int type.
}

type IkeSerializer struct {
	a   IkeA
	buf *bytes.Buffer
}

func (s *IkeSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
	err = ikea.Pack(s.buf, a)
	if err != nil {
		return
	}
	buf = s.buf.Bytes()
	s.buf.Reset()
	return
}

func (s *IkeSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	s.buf.Reset()
	s.buf.Write(bs)
	err = ikea.Unpack(s.buf, a)
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

func newIkeSerializer() *IkeSerializer {
	return &IkeSerializer{buf: bytes.NewBuffer(make([]byte, 0, 1024))}
}
