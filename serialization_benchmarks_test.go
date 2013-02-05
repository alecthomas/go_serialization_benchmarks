package goserbench

import (
	vitessbson "code.google.com/p/vitess/go/bson"
	"encoding/json"
	"fmt"
	ugorji "github.com/ugorji/go-msgpack"
	vmihailenco "github.com/vmihailenco/msgpack"
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
	Money    float64
	Tags     map[string]string
	Aliases  []string
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
			Money:    rand.Float64(),
			Tags:     map[string]string{"a": "a", "b": "b", "c": "c"},
			Aliases:  []string{randString(5), randString(5)},
		})
	}
	return a
}

type Serializer interface {
	Marshal(o interface{}) []byte
	Unmarshal(d []byte, o interface{}) error
	String() string
}

type UgorjiMsgpackSerializer int

func (m UgorjiMsgpackSerializer) Marshal(o interface{}) []byte {
	d, _ := ugorji.Marshal(o)
	return d
}

func (m UgorjiMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return ugorji.Unmarshal(d, o, nil)
}

func (m UgorjiMsgpackSerializer) String() string {
	return "ugorji-msgpack"
}

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

func benchMarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	data := generate()
	b.StartTimer()
	var size uint64 = 0
	for i := 0; i < b.N; i++ {
		b := s.Marshal(data[rand.Intn(len(data))])
		size += uint64(len(b))
	}
	// println(fmt.Sprintf("average size of %s serialized structure is %d", s, size/uint64(b.N)))
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
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s", s, err)
		}
		// Validate unmarshalled data.
		// i := data[n]
		// // o.BirthDay.String() == i.BirthDay.String() &&
		// correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
		// if !correct {
		// 	b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
		// }
	}
}

func TestMessage(t *testing.T) {
	println(`
A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.
`)

}
func BenchmarkUgorjiMsgpackMarshal(b *testing.B) {
	benchMarshal(b, UgorjiMsgpackSerializer(0))
}

func BenchmarkUgorjiMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, UgorjiMsgpackSerializer(0))
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
