package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"github.com/philhofer/msgp/msgp"
)


// MarshalMsg marshals a A into MessagePack
func (z *A) MarshalMsg() ([]byte, error) {
	o := make([]byte, 0, z.Maxsize())
	return z.AppendMsg(o)
}

// AppendMsg marshals a A onto the end of a []byte
func (z *A) AppendMsg(b []byte) (o []byte, err error) {
	o = b

	o = msgp.AppendMapHeader(o, 6)

	o = msgp.AppendString(o, "Name")

	o = msgp.AppendString(o, z.Name)

	o = msgp.AppendString(o, "BirthDay")

	o = msgp.AppendTime(o, z.BirthDay)

	o = msgp.AppendString(o, "Phone")

	o = msgp.AppendString(o, z.Phone)

	o = msgp.AppendString(o, "Siblings")

	o = msgp.AppendInt(o, z.Siblings)

	o = msgp.AppendString(o, "Spouse")

	o = msgp.AppendBool(o, z.Spouse)

	o = msgp.AppendString(o, "Money")

	o = msgp.AppendFloat64(o, z.Money)

	return
}
// UnmarshalMsg unmarshals a A from MessagePack, returning any extra bytes
// and any errors encountered
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadStringZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Name":

			z.Name, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "BirthDay":

			z.BirthDay, bts, err = msgp.ReadTimeBytes(bts)

			if err != nil {
				return
			}

		case "Phone":

			z.Phone, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "Siblings":

			z.Siblings, bts, err = msgp.ReadIntBytes(bts)

			if err != nil {
				return
			}

		case "Spouse":

			z.Spouse, bts, err = msgp.ReadBoolBytes(bts)

			if err != nil {
				return
			}

		case "Money":

			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}

		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// Maxsize returns the encoded size of the object when messagepack encoded.
// This value is guaranteed to be larger than the encoded size
// of the A unless it contains any non-concrete
// (e.g. "interface{}") fields.
func (z *A) Maxsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 4

	s += msgp.StringPrefixSize + len(z.Name)
	s += msgp.StringPrefixSize + 8

	s += msgp.TimeSize
	s += msgp.StringPrefixSize + 5

	s += msgp.StringPrefixSize + len(z.Phone)
	s += msgp.StringPrefixSize + 8

	s += msgp.IntSize
	s += msgp.StringPrefixSize + 6

	s += msgp.BoolSize
	s += msgp.StringPrefixSize + 5

	s += msgp.Float64Size

	return
}
