// Code generated by protoc-gen-go.
// source: kodec.proto
// DO NOT EDIT!

package kodec

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Msg_Type int32

const (
	Msg_UNKNOWN Msg_Type = 0
	Msg_TXT     Msg_Type = 1
	Msg_IMG     Msg_Type = 2
	Msg_VOICE   Msg_Type = 3
	Msg_SYS     Msg_Type = 4
	Msg_CARD    Msg_Type = 5
	Msg_GIF     Msg_Type = 6
	Msg_NEWS    Msg_Type = 7
)

var Msg_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "TXT",
	2: "IMG",
	3: "VOICE",
	4: "SYS",
	5: "CARD",
	6: "GIF",
	7: "NEWS",
}
var Msg_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"TXT":     1,
	"IMG":     2,
	"VOICE":   3,
	"SYS":     4,
	"CARD":    5,
	"GIF":     6,
	"NEWS":    7,
}

func (x Msg_Type) Enum() *Msg_Type {
	p := new(Msg_Type)
	*p = x
	return p
}
func (x Msg_Type) String() string {
	return proto.EnumName(Msg_Type_name, int32(x))
}
func (x Msg_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Msg_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Msg_Type_value, data, "Msg_Type")
	if err != nil {
		return err
	}
	*x = Msg_Type(value)
	return nil
}

type Cmd_Type int32

const (
	Cmd_UNKNOWN      Cmd_Type = 0
	Cmd_UNAUTHORIZED Cmd_Type = 1
	Cmd_PING         Cmd_Type = 2
	Cmd_VISIT        Cmd_Type = 3
	Cmd_SYNC         Cmd_Type = 4
)

var Cmd_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "UNAUTHORIZED",
	2: "PING",
	3: "VISIT",
	4: "SYNC",
}
var Cmd_Type_value = map[string]int32{
	"UNKNOWN":      0,
	"UNAUTHORIZED": 1,
	"PING":         2,
	"VISIT":        3,
	"SYNC":         4,
}

func (x Cmd_Type) Enum() *Cmd_Type {
	p := new(Cmd_Type)
	*p = x
	return p
}
func (x Cmd_Type) String() string {
	return proto.EnumName(Cmd_Type_name, int32(x))
}
func (x Cmd_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Cmd_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Cmd_Type_value, data, "Cmd_Type")
	if err != nil {
		return err
	}
	*x = Cmd_Type(value)
	return nil
}

type Meta_Type int32

const (
	Meta_UNKNOWN      Meta_Type = 0
	Meta_JOIN_TOK     Meta_Type = 1
	Meta_LEAVE_TOK    Meta_Type = 2
	Meta_DISMISS_TOK  Meta_Type = 3
	Meta_EVENT_CHANGE Meta_Type = 4
	Meta_ADD_FRIEND   Meta_Type = 5
	Meta_DEL_FRIEND   Meta_Type = 6
	Meta_USER_UPDATE  Meta_Type = 7
	Meta_SETUP_TOK    Meta_Type = 8
	Meta_UPDATE_TOK   Meta_Type = 9
	Meta_NEWS         Meta_Type = 10
)

var Meta_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "JOIN_TOK",
	2:  "LEAVE_TOK",
	3:  "DISMISS_TOK",
	4:  "EVENT_CHANGE",
	5:  "ADD_FRIEND",
	6:  "DEL_FRIEND",
	7:  "USER_UPDATE",
	8:  "SETUP_TOK",
	9:  "UPDATE_TOK",
	10: "NEWS",
}
var Meta_Type_value = map[string]int32{
	"UNKNOWN":      0,
	"JOIN_TOK":     1,
	"LEAVE_TOK":    2,
	"DISMISS_TOK":  3,
	"EVENT_CHANGE": 4,
	"ADD_FRIEND":   5,
	"DEL_FRIEND":   6,
	"USER_UPDATE":  7,
	"SETUP_TOK":    8,
	"UPDATE_TOK":   9,
	"NEWS":         10,
}

