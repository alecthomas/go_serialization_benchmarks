package baseline

import (
	"unsafe"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

// maxSmallStructSerializeSize is the max size of a small struct serialized
// with the baseline serializer.
const maxSmallStructSerializeSize = 8 + 8 + 4 + 1 + goserbench.MaxSmallStructPhoneSize + goserbench.MaxSmallStructNameSize

// appendBool appends a bool to b.
func appendBool(b []byte, v bool) []byte {
	if v {
		return append(b, 1)
	} else {
		return append(b, 0)
	}
}

// getBool reads the next bool from b.
func getBool(b []byte) bool {
	if b[0] == 0 {
		return false
	} else {
		return true
	}
}

// unsafeStringToSlice converts a string to a byte slice in an unsafe way
// (modifications to the byte slice modify the string).
func unsafeStringToSlice(v string) []byte {
	return unsafe.Slice(unsafe.StringData(v), len(v))
}

// unsafeSliceToString converts a byte slice to a string in an unsafe way
// (modifications to the byte slice modify the string).
func unsafeSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var (
	boolZeroSlice = []byte{0}
	boolOneSlice  = []byte{1}
)

// boolToByteSlice returns a byte slice with one element, either 0 or 1.
func boolToByteSlice(b bool) []byte {
	s := boolZeroSlice
	if b {
		s = boolOneSlice
	}
	return s
}

// simpleBufferWriter is a very simple bytes writer that allows aliasing the
// writing buffer b.
type simpleBufferWriter struct {
	b []byte
}

func (sbw *simpleBufferWriter) Write(b []byte) (int, error) {
	sbw.b = append(sbw.b, b...)
	return len(b), nil
}

func (sbw *simpleBufferWriter) Reset() {
	sbw.b = sbw.b[:0]
}

func (sbw *simpleBufferWriter) Bytes() []byte {
	return sbw.b
}
