package capnproto

import (
	"bytes"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	capn "github.com/glycerine/go-capnproto"
)

type CapNProtoCodec struct{ goserbench.NativeSmallStruct }

func (x CapNProtoCodec) Marshal(a *goserbench.SmallStruct) ([]byte, error) {
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

func (x CapNProtoCodec) Unmarshal(d []byte, a *goserbench.SmallStruct) error {
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

func NewCapNProtoCodec() goserbench.Codec[goserbench.SmallStruct] {
	return CapNProtoCodec{}
}
