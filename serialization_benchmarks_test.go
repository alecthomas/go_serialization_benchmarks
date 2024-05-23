package goserbench

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

var (
	validate = os.Getenv("VALIDATE") != ""
)

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
			goserbench.BenchMarshalSmallStruct(b, bc.New())
		})
		b.Run("unmarshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchUnmarshalSmallStruct(b, bc.New(), validate)
		})
	}
}
