package protobuf

import (
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type ProtobufCodec struct {
	// marshaller and unmarshaller are set on creation to either binary
	// or json marshallers.
	marshaller   func(proto.Message) ([]byte, error)
	unmarshaller func([]byte, proto.Message) error
}

func (s *ProtobufCodec) Marshal(a *ProtobufGo) ([]byte, error) {
	return s.marshaller(a)
}

func (s *ProtobufCodec) Unmarshal(d []byte, a *ProtobufGo) error {
	return s.unmarshaller(d, a)
}

func (s *ProtobufCodec) FromSmall(v *goserbench.SmallStruct, a *ProtobufGo) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (s *ProtobufCodec) ToSmall(a *ProtobufGo, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewProtobufCodec() goserbench.Codec[ProtobufGo] {
	return &ProtobufCodec{
		marshaller:   proto.Marshal,
		unmarshaller: proto.Unmarshal,
	}
}
