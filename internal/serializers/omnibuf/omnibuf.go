package omnibuf

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	omnibufwire "github.com/omnibuf/omnibuf-go"
)

type OmniBufGo struct {
	Name     string  `omnibuf:"1,name=name"`
	BirthDay int64   `omnibuf:"2,name=birthDay"`
	Phone    string  `omnibuf:"3,name=phone"`
	Siblings int32   `omnibuf:"4,name=siblings"`
	Spouse   bool    `omnibuf:"5,name=spouse"`
	Money    float64 `omnibuf:"6,name=money"`
}

// omniBufConversions converts between OmniBufGo and goserbench.SmallStruct
// for both OmniBuf codecs.
type omniBufConversions struct{}

func (omniBufConversions) FromSmall(v *goserbench.SmallStruct, a *OmniBufGo) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (omniBufConversions) ToSmall(a *OmniBufGo, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

type OmniBufCodec struct {
	omniBufConversions
	buf      []byte
	reuse    bool
	zeroCopy bool
}

func (s *OmniBufCodec) Marshal(a *OmniBufGo) ([]byte, error) {
	if s.reuse {
		s.buf = a.MarshalOmniBufAppend(s.buf[:0])
		return s.buf, nil
	}
	return a.MarshalOmniBufAppend(nil), nil
}

func (s *OmniBufCodec) Unmarshal(d []byte, a *OmniBufGo) error {
	return a.UnmarshalOmniBuf(d, s.zeroCopy)
}

func NewOmniBufCodec() goserbench.Codec[OmniBufGo] {
	return &OmniBufCodec{}
}

func NewOmniBufUnsafeReuseCodec() goserbench.Codec[OmniBufGo] {
	return &OmniBufCodec{
		buf:      make([]byte, 0, 128),
		reuse:    true,
		zeroCopy: true,
	}
}

type OmniBufReflectCodec struct {
	omniBufConversions
}

func (s *OmniBufReflectCodec) Marshal(a *OmniBufGo) ([]byte, error) {
	return omnibufwire.Marshal(a)
}

func (s *OmniBufReflectCodec) Unmarshal(d []byte, a *OmniBufGo) error {
	return omnibufwire.Unmarshal(d, a)
}

func NewOmniBufReflectCodec() goserbench.Codec[OmniBufGo] {
	return &OmniBufReflectCodec{}
}
