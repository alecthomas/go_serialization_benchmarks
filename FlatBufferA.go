// automatically generated, do not modify

package goserbench

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatBufferA struct {
	_tab flatbuffers.Table
}

func (rcv *FlatBufferA) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlatBufferA) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatBufferA) BirthDay() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FlatBufferA) Phone() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatBufferA) Siblings() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FlatBufferA) Spouse() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FlatBufferA) Money() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0
}

func FlatBufferAStart(builder *flatbuffers.Builder) { builder.StartObject(6) }
func FlatBufferAAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func FlatBufferAAddBirthDay(builder *flatbuffers.Builder, birthDay int64) {
	builder.PrependInt64Slot(1, birthDay, 0)
}
func FlatBufferAAddPhone(builder *flatbuffers.Builder, phone flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(phone), 0)
}
func FlatBufferAAddSiblings(builder *flatbuffers.Builder, siblings int32) {
	builder.PrependInt32Slot(3, siblings, 0)
}
func FlatBufferAAddSpouse(builder *flatbuffers.Builder, spouse byte) {
	builder.PrependByteSlot(4, spouse, 0)
}
func FlatBufferAAddMoney(builder *flatbuffers.Builder, money float64) {
	builder.PrependFloat64Slot(5, money, 0)
}
func FlatBufferAEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
