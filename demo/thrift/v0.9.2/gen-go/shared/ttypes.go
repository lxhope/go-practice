// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package shared

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type SharedStruct struct {
	Key   int32  `thrift:"key,1" json:"key"`
	Value string `thrift:"value,2" json:"value"`
}

func NewSharedStruct() *SharedStruct {
	return &SharedStruct{}
}

func (p *SharedStruct) GetKey() int32 {
	return p.Key
}

func (p *SharedStruct) GetValue() string {
	return p.Value
}
func (p *SharedStruct) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SharedStruct) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *SharedStruct) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *SharedStruct) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SharedStruct"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *SharedStruct) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Key)); err != nil {
		return fmt.Errorf("%T.key (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:key: %s", p, err)
	}
	return err
}

func (p *SharedStruct) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:value: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Value)); err != nil {
		return fmt.Errorf("%T.value (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:value: %s", p, err)
	}
	return err
}

func (p *SharedStruct) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SharedStruct(%+v)", *p)
}