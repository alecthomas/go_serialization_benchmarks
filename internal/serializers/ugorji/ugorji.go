package ugorji

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/ugorji/go/codec"
)

type UgorjiCodec struct {
	goserbench.NativeSmallStruct

	codec.Handle
}

func (u *UgorjiCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	var bs []byte
	return bs, codec.NewEncoderBytes(&bs, u.Handle).Encode(o)
}

func (u *UgorjiCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return codec.NewDecoderBytes(d, u.Handle).Decode(o)
}

func NewUgorjiMsgPackCodec() goserbench.Codec[goserbench.SmallStruct] {
	return &UgorjiCodec{Handle: &codec.MsgpackHandle{}}
}

func NewUgorjiBincCodec() goserbench.Codec[goserbench.SmallStruct] {
	h := &codec.BincHandle{}
	h.AsSymbols = 0
	return &UgorjiCodec{Handle: h}
}
