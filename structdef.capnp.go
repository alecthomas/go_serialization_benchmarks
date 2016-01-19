package goserbench

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	capnp "zombiezen.com/go/capnproto2"
)

type CapnpA struct{ capnp.Struct }

func NewCapnpA(s *capnp.Segment) (CapnpA, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	if err != nil {
		return CapnpA{}, err
	}
	return CapnpA{st}, nil
}

func NewRootCapnpA(s *capnp.Segment) (CapnpA, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	if err != nil {
		return CapnpA{}, err
	}
	return CapnpA{st}, nil
}

func ReadRootCapnpA(msg *capnp.Message) (CapnpA, error) {
	root, err := msg.Root()
	if err != nil {
		return CapnpA{}, err
	}
	st := capnp.ToStruct(root)
	return CapnpA{st}, nil
}

func (s CapnpA) Name() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s CapnpA) NameBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s CapnpA) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s CapnpA) BirthDay() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s CapnpA) SetBirthDay(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s CapnpA) Phone() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s CapnpA) PhoneBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s CapnpA) SetPhone(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s CapnpA) Siblings() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s CapnpA) SetSiblings(v int32) {

	s.Struct.SetUint32(8, uint32(v))
}

func (s CapnpA) Spouse() bool {
	return s.Struct.Bit(96)
}

func (s CapnpA) SetSpouse(v bool) {

	s.Struct.SetBit(96, v)
}

func (s CapnpA) Money() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s CapnpA) SetMoney(v float64) {

	s.Struct.SetUint64(16, math.Float64bits(v))
}

// CapnpA_List is a list of CapnpA.
type CapnpA_List struct{ capnp.List }

// NewCapnpA creates a new list of CapnpA.
func NewCapnpA_List(s *capnp.Segment, sz int32) (CapnpA_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	if err != nil {
		return CapnpA_List{}, err
	}
	return CapnpA_List{l}, nil
}

func (s CapnpA_List) At(i int) CapnpA           { return CapnpA{s.List.Struct(i)} }
func (s CapnpA_List) Set(i int, v CapnpA) error { return s.List.SetStruct(i, v.Struct) }

// CapnpA_Promise is a wrapper for a CapnpA promised by a client call.
type CapnpA_Promise struct{ *capnp.Pipeline }

func (p CapnpA_Promise) Struct() (CapnpA, error) {
	s, err := p.Pipeline.Struct()
	return CapnpA{s}, err
}
