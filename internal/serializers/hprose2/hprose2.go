package hprose2

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	hprose2 "github.com/hprose/hprose-golang/io"
)

type Hprose2Codec struct {
	goserbench.NativeSmallStruct

	writer *hprose2.Writer
	reader *hprose2.Reader
}

func (s Hprose2Codec) Marshal(a *goserbench.SmallStruct) ([]byte, error) {
	writer := s.writer
	writer.Clear()
	writer.WriteString(a.Name)
	writer.WriteTime(&a.BirthDay)
	writer.WriteString(a.Phone)
	writer.WriteInt(int64(a.Siblings))
	writer.WriteBool(a.Spouse)
	writer.WriteFloat(a.Money, 64)
	return writer.Bytes(), nil
}

func (s Hprose2Codec) Unmarshal(d []byte, o *goserbench.SmallStruct) error {
	reader := s.reader
	reader.Init(d)
	o.Name = reader.ReadString()
	o.BirthDay = reader.ReadTime()
	o.Phone = reader.ReadString()
	o.Siblings = int(reader.ReadInt())
	o.Spouse = reader.ReadBool()
	o.Money = reader.ReadFloat64()
	return nil
}

func NewHProse2Codec() goserbench.Codec[goserbench.SmallStruct] {
	writer := hprose2.NewWriter(true)
	reader := hprose2.NewReader(nil, true)
	return Hprose2Codec{writer: writer, reader: reader}
}
