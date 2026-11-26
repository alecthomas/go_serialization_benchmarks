package goserbench

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	// MaxSmallStructNameSize is the max size of a name used in the small
	// struct benchmarks.
	MaxSmallStructNameSize = 16

	// MaxSmallStructPhoneSize is the max size of a phone used in the small
	// struct benchmarks.
	MaxSmallStructPhoneSize = 10
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generateSmallStruct() []*SmallStruct {
	a := make([]*SmallStruct, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &SmallStruct{
			Name:     randString(MaxSmallStructNameSize),
			BirthDay: time.Now(),
			Phone:    randString(MaxSmallStructPhoneSize),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

// Benchmarks are the marshal and unmarshal benchmarks for one codec, with the
// codec's native type erased so codecs of different types can share a
// registry.
type Benchmarks struct {
	Marshal   func(b *testing.B)
	Unmarshal func(b *testing.B, validate bool)
}

// NewBenchmarks creates the benchmarks for codecs created by newCodec.
func NewBenchmarks[T any](newCodec func() Codec[T]) Benchmarks {
	return Benchmarks{
		Marshal:   func(b *testing.B) { BenchMarshal(b, newCodec()) },
		Unmarshal: func(b *testing.B, validate bool) { BenchUnmarshal(b, newCodec(), validate) },
	}
}

// serializerCodec adapts a [Serializer] to a [Codec] over [SmallStruct].
type serializerCodec struct {
	NativeSmallStruct
	s Serializer
}

func (c serializerCodec) Marshal(v *SmallStruct) ([]byte, error)   { return c.s.Marshal(v) }
func (c serializerCodec) Unmarshal(d []byte, v *SmallStruct) error { return c.s.Unmarshal(d, v) }

// option returns c as the optional interface I. The adapter does not
// implement the optional interfaces itself, so it is unwrapped to the
// adapted serializer first.
func option[I any](c any) (I, bool) {
	if sc, ok := c.(serializerCodec); ok {
		c = sc.s
	}
	i, ok := c.(I)
	return i, ok
}

// BenchMarshalSmallStruct benchmarks marshalling the [SmallStruct] type.
func BenchMarshalSmallStruct(b *testing.B, s Serializer) {
	b.Helper()
	BenchMarshal[SmallStruct](b, serializerCodec{s: s})
}

// BenchUnmarshalSmallStruct benchmarks unmarshalling the [SmallStruct] type.
// If validate is true, then the unmarshalled struct is verified to be correct
// against the source struct.
func BenchUnmarshalSmallStruct(b *testing.B, s Serializer, validate bool) {
	b.Helper()
	BenchUnmarshal[SmallStruct](b, serializerCodec{s: s}, validate)
}

// BenchMarshal benchmarks marshalling the codec's native type. Test data is
// converted to the native type before timing starts.
func BenchMarshal[T any](b *testing.B, c Codec[T]) {
	b.Helper()
	data := generateSmallStruct()
	native := make([]T, len(data))
	for i, d := range data {
		c.FromSmall(d, &native[i])
	}

	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		o := &native[rand.Intn(len(native))]
		bytes, err := c.Marshal(o)
		if err != nil {
			b.Fatalf("marshal error %s for %#v", err, o)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

// BenchUnmarshal benchmarks unmarshalling the codec's native type. If
// validate is true, then the unmarshalled value is converted back to a
// [SmallStruct] and verified against the source struct.
func BenchUnmarshal[T any](b *testing.B, c Codec[T], validate bool) {
	b.Helper()

	var timePrecision time.Duration
	if stp, ok := option[SerializerTimePrecision](c); ok {
		timePrecision = stp.TimePrecision()
	}
	var forcesUTC bool
	if set, ok := option[SerializerEnforcesTimezone](c); ok {
		forcesUTC = set.ForcesUTC()
	}

	data := generateSmallStruct()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		// Reduce the precision of the Birthday field when the
		// serializer cannot represent time with nanosecond precision.
		if timePrecision > 0 {
			d.BirthDay = d.BirthDay.Truncate(timePrecision)
		}

		// Enforce Timezone when serializer requires it.
		if forcesUTC {
			d.BirthDay = d.BirthDay.UTC()
		}

		// Reduce precision of fractional fields when the serializer
		// cannot represent the full float64 range.
		if slfp, ok := option[SerializerLimitsFloat64Precision](c); ok {
			fracDigits := slfp.ReduceFloat64Precision()
			i, f := math.Modf(d.Money)
			power := math.Pow(10, float64(fracDigits))
			newf := math.Trunc(f*power) / power
			d.Money = float64(i) + newf
		}

		var n T
		c.FromSmall(d, &n)
		o, err := c.Marshal(&n)
		if err != nil {
			b.Fatal(err)
		}
		t := make([]byte, len(o))
		serialSize += copy(t, o)
		ser[i] = t
	}
	o := new(T)
	var got SmallStruct

	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		var zero T
		*o = zero
		err := c.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("unmarshal error %s for %#x / %q", err, ser[n], ser[n])
		}
		// Validate unmarshalled data.
		if validate {
			c.ToSmall(o, &got)
			i := data[n]
			if got.Name != i.Name {
				b.Fatalf("wrong name: got %q, want %q", got.Name, i.Name)
			}
			if got.Phone != i.Phone {
				b.Fatalf("wrong phone: got %q, want %q", got.Phone, i.Phone)
			}
			if got.Siblings != i.Siblings {
				b.Fatalf("wrong siblings: got %v, want %v", got.Siblings, i.Siblings)
			}
			if got.Spouse != i.Spouse {
				b.Fatalf("wrong spouse : got %v, want %v", got.Spouse, i.Spouse)
			}
			if !got.BirthDay.Equal(i.BirthDay) {
				b.Fatalf("wrong birthday: got %v, want %v", got.BirthDay, i.BirthDay)
			}
			if got.Money != i.Money {
				b.Fatalf("wrong money: got %v, want %v", got.Money, i.Money)
			}
		}
	}
}
