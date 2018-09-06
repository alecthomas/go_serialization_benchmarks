package goserbench

import (
	"fmt"
	"math/rand"
	"testing"
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate(new func() Object) []Object {
	a := make([]Object, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = new()
	}
	return a
}

func bench(s Serializer, randNew func() Object, validate bool) *serializeBenchResault {
	ret := &serializeBenchResault{
		serializer: s,
	}
	data := generate(randNew)
	ser := make([][]byte, len(data))
	for i, d := range data {
		var buf []byte
		buf, ret.merr = s.Marshal(d)
		if validate && ret.merr != nil {
			return ret
		}
		t := make([]byte, len(buf))
		copy(t, buf)
		ser[i] = t
	}
	var sumlen int
	ret.marshalResult = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf, _ := s.Marshal(data[rand.Intn(len(data))])
			sumlen += len(buf)
		}
	})
	ret.length = sumlen
	o := randNew()
	datalen := len(data)
	ret.unmarshalResult = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := rand.Intn(datalen)
			o.Reset()
			err := s.Unmarshal(ser[n], o)
			if err != nil {
				ret.unmerr = fmt.Errorf("%s failed to unmarshal: %s", s, err)
				b.Error(err)
				return
			}
			if validate {
				if ok, err := o.AssertEqual(data[n]); !ok {
					ret.unmerr = fmt.Errorf("%s failed to unmarshal: %s", s, err)
					b.Error(err)
					return
				}
			}
		}
	})
	return ret
}
