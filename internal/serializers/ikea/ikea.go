package ikea

import (
	"bytes"
	"math"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	ikea "github.com/ikkerens/ikeapack"
)

type IkeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64
}

type IkeaSerializer struct {
	a   IkeA
	buf *bytes.Buffer
}

func (s *IkeaSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*goserbench.SmallStruct)
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

func (s *IkeaSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	s.buf.Reset()
	s.buf.Write(bs)
	err = ikea.Unpack(s.buf, a)
	if err != nil {
		return
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
	return
}

func NewIkeaSerializer() goserbench.Serializer {
	return &IkeaSerializer{buf: bytes.NewBuffer(make([]byte, 0, 1024))}
}
