package mus

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/unsafe"
)

type MUSSerializer struct{}

func (s MUSSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	n := ord.SizeString(v.Name, nil)
	n += unsafe.SizeInt64(v.BirthDay.UnixNano())
	n += ord.SizeString(v.Phone, nil)
	n += unsafe.SizeInt32(int32(v.Siblings))
	n += unsafe.SizeBool(v.Spouse)
	n += unsafe.SizeFloat64(v.Money)
	buf := make([]byte, n)
	n = ord.MarshalString(v.Name, nil, buf)
	n += unsafe.MarshalInt64(v.BirthDay.UnixNano(), buf[n:])
	n += ord.MarshalString(v.Phone, nil, buf[n:])
	n += unsafe.MarshalInt32(int32(v.Siblings), buf[n:])
	n += unsafe.MarshalBool(v.Spouse, buf[n:])
	raw.MarshalFloat64(v.Money, buf[n:])
	return buf, nil
}

func (s MUSSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	v := o.(*goserbench.SmallStruct)
	var n int
	v.Name, n, err = ord.UnmarshalString(nil, bs)
	if err != nil {
		return
	}
	var (
		n1         int
		birthDay64 int64
	)
	birthDay64, n1, err = unsafe.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.BirthDay = time.Unix(0, birthDay64)
	v.Phone, n1, err = ord.UnmarshalString(nil, bs[n:])
	n += n1
	if err != nil {
		return
	}
	var siblings32 int32
	siblings32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Siblings = int(siblings32)
	v.Spouse, n1, err = unsafe.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	return
}

func NewMUSSerializer() goserbench.Serializer {
	return MUSSerializer{}
}

type MUSUnsafeReuseSerializer struct {
	buf []byte
}

func (s MUSUnsafeReuseSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	n := unsafe.MarshalString(v.Name, nil, s.buf)
	n += unsafe.MarshalInt64(v.BirthDay.UnixNano(), s.buf[n:])
	n += unsafe.MarshalString(v.Phone, nil, s.buf[n:])
	n += unsafe.MarshalInt32(int32(v.Siblings), s.buf[n:])
	n += unsafe.MarshalBool(v.Spouse, s.buf[n:])
	n += unsafe.MarshalFloat64(v.Money, s.buf[n:])
	return s.buf[:n], nil
}

func (s MUSUnsafeReuseSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	v := o.(*goserbench.SmallStruct)
	var n int
	v.Name, n, err = unsafe.UnmarshalString(nil, bs)
	if err != nil {
		return
	}
	var (
		n1         int
		birthDay64 int64
	)
	birthDay64, n1, err = unsafe.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.BirthDay = time.Unix(0, birthDay64)
	v.Phone, n1, err = unsafe.UnmarshalString(nil, bs[n:])
	n += n1
	if err != nil {
		return
	}
	var siblings32 int32
	siblings32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Siblings = int(siblings32)
	v.Spouse, n1, err = unsafe.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	return
}

func NewMUSUnsafeReuseSerializer() goserbench.Serializer {
	return MUSUnsafeReuseSerializer{buf: make([]byte, 1024)}
}
