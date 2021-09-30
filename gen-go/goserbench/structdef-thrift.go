// Code generated by Thrift Compiler (0.15.0). DO NOT EDIT.

package goserbench

import (
	"bytes"
	"context"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - Name
//  - BirthDay
//  - Phone
//  - Siblings
//  - Spouse
//  - Money
type ThriftStructA struct {
  Name string `thrift:"name,1,required" db:"name" json:"name"`
  BirthDay int64 `thrift:"birthDay,2,required" db:"birthDay" json:"birthDay"`
  Phone string `thrift:"phone,3,required" db:"phone" json:"phone"`
  Siblings int32 `thrift:"siblings,4,required" db:"siblings" json:"siblings"`
  Spouse bool `thrift:"spouse,5,required" db:"spouse" json:"spouse"`
  Money float64 `thrift:"money,6,required" db:"money" json:"money"`
}

func NewThriftStructA() *ThriftStructA {
  return &ThriftStructA{}
}


func (p *ThriftStructA) GetName() string {
  return p.Name
}

func (p *ThriftStructA) GetBirthDay() int64 {
  return p.BirthDay
}

func (p *ThriftStructA) GetPhone() string {
  return p.Phone
}

func (p *ThriftStructA) GetSiblings() int32 {
  return p.Siblings
}

func (p *ThriftStructA) GetSpouse() bool {
  return p.Spouse
}

func (p *ThriftStructA) GetMoney() float64 {
  return p.Money
}
func (p *ThriftStructA) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetName bool = false;
  var issetBirthDay bool = false;
  var issetPhone bool = false;
  var issetSiblings bool = false;
  var issetSpouse bool = false;
  var issetMoney bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetName = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetBirthDay = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetPhone = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
        issetSiblings = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField5(ctx, iprot); err != nil {
          return err
        }
        issetSpouse = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.DOUBLE {
        if err := p.ReadField6(ctx, iprot); err != nil {
          return err
        }
        issetMoney = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetName{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Name is not set"));
  }
  if !issetBirthDay{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BirthDay is not set"));
  }
  if !issetPhone{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Phone is not set"));
  }
  if !issetSiblings{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Siblings is not set"));
  }
  if !issetSpouse{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Spouse is not set"));
  }
  if !issetMoney{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Money is not set"));
  }
  return nil
}

func (p *ThriftStructA)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *ThriftStructA)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.BirthDay = v
}
  return nil
}

func (p *ThriftStructA)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Phone = v
}
  return nil
}

func (p *ThriftStructA)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Siblings = v
}
  return nil
}

func (p *ThriftStructA)  ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(ctx); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Spouse = v
}
  return nil
}

func (p *ThriftStructA)  ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(ctx); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.Money = v
}
  return nil
}

func (p *ThriftStructA) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ThriftStructA"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
    if err := p.writeField5(ctx, oprot); err != nil { return err }
    if err := p.writeField6(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ThriftStructA) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err) }
  return err
}

func (p *ThriftStructA) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "birthDay", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:birthDay: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.BirthDay)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.birthDay (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:birthDay: ", p), err) }
  return err
}

func (p *ThriftStructA) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "phone", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:phone: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Phone)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.phone (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:phone: ", p), err) }
  return err
}

func (p *ThriftStructA) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "siblings", thrift.I32, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:siblings: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Siblings)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.siblings (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:siblings: ", p), err) }
  return err
}

func (p *ThriftStructA) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "spouse", thrift.BOOL, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:spouse: ", p), err) }
  if err := oprot.WriteBool(ctx, bool(p.Spouse)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.spouse (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:spouse: ", p), err) }
  return err
}

func (p *ThriftStructA) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "money", thrift.DOUBLE, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:money: ", p), err) }
  if err := oprot.WriteDouble(ctx, float64(p.Money)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.money (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:money: ", p), err) }
  return err
}

func (p *ThriftStructA) Equals(other *ThriftStructA) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Name != other.Name { return false }
  if p.BirthDay != other.BirthDay { return false }
  if p.Phone != other.Phone { return false }
  if p.Siblings != other.Siblings { return false }
  if p.Spouse != other.Spouse { return false }
  if p.Money != other.Money { return false }
  return true
}

func (p *ThriftStructA) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ThriftStructA(%+v)", *p)
}

