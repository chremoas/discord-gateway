// Automatically generated by MockGen. DO NOT EDIT!
// Source: gateway.pb.go

package discord_gateway

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	context "golang.org/x/net/context"
)

// Mock of DiscordGatewayClient interface
type MockDiscordGatewayClient struct {
	ctrl     *gomock.Controller
	recorder *_MockDiscordGatewayClientRecorder
}

// Recorder for MockDiscordGatewayClient (not exported)
type _MockDiscordGatewayClientRecorder struct {
	mock *MockDiscordGatewayClient
}

func NewMockDiscordGatewayClient(ctrl *gomock.Controller) *MockDiscordGatewayClient {
	mock := &MockDiscordGatewayClient{ctrl: ctrl}
	mock.recorder = &_MockDiscordGatewayClientRecorder{mock}
	return mock
}

func (_m *MockDiscordGatewayClient) EXPECT() *_MockDiscordGatewayClientRecorder {
	return _m.recorder
}

func (_m *MockDiscordGatewayClient) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...client.CallOption) (*UpdateMemberResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "UpdateMember", _s...)
	ret0, _ := ret[0].(*UpdateMemberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordGatewayClientRecorder) UpdateMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMember", _s...)
}

func (_m *MockDiscordGatewayClient) GetAllMembers(ctx context.Context, in *GetAllMembersRequest, opts ...client.CallOption) (*GetMembersResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetAllMembers", _s...)
	ret0, _ := ret[0].(*GetMembersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordGatewayClientRecorder) GetAllMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllMembers", _s...)
}

func (_m *MockDiscordGatewayClient) GetAllRoles(ctx context.Context, in *GuildObjectRequest, opts ...client.CallOption) (*GetRoleResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetAllRoles", _s...)
	ret0, _ := ret[0].(*GetRoleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordGatewayClientRecorder) GetAllRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllRoles", _s...)
}

func (_m *MockDiscordGatewayClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...client.CallOption) (*CreateRolesResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateRole", _s...)
	ret0, _ := ret[0].(*CreateRolesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordGatewayClientRecorder) CreateRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRole", _s...)
}

func (_m *MockDiscordGatewayClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...client.CallOption) (*DeleteRoleResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeleteRole", _s...)
	ret0, _ := ret[0].(*DeleteRoleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordGatewayClientRecorder) DeleteRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRole", _s...)
}

func (_m *MockDiscordGatewayClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetUser", _s...)
	ret0, _ := ret[0].(*GetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordGatewayClientRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUser", _s...)
}

// Mock of DiscordGatewayHandler interface
type MockDiscordGatewayHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockDiscordGatewayHandlerRecorder
}

// Recorder for MockDiscordGatewayHandler (not exported)
type _MockDiscordGatewayHandlerRecorder struct {
	mock *MockDiscordGatewayHandler
}

func NewMockDiscordGatewayHandler(ctrl *gomock.Controller) *MockDiscordGatewayHandler {
	mock := &MockDiscordGatewayHandler{ctrl: ctrl}
	mock.recorder = &_MockDiscordGatewayHandlerRecorder{mock}
	return mock
}

func (_m *MockDiscordGatewayHandler) EXPECT() *_MockDiscordGatewayHandlerRecorder {
	return _m.recorder
}

func (_m *MockDiscordGatewayHandler) UpdateMember(_param0 context.Context, _param1 *UpdateMemberRequest, _param2 *UpdateMemberResponse) error {
	ret := _m.ctrl.Call(_m, "UpdateMember", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordGatewayHandlerRecorder) UpdateMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMember", arg0, arg1, arg2)
}

func (_m *MockDiscordGatewayHandler) GetAllMembers(_param0 context.Context, _param1 *GetAllMembersRequest, _param2 *GetMembersResponse) error {
	ret := _m.ctrl.Call(_m, "GetAllMembers", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordGatewayHandlerRecorder) GetAllMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllMembers", arg0, arg1, arg2)
}

func (_m *MockDiscordGatewayHandler) GetAllRoles(_param0 context.Context, _param1 *GuildObjectRequest, _param2 *GetRoleResponse) error {
	ret := _m.ctrl.Call(_m, "GetAllRoles", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordGatewayHandlerRecorder) GetAllRoles(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllRoles", arg0, arg1, arg2)
}

func (_m *MockDiscordGatewayHandler) CreateRole(_param0 context.Context, _param1 *CreateRoleRequest, _param2 *CreateRolesResponse) error {
	ret := _m.ctrl.Call(_m, "CreateRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordGatewayHandlerRecorder) CreateRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRole", arg0, arg1, arg2)
}

func (_m *MockDiscordGatewayHandler) DeleteRole(_param0 context.Context, _param1 *DeleteRoleRequest, _param2 *DeleteRoleResponse) error {
	ret := _m.ctrl.Call(_m, "DeleteRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordGatewayHandlerRecorder) DeleteRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRole", arg0, arg1, arg2)
}

func (_m *MockDiscordGatewayHandler) GetUser(_param0 context.Context, _param1 *GetUserRequest, _param2 *GetUserResponse) error {
	ret := _m.ctrl.Call(_m, "GetUser", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordGatewayHandlerRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUser", arg0, arg1, arg2)
}
