package goserbench

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	capnp "zombiezen.com/go/capnproto2"
)

type Capnp2A struct{ capnp.Struct }

func NewCapnp2A(s *capnp.Segment) (Capnp2A, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	if err != nil {
		return Capnp2A{}, err
	}
	return Capnp2A{st}, nil
}

func NewRootCapnp2A(s *capnp.Segment) (Capnp2A, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	if err != nil {
		return Capnp2A{}, err
	}
	return Capnp2A{st}, nil
}

func ReadRootCapnp2A(msg *capnp.Message) (Capnp2A, error) {
	root, err := msg.Root()
	if err != nil {
		return Capnp2A{}, err
	}
	st := capnp.ToStruct(root)
	return Capnp2A{st}, nil
}

func (s Capnp2A) Name() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Capnp2A) NameBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s Capnp2A) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Capnp2A) BirthDay() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Capnp2A) SetBirthDay(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s Capnp2A) Phone() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Capnp2A) PhoneBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s Capnp2A) SetPhone(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s Capnp2A) Siblings() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s Capnp2A) SetSiblings(v int32) {

	s.Struct.SetUint32(8, uint32(v))
}

func (s Capnp2A) Spouse() bool {
	return s.Struct.Bit(96)
}

func (s Capnp2A) SetSpouse(v bool) {

	s.Struct.SetBit(96, v)
}

func (s Capnp2A) Money() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s Capnp2A) SetMoney(v float64) {

	s.Struct.SetUint64(16, math.Float64bits(v))
}

// Capnp2A_List is a list of Capnp2A.
type Capnp2A_List struct{ capnp.List }

// NewCapnp2A creates a new list of Capnp2A.
func NewCapnp2A_List(s *capnp.Segment, sz int32) (Capnp2A_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	if err != nil {
		return Capnp2A_List{}, err
	}
	return Capnp2A_List{l}, nil
}

func (s Capnp2A_List) At(i int) Capnp2A           { return Capnp2A{s.List.Struct(i)} }
func (s Capnp2A_List) Set(i int, v Capnp2A) error { return s.List.SetStruct(i, v.Struct) }

// Capnp2A_Promise is a wrapper for a Capnp2A promised by a client call.
type Capnp2A_Promise struct{ *capnp.Pipeline }

func (p Capnp2A_Promise) Struct() (Capnp2A, error) {
	s, err := p.Pipeline.Struct()
	return Capnp2A{s}, err
}
