package goserbench

import (
	"bytes"
	"time"

	capn "github.com/glycerine/go-capnproto"
)

type CapNProtoSerializer struct{}

func (x CapNProtoSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	seg := capn.NewBuffer(nil)
	c := NewRootCapnpA(seg)
	c.SetName(a.Name)
	c.SetBirthDay(a.BirthDay.UnixNano())
	c.SetPhone(a.Phone)
	c.SetSiblings(int32(a.Siblings))
	c.SetSpouse(a.Spouse)
	c.SetMoney(a.Money)
	var buf bytes.Buffer
	_, err := seg.WriteTo(&buf)
	return buf.Bytes(), err
}

func (x CapNProtoSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	s, _, err := capn.ReadFromMemoryZeroCopy(d)
	if err != nil {
		return err
	}
	o := ReadRootCapnpA(s)
	a.Name = o.Name()
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone = o.Phone()
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func NewCapNProtoSerializer() CapNProtoSerializer {
	return CapNProtoSerializer{}
}
