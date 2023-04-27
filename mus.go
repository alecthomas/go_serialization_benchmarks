package goserbench

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
)

func MarshalMUS(v MUSA) (buf []byte) {
	n := ord.SizeString(v.Name)
	n += raw.SizeInt64(v.BirthDay)
	n += ord.SizeString(v.Phone)
	n += varint.SizeInt32(v.Siblings)
	n += ord.SizeBool(v.Spouse)
	n += raw.SizeFloat64(v.Money)
	buf = make([]byte, n)
	n = ord.MarshalString(v.Name, buf)
	n += raw.MarshalInt64(v.BirthDay, buf[n:])
	n += ord.MarshalString(v.Phone, buf[n:])
	n += varint.MarshalInt32(v.Siblings, buf[n:])
	n += ord.MarshalBool(v.Spouse, buf[n:])
	raw.MarshalFloat64(v.Money, buf[n:])
	return
}

func UnmarshalMUS(bs []byte) (v MUSA, n int, err error) {
	v.Name, n, err = ord.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	v.BirthDay, n1, err = raw.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = ord.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Siblings, n1, err = varint.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Spouse, n1, err = ord.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = raw.UnmarshalFloat64(bs[n:])
	n += n1
	return
}

func MarshalMUSUnsafe(v MUSA) (buf []byte) {
	n := unsafe.SizeString(v.Name)
	n += unsafe.SizeInt64(v.BirthDay)
	n += unsafe.SizeString(v.Phone)
	n += unsafe.SizeInt32(v.Siblings)
	n += unsafe.SizeBool(v.Spouse)
	n += unsafe.SizeFloat64(v.Money)
	buf = make([]byte, n)
	n = unsafe.MarshalString(v.Name, buf)
	n += unsafe.MarshalInt64(v.BirthDay, buf[n:])
	n += unsafe.MarshalString(v.Phone, buf[n:])
	n += unsafe.MarshalInt32(v.Siblings, buf[n:])
	n += unsafe.MarshalBool(v.Spouse, buf[n:])
	unsafe.MarshalFloat64(v.Money, buf[n:])
	return
}

func UnmarshalMUSUnsafe(bs []byte) (v MUSA, n int, err error) {
	v.Name, n, err = unsafe.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	v.BirthDay, n1, err = unsafe.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = unsafe.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Siblings, n1, err = unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Spouse, n1, err = unsafe.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	return
}
