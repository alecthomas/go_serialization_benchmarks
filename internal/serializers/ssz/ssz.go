package ssz

import (
	"math"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	ssz "github.com/prysmaticlabs/go-ssz"
)

type SSZA struct {
	Name     string
	BirthDay uint64 // ssz does not support int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64 // ssz does not support float64
}

type SSZSerializer struct {
	a SSZA
}

func (s *SSZSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = uint64(v.BirthDay.UnixNano())
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
	return ssz.Marshal(a)
}

func (s *SSZSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	err = ssz.Unmarshal(bs, a)
	if err != nil {
		return
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, int64(a.BirthDay))
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
	return
}

func NewSSZSerializer() goserbench.Serializer {
	return &SSZSerializer{}
}
