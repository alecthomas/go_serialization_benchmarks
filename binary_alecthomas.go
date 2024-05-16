package goserbench

import "github.com/alecthomas/binary"

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o interface{}) ([]byte, error) {
	return binary.Marshal(o)
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func newBinarySerializer() BinarySerializer {
	return BinarySerializer{}
}
