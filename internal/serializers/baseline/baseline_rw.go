package baseline

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

// bytesMarshable is an interface for an io.Writer that can also generate a byte
// slice and reset its size (i.e. a subset of *bytes.Buffer).
type bytesMarshable interface {
	io.Writer
	Bytes() []byte
	Reset()
}

// alwaysNilWriter is always nil. It is used to ensure the compiler does not
// devirtualize calls to bytes.Buffer.Write().
var alwaysNilWriter bytesMarshable

// awlaysNilReader is always nil. It is used to ensure the compiler does not
// devirtualize calls to bytes.Reader.Read().
var awlaysNilReader io.Reader

// BaselineReaderWriter is a hand-written baseline serializer that reads/writes
// using standard io.Reader and io.Writer interfaces.
//
// While goserbench utlimately generates a []byte as a result of Marshal (and
// reads from a []byte as well), this verifies the performance of using the
// generic io interface, which could be used to write to other streams directly
// (e.g. file or socket).
type BaselineReaderWriter struct {
	// Store r to avoid this allocation. This is meant to simulate the fact
	// that Unmarshal() receives a Reader directly (instead of a []byte).
	r *bytes.Reader

	// Store w to avoid this allocation. This is meant to simulate the fact
	// that Marshal() receives a Writer directly (instead of a []byte).
	w *simpleBufferWriter

	// aux buffer to read native values (int64, etc). Assume that the caller
	// can amortize this allocation somehow.
	auxBuf []byte

	// aux buffer to read strings.
	strBuf []byte
}

