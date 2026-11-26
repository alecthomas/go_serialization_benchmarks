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

type IkeaCodec struct {
	buf *bytes.Buffer
}

func (s *IkeaCodec) Marshal(a *IkeA) (buf []byte, err error) {
	err = ikea.Pack(s.buf, a)
	if err != nil {
		return
	}
	buf = s.buf.Bytes()
	s.buf.Reset()
	return
}

func (s *IkeaCodec) Unmarshal(bs []byte, a *IkeA) error {
	s.buf.Reset()
	s.buf.Write(bs)
	return ikea.Unpack(s.buf, a)
}

func (s *IkeaCodec) FromSmall(v *goserbench.SmallStruct, a *IkeA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
}

func (s *IkeaCodec) ToSmall(a *IkeA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
}

func NewIkeaCodec() goserbench.Codec[IkeA] {
	return &IkeaCodec{buf: bytes.NewBuffer(make([]byte, 0, 1024))}
}
