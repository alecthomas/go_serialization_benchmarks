package pulsar

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	pproto "google.golang.org/protobuf/proto"
)

type PulsarCodec struct{}

var _ goserbench.SerializerEnforcesTimezone = PulsarCodec{}

func (PulsarCodec) ForcesUTC() bool {
	return true
}

func (PulsarCodec) Marshal(a *PulsarBufA) ([]byte, error) {
	return pproto.Marshal(a)
}

func (PulsarCodec) Unmarshal(bs []byte, a *PulsarBufA) error {
	return pproto.Unmarshal(bs, a)
}

func (PulsarCodec) FromSmall(v *goserbench.SmallStruct, a *PulsarBufA) {
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
}

func (PulsarCodec) ToSmall(a *PulsarBufA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
}

func NewPulsarCodec() goserbench.Codec[PulsarBufA] {
	return PulsarCodec{}
}
