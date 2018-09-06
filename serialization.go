package goserbench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"math/rand"
	"reflect"
	"time"
	"unsafe"

	"github.com/mailru/easyjson"
	"github.com/niubaoshu/gotiny"
	"zombiezen.com/go/capnproto2"

	"github.com/DeDiS/protobuf"
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/glycerine/go-capnproto"
	"github.com/gogo/protobuf/proto"
	"github.com/google/flatbuffers/go"
	"github.com/hprose/hprose-go"
	hprose2 "github.com/hprose/hprose-golang/io"
	"github.com/ikkerens/ikeapack"
	"github.com/json-iterator/go"
	"github.com/tinylib/msgp/msgp"
	"github.com/ugorji/go/codec"
	"gopkg.in/mgo.v2/bson"
	vmihailenco "gopkg.in/vmihailenco/msgpack.v2"

	"github.com/alecthomas/binary"
)

var B = NewBenchSerializer()

type NeedGeneratingCodeFalse struct{}

func (NeedGeneratingCodeFalse) NeedGeneratingCode() bool {
	return false
}

type NeedGeneratingCodeTrue struct{}

func (NeedGeneratingCodeTrue) NeedGeneratingCode() bool {
	return true
}

// github.com/niubaoshu/gotiny

type GotinySerializer struct {
	NeedGeneratingCodeFalse
	name string
	enc  *gotiny.Encoder
	dec  *gotiny.Decoder
}

func (g GotinySerializer) Marshal(o interface{}) ([]byte, error) {
	return g.enc.Encode(o), nil
}

func (g GotinySerializer) Unmarshal(d []byte, o interface{}) error {
	g.dec.Decode(d, o)
	return nil
}
func (g GotinySerializer) String() string { return g.name }

func NewGotinySerializer(o interface{}, name string) Serializer {
	ot := reflect.TypeOf(o)
	return GotinySerializer{
		name: name,
		enc:  gotiny.NewEncoderWithType(ot),
		dec:  gotiny.NewDecoderWithType(ot),
	}
}

func init() { B.AddSerializer(NewGotinySerializer(A{}, "gotiny"), NewA) }
func init() { B.AddSerializer(NewGotinySerializer(NoTimeA{}, "gotiny_notime"), NewNoTimeA) }

// github.com/tinylib/msgp

type MsgpSerializer struct {
	NeedGeneratingCodeTrue
}

func (MsgpSerializer) Marshal(o interface{}) ([]byte, error) {
	return o.(msgp.Marshaler).MarshalMsg(nil)
}

func (MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(msgp.Unmarshaler).UnmarshalMsg(d)
	return err
}

func (MsgpSerializer) String() string { return "msgp" }

func init() { B.AddSerializer(MsgpSerializer{}, NewA) }

// gopkg.in/vmihailenco/msgpack.v2

type VmihailencoMsgpackSerializer struct {
	NeedGeneratingCodeFalse
}

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return vmihailenco.Marshal(o)
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func (m VmihailencoMsgpackSerializer) String() string {
	return "msgpack"
}

func init() { B.AddSerializer(VmihailencoMsgpackSerializer{}, NewA) }

// encoding/json

type JsonSerializer struct {
	NeedGeneratingCodeFalse
}

