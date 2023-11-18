package goserbench

import (
	bstd "github.com/deneonet/benc"
	"github.com/deneonet/benc/bunsafe"
)

func MarshalBENC(v BENC) (buf []byte, err error) {
	n := bstd.SizeString(v.Name)
	n += bstd.SizeInt64()
	n += bstd.SizeString(v.Phone)
	n += bstd.SizeInt32()
	n += bstd.SizeBool()
	n += bstd.SizeFloat64()
	n, buf = bstd.Marshal(n)
	n = bstd.MarshalString(n, buf, v.Name)
	n = bstd.MarshalInt64(n, buf, v.BirthDay)
	n = bstd.MarshalString(n, buf, v.Phone)
	n = bstd.MarshalInt32(n, buf, v.Siblings)
	n = bstd.MarshalBool(n, buf, v.Spouse)
	n = bstd.MarshalFloat64(n, buf, v.Money)
	err = bstd.VerifyMarshal(n, buf)
	return
}

func UnmarshalBENC(bs []byte) (v BENC, n int, err error) {
	n, v.Name, err = bstd.UnmarshalString(0, bs)
	if err != nil {
		return
	}
	n, v.BirthDay, err = bstd.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	n, v.Phone, err = bstd.UnmarshalString(n, bs)
	if err != nil {
		return
	}
	n, v.Siblings, err = bstd.UnmarshalInt32(n, bs)
	if err != nil {
		return
	}
	n, v.Spouse, err = bstd.UnmarshalBool(n, bs)
	if err != nil {
		return
	}
	n, v.Money, err = bstd.UnmarshalFloat64(n, bs)
	err = bstd.VerifyUnmarshal(n, bs)
	return
}

func MarshalBENC_UnsafeStringConversion(v BENC) (buf []byte, err error) {
	n := bstd.SizeString(v.Name)
	n += bstd.SizeInt64()
	n += bstd.SizeString(v.Phone)
	n += bstd.SizeInt32()
	n += bstd.SizeBool()
	n += bstd.SizeFloat64()
	n, buf = bstd.Marshal(n)
	n = bunsafe.MarshalString(n, buf, v.Name)
	n = bstd.MarshalInt64(n, buf, v.BirthDay)
	n = bunsafe.MarshalString(n, buf, v.Phone)
	n = bstd.MarshalInt32(n, buf, v.Siblings)
	n = bstd.MarshalBool(n, buf, v.Spouse)
	n = bstd.MarshalFloat64(n, buf, v.Money)
	err = bstd.VerifyMarshal(n, buf)
	return
}

func UnmarshalBENC_UnsafeStringConversion(bs []byte) (v BENC, n int, err error) {
	n, v.Name, err = bunsafe.UnmarshalString(0, bs)
	if err != nil {
		return
	}
	n, v.BirthDay, err = bstd.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	n, v.Phone, err = bunsafe.UnmarshalString(n, bs)
	if err != nil {
		return
	}
	n, v.Siblings, err = bstd.UnmarshalInt32(n, bs)
	if err != nil {
		return
	}
	n, v.Spouse, err = bstd.UnmarshalBool(n, bs)
	if err != nil {
		return
	}
	n, v.Money, err = bstd.UnmarshalFloat64(n, bs)
	err = bstd.VerifyUnmarshal(n, bs)
	return
}
