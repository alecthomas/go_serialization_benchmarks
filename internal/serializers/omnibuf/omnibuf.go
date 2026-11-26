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

type OmniBufCodegen struct {
	a        OmniBufGo
	buf      []byte
	reuse    bool
	zeroCopy bool
}

func (s *OmniBufCodegen) Marshal(o any) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	if s.reuse {
		s.buf = a.MarshalOmniBufAppend(s.buf[:0])
		return s.buf, nil
	}
	return a.MarshalOmniBufAppend(nil), nil
}

func (s *OmniBufCodegen) Unmarshal(d []byte, o any) error {
	a := &s.a
	if err := a.UnmarshalOmniBuf(d, s.zeroCopy); err != nil {
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

func NewOmniBufSerializer() goserbench.Serializer {
	return &OmniBufCodegen{}
}

func NewOmniBufUnsafeReuseSerializer() goserbench.Serializer {
	return &OmniBufCodegen{
		buf:      make([]byte, 0, 128),
		reuse:    true,
		zeroCopy: true,
	}
}

type OmniBufReflect struct {
	a OmniBufGo
}

func (s *OmniBufReflect) Marshal(o any) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return omnibufwire.Marshal(a)
}

func (s *OmniBufReflect) Unmarshal(d []byte, o any) error {
	a := &s.a
	*a = OmniBufGo{}
	if err := omnibufwire.Unmarshal(d, a); err != nil {
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

func NewOmniBufReflectSerializer() goserbench.Serializer {
	return &OmniBufReflect{}
}
