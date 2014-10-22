package goserbench

import (
	"time"
)

//go:generate msgp -o msgp_gen.go

type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}
