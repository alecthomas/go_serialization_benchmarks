package fastape

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/nazarifard/fastape"
)

type smallStructTape struct {
	NameTape     fastape.StringTape
	BirthDayTape fastape.UnitTape[int64]
	PhoneTape    fastape.StringTape
	SiblingsTape fastape.UnitTape[byte]
	SpouseTape   fastape.UnitTape[bool]
	MoneyTape    fastape.UnitTape[float64]
	buf          []byte
}

func (cp *smallStructTape) Marshal(o interface{}) (buf []byte, err error) {
	p := o.(*goserbench.SmallStruct)

	birthday := p.BirthDay.UnixNano()
	if cp.buf != nil {
		buf = cp.buf //bufer reuse
	} else {
		sizeof := cp.NameTape.Sizeof(p.Name) +
			cp.BirthDayTape.Sizeof(birthday) +
			cp.PhoneTape.Sizeof(p.Phone) +
			cp.SiblingsTape.Sizeof(byte(p.Siblings)) +
			cp.SpouseTape.Sizeof(p.Spouse) +
			cp.MoneyTape.Sizeof(p.Money)

		buf = make([]byte, sizeof)
	}

	k, n := 0, 0
	k, _ = cp.NameTape.Roll(p.Name, buf[n:])
	n += k
	k, _ = cp.BirthDayTape.Roll(birthday, buf[n:])
	n += k
	k, _ = cp.PhoneTape.Roll(p.Phone, buf[n:])
	n += k
	k, _ = cp.SiblingsTape.Roll(byte(p.Siblings), buf[n:])
	n += k
	k, _ = cp.SpouseTape.Roll(p.Spouse, buf[n:])
	n += k
	k, _ = cp.MoneyTape.Roll(p.Money, buf[n:])
	n += k
	return
}

func (cp smallStructTape) Unmarshal(bs []byte, o interface{}) (err error) {
	p := o.(*goserbench.SmallStruct)

	k, n := 0, 0
	k, err = cp.NameTape.Unroll(bs[n:], &p.Name)
	n += k
	if err != nil {
		return err
	}

	var nano int64
	k, err = cp.BirthDayTape.Unroll(bs[n:], &nano)
	p.BirthDay = time.Unix(0, nano)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.PhoneTape.Unroll(bs[n:], &p.Phone)
	n += k
	if err != nil {
		return err
	}

	var sib byte
	k, err = cp.SiblingsTape.Unroll(bs[n:], &sib)
	p.Siblings = int(sib)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.SpouseTape.Unroll(bs[n:], &p.Spouse)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.MoneyTape.Unroll(bs[n:], &p.Money)
	n += k
	if err != nil {
		return err
	}

	return
}

func NewTape() goserbench.Serializer {
	return &smallStructTape{}
}

func NewTape_Reuse() goserbench.Serializer {
	const maxSize = 0 +
		8 + //date
		8 + //money
		1 + //sibling
		1 + //spouse
		1 + goserbench.MaxSmallStructPhoneSize +
		1 + goserbench.MaxSmallStructNameSize

	return &smallStructTape{
		buf: make([]byte, maxSize),
	}
}
