package goserbench

import (
	"bytes"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

type GogoProtoSerializer struct {
	a GogoProtoBufA

	// marshaller and unmarshaller are set on creation to either binary
	// or json marshallers.
	marshaller   func(proto.Message) ([]byte, error)
	unmarshaller func([]byte, proto.Message) error
}

func (s *GogoProtoSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return s.marshaller(a)
}

func (s *GogoProtoSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	// NOTE: gogoproto serialization in jsonpb mode does not automatically
	// clear fields with their empty value.
	a := &s.a
	*a = GogoProtoBufA{}

	err = s.unmarshaller(bs, a)
	if err != nil {
		return
	}

	v := o.(*A)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}

func newGogoProtoSerializer() *GogoProtoSerializer {
	return &GogoProtoSerializer{
		marshaller:   proto.Marshal,
		unmarshaller: proto.Unmarshal,
	}
}

func newGogoJsonSerializer() *GogoProtoSerializer {
	marshaller := &jsonpb.Marshaler{}
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	return &GogoProtoSerializer{
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
