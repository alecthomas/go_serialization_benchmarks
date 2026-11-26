package gogo

import (
	"bytes"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

type GogoProtoCodec struct {
	// marshaller and unmarshaller are set on creation to either binary
	// or json marshallers.
	marshaller   func(proto.Message) ([]byte, error)
	unmarshaller func([]byte, proto.Message) error
}

func (s *GogoProtoCodec) Marshal(a *GogoProtoBufA) ([]byte, error) {
	return s.marshaller(a)
}

// Unmarshal relies on the harness zeroing a first, because gogoproto in
// jsonpb mode does not clear fields that are absent from the input.
func (s *GogoProtoCodec) Unmarshal(bs []byte, a *GogoProtoBufA) error {
	return s.unmarshaller(bs, a)
}

func (s *GogoProtoCodec) FromSmall(v *goserbench.SmallStruct, a *GogoProtoBufA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (s *GogoProtoCodec) ToSmall(a *GogoProtoBufA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewGogoProtoCodec() goserbench.Codec[GogoProtoBufA] {
	return &GogoProtoCodec{
		marshaller:   proto.Marshal,
		unmarshaller: proto.Unmarshal,
	}
}

func NewGogoJsonCodec() goserbench.Codec[GogoProtoBufA] {
	marshaller := &jsonpb.Marshaler{}
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	return &GogoProtoCodec{
		marshaller: func(a proto.Message) ([]byte, error) {
			buf.Reset()
			err := marshaller.Marshal(buf, a)
			return buf.Bytes(), err
		},
		unmarshaller: func(bs []byte, a proto.Message) error {
			return jsonpb.Unmarshal(bytes.NewBuffer(bs), a)
		},
	}
}
