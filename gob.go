package goserbench

import (
	"bytes"
	"encoding/gob"
)

type GobSerializer struct{}

func (g *GobSerializer) Marshal(o interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(o)
	return buf.Bytes(), err
}

func (g *GobSerializer) Unmarshal(d []byte, o interface{}) error {
	return gob.NewDecoder(bytes.NewReader(d)).Decode(o)
}

func NewGobSerializer() *GobSerializer {
	// registration required before first use
	gob.Register(A{})
	return &GobSerializer{}
}
