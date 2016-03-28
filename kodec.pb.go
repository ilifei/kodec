// Code generated by protoc-gen-go.
// source: kodec.proto
// DO NOT EDIT!

/*
Package kodec is a generated protocol buffer package.

It is generated from these files:
	kodec.proto

It has these top-level messages:
	Msg
	Card
	File
	Cmd
	Meta
	Ack
*/
package kodec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

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
	Msg_FILE    Msg_Type = 8
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
	8: "FILE",
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
	"FILE":    8,
}

func (x Msg_Type) Enum() *Msg_Type {
	p := new(Msg_Type)
	*p = x
	return p
}
func (x Msg_Type) String() string {
	return proto.EnumName(Msg_Type_name, int32(x))
}
func (x *Msg_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Msg_Type_value, data, "Msg_Type")
	if err != nil {
		return err
	}
	*x = Msg_Type(value)
	return nil
}
func (Msg_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type File_Type int32

const (
	File_UNKNOWN File_Type = 0
	File_PDF     File_Type = 1
	File_DOC     File_Type = 2
	File_XLS     File_Type = 3
	File_PPT     File_Type = 4
	File_IMG     File_Type = 5
	File_TXT     File_Type = 6
)

var File_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PDF",
	2: "DOC",
	3: "XLS",
	4: "PPT",
	5: "IMG",
	6: "TXT",
}
var File_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"PDF":     1,
	"DOC":     2,
	"XLS":     3,
	"PPT":     4,
	"IMG":     5,
	"TXT":     6,
}

