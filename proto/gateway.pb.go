// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

/*
Package discord_gateway is a generated protocol buffer package.

It is generated from these files:
	gateway.proto

It has these top-level messages:
	NilMessage
	SendMessageRequest
	UpdateMemberRequest
	UpdateMemberResponse
	GetAllMembersRequest
	GuildObjectRequest
	GetMembersResponse
	GetRoleResponse
	CreateRoleRequest
	CreateRolesResponse
	DeleteRoleRequest
	DeleteRoleResponse
	GetUserRequest
	GetUserResponse
	Role
	Member
	User
*/
package discord_gateway

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MemberUpdateOperation int32

const (
	MemberUpdateOperation_ADD_OR_UPDATE_ROLES MemberUpdateOperation = 0
	MemberUpdateOperation_REMOVE_ROLES        MemberUpdateOperation = 1
)

var MemberUpdateOperation_name = map[int32]string{
	0: "ADD_OR_UPDATE_ROLES",
	1: "REMOVE_ROLES",
}
var MemberUpdateOperation_value = map[string]int32{
	"ADD_OR_UPDATE_ROLES": 0,
	"REMOVE_ROLES":        1,
}

func (x MemberUpdateOperation) String() string {
	return proto.EnumName(MemberUpdateOperation_name, int32(x))
}
func (MemberUpdateOperation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NilMessage struct {
}

func (m *NilMessage) Reset()                    { *m = NilMessage{} }
func (m *NilMessage) String() string            { return proto.CompactTextString(m) }
func (*NilMessage) ProtoMessage()               {}
func (*NilMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SendMessageRequest struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channelId" json:"channelId,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *SendMessageRequest) Reset()                    { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()               {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SendMessageRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateMemberRequest struct {
	Operation MemberUpdateOperation `protobuf:"varint,2,opt,name=Operation,enum=discord.gateway.MemberUpdateOperation" json:"Operation,omitempty"`
	UserId    string                `protobuf:"bytes,3,opt,name=UserId" json:"UserId,omitempty"`
	RoleIds   []string              `protobuf:"bytes,4,rep,name=RoleIds" json:"RoleIds,omitempty"`
}

func (m *UpdateMemberRequest) Reset()                    { *m = UpdateMemberRequest{} }
func (m *UpdateMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateMemberRequest) ProtoMessage()               {}
func (*UpdateMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateMemberRequest) GetOperation() MemberUpdateOperation {
	if m != nil {
		return m.Operation
	}
	return MemberUpdateOperation_ADD_OR_UPDATE_ROLES
}

func (m *UpdateMemberRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateMemberRequest) GetRoleIds() []string {
	if m != nil {
		return m.RoleIds
	}
	return nil
}

type UpdateMemberResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *UpdateMemberResponse) Reset()                    { *m = UpdateMemberResponse{} }
func (m *UpdateMemberResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateMemberResponse) ProtoMessage()               {}
func (*UpdateMemberResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateMemberResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateMemberResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetAllMembersRequest struct {
	After         string `protobuf:"bytes,2,opt,name=After" json:"After,omitempty"`
	NumberPerPage int32  `protobuf:"varint,3,opt,name=NumberPerPage" json:"NumberPerPage,omitempty"`
}

func (m *GetAllMembersRequest) Reset()                    { *m = GetAllMembersRequest{} }
func (m *GetAllMembersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllMembersRequest) ProtoMessage()               {}
func (*GetAllMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetAllMembersRequest) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

func (m *GetAllMembersRequest) GetNumberPerPage() int32 {
	if m != nil {
		return m.NumberPerPage
	}
	return 0
}

type GuildObjectRequest struct {
}

func (m *GuildObjectRequest) Reset()                    { *m = GuildObjectRequest{} }
func (m *GuildObjectRequest) String() string            { return proto.CompactTextString(m) }
func (*GuildObjectRequest) ProtoMessage()               {}
func (*GuildObjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetMembersResponse struct {
	Members []*Member `protobuf:"bytes,1,rep,name=Members" json:"Members,omitempty"`
}

func (m *GetMembersResponse) Reset()                    { *m = GetMembersResponse{} }
func (m *GetMembersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMembersResponse) ProtoMessage()               {}
func (*GetMembersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetMembersResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type GetRoleResponse struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *GetRoleResponse) Reset()                    { *m = GetRoleResponse{} }
func (m *GetRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRoleResponse) ProtoMessage()               {}
func (*GetRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetRoleResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type CreateRoleRequest struct {
	Name        string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Color       int32  `protobuf:"varint,3,opt,name=Color" json:"Color,omitempty"`
	Hoist       bool   `protobuf:"varint,4,opt,name=Hoist" json:"Hoist,omitempty"`
	Permissions int32  `protobuf:"varint,5,opt,name=Permissions" json:"Permissions,omitempty"`
	Mentionable bool   `protobuf:"varint,6,opt,name=Mentionable" json:"Mentionable,omitempty"`
}

func (m *CreateRoleRequest) Reset()                    { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()               {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateRoleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRoleRequest) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *CreateRoleRequest) GetHoist() bool {
	if m != nil {
		return m.Hoist
	}
	return false
}

func (m *CreateRoleRequest) GetPermissions() int32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *CreateRoleRequest) GetMentionable() bool {
	if m != nil {
		return m.Mentionable
	}
	return false
}

type CreateRolesResponse struct {
	RoleId string `protobuf:"bytes,1,opt,name=RoleId" json:"RoleId,omitempty"`
}

func (m *CreateRolesResponse) Reset()                    { *m = CreateRolesResponse{} }
func (m *CreateRolesResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRolesResponse) ProtoMessage()               {}
func (*CreateRolesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreateRolesResponse) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type DeleteRoleRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *DeleteRoleRequest) Reset()                    { *m = DeleteRoleRequest{} }
func (m *DeleteRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoleRequest) ProtoMessage()               {}
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteRoleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteRoleResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteRoleResponse) Reset()                    { *m = DeleteRoleResponse{} }
func (m *DeleteRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoleResponse) ProtoMessage()               {}
func (*DeleteRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteRoleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetUserRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=User" json:"User,omitempty"`
}

func (m *GetUserResponse) Reset()                    { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()               {}
func (*GetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Role struct {
	Id          string    `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Managed     bool      `protobuf:"varint,3,opt,name=Managed" json:"Managed,omitempty"`
	Mentionable bool      `protobuf:"varint,4,opt,name=Mentionable" json:"Mentionable,omitempty"`
	Hoist       bool      `protobuf:"varint,5,opt,name=Hoist" json:"Hoist,omitempty"`
	Color       int32     `protobuf:"varint,6,opt,name=Color" json:"Color,omitempty"`
	Position    int32     `protobuf:"varint,7,opt,name=Position" json:"Position,omitempty"`
	Permissions int32     `protobuf:"varint,8,opt,name=Permissions" json:"Permissions,omitempty"`
	Members     []*Member `protobuf:"bytes,9,rep,name=Members" json:"Members,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetManaged() bool {
	if m != nil {
		return m.Managed
	}
	return false
}

func (m *Role) GetMentionable() bool {
	if m != nil {
		return m.Mentionable
	}
	return false
}

func (m *Role) GetHoist() bool {
	if m != nil {
		return m.Hoist
	}
	return false
}

func (m *Role) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Role) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Role) GetPermissions() int32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *Role) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type Member struct {
	GuildId  string  `protobuf:"bytes,1,opt,name=GuildId" json:"GuildId,omitempty"`
	JoinedAt string  `protobuf:"bytes,2,opt,name=JoinedAt" json:"JoinedAt,omitempty"`
	Nick     string  `protobuf:"bytes,3,opt,name=Nick" json:"Nick,omitempty"`
	Deaf     bool    `protobuf:"varint,4,opt,name=Deaf" json:"Deaf,omitempty"`
	Mute     bool    `protobuf:"varint,5,opt,name=Mute" json:"Mute,omitempty"`
	User     *User   `protobuf:"bytes,6,opt,name=User" json:"User,omitempty"`
	Roles    []*Role `protobuf:"bytes,7,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Member) GetGuildId() string {
	if m != nil {
		return m.GuildId
	}
	return ""
}

func (m *Member) GetJoinedAt() string {
	if m != nil {
		return m.JoinedAt
	}
	return ""
}

func (m *Member) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *Member) GetDeaf() bool {
	if m != nil {
		return m.Deaf
	}
	return false
}

func (m *Member) GetMute() bool {
	if m != nil {
		return m.Mute
	}
	return false
}

func (m *Member) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Member) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type User struct {
	Id            string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=Email" json:"Email,omitempty"`
	Username      string `protobuf:"bytes,3,opt,name=Username" json:"Username,omitempty"`
	Avatar        string `protobuf:"bytes,4,opt,name=Avatar" json:"Avatar,omitempty"`
	Discriminator string `protobuf:"bytes,5,opt,name=Discriminator" json:"Discriminator,omitempty"`
	Token         string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
	Verified      bool   `protobuf:"varint,7,opt,name=Verified" json:"Verified,omitempty"`
	MFAEnabled    bool   `protobuf:"varint,8,opt,name=MFAEnabled" json:"MFAEnabled,omitempty"`
	Bot           bool   `protobuf:"varint,9,opt,name=Bot" json:"Bot,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetDiscriminator() string {
	if m != nil {
		return m.Discriminator
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *User) GetMFAEnabled() bool {
	if m != nil {
		return m.MFAEnabled
	}
	return false
}

func (m *User) GetBot() bool {
	if m != nil {
		return m.Bot
	}
	return false
}

func init() {
	proto.RegisterType((*NilMessage)(nil), "discord.gateway.NilMessage")
	proto.RegisterType((*SendMessageRequest)(nil), "discord.gateway.SendMessageRequest")
	proto.RegisterType((*UpdateMemberRequest)(nil), "discord.gateway.UpdateMemberRequest")
	proto.RegisterType((*UpdateMemberResponse)(nil), "discord.gateway.UpdateMemberResponse")
	proto.RegisterType((*GetAllMembersRequest)(nil), "discord.gateway.GetAllMembersRequest")
	proto.RegisterType((*GuildObjectRequest)(nil), "discord.gateway.GuildObjectRequest")
	proto.RegisterType((*GetMembersResponse)(nil), "discord.gateway.GetMembersResponse")
	proto.RegisterType((*GetRoleResponse)(nil), "discord.gateway.GetRoleResponse")
	proto.RegisterType((*CreateRoleRequest)(nil), "discord.gateway.CreateRoleRequest")
	proto.RegisterType((*CreateRolesResponse)(nil), "discord.gateway.CreateRolesResponse")
	proto.RegisterType((*DeleteRoleRequest)(nil), "discord.gateway.DeleteRoleRequest")
	proto.RegisterType((*DeleteRoleResponse)(nil), "discord.gateway.DeleteRoleResponse")
	proto.RegisterType((*GetUserRequest)(nil), "discord.gateway.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "discord.gateway.GetUserResponse")
	proto.RegisterType((*Role)(nil), "discord.gateway.Role")
	proto.RegisterType((*Member)(nil), "discord.gateway.Member")
	proto.RegisterType((*User)(nil), "discord.gateway.User")
	proto.RegisterEnum("discord.gateway.MemberUpdateOperation", MemberUpdateOperation_name, MemberUpdateOperation_value)
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 908 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0xb6, 0x12, 0xdb, 0xb1, 0x3b, 0x3f, 0x9b, 0x9d, 0x78, 0x77, 0x55, 0x86, 0x02, 0xd7, 0x10,
	0x20, 0x40, 0xe1, 0x2a, 0xc2, 0x95, 0xa2, 0xca, 0xbb, 0x32, 0x26, 0x5b, 0xeb, 0xd8, 0x4c, 0x36,
	0x81, 0xcb, 0x56, 0x6a, 0x62, 0x75, 0xc2, 0xb0, 0xb2, 0x14, 0x34, 0x63, 0x28, 0xce, 0x9c, 0x39,
	0xf3, 0x5e, 0x3c, 0x06, 0xcf, 0xc0, 0x81, 0x9a, 0x1f, 0xfd, 0x38, 0x12, 0x24, 0xb7, 0xf9, 0x5a,
	0xdd, 0xad, 0xee, 0xef, 0xeb, 0x69, 0x09, 0x76, 0x6f, 0xb8, 0xc2, 0x5f, 0xf9, 0x6f, 0xc3, 0xdb,
	0x34, 0x51, 0x09, 0x79, 0x14, 0x0a, 0xb9, 0x48, 0xd2, 0x70, 0xe8, 0xcc, 0x74, 0x07, 0xe0, 0x54,
	0x44, 0x53, 0x94, 0x92, 0xdf, 0x20, 0x7d, 0x05, 0xe4, 0x0c, 0xe3, 0xd0, 0x41, 0x86, 0x3f, 0xaf,
	0x50, 0x2a, 0xf2, 0x2e, 0x74, 0x17, 0x3f, 0xf2, 0x38, 0xc6, 0xe8, 0x24, 0xf4, 0xbd, 0x81, 0x77,
	0xd4, 0x65, 0x85, 0x81, 0xf8, 0xb0, 0xb5, 0xb4, 0xfe, 0xfe, 0x86, 0x79, 0x96, 0x41, 0xfa, 0x87,
	0x07, 0x07, 0xe7, 0xb7, 0x21, 0x57, 0x38, 0xc5, 0xe5, 0x15, 0xa6, 0x59, 0xbe, 0x00, 0xba, 0xb3,
	0x5b, 0x4c, 0xb9, 0x12, 0x49, 0x6c, 0x62, 0xf6, 0x8e, 0x3f, 0x1a, 0xde, 0x29, 0x6c, 0x68, 0x43,
	0x6c, 0x78, 0xee, 0xcd, 0x8a, 0x40, 0xf2, 0x14, 0xda, 0xe7, 0x12, 0xd3, 0x93, 0xd0, 0xdf, 0x34,
	0xaf, 0x75, 0x48, 0xd7, 0xc3, 0x92, 0x08, 0x4f, 0x42, 0xe9, 0x37, 0x07, 0x9b, 0xba, 0x1e, 0x07,
	0xe9, 0x4b, 0xe8, 0xad, 0x97, 0x23, 0x6f, 0x93, 0x58, 0xa2, 0x8e, 0x38, 0x5b, 0x2d, 0x16, 0x28,
	0xa5, 0xe9, 0xae, 0xc3, 0x32, 0xa8, 0x9f, 0x4c, 0xd7, 0x7b, 0xcb, 0x98, 0x62, 0xd0, 0x9b, 0xa0,
	0x1a, 0x45, 0x91, 0xcd, 0x25, 0xb3, 0xde, 0x7a, 0xd0, 0x1a, 0x5d, 0x2b, 0x4c, 0x9d, 0xbf, 0x05,
	0xe4, 0x10, 0x76, 0x4f, 0x57, 0xda, 0x6f, 0x8e, 0xe9, 0x5c, 0x67, 0xd3, 0x25, 0xb7, 0xd8, 0xba,
	0x91, 0xf6, 0x80, 0x4c, 0x56, 0x22, 0x0a, 0x67, 0x57, 0x3f, 0xe1, 0x42, 0xb9, 0x8c, 0x74, 0x02,
	0x64, 0x82, 0x2a, 0x7f, 0x8d, 0xab, 0xf9, 0x0b, 0x5d, 0x99, 0x31, 0xf9, 0xde, 0x60, 0xf3, 0x68,
	0xfb, 0xf8, 0xd9, 0x7f, 0x30, 0xc8, 0x32, 0x3f, 0xfa, 0x35, 0x3c, 0x9a, 0xa0, 0xd2, 0x64, 0xe4,
	0x59, 0x3e, 0x83, 0x96, 0xc6, 0x59, 0x8e, 0x27, 0x95, 0x1c, 0xc6, 0xdb, 0xfa, 0xd0, 0x3f, 0x3d,
	0x78, 0xfc, 0x22, 0x45, 0xae, 0xd0, 0xe6, 0xb0, 0x0d, 0x13, 0x68, 0x9e, 0xf2, 0x65, 0xc6, 0x8f,
	0x39, 0x6b, 0x12, 0x5e, 0x24, 0x51, 0x92, 0xba, 0x36, 0x2d, 0xd0, 0xd6, 0x6f, 0x13, 0x21, 0x95,
	0xdf, 0x34, 0x24, 0x5b, 0x40, 0x06, 0xb0, 0x3d, 0xc7, 0x74, 0x29, 0xa4, 0x14, 0x49, 0x2c, 0xfd,
	0x96, 0x89, 0x28, 0x9b, 0xb4, 0xc7, 0x14, 0x63, 0xad, 0x39, 0xbf, 0x8a, 0xd0, 0x6f, 0x9b, 0xe8,
	0xb2, 0x89, 0x7e, 0x0e, 0x07, 0x45, 0x61, 0x05, 0x47, 0x4f, 0xa1, 0x6d, 0xa5, 0x77, 0x43, 0xeb,
	0x10, 0xfd, 0x18, 0x1e, 0x07, 0x18, 0x61, 0x7d, 0x1f, 0x5e, 0xd1, 0x07, 0x1d, 0x02, 0x29, 0x3b,
	0xde, 0x37, 0x2e, 0xf4, 0x08, 0xf6, 0x26, 0xa8, 0xf4, 0x1c, 0x66, 0x59, 0x8b, 0x21, 0xf5, 0xca,
	0x43, 0x4a, 0xbf, 0x32, 0x5a, 0x58, 0x4f, 0x97, 0xf6, 0x13, 0x68, 0x6a, 0x6c, 0x1c, 0xeb, 0xa4,
	0x30, 0xce, 0xc6, 0x85, 0xfe, 0xbe, 0x01, 0x4d, 0x5d, 0x12, 0xd9, 0x83, 0x8d, 0x3c, 0xf5, 0xc6,
	0x49, 0x58, 0x2b, 0x86, 0x9e, 0x61, 0x1e, 0xf3, 0x1b, 0xb4, 0x17, 0xa5, 0xc3, 0x32, 0x78, 0x97,
	0xd8, 0x66, 0x85, 0xd8, 0x42, 0xb2, 0x56, 0x59, 0xb2, 0x5c, 0xde, 0x76, 0x59, 0xde, 0x3e, 0x74,
	0xe6, 0x89, 0x14, 0xe6, 0x52, 0x6f, 0x99, 0x07, 0x39, 0xbe, 0x2b, 0x72, 0xa7, 0x2a, 0x72, 0x69,
	0x9e, 0xbb, 0x0f, 0x9c, 0xe7, 0xbf, 0x3c, 0x68, 0xdb, 0xb3, 0xee, 0xd1, 0xdc, 0x9c, 0x9c, 0x8c,
	0x0c, 0xea, 0xaa, 0x5e, 0x26, 0x22, 0xc6, 0x70, 0xa4, 0x1c, 0x2b, 0x39, 0x36, 0x6c, 0x89, 0xc5,
	0x5b, 0xb7, 0x3f, 0xcc, 0x59, 0xdb, 0x02, 0xe4, 0xd7, 0x8e, 0x0c, 0x73, 0xd6, 0xb6, 0xe9, 0x4a,
	0xa1, 0x23, 0xc1, 0x9c, 0x73, 0xb5, 0xda, 0xf7, 0xaa, 0x55, 0x5c, 0xb2, 0xad, 0x07, 0x5c, 0xb2,
	0xbf, 0x3d, 0x9b, 0xb8, 0x22, 0x6d, 0x0f, 0x5a, 0xe3, 0x25, 0x17, 0x51, 0xb6, 0x58, 0x0c, 0xd0,
	0xed, 0x69, 0xef, 0x58, 0x8b, 0x6e, 0xdb, 0xc8, 0xb1, 0x9e, 0xbd, 0xd1, 0x2f, 0x5c, 0xf1, 0xd4,
	0x34, 0xd3, 0x65, 0x0e, 0xe9, 0x65, 0x14, 0x08, 0xb9, 0x48, 0xc5, 0x52, 0xc4, 0x5c, 0x25, 0xa9,
	0xe9, 0xab, 0xcb, 0xd6, 0x8d, 0xfa, 0x7d, 0xaf, 0x93, 0xb7, 0x18, 0x9b, 0x0e, 0xbb, 0xcc, 0x02,
	0xfd, 0xbe, 0x0b, 0x4c, 0xc5, 0xb5, 0xc0, 0xd0, 0x88, 0xdc, 0x61, 0x39, 0x26, 0xef, 0x01, 0x4c,
	0xbf, 0x19, 0x8d, 0xcd, 0xe4, 0x84, 0x46, 0xe3, 0x0e, 0x2b, 0x59, 0xc8, 0x3e, 0x6c, 0x3e, 0x4f,
	0x94, 0xdf, 0x35, 0x0f, 0xf4, 0xf1, 0xd3, 0xe7, 0xf0, 0xa4, 0x76, 0xcd, 0x93, 0x67, 0x70, 0x30,
	0x0a, 0x82, 0xcb, 0x19, 0xbb, 0x3c, 0x9f, 0x07, 0xa3, 0xd7, 0xe3, 0x4b, 0x36, 0x7b, 0x35, 0x3e,
	0xdb, 0x6f, 0x90, 0x7d, 0xd8, 0x61, 0xe3, 0xe9, 0xec, 0x22, 0xb3, 0x78, 0xc7, 0xff, 0x34, 0x61,
	0x2f, 0xb0, 0x84, 0x4e, 0x2c, 0x9f, 0xe4, 0x0d, 0xec, 0x94, 0xf7, 0x3c, 0x39, 0xac, 0xaa, 0x53,
	0xfd, 0x2a, 0xf5, 0x3f, 0xbc, 0xc7, 0xcb, 0x5e, 0x53, 0xda, 0x20, 0x6f, 0x60, 0x77, 0x6d, 0xf5,
	0x93, 0x6a, 0x64, 0xdd, 0xa7, 0xa1, 0xff, 0x41, 0x9d, 0xdb, 0x9d, 0xbd, 0x4e, 0x1b, 0xe4, 0x02,
	0xb6, 0x6d, 0xb8, 0x19, 0x08, 0x52, 0x13, 0x55, 0xf9, 0x46, 0xf4, 0x07, 0x75, 0xa9, 0xcb, 0x4b,
	0x8b, 0x36, 0xc8, 0x0f, 0x00, 0xc5, 0x92, 0x24, 0xb4, 0x12, 0x51, 0x59, 0xed, 0xfd, 0xc3, 0xff,
	0xf1, 0x29, 0x57, 0xfc, 0x3d, 0x40, 0xb1, 0x26, 0x6b, 0x32, 0x57, 0x96, 0x6d, 0x0d, 0x15, 0xd5,
	0x3d, 0x4b, 0x1b, 0xe4, 0x14, 0xb6, 0xdc, 0x96, 0x24, 0xef, 0xd7, 0x75, 0x58, 0xda, 0xb4, 0xf5,
	0x14, 0x94, 0x17, 0x2c, 0x6d, 0x90, 0xef, 0x60, 0xbb, 0xf4, 0x7b, 0x53, 0x43, 0x6d, 0xf5, 0xe7,
	0xa7, 0xff, 0x4e, 0xc5, 0xa9, 0xf4, 0xbf, 0xd4, 0xb8, 0x6a, 0x9b, 0xff, 0xaa, 0x2f, 0xff, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x0a, 0xdf, 0x69, 0x68, 0x09, 0x00, 0x00,
}
