// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

/*
Package discord_gateway is a generated protocol buffer package.

It is generated from these files:
	gateway.proto

It has these top-level messages:
	UpdateMemberRequest
	UpdateMemberResponse
	GetAllMembersRequest
	GuildObjectRequest
	GetMembersResponse
	GetRoleResponse
	CreateRoleRequest
	CreateRolesResponse
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

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

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

type UpdateMemberRequest struct {
	Operation MemberUpdateOperation `protobuf:"varint,2,opt,name=Operation,enum=discord.gateway.MemberUpdateOperation" json:"Operation,omitempty"`
	UserId    string                `protobuf:"bytes,3,opt,name=UserId" json:"UserId,omitempty"`
	RoleIds   []string              `protobuf:"bytes,4,rep,name=RoleIds" json:"RoleIds,omitempty"`
}

func (m *UpdateMemberRequest) Reset()                    { *m = UpdateMemberRequest{} }
func (m *UpdateMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateMemberRequest) ProtoMessage()               {}
func (*UpdateMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (*UpdateMemberResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
func (*GetAllMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*GuildObjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetMembersResponse struct {
	Members []*Member `protobuf:"bytes,1,rep,name=Members" json:"Members,omitempty"`
}

func (m *GetMembersResponse) Reset()                    { *m = GetMembersResponse{} }
func (m *GetMembersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMembersResponse) ProtoMessage()               {}
func (*GetMembersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*GetRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*CreateRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*CreateRolesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateRolesResponse) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type GetUserRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*GetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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
	proto.RegisterType((*UpdateMemberRequest)(nil), "discord.gateway.UpdateMemberRequest")
	proto.RegisterType((*UpdateMemberResponse)(nil), "discord.gateway.UpdateMemberResponse")
	proto.RegisterType((*GetAllMembersRequest)(nil), "discord.gateway.GetAllMembersRequest")
	proto.RegisterType((*GuildObjectRequest)(nil), "discord.gateway.GuildObjectRequest")
	proto.RegisterType((*GetMembersResponse)(nil), "discord.gateway.GetMembersResponse")
	proto.RegisterType((*GetRoleResponse)(nil), "discord.gateway.GetRoleResponse")
	proto.RegisterType((*CreateRoleRequest)(nil), "discord.gateway.CreateRoleRequest")
	proto.RegisterType((*CreateRolesResponse)(nil), "discord.gateway.CreateRolesResponse")
	proto.RegisterType((*GetUserRequest)(nil), "discord.gateway.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "discord.gateway.GetUserResponse")
	proto.RegisterType((*Role)(nil), "discord.gateway.Role")
	proto.RegisterType((*Member)(nil), "discord.gateway.Member")
	proto.RegisterType((*User)(nil), "discord.gateway.User")
	proto.RegisterEnum("discord.gateway.MemberUpdateOperation", MemberUpdateOperation_name, MemberUpdateOperation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DiscordGateway service

type DiscordGatewayClient interface {
	UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...client.CallOption) (*UpdateMemberResponse, error)
	GetAllMembers(ctx context.Context, in *GetAllMembersRequest, opts ...client.CallOption) (*GetMembersResponse, error)
	GetAllRoles(ctx context.Context, in *GuildObjectRequest, opts ...client.CallOption) (*GetRoleResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*CreateRolesResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
}

type discordGatewayClient struct {
	c           client.Client
	serviceName string
}

func NewDiscordGatewayClient(serviceName string, c client.Client) DiscordGatewayClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "discord.gateway"
	}
	return &discordGatewayClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *discordGatewayClient) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...client.CallOption) (*UpdateMemberResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DiscordGateway.UpdateMember", in)
	out := new(UpdateMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordGatewayClient) GetAllMembers(ctx context.Context, in *GetAllMembersRequest, opts ...client.CallOption) (*GetMembersResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DiscordGateway.GetAllMembers", in)
	out := new(GetMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordGatewayClient) GetAllRoles(ctx context.Context, in *GuildObjectRequest, opts ...client.CallOption) (*GetRoleResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DiscordGateway.GetAllRoles", in)
	out := new(GetRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordGatewayClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*CreateRolesResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DiscordGateway.CreateRole", in)
	out := new(CreateRolesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordGatewayClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DiscordGateway.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DiscordGateway service

type DiscordGatewayHandler interface {
	UpdateMember(context.Context, *UpdateMemberRequest, *UpdateMemberResponse) error
	GetAllMembers(context.Context, *GetAllMembersRequest, *GetMembersResponse) error
	GetAllRoles(context.Context, *GuildObjectRequest, *GetRoleResponse) error
	CreateRole(context.Context, *CreateRoleRequest, *CreateRolesResponse) error
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
}

func RegisterDiscordGatewayHandler(s server.Server, hdlr DiscordGatewayHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DiscordGateway{hdlr}, opts...))
}

type DiscordGateway struct {
	DiscordGatewayHandler
}

func (h *DiscordGateway) UpdateMember(ctx context.Context, in *UpdateMemberRequest, out *UpdateMemberResponse) error {
	return h.DiscordGatewayHandler.UpdateMember(ctx, in, out)
}

func (h *DiscordGateway) GetAllMembers(ctx context.Context, in *GetAllMembersRequest, out *GetMembersResponse) error {
	return h.DiscordGatewayHandler.GetAllMembers(ctx, in, out)
}

func (h *DiscordGateway) GetAllRoles(ctx context.Context, in *GuildObjectRequest, out *GetRoleResponse) error {
	return h.DiscordGatewayHandler.GetAllRoles(ctx, in, out)
}

func (h *DiscordGateway) CreateRole(ctx context.Context, in *CreateRoleRequest, out *CreateRolesResponse) error {
	return h.DiscordGatewayHandler.CreateRole(ctx, in, out)
}

func (h *DiscordGateway) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.DiscordGatewayHandler.GetUser(ctx, in, out)
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0xf5, 0xcf, 0xb1, 0x2d, 0xbb, 0x6b, 0xd9, 0x26, 0x74, 0x68, 0x85, 0xad, 0x5b, 0xa8,
	0x2d, 0x2a, 0xa0, 0xee, 0x35, 0x08, 0x20, 0x5b, 0x8a, 0x62, 0x23, 0xb2, 0x84, 0xf5, 0x0f, 0x72,
	0x31, 0x0c, 0x4a, 0x1c, 0x1b, 0x8c, 0x25, 0x52, 0xe1, 0xae, 0x12, 0xe4, 0x9c, 0x73, 0xce, 0x79,
	0xad, 0x20, 0x8f, 0x91, 0xa7, 0x08, 0x76, 0x97, 0xa4, 0x28, 0x91, 0x89, 0x7d, 0xdb, 0x6f, 0xf8,
	0xcd, 0xec, 0xcc, 0x37, 0xb3, 0x43, 0xd8, 0xba, 0xb7, 0x05, 0xbe, 0xb7, 0x3f, 0xb4, 0xe7, 0x81,
	0x2f, 0x7c, 0xb2, 0xed, 0xb8, 0x7c, 0xe2, 0x07, 0x4e, 0x3b, 0x34, 0xd3, 0x4f, 0x06, 0xec, 0x5e,
	0xcd, 0x1d, 0x5b, 0xe0, 0x00, 0x67, 0x63, 0x0c, 0x18, 0xbe, 0x5d, 0x20, 0x17, 0xa4, 0x0b, 0xe6,
	0x70, 0x8e, 0x81, 0x2d, 0x5c, 0xdf, 0xb3, 0xf2, 0x4d, 0xa3, 0x55, 0x3b, 0xfa, 0xb3, 0xbd, 0xe6,
	0xdc, 0xd6, 0x2e, 0xda, 0x3d, 0x66, 0xb3, 0xa5, 0x23, 0xd9, 0x87, 0xf2, 0x15, 0xc7, 0xe0, 0xd4,
	0xb1, 0x0a, 0x4d, 0xa3, 0x65, 0xb2, 0x10, 0x11, 0x0b, 0x2a, 0xcc, 0x9f, 0xe2, 0xa9, 0xc3, 0xad,
	0x62, 0xb3, 0xd0, 0x32, 0x59, 0x04, 0xe9, 0x19, 0xd4, 0x57, 0xd3, 0xe1, 0x73, 0xdf, 0xe3, 0x28,
	0x3d, 0x2e, 0x16, 0x93, 0x09, 0x72, 0x6e, 0x19, 0x4d, 0xa3, 0x55, 0x65, 0x11, 0x94, 0x5f, 0x06,
	0xc8, 0xb9, 0x7d, 0x8f, 0x2a, 0x4f, 0x93, 0x45, 0x90, 0x32, 0xa8, 0xf7, 0x51, 0x74, 0xa6, 0x53,
	0x1d, 0x8b, 0x47, 0xb5, 0xd5, 0xa1, 0xd4, 0xb9, 0x13, 0x18, 0x84, 0x7c, 0x0d, 0xc8, 0x21, 0x6c,
	0x9d, 0x2f, 0x24, 0x6f, 0x84, 0xc1, 0x48, 0x46, 0x93, 0x29, 0x97, 0xd8, 0xaa, 0x91, 0xd6, 0x81,
	0xf4, 0x17, 0xee, 0xd4, 0x19, 0x8e, 0xdf, 0xe0, 0x44, 0x84, 0x11, 0x69, 0x1f, 0x48, 0x1f, 0x45,
	0x7c, 0x4d, 0x98, 0xf3, 0x7f, 0x32, 0x33, 0x65, 0xb2, 0x8c, 0x66, 0xa1, 0xb5, 0x71, 0x74, 0xf0,
	0x03, 0x05, 0x59, 0xc4, 0xa3, 0xcf, 0x61, 0xbb, 0x8f, 0x42, 0x8a, 0x11, 0x47, 0xf9, 0x07, 0x4a,
	0x12, 0x47, 0x31, 0xf6, 0x52, 0x31, 0x14, 0x5b, 0x73, 0xe8, 0x67, 0x03, 0x7e, 0x39, 0x09, 0xd0,
	0x16, 0xa8, 0x63, 0xe8, 0x82, 0x09, 0x14, 0xcf, 0xed, 0x59, 0xa4, 0x8f, 0x3a, 0x4b, 0x11, 0x4e,
	0xfc, 0xa9, 0x1f, 0x84, 0x65, 0x6a, 0x20, 0xad, 0x2f, 0x7d, 0x97, 0x0b, 0xab, 0xa8, 0x44, 0xd6,
	0x80, 0x34, 0x61, 0x63, 0x84, 0xc1, 0xcc, 0xe5, 0xdc, 0xf5, 0x3d, 0x6e, 0x95, 0x94, 0x47, 0xd2,
	0x24, 0x19, 0x03, 0xf4, 0x64, 0xcf, 0xed, 0xf1, 0x14, 0xad, 0xb2, 0xf2, 0x4e, 0x9a, 0xe8, 0xbf,
	0xb0, 0xbb, 0x4c, 0x6c, 0xa9, 0xd1, 0x3e, 0x94, 0x75, 0xeb, 0x55, 0x5b, 0x4d, 0x16, 0x22, 0xda,
	0x82, 0x5a, 0x1f, 0x85, 0x1c, 0x97, 0xa8, 0x88, 0xe5, 0x2c, 0x19, 0xc9, 0x59, 0xa2, 0xcf, 0x94,
	0x64, 0x9a, 0x19, 0x06, 0xfd, 0x0b, 0x8a, 0x12, 0x2b, 0x62, 0x96, 0x62, 0x8a, 0xac, 0x28, 0xf4,
	0x63, 0x1e, 0x8a, 0xf2, 0x4a, 0x52, 0x83, 0x7c, 0x1c, 0x3a, 0x7f, 0xea, 0x64, 0x6a, 0x26, 0x47,
	0xcd, 0xf6, 0xec, 0x7b, 0xd4, 0xf3, 0x5c, 0x65, 0x11, 0x5c, 0xaf, 0xbf, 0x98, 0xaa, 0x7f, 0xa9,
	0x6c, 0x29, 0xa9, 0x6c, 0xdc, 0x85, 0x72, 0xb2, 0x0b, 0x0d, 0xa8, 0x8e, 0x7c, 0xee, 0xaa, 0xb7,
	0x57, 0x51, 0x1f, 0x62, 0xbc, 0xde, 0x8b, 0x6a, 0xba, 0x17, 0x89, 0xb1, 0x33, 0x9f, 0x38, 0x76,
	0x5f, 0x0d, 0x28, 0xeb, 0xb3, 0xac, 0x51, 0x0d, 0x78, 0x2c, 0x46, 0x04, 0x65, 0x56, 0x67, 0xbe,
	0xeb, 0xa1, 0xd3, 0x11, 0xa1, 0x2a, 0x31, 0x56, 0x6a, 0xb9, 0x93, 0x87, 0xf0, 0x99, 0xab, 0xb3,
	0xb4, 0x75, 0xd1, 0xbe, 0x0b, 0xc5, 0x50, 0x67, 0x69, 0x1b, 0x2c, 0x04, 0x86, 0x22, 0xa8, 0x73,
	0xdc, 0xad, 0xf2, 0xa3, 0xdd, 0x5a, 0xbe, 0x85, 0xca, 0x13, 0xde, 0xc2, 0x37, 0x43, 0x07, 0x4e,
	0xb5, 0xb6, 0x0e, 0xa5, 0xde, 0xcc, 0x76, 0xa7, 0xd1, 0xfb, 0x57, 0x40, 0x96, 0x27, 0xd9, 0x9e,
	0x6c, 0xba, 0x2e, 0x23, 0xc6, 0x72, 0xf6, 0x3a, 0xef, 0x6c, 0x61, 0x07, 0xaa, 0x18, 0x93, 0x85,
	0x48, 0xee, 0x8c, 0xae, 0xcb, 0x27, 0x81, 0x3b, 0x73, 0x3d, 0x5b, 0xf8, 0x81, 0xaa, 0xcb, 0x64,
	0xab, 0x46, 0x79, 0xdf, 0xa5, 0xff, 0x80, 0x9e, 0xaa, 0xd0, 0x64, 0x1a, 0xc8, 0xfb, 0xae, 0x31,
	0x70, 0xef, 0x5c, 0x74, 0x54, 0x93, 0xab, 0x2c, 0xc6, 0xe4, 0x57, 0x80, 0xc1, 0x8b, 0x4e, 0x4f,
	0x4d, 0x8e, 0xa3, 0x7a, 0x5c, 0x65, 0x09, 0x0b, 0xd9, 0x81, 0xc2, 0xb1, 0x2f, 0x2c, 0x53, 0x7d,
	0x90, 0xc7, 0xbf, 0x8f, 0x61, 0x2f, 0x73, 0x1b, 0x93, 0x03, 0xd8, 0xed, 0x74, 0xbb, 0xb7, 0x43,
	0x76, 0x7b, 0x35, 0xea, 0x76, 0x2e, 0x7b, 0xb7, 0x6c, 0xf8, 0xaa, 0x77, 0xb1, 0x93, 0x23, 0x3b,
	0xb0, 0xc9, 0x7a, 0x83, 0xe1, 0x75, 0x64, 0x31, 0x8e, 0xbe, 0x14, 0xa0, 0xd6, 0xd5, 0x82, 0xf6,
	0xb5, 0x9e, 0xe4, 0x06, 0x36, 0x93, 0xeb, 0x98, 0x1c, 0xa6, 0xbb, 0x93, 0xfe, 0x79, 0x34, 0xfe,
	0x78, 0x84, 0xa5, 0x9f, 0x29, 0xcd, 0x91, 0x1b, 0xd8, 0x5a, 0xd9, 0xd0, 0x24, 0xed, 0x99, 0xb5,
	0xc1, 0x1b, 0xbf, 0x67, 0xd1, 0xd6, 0xd6, 0x2f, 0xcd, 0x91, 0x6b, 0xd8, 0xd0, 0xee, 0x6a, 0x20,
	0x48, 0x86, 0x57, 0x6a, 0x95, 0x37, 0x9a, 0x59, 0xa1, 0x93, 0x0b, 0x99, 0xe6, 0xc8, 0x6b, 0x80,
	0xe5, 0x2e, 0x23, 0x34, 0xe5, 0x91, 0xda, 0xc0, 0x8d, 0xc3, 0x9f, 0x70, 0x92, 0x19, 0x9f, 0x43,
	0x25, 0x5c, 0x66, 0xe4, 0xb7, 0xac, 0x44, 0x12, 0x0b, 0x31, 0x3b, 0xd3, 0xe4, 0x1e, 0xa4, 0xb9,
	0x71, 0x59, 0xfd, 0xf6, 0xff, 0xff, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x32, 0xce, 0x60, 0xed, 0x07,
	0x08, 0x00, 0x00,
}
