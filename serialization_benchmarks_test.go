package goserbench

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/niubaoshu/gotiny"
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

type Serializer interface {
	Marshal(o interface{}) ([]byte, error)
	Unmarshal(d []byte, o interface{}) error
}

// SerializerTimePrecision is a serializer that specifies the max precision that
// a time.Time is encodable. This will be used to truncate time fields.
type SerializerTimePrecision interface {
	// TimePrecision is the max precision a time.Time may be encoded.  When
	// greater than zero, all time.Time fields are truncated down to this
	// precision.
	TimePrecision() time.Duration
}

// SerializerEnforcesTimezone is a serializer that enforces a specific timezone
// when marshalling/unmarshalling time.Time fields.
type SerializerEnforcesTimezone interface {
	// ForcesUTC is true when the serializes forces a UTC timezone.
	ForcesUTC() bool
}

// SerializerLimitsFloat64Precision is a serializer that enforces a maximum
// precision when marshalling/unmarshalling float64 fields.
type SerializerLimitsFloat64Precision interface {
	// FractionalDigits returns the max number of fractional digits that
	// the serializer may encode.
	ReduceFloat64Precision() uint
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

	var timePrecision time.Duration
	if stp, ok := s.(SerializerTimePrecision); ok {
		timePrecision = stp.TimePrecision()
	}
	var forcesUTC bool
	if set, ok := s.(SerializerEnforcesTimezone); ok {
		forcesUTC = set.ForcesUTC()
	}

	data := generate()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		// Reduce the precision of the Birthday field when the
		// serializer cannot represent time with nanosecond precision.
		if timePrecision > 0 {
			d.BirthDay = d.BirthDay.Truncate(timePrecision)
		}

		// Enforce Timezone when serializer requires it.
		if forcesUTC {
			d.BirthDay = d.BirthDay.UTC()
		}

		// Reduce precision of fractional fields when the serializer
		// cannot represent the full float64 range.
		if slfp, ok := s.(SerializerLimitsFloat64Precision); ok {
			fracDigits := slfp.ReduceFloat64Precision()
			i, f := math.Modf(d.Money)
			power := math.Pow(10, float64(fracDigits))
			newf := math.Trunc(f*power) / power
			d.Money = float64(i) + newf
		}

		o, err := s.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		t := make([]byte, len(o))
		serialSize += copy(t, o)
		ser[i] = t
	}
	o := &A{}

	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		*o = A{}
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
	fmt.Print(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

`)
}

// github.com/niubaoshu/gotiny

type GotinySerializer struct {
	dec *gotiny.Decoder
}

func (g GotinySerializer) Marshal(o interface{}) ([]byte, error) {
	return gotiny.Marshal(o), nil
}

func (g GotinySerializer) Unmarshal(d []byte, o interface{}) error {
	g.dec.Decode(d, o)
	return nil
}

func NewGotinySerializer(o interface{}) Serializer {
	ot := reflect.TypeOf(o)
	return GotinySerializer{
		dec: gotiny.NewDecoderWithType(ot),
	}
}

func Benchmark_Gotiny_Marshal(b *testing.B) {
	benchMarshal(b, NewGotinySerializer(A{}))
}

func Benchmark_Gotiny_Unmarshal(b *testing.B) {
	benchUnmarshal(b, NewGotinySerializer(A{}))
}

// github.com/tinylib/msgp

func Benchmark_Msgp_Marshal(b *testing.B) {
	benchMarshal(b, newMsgpSerializer())
}

func Benchmark_Msgp_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newMsgpSerializer())
}

// github.com/vmihailenco/msgpack

func Benchmark_VmihailencoMsgpack_Marshal(b *testing.B) {
	benchMarshal(b, newVmihailencoMsgpackSerialier())
}

func Benchmark_VmihailencoMsgpack_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newVmihailencoMsgpackSerialier())
}

// encoding/json

func Benchmark_Json_Marshal(b *testing.B) {
	benchMarshal(b, newJSONSerializer())
}

func Benchmark_Json_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newJSONSerializer())
}

// github.com/json-iterator/go

func Benchmark_JsonIter_Marshal(b *testing.B) {
	benchMarshal(b, newJSONIterSerializer())
}

func Benchmark_JsonIter_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newJSONIterSerializer())
}

// github.com/mailru/easyjson

func Benchmark_EasyJson_Marshal(b *testing.B) {
	benchMarshal(b, newEasyJSONSerializer())
}

func Benchmark_EasyJson_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newEasyJSONSerializer())
}

// gopkg.in/mgo.v2/bson

func Benchmark_Bson_Marshal(b *testing.B) {
	benchMarshal(b, newBsonSerializer())
}

func Benchmark_Bson_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newBsonSerializer())
}

// go.mongodb.org/mongo-driver/mongo

func Benchmark_MongoBson_Marshal(b *testing.B) {
	benchMarshal(b, newMongoBSONSerializer())
}

func Benchmark_MongoBson_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newMongoBSONSerializer())
}

// encoding/gob

func Benchmark_Gob_Marshal(b *testing.B) {
	benchMarshal(b, NewGobSerializer())
}

func Benchmark_Gob_Unmarshal(b *testing.B) {
	benchUnmarshal(b, NewGobSerializer())
}

// github.com/davecgh/go-xdr/xdr

func Benchmark_XDRDavecgh_Marshal(b *testing.B) {
	benchMarshal(b, newXDRDavecghSerializer())
}

func Benchmark_XDRDavecgh_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newXDRDavecghSerializer())
}

// github.com/ugorji/go/codec

func Benchmark_UgorjiCodecMsgpack_Marshal(b *testing.B) {
	benchMarshal(b, newUgorjiCodecMsgPack())
}

func Benchmark_UgorjiCodecMsgpack_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newUgorjiCodecMsgPack())
}

func Benchmark_UgorjiCodecBinc_Marshal(b *testing.B) {
	benchMarshal(b, newUgorjiCodecBinc())
}

func Benchmark_UgorjiCodecBinc_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newUgorjiCodecBinc())
}

// github.com/Sereal/Sereal/Go/sereal

func Benchmark_Sereal_Marshal(b *testing.B) {
	benchMarshal(b, newSerealSerializer())
}

func Benchmark_Sereal_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newSerealSerializer())
}

// github.com/alecthomas/binary

func Benchmark_Binary_Marshal(b *testing.B) {
	benchMarshal(b, newBinarySerializer())
}

func Benchmark_Binary_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newBinarySerializer())
}

// github.com/google/flatbuffers/go

func Benchmark_FlatBuffers_Marshal(b *testing.B) {
	benchMarshal(b, newFlatBuffersSerializer())
}

func Benchmark_FlatBuffers_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newFlatBuffersSerializer())
}

// github.com/glycerine/go-capnproto

func Benchmark_CapNProto_Marshal(b *testing.B) {
	benchMarshal(b, newCapNProtoSerializer())
}

func Benchmark_CapNProto_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newCapNProtoSerializer())
}

// github.com/hprose/hprose-go/io

func Benchmark_Hprose_Marshal(b *testing.B) {
	benchMarshal(b, newHproseSerializer())
}

func Benchmark_Hprose_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newHproseSerializer())
}

// github.com/hprose/hprose-golang/io

func Benchmark_Hprose2_Marshal(b *testing.B) {
	benchMarshal(b, newHProse2Serializer())
}

func Benchmark_Hprose2_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newHProse2Serializer())
}

// go.dedis.ch/protobuf

func Benchmark_Protobuf_Marshal(b *testing.B) {
	benchMarshal(b, newProtobufSerializer())
}

func Benchmark_Protobuf_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newProtobufSerializer())
}

// github.com/cosmos/cosmos-proto

func Benchmark_Pulsar_Marshal(b *testing.B) {
	benchMarshal(b, newPulsarSerializer())
}

func Benchmark_Pulsar_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newPulsarSerializer())
}

// github.com/gogo/protobuf/proto

func Benchmark_Gogoprotobuf_Marshal(b *testing.B) {
	benchMarshal(b, newGogoProtoSerializer())
}

func Benchmark_Gogoprotobuf_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newGogoProtoSerializer())
}

func Benchmark_Gogojsonpb_Marshal(b *testing.B) {
	benchMarshal(b, newGogoJsonSerializer())
}

func Benchmark_Gogojsonpb_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newGogoJsonSerializer())
}

// github.com/pascaldekloe/colfer

func Benchmark_Colfer_Marshal(b *testing.B) {
	benchMarshal(b, newColferSerializer())
}

func Benchmark_Colfer_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newColferSerializer())
}

// github.com/andyleap/gencode

func Benchmark_Gencode_Marshal(b *testing.B) {
	benchMarshal(b, newGencodeSerializer())
}

func Benchmark_Gencode_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newGencodeSerializer())
}

func Benchmark_GencodeUnsafe_Marshal(b *testing.B) {
	benchMarshal(b, newGencodeUnsafeSerializer())
}

func Benchmark_GencodeUnsafe_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newGencodeUnsafeSerializer())
}

// github.com/calmh/xdr

func Benchmark_XDRCalmh_Marshal(b *testing.B) {
	benchMarshal(b, newXDRCalmhSerializer())
}

func Benchmark_XDRCalmh_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newXDRCalmhSerializer())
}

// gopkg.in/linkedin/goavro.v1

func Benchmark_GoAvro_Marshal(b *testing.B) {
	benchMarshal(b, NewAvroA())
}

func Benchmark_GoAvro_Unmarshal(b *testing.B) {
	benchUnmarshal(b, NewAvroA())
}

// github.com/linkedin/goavro

func Benchmark_GoAvro2Text_Marshal(b *testing.B) {
	benchMarshal(b, NewAvro2Txt())
}

func Benchmark_GoAvro2Text_Unmarshal(b *testing.B) {
	benchUnmarshal(b, NewAvro2Txt())
}

func Benchmark_GoAvro2Binary_Marshal(b *testing.B) {
	benchMarshal(b, NewAvro2Bin())
}

func Benchmark_GoAvro2Binary_Unmarshal(b *testing.B) {
	benchUnmarshal(b, NewAvro2Bin())
}

// github.com/ikkerens/ikeapack

func Benchmark_Ikea_Marshal(b *testing.B) {
	benchMarshal(b, newIkeSerializer())
}

func Benchmark_Ikea_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newIkeSerializer())
}

// github.com/shamaton/msgpack - as map

func Benchmark_ShamatonMapMsgpack_Marshal(b *testing.B) {
	benchMarshal(b, newShamatonMapMsgpackSerializer())
}

func Benchmark_ShamatonMapMsgpack_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newShamatonMapMsgpackSerializer())
}

// github.com/shamaton/msgpack - as array

func Benchmark_ShamatonArrayMsgpack_Marshal(b *testing.B) {
	benchMarshal(b, newShamatonArrayMsgPackSerializer())
}

func Benchmark_ShamatonArrayMsgpack_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newShamatonArrayMsgPackSerializer())
}

// github.com/shamaton/msgpackgen - as map

func Benchmark_ShamatonMapMsgpackgen_Marshal(b *testing.B) {
	benchMarshal(b, newShamatonMapMsgPackgenSerializer())
}

func Benchmark_ShamatonMapMsgpackgen_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newShamatonMapMsgPackgenSerializer())
}

// github.com/shamaton/msgpackgen - as array

func Benchmark_ShamatonArrayMsgpackgen_Marshal(b *testing.B) {
	benchMarshal(b, newShamatonArrayMsgpackgenSerializer())
}

func Benchmark_ShamatonArrayMsgpackgen_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newShamatonArrayMsgpackgenSerializer())
}

// github.com/prysmaticlabs/go-ssz

func Benchmark_SSZ_Marshal(b *testing.B) {
	benchMarshal(b, newSSZSerializer())
}

func Benchmark_SSZ_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newSSZSerializer())
}

// github.com/200sc/bebop

func Benchmark_Bebop200sc_Marshal(b *testing.B) {
	benchMarshal(b, newBebop200ScSerializer())
}

func Benchmark_Bebop200sc_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newBebop200ScSerializer())
}

// wellquite.org/bebop

func Benchmark_BebopWellquite_Marshal(b *testing.B) {
	benchMarshal(b, newBebopWellquiteSerializer())
}

func Benchmark_BebopWellquite_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newBebopWellquiteSerializer())
}

// github.com/valyala/fastjson

func Benchmark_FastJson_Marshal(b *testing.B) {
	benchMarshal(b, newFastJSONSerializer())
}

func Benchmark_FastJson_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newFastJSONSerializer())
}

// github.com/deneonet/benc

func Benchmark_BENC_Marshal(b *testing.B) {
	benchMarshal(b, BENCSerializer{})
}
func Benchmark_BENC_Unmarshal(b *testing.B) {
	benchUnmarshal(b, BENCSerializer{})
}

func Benchmark_BENCUnsafe_Marshal(b *testing.B) {
	benchMarshal(b, BENCUnsafeSerializer{})
}

func Benchmark_BENCUnsafe_Unmarshal(b *testing.B) {
	benchUnmarshal(b, BENCUnsafeSerializer{})
}

// github.com/mus-format/mus-go

func Benchmark_MUS_Marshal(b *testing.B) {
	benchMarshal(b, MUSSerializer{})
}

func Benchmark_MUS_Unmarshal(b *testing.B) {
	benchUnmarshal(b, MUSSerializer{})
}

func Benchmark_MUSUnsafe_Marshal(b *testing.B) {
	benchMarshal(b, MUSUnsafeSerializer{})
}

func Benchmark_MUSUnsafe_Unmarshal(b *testing.B) {
	benchUnmarshal(b, MUSUnsafeSerializer{})
}

func Benchmark_Baseline_Marshal(b *testing.B) {
	benchMarshal(b, newBaselineSerializer())
}

func Benchmark_Baseline_Unmarshal(b *testing.B) {
	benchUnmarshal(b, newBaselineSerializer())
}
