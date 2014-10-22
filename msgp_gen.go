package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"bytes"
	"github.com/philhofer/msgp/enc"
	"io"
)

// MarshalMsg marshals a A into MessagePack
func (z *A) MarshalMsg() ([]byte, error) {
	var buf bytes.Buffer
	_, err := z.EncodeMsg(&buf)
	return buf.Bytes(), err
}

// UnmarshalMsg unmarshals a A from MessagePack, returning any extra bytes
// and any errors encountered
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = enc.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = enc.ReadStringZC(bts)
		if err != nil {
			return
		}
		switch enc.UnsafeString(field) {

		case "Name":

			z.Name, bts, err = enc.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "BirthDay":

			z.BirthDay, bts, err = enc.ReadTimeBytes(bts)

			if err != nil {
				return
			}

		case "Phone":

			z.Phone, bts, err = enc.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "Siblings":

			z.Siblings, bts, err = enc.ReadIntBytes(bts)

			if err != nil {
				return
			}

		case "Spouse":

			z.Spouse, bts, err = enc.ReadBoolBytes(bts)

			if err != nil {
				return
			}

		case "Money":

			z.Money, bts, err = enc.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}

		default:
			bts, err = enc.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// EncodeMsg encodes a A as MessagePack to the supplied io.Writer,
// returning the number of bytes written and any errors encountered
func (z *A) EncodeMsg(w io.Writer) (n int, err error) {
	en := enc.NewEncoder(w)
	return z.EncodeTo(en)
}

// EncodeTo encodes a A as MessagePack using the provided encoder,
// returning the number of bytes written and any errors encountered
func (z *A) EncodeTo(en *enc.MsgWriter) (n int, err error) {
	var nn int
	_ = nn

	nn, err = en.WriteMapHeader(6)
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Name")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString(z.Name)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("BirthDay")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteTime(z.BirthDay)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Phone")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString(z.Phone)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Siblings")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteInt(z.Siblings)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Spouse")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteBool(z.Spouse)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Money")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteFloat64(z.Money)

	n += nn
	if err != nil {
		return
	}

	return
}

// DecodeMsg decodes MessagePack from the provided io.Reader into the A,
// returning the number of bytes read and any errors encountered
func (z *A) DecodeMsg(r io.Reader) (n int, err error) {
	dc := enc.NewDecoder(r)
	n, err = z.DecodeFrom(dc)
	enc.Done(dc)
	return
}

// DecodeFrom deocdes MessagePack from the provided decoder into the A,
// returning the number of bytes read and any errors encountered.
func (z *A) DecodeFrom(dc *enc.MsgReader) (n int, err error) {
	var nn int
	var field []byte
	_ = nn
	_ = field

	var isz uint32
	isz, nn, err = dc.ReadMapHeader()
	n += nn
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, nn, err = dc.ReadStringAsBytes(field)
		n += nn
		if err != nil {
			return
		}
		switch enc.UnsafeString(field) {

		case "Name":

			z.Name, nn, err = dc.ReadString()

			n += nn
			if err != nil {
				return
			}

		case "BirthDay":

			z.BirthDay, nn, err = dc.ReadTime()

			n += nn
			if err != nil {
				return
			}

		case "Phone":

			z.Phone, nn, err = dc.ReadString()

			n += nn
			if err != nil {
				return
			}

		case "Siblings":

			z.Siblings, nn, err = dc.ReadInt()

			n += nn
			if err != nil {
				return
			}

		case "Spouse":

			z.Spouse, nn, err = dc.ReadBool()

			n += nn
			if err != nil {
				return
			}

		case "Money":

			z.Money, nn, err = dc.ReadFloat64()

			n += nn
			if err != nil {
				return
			}

		default:
			nn, err = dc.Skip()
			n += nn
			if err != nil {
				return
			}
		}
	}

	return
}
