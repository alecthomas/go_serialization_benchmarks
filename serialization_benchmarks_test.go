package goserbench

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/avro"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/baseline"
	bebop200sc "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/bebop_200sc"
	bebopwellquite "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/bebop_wellquite"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/benc"
	binaryalecthomas "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/binary_alecthomas"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/bson"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/capnproto"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/colfer"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/easyjson"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/fastjson"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/flatbuffers"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/gencode"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/gogo"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/hprose"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/hprose2"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/ikea"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/jsoniter"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/mongobson"
	msgpacktinylib "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/msgpack_tinylib"
	msgpackvmihailenco "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/msgpack_vmihailenco"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/mus"
	protobufdedis "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/protobuf_dedis"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/pulsar"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/sereal"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/shamaton"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/ssz"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/stdlib"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/ugorji"
	xdrcalmh "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/xdr_calmh"
	xdrdavecgh "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/xdr_davecgh"
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
		New:  msgpacktinylib.NewMsgpSerializer,
	}, {
		Name: "msgpack",
		URL:  "github.com/vmihailenco/msgpack",
		New:  msgpackvmihailenco.NewVmihailencoMsgpackSerialier,
	}, {
		Name: "json",
		URL:  "pkg.go/dev/encoding/json",
		New:  stdlib.NewJSONSerializer,
	}, {
		Name: "jsoniter",
		URL:  "github.com/json-iterator/go",
		New:  jsoniter.NewJSONIterSerializer,
	}, {
		Name: "easyjson",
		URL:  "github.com/mailru/easyjson",
		New:  easyjson.NewEasyJSONSerializer,
	}, {
		Name: "bson",
		URL:  "gopkg.in/mgo.v2/bson",
		New:  bson.NewBsonSerializer,
	}, {
		Name: "mongobson",
		URL:  "go.mongodb.org/mongo-driver/mongo",
		New:  mongobson.NewMongoBSONSerializer,
	}, {
		Name: "gob",
		URL:  "pkg.go.dev/encoding/gob",
		New:  stdlib.NewGobSerializer,
	}, {
		Name: "davecgh/xdr",
		URL:  "github.com/davecgh/go-xdr/xdr",
		New:  xdrdavecgh.NewXDRDavecghSerializer,
	}, {
		Name: "ugorji/msgpack",
		URL:  "github.com/ugorji/go/codec",
		New:  ugorji.NewUgorjiCodecMsgPack,
	}, {
		Name: "ugorji/binc",
		URL:  "github.com/ugorji/go/codec",
		New:  ugorji.NewUgorjiCodecBinc,
	}, {
		Name: "sereal",
		URL:  "github.com/Sereal/Sereal/Go/sereal",
		New:  sereal.NewSerealSerializer,
	}, {
		Name: "alecthomas/binary",
		URL:  "github.com/alecthomas/binary",
		New:  binaryalecthomas.NewBinarySerializer,
	}, {
		Name: "flatbuffers",
		URL:  "github.com/google/flatbuffers/go",
		New:  flatbuffers.NewFlatBuffersSerializer,
	}, {
		Name: "capnproto",
		URL:  "github.com/glycerine/go-capnproto",
		New:  capnproto.NewCapNProtoSerializer,
	}, {
		Name: "hprose",
		URL:  "github.com/hprose/hprose-go/io",
		New:  hprose.NewHproseSerializer,
	}, {
		Name: "hprose2",
		URL:  "github.com/hprose/hprose-golang/io",
		New:  hprose2.NewHProse2Serializer,
	}, {
		Name: "dedis/protobuf",
		URL:  "go.dedis.ch/protobuf",
		New:  protobufdedis.NewProtobufSerializer,
	}, {
		Name: "pulsar",
		URL:  "github.com/cosmos/cosmos-proto",
		New:  pulsar.NewPulsarSerializer,
	}, {
		Name: "gogo/protobuf",
		URL:  "github.com/gogo/protobuf/proto",
		New:  gogo.NewGogoProtoSerializer,
	}, {
		Name: "gogo/jsonpb",
		URL:  "github.com/gogo/protobuf/proto",
		New:  gogo.NewGogoJsonSerializer,
	}, {
		Name: "colfer",
		URL:  "github.com/pascaldekloe/colfer",
		New:  colfer.NewColferSerializer,
	}, {
		Name: "gencode",
		URL:  "github.com/andyleap/gencode",
		New:  gencode.NewGencodeSerializer,
	}, {
		Name: "gencode/unsafe",
		URL:  "github.com/andyleap/gencode",
		New:  gencode.NewGencodeUnsafeSerializer,
	}, {
		Name: "calmh/xdr",
		URL:  "github.com/calmh/xdr",
		New:  xdrcalmh.NewXDRCalmhSerializer,
	}, {
		Name: "goavro",
		URL:  "gopkg.in/linkedin/goavro.v1",
		New:  avro.NewAvroA,
	}, {
		Name: "avro2/text",
		URL:  "github.com/linkedin/goavro",
		New:  avro.NewAvro2Txt,
	}, {
		Name: "avro2/binary",
		URL:  "github.com/linkedin/goavro",
		New:  avro.NewAvro2Bin,
	}, {
		Name: "ikea",
		URL:  "github.com/ikkerens/ikeapack",
		New:  ikea.NewIkeaSerializer,
	}, {
		Name: "shamaton/msgpack/map",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonMapMsgpackSerializer,
	}, {
		Name: "shamaton/msgpack/array",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonArrayMsgPackSerializer,
	}, {
		Name: "shamaton/msgpackgen/map",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonMapMsgPackgenSerializer,
	}, {
		Name: "shamaton/msgpackgen/array",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonArrayMsgpackgenSerializer,
	}, {
		Name: "ssz",
		URL:  "github.com/prysmaticlabs/go-ssz",
		New:  ssz.NewSSZSerializer,
	}, {
		Name: "200sc/bebop",
		URL:  "github.com/200sc/bebop",
		New:  bebop200sc.NewBebop200ScSerializer,
	}, {
		Name: "wellquite/bebop",
		URL:  "wellquite.org/bebop",
		New:  bebopwellquite.NewBebopWellquiteSerializer,
	}, {
		Name: "fastjson",
		URL:  "github.com/valyala/fastjson",
		New:  fastjson.NewFastJSONSerializer,
	}, {
		Name: "benc",
		URL:  "github.com/deneonet/benc",
		New:  benc.NewBENCSerializer,
	}, {
		Name: "benc/usafe",
		URL:  "github.com/deneonet/benc",
		New:  benc.NewBENCUnsafeSerializer,
	}, {
		Name: "mus",
		URL:  "github.com/mus-format/mus-go",
		New:  mus.NewMUSSerializer,
	}, {
		Name: "mus/unsafe",
		URL:  "github.com/mus-format/mus-go",
		New:  mus.NewMUSUnsafeSerializer,
	}, {
		Name: "baseline",
		URL:  "",
		New:  baseline.NewBaselineSerializer,
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
