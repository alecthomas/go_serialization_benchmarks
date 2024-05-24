package gencode

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type GencodeA struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64
}

func (d *GencodeA) Size() (s uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Phone))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{

		t := uint32(d.Siblings)
		t <<= 1
		if d.Siblings < 0 {
			t = ^t
		}
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	s += 24
	return
}
func (d *GencodeA) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	{
		b, err := d.BirthDay.MarshalBinary()
		if err != nil {
			return nil, err
		}
		copy(buf[i+0:], b)
	}
	{
		l := uint64(len(d.Phone))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+15] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+15] = byte(t)
			i++

		}
		copy(buf[i+15:], d.Phone)
		i += l
	}
	{

		t := uint32(d.Siblings)

		t <<= 1
		if d.Siblings < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+15] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+15] = byte(t)
		i++

	}
	{
		if d.Spouse {
			buf[i+15] = 1
		} else {
			buf[i+15] = 0
		}
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Money)))

		buf[i+0+16] = byte(v >> 0)

		buf[i+1+16] = byte(v >> 8)

		buf[i+2+16] = byte(v >> 16)

		buf[i+3+16] = byte(v >> 24)

		buf[i+4+16] = byte(v >> 32)

		buf[i+5+16] = byte(v >> 40)

		buf[i+6+16] = byte(v >> 48)

		buf[i+7+16] = byte(v >> 56)

	}
	return buf[:i+24], nil
}

func (d *GencodeA) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		d.BirthDay.UnmarshalBinary(buf[i+0 : i+0+15])
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+15] & 0x7F)
			for buf[i+15]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+15]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Phone = string(buf[i+15 : i+15+l])
		i += l
	}
	{

		bs := uint8(7)
		t := uint32(buf[i+15] & 0x7F)
		for buf[i+15]&0x80 == 0x80 {
			i++
			t |= uint32(buf[i+15]&0x7F) << bs
			bs += 7
		}
		i++

		d.Siblings = int32(t >> 1)
		if t&1 != 0 {
			d.Siblings = ^d.Siblings
		}

	}
	{
		d.Spouse = buf[i+15] == 1
	}
	{

		v := 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)
		d.Money = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 24, nil
}
