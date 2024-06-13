package mus

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
)

type MUSSerializer struct {
	buf []byte
}

func (s MUSSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	n := ord.MarshalString(v.Name, s.buf)
	n += raw.MarshalInt64(v.BirthDay.UnixNano(), s.buf[n:])
	n += ord.MarshalString(v.Phone, s.buf[n:])
	n += varint.MarshalInt32(int32(v.Siblings), s.buf[n:])
	n += ord.MarshalBool(v.Spouse, s.buf[n:])
	n += raw.MarshalFloat64(v.Money, s.buf[n:])
	return s.buf[:n], nil
}

func (s MUSSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	var (
		n int
		v = o.(*goserbench.SmallStruct)
	)
	v.Name, n, err = ord.UnmarshalString(bs)
	if err != nil {
		return
	}
	var (
		n1       int
		bdayNano int64
	)
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
	return MUSSerializer{buf: make([]byte, 1024)}
}

type MUSUnsafeSerializer struct {
	buf []byte
}

func (s MUSUnsafeSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	n := unsafe.MarshalString(v.Name, s.buf)
	n += unsafe.MarshalInt64(v.BirthDay.UnixNano(), s.buf[n:])
	n += unsafe.MarshalString(v.Phone, s.buf[n:])
	n += unsafe.MarshalInt32(int32(v.Siblings), s.buf[n:])
	n += unsafe.MarshalBool(v.Spouse, s.buf[n:])
	n += unsafe.MarshalFloat64(v.Money, s.buf[n:])
	return s.buf[:n], nil
}

func (s MUSUnsafeSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	var (
		n int
		v = o.(*goserbench.SmallStruct)
	)
	v.Name, n, err = unsafe.UnmarshalString(bs)
	if err != nil {
		return
	}
	var (
		n1       int
		bdayNano int64
	)
	bdayNano, n1, err = unsafe.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.BirthDay = time.Unix(0, bdayNano)
	v.Phone, n1, err = unsafe.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	var sibInt32 int32
	sibInt32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Siblings = int(sibInt32)
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
	return MUSUnsafeSerializer{buf: make([]byte, 1024)}
}
