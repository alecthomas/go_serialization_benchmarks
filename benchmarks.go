package goserbench

import (
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
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/fastape"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/fastjson"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/flatbuffers"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/gencode"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/gogo"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/gotiny"
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

// TimeSupport is the type of support for time.Time values.
type TimeSupport string

const (
	// TSUnixNs means the serializer supports time by encoding into an int64
	// nanoseconds since unix.
	TSUnixNs TimeSupport = "unix-ns"

	// TSUnixMs means the serializer supports time by encoding into an int64
	// milliseconds since unix.
	TSUnixMs TimeSupport = "unix-ms"

	// TSFullRange means the serializer supports the full time.Time range,
	// but without timezone information.
	TSFullRange TimeSupport = "fullrange"

	// TSFullTzOffset means the serializer supports the full time.Time range,
	// including timezone offset information (but not timezone names).
	TSFullTzOffset TimeSupport = "fulltzoffset"

	// TSRFC3339Ns means the serializer supports encoding time values by
	// using RFC3339 strings with nanosecond precision.
	TSRFC3339Ns TimeSupport = "RFC3339ns"

	// TSCustom means the serializer has custom restrictions on time.Time
	// fields.
	TSCustom TimeSupport = "custom"

	// TSNoSupport means the serializer does not natively encode time.Time
	// fields. Such fields need to be converted to an appropriately
	// supported primitive field by the user.
	TSNoSupport TimeSupport = "no"

	// TSUnknown means the serializer has not documented the encoding or
	// the limitations on encoding time.Time values.
	TSUnknown TimeSupport = "unknown"
)

// APIKind is the type of API needed to interact with the serializer.
type APIKind string

const (
	// AKReflect means the serializer uses Go's reflect API to marshal/
	// unmarshal objects.
	AKReflect APIKind = "reflect"

	// AKCodegen means there's a code generation step that is used to
	// generate the client code used for marshalling/unmarshalling.
	AKCodegen APIKind = "codegen"

	// AKManual means the user must manually write the code to
	// marshal/unmarshal structures.
	AKManual APIKind = "manual"
)

type BenchmarkCase struct {
	Name string
	URL  string
	New  func() goserbench.Serializer

	// UnsafeStringUnmarshal records whether the serializer uses usafe
	// strings during unmarshal. Unsafe strings are modified if the
	// underlying byte slice is modified.
	UnsafeStringUnmarshal bool

	// BufferReuseMarshal reocrds whether the serializer re-uses the
	// marshalling buffer.
	BufferReuseMarshal bool

	// TimeSupport records the type of support time.Time values have on
	// this encoder.
	TimeSupport TimeSupport

	// APIKind records the type of user code needed to marshal/unmarshal
	// values with the serializer.
	APIKind APIKind

	// Notes are notes about specific limitations for the serializer.
	Notes []string
}

var benchmarkCases = []BenchmarkCase{
	{
		Name: "gotiny",
		URL:  "github.com/niubaoshu/gotiny",
		New:  gotiny.NewGotinySerializer,

		TimeSupport: TSUnixNs,
		APIKind:     AKReflect,
	}, {
		Name: "msgp",
		URL:  "github.com/tinylib/msgp",
		New:  msgpacktinylib.NewMsgpSerializer,

		TimeSupport: TSFullRange,
		APIKind:     AKCodegen,
	}, {
		Name: "msgpack",
		URL:  "github.com/vmihailenco/msgpack",
		New:  msgpackvmihailenco.NewVmihailencoMsgpackSerialier,

		TimeSupport: TSFullRange,
		APIKind:     AKReflect,
	}, {
		Name: "json",
		URL:  "pkg.go/dev/encoding/json",
		New:  stdlib.NewJSONSerializer,

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKReflect,
	}, {
		Name: "jsoniter",
		URL:  "github.com/json-iterator/go",
		New:  jsoniter.NewJSONIterSerializer,

		TimeSupport: TSCustom,
		APIKind:     AKReflect,
	}, {
		Name: "easyjson",
		URL:  "github.com/mailru/easyjson",
		New:  easyjson.NewEasyJSONSerializer,

		TimeSupport: TSUnknown,
		APIKind:     AKCodegen,
	}, {
		Name: "bson",
		URL:  "gopkg.in/mgo.v2/bson",
		New:  bson.NewBsonSerializer,

		TimeSupport: TSUnixMs,
		APIKind:     AKReflect,
	}, {
		Name: "mongobson",
		URL:  "go.mongodb.org/mongo-driver/mongo",
		New:  mongobson.NewMongoBSONSerializer,

		TimeSupport: TSUnixMs,
		APIKind:     AKReflect,
	}, {
		Name: "gob",
		URL:  "pkg.go.dev/encoding/gob",
		New:  stdlib.NewGobSerializer,

		TimeSupport: TSFullTzOffset,
		APIKind:     AKReflect,
	}, {
		Name: "davecgh/xdr",
		URL:  "github.com/davecgh/go-xdr/xdr",
		New:  xdrdavecgh.NewXDRDavecghSerializer,

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKReflect,
	}, {
		Name: "ugorji/msgpack",
		URL:  "github.com/ugorji/go/codec",
		New:  ugorji.NewUgorjiCodecMsgPack,

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "ugorji/binc",
		URL:  "github.com/ugorji/go/codec",
		New:  ugorji.NewUgorjiCodecBinc,

		TimeSupport: TSFullTzOffset,
		APIKind:     AKReflect,
	}, {
		Name: "sereal",
		URL:  "github.com/Sereal/Sereal/Go/sereal",
		New:  sereal.NewSerealSerializer,

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "alecthomas/binary",
		URL:  "github.com/alecthomas/binary",
		New:  binaryalecthomas.NewBinarySerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKReflect,
	}, {
		Name: "flatbuffers",
		URL:  "github.com/google/flatbuffers/go",
		New:  flatbuffers.NewFlatBuffersSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "capnproto",
		URL:  "github.com/glycerine/go-capnproto",
		New:  capnproto.NewCapNProtoSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "hprose",
		URL:  "github.com/hprose/hprose-go/io",
		New:  hprose.NewHproseSerializer,

		TimeSupport: TSCustom,
		APIKind:     AKManual,
	}, {
		Name: "hprose2",
		URL:  "github.com/hprose/hprose-golang/io",
		New:  hprose2.NewHProse2Serializer,

		TimeSupport: TSCustom,
		APIKind:     AKManual,
	}, {
		Name: "dedis/protobuf",
		URL:  "go.dedis.ch/protobuf",
		New:  protobufdedis.NewProtobufSerializer,

		TimeSupport: TSUnixNs,
		APIKind:     AKReflect,
	}, {
		Name: "pulsar",
		URL:  "github.com/cosmos/cosmos-proto",
		New:  pulsar.NewPulsarSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "gogo/protobuf",
		URL:  "github.com/gogo/protobuf/proto",
		New:  gogo.NewGogoProtoSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "gogo/jsonpb",
		URL:  "github.com/gogo/protobuf/proto",
		New:  gogo.NewGogoJsonSerializer,

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKCodegen,
	}, {
		Name: "colfer",
		URL:  "github.com/pascaldekloe/colfer",
		New:  colfer.NewColferSerializer,

		TimeSupport: TSCustom,
		APIKind:     AKCodegen,
	}, {
		Name: "gencode",
		URL:  "github.com/andyleap/gencode",
		New:  gencode.NewGencodeSerializer,

		TimeSupport: TSFullTzOffset,
		APIKind:     AKCodegen,
	}, {
		Name: "gencode/unsafe_reuse",
		URL:  "github.com/andyleap/gencode",
		New:  gencode.NewGencodeUnsafeSerializer,

		BufferReuseMarshal:    true,
		UnsafeStringUnmarshal: true,
		TimeSupport:           TSFullTzOffset,
		APIKind:               AKCodegen,
	}, {
		Name: "calmh/xdr",
		URL:  "github.com/calmh/xdr",
		New:  xdrcalmh.NewXDRCalmhSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "goavro",
		URL:  "gopkg.in/linkedin/goavro.v1",
		New:  avro.NewAvroA,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "avro2/text",
		URL:  "github.com/linkedin/goavro",
		New:  avro.NewAvro2Txt,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "avro2/binary",
		URL:  "github.com/linkedin/goavro",
		New:  avro.NewAvro2Bin,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "ikea",
		URL:  "github.com/ikkerens/ikeapack",
		New:  ikea.NewIkeaSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "shamaton/msgpack/map",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonMapMsgpackSerializer,

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "shamaton/msgpack/array",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonArrayMsgPackSerializer,

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "shamaton/msgpackgen/map",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonMapMsgPackgenSerializer,

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "shamaton/msgpackgen/array",
		URL:  "github.com/shamaton/msgpack",
		New:  shamaton.NewShamatonArrayMsgpackgenSerializer,

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "ssz",
		URL:  "github.com/prysmaticlabs/go-ssz",
		New:  ssz.NewSSZSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "200sc/bebop",
		URL:  "github.com/200sc/bebop",
		New:  bebop200sc.NewBebop200ScSerializer,

		TimeSupport: TSCustom,
		APIKind:     AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "200sc/bebop/reuse",
		URL:  "github.com/200sc/bebop",
		New:  bebop200sc.NewBebop200ScReuseSerializer,

		BufferReuseMarshal: true,
		TimeSupport:        TSCustom,
		APIKind:            AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "wellquite/bebop",
		URL:  "wellquite.org/bebop",
		New:  bebopwellquite.NewBebopWellquiteSerializer,

		TimeSupport: TSCustom,
		APIKind:     AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "wellquite/bebop/reuse",
		URL:  "wellquite.org/bebop",
		New:  bebopwellquite.NewBebopWellquiteReuseSerializer,

		BufferReuseMarshal: true,
		TimeSupport:        TSCustom,
		APIKind:            AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "fastjson",
		URL:  "github.com/valyala/fastjson",
		New:  fastjson.NewFastJSONSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "fastjson/reuse",
		URL:  "github.com/valyala/fastjson",
		New:  fastjson.NewFastJSONReuseSerializer,

		BufferReuseMarshal: true,
		TimeSupport:        TSNoSupport,
		APIKind:            AKManual,
	}, {
		Name: "benc",
		URL:  "github.com/deneonet/benc",
		New:  benc.NewBENCSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "benc/usafe",
		URL:  "github.com/deneonet/benc",
		New:  benc.NewBENCUnsafeSerializer,

		UnsafeStringUnmarshal: true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
	}, {
		Name: "mus",
		URL:  "github.com/mus-format/mus-go",
		New:  mus.NewMUSSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "mus/unsafe_reuse",
		URL:  "github.com/mus-format/mus-go",
		New:  mus.NewMUSUnsafeSerializer,

		BufferReuseMarshal:    true,
		UnsafeStringUnmarshal: true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
	}, {
		Name: "baseline",
		URL:  "",
		New:  baseline.NewBaselineSerializer,

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
		Notes: []string{
			"This is a manually written encoding, designed to be the fastest possible for this benchmark.",
		},
	}, {
		Name: "baseline/unsafe_reuse",
		URL:  "",
		New:  baseline.NewBaselineUnsafeSerializer,

		UnsafeStringUnmarshal: true,
		BufferReuseMarshal:    true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
		Notes: []string{
			"This is a manually written encoding, designed to be the fastest possible for this benchmark.",
		},
	}, {
		Name: "fastape",
		URL:  "github.com/nazarifard/fastape",
		New:  fastape.NewTape,

		UnsafeStringUnmarshal: true,
		TimeSupport:           TSUnixNs,
		APIKind:               AKManual,
	},
}
