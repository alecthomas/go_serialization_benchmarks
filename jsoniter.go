package goserbench

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	jsoniterFast = jsoniter.ConfigFastest
)

type JsonIterSerializer struct{}

func (j JsonIterSerializer) Marshal(o interface{}) ([]byte, error) {
	return jsoniterFast.Marshal(o)
}

func (j JsonIterSerializer) Unmarshal(d []byte, o interface{}) error {
	return jsoniterFast.Unmarshal(d, o)
}

func (j JsonIterSerializer) ReduceFloat64Precision() uint {
	return 6
}

func NewJSONIterSerializer() Serializer {
	return JsonIterSerializer{}
}
