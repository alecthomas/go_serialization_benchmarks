package goserbench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/binary"
	"github.com/davecgh/go-xdr/xdr"
	capn "github.com/glycerine/go-capnproto"
	"github.com/gogo/protobuf/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/hprose/hprose-go"
	hprose2 "github.com/hprose/hprose-golang/io"
	ikea "github.com/ikkerens/ikeapack"
	jsoniter "github.com/json-iterator/go"
	easyjson "github.com/mailru/easyjson"
	"github.com/niubaoshu/gotiny"
	ssz "github.com/prysmaticlabs/go-ssz"
	shamaton "github.com/shamaton/msgpack"
	"github.com/tinylib/msgp/msgp"
	"github.com/ugorji/go/codec"
	vmihailenco "github.com/vmihailenco/msgpack/v4"
	"go.dedis.ch/protobuf"
	"gopkg.in/mgo.v2/bson"
	capnp "zombiezen.com/go/capnproto2"
)

var (
	validate     = os.Getenv("VALIDATE")
	jsoniterFast = jsoniter.ConfigFastest
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate() []*A {
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

type Serializer interface {
	Marshal(o interface{}) ([]byte, error)
	Unmarshal(d []byte, o interface{}) error
}

func benchMarshal(b *testing.B, s Serializer) {
	b.Helper()
	data := generate()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		o := data[rand.Intn(len(data))]
		bytes, err := s.Marshal(o)
		if err != nil {
			b.Fatalf("marshal error %s for %#v", err, o)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func cmpTags(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}

func cmpAliases(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func benchUnmarshal(b *testing.B, s Serializer) {
	b.Helper()
	b.StopTimer()
	data := generate()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		o, err := s.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		t := make([]byte, len(o))
		serialSize += copy(t, o)
		ser[i] = t
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("unmarshal error %s for %#x / %q", err, ser[n], ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.Equal(i.BirthDay) //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

func TestMessage(t *testing.T) {
	println(`
A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.
`)

}

// github.com/niubaoshu/gotiny

type GotinySerializer struct {
	enc *gotiny.Encoder
	dec *gotiny.Decoder
}

func (g GotinySerializer) Marshal(o interface{}) ([]byte, error) {
	return g.enc.Encode(o), nil
}

func (g GotinySerializer) Unmarshal(d []byte, o interface{}) error {
	g.dec.Decode(d, o)
	return nil
}

func NewGotinySerializer(o interface{}) Serializer {
	ot := reflect.TypeOf(o)
	return GotinySerializer{
		enc: gotiny.NewEncoderWithType(ot),
		dec: gotiny.NewDecoderWithType(ot),
	}
}

func BenchmarkGotinyMarshal(b *testing.B) {
	benchMarshal(b, NewGotinySerializer(A{}))
}

func BenchmarkGotinyUnmarshal(b *testing.B) {
	benchUnmarshal(b, NewGotinySerializer(A{}))
}

func generateNoTimeA() []*NoTimeA {
	a := make([]*NoTimeA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &NoTimeA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGotinyNoTimeMarshal(b *testing.B) {
	s := NewGotinySerializer(NoTimeA{})
	data := generateNoTimeA()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := s.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkGotinyNoTimeUnmarshal(b *testing.B) {
	b.StopTimer()
	s := NewGotinySerializer(NoTimeA{})
	data := generateNoTimeA()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		o, err := s.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		t := make([]byte, len(o))
		serialSize += copy(t, o)
		ser[i] = t
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &NoTimeA{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/tinylib/msgp

type MsgpSerializer struct{}

func (m MsgpSerializer) Marshal(o interface{}) ([]byte, error) {
	return o.(msgp.Marshaler).MarshalMsg(nil)
}

func (m MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(msgp.Unmarshaler).UnmarshalMsg(d)
	return err
}

func BenchmarkMsgpMarshal(b *testing.B) {
	benchMarshal(b, MsgpSerializer{})
}

func BenchmarkMsgpUnmarshal(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{})
}

// github.com/vmihailenco/msgpack

type VmihailencoMsgpackSerializer struct{}

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return vmihailenco.Marshal(o)
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func BenchmarkVmihailencoMsgpackMarshal(b *testing.B) {
	benchMarshal(b, VmihailencoMsgpackSerializer{})
}

func BenchmarkVmihailencoMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, VmihailencoMsgpackSerializer{})
}

// encoding/json

type JsonSerializer struct{}

func (j JsonSerializer) Marshal(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func (j JsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func BenchmarkJsonMarshal(b *testing.B) {
	benchMarshal(b, JsonSerializer{})
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{})
}

// github.com/json-iterator/go

type JsonIterSerializer struct{}

func (j JsonIterSerializer) Marshal(o interface{}) ([]byte, error) {
	return jsoniterFast.Marshal(o)
}

func (j JsonIterSerializer) Unmarshal(d []byte, o interface{}) error {
	return jsoniterFast.Unmarshal(d, o)
}

func BenchmarkJsonIterMarshal(b *testing.B) {
	benchMarshal(b, JsonIterSerializer{})
}

func BenchmarkJsonIterUnmarshal(b *testing.B) {
	benchUnmarshal(b, JsonIterSerializer{})
}

// github.com/mailru/easyjson

type EasyJSONSerializer struct{}

func (m EasyJSONSerializer) Marshal(o interface{}) ([]byte, error) {
	return easyjson.Marshal(o.(easyjson.Marshaler))
}

func (m EasyJSONSerializer) Unmarshal(d []byte, o interface{}) error {
	return easyjson.Unmarshal(d, o.(*A))
}

func BenchmarkEasyJsonMarshal(b *testing.B) {
	benchMarshal(b, EasyJSONSerializer{})
}

func BenchmarkEasyJsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, EasyJSONSerializer{})
}

// gopkg.in/mgo.v2/bson

type BsonSerializer struct{}

func (m BsonSerializer) Marshal(o interface{}) ([]byte, error) {
	return bson.Marshal(o)
}

func (m BsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return bson.Unmarshal(d, o)
}

func BenchmarkBsonMarshal(b *testing.B) {
	benchMarshal(b, BsonSerializer{})
}

func BenchmarkBsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, BsonSerializer{})
}

// encoding/gob

type GobSerializer struct {
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
	err := g.dec.Decode(o)
	return err
}

func NewGobSerializer() *GobSerializer {
	s := &GobSerializer{}
	s.enc = gob.NewEncoder(&s.b)
	s.dec = gob.NewDecoder(&s.b)
	err := s.enc.Encode(A{})
	if err != nil {
		panic(err)
	}
	var a A
	err = s.dec.Decode(&a)
	if err != nil {
		panic(err)
	}
	return s
}

func BenchmarkGobMarshal(b *testing.B) {
	s := NewGobSerializer()
	benchMarshal(b, s)
}

func BenchmarkGobUnmarshal(b *testing.B) {
	s := NewGobSerializer()
	benchUnmarshal(b, s)
}

// github.com/davecgh/go-xdr/xdr

type XDRSerializer struct{}

func (x XDRSerializer) Marshal(o interface{}) ([]byte, error) {
	return xdr.Marshal(o)
}

func (x XDRSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func BenchmarkXDRMarshal(b *testing.B) {
	benchMarshal(b, XDRSerializer{})
}

func BenchmarkXDRUnmarshal(b *testing.B) {
	benchUnmarshal(b, XDRSerializer{})
}

// github.com/ugorji/go/codec

type UgorjiCodecSerializer struct {
	codec.Handle
}

func (u *UgorjiCodecSerializer) Marshal(o interface{}) ([]byte, error) {
	var bs []byte
	return bs, codec.NewEncoderBytes(&bs, u.Handle).Encode(o)
}

func (u *UgorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	return codec.NewDecoderBytes(d, u.Handle).Decode(o)
}

func BenchmarkUgorjiCodecMsgpackMarshal(b *testing.B) {
	benchMarshal(b, &UgorjiCodecSerializer{&codec.MsgpackHandle{}})
}

func BenchmarkUgorjiCodecMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, &UgorjiCodecSerializer{&codec.MsgpackHandle{}})
}

func BenchmarkUgorjiCodecBincMarshal(b *testing.B) {
	h := &codec.BincHandle{}
	h.AsSymbols = 0
	benchMarshal(b, &UgorjiCodecSerializer{h})
}

func BenchmarkUgorjiCodecBincUnmarshal(b *testing.B) {
	h := &codec.BincHandle{}
	h.AsSymbols = 0
	benchUnmarshal(b, &UgorjiCodecSerializer{h})
}

// github.com/Sereal/Sereal/Go/sereal

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o interface{}) ([]byte, error) {
	return sereal.Marshal(o)
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func BenchmarkSerealMarshal(b *testing.B) {
	benchMarshal(b, SerealSerializer{})
}

func BenchmarkSerealUnmarshal(b *testing.B) {
	benchUnmarshal(b, SerealSerializer{})
}

// github.com/alecthomas/binary

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o interface{}) ([]byte, error) {
	return binary.Marshal(o)
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func BenchmarkBinaryMarshal(b *testing.B) {
	benchMarshal(b, BinarySerializer{})
}

func BenchmarkBinaryUnmarshal(b *testing.B) {
	benchUnmarshal(b, BinarySerializer{})
}

// github.com/google/flatbuffers/go

type FlatBufferSerializer struct {
	builder *flatbuffers.Builder
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

func BenchmarkFlatBuffersMarshal(b *testing.B) {
	benchMarshal(b, &FlatBufferSerializer{flatbuffers.NewBuilder(0)})
}

func BenchmarkFlatBuffersUnmarshal(b *testing.B) {
	benchUnmarshal(b, &FlatBufferSerializer{flatbuffers.NewBuilder(0)})
}

// github.com/glycerine/go-capnproto

type CapNProtoSerializer struct {
	buf []byte
	out *bytes.Buffer
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
	_, err := s.WriteTo(x.out)
	x.buf = []byte(s.Data)[:0]
	return x.out.Bytes(), err
}

func (x *CapNProtoSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	s, _, err := capn.ReadFromMemoryZeroCopy(d)
	if err != nil {
		return err
	}
	o := ReadRootCapnpA(s)
	a.Name = o.Name()
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone = o.Phone()
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func BenchmarkCapNProtoMarshal(b *testing.B) {
	benchMarshal(b, &CapNProtoSerializer{nil, &bytes.Buffer{}})
}

func BenchmarkCapNProtoUnmarshal(b *testing.B) {
	benchUnmarshal(b, &CapNProtoSerializer{nil, &bytes.Buffer{}})
}

// zombiezen.com/go/capnproto2

type CapNProto2Serializer struct {
	arena capnp.Arena
}

func (x *CapNProto2Serializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	m, s, err := capnp.NewMessage(x.arena)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	c, err := NewRootCapnp2A(s)
	if err != nil {
		return nil, err
	}
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
	m, err := capnp.Unmarshal(d)
	if err != nil {
		return err
	}
	o, err := ReadRootCapnp2A(m)
	if err != nil {
		return err
	}
	a.Name, err = o.Name()
	if err != nil {
		return err
	}
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone, err = o.Phone()
	if err != nil {
		return err
	}
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func BenchmarkCapNProto2Marshal(b *testing.B) {
	benchMarshal(b, &CapNProto2Serializer{capnp.SingleSegment(nil)})
}

func BenchmarkCapNProto2Unmarshal(b *testing.B) {
	benchUnmarshal(b, &CapNProto2Serializer{capnp.SingleSegment(nil)})
}

// github.com/hprose/hprose-go/io

type HproseSerializer struct {
	writer *hprose.Writer
	reader *hprose.Reader
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

func (s *HproseSerializer) Unmarshal(d []byte, i interface{}) (err error) {
	o := i.(*A)
	reader := s.reader
	reader.Stream = &hprose.BytesReader{Bytes: d, Pos: 0}
	o.Name, err = reader.ReadString()
	if err != nil {
		return err
	}
	o.BirthDay, err = reader.ReadDateTime()
	if err != nil {
		return err
	}
	o.Phone, err = reader.ReadString()
	if err != nil {
		return err
	}
	o.Siblings, err = reader.ReadInt()
	if err != nil {
		return err
	}
	o.Spouse, err = reader.ReadBool()
	if err != nil {
		return err
	}
	o.Money, err = reader.ReadFloat64()
	return err
}

func BenchmarkHproseMarshal(b *testing.B) {
	buf := new(bytes.Buffer)
	writer := hprose.NewWriter(buf, true)
	benchMarshal(b, &HproseSerializer{writer: writer})
}

func BenchmarkHproseUnmarshal(b *testing.B) {
	buf := new(bytes.Buffer)
	reader := hprose.NewReader(buf, true)
	bufw := new(bytes.Buffer)
	writer := hprose.NewWriter(bufw, true)
	benchUnmarshal(b, &HproseSerializer{writer: writer, reader: reader})
}

// github.com/hprose/hprose-golang/io

type Hprose2Serializer struct {
	writer *hprose2.Writer
	reader *hprose2.Reader
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

func BenchmarkHprose2Marshal(b *testing.B) {
	writer := hprose2.NewWriter(true)
	benchMarshal(b, Hprose2Serializer{writer: writer})
}

func BenchmarkHprose2Unmarshal(b *testing.B) {
	writer := hprose2.NewWriter(true)
	reader := hprose2.NewReader(nil, true)
	benchUnmarshal(b, &Hprose2Serializer{writer: writer, reader: reader})
}

// go.dedis.ch/protobuf

type ProtobufSerializer struct{}

func (m ProtobufSerializer) Marshal(o interface{}) ([]byte, error) {
	return protobuf.Encode(o)
}

func (m ProtobufSerializer) Unmarshal(d []byte, o interface{}) error {
	return protobuf.Decode(d, o)
}

func BenchmarkProtobufMarshal(b *testing.B) {
	benchMarshal(b, ProtobufSerializer{})
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{})
}

// github.com/golang/protobuf

func generateProto() []*ProtoBufA {
	a := make([]*ProtoBufA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &ProtoBufA{
			Name:     proto.String(randString(16)),
			BirthDay: proto.Int64(time.Now().UnixNano()),
			Phone:    proto.String(randString(10)),
			Siblings: proto.Int32(rand.Int31n(5)),
			Spouse:   proto.Bool(rand.Intn(2) == 1),
			Money:    proto.Float64(rand.Float64()),
		})
	}
	return a
}

func BenchmarkGoprotobufMarshal(b *testing.B) {
	data := generateProto()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkGoprotobufUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateProto()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &ProtoBufA{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := *o.Name == *i.Name && *o.Phone == *i.Phone && *o.Siblings == *i.Siblings && *o.Spouse == *i.Spouse && *o.Money == *i.Money && *o.BirthDay == *i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/gogo/protobuf/proto

func generateGogoProto() []*GogoProtoBufA {
	a := make([]*GogoProtoBufA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GogoProtoBufA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGogoprotobufMarshal(b *testing.B) {
	data := generateGogoProto()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkGogoprotobufUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoProto()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GogoProtoBufA{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/pascaldekloe/colfer

func generateColfer() []*ColferA {
	a := make([]*ColferA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &ColferA{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkColferMarshal(b *testing.B) {
	data := generateColfer()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].MarshalBinary()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkColferUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateColfer()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = d.MarshalBinary()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[0])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &ColferA{}
		if err := o.UnmarshalBinary(ser[n]); err != nil {
			b.Fatalf("Colfer failed to unmarshal %#v: %s", data[n], err)
		}
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.Equal(i.BirthDay)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/andyleap/gencode

func generateGencode() []*GencodeA {
	a := make([]*GencodeA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GencodeA{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGencodeMarshal(b *testing.B) {
	data := generateGencode()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].Marshal(nil)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkGencodeUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGencode()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = d.Marshal(nil)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GencodeA{}
		_, err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("gencode failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.Equal(i.BirthDay) //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

func generateGencodeUnsafe() []*GencodeUnsafeA {
	a := make([]*GencodeUnsafeA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GencodeUnsafeA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGencodeUnsafeMarshal(b *testing.B) {
	data := generateGencodeUnsafe()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].Marshal(nil)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkGencodeUnsafeUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGencodeUnsafe()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = d.Marshal(nil)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GencodeUnsafeA{}
		_, err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("gencode failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/calmh/xdr

func generateXDR() []*XDRA {
	a := make([]*XDRA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &XDRA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    math.Float64bits(rand.Float64()),
		})
	}
	return a
}

func BenchmarkXDR2Marshal(b *testing.B) {
	data := generateXDR()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].MarshalXDR()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkXDR2Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateXDR()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		ser[i] = d.MustMarshalXDR()
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := XDRA{}
		err := o.UnmarshalXDR(ser[n])
		if err != nil {
			b.Fatalf("xdr failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// gopkg.in/linkedin/goavro.v1

func BenchmarkGoAvroMarshal(b *testing.B) {
	benchMarshal(b, NewAvroA())
}

func BenchmarkGoAvroUnmarshal(b *testing.B) {
	benchUnmarshal(b, NewAvroA())
}

// github.com/linkedin/goavro

func BenchmarkGoAvro2TextMarshal(b *testing.B) {
	benchMarshal(b, NewAvro2Txt())
}

func BenchmarkGoAvro2TextUnmarshal(b *testing.B) {
	benchUnmarshal(b, NewAvro2Txt())
}

func BenchmarkGoAvro2BinaryMarshal(b *testing.B) {
	benchMarshal(b, NewAvro2Bin())
}

func BenchmarkGoAvro2BinaryUnmarshal(b *testing.B) {
	benchUnmarshal(b, NewAvro2Bin())
}

// github.com/ikkerens/ikeapack

type IkeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64
}

func generateIkeA() []*IkeA {
	a := make([]*IkeA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &IkeA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    math.Float64bits(rand.Float64()),
		})
	}
	return a
}

func BenchmarkIkeaMarshal(b *testing.B) {
	buf := new(bytes.Buffer)
	buf.Grow(100)
	data := generateIkeA()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		ikea.Pack(buf, data[rand.Intn(len(data))])
		serialSize += buf.Len()
		buf.Reset()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkIkeaUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateIkeA()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		buf := new(bytes.Buffer)
		ikea.Pack(buf, d)
		ser[i] = buf.Bytes()
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	buf := new(bytes.Buffer)
	buf.Grow(100)
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := IkeA{}
		buf.Reset()
		buf.Write(ser[n])
		err := ikea.Unpack(buf, &o)
		if err != nil {
			b.Fatalf("ikea failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/shamaton/msgpack - as map

type ShamatonMapMsgpackSerializer struct{}

func (m ShamatonMapMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamaton.EncodeStructAsMap(o)
}

func (m ShamatonMapMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamaton.DecodeStructAsMap(d, o)
}

func BenchmarkShamatonMapMsgpackMarshal(b *testing.B) {
	benchMarshal(b, ShamatonMapMsgpackSerializer{})
}

func BenchmarkShamatonMapMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, ShamatonMapMsgpackSerializer{})
}

// github.com/shamaton/msgpack - as array

type ShamatonArrayMsgpackSerializer struct{}

func (m ShamatonArrayMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return shamaton.EncodeStructAsArray(o)
}

func (m ShamatonArrayMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return shamaton.DecodeStructAsArray(d, o)
}

func BenchmarkShamatonArrayMsgpackMarshal(b *testing.B) {
	benchMarshal(b, ShamatonArrayMsgpackSerializer{})
}

func BenchmarkShamatonArrayMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, ShamatonArrayMsgpackSerializer{})
}

// github.com/prysmaticlabs/go-ssz

func generateNoTimeNoStringNoFloatA() []*NoTimeNoStringNoFloatA {
	a := make([]*NoTimeNoStringNoFloatA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &NoTimeNoStringNoFloatA{
			Name:     []byte(randString(16)),
			BirthDay: uint64(time.Now().UnixNano()),
			Phone:    []byte(randString(10)),
			Siblings: uint32(rand.Intn(5)),
			Spouse:   rand.Intn(2) == 1,
			Money:    math.Float64bits(rand.Float64()),
		})
	}
	return a
}

func BenchmarkSSZNoTimeNoStringNoFloatAMarshal(b *testing.B) {
	data := generateNoTimeNoStringNoFloatA()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := ssz.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkSSZNoTimeNoStringNoFloatAUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateNoTimeNoStringNoFloatA()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		o, err := ssz.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		t := make([]byte, len(o))
		serialSize += copy(t, o)
		ser[i] = t
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &NoTimeNoStringNoFloatA{}
		err := ssz.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", "ssz", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := bytes.Equal(o.Name, i.Name) && bytes.Equal(o.Phone, i.Phone) && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/itsmontoya/mum
func BenchmarkMumMarshal(b *testing.B) {
	s := newMumSerializer()
	benchMarshal(b, s)
}

func BenchmarkMumUnmarshal(b *testing.B) {
	s := newMumSerializer()
	benchUnmarshal(b, s)
}

// github.com/200sc/bebop

func generateBebopA() []*BebopBufA {
	a := make([]*BebopBufA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &BebopBufA{
			Name: randString(16),
			// bebop does support times, but as 100-nanosecond ticks, losing some precision
			BirthDay: uint64(time.Now().UnixNano()),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkBebopaMarshal(b *testing.B) {
	data := generateBebopA()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		out := data[rand.Intn(len(data))].MarshalBebop()
		serialSize += len(out)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkBebopaUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateBebopA()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		ser[i] = d.MarshalBebop()
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	buf := new(bytes.Buffer)
	buf.Grow(100)
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := BebopBufA{}
		err := o.UnmarshalBebop(ser[n])
		if err != nil {
			b.Fatalf("bebop failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}
