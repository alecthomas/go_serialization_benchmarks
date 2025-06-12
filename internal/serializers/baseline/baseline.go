package baseline

import (
	"encoding/binary"
	"math"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

// BaselineSerializer is a baseline test for marshalling/unmarshalling. It
// assumes names are 16 bytes and phones are 10 bytes and may panic or produce
// incorrect results if this is not the case.
//
// This is useful as an upper bound on the performance achievable for standard
// marshalling/unmarshalling operations.
type BaselineSerializer struct{}

func (b *BaselineSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*goserbench.SmallStruct)
	buf := make([]byte, 0, maxSmallStructSerializeSize)
	buf = binary.LittleEndian.AppendUint64(buf, uint64(a.BirthDay.UnixNano()))
	buf = binary.LittleEndian.AppendUint64(buf, math.Float64bits(a.Money))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(a.Siblings))
	buf = append(buf, []byte(a.Name)...)
	buf = append(buf, []byte(a.Phone)...)
	buf = appendBool(buf, a.Spouse)
	return buf, nil
}

func (b *BaselineSerializer) Unmarshal(d []byte, o interface{}) error {
	a := o.(*goserbench.SmallStruct)
	a.BirthDay = time.Unix(0, int64(binary.LittleEndian.Uint64(d[:8])))
	a.Money = math.Float64frombits(binary.LittleEndian.Uint64(d[8:16]))
	a.Siblings = int(binary.LittleEndian.Uint32(d[16:20]))
	a.Name = string(d[20:36])
	a.Phone = string(d[36:46])
	a.Spouse = getBool(d[46:])
	return nil
}

func NewBaselineSerializer() goserbench.Serializer {
	return &BaselineSerializer{}
}

// BaselineUnsafeSerializer is similar to BaselizeSerializer, but it reuses
// the marshalling buffer and unmarshals unsafe strings.
type BaselineUnsafeSerializer struct {
	buf []byte
}

func (b *BaselineUnsafeSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*goserbench.SmallStruct)
	buf := b.buf
	buf = binary.LittleEndian.AppendUint64(buf, uint64(a.BirthDay.UnixNano()))
	buf = binary.LittleEndian.AppendUint64(buf, math.Float64bits(a.Money))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(a.Siblings))
	buf = append(buf, []byte(a.Name)...)
	buf = append(buf, []byte(a.Phone)...)
	buf = appendBool(buf, a.Spouse)
	return buf, nil
}

func (b *BaselineUnsafeSerializer) Unmarshal(d []byte, o interface{}) error {
	a := o.(*goserbench.SmallStruct)
	a.BirthDay = time.Unix(0, int64(binary.LittleEndian.Uint64(d[:8])))
	a.Money = math.Float64frombits(binary.LittleEndian.Uint64(d[8:16]))
	a.Siblings = int(binary.LittleEndian.Uint32(d[16:20]))
	a.Name = unsafeSliceToString(d[20:36])
	a.Phone = unsafeSliceToString(d[36:46])
	a.Spouse = getBool(d[46:])
	return nil
}

func NewBaselineUnsafeSerializer() goserbench.Serializer {
	return &BaselineUnsafeSerializer{
		buf: make([]byte, 0, maxSmallStructSerializeSize),
	}
}
