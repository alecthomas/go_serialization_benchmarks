package capnproto

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"math"

	C "github.com/glycerine/go-capnproto"
)

type CapnpA C.Struct

func NewCapnpA(s *C.Segment) CapnpA      { return CapnpA(s.NewStruct(24, 2)) }
func NewRootCapnpA(s *C.Segment) CapnpA  { return CapnpA(s.NewRootStruct(24, 2)) }
func AutoNewCapnpA(s *C.Segment) CapnpA  { return CapnpA(s.NewStructAR(24, 2)) }
func ReadRootCapnpA(s *C.Segment) CapnpA { return CapnpA(s.Root(0).ToStruct()) }
func (s CapnpA) Name() string            { return C.Struct(s).GetObject(0).ToText() }
func (s CapnpA) NameBytes() []byte       { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s CapnpA) SetName(v string)        { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s CapnpA) BirthDay() int64         { return int64(C.Struct(s).Get64(0)) }
func (s CapnpA) SetBirthDay(v int64)     { C.Struct(s).Set64(0, uint64(v)) }
func (s CapnpA) Phone() string           { return C.Struct(s).GetObject(1).ToText() }
func (s CapnpA) PhoneBytes() []byte      { return C.Struct(s).GetObject(1).ToDataTrimLastByte() }
func (s CapnpA) SetPhone(v string)       { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s CapnpA) Siblings() int32         { return int32(C.Struct(s).Get32(8)) }
func (s CapnpA) SetSiblings(v int32)     { C.Struct(s).Set32(8, uint32(v)) }
func (s CapnpA) Spouse() bool            { return C.Struct(s).Get1(96) }
func (s CapnpA) SetSpouse(v bool)        { C.Struct(s).Set1(96, v) }
func (s CapnpA) Money() float64          { return math.Float64frombits(C.Struct(s).Get64(16)) }
func (s CapnpA) SetMoney(v float64)      { C.Struct(s).Set64(16, math.Float64bits(v)) }
func (s CapnpA) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"name\":")
	if err != nil {
		return err
	}
	{
		s := s.Name()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"birthDay\":")
	if err != nil {
		return err
	}
	{
		s := s.BirthDay()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"phone\":")
	if err != nil {
		return err
	}
	{
		s := s.Phone()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"siblings\":")
	if err != nil {
		return err
	}
	{
		s := s.Siblings()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"spouse\":")
	if err != nil {
		return err
	}
	{
		s := s.Spouse()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"money\":")
	if err != nil {
		return err
	}
	{
		s := s.Money()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s CapnpA) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s CapnpA) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("name = ")
	if err != nil {
		return err
	}
	{
		s := s.Name()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("birthDay = ")
	if err != nil {
		return err
	}
	{
		s := s.BirthDay()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("phone = ")
	if err != nil {
		return err
	}
	{
		s := s.Phone()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("siblings = ")
	if err != nil {
		return err
	}
	{
		s := s.Siblings()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("spouse = ")
	if err != nil {
		return err
	}
	{
		s := s.Spouse()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("money = ")
	if err != nil {
		return err
	}
	{
		s := s.Money()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s CapnpA) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type CapnpA_List C.PointerList

func NewCapnpAList(s *C.Segment, sz int) CapnpA_List {
	return CapnpA_List(s.NewCompositeList(24, 2, sz))
}
func (s CapnpA_List) Len() int        { return C.PointerList(s).Len() }
func (s CapnpA_List) At(i int) CapnpA { return CapnpA(C.PointerList(s).At(i).ToStruct()) }
func (s CapnpA_List) ToArray() []CapnpA {
	n := s.Len()
	a := make([]CapnpA, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s CapnpA_List) Set(i int, item CapnpA) { C.PointerList(s).Set(i, C.Object(item)) }
