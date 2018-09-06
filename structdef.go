package goserbench

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//go:generate msgp -o msgp_gen.go -io=false -tests=false
//easyjson:json
type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

func NewA() Object {
	return &A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}

func (a *A) Reset() { *a = A{} }
func (a *A) AssertEqual(i interface{}) (bool, error) {
	o, ok := i.(*A)
	if !ok {
		return false, errors.New("not A type")
	}
	isEqual := o.Name == a.Name && o.Phone == a.Phone && o.Siblings == a.Siblings && o.Spouse == a.Spouse && o.Money == a.Money && o.BirthDay.Equal(a.BirthDay)
	var err error
	if !isEqual {
		err = fmt.Errorf("\n%#v\n%#v", *o, *a)
	}
	return isEqual, err
}
