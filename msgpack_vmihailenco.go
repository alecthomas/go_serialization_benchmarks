package goserbench

import (
	vmihailenco "github.com/vmihailenco/msgpack/v5"
)

type VmihailencoMsgpackSerializer struct{}

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	return vmihailenco.Marshal(o)
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func NewVmihailencoMsgpackSerialier() VmihailencoMsgpackSerializer {
	return VmihailencoMsgpackSerializer{}
}
