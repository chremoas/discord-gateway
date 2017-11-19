// Automatically generated by MockGen. DO NOT EDIT!
// Source: discord.go

package discord

import (
	discordgo "github.com/bwmarrin/discordgo"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DiscordClient interface
type MockDiscordClient struct {
	ctrl     *gomock.Controller
	recorder *_MockDiscordClientRecorder
}

// Recorder for MockDiscordClient (not exported)
type _MockDiscordClientRecorder struct {
	mock *MockDiscordClient
}

func NewMockDiscordClient(ctrl *gomock.Controller) *MockDiscordClient {
	mock := &MockDiscordClient{ctrl: ctrl}
	mock.recorder = &_MockDiscordClientRecorder{mock}
	return mock
}

func (_m *MockDiscordClient) EXPECT() *_MockDiscordClientRecorder {
	return _m.recorder
}

func (_m *MockDiscordClient) UpdateMember(guildID string, userID string, roles []string) error {
	ret := _m.ctrl.Call(_m, "UpdateMember", guildID, userID, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordClientRecorder) UpdateMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMember", arg0, arg1, arg2)
}

func (_m *MockDiscordClient) RemoveMemberRole(guildID string, userID string, role string) error {
	ret := _m.ctrl.Call(_m, "RemoveMemberRole", guildID, userID, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDiscordClientRecorder) RemoveMemberRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveMemberRole", arg0, arg1, arg2)
}

func (_m *MockDiscordClient) GetAllMembers(guildID string, after string, limit int) ([]*discordgo.Member, error) {
	ret := _m.ctrl.Call(_m, "GetAllMembers", guildID, after, limit)
	ret0, _ := ret[0].([]*discordgo.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordClientRecorder) GetAllMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllMembers", arg0, arg1, arg2)
}

func (_m *MockDiscordClient) GetAllRoles(guildID string) ([]*discordgo.Role, error) {
	ret := _m.ctrl.Call(_m, "GetAllRoles", guildID)
	ret0, _ := ret[0].([]*discordgo.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordClientRecorder) GetAllRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllRoles", arg0)
}

func (_m *MockDiscordClient) GetUser(userID string) (*discordgo.User, error) {
	ret := _m.ctrl.Call(_m, "GetUser", userID)
	ret0, _ := ret[0].(*discordgo.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDiscordClientRecorder) GetUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUser", arg0)
}
