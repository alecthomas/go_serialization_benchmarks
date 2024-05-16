package goserbench

import "github.com/ugorji/go/codec"

type UgorjiCodecSerializer struct {
	codec.Handle
}

func (u *UgorjiCodecSerializer) Marshal(o interface{}) ([]byte, error) {
	var bs []byte
	return bs, codec.NewEncoderBytes(&bs, u.Handle).Encode(o)
}

func (u *UgorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	return codec.NewDecoderBytes(d, u.Handle).Decode(o)
}

func NewUgorjiCodecMsgPack() *UgorjiCodecSerializer {
	return &UgorjiCodecSerializer{&codec.MsgpackHandle{}}
}

func NewUgorjiCodecBinc() *UgorjiCodecSerializer {
	h := &codec.BincHandle{}
	h.AsSymbols = 0
	return &UgorjiCodecSerializer{h}
}
