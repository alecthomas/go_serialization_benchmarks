module github.com/alecthomas/go_serialization_benchmarks

go 1.19

require (
	github.com/200sc/bebop v0.0.8
	github.com/Sereal/Sereal v0.0.0-20190618215532-0b8ac451a863
	github.com/alecthomas/binary v0.0.0-20171101084825-6e8df1b1fb9d
	github.com/calmh/xdr v1.1.0
	github.com/davecgh/go-xdr v0.0.0-20161123171359-e6a2ba005892
	github.com/glycerine/go-capnproto v0.0.0-20190118050403-2d07de3aa7fc
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.3.2
	github.com/google/flatbuffers v1.11.0
	github.com/hprose/hprose-go v0.0.0-20161031134501-83de97da5004
	github.com/hprose/hprose-golang v2.0.4+incompatible
	github.com/ikkerens/ikeapack v1.5.1
	github.com/itsmontoya/mum v0.3.2
	github.com/json-iterator/go v1.1.7
	github.com/linkedin/goavro v2.1.0+incompatible
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e
	github.com/niubaoshu/gotiny v0.0.3
	github.com/prysmaticlabs/go-ssz v0.0.0-20190827151743-72881c4223d8
	github.com/shamaton/msgpack/v2 v2.0.0
	github.com/shamaton/msgpackgen v0.1.1
	github.com/tinylib/msgp v1.1.0
	github.com/ugorji/go/codec v1.1.7
	github.com/valyala/fastjson v1.6.3
	github.com/vmihailenco/msgpack/v4 v4.3.0
	github.com/ymz-ncnk/musgo v0.2.0
	go.dedis.ch/protobuf v1.0.6
	go.mongodb.org/mongo-driver v1.5.1
	gopkg.in/linkedin/goavro.v1 v1.0.5
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	zombiezen.com/go/capnproto2 v2.17.0+incompatible
)

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/cybriq/gotiny v0.0.5 // indirect
	github.com/glycerine/goconvey v0.0.0-20190410193231-58a59202ab31 // indirect
	github.com/glycerine/rbtree v0.0.0-20190406191118-ceb71889d809 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20210603182125-eeedf4a0e899 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/minio/sha256-simd v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/protolambda/zssz v0.1.1 // indirect
	github.com/prysmaticlabs/go-bitfield v0.0.0-20190825002834-fb724e897364 // indirect
	github.com/stretchrcom/testify v1.7.1 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/itsmontoya/mum v0.5.6 => github.com/mojura/enkodo v0.5.6
	github.com/niubaoshu/gotiny v0.0.3 => github.com/cybriq/gotiny v0.0.6-0.20220412231127-0a1864225fc8
	github.com/stretchrcom/testify v1.7.1 => github.com/stretchr/testify v1.7.1
)