func (x File_Type) Enum() *File_Type {
	p := new(File_Type)
	*p = x
	return p
}
func (x File_Type) String() string {
	return proto.EnumName(File_Type_name, int32(x))
}
func (x *File_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(File_Type_value, data, "File_Type")
	if err != nil {
		return err
	}
	*x = File_Type(value)
	return nil
}
func (File_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

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
func (x *Cmd_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Cmd_Type_value, data, "Cmd_Type")
	if err != nil {
		return err
	}
	*x = Cmd_Type(value)
	return nil
}
func (Cmd_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Meta_Type int32

const (
	Meta_UNKNOWN        Meta_Type = 0
	Meta_JOIN_TOK       Meta_Type = 1
	Meta_LEAVE_TOK      Meta_Type = 2
	Meta_DISMISS_TOK    Meta_Type = 3
	Meta_EVENT_CHANGE   Meta_Type = 4
	Meta_ADD_FRIEND     Meta_Type = 5
	Meta_DEL_FRIEND     Meta_Type = 6
	Meta_USER_UPDATE    Meta_Type = 7
	Meta_SETUP_TOK      Meta_Type = 8
	Meta_UPDATE_TOK     Meta_Type = 9
	Meta_NEWS           Meta_Type = 10
	Meta_LANDING_PHONE  Meta_Type = 11
	Meta_VOIP           Meta_Type = 12
	Meta_VIDEO          Meta_Type = 13
	Meta_CONSULAR_PHONE Meta_Type = 14
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
	11: "LANDING_PHONE",
	12: "VOIP",
	13: "VIDEO",
	14: "CONSULAR_PHONE",
}
var Meta_Type_value = map[string]int32{
	"UNKNOWN":        0,
	"JOIN_TOK":       1,
	"LEAVE_TOK":      2,
	"DISMISS_TOK":    3,
	"EVENT_CHANGE":   4,
	"ADD_FRIEND":     5,
	"DEL_FRIEND":     6,
	"USER_UPDATE":    7,
	"SETUP_TOK":      8,
	"UPDATE_TOK":     9,
	"NEWS":           10,
	"LANDING_PHONE":  11,
	"VOIP":           12,
	"VIDEO":          13,
	"CONSULAR_PHONE": 14,
}

func (x Meta_Type) Enum() *Meta_Type {
	p := new(Meta_Type)
	*p = x
	return p
}
func (x Meta_Type) String() string {
	return proto.EnumName(Meta_Type_name, int32(x))
}
func (x *Meta_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Meta_Type_value, data, "Meta_Type")
	if err != nil {
		return err
	}
	*x = Meta_Type(value)
	return nil
}
func (Meta_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// bag id = 0
type Msg struct {
	To   *string   `protobuf:"bytes,1,req,name=to" json:"to,omitempty"`
	From *int64    `protobuf:"varint,2,opt,name=from" json:"from,omitempty"`
	Tp   *Msg_Type `protobuf:"varint,3,opt,name=tp,enum=Msg_Type,def=0" json:"tp,omitempty"`
	Desc *string   `protobuf:"bytes,4,opt,name=desc,def=" json:"desc,omitempty"`
	// d has different meanings for different type:
	// TXT & SYS: text content
	// IMG: image binary
	// VOICE: voice binary
	// CARD: Card binary
	D                []byte  `protobuf:"bytes,5,req,name=d" json:"d,omitempty"`
	Ct               *int64  `protobuf:"varint,6,opt,name=ct" json:"ct,omitempty"`
	Meta             []*Meta `protobuf:"bytes,7,rep,name=meta" json:"meta,omitempty"`
	Id               *string `protobuf:"bytes,8,opt,name=id,def=" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Msg) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

// the user card
type Card struct {
	Uid              *int64  `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Icon             *string `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Card) Reset()                    { *m = Card{} }
func (m *Card) String() string            { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()               {}
func (*Card) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type File struct {
	Name *string    `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Size *int64     `protobuf:"varint,2,req,name=size" json:"size,omitempty"`
	Dl   *string    `protobuf:"bytes,3,req,name=dl" json:"dl,omitempty"`
	Tp   *File_Type `protobuf:"varint,4,opt,name=tp,enum=File_Type,def=0" json:"tp,omitempty"`
	D    []byte     `protobuf:"bytes,5,opt,name=d" json:"d,omitempty"`
	// 原始文件，默认值true
	// 除图片以外的文件，值都是true
	// 当原始图片size大于一个数量级，此值是true, 否则是false
	Original         *bool  `protobuf:"varint,6,opt,name=original,def=1" json:"original,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_File_Tp File_Type = File_UNKNOWN
const Default_File_Original bool = true

func (m *File) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *File) GetDl() string {
	if m != nil && m.Dl != nil {
		return *m.Dl
	}
	return ""
}

func (m *File) GetTp() File_Type {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return Default_File_Tp
}

func (m *File) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *File) GetOriginal() bool {
	if m != nil && m.Original != nil {
		return *m.Original
	}
	return Default_File_Original
}

// bag id = 1
type Cmd struct {
	Tp               *Cmd_Type `protobuf:"varint,1,opt,name=tp,enum=Cmd_Type,def=0" json:"tp,omitempty"`
	Ct               *int64    `protobuf:"varint,2,req,name=ct" json:"ct,omitempty"`
	Txt              *string   `protobuf:"bytes,3,opt,name=txt,def=" json:"txt,omitempty"`
	Meta             []*Meta   `protobuf:"bytes,4,rep,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Cmd) Reset()                    { *m = Cmd{} }
func (m *Cmd) String() string            { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()               {}
func (*Cmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_Cmd_Tp Cmd_Type = Cmd_UNKNOWN

func (m *Cmd) GetTp() Cmd_Type {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return Default_Cmd_Tp
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
	Tp *Meta_Type `protobuf:"varint,1,opt,name=tp,enum=Meta_Type,def=0" json:"tp,omitempty"`
	// for JOIN_TOK/LEAVE_TOK, it's comma separated uids
	// for event change it's event id
	// for ADD_FRIEND/DEL_FRIEND/USER_UPDATE, it's comma separated uids
	// for USER_UPDATE, it's uid
	// for SETUP_TOK, UPDATE_TOK, it's nothing
	Txt              *string `protobuf:"bytes,2,req,name=txt" json:"txt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Meta) Reset()                    { *m = Meta{} }
func (m *Meta) String() string            { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()               {}
func (*Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_Meta_Tp Meta_Type = Meta_UNKNOWN

func (m *Meta) GetTp() Meta_Type {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return Default_Meta_Tp
}

func (m *Meta) GetTxt() string {
	if m != nil && m.Txt != nil {
		return *m.Txt
	}
	return ""
}

type Ack struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Ack) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "Msg")
	proto.RegisterType((*Card)(nil), "Card")
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*Cmd)(nil), "Cmd")
	proto.RegisterType((*Meta)(nil), "Meta")
	proto.RegisterType((*Ack)(nil), "Ack")
	proto.RegisterEnum("Msg_Type", Msg_Type_name, Msg_Type_value)
	proto.RegisterEnum("File_Type", File_Type_name, File_Type_value)
	proto.RegisterEnum("Cmd_Type", Cmd_Type_name, Cmd_Type_value)
	proto.RegisterEnum("Meta_Type", Meta_Type_name, Meta_Type_value)
}

var fileDescriptor0 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0x9b, 0x40,
	0x10, 0xc7, 0xc3, 0x87, 0x6d, 0x18, 0x7f, 0x74, 0x43, 0xa5, 0x96, 0x4b, 0x9b, 0x88, 0x53, 0x4e,
	0x54, 0xca, 0x31, 0x37, 0x02, 0xeb, 0x64, 0x1b, 0x7b, 0x41, 0x06, 0xf2, 0xd1, 0x8b, 0xeb, 0x18,
	0x9a, 0xa2, 0xc4, 0x21, 0x25, 0x58, 0x4a, 0xfb, 0x2a, 0xbd, 0xf4, 0x45, 0xfa, 0x40, 0x55, 0x5f,
	0xa2, 0xb3, 0x0b, 0x8e, 0xa2, 0x28, 0xbd, 0x31, 0xb3, 0xf3, 0xf9, 0xfb, 0x0f, 0xd0, 0xbf, 0x2e,
	0xb3, 0x7c, 0xe9, 0xde, 0x55, 0x65, 0x5d, 0x3a, 0x7f, 0x15, 0xd0, 0xa6, 0xf7, 0x57, 0x16, 0x80,
	0x5a, 0x97, 0xb6, 0xb2, 0xab, 0xee, 0x99, 0xd6, 0x00, 0xf4, 0x2f, 0x55, 0xb9, 0xb2, 0xd5, 0x5d,
	0x65, 0x4f, 0xb3, 0xde, 0xe3, 0xcb, 0x9d, 0xad, 0xe1, 0xf7, 0x68, 0xdf, 0x74, 0x31, 0xd6, 0x4d,
	0xbe, 0xdf, 0xe5, 0x07, 0xbd, 0x94, 0x9f, 0xf0, 0xf0, 0x8c, 0x5b, 0x23, 0xd0, 0xb3, 0xfc, 0x7e,
	0x69, 0xeb, 0x18, 0x61, 0x1e, 0x6c, 0x59, 0x26, 0x28, 0x99, 0xdd, 0xc1, 0x42, 0x03, 0x51, 0x74,
	0x59, 0xdb, 0x5d, 0x59, 0xe6, 0x35, 0xe8, 0xab, 0xbc, 0x5e, 0xd8, 0xbd, 0x5d, 0x6d, 0xaf, 0xbf,
	0xdf, 0x71, 0xa7, 0x68, 0x60, 0x27, 0xb5, 0xc8, 0x6c, 0xa3, 0xc9, 0x74, 0x3e, 0x83, 0x2e, 0x4a,
	0x5b, 0x7d, 0xd8, 0x14, 0x27, 0x5b, 0x56, 0x0f, 0xb4, 0xe4, 0x3c, 0x21, 0x8a, 0xf8, 0x60, 0xd3,
	0x23, 0xa2, 0x62, 0x83, 0xce, 0x69, 0xc8, 0x7c, 0x4a, 0x34, 0xe1, 0x8b, 0x2f, 0x62, 0xa2, 0x5b,
	0x06, 0xe8, 0xbe, 0x37, 0x0b, 0x48, 0x47, 0xb8, 0x8e, 0xd8, 0x98, 0x74, 0x85, 0x8b, 0xd3, 0xb3,
	0x98, 0xf4, 0xc4, 0xd7, 0x98, 0x4d, 0x28, 0x31, 0x9c, 0x0f, 0x18, 0xb6, 0xa8, 0x32, 0xec, 0xa0,
	0xad, 0xb1, 0xb1, 0x58, 0x57, 0x13, 0xeb, 0xde, 0x2e, 0x56, 0x39, 0xae, 0xdb, 0x2e, 0x5f, 0x2c,
	0xcb, 0x5b, 0xb9, 0xb0, 0xe9, 0xfc, 0x56, 0x30, 0xb7, 0xb8, 0xc9, 0x1f, 0x83, 0x1e, 0x09, 0xdd,
	0x17, 0x3f, 0x9a, 0x14, 0x4d, 0xac, 0x99, 0xdd, 0x60, 0x82, 0x78, 0xd9, 0x91, 0xb4, 0x74, 0x49,
	0x0b, 0x5c, 0x91, 0xfa, 0x0c, 0x57, 0x8b, 0x47, 0x41, 0x3c, 0x6f, 0xc0, 0x28, 0xab, 0xe2, 0xaa,
	0xb8, 0x5d, 0xdc, 0x48, 0x48, 0xc6, 0x81, 0x5e, 0x57, 0xeb, 0xdc, 0x61, 0xff, 0xe1, 0x10, 0x05,
	0xe3, 0x86, 0x43, 0x10, 0xfa, 0xc8, 0x01, 0x3f, 0xce, 0x27, 0x71, 0x43, 0x21, 0x8a, 0x12, 0xa4,
	0xd0, 0x22, 0xea, 0x6c, 0xa0, 0x75, 0x9d, 0x9f, 0x28, 0xaf, 0xbf, 0xca, 0x5a, 0x11, 0x95, 0x56,
	0x44, 0xf4, 0x3c, 0x9b, 0xaa, 0x51, 0xaa, 0x59, 0x67, 0x08, 0x5a, 0xfd, 0x50, 0x37, 0x00, 0x50,
	0xcf, 0x8d, 0x70, 0xfa, 0x13, 0xe1, 0x9c, 0xe0, 0xa5, 0x11, 0x09, 0x0c, 0x52, 0xee, 0xa5, 0xc9,
	0x71, 0x38, 0x63, 0x9f, 0x68, 0x80, 0xb3, 0x22, 0xf9, 0x88, 0xf1, 0x8d, 0x68, 0x2c, 0x66, 0x09,
	0x8e, 0x8b, 0xce, 0xf8, 0x82, 0xfb, 0x44, 0x77, 0x7e, 0xa9, 0xa0, 0xcb, 0x3b, 0xd8, 0x79, 0x32,
	0x1e, 0xc8, 0x0e, 0xcf, 0xe6, 0xeb, 0x37, 0x33, 0x49, 0x89, 0x9c, 0x3f, 0xca, 0x4b, 0xdd, 0x07,
	0x60, 0x7c, 0x0c, 0x19, 0x9f, 0x27, 0xe1, 0x09, 0x76, 0x1e, 0x82, 0x39, 0xa1, 0xde, 0x29, 0x95,
	0xa6, 0x6a, 0xbd, 0x82, 0x7e, 0xc0, 0xe2, 0x29, 0x8b, 0x63, 0xe9, 0xd0, 0xc4, 0xac, 0xf4, 0x94,
	0xf2, 0x64, 0xee, 0x1f, 0x7b, 0xfc, 0x88, 0x22, 0xbc, 0x11, 0x80, 0x17, 0x04, 0xf3, 0xf1, 0x8c,
	0x51, 0x2e, 0x0e, 0x09, 0xed, 0x80, 0x4e, 0x36, 0x76, 0x57, 0x94, 0x48, 0x63, 0x3a, 0x9b, 0xa7,
	0x51, 0xe0, 0x25, 0x14, 0xcf, 0x0a, 0x5b, 0xc4, 0x34, 0x49, 0x23, 0x59, 0xd1, 0x10, 0xf1, 0xcd,
	0x93, 0xb4, 0xcd, 0xc7, 0xfb, 0x03, 0x6b, 0x1b, 0x86, 0x13, 0x8f, 0x07, 0x08, 0x62, 0x1e, 0x1d,
	0x87, 0x9c, 0x92, 0xbe, 0x78, 0xc4, 0x1b, 0x8e, 0xc8, 0xa0, 0x01, 0x13, 0xd0, 0x90, 0x0c, 0x2d,
	0x0b, 0x46, 0x7e, 0xc8, 0xe3, 0x74, 0xe2, 0xcd, 0xda, 0xc0, 0x91, 0xb3, 0x0d, 0x9a, 0xb7, 0xbc,
	0x16, 0xfa, 0xb4, 0xf7, 0x6a, 0x1e, 0xbe, 0x83, 0xb7, 0xcb, 0x72, 0xe5, 0x5e, 0x15, 0xf5, 0xd7,
	0xf5, 0xa5, 0xfb, 0x6d, 0x9d, 0x3f, 0xe4, 0x95, 0x2b, 0xff, 0xe9, 0x43, 0x35, 0xba, 0xfc, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x80, 0x1b, 0x31, 0xdf, 0x03, 0x00, 0x00,
}
