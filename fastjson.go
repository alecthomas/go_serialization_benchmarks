package goserbench

import (
	"time"

	"github.com/valyala/fastjson"
)

type FastJSONSerializer struct {
	arena  fastjson.Arena
	object *fastjson.Value
	buf    []byte
}

func (s *FastJSONSerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*A)
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

func (s *FastJSONSerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	v := o.(*A)
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

func NewFastJSONSerializer() *FastJSONSerializer {
	var arena fastjson.Arena
	return &FastJSONSerializer{
		object: arena.NewObject(),
		arena:  arena,
		buf:    make([]byte, 0, 1024),
	}
}
