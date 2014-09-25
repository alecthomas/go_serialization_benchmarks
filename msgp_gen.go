package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"github.com/philhofer/msgp/enc"
	"io"
	"bytes"
)

func (z *A) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	_, err := z.WriteTo(&buf)
	return buf.Bytes(), err
}

func (z *A) Unmarshal(b []byte) error {
	_, err := z.ReadFrom(bytes.NewReader(b))
	return err
}

func (z *A) WriteTo(w io.Writer) (n int64, err error) {
	var nn int
	en := enc.NewEncoder(w)
	_ = nn
	_ = en

	if z == nil {
		nn, err = en.WriteNil()
		n += int64(nn)
		if err != nil {
			return
		}
	} else {

		nn, err = en.WriteMapHeader(uint32(6))
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString("Name")
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString(z.Name)

		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString("BirthDay")
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteTime(z.BirthDay)

		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString("Phone")
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString(z.Phone)

		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString("Siblings")
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteInt(z.Siblings)

		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString("Spouse")
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteBool(z.Spouse)

		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteString("Money")
		n += int64(nn)
		if err != nil {
			return
		}

		nn, err = en.WriteFloat64(z.Money)

		n += int64(nn)
		if err != nil {
			return
		}

	}

	return
}

func (z *A) ReadFrom(r io.Reader) (n int64, err error) {
	var sz uint32
	var nn int
	var field []byte
	dc := enc.NewDecoder(r)
	_ = sz
	_ = nn
	_ = field

	if dc.IsNil() {
		nn, err = dc.ReadNil()
		n += int64(nn)
		if err != nil {
			return
		}
		z = nil
	} else {
		if z == nil {
			z = new(A)
		}

		var isz uint32
		isz, nn, err = dc.ReadMapHeader()
		n += int64(nn)
		if err != nil {
			return
		}
		for xplz := uint32(0); xplz < isz; xplz++ {
			field, nn, err = dc.ReadStringAsBytes(field)
			n += int64(nn)
			if err != nil {
				return
			}
			switch enc.UnsafeString(field) {

			case "Name":

				z.Name, nn, err = dc.ReadString()

				n += int64(nn)
				if err != nil {
					return
				}

			case "BirthDay":

				z.BirthDay, nn, err = dc.ReadTime()

				n += int64(nn)
				if err != nil {
					return
				}

			case "Phone":

				z.Phone, nn, err = dc.ReadString()

				n += int64(nn)
				if err != nil {
					return
				}

			case "Siblings":

				z.Siblings, nn, err = dc.ReadInt()

				n += int64(nn)
				if err != nil {
					return
				}

			case "Spouse":

				z.Spouse, nn, err = dc.ReadBool()

				n += int64(nn)
				if err != nil {
					return
				}

			case "Money":

				z.Money, nn, err = dc.ReadFloat64()

				n += int64(nn)
				if err != nil {
					return
				}

			default:
				nn, err = dc.Skip()
				n += int64(nn)
				if err != nil {
					return
				}
			}
		}

	}

	enc.Done(dc)
	return
}

