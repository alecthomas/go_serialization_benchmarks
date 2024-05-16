package goserbench

import (
	"time"

	bstd "github.com/deneonet/benc"
	"github.com/deneonet/benc/bunsafe"
)

type BENCSerializer struct{}

func (s BENCSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	n := bstd.SizeString(v.Name)
	n += bstd.SizeInt64()
	n += bstd.SizeString(v.Phone)
	n += bstd.SizeInt32()
	n += bstd.SizeBool()
	n += bstd.SizeFloat64()
	n, buf = bstd.Marshal(n)
	n = bstd.MarshalString(n, buf, v.Name)
	n = bstd.MarshalInt64(n, buf, v.BirthDay.UnixNano())
	n = bstd.MarshalString(n, buf, v.Phone)
	n = bstd.MarshalInt32(n, buf, int32(v.Siblings))
	n = bstd.MarshalBool(n, buf, v.Spouse)
	n = bstd.MarshalFloat64(n, buf, v.Money)
	err = bstd.VerifyMarshal(n, buf)
	return
}

func (s BENCSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	v := o.(*A)
	var n int
	n, v.Name, err = bstd.UnmarshalString(0, bs)
	if err != nil {
		return
	}
	var bday int64
	n, bday, err = bstd.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	v.BirthDay = time.Unix(0, bday)
	n, v.Phone, err = bstd.UnmarshalString(n, bs)
	if err != nil {
		return
	}
	var sib int32
	n, sib, err = bstd.UnmarshalInt32(n, bs)
	if err != nil {
		return
	}
	v.Siblings = int(sib)
	n, v.Spouse, err = bstd.UnmarshalBool(n, bs)
	if err != nil {
		return
	}
	n, v.Money, err = bstd.UnmarshalFloat64(n, bs)
	err = bstd.VerifyUnmarshal(n, bs)
	return
}

func NewBENCSerializer() Serializer {
	return BENCSerializer{}
}

type BENCUnsafeSerializer struct{}

func (s BENCUnsafeSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	n := bstd.SizeString(v.Name)
	n += bstd.SizeInt64()
	n += bstd.SizeString(v.Phone)
	n += bstd.SizeInt32()
	n += bstd.SizeBool()
	n += bstd.SizeFloat64()
	n, buf = bstd.Marshal(n)
	n = bunsafe.MarshalString(n, buf, v.Name)
	n = bstd.MarshalInt64(n, buf, v.BirthDay.UnixNano())
	n = bunsafe.MarshalString(n, buf, v.Phone)
	n = bstd.MarshalInt32(n, buf, int32(v.Siblings))
	n = bstd.MarshalBool(n, buf, v.Spouse)
	n = bstd.MarshalFloat64(n, buf, v.Money)
	err = bstd.VerifyMarshal(n, buf)
	return
}

func (s BENCUnsafeSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	var n int
	v := o.(*A)
	n, v.Name, err = bunsafe.UnmarshalString(0, bs)
	if err != nil {
		return
	}
	var bday int64
	n, bday, err = bstd.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	v.BirthDay = time.Unix(0, bday)
	n, v.Phone, err = bunsafe.UnmarshalString(n, bs)
	if err != nil {
		return
	}
	var sib int32
	n, sib, err = bstd.UnmarshalInt32(n, bs)
	if err != nil {
		return
	}
	v.Siblings = int(sib)
	n, v.Spouse, err = bstd.UnmarshalBool(n, bs)
	if err != nil {
		return
	}
	n, v.Money, err = bstd.UnmarshalFloat64(n, bs)
	err = bstd.VerifyUnmarshal(n, bs)
	return
}

func NewBENCUnsafeSerializer() Serializer {
	return BENCUnsafeSerializer{}
}
