package goserbench

import (
	"time"
)

//go:generate msgp -o msgp_gen.go -io=false -tests=false
//go:generate protoc --gogo_out=. -I. -I$GOPATH/src  -I$GOPATH/src/github.com/gogo/protobuf/protobuf structdef-gogo.proto

type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}
