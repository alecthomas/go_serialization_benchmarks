package flatbuffers

import (
	"time"
	"unsafe"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatBufferCodec struct {
	goserbench.NativeSmallStruct

	unsafeReuse bool
	builder     *flatbuffers.Builder
}

func (s *FlatBufferCodec) Marshal(a *goserbench.SmallStruct) ([]byte, error) {
	builder := s.builder
	if !s.unsafeReuse {
		builder.Bytes = nil // free
	}
	builder.Reset()

	name := builder.CreateString(a.Name)
	phone := builder.CreateString(a.Phone)

	FlatBufferAStart(builder)
	FlatBufferAAddName(builder, name)
	FlatBufferAAddPhone(builder, phone)
	FlatBufferAAddBirthDay(builder, a.BirthDay.UnixNano())
	FlatBufferAAddSiblings(builder, int32(a.Siblings))
	FlatBufferAAddSpouse(builder, a.Spouse)
	FlatBufferAAddMoney(builder, a.Money)
	builder.Finish(FlatBufferAEnd(builder))
	return builder.FinishedBytes(), nil
}

func (s *FlatBufferCodec) Unmarshal(d []byte, a *goserbench.SmallStruct) error {
	o := FlatBufferA{}
	o.Init(d, flatbuffers.GetUOffsetT(d))
	if s.unsafeReuse {
		a.Name = unsafeSliceToString(o.Name())
		a.Phone = unsafeSliceToString(o.Phone())
	} else {
		a.Name = string(o.Name())
		a.Phone = string(o.Phone())
	}
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func NewFlatBuffersCodec() goserbench.Codec[goserbench.SmallStruct] {
	return &FlatBufferCodec{builder: flatbuffers.NewBuilder(0), unsafeReuse: false}
}

func NewFlatBuffersUnsafeReuseCodec() goserbench.Codec[goserbench.SmallStruct] {
	const maxSerSize = 128
	return &FlatBufferCodec{builder: flatbuffers.NewBuilder(maxSerSize), unsafeReuse: true}
}

// unsafeSliceToString converts a byte slice to a string in an unsafe way
// (modifications to the byte slice modify the string).
func unsafeSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
