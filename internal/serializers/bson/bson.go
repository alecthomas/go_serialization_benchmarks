package bson

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"gopkg.in/mgo.v2/bson"
)

type BsonCodec struct{ goserbench.NativeSmallStruct }

func (m BsonCodec) TimePrecision() time.Duration {
	return time.Millisecond
}

func (m BsonCodec) Marshal(o *goserbench.SmallStruct) ([]byte, error) {
	return bson.Marshal(o)
}

func (m BsonCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	return bson.Unmarshal(d, o)
}

func NewBsonCodec() goserbench.Codec[goserbench.SmallStruct] {
	return BsonCodec{}
}
