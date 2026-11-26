package benc

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	benclib "github.com/deneonet/benc"
	bstd "github.com/deneonet/benc/std"
)

type BENCCodec struct{ goserbench.NativeSmallStruct }

func (s BENCCodec) Marshal(v *goserbench.SmallStruct) (buf []byte, err error) {
	n := bstd.SizeString(v.Name)
	n += bstd.SizeInt64()
	n += bstd.SizeString(v.Phone)
	n += bstd.SizeInt32()
	n += bstd.SizeBool()
	n += bstd.SizeFloat64()
	buf = make([]byte, n)
	n = 0
	n = bstd.MarshalString(n, buf, v.Name)
	n = bstd.MarshalInt64(n, buf, v.BirthDay.UnixNano())
	n = bstd.MarshalString(n, buf, v.Phone)
	n = bstd.MarshalInt32(n, buf, int32(v.Siblings))
	n = bstd.MarshalBool(n, buf, v.Spouse)
	n = bstd.MarshalFloat64(n, buf, v.Money)
	err = benclib.VerifyMarshal(n, buf)
	return
}

func (s BENCCodec) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
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
	err = benclib.VerifyUnmarshal(n, bs)
	return
}

func NewBENCCodec() goserbench.Codec[goserbench.SmallStruct] {
	return BENCCodec{}
}

type BENCUnsafeCodec struct{ goserbench.NativeSmallStruct }

func (s BENCUnsafeCodec) Marshal(v *goserbench.SmallStruct) (buf []byte, err error) {
	n := bstd.SizeString(v.Name)
	n += bstd.SizeInt64()
	n += bstd.SizeString(v.Phone)
	n += bstd.SizeInt32()
	n += bstd.SizeBool()
	n += bstd.SizeFloat64()
	buf = make([]byte, n)
	n = 0
	n = bstd.MarshalUnsafeString(n, buf, v.Name)
	n = bstd.MarshalInt64(n, buf, v.BirthDay.UnixNano())
	n = bstd.MarshalUnsafeString(n, buf, v.Phone)
	n = bstd.MarshalInt32(n, buf, int32(v.Siblings))
	n = bstd.MarshalBool(n, buf, v.Spouse)
	n = bstd.MarshalFloat64(n, buf, v.Money)
	err = benclib.VerifyMarshal(n, buf)
	return
}

func (s BENCUnsafeCodec) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	var n int
	n, v.Name, err = bstd.UnmarshalUnsafeString(0, bs)
	if err != nil {
		return
	}
	var bday int64
	n, bday, err = bstd.UnmarshalInt64(n, bs)
	if err != nil {
		return
	}
	v.BirthDay = time.Unix(0, bday)
	n, v.Phone, err = bstd.UnmarshalUnsafeString(n, bs)
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
	err = benclib.VerifyUnmarshal(n, bs)
	return
}

func NewBENCUnsafeCodec() goserbench.Codec[goserbench.SmallStruct] {
	return BENCUnsafeCodec{}
}