func (j JsonSerializer) Marshal(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func (j JsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func (j JsonSerializer) String() string {
	return "json"
}

func init() { B.AddSerializer(JsonSerializer{}, NewA) }

// github.com/json-iterator/go

type JsonIterSerializer struct {
	NeedGeneratingCodeFalse
}

func (j JsonIterSerializer) Marshal(o interface{}) ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(o)
}

func (j JsonIterSerializer) Unmarshal(d []byte, o interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(d, o)
}

func (j JsonIterSerializer) String() string {
	return "jsoniter"
}

func init() { B.AddSerializer(JsonIterSerializer{}, NewA) }

// github.com/mailru/easyjson

type EasyJSONSerializer struct {
	NeedGeneratingCodeTrue
}

func (m EasyJSONSerializer) Marshal(o interface{}) ([]byte, error) {
	return easyjson.Marshal(o.(easyjson.Marshaler))
}

func (m EasyJSONSerializer) Unmarshal(d []byte, o interface{}) error {
	return easyjson.Unmarshal(d, o.(easyjson.Unmarshaler))
}

func (m EasyJSONSerializer) String() string { return "easyJson" }

func init() { B.AddSerializer(EasyJSONSerializer{}, NewA) }

// gopkg.in/mgo.v2/bson

type BsonSerializer struct {
	NeedGeneratingCodeFalse
}

func (m BsonSerializer) Marshal(o interface{}) ([]byte, error) {
	return bson.Marshal(o)
}

func (m BsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return bson.Unmarshal(d, o)
}

func (j BsonSerializer) String() string {
	return "bson"
}

func init() { B.AddSerializer(BsonSerializer{}, NewA) }

// encoding/gob

type GobSerializer struct {
	NeedGeneratingCodeFalse
	b   bytes.Buffer
	enc *gob.Encoder
	dec *gob.Decoder
}

func (g *GobSerializer) Marshal(o interface{}) ([]byte, error) {
	g.b.Reset()
	err := g.enc.Encode(o)
	return g.b.Bytes(), err
}

func (g *GobSerializer) Unmarshal(d []byte, o interface{}) error {
	g.b.Reset()
	g.b.Write(d)
	return g.dec.Decode(o)
}

func (g GobSerializer) String() string {
	return "gob"
}

func NewGobSerializer() Serializer {
	s := &GobSerializer{}
	s.enc = gob.NewEncoder(&s.b)
	s.dec = gob.NewDecoder(&s.b)
	var a A
	s.enc.Encode(a)
	s.dec.Decode(&a)
	return s
}

func init() { B.AddSerializer(NewGobSerializer(), NewA) }

// github.com/davecgh/go-xdr/xdr

type XdrSerializer struct {
	NeedGeneratingCodeFalse
}

func (x XdrSerializer) Marshal(o interface{}) ([]byte, error) {
	return xdr.Marshal(o)
}

func (x XdrSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func (x XdrSerializer) String() string {
	return "xdr"
}

func init() { B.AddSerializer(XdrSerializer{}, NewA) }

// github.com/ugorji/go/codec

type UgorjiCodecSerializer struct {
	NeedGeneratingCodeFalse
	h codec.Handle
}

func NewUgorjiCodecSerializer() Serializer {
	return &UgorjiCodecSerializer{
		h: &codec.BincHandle{},
	}
}

func (u *UgorjiCodecSerializer) Marshal(o interface{}) ([]byte, error) {
	var bs []byte
	return bs, codec.NewEncoderBytes(&bs, u.h).Encode(o)
}

func (u *UgorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	return codec.NewDecoderBytes(d, u.h).Decode(o)
}

func (u *UgorjiCodecSerializer) String() string {
	return "ugorjicodec"
}

func init() { B.AddSerializer(NewUgorjiCodecSerializer(), NewA) }

// github.com/Sereal/Sereal/Go/sereal

type SerealSerializer struct {
	NeedGeneratingCodeFalse
}

func (s SerealSerializer) Marshal(o interface{}) ([]byte, error) {
	return sereal.Marshal(o)
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	return sereal.Unmarshal(d, o)
}

func (s SerealSerializer) String() string {
	return "sereal"
}

func init() { B.AddSerializer(SerealSerializer{}, NewA) }

// github.com/alecthomas/binary

type BinarySerializer struct {
	NeedGeneratingCodeFalse
}

func (b BinarySerializer) Marshal(o interface{}) ([]byte, error) {
	return binary.Marshal(o)
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func (b BinarySerializer) String() string {
	return "alecthomas"
}

func init() { B.AddSerializer(BinarySerializer{}, NewA) }

// github.com/google/flatbuffers/go

type FlatBufferSerializer struct {
	builder *flatbuffers.Builder
	NeedGeneratingCodeTrue
}

func (s *FlatBufferSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	builder := s.builder

	builder.Reset()

	name := builder.CreateString(a.Name)
	phone := builder.CreateString(a.Phone)

	FlatBufferAStart(builder)
	FlatBufferAAddName(builder, name)
	FlatBufferAAddPhone(builder, phone)
	FlatBufferAAddBirthDay(builder, a.BirthDay.UnixNano())
	FlatBufferAAddSiblings(builder, int32(a.Siblings))
	var spouse byte
	if a.Spouse {
		spouse = byte(1)
	}
	FlatBufferAAddSpouse(builder, spouse)
	FlatBufferAAddMoney(builder, a.Money)
	builder.Finish(FlatBufferAEnd(builder))
	return builder.Bytes[builder.Head():], nil
}

func (s *FlatBufferSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	o := FlatBufferA{}
	o.Init(d, flatbuffers.GetUOffsetT(d))
	a.Name = string(o.Name())
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone = string(o.Phone())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse() == byte(1)
	a.Money = o.Money()
	return nil
}

func (s *FlatBufferSerializer) String() string {
	return "FlatBuffer"
}

func NewFlatBufferSerializer() Serializer {
	return &FlatBufferSerializer{builder: flatbuffers.NewBuilder(0)}
}

func init() { B.AddSerializer(NewFlatBufferSerializer(), NewA) }

// github.com/glycerine/go-capnproto

type CapNProtoSerializer struct {
	buf []byte
	out *bytes.Buffer
	NeedGeneratingCodeTrue
}

func (x *CapNProtoSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	s := capn.NewBuffer(x.buf)
	c := NewRootCapnpA(s)
	c.SetName(a.Name)
	c.SetBirthDay(a.BirthDay.UnixNano())
	c.SetPhone(a.Phone)
	c.SetSiblings(int32(a.Siblings))
	c.SetSpouse(a.Spouse)
	c.SetMoney(a.Money)
	x.out.Reset()
	s.WriteTo(x.out)
	x.buf = s.Data[:0]
	return x.out.Bytes(), nil
}

func (x *CapNProtoSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	s, _, _ := capn.ReadFromMemoryZeroCopy(d)
	o := ReadRootCapnpA(s)
	a.Name = string(o.NameBytes())
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone = string(o.PhoneBytes())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func (x *CapNProtoSerializer) String() string {
	return "CapNProto"
}

func NewCapNProtoSerializer() Serializer {
	return &CapNProtoSerializer{
		out: new(bytes.Buffer),
	}
}

func init() { B.AddSerializer(NewCapNProtoSerializer(), NewA) }

// zombiezen.com/go/capnproto2

type CapNProto2Serializer struct {
	arena capnp.Arena
	NeedGeneratingCodeTrue
}

func (x *CapNProto2Serializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	m, s, _ := capnp.NewMessage(x.arena)
	c, _ := NewRootCapnp2A(s)
	c.SetName(a.Name)
	c.SetBirthDay(a.BirthDay.UnixNano())
	c.SetPhone(a.Phone)
	c.SetSiblings(int32(a.Siblings))
	c.SetSpouse(a.Spouse)
	c.SetMoney(a.Money)
	return m.Marshal()
}

func (x *CapNProto2Serializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	m, _ := capnp.Unmarshal(d)
	o, _ := ReadRootCapnp2A(m)
	a.Name, _ = o.Name()
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone, _ = o.Phone()
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func (x *CapNProto2Serializer) String() string {
	return "CapNProto2"
}

func NewCapNProto2Serializer() Serializer {
	return &CapNProto2Serializer{arena: capnp.SingleSegment(nil)}
}

func init() { B.AddSerializer(NewCapNProto2Serializer(), NewA) }

// github.com/hprose/hprose-go/io

type HproseSerializer struct {
	writer *hprose.Writer
	reader *hprose.Reader
	NeedGeneratingCodeFalse
}

func (s *HproseSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	writer := s.writer
	buf := writer.Stream.(*bytes.Buffer)
	l := buf.Len()
	writer.WriteString(a.Name)
	writer.WriteTime(a.BirthDay)
	writer.WriteString(a.Phone)
	writer.WriteInt64(int64(a.Siblings))
	writer.WriteBool(a.Spouse)
	writer.WriteFloat64(a.Money)
	return buf.Bytes()[l:], nil
}

func (s *HproseSerializer) Unmarshal(d []byte, i interface{}) error {
	o := i.(*A)
	reader := s.reader
	reader.Stream = &hprose.BytesReader{d, 0}
	o.Name, _ = reader.ReadString()
	o.BirthDay, _ = reader.ReadDateTime()
	o.Phone, _ = reader.ReadString()
	o.Siblings, _ = reader.ReadInt()
	o.Spouse, _ = reader.ReadBool()
	o.Money, _ = reader.ReadFloat64()
	return nil
}

func (s *HproseSerializer) String() string {
	return "Hprose"
}

func NewHproseSerializer() Serializer {
	buf := new(bytes.Buffer)
	reader := hprose.NewReader(buf, true)
	bufw := new(bytes.Buffer)
	writer := hprose.NewWriter(bufw, true)
	return &HproseSerializer{writer: writer, reader: reader}
}

func init() { B.AddSerializer(NewHproseSerializer(), NewA) }

// github.com/hprose/hprose-golang/io

type Hprose2Serializer struct {
	writer *hprose2.Writer
	reader *hprose2.Reader
	NeedGeneratingCodeFalse
}

func (s Hprose2Serializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	writer := s.writer
	writer.Clear()
	writer.WriteString(a.Name)
	writer.WriteTime(&a.BirthDay)
	writer.WriteString(a.Phone)
	writer.WriteInt(int64(a.Siblings))
	writer.WriteBool(a.Spouse)
	writer.WriteFloat(a.Money, 64)
	return writer.Bytes(), nil
}

func (s Hprose2Serializer) Unmarshal(d []byte, i interface{}) error {
	o := i.(*A)
	reader := s.reader
	reader.Init(d)
	o.Name = reader.ReadString()
	o.BirthDay = reader.ReadTime()
	o.Phone = reader.ReadString()
	o.Siblings = int(reader.ReadInt())
	o.Spouse = reader.ReadBool()
	o.Money = reader.ReadFloat64()
	return nil
}

func (s Hprose2Serializer) String() string {
	return "Hprose2"
}

func NewHprose2Serializer() Serializer {
	writer := hprose2.NewWriter(true)
	reader := hprose2.NewReader(nil, true)
	return &Hprose2Serializer{writer: writer, reader: reader}
}

func init() { B.AddSerializer(NewHprose2Serializer(), NewA) }

// github.com/DeDiS/protobuf

type ProtobufSerializer struct {
	NeedGeneratingCodeFalse
}

func (m ProtobufSerializer) Marshal(o interface{}) ([]byte, error) {
	return protobuf.Encode(o)
}

func (m ProtobufSerializer) Unmarshal(d []byte, o interface{}) error {
	return protobuf.Decode(d, o)
}

func (m ProtobufSerializer) String() string {
	return "DeDiS-protobuf"
}

func init() { B.AddSerializer(ProtobufSerializer{}, NewA) }

// github.com/golang/protobuf

func NewProtoBufA() Object {
	return &ProtoBufA{
		Name:     proto.String(randString(16)),
		BirthDay: proto.Int64(time.Now().UnixNano()),
		Phone:    proto.String(randString(10)),
		Siblings: proto.Int32(rand.Int31n(5)),
		Spouse:   proto.Bool(rand.Intn(2) == 1),
		Money:    proto.Float64(rand.Float64()),
	}
}

func (i *ProtoBufA) AssertEqual(a interface{}) (bool, error) {
	o, ok := a.(*ProtoBufA)
	if !ok {
		return false, errors.New("not A type")
	}
	correct := *o.Name == *i.Name && *o.Phone == *i.Phone && *o.Siblings == *i.Siblings && *o.Spouse == *i.Spouse && *o.Money == *i.Money && *o.BirthDay == *i.BirthDay
	return correct, nil
}

type GoprotobufSerializer struct {
	NeedGeneratingCodeTrue
}

func (GoprotobufSerializer) Marshal(o interface{}) ([]byte, error) {
	return proto.Marshal(o.(proto.Message))
}

func (GoprotobufSerializer) Unmarshal(d []byte, o interface{}) error {
	return proto.Unmarshal(d, o.(proto.Message))
}

func (GoprotobufSerializer) String() string {
	return "golang-protobuf"
}

func init() { B.AddSerializer(GoprotobufSerializer{}, NewProtoBufA) }

// github.com/gogo/protobuf/proto

func NewGogoProtoBufA() Object {
	return &GogoProtoBufA{
		Name:     randString(16),
		BirthDay: time.Now().UnixNano(),
		Phone:    randString(10),
		Siblings: rand.Int31n(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}

func (i *GogoProtoBufA) AssertEqual(a interface{}) (bool, error) {
	o, ok := a.(*GogoProtoBufA)
	if !ok {
		return false, errors.New("not A type")
	}
	return *i == *o, nil
}

type GogoProtoBufSerializer struct {
	GoprotobufSerializer
}

func (GogoProtoBufSerializer) String() string {
	return "gogo-protobuf"
}

func init() { B.AddSerializer(GogoProtoBufSerializer{}, NewGogoProtoBufA) }

// github.com/pascaldekloe/colfer

func NewColferA() Object {
	return &ColferA{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Int31n(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}

func (a *ColferA) Reset() { *a = ColferA{} }

func (a *ColferA) AssertEqual(i interface{}) (bool, error) {
	return (*A)(unsafe.Pointer(a)).AssertEqual((*A)(unsafe.Pointer(i.(*ColferA))))
}

type ColferSerializer struct {
	NeedGeneratingCodeTrue
}

func (ColferSerializer) Marshal(o interface{}) ([]byte, error) {
	return o.(*ColferA).MarshalBinary()
}

func (ColferSerializer) Unmarshal(d []byte, o interface{}) error {
	return o.(*ColferA).UnmarshalBinary(d)
}

func (ColferSerializer) String() string {
	return "colfer"
}

func init() { B.AddSerializer(ColferSerializer{}, NewColferA) }

// github.com/andyleap/gencode

func NewGencodeA() Object {
	return &GencodeA{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Int63n(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}

func (a *GencodeA) Reset() { *a = GencodeA{} }

func (a *GencodeA) AssertEqual(i interface{}) (bool, error) {
	return (*A)(unsafe.Pointer(a)).AssertEqual((*A)(unsafe.Pointer(i.(*GencodeA))))
}

type GencodeSerializer struct {
	NeedGeneratingCodeTrue
}

func (GencodeSerializer) Marshal(o interface{}) ([]byte, error) {
	return o.(interface{ Marshal([]byte) ([]byte, error) }).Marshal(nil)
}

func (GencodeSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(interface{ Unmarshal([]byte) (uint64, error) }).Unmarshal(d)
	return err
}

func (GencodeSerializer) String() string {
	return "gencode"
}

func init() { B.AddSerializer(GencodeSerializer{}, NewGencodeA) }

func NewGencodeUnsafeA() Object {
	return &GencodeUnsafeA{
		Name:     randString(16),
		BirthDay: time.Now().UnixNano(),
		Phone:    randString(10),
		Siblings: rand.Int63n(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}

func (a *GencodeUnsafeA) Reset() { *a = GencodeUnsafeA{} }

func (a *GencodeUnsafeA) AssertEqual(i interface{}) (bool, error) {
	o, ok := i.(*GencodeUnsafeA)
	if !ok {
		return false, errors.New("not A type")
	}
	return *a == *o, nil
}

type GencodeUnsafeSerializer struct {
	GencodeSerializer
}

func (GencodeUnsafeSerializer) String() string {
	return "unsafe-gencode"
}

func init() { B.AddSerializer(GencodeUnsafeSerializer{}, NewGencodeUnsafeA) }

// github.com/calmh/xdr

type XDRSerializer struct {
	NeedGeneratingCodeTrue
}

func (XDRSerializer) Marshal(o interface{}) ([]byte, error) {
	return o.(*XDRA).MarshalXDR()
}

func (XDRSerializer) Unmarshal(d []byte, o interface{}) error {
	return o.(*XDRA).UnmarshalXDR(d)
}

func (XDRSerializer) String() string {
	return "XDR"
}
func init() { B.AddSerializer(XDRSerializer{}, NewXDRA) }

// gopkg.in/linkedin/goavro.v1

func init() { B.AddSerializer(NewAvroA(), NewA) }

func init() { B.AddSerializer(NewAvro2Txt(), NewA) }

func init() { B.AddSerializer(NewAvro2Bin(), NewA) }

// github.com/ikkerens/ikeapack

type IkeapackSerializer struct {
	NeedGeneratingCodeFalse
	buf bytes.Buffer
}

func (i IkeapackSerializer) Marshal(o interface{}) ([]byte, error) {
	i.buf.Reset()
	err := ikea.Pack(&i.buf, o)
	return i.buf.Bytes(), err
}

func (i IkeapackSerializer) Unmarshal(d []byte, o interface{}) error {
	i.buf.Reset()
	i.buf.Write(d)
	return ikea.Unpack(&i.buf, o)
}

func (IkeapackSerializer) String() string { return "Ikeapack" }

func init() { B.AddSerializer(IkeapackSerializer{}, NewGencodeUnsafeA) }
