package goserbench

import (
	"bytes"
	"fmt"
	"github.com/linkedin/goavro"
	"time"
)

type AvroA struct {
	record *goavro.Record
	codec  goavro.Codec
}

var (
	avroSchemaJSON = `
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
)

func NewAvroA() *AvroA {
	rec, err := goavro.NewRecord(goavro.RecordSchema(avroSchemaJSON))
	if err != nil {
		fmt.Printf("Avro record initialization error: %v\n", err)
		return nil
	}
	codec, err := goavro.NewCodec(avroSchemaJSON)
	if err != nil {
		fmt.Printf("Avro codec initialization error: %v\n", err)
		return nil
	}
	return &AvroA{record: rec, codec: codec}
}

func (a *AvroA) Marshal(o interface{}) []byte {
	object := o.(*A)
	a.record.Set("name", object.Name)
	a.record.Set("birthday", int64(object.BirthDay.UnixNano()))
	a.record.Set("phone", object.Phone)
	a.record.Set("siblings", int32(object.Siblings))
	a.record.Set("spouse", object.Spouse)
	a.record.Set("money", object.Money)
	b := new(bytes.Buffer)
	if err := a.codec.Encode(b, a.record); err != nil {
		fmt.Printf("Avro encoding error: %v\n", err)
		return []byte("")
	}
	return b.Bytes()
}

func (a *AvroA) Unmarshal(d []byte, o interface{}) error {
	object := o.(*A)
	b := bytes.NewBuffer(d)
	i, err := a.codec.Decode(b)
	if err != nil {
		fmt.Printf("Avro decoding error: %v\n", err)
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
