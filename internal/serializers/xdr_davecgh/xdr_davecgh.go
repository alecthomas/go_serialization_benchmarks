package xdrdavecgh

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/davecgh/go-xdr/xdr"
)

type XDRDavecghCodec struct{ goserbench.NativeSmallStruct }

func (x XDRDavecghCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return xdr.Marshal(o)
}

func (x XDRDavecghCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func NewXDRDavecghCodec() goserbench.Codec[goserbench.SmallStruct] {
	return XDRDavecghCodec{}
}
