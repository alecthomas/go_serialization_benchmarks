package copi

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/nazarifard/copi"
)

type smallStructCp struct {
	NameCp     copi.StringCp
	BirthDayCp copi.TimeCp
	PhoneCp    copi.StringCp
	SiblingsCp copi.UnitCp[int]
	SpouseCp   copi.UnitCp[bool]
	MoneyCp    copi.UnitCp[float64]
}

func (cp smallStructCp) SizeOf(p goserbench.SmallStruct) int {
	return cp.NameCp.SizeOf(p.Name) +
		cp.BirthDayCp.SizeOf(p.BirthDay) +
		cp.PhoneCp.SizeOf(p.Phone) +
		cp.SiblingsCp.SizeOf(p.Siblings) +
		cp.SpouseCp.SizeOf(p.Spouse) +
		cp.MoneyCp.SizeOf(p.Money)
}

func (cp smallStructCp) Marshal(o interface{}) (buf []byte, err error) {
	p := o.(*goserbench.SmallStruct)
	buf = make([]byte, cp.SizeOf(*p))

	k, n := 0, 0
	k, _ = cp.NameCp.Marshal(p.Name, buf[n:])
	n += k
	k, _ = cp.BirthDayCp.Marshal(p.BirthDay, buf[n:])
	n += k
	k, _ = cp.PhoneCp.Marshal(p.Phone, buf[n:])
	n += k
	k, _ = cp.SiblingsCp.Marshal(p.Siblings, buf[n:])
	n += k
	k, _ = cp.SpouseCp.Marshal(p.Spouse, buf[n:])
	n += k
	k, _ = cp.MoneyCp.Marshal(p.Money, buf[n:])
	n += k
	return
}

func (cp smallStructCp) Unmarshal(bs []byte, o interface{}) (err error) {
	p := o.(*goserbench.SmallStruct)

	k, n := 0, 0
	k, err = cp.NameCp.Unmarshal(bs[n:], &p.Name)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.BirthDayCp.Unmarshal(bs[n:], &p.BirthDay)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.PhoneCp.Unmarshal(bs[n:], &p.Phone)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.SiblingsCp.Unmarshal(bs[n:], &p.Siblings)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.SpouseCp.Unmarshal(bs[n:], &p.Spouse)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.MoneyCp.Unmarshal(bs[n:], &p.Money)
	n += k
	if err != nil {
		return err
	}

	return
}

func NewCopiSerializer() goserbench.Serializer {
	return smallStructCp{}
}
