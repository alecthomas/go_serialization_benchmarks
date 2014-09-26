package goserbench

import (
	"time"
)

//go:generate msgp

type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}
