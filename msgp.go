package goserbench

import "github.com/tinylib/msgp/msgp"

type MsgpSerializer struct{}

func (m MsgpSerializer) Marshal(o interface{}) ([]byte, error) {
	return o.(msgp.Marshaler).MarshalMsg(nil)
}

func (m MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(msgp.Unmarshaler).UnmarshalMsg(d)
	return err
}

func NewMsgpSerializer() Serializer {
	return MsgpSerializer{}
}
