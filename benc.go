package goserbench

import (
	"github.com/deneonet/benc/benc"
)

func MarshalBENC(v BENC) (buf []byte, err error) {
	n := benc.SizeString(v.Name)
	n += benc.SizeInt64()
	n += benc.SizeString(v.Phone)
	n += benc.SizeInt32()
	n += benc.SizeBool()
	n += benc.SizeFloat64()
	n, buf = benc.Marshal(n)
	n, buf = benc.MarshalString(n, buf, v.Name)
	n, buf = benc.MarshalInt64(n, buf, v.BirthDay)
	n, buf = benc.MarshalString(n, buf, v.Phone)
	n, buf = benc.MarshalInt32(n, buf, v.Siblings)
	n, buf = benc.MarshalBool(n, buf, v.Spouse)
	n, buf = benc.MarshalFloat64(n, buf, v.Money)
	err = benc.VerifyMarshal(n, buf)
	return
}

func UnmarshalBENC(bs []byte) (v BENC, n int, err error) {
	n, v.Name, err = benc.UnmarshalString(0, bs)
	if err != nil {
		return
	}
	n, v.BirthDay, err = benc.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	n, v.Phone, err = benc.UnmarshalString(n, bs)
	if err != nil {
		return
	}
	n, v.Siblings, err = benc.UnmarshalInt32(n, bs)
	if err != nil {
		return
	}
	n, v.Spouse, err = benc.UnmarshalBool(n, bs)
	if err != nil {
		return
	}
	n, v.Money, err = benc.UnmarshalFloat64(n, bs)
	err = benc.VerifyUnmarshal(n, bs)
	return
}
