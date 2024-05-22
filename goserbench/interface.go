package goserbench

import "time"

// SmallStruct is a test structure of small size.
type SmallStruct struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

// Serializer is the main interface for the serializer benchmarks.
type Serializer interface {
	// Marshal should marshal the given object (a *A) into a byte slice.
	// The byte slice may be used across invocations of Marshal().
	Marshal(o interface{}) ([]byte, error)

	// Unmarshal should unmarshal the given byte slice into a *A object.
	Unmarshal(d []byte, o interface{}) error
}

// SerializerTimePrecision is a serializer that specifies the max precision that
// a time.Time is encodable. This will be used to truncate time fields.
type SerializerTimePrecision interface {
	// TimePrecision is the max precision a time.Time may be encoded.  When
	// greater than zero, all time.Time fields are truncated down to this
	// precision.
	TimePrecision() time.Duration
}

// SerializerEnforcesTimezone is a serializer that enforces a specific timezone
// when marshalling/unmarshalling time.Time fields.
type SerializerEnforcesTimezone interface {
	// ForcesUTC is true when the serializes forces a UTC timezone.
	ForcesUTC() bool
}

// SerializerLimitsFloat64Precision is a serializer that enforces a maximum
// precision when marshalling/unmarshalling float64 fields.
type SerializerLimitsFloat64Precision interface {
	// FractionalDigits returns the max number of fractional digits that
	// the serializer may encode.
	ReduceFloat64Precision() uint
}
