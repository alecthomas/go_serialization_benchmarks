all: FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go vitess_test.go

FlatBufferA.go: flatbuffers-structdef.fbs
	flatc -g flatbuffers-structdef.fbs
	mv flatbuffersmodels/FlatBufferA.go FlatBufferA.go
	rmdir flatbuffersmodels
	sed -i '' 's/flatbuffersmodels/goserbench/' FlatBufferA.go

msgp_gen.go: structdef.go
	go generate

structdef-gogo.pb.go: structdef-gogo.proto
	protoc --gogo_out=. -I. -I${GOPATH}/src  -I${GOPATH}/src/github.com/gogo/protobuf/protobuf structdef-gogo.proto

structdef.pb.go: structdef.proto
	protoc --go_out=. structdef.proto

vitess_test.go: structdef.go
	bsongen -file=structdef.go -o=vitess_test.go -type=A

.PHONY: clean
clean:
	rm -f FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go vitess_test.go

.PHONY: install
install:
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/tinylib/msgp
	go get -u github.com/youtube/vitess/go/cmd/bsongen

	go get -u github.com/DeDiS/protobuf
	go get -u github.com/Sereal/Sereal/Go/sereal
	go get -u github.com/alecthomas/binary
	go get -u github.com/davecgh/go-xdr/xdr
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/google/flatbuffers/go
	go get -u github.com/tinylib/msgp/msgp
	go get -u github.com/ugorji/go/codec
	go get -u github.com/youtube/vitess/go/bson
	go get -u gopkg.in/mgo.v2/bson
	go get -u gopkg.in/vmihailenco/msgpack.v2
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/hprose/hprose-go/io
