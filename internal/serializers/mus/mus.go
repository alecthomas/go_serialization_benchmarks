package mus

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
)

type MUSSerializer struct{}

func (s MUSSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	n := ord.SizeString(v.Name)
	n += raw.SizeInt64(v.BirthDay.UnixNano())
	n += ord.SizeString(v.Phone)
	n += varint.SizeInt32(int32(v.Siblings))
	n += ord.SizeBool(v.Spouse)
	n += raw.SizeFloat64(v.Money)
	buf := make([]byte, n)
	n = ord.MarshalString(v.Name, buf)
	n += raw.MarshalInt64(v.BirthDay.UnixNano(), buf[n:])
	n += ord.MarshalString(v.Phone, buf[n:])
	n += varint.MarshalInt32(int32(v.Siblings), buf[n:])
	n += ord.MarshalBool(v.Spouse, buf[n:])
	raw.MarshalFloat64(v.Money, buf[n:])
	return buf, nil
}

func (s MUSSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	v := o.(*goserbench.SmallStruct)

	var n int

	v.Name, n, err = ord.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	var bdayNano int64
	bdayNano, n1, err = raw.UnmarshalInt64(bs[n:])
	v.BirthDay = time.Unix(0, bdayNano)
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = ord.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	var sibInt32 int32
	sibInt32, n1, err = varint.UnmarshalInt32(bs[n:])
	v.Siblings = int(sibInt32)
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

func NewMUSSerializer() goserbench.Serializer {
	return MUSSerializer{}
}

type MUSUnsafeSerializer struct{}

func (s MUSUnsafeSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	n := unsafe.SizeString(v.Name)
	n += unsafe.SizeInt64(v.BirthDay.UnixNano())
	n += unsafe.SizeString(v.Phone)
	n += unsafe.SizeInt32(int32(v.Siblings))
	n += unsafe.SizeBool(v.Spouse)
	n += unsafe.SizeFloat64(v.Money)
	buf := make([]byte, n)
	n = unsafe.MarshalString(v.Name, buf)
	n += unsafe.MarshalInt64(v.BirthDay.UnixNano(), buf[n:])
	n += unsafe.MarshalString(v.Phone, buf[n:])
	n += unsafe.MarshalInt32(int32(v.Siblings), buf[n:])
	n += unsafe.MarshalBool(v.Spouse, buf[n:])
	unsafe.MarshalFloat64(v.Money, buf[n:])
	return buf, nil
}

func (s MUSUnsafeSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	var n int
	v := o.(*goserbench.SmallStruct)

	v.Name, n, err = unsafe.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	var bdayNano int64
	bdayNano, n1, err = unsafe.UnmarshalInt64(bs[n:])
	v.BirthDay = time.Unix(0, bdayNano)
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = unsafe.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	var sibInt32 int32
	sibInt32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	v.Siblings = int(sibInt32)
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

func NewMUSUnsafeSerializer() goserbench.Serializer {
	return MUSUnsafeSerializer{}
}
