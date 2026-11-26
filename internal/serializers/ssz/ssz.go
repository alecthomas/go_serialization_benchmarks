package ssz

import (
	"math"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	ssz "github.com/prysmaticlabs/go-ssz"
)

type SSZA struct {
	Name     string
	BirthDay uint64 // ssz does not support int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64 // ssz does not support float64
}

type SSZCodec struct{}

func (SSZCodec) Marshal(a *SSZA) ([]byte, error) {
	return ssz.Marshal(a)
}

func (SSZCodec) Unmarshal(bs []byte, a *SSZA) error {
	return ssz.Unmarshal(bs, a)
}

func (SSZCodec) FromSmall(v *goserbench.SmallStruct, a *SSZA) {
	a.Name = v.Name
	a.BirthDay = uint64(v.BirthDay.UnixNano())
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
}

func (SSZCodec) ToSmall(a *SSZA, v *goserbench.SmallStruct) {
	v.Name = a.Name
	v.BirthDay = time.Unix(0, int64(a.BirthDay))
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
}

func NewSSZCodec() goserbench.Codec[SSZA] {
	return SSZCodec{}
}
