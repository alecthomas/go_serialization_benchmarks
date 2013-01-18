package goserbench

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/ugorji/go-msgpack"
	"labix.org/v2/mgo/bson"
	"math/rand"
	"testing"
	"time"
)

type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
}

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
		})
	}
	return a
}

type Serializer interface {
	Marshal(o interface{}) []byte
	Unmarshal(d []byte, o interface{})
	String() string
}

type MsgpackSerializer int

func (m MsgpackSerializer) Marshal(o interface{}) []byte {
	d, _ := msgpack.Marshal(o)
	return d
}

func (m MsgpackSerializer) Unmarshal(d []byte, o interface{}) {
	msgpack.Unmarshal(d, o, nil)
}

func (m MsgpackSerializer) String() string {
	return "msgpack"
}

type JsonSerializer int

func (m JsonSerializer) Marshal(o interface{}) []byte {
	d, _ := json.Marshal(o)
	return d
}

func (m JsonSerializer) Unmarshal(d []byte, o interface{}) {
	json.Unmarshal(d, o)
}

func (j JsonSerializer) String() string {
	return "json"
}

type XmlSerializer int

func (m XmlSerializer) Marshal(o interface{}) []byte {
	d, _ := xml.Marshal(o)
	return d
}

func (m XmlSerializer) Unmarshal(d []byte, o interface{}) {
	xml.Unmarshal(d, o)
}

func (j XmlSerializer) String() string {
	return "xml"
}

type BsonSerializer int

func (m BsonSerializer) Marshal(o interface{}) []byte {
	d, _ := bson.Marshal(o)
	return d
}

func (m BsonSerializer) Unmarshal(d []byte, o interface{}) {
	bson.Unmarshal(d, o)
}

func (j BsonSerializer) String() string {
	return "bson"
}

func benchMarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	data := generate()
	b.StartTimer()
	var size uint64 = 0
	for i := 0; i < b.N; i++ {
		b := s.Marshal(data[rand.Intn(len(data))])
		size += uint64(len(b))
	}
	// fmt.Fprintf(os.Stderr, "average size of %s serialized structure is %d\n", s, size/uint64(b.N))
}

func benchUnmarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	data := generate()
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i] = s.Marshal(d)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		s.Unmarshal(ser[n], o)
	}
}

func TestMessage(t *testing.T) {
	println(`
A test suite for benchmarking various Go serialization methods.
To run:

     go get github.com/ugorji/go-msgpack labix.org/v2/mgo/bson
     go test -bench='.*' ./

`)

}

func BenchmarkMsgpackMarshal(b *testing.B) {
	benchMarshal(b, MsgpackSerializer(0))
}

func BenchmarkMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, MsgpackSerializer(0))
}

func BenchmarkJsonMarshal(b *testing.B) {
	benchMarshal(b, JsonSerializer(0))
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, JsonSerializer(0))
}

func BenchmarkXmlMarshal(b *testing.B) {
	benchMarshal(b, XmlSerializer(0))
}

func BenchmarkXmlUnmarshal(b *testing.B) {
	benchUnmarshal(b, XmlSerializer(0))
}

func BenchmarkBsonMarshal(b *testing.B) {
	benchMarshal(b, BsonSerializer(0))
}

func BenchmarkBsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, BsonSerializer(0))
}
