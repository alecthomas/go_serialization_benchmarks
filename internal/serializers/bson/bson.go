package bson

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"gopkg.in/mgo.v2/bson"
)

type BsonSerializer struct{}

func (m BsonSerializer) TimePrecision() time.Duration {
	return time.Millisecond
}

func (m BsonSerializer) Marshal(o any) ([]byte, error) {
	return bson.Marshal(o)
}

func (m BsonSerializer) Unmarshal(d []byte, o any) error {
	return bson.Unmarshal(d, o)
}

func NewBsonSerializer() goserbench.Serializer {
	return BsonSerializer{}
}
