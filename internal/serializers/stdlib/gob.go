package stdlib

import (
	"bytes"
	"encoding/gob"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type GobCodec struct{ goserbench.NativeSmallStruct }

func (g *GobCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(o)
	return buf.Bytes(), err
}

func (g *GobCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return gob.NewDecoder(bytes.NewReader(d)).Decode(o)
}

func NewGobCodec() goserbench.Codec[goserbench.SmallStruct] {
	// registration required before first use
	gob.Register(goserbench.SmallStruct{})
	return &GobCodec{}
}
