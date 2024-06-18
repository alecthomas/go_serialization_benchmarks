package idr

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/chmike/ditp/idr/low"
)

type IDRSerializer struct {
	e low.Encoder
}

func (s IDRSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	e := low.Reset(s.e)
	e = low.PutInt64(e, v.BirthDay.UnixNano())
	e = low.PutFloat64(e, v.Money)
	e = low.PutInt32(e, int32(v.Siblings))
	e = low.PutBool(e, v.Spouse)
	e = low.PutString(e, v.Name)
	e = low.PutString(e, v.Phone)
	return e, nil
}

func (s IDRSerializer) Unmarshal(d []byte, o interface{}) (err error) {
	v := o.(*goserbench.SmallStruct)
	_ = d[21]
	d, v0 := low.Int64(d)
	v.BirthDay = time.Unix(0, v0)
	d, v.Money = low.Float64(d)
	d, v1 := low.Int32(d)
	v.Siblings = int(v1)
	d, v.Spouse = low.Bool(d)
	d, v.Name = low.String(d, 1<<14)
	_, v.Phone = low.String(d, 20)
	return nil
}

func NewIDRSerializer() goserbench.Serializer {
	return IDRSerializer{e: make([]byte, 256)} // set initial buffer to 256 bytes
}
