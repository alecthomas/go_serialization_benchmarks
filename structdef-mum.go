package goserbench

import (
	"fmt"
	"time"

	"github.com/itsmontoya/mum"
)

func newMumSerializer() *mumSerializer {
	var m mumSerializer
	m.w = mum.NewWriter(make([]byte, 0, 67))
	m.r = mum.NewReader(nil)
	return &m
}

type mumSerializer struct {
	w *mum.Writer
	r *mum.Reader
}

func (m *mumSerializer) Marshal(value interface{}) (bs []byte, err error) {
	encodee, ok := value.(mum.Encodee)
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

func (m *mumSerializer) Unmarshal(bs []byte, value interface{}) (err error) {
	decodee, ok := value.(mum.Decodee)
	if !ok {
		err = fmt.Errorf("invalid type, expected %t and received %t", decodee, value)
		return
	}

	m.r.SetBuffer(bs)
	return m.r.Decode(decodee)
}

// MarshalMum is an encoding helper func
func (a *A) MarshalMum(enc *mum.Encoder) {
	enc.String(a.Name)
	enc.Int64(a.BirthDay.UnixNano())
	enc.String(a.Phone)
	enc.Int(a.Siblings)
	enc.Bool(a.Spouse)
	enc.Float64(a.Money)
	return
}

// UnmarshalMum is an decoding helper func
func (a *A) UnmarshalMum(dec *mum.Decoder) (err error) {
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
