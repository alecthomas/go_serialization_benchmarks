package goserbench

import (
	"math"
	"reflect"
	"unsafe"

	"github.com/ymz-ncnk/musgo/errs"
)

// MarshalMUSUnsafe fills buf with the MUS encoding of v.
func (v MusgoA) MarshalMUSUnsafe(buf []byte) int {
	i := 0
	{
		length := len(v.Name)
		{
			uv := uint64(length)
			if length < 0 {
				uv = ^(uv << 1)
			} else {
				uv = uv << 1
			}
			{
				for uv >= 0x80 {
					buf[i] = byte(uv) | 0x80
					uv >>= 7
					i++
				}
				buf[i] = byte(uv)
				i++
			}
		}
		i += copy(buf[i:], v.Name)
	}
	{
		{
			*(*int64)(unsafe.Pointer(&buf[i])) = v.BirthDay
			i += 8
		}
	}
	{
		length := len(v.Phone)
		{
			uv := uint64(length)
			if length < 0 {
				uv = ^(uv << 1)
			} else {
				uv = uv << 1
			}
			{
				for uv >= 0x80 {
					buf[i] = byte(uv) | 0x80
					uv >>= 7
					i++
				}
				buf[i] = byte(uv)
				i++
			}
		}
		i += copy(buf[i:], v.Phone)
	}
	{
		uv := uint32(v.Siblings)
		if v.Siblings < 0 {
			uv = ^(uv << 1)
		} else {
			uv = uv << 1
		}
		{
			for uv >= 0x80 {
				buf[i] = byte(uv) | 0x80
				uv >>= 7
				i++
			}
			buf[i] = byte(uv)
			i++
		}
	}
	{
		if v.Spouse {
			buf[i] = 0x01
		} else {
			buf[i] = 0x00
		}
		i++
	}
	{
		uv := math.Float64bits(float64(v.Money))
		{
			*(*uint64)(unsafe.Pointer(&buf[i])) = uv
			i += 8
		}
	}
	return i
}

// UnmarshalMUSUnsafe parses the MUS-encoded buf, and sets the result to *v.
func (v *MusgoA) UnmarshalMUSUnsafe(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				}
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l == 9 && b > 1 {
						return i, errs.ErrOverflow
					}
					if b < 0x80 {
						uv = uv | uint64(b)<<shift
						done = true
						i += l + 1
						break
					}
					uv = uv | uint64(b&0x7F)<<shift
					shift += 7
				}
				if !done {
					return i, errs.ErrSmallBuf
				}
			}
			if uv&1 == 1 {
				uv = ^(uv >> 1)
			} else {
				uv = uv >> 1
			}
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if len(buf) < i+length {
			return i, errs.ErrSmallBuf
		}
		content := buf[i : i+length]
		slcHeader := (*reflect.SliceHeader)(unsafe.Pointer(&content))
		strHeader := (*reflect.StringHeader)(unsafe.Pointer(&v.Name))
		strHeader.Data = slcHeader.Data
		strHeader.Len = slcHeader.Len
		i += length
	}
	if err != nil {
		return i, errs.NewFieldError("Name", err)
	}
	{
		{
			if len(buf) < 8 {
				return i, errs.ErrSmallBuf
			}
			v.BirthDay = *(*int64)(unsafe.Pointer(&buf[i]))
			i += 8
		}
	}
	if err != nil {
		return i, errs.NewFieldError("BirthDay", err)
	}
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				}
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l == 9 && b > 1 {
						return i, errs.ErrOverflow
					}
					if b < 0x80 {
						uv = uv | uint64(b)<<shift
						done = true
						i += l + 1
						break
					}
					uv = uv | uint64(b&0x7F)<<shift
					shift += 7
				}
				if !done {
					return i, errs.ErrSmallBuf
				}
			}
			if uv&1 == 1 {
				uv = ^(uv >> 1)
			} else {
				uv = uv >> 1
			}
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if len(buf) < i+length {
			return i, errs.ErrSmallBuf
		}
		content := buf[i : i+length]
		slcHeader := (*reflect.SliceHeader)(unsafe.Pointer(&content))
		strHeader := (*reflect.StringHeader)(unsafe.Pointer(&v.Phone))
		strHeader.Data = slcHeader.Data
		strHeader.Len = slcHeader.Len
		i += length
	}
	if err != nil {
		return i, errs.NewFieldError("Phone", err)
	}
	{
		var uv uint32
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			}
			shift := 0
			done := false
			for l, b := range buf[i:] {
				if l == 4 && b > 15 {
					return i, errs.ErrOverflow
				}
				if b < 0x80 {
					uv = uv | uint32(b)<<shift
					done = true
					i += l + 1
					break
				}
				uv = uv | uint32(b&0x7F)<<shift
				shift += 7
			}
			if !done {
				return i, errs.ErrSmallBuf
			}
		}
		if uv&1 == 1 {
			uv = ^(uv >> 1)
		} else {
			uv = uv >> 1
		}
		v.Siblings = int32(uv)
	}
	if err != nil {
		return i, errs.NewFieldError("Siblings", err)
	}
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		}
		if buf[i] == 0x01 {
			v.Spouse = true
			i++
		} else if buf[i] == 0x00 {
			v.Spouse = false
			i++
		} else {
			err = errs.ErrWrongByte
		}
	}
	if err != nil {
		return i, errs.NewFieldError("Spouse", err)
	}
	{
		var uv uint64
		{
			if len(buf) < 8 {
				return i, errs.ErrSmallBuf
			}
			uv = *(*uint64)(unsafe.Pointer(&buf[i]))
			i += 8
		}
		v.Money = float64(math.Float64frombits(uv))
	}
	if err != nil {
		return i, errs.NewFieldError("Money", err)
	}
	return i, err
}

// SizeMUSUnsafe returns the size of the MUS-encoded v.
func (v MusgoA) SizeMUSUnsafe() int {
	size := 0
	{
		length := len(v.Name)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					uv >>= 7
					size++
				}
				size++
			}
		}
		size += len(v.Name)
	}
	{
		{
			_ = v.BirthDay
			size += 8
		}
	}
	{
		length := len(v.Phone)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					uv >>= 7
					size++
				}
				size++
			}
		}
		size += len(v.Phone)
	}
	{
		uv := uint32(v.Siblings<<1) ^ uint32(v.Siblings>>31)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}
	}
	{
		_ = v.Spouse
		size++
	}
	{
		{
			_ = v.Money
			size += 8
		}

	}
	return size
}
