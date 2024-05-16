package goserbench

import "github.com/Sereal/Sereal/Go/sereal"

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o interface{}) ([]byte, error) {
	return sereal.Marshal(o)
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func NewSerealSerializer() SerealSerializer {
	return SerealSerializer{}
}
