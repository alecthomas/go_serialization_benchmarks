package fastjson

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/valyala/fastjson"
)

type FastJSONCodec struct {
	goserbench.NativeSmallStruct

	arena  fastjson.Arena
	object *fastjson.Value
	buf    []byte
}

func (s *FastJSONCodec) Marshal(v *goserbench.SmallStruct) (buf []byte, err error) {
	object, arena := s.object, s.arena
	object.Set("name", arena.NewString(v.Name))
	object.Set("birthday", arena.NewNumberInt(int(v.BirthDay.UnixNano())))
	object.Set("phone", arena.NewString(v.Phone))
	object.Set("siblings", arena.NewNumberInt(v.Siblings))
	var spouse *fastjson.Value
	if v.Spouse {
		spouse = arena.NewTrue()
	} else {
		spouse = arena.NewFalse()
	}
	object.Set("spouse", spouse)
	object.Set("money", arena.NewNumberFloat64(v.Money))
	return object.MarshalTo(s.buf), nil
}

func (s *FastJSONCodec) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	val, err := fastjson.ParseBytes(bs)
	if err != nil {
		return err
	}
	v.Name = string(val.GetStringBytes("name"))
	v.BirthDay = time.Unix(0, val.GetInt64("birthday"))
	v.Phone = string(val.GetStringBytes("phone"))
	v.Siblings = val.GetInt("siblings")
	v.Spouse = val.GetBool("spouse")
	v.Money = val.GetFloat64("money")
	return nil
}

func NewFastJSONCodec() goserbench.Codec[goserbench.SmallStruct] {
	var arena fastjson.Arena
	return &FastJSONCodec{
		object: arena.NewObject(),
		arena:  arena,
	}
}

func NewFastJSONReuseCodec() goserbench.Codec[goserbench.SmallStruct] {
	var arena fastjson.Arena
	return &FastJSONCodec{
		object: arena.NewObject(),
		arena:  arena,
		buf:    make([]byte, 0, 1024),
	}

}
