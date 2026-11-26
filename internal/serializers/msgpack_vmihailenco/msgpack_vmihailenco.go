package msgpackvmihailenco

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	vmihailenco "github.com/vmihailenco/msgpack/v5"
)

type VmihailencoMsgpackSerializer struct{}

func (m VmihailencoMsgpackSerializer) Marshal(o any) ([]byte, error) {
	return vmihailenco.Marshal(o)
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o any) error {
	return vmihailenco.Unmarshal(d, o)
}

func NewVmihailencoMsgpackSerialier() goserbench.Serializer {
	return VmihailencoMsgpackSerializer{}
}
