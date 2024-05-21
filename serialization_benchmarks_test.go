package goserbench

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
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

type BenchmarkCase struct {
	Name string
	URL  string
	New  func() Serializer
}

var benchmarkCases = []BenchmarkCase{
	{
		Name: "gotiny",
		URL:  "github.com/niubaoshu/gotiny",
		New:  NewGotinySerializer,
	}, {
		Name: "msgp",
		URL:  "github.com/tinylib/msgp",
		New:  NewMsgpSerializer,
	}, {
		Name: "msgpack",
		URL:  "github.com/vmihailenco/msgpack",
		New:  NewVmihailencoMsgpackSerialier,
	}, {
		Name: "json",
		URL:  "pkg.go/dev/encoding/json",
		New:  NewJSONSerializer,
	}, {
		Name: "jsoniter",
		URL:  "github.com/json-iterator/go",
		New:  NewJSONIterSerializer,
	}, {
		Name: "easyjson",
		URL:  "github.com/mailru/easyjson",
		New:  NewEasyJSONSerializer,
	}, {
		Name: "bson",
		URL:  "gopkg.in/mgo.v2/bson",
		New:  NewBsonSerializer,
	}, {
		Name: "mongobson",
		URL:  "go.mongodb.org/mongo-driver/mongo",
		New:  NewMongoBSONSerializer,
	}, {
		Name: "gob",
		URL:  "pkg.go.dev/encoding/gob",
		New:  NewGobSerializer,
	}, {
		Name: "davecgh/xdr",
		URL:  "github.com/davecgh/go-xdr/xdr",
		New:  NewXDRDavecghSerializer,
	}, {
		Name: "ugorji/msgpack",
		URL:  "github.com/ugorji/go/codec",
		New:  NewUgorjiCodecMsgPack,
	}, {
		Name: "ugorji/binc",
		URL:  "github.com/ugorji/go/codec",
		New:  NewUgorjiCodecBinc,
	}, {
		Name: "sereal",
		URL:  "github.com/Sereal/Sereal/Go/sereal",
		New:  NewSerealSerializer,
	}, {
		Name: "alecthomas/binary",
		URL:  "github.com/alecthomas/binary",
		New:  NewBinarySerializer,
	}, {
		Name: "flatbuffers",
		URL:  "github.com/google/flatbuffers/go",
		New:  NewFlatBuffersSerializer,
	}, {
		Name: "capnproto",
		URL:  "github.com/glycerine/go-capnproto",
		New:  NewCapNProtoSerializer,
	}, {
		Name: "hprose",
		URL:  "github.com/hprose/hprose-go/io",
		New:  NewHproseSerializer,
	}, {
		Name: "hprose2",
		URL:  "github.com/hprose/hprose-golang/io",
		New:  NewHProse2Serializer,
	}, {
		Name: "dedis/protobuf",
		URL:  "go.dedis.ch/protobuf",
		New:  NewProtobufSerializer,
	}, {
		Name: "pulsar",
		URL:  "github.com/cosmos/cosmos-proto",
		New:  NewPulsarSerializer,
	}, {
		Name: "gogo/protobuf",
		URL:  "github.com/gogo/protobuf/proto",
		New:  NewGogoProtoSerializer,
	}, {
		Name: "gogo/jsonpb",
		URL:  "github.com/gogo/protobuf/proto",
		New:  NewGogoJsonSerializer,
	}, {
		Name: "colfer",
		URL:  "github.com/pascaldekloe/colfer",
		New:  NewColferSerializer,
	}, {
		Name: "gencode",
		URL:  "github.com/andyleap/gencode",
		New:  NewGencodeSerializer,
	}, {
		Name: "gencode/unsafe",
		URL:  "github.com/andyleap/gencode",
		New:  NewGencodeUnsafeSerializer,
	}, {
		Name: "calmh/xdr",
		URL:  "github.com/calmh/xdr",
		New:  NewXDRCalmhSerializer,
	}, {
		Name: "goavro",
		URL:  "gopkg.in/linkedin/goavro.v1",
		New:  NewAvroA,
	}, {
		Name: "avro2/text",
		URL:  "github.com/linkedin/goavro",
		New:  NewAvro2Txt,
	}, {
		Name: "avro2/binary",
		URL:  "github.com/linkedin/goavro",
		New:  NewAvro2Bin,
	}, {
		Name: "ikea",
		URL:  "github.com/ikkerens/ikeapack",
		New:  NewIkeSerializer,
	}, {
		Name: "shamaton/msgpack/map",
		URL:  "github.com/shamaton/msgpack",
		New:  NewShamatonMapMsgpackSerializer,
	}, {
		Name: "shamaton/msgpack/array",
		URL:  "github.com/shamaton/msgpack",
		New:  NewShamatonArrayMsgPackSerializer,
	}, {
		Name: "shamaton/msgpackgen/map",
		URL:  "github.com/shamaton/msgpack",
		New:  NewShamatonMapMsgPackgenSerializer,
	}, {
		Name: "shamaton/msgpackgen/array",
		URL:  "github.com/shamaton/msgpack",
		New:  NewShamatonArrayMsgpackgenSerializer,
	}, {
		Name: "ssz",
		URL:  "github.com/prysmaticlabs/go-ssz",
		New:  NewSSZSerializer,
	}, {
		Name: "200sc/bebop",
		URL:  "github.com/200sc/bebop",
		New:  NewBebop200ScSerializer,
	}, {
		Name: "wellquite/bebop",
		URL:  "wellquite.org/bebop",
		New:  NewBebopWellquiteSerializer,
	}, {
		Name: "fastjson",
		URL:  "github.com/valyala/fastjson",
		New:  NewFastJSONSerializer,
	}, {
		Name: "benc",
		URL:  "github.com/deneonet/benc",
		New:  NewBENCSerializer,
	}, {
		Name: "benc/usafe",
		URL:  "github.com/deneonet/benc",
		New:  NewBENCUnsafeSerializer,
	}, {
		Name: "mus",
		URL:  "github.com/mus-format/mus-go",
		New:  NewMUSSerializer,
	}, {
		Name: "mus/unsafe",
		URL:  "github.com/mus-format/mus-go",
		New:  NewMUSUnsafeSerializer,
	}, {
		Name: "baseline",
		URL:  "",
		New:  NewBaselineSerializer,
	},
}

func BenchmarkSerializers(b *testing.B) {
	for i := range benchmarkCases {
		bc := benchmarkCases[i]
		b.Run("marshal/"+bc.Name, func(b *testing.B) {
			benchMarshal(b, bc.New())
		})
		b.Run("unmarshal/"+bc.Name, func(b *testing.B) {
			benchUnmarshal(b, bc.New())
		})
	}
}
