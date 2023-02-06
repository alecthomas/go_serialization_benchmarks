package goserbench

import (
	"bytes"
	"fmt"
	"time"

	"github.com/mojura/enkodo"
)

func newEnkodoSerializer() *enkodoSerializer {
	var m enkodoSerializer
	m.w = enkodo.NewWriter(bytes.NewBuffer(nil))
	return &m
}

type enkodoSerializer struct {
	w *enkodo.Writer
}

func (m *enkodoSerializer) Marshal(value interface{}) (bs []byte, err error) {
	encodee, ok := value.(enkodo.Encodee)
	if !ok {
		err = fmt.Errorf("invalid type, expected %t and received %t", encodee, value)
		return
	}

	if err = m.w.Encode(encodee); err != nil {
		return
	}

	bs = m.w.Bytes()
	m.w.Reset()
	return
}

func (m *enkodoSerializer) Unmarshal(bs []byte, value interface{}) (err error) {
	decodee, ok := value.(enkodo.Decodee)
	if !ok {
		err = fmt.Errorf("invalid type, expected %t and received %t", decodee, value)
		return
	}
	return enkodo.Unmarshal(bs, decodee)
}

// MarshalEnkodo is an encoding helper func
func (a *A) MarshalEnkodo(enc *enkodo.Encoder) (err error) {
	enc.String(a.Name)
	enc.Int64(a.BirthDay.UnixNano())
	enc.String(a.Phone)
	enc.Int(a.Siblings)
	enc.Bool(a.Spouse)
	enc.Float64(a.Money)
	return
}

// UnmarshalEnkodo is an decoding helper func
func (a *A) UnmarshalEnkodo(dec *enkodo.Decoder) (err error) {
	if a.Name, err = dec.String(); err != nil {
		return
	}

	var unixNano int64
	if unixNano, err = dec.Int64(); err != nil {
		return
	}

	a.BirthDay = time.Unix(0, unixNano)

	if a.Phone, err = dec.String(); err != nil {
		return
	}

	if a.Siblings, err = dec.Int(); err != nil {
		return
	}

	if a.Spouse, err = dec.Bool(); err != nil {
		return
	}

	if a.Money, err = dec.Float64(); err != nil {
		return
	}

	return
}
