package goserbench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"code.google.com/p/goprotobuf/proto"
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/binary"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/philhofer/msgp/msgp"
	"github.com/ugorji/go/codec"
	vmihailenco "github.com/vmihailenco/msgpack"
	vitessbson "github.com/youtube/vitess/go/bson"
	"gopkg.in/mgo.v2/bson"
)

var (
	validate = os.Getenv("VALIDATE")
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

func generateProto() []*ProtoBufA {
	a := make([]*ProtoBufA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &ProtoBufA{
			Name:     proto.String(randString(16)),
			BirthDay: proto.Int64(time.Now().Unix()),
			Phone:    proto.String(randString(10)),
			Siblings: proto.Int32(rand.Int31n(5)),
			Spouse:   proto.Bool(rand.Intn(2) == 1),
			Money:    proto.Float64(rand.Float64()),
		})
	}
	return a
}

type Serializer interface {
	Marshal(o interface{}) []byte
	Unmarshal(d []byte, o interface{}) error
	String() string
}

type MsgpSerializer struct{}

func (m MsgpSerializer) Marshal(o interface{}) []byte {
	bts, _ := o.(msgp.Marshaler).MarshalMsg()
	return bts
}

func (m MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(msgp.Unmarshaler).UnmarshalMsg(d)
	return err
}

func (m MsgpSerializer) String() string { return "Msgp" }

type VmihailencoMsgpackSerializer int

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) []byte {
	d, _ := vmihailenco.Marshal(o)
	return d
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func (m VmihailencoMsgpackSerializer) String() string {
	return "vmihailenco-msgpack"
}

type JsonSerializer int

func (m JsonSerializer) Marshal(o interface{}) []byte {
	d, _ := json.Marshal(o)
	return d
}

func (m JsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func (j JsonSerializer) String() string {
	return "json"
}

type BsonSerializer int

func (m BsonSerializer) Marshal(o interface{}) []byte {
	d, _ := bson.Marshal(o)
	return d
}

func (m BsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return bson.Unmarshal(d, o)
}

func (j BsonSerializer) String() string {
	return "bson"
}

type VitessBsonSerializer int

func (m VitessBsonSerializer) Marshal(o interface{}) []byte {
	d, _ := vitessbson.Marshal(o)
	return d
}

func (m VitessBsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return vitessbson.Unmarshal(d, o)
}

func (j VitessBsonSerializer) String() string {
	return "vitessbson"
}

type GobSerializer int

func (g GobSerializer) Marshal(o interface{}) []byte {
	b := &bytes.Buffer{}
	e := gob.NewEncoder(b)
	err := e.Encode(o)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func (g GobSerializer) Unmarshal(d []byte, o interface{}) error {
	b := bytes.NewBuffer(d)
	e := gob.NewDecoder(b)
	err := e.Decode(o)
	return err
}

func (g GobSerializer) String() string {
	return "gob"
}

type XdrSerializer int

func (x XdrSerializer) Marshal(o interface{}) []byte {
	d, _ := xdr.Marshal(o)
	return d
}

func (x XdrSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func (x XdrSerializer) String() string {
	return "xdr"
}

type ugorjiCodecSerializer struct {
	name string
	h    codec.Handle
}

func NewugorjiCodecSerializer(name string, h codec.Handle) *ugorjiCodecSerializer {
	return &ugorjiCodecSerializer{
		name: name,
		h:    h,
	}
}

func (u *ugorjiCodecSerializer) Marshal(o interface{}) []byte {
	var bs []byte
	codec.NewEncoderBytes(&bs, u.h).Encode(o)
	return bs
}

func (u *ugorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	return codec.NewDecoderBytes(d, u.h).Decode(o)
}

func (u *ugorjiCodecSerializer) String() string {
	return "ugorjicodec-" + u.name
}

type SerealSerializer int

func (s SerealSerializer) Marshal(o interface{}) []byte {
	d, _ := sereal.Marshal(o)
	return d
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func (s SerealSerializer) String() string {
	return "sereal"
}

type BinarySerializer int

func (b BinarySerializer) Marshal(o interface{}) []byte {
	d, _ := binary.Marshal(o)
	return d
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func (b BinarySerializer) String() string {
	return "binary"
}

func benchMarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	data := generate()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Marshal(data[rand.Intn(len(data))])
	}
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
	b.StopTimer()
	data := generate()
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i] = s.Marshal(d)
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String() //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
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

func BenchmarkVmihailencoMsgpackMarshal(b *testing.B) {
	benchMarshal(b, VmihailencoMsgpackSerializer(0))
}

func BenchmarkVmihailencoMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, VmihailencoMsgpackSerializer(0))
}

func BenchmarkJsonMarshal(b *testing.B) {
	benchMarshal(b, JsonSerializer(0))
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, JsonSerializer(0))
}

func BenchmarkBsonMarshal(b *testing.B) {
	benchMarshal(b, BsonSerializer(0))
}

func BenchmarkBsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, BsonSerializer(0))
}

func BenchmarkVitessBsonMarshal(b *testing.B) {
	benchMarshal(b, VitessBsonSerializer(0))
}

func BenchmarkVitessBsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, VitessBsonSerializer(0))
}

func BenchmarkGobMarshal(b *testing.B) {
	benchMarshal(b, GobSerializer(0))
}

func BenchmarkGobUnmarshal(b *testing.B) {
	benchUnmarshal(b, GobSerializer(0))
}

func BenchmarkXdrMarshal(b *testing.B) {
	benchMarshal(b, XdrSerializer(0))
}

func BenchmarkXdrUnmarshal(b *testing.B) {
	benchUnmarshal(b, XdrSerializer(0))
}

func BenchmarkugorjiCodecMsgpackMarshal(b *testing.B) {
	s := NewugorjiCodecSerializer("msgpack", &codec.MsgpackHandle{})
	benchMarshal(b, s)
}

func BenchmarkugorjiCodecMsgpackUnmarshal(b *testing.B) {
	s := NewugorjiCodecSerializer("msgpack", &codec.MsgpackHandle{})
	benchUnmarshal(b, s)
}

func BenchmarkugorjiCodecBincMarshal(b *testing.B) {
	s := NewugorjiCodecSerializer("binc", &codec.BincHandle{})
	benchMarshal(b, s)
}

func BenchmarkugorjiCodecBincUnmarshal(b *testing.B) {
	s := NewugorjiCodecSerializer("binc", &codec.BincHandle{})
	benchUnmarshal(b, s)
}

func BenchmarkSerealMarshal(b *testing.B) {
	benchMarshal(b, SerealSerializer(0))
}

func BenchmarkSerealUnmarshal(b *testing.B) {
	benchUnmarshal(b, SerealSerializer(0))
}

func BenchmarkBinaryMarshal(b *testing.B) {
	benchMarshal(b, BinarySerializer(0))
}

func BenchmarkBinaryUnmarshal(b *testing.B) {
	benchUnmarshal(b, BinarySerializer(0))
}

func BenchmarkMsgpMarshal(b *testing.B) {
	benchMarshal(b, MsgpSerializer{})
}

func BenchmarkMsgpUnmarshal(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{})
}

func BenchmarkGoprotobufMarshal(b *testing.B) {
	b.StopTimer()
	data := generateProto()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(data[rand.Intn(len(data))])
	}
}

func BenchmarkGoprotobufUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateProto()
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = proto.Marshal(d)
	}
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
