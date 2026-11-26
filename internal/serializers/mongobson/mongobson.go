package mongobson

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	mongobson "go.mongodb.org/mongo-driver/bson"
)

type MongoBsonCodec struct{ goserbench.NativeSmallStruct }

func (m MongoBsonCodec) TimePrecision() time.Duration {
	return time.Millisecond
}

func (m MongoBsonCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return mongobson.Marshal(o)
}

func (m MongoBsonCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return mongobson.Unmarshal(d, o)
}

func NewMongoBSONCodec() goserbench.Codec[goserbench.SmallStruct] {
	return MongoBsonCodec{}
}