func (b *BaselineReaderWriter) Marshal(o interface{}) ([]byte, error) {
	// Create a new bytes buffer. We do it like this (referencing
	// alwaysNilWriter first) to avoid the compiler de-virtualizing and
	// inlining the calls to w.Write() directly into bytes.Buffer.Write()
	// calls, which simulates the worst case scenario of the overhead in
	// using Write() calls.
	w := alwaysNilWriter
	if w == nil {
		// Create with the pre-sized buffer. We assume the caller will
		// be able to hint to the writer either the final size or at
		// least an upper bound of what will be written, and that the
		// writer will be able to use that hint (e.g. in case of a file,
		// it could grow the file to allow adding that many bytes,
		// create a kernel buffer, memory map the file, etc).
		b.w.b = make([]byte, 0, maxSmallStructSerializeSize)
		w = b.w
	}

	a := o.(*goserbench.SmallStruct)
	extra := b.auxBuf[:0]
	if _, err := w.Write(binary.LittleEndian.AppendUint64(extra, uint64(a.BirthDay.UnixNano()))); err != nil {
		return nil, err
	}
	if _, err := w.Write(binary.LittleEndian.AppendUint64(extra, math.Float64bits(a.Money))); err != nil {
		return nil, err
	}
	if _, err := w.Write(binary.LittleEndian.AppendUint32(extra, uint32(a.Siblings))); err != nil {
		return nil, err
	}
	if _, err := w.Write([]byte(a.Name)); err != nil {
		return nil, err
	}
	if _, err := w.Write([]byte(a.Phone)); err != nil {
		return nil, err
	}
	if _, err := w.Write(boolToByteSlice(a.Spouse)); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func (b *BaselineReaderWriter) Unmarshal(d []byte, o interface{}) error {
	// Create reader with passed buffer. We do it like this (referencing
	// alwaysNilReader first) to ensure the compiler does not devirtualize
	// and inline the Read() calls direcly to the *bytes.Buffer.
	r := awlaysNilReader
	if r == nil {
		b.r.Reset(d)
		r = b.r
	}

	a := o.(*goserbench.SmallStruct)
	aux := b.auxBuf
	strs := b.strBuf

	if _, err := io.ReadFull(r, aux[:8]); err != nil {
		return err
	}
	a.BirthDay = time.Unix(0, int64(binary.LittleEndian.Uint64(aux)))

	if _, err := io.ReadFull(r, aux[:8]); err != nil {
		return err
	}
	a.Money = math.Float64frombits(binary.LittleEndian.Uint64(aux))

	if _, err := io.ReadFull(r, aux[:4]); err != nil {
		return err
	}
	a.Siblings = int(binary.LittleEndian.Uint32(aux))

	if _, err := io.ReadFull(r, strs[:goserbench.MaxSmallStructNameSize]); err != nil {
		return err
	}
	a.Name = string(strs[:goserbench.MaxSmallStructNameSize])

	if _, err := io.ReadFull(r, strs[:goserbench.MaxSmallStructPhoneSize]); err != nil {
		return err
	}
	a.Phone = string(strs[:goserbench.MaxSmallStructPhoneSize])

	if _, err := io.ReadFull(r, aux[:1]); err != nil {
		return err
	}
	a.Spouse = aux[0] != 0
	return nil
}

func NewBaselineReaderWriter() goserbench.Serializer {
	return &BaselineReaderWriter{
		r:      bytes.NewReader(nil),
		w:      &simpleBufferWriter{},
		auxBuf: make([]byte, 8),
		strBuf: make([]byte, max(goserbench.MaxSmallStructNameSize, goserbench.MaxSmallStructPhoneSize)),
	}
}

// BaselineReaderWriterUnsafeReuse is a hand-written baseline serializer that
// reads/writes using standard io.Reader and io.Writer interfaces. It reuses
// buffers and unsafe strings.
type BaselineReaderWriterUnsafeReuse struct {
	// w saves the buffer for writing.
	w bytesMarshable

	// r avoids having to allocate a *bytes.Reader reference.
	r *bytes.Reader

	// aux buffer to read native values (int64, etc).
	auxBuf []byte

	// aux buffer to read strings.
	strBuf []byte
}

func (b *BaselineReaderWriterUnsafeReuse) Marshal(o interface{}) ([]byte, error) {
	// Get a reference to the writer and reset its buffer. We do it like
	// this (referencing alwaysNilWriter first) to ensure the compiler does
	// not devirtualize and inline the Write() calls direcly to the
	// *bytes.Buffer.
	w := alwaysNilWriter
	if w == nil {
		w = b.w
		w.Reset()
	}

	a := o.(*goserbench.SmallStruct)
	extra := b.auxBuf[:0]
	if _, err := w.Write(binary.LittleEndian.AppendUint64(extra, uint64(a.BirthDay.UnixNano()))); err != nil {
		return nil, err
	}
	if _, err := w.Write(binary.LittleEndian.AppendUint64(extra, math.Float64bits(a.Money))); err != nil {
		return nil, err
	}
	if _, err := w.Write(binary.LittleEndian.AppendUint32(extra, uint32(a.Siblings))); err != nil {
		return nil, err
	}
	if _, err := w.Write(unsafeStringToSlice(a.Name)); err != nil {
		return nil, err
	}
	if _, err := w.Write(unsafeStringToSlice(a.Phone)); err != nil {
		return nil, err
	}
	if _, err := w.Write(boolToByteSlice(a.Spouse)); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func (b *BaselineReaderWriterUnsafeReuse) Unmarshal(d []byte, o interface{}) error {
	// Reset reader to passed buffer. We do it like this (referencing
	// alwaysNilReader first) to ensure the compiler does not devirtualize
	// and inline the Read() calls direcly to the *bytes.Buffer.
	r := awlaysNilReader
	if r == nil {
		b.r.Reset(d)
		r = b.r
	}

	a := o.(*goserbench.SmallStruct)
	aux := b.auxBuf
	strs := b.strBuf

	if _, err := io.ReadFull(r, aux[:8]); err != nil {
		return err
	}
	a.BirthDay = time.Unix(0, int64(binary.LittleEndian.Uint64(aux)))

	if _, err := io.ReadFull(r, aux[:8]); err != nil {
		return err
	}
	a.Money = math.Float64frombits(binary.LittleEndian.Uint64(aux))

	if _, err := io.ReadFull(r, aux[:4]); err != nil {
		return err
	}
	a.Siblings = int(binary.LittleEndian.Uint32(aux))

	if _, err := io.ReadFull(r, strs[:goserbench.MaxSmallStructNameSize]); err != nil {
		return err
	}
	a.Name = unsafeSliceToString(strs[:goserbench.MaxSmallStructNameSize])
	strs = strs[goserbench.MaxSmallStructNameSize:]

	if _, err := io.ReadFull(r, strs[:goserbench.MaxSmallStructPhoneSize]); err != nil {
		return err
	}
	a.Phone = unsafeSliceToString(strs[:goserbench.MaxSmallStructPhoneSize])

	if _, err := io.ReadFull(r, aux[:1]); err != nil {
		return err
	}
	a.Spouse = aux[0] != 0
	return nil
}

func NewBaselineReaderWriterUnsafeReuse() goserbench.Serializer {
	return &BaselineReaderWriterUnsafeReuse{
		w:      bytes.NewBuffer(make([]byte, 0, maxSmallStructSerializeSize)),
		r:      bytes.NewReader(nil),
		auxBuf: make([]byte, 8),
		strBuf: make([]byte, goserbench.MaxSmallStructNameSize+goserbench.MaxSmallStructPhoneSize),
	}
}
