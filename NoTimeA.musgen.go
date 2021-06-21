package goserbench

import (
	"math"

	"github.com/ymz-ncnk/musgo/errs"
)

func (v NoTimeA) MarshalMUS(buf []byte) int {
	i := 0
	{
		length := len(v.Name)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
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
		uv := uint64(v.BirthDay<<1) ^ uint64(v.BirthDay>>63)
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
		length := len(v.Phone)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
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
		uv := uint64(v.Siblings<<1) ^ uint64(v.Siblings>>63)
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
		uv = (uv << 32) | (uv >> 32)
		uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
		uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
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
	return i
}

func (v *NoTimeA) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				} else {
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
			}
			uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if len(buf) < i+length {
			return i, errs.ErrSmallBuf
		}
		v.Name = string(buf[i : i+length])
		i += length
	}
	if err != nil {
		return i, errs.NewFieldError("Name", err)
	}
	{
		var uv uint64
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			} else {
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
		}
		uv = (uv >> 1) ^ uint64((int64(uv&1)<<63)>>63)
		v.BirthDay = int64(uv)
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
				} else {
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
			}
			uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if len(buf) < i+length {
			return i, errs.ErrSmallBuf
		}
		v.Phone = string(buf[i : i+length])
		i += length
	}
	if err != nil {
		return i, errs.NewFieldError("Phone", err)
	}
	{
		var uv uint64
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			} else {
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
		}
		uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
		v.Siblings = int(uv)
	}
	if err != nil {
		return i, errs.NewFieldError("Siblings", err)
	}
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		} else {
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
	}
	if err != nil {
		return i, errs.NewFieldError("Spouse", err)
	}
	{
		var uv uint64
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			} else {
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
		}
		uv = (uv << 32) | (uv >> 32)
		uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
		uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
		v.Money = float64(math.Float64frombits(uv))
	}
	if err != nil {
		return i, errs.NewFieldError("Money", err)
	}
	return i, err
}

func (v NoTimeA) SizeMUS() int {
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
		uv := uint64(v.BirthDay<<1) ^ uint64(v.BirthDay>>63)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
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
		uv := uint64(v.Siblings<<1) ^ uint64(v.Siblings>>63)
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
		uv := math.Float64bits(float64(v.Money))
		uv = (uv << 32) | (uv >> 32)
		uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
		uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}
	}
	return size
}
