package goserbench

import (
	"bytes"

	"github.com/hprose/hprose-go"
)

type HproseSerializer struct {
	writer *hprose.Writer
	reader *hprose.Reader
}

func (s *HproseSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*A)
	writer := s.writer
	buf := writer.Stream.(*bytes.Buffer)
	l := buf.Len()
	writer.WriteString(a.Name)
	writer.WriteTime(a.BirthDay)
	writer.WriteString(a.Phone)
	writer.WriteInt64(int64(a.Siblings))
	writer.WriteBool(a.Spouse)
	writer.WriteFloat64(a.Money)
	return buf.Bytes()[l:], nil
}

func (s *HproseSerializer) Unmarshal(d []byte, i interface{}) (err error) {
	o := i.(*A)
	reader := s.reader
	reader.Stream = &hprose.BytesReader{Bytes: d, Pos: 0}
	o.Name, err = reader.ReadString()
	if err != nil {
		return err
	}
	o.BirthDay, err = reader.ReadDateTime()
	if err != nil {
		return err
	}
	o.Phone, err = reader.ReadString()
	if err != nil {
		return err
	}
	o.Siblings, err = reader.ReadInt()
	if err != nil {
		return err
	}
	o.Spouse, err = reader.ReadBool()
	if err != nil {
		return err
	}
	o.Money, err = reader.ReadFloat64()
	return err
}

func newHproseSerializer() *HproseSerializer {
	buf := new(bytes.Buffer)
	reader := hprose.NewReader(buf, true)
	bufw := new(bytes.Buffer)
	writer := hprose.NewWriter(bufw, true)
	return &HproseSerializer{writer: writer, reader: reader}
}
