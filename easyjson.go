package goserbench

import easyjson "github.com/mailru/easyjson"

type EasyJSONSerializer struct{}

func (m EasyJSONSerializer) Marshal(o interface{}) ([]byte, error) {
	return easyjson.Marshal(o.(easyjson.Marshaler))
}

func (m EasyJSONSerializer) Unmarshal(d []byte, o interface{}) error {
	return easyjson.Unmarshal(d, o.(*A))
}

func newEasyJSONSerializer() EasyJSONSerializer {
	return EasyJSONSerializer{}
}
