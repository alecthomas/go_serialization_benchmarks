package goserbench

import (
	"errors"
	"math/rand"
	"time"
)

type NoTimeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

func NewNoTimeA() Object {
	return &NoTimeA{
		Name:     randString(16),
		BirthDay: time.Now().Unix(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}

func (a *NoTimeA) Reset() { *a = NoTimeA{} }

func (a *NoTimeA) AssertEqual(i interface{}) (bool, error) {
	o, ok := i.(*NoTimeA)
	if !ok {
		return false, errors.New("not A type")
	}
	return *a == *o, nil
}
