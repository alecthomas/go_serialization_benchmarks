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

type GencodeUnsafeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64
}

func (d *GencodeUnsafeA) Size() (s uint64) {

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
	s += 17
	return
}
func (d *GencodeUnsafeA) Marshal(buf []byte) ([]byte, error) {
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

		*(*int64)(unsafe.Pointer(&buf[i+0])) = d.BirthDay

	}
	{
		l := uint64(len(d.Phone))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Phone)
		i += l
	}
	{

		t := uint32(d.Siblings)

		t <<= 1
		if d.Siblings < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+8] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+8] = byte(t)
		i++

	}
	{
		if d.Spouse {
			buf[i+8] = 1
		} else {
			buf[i+8] = 0
		}
	}
	{

		*(*float64)(unsafe.Pointer(&buf[i+9])) = d.Money

	}
	return buf[:i+17], nil
}

func (d *GencodeUnsafeA) Unmarshal(buf []byte) (uint64, error) {
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

		d.BirthDay = *(*int64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Phone = string(buf[i+8 : i+8+l])
		i += l
	}
	{

		bs := uint8(7)
		t := uint32(buf[i+8] & 0x7F)
		for buf[i+8]&0x80 == 0x80 {
			i++
			t |= uint32(buf[i+8]&0x7F) << bs
			bs += 7
		}
		i++

		d.Siblings = int32(t >> 1)
		if t&1 != 0 {
			d.Siblings = ^d.Siblings
		}

	}
	{
		d.Spouse = buf[i+8] == 1
	}
	{

		d.Money = *(*float64)(unsafe.Pointer(&buf[i+9]))

	}
	return i + 17, nil
}
