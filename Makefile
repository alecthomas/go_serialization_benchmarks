# This is necessary due to the use of two conflicting generator commands for capnproto
.NOTPARALLEL:

all: Colfer.go FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go structdefxdr_generated.go structdef-bebop.go structdef_msgpackgen.go musgo.go structdef.pulsar.go

Colfer.go:
	go run github.com/pascaldekloe/colfer/cmd/colf@latest go
	mv goserbench/Colfer.go .
	rmdir goserbench

FlatBufferA.go: flatbuffers-structdef.fbs
	flatc --go --go-namespace goserbench flatbuffers-structdef.fbs
	mv goserbench/FlatBufferA.go FlatBufferA.go
	rmdir goserbench

msgp_gen.go: structdef.go
	go run github.com/tinylib/msgp@latest -o msgp_gen.go -file structdef.go -io=false -tests=false

structdef_easyjson.go: structdef.go
	go run github.com/mailru/easyjson/easyjson@latest -no_std_marshalers -all structdef.go

structdef-gogo.pb.go: structdef-gogo.proto
	protoc --gogofaster_opt=paths=source_relative --gogofaster_out=. -I. -I${GOPATH}/src  -I${GOPATH}/src/github.com/gogo/protobuf/protobuf structdef-gogo.proto

structdef.pb.go: structdef.proto
	protoc --go_opt=paths=source_relative --go_out=. structdef.proto

structdef-pulsar.pulsar.go: structdef-pulsar.proto
	protoc --go-pulsar_out=. --go-pulsar_opt=paths=source_relative --go-pulsar_opt=features=protoc+fast -I . structdef-pulsar.proto

#structdef.capnp2.go: structdef.capnp2
#	go run zombiezen.com/go/capnproto2/capnpc-go@latest compile -I${GOPATH}/src -ogo structdef.capnp2

#structdef.capnp.go: structdef.capnp
#	go run github.com/glycerine/go-capnproto/capnpc-go@latest compile -I${GOPATH}/src -ogo structdef.capnp

#gencode.schema.gen.go: gencode.schema
#	go run github.com/andyleap/gencode@latest go -schema=gencode.schema -package=goserbench
#
#gencode-unsafe.schema.gen.go: gencode-unsafe.schema
#	go run github.com/andyleap/gencode@latest go -schema=gencode-unsafe.schema -package=goserbench -unsafe

structdefxdr_generated.go: structdefxdr.go
	go run github.com/calmh/xdr/cmd/genxdr@latest -o structdefxdr_generated.go structdefxdr.go

structdef-bebop.go:
	go run github.com/200sc/bebop/main/bebopc-go@latest -i structdef-bebop.bop -o structdef-bebop.go --package goserbench

structdef_msgpackgen.go: structdef.go
	go run github.com/shamaton/msgpackgen@latest -input-file structdef.go -output-file structdef_msgpackgen.go -strict

musgo.go: structdef.go
	go run ./musgo-gen.go

.PHONY: clean
clean:
	rm -f Colfer.go FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go structdefxdr_generated.go structdef_msgpackgen.go NoTimeA.musgen.go NoTimeAUnsafe.musgen.go
.PHONY: install
install:
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go install github.com/gogo/protobuf/gogoproto@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/tinylib/msgp@latest
	go install github.com/andyleap/gencode@latest
	go install github.com/mailru/easyjson@latest
	go install go.dedis.ch/protobuf@latest
	go install github.com/Sereal/Sereal/Go/sereal@latest
	go install github.com/alecthomas/binary@latest
	go install github.com/davecgh/go-xdr/xdr@latest
	go install github.com/gogo/protobuf/proto@latest
	go install github.com/google/flatbuffers/go@latest
	go install github.com/tinylib/msgp/msgp@latest
	go install github.com/ugorji/go/codec@latest
	go install gopkg.in/mgo.v2/bson@latest
	go install github.com/vmihailenco/msgpack/v5@latest
	go install github.com/golang/protobuf/proto@latest
	go install github.com/hprose/hprose-go/io@latest
	go install github.com/pascaldekloe/colfer/cmd/colf@latest
	go install github.com/calmh/xdr@latest
	go install github.com/niubaoshu/gotiny@latest
	go install github.com/200sc/bebop@latest
	go install github.com/200sc/bebop/main/bebopc-go@latest
	go install github.com/shamaton/msgpackgen@latest
	go install github.com/ymz-ncnk/musgo@latest
	go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest