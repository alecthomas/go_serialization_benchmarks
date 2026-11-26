package msgpackvmihailenco

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	vmihailenco "github.com/vmihailenco/msgpack/v5"
)

type VmihailencoMsgpackCodec struct{ goserbench.NativeSmallStruct }

func (m VmihailencoMsgpackCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return vmihailenco.Marshal(o)
}

func (m VmihailencoMsgpackCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return vmihailenco.Unmarshal(d, o)
}

func NewVmihailencoMsgpackCodec() goserbench.Codec[goserbench.SmallStruct] {
	return VmihailencoMsgpackCodec{}
}
