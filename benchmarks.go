package main

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/avro"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/baseline"
	bebop200sc "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/bebop_200sc"
	bebopwellquite "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/bebop_wellquite"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/benc"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/bencgen"
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
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/idr"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/ikea"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/jsoniter"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/mongobson"
	msgpacktinylib "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/msgpack_tinylib"
	msgpackvmihailenco "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/msgpack_vmihailenco"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/mus"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/omnibuf"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/protobuf"
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
	// Bench holds the serializer's marshal and unmarshal benchmarks.
	Bench goserbench.Benchmarks

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
		Bench: goserbench.NewBenchmarks(gotiny.NewGotinyCodec),

		TimeSupport: TSUnixNs,
		APIKind:     AKReflect,
	}, {
		Name: "msgp",
		URL:  "github.com/tinylib/msgp",
		Bench: goserbench.NewBenchmarks(msgpacktinylib.NewMsgpCodec),

		TimeSupport: TSFullRange,
		APIKind:     AKCodegen,
	}, {
		Name: "msgpack",
		URL:  "github.com/vmihailenco/msgpack",
		Bench: goserbench.NewBenchmarks(msgpackvmihailenco.NewVmihailencoMsgpackCodec),

		TimeSupport: TSFullRange,
		APIKind:     AKReflect,
	}, {
		Name: "json",
		URL:  "pkg.go/dev/encoding/json",
		Bench: goserbench.NewBenchmarks(stdlib.NewJSONCodec),

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKReflect,
	}, {
		Name: "jsoniter",
		URL:  "github.com/json-iterator/go",
		Bench: goserbench.NewBenchmarks(jsoniter.NewJSONIterCodec),

		TimeSupport: TSCustom,
		APIKind:     AKReflect,
	}, {
		Name: "easyjson",
		URL:  "github.com/mailru/easyjson",
		Bench: goserbench.NewBenchmarks(easyjson.NewEasyJSONCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKCodegen,
	}, {
		Name: "bson",
		URL:  "gopkg.in/mgo.v2/bson",
		Bench: goserbench.NewBenchmarks(bson.NewBsonCodec),

		TimeSupport: TSUnixMs,
		APIKind:     AKReflect,
	}, {
		Name: "mongobson",
		URL:  "go.mongodb.org/mongo-driver/mongo",
		Bench: goserbench.NewBenchmarks(mongobson.NewMongoBSONCodec),

		TimeSupport: TSUnixMs,
		APIKind:     AKReflect,
	}, {
		Name: "gob",
		URL:  "pkg.go.dev/encoding/gob",
		Bench: goserbench.NewBenchmarks(stdlib.NewGobCodec),

		TimeSupport: TSFullTzOffset,
		APIKind:     AKReflect,
	}, {
		Name: "davecgh/xdr",
		URL:  "github.com/davecgh/go-xdr/xdr",
		Bench: goserbench.NewBenchmarks(xdrdavecgh.NewXDRDavecghCodec),

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKReflect,
	}, {
		Name: "ugorji/msgpack",
		URL:  "github.com/ugorji/go/codec",
		Bench: goserbench.NewBenchmarks(ugorji.NewUgorjiMsgPackCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "ugorji/binc",
		URL:  "github.com/ugorji/go/codec",
		Bench: goserbench.NewBenchmarks(ugorji.NewUgorjiBincCodec),

		TimeSupport: TSFullTzOffset,
		APIKind:     AKReflect,
	}, {
		Name: "sereal",
		URL:  "github.com/Sereal/Sereal/Go/sereal",
		Bench: goserbench.NewBenchmarks(sereal.NewSerealCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "alecthomas/binary",
		URL:  "github.com/alecthomas/binary",
		Bench: goserbench.NewBenchmarks(binaryalecthomas.NewBinaryCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKReflect,
	}, {
		Name: "flatbuffers",
		URL:  "github.com/google/flatbuffers/go",
		Bench: goserbench.NewBenchmarks(flatbuffers.NewFlatBuffersCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "flatbuffers/unsafe_reuse",
		URL:  "github.com/google/flatbuffers/go",
		Bench: goserbench.NewBenchmarks(flatbuffers.NewFlatBuffersUnsafeReuseCodec),

		TimeSupport:           TSNoSupport,
		APIKind:               AKCodegen,
		UnsafeStringUnmarshal: true,
		BufferReuseMarshal:    true,
	}, {
		Name: "capnproto",
		URL:  "github.com/glycerine/go-capnproto",
		Bench: goserbench.NewBenchmarks(capnproto.NewCapNProtoCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "hprose",
		URL:  "github.com/hprose/hprose-go/io",
		Bench: goserbench.NewBenchmarks(hprose.NewHproseCodec),

		TimeSupport: TSCustom,
		APIKind:     AKManual,
	}, {
		Name: "hprose2",
		URL:  "github.com/hprose/hprose-golang/io",
		Bench: goserbench.NewBenchmarks(hprose2.NewHProse2Codec),

		TimeSupport: TSCustom,
		APIKind:     AKManual,
	}, {
		Name: "dedis/protobuf",
		URL:  "go.dedis.ch/protobuf",
		Bench: goserbench.NewBenchmarks(protobufdedis.NewProtobufCodec),

		TimeSupport: TSUnixNs,
		APIKind:     AKReflect,
	}, {
		Name: "pulsar",
		URL:  "github.com/cosmos/cosmos-proto",
		Bench: goserbench.NewBenchmarks(pulsar.NewPulsarCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "protobuf-go",
		URL:  "github.com/protocolbuffers/protobuf-go",
		Bench: goserbench.NewBenchmarks(protobuf.NewProtobufCodec),

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKCodegen,
	}, {
		Name: "omnibuf",
		URL:  "github.com/omnibuf/omnibuf-go",
		Bench: goserbench.NewBenchmarks(omnibuf.NewOmniBufCodec),

		TimeSupport: TSUnixNs,
		APIKind:     AKCodegen,
	}, {
		Name: "omnibuf/unsafe_reuse",
		URL:  "github.com/omnibuf/omnibuf-go",
		Bench: goserbench.NewBenchmarks(omnibuf.NewOmniBufUnsafeReuseCodec),

		UnsafeStringUnmarshal: true,
		BufferReuseMarshal:    true,
		TimeSupport:           TSUnixNs,
		APIKind:               AKCodegen,
	}, {
		Name: "omnibuf/reflect",
		URL:  "github.com/omnibuf/omnibuf-go",
		Bench: goserbench.NewBenchmarks(omnibuf.NewOmniBufReflectCodec),

		TimeSupport: TSUnixNs,
		APIKind:     AKReflect,
	}, {
		Name: "gogo/protobuf",
		URL:  "github.com/gogo/protobuf/proto",
		Bench: goserbench.NewBenchmarks(gogo.NewGogoProtoCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "gogo/jsonpb",
		URL:  "github.com/gogo/protobuf/proto",
		Bench: goserbench.NewBenchmarks(gogo.NewGogoJsonCodec),

		TimeSupport: TSRFC3339Ns,
		APIKind:     AKCodegen,
	}, {
		Name: "colfer",
		URL:  "github.com/pascaldekloe/colfer",
		Bench: goserbench.NewBenchmarks(colfer.NewColferCodec),

		TimeSupport: TSCustom,
		APIKind:     AKCodegen,
	}, {
		Name: "gencode",
		URL:  "github.com/andyleap/gencode",
		Bench: goserbench.NewBenchmarks(gencode.NewGencodeCodec),

		TimeSupport: TSFullTzOffset,
		APIKind:     AKCodegen,
	}, {
		Name: "gencode/unsafe_reuse",
		URL:  "github.com/andyleap/gencode",
		Bench: goserbench.NewBenchmarks(gencode.NewGencodeUnsafeCodec),

		BufferReuseMarshal:    true,
		UnsafeStringUnmarshal: true,
		TimeSupport:           TSFullTzOffset,
		APIKind:               AKCodegen,
	}, {
		Name: "calmh/xdr",
		URL:  "github.com/calmh/xdr",
		Bench: goserbench.NewBenchmarks(xdrcalmh.NewXDRCalmhCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "goavro",
		URL:  "gopkg.in/linkedin/goavro.v1",
		Bench: goserbench.NewBenchmarks(avro.NewGoAvroCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "avro2/text",
		URL:  "github.com/linkedin/goavro",
		Bench: goserbench.NewBenchmarks(avro.NewAvro2TextCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "avro2/binary",
		URL:  "github.com/linkedin/goavro",
		Bench: goserbench.NewBenchmarks(avro.NewAvro2BinaryCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "ikea",
		URL:  "github.com/ikkerens/ikeapack",
		Bench: goserbench.NewBenchmarks(ikea.NewIkeaCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "shamaton/msgpack/map",
		URL:  "github.com/shamaton/msgpack",
		Bench: goserbench.NewBenchmarks(shamaton.NewShamatonMapMsgpackCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "shamaton/msgpack/array",
		URL:  "github.com/shamaton/msgpack",
		Bench: goserbench.NewBenchmarks(shamaton.NewShamatonArrayMsgPackCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "shamaton/msgpackgen/map",
		URL:  "github.com/shamaton/msgpack",
		Bench: goserbench.NewBenchmarks(shamaton.NewShamatonMapMsgPackgenCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "shamaton/msgpackgen/array",
		URL:  "github.com/shamaton/msgpack",
		Bench: goserbench.NewBenchmarks(shamaton.NewShamatonArrayMsgpackgenCodec),

		TimeSupport: TSUnknown,
		APIKind:     AKReflect,
	}, {
		Name: "ssz",
		URL:  "github.com/prysmaticlabs/go-ssz",
		Bench: goserbench.NewBenchmarks(ssz.NewSSZCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "200sc/bebop",
		URL:  "github.com/200sc/bebop",
		Bench: goserbench.NewBenchmarks(bebop200sc.NewBebop200ScCodec),

		TimeSupport: TSCustom,
		APIKind:     AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "200sc/bebop/reuse",
		URL:  "github.com/200sc/bebop",
		Bench: goserbench.NewBenchmarks(bebop200sc.NewBebop200ScReuseCodec),

		BufferReuseMarshal: true,
		TimeSupport:        TSCustom,
		APIKind:            AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "wellquite/bebop",
		URL:  "wellquite.org/bebop",
		Bench: goserbench.NewBenchmarks(bebopwellquite.NewBebopWellquiteCodec),

		TimeSupport: TSCustom,
		APIKind:     AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "wellquite/bebop/reuse",
		URL:  "wellquite.org/bebop",
		Bench: goserbench.NewBenchmarks(bebopwellquite.NewBebopWellquiteReuseCodec),

		BufferReuseMarshal: true,
		TimeSupport:        TSCustom,
		APIKind:            AKCodegen,
		Notes: []string{
			"time.Time values are encoded with 100 nanosecond precision.",
		},
	}, {
		Name: "fastjson",
		URL:  "github.com/valyala/fastjson",
		Bench: goserbench.NewBenchmarks(fastjson.NewFastJSONCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "fastjson/reuse",
		URL:  "github.com/valyala/fastjson",
		Bench: goserbench.NewBenchmarks(fastjson.NewFastJSONReuseCodec),

		BufferReuseMarshal: true,
		TimeSupport:        TSNoSupport,
		APIKind:            AKManual,
	}, {
		Name: "benc",
		URL:  "github.com/deneonet/benc",
		Bench: goserbench.NewBenchmarks(benc.NewBENCCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "benc/usafe",
		URL:  "github.com/deneonet/benc",
		Bench: goserbench.NewBenchmarks(benc.NewBENCUnsafeCodec),

		UnsafeStringUnmarshal: true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
	}, {
		Name: "bencgen",
		URL:  "github.com/deneonet/benc",
		Bench: goserbench.NewBenchmarks(bencgen.NewBENCGenCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKCodegen,
	}, {
		Name: "bencgen/unsafe",
		URL:  "github.com/deneonet/benc",
		Bench: goserbench.NewBenchmarks(bencgen.NewBENCGenUnsafeCodec),

		UnsafeStringUnmarshal: true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKCodegen,
	}, {
		Name: "mus",
		URL:  "github.com/mus-format/mus-go",
		Bench: goserbench.NewBenchmarks(mus.NewMUSCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
	}, {
		Name: "mus/unsafe_reuse",
		URL:  "github.com/mus-format/mus-go",
		Bench: goserbench.NewBenchmarks(mus.NewMUSUnsafeReuseCodec),

		BufferReuseMarshal:    true,
		UnsafeStringUnmarshal: true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
	}, {
		Name: "idr",
		URL:  "github.com/chmike/ditp",
		Bench: goserbench.NewBenchmarks(idr.NewIDRCodec),

		TimeSupport: TSFullTzOffset,
		APIKind:     AKManual,
		Notes:       []string{"low level IDR encoding demo with same nsec time encoding as benc"},
	}, {
		Name: "idr/reuse",
		URL:  "github.com/chmike/ditp",
		Bench: goserbench.NewBenchmarks(idr.NewIDRReuseCodec),

		BufferReuseMarshal: true,
		TimeSupport:        TSFullTzOffset,
		APIKind:            AKManual,
		Notes:              []string{"low level IDR encoding demo with same nsec time encoding as benc"},
	}, {
		Name: "baseline",
		URL:  "",
		Bench: goserbench.NewBenchmarks(baseline.NewBaselineCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
		Notes: []string{
			"This is a manually written encoding, designed to be the fastest possible for this benchmark.",
		},
	}, {
		Name: "baseline/unsafe_reuse",
		URL:  "",
		Bench: goserbench.NewBenchmarks(baseline.NewBaselineUnsafeCodec),

		UnsafeStringUnmarshal: true,
		BufferReuseMarshal:    true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
		Notes: []string{
			"This is a manually written encoding, designed to be the fastest possible for this benchmark.",
		},
	}, {
		Name: "baseline_rw",
		URL:  "",
		Bench: goserbench.NewBenchmarks(baseline.NewBaselineReaderWriterCodec),

		TimeSupport: TSNoSupport,
		APIKind:     AKManual,
		Notes: []string{
			"This is a manually written encoding, designed to be the fastest possible for this benchmark, using an io.Reader/io.Writer as the API.",
		},
	}, {
		Name: "baseline_rw/unsafe_reuse",
		URL:  "",
		Bench: goserbench.NewBenchmarks(baseline.NewBaselineReaderWriterUnsafeReuseCodec),

		UnsafeStringUnmarshal: true,
		BufferReuseMarshal:    true,
		TimeSupport:           TSNoSupport,
		APIKind:               AKManual,
		Notes: []string{
			"This is a manually written encoding, designed to be the fastest possible for this benchmark, using an io.Reader/io.Writer as the API.",
		},
	}, {
		Name: "fastape",
		URL:  "github.com/nazarifard/fastape",
		Bench: goserbench.NewBenchmarks(fastape.NewFastapeCodec),

		UnsafeStringUnmarshal: true,
		TimeSupport:           TSUnixNs,
		APIKind:               AKManual,
	},
}
