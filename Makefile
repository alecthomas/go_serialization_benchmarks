# This is necessary due to the use of two conflicting generator commands for capnproto
.NOTPARALLEL:

all: Colfer.go FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go structdef.capnp.go structdef.capnp2.go gencode.schema.gen.go gencode-unsafe.schema.gen.go structdefxdr_generated.go

Colfer.go:
	go run github.com/pascaldekloe/colfer/cmd/colf go
	mv goserbench/Colfer.go .
	rmdir goserbench

FlatBufferA.go: flatbuffers-structdef.fbs
	flatc -g flatbuffers-structdef.fbs
	mv flatbuffersmodels/FlatBufferA.go FlatBufferA.go
	rmdir flatbuffersmodels
	sed -i '' 's/flatbuffersmodels/goserbench/' FlatBufferA.go

msgp_gen.go: structdef.go
	go run github.com/tinylib/msgp -o msgp_gen.go -file structdef.go -io=false -tests=false

structdef_easyjson.go: structdef.go
	go run github.com/mailru/easyjson/easyjson -all structdef.go

structdef-gogo.pb.go: structdef-gogo.proto
	protoc --gogofaster_out=. -I. -I${GOPATH}/src  -I${GOPATH}/src/github.com/gogo/protobuf/protobuf structdef-gogo.proto

structdef.pb.go: structdef.proto
	protoc --go_out=. structdef.proto

structdef.capnp2.go: structdef.capnp2
	go get -u zombiezen.com/go/capnproto2/... # conflicts with go-capnproto
	capnp compile -I${GOPATH}/src -ogo structdef.capnp2

structdef.capnp.go: structdef.capnp
	go get -u github.com/glycerine/go-capnproto/capnpc-go # conflicts with capnproto2
	capnp compile -I${GOPATH}/src -ogo structdef.capnp

gencode.schema.gen.go: gencode.schema
	go run github.com/andyleap/gencode go -schema=gencode.schema -package=goserbench

gencode-unsafe.schema.gen.go: gencode-unsafe.schema
	go run github.com/andyleap/gencode go -schema=gencode-unsafe.schema -package=goserbench -unsafe

structdefxdr_generated.go: structdefxdr.go
	go run github.com/calmh/xdr/cmd/genxdr -o structdefxdr_generated.go structdefxdr.go

.PHONY: clean
clean:
	rm -f Colfer.go FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go structdef.capnp.go structdef.capnp2.go gencode.schema.gen.go gencode-unsafe.schema.gen.go structdefxdr_generated.go

.PHONY: install
install:
	go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/tinylib/msgp
	go get -u github.com/andyleap/gencode
	go get -u github.com/mailru/easyjson/...
	go get -u go.dedis.ch/protobuf
	go get -u github.com/Sereal/Sereal/Go/sereal
	go get -u github.com/alecthomas/binary
	go get -u github.com/davecgh/go-xdr/xdr
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/google/flatbuffers/go
	go get -u github.com/tinylib/msgp/msgp
	go get -u github.com/ugorji/go/codec
	go get -u gopkg.in/mgo.v2/bson
	go get -u gopkg.in/vmihailenco/msgpack.v2
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/hprose/hprose-go/io
	go get -u github.com/pascaldekloe/colfer/cmd/colf
	go get -u github.com/calmh/xdr
	go get -u github.com/niubaoshu/gotiny
