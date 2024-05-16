package goserbench

import (
	"bytes"
	"time"

	goavro "github.com/linkedin/goavro"
	goavro2 "github.com/linkedin/goavro/v2"
)

type AvroA struct {
	record *goavro.Record
	codec  goavro.Codec
}

type Avro2Txt struct {
	codec *goavro2.Codec
}

type Avro2Bin struct {
	codec *goavro2.Codec
}

var avroSchemaJSON = `
		{
		  "type": "record",
		  "name": "AvroA",
		  "doc:": "Schema for encoding/decoding sample message",
		  "namespace": "com.example",
		  "fields": [
		    {
		      "name": "name",
		      "type": "string"
		    },
		    {
		      "name": "birthday",
		      "type": "long"
		    },
		    {
		      "name": "phone",
		      "type": "string"
		    },
		    {
		      "name": "siblings",
		      "type": "int"
		    },
		    {
		      "name": "spouse",
		      "type": "boolean"
		    },
		    {
		      "name": "money",
		      "type": "double"
		    }
		  ]
		}
	`

func NewAvroA() Serializer {
	rec, err := goavro.NewRecord(goavro.RecordSchema(avroSchemaJSON))
	if err != nil {
		panic(err)
	}
	codec, err := goavro.NewCodec(avroSchemaJSON)
	if err != nil {
		panic(err)
	}
	return &AvroA{record: rec, codec: codec}
}

func (a *AvroA) Marshal(o interface{}) ([]byte, error) {
	object := o.(*A)
	a.record.Set("name", object.Name)
	a.record.Set("birthday", int64(object.BirthDay.UnixNano()))
	a.record.Set("phone", object.Phone)
	a.record.Set("siblings", int32(object.Siblings))
	a.record.Set("spouse", object.Spouse)
	a.record.Set("money", object.Money)
	b := new(bytes.Buffer)
	if err := a.codec.Encode(b, a.record); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (a *AvroA) Unmarshal(d []byte, o interface{}) error {
	object := o.(*A)
	b := bytes.NewBuffer(d)
	i, err := a.codec.Decode(b)
	if err != nil {
		return err
	}
	rec := i.(*goavro.Record)
	temp, _ := rec.Get("name")
	object.Name = temp.(string)
	temp, _ = rec.Get("birthday")
	object.BirthDay = time.Unix(0, temp.(int64))
	temp, _ = rec.Get("phone")
	object.Phone = temp.(string)
	temp, _ = rec.Get("siblings")
	object.Siblings = int(temp.(int32))
	temp, _ = rec.Get("spouse")
	object.Spouse = temp.(bool)
	temp, _ = rec.Get("money")
	object.Money = temp.(float64)
	return nil
}

func (a *AvroA) String() string {
	return "GoAvro"
}

func avroMarshal(o interface{}, marshalFunc func([]byte, interface{}) ([]byte, error)) ([]byte, error) {
	object := o.(*A)
	return marshalFunc(nil, map[string]interface{}{
		"name":     object.Name,
		"birthday": int64(object.BirthDay.UnixNano()),
		"phone":    object.Phone,
		"siblings": int32(object.Siblings),
		"spouse":   object.Spouse,
		"money":    object.Money,
	})
}

func avroUnmarshal(d []byte, o interface{}, unmarshalFunc func([]byte) (interface{}, []byte, error)) error {
	object := o.(*A)
	r, _, err := unmarshalFunc(d)
	if err != nil {
		return err
	}
	m := r.(map[string]interface{})
	object.Name = m["name"].(string)
	object.BirthDay = time.Unix(0, m["birthday"].(int64))
	object.Phone = m["phone"].(string)
	object.Siblings = int(m["siblings"].(int32))
	object.Spouse = m["spouse"].(bool)
	object.Money = m["money"].(float64)
	return nil
}

func NewAvro2Txt() Serializer {
	codec, err := goavro2.NewCodec(avroSchemaJSON)
	if err != nil {
		panic(err)
	}
	return &Avro2Txt{codec: codec}
}

func (a *Avro2Txt) Marshal(o interface{}) ([]byte, error) {
	return avroMarshal(o, a.codec.TextualFromNative)
}

func (a *Avro2Txt) Unmarshal(d []byte, o interface{}) error {
	return avroUnmarshal(d, o, a.codec.NativeFromTextual)
}

func (a *Avro2Txt) String() string {
	return "GoAvro2Text"
}

func NewAvro2Bin() Serializer {
	codec, err := goavro2.NewCodec(avroSchemaJSON)
	if err != nil {
		panic(err)
	}
	return &Avro2Bin{codec: codec}
}

func (a *Avro2Bin) Marshal(o interface{}) ([]byte, error) {
	return avroMarshal(o, a.codec.BinaryFromNative)
}

func (a *Avro2Bin) Unmarshal(d []byte, o interface{}) error {
	return avroUnmarshal(d, o, a.codec.NativeFromBinary)
}

func (a *Avro2Bin) String() string {
	return "GoAvro2Binary"
}