func (x Meta_Type) Enum() *Meta_Type {
	p := new(Meta_Type)
	*p = x
	return p
}
func (x Meta_Type) String() string {
	return proto.EnumName(Meta_Type_name, int32(x))
}
func (x Meta_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Meta_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Meta_Type_value, data, "Meta_Type")
	if err != nil {
		return err
	}
	*x = Meta_Type(value)
	return nil
}

// bag id = 0
type Msg struct {
	To   *string   `protobuf:"bytes,1,req,name=to" json:"to,omitempty"`
	From *int64    `protobuf:"varint,2,opt,name=from" json:"from,omitempty"`
	Tp   *Msg_Type `protobuf:"varint,3,opt,name=tp,enum=Msg_Type,def=0" json:"tp,omitempty"`
	Desc *string   `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	// d has different meanings for different type:
	// TXT & SYS: text content
	// IMG: image binary
	// VOICE: voice binary
	// CARD: Card binary
	D                []byte  `protobuf:"bytes,5,req,name=d" json:"d,omitempty"`
	Ct               *int64  `protobuf:"varint,6,opt,name=ct" json:"ct,omitempty"`
	Meta             []*Meta `protobuf:"bytes,7,rep,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}

const Default_Msg_Tp Msg_Type = Msg_UNKNOWN

func (m *Msg) GetTo() string {
	if m != nil && m.To != nil {
		return *m.To
	}
	return ""
}

func (m *Msg) GetFrom() int64 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *Msg) GetTp() Msg_Type {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return Default_Msg_Tp
}

func (m *Msg) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func (m *Msg) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *Msg) GetCt() int64 {
	if m != nil && m.Ct != nil {
		return *m.Ct
	}
	return 0
}

func (m *Msg) GetMeta() []*Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// the user card
type Card struct {
	Uid              *int64  `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Icon             *string `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}

func (m *Card) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *Card) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Card) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

// bag id = 1
type Cmd struct {
	Tp               *Cmd_Type `protobuf:"varint,1,req,name=tp,enum=Cmd_Type" json:"tp,omitempty"`
	Ct               *int64    `protobuf:"varint,2,req,name=ct" json:"ct,omitempty"`
	Txt              *string   `protobuf:"bytes,3,opt,name=txt" json:"txt,omitempty"`
	Meta             []*Meta   `protobuf:"bytes,4,rep,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}

func (m *Cmd) GetTp() Cmd_Type {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return 0
}

func (m *Cmd) GetCt() int64 {
	if m != nil && m.Ct != nil {
		return *m.Ct
	}
	return 0
}

func (m *Cmd) GetTxt() string {
	if m != nil && m.Txt != nil {
		return *m.Txt
	}
	return ""
}

func (m *Cmd) GetMeta() []*Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// meta, add invisible payload to message & command
type Meta struct {
	Tp *Meta_Type `protobuf:"varint,1,req,name=tp,enum=Meta_Type" json:"tp,omitempty"`
	// for JOIN_TOK/LEAVE_TOK, it's comma separated uids
	// for event change it's event id
	// for ADD_FRIEND/DEL_FRIEND/USER_UPDATE, it's comma separated uids
	// for USER_UPDATE, it's uid
	// for SETUP_TOK, UPDATE_TOK, it's nothing
	Txt              *string `protobuf:"bytes,2,req,name=txt" json:"txt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}

func (m *Meta) GetTp() Meta_Type {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return 0
}

func (m *Meta) GetTxt() string {
	if m != nil && m.Txt != nil {
		return *m.Txt
	}
	return ""
}

func init() {
	proto.RegisterEnum("Msg_Type", Msg_Type_name, Msg_Type_value)
	proto.RegisterEnum("Cmd_Type", Cmd_Type_name, Cmd_Type_value)
	proto.RegisterEnum("Meta_Type", Meta_Type_name, Meta_Type_value)
}
