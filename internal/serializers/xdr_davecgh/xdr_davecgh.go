package xdrdavecgh

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/davecgh/go-xdr/xdr"
)

type XDRDavecghSerializer struct{}

func (x XDRDavecghSerializer) Marshal(o interface{}) ([]byte, error) {
	return xdr.Marshal(o)
}

func (x XDRDavecghSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func NewXDRDavecghSerializer() goserbench.Serializer {
	return XDRDavecghSerializer{}
}
