package hprose

import (
	"bytes"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/hprose/hprose-go"
)

type HproseCodec struct {
	goserbench.NativeSmallStruct

	writer *hprose.Writer
	reader *hprose.Reader
}

func (s *HproseCodec) Marshal(a *goserbench.SmallStruct) ([]byte, error) {
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

func (s *HproseCodec) Unmarshal(d []byte, o *goserbench.SmallStruct) (err error) {
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

func NewHproseCodec() goserbench.Codec[goserbench.SmallStruct] {
	buf := new(bytes.Buffer)
	reader := hprose.NewReader(buf, true)
	bufw := new(bytes.Buffer)
	writer := hprose.NewWriter(bufw, true)
	return &HproseCodec{writer: writer, reader: reader}
}
