package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type mockError struct {
	message string
}

func (me *mockError) Error() string {
	return me.message
}

func TestUpdateRoles(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return([]*discordgo.Role{
			{
				Name: "TEST ROLE 1",
				ID:   "0123456789",
			},
			{
				Name: "TEST ROLE 2",
				ID:   "0234567890",
			},
			{
				Name: "TEST ROLE 3",
				ID:   "0345678901",
			},
		}, nil),
	)

	err := roleMap.UpdateRoles()

	if err != nil {
		t.Fatalf("Received an error when none was expected: (%s)", err)
	}

	if len(roleMap.rolesByName) == 0 {
		t.Fatal("Expected more than zero roles")
	}

	if roleMap.rolesByName["TEST ROLE 1"] == nil {
		t.Fatal("Role 1 was not properly put into the map")
	}

	if roleMap.rolesByName["TEST ROLE 2"] == nil {
		t.Fatal("Role 2 was not properly put into the map")
	}

	if roleMap.rolesByName["TEST ROLE 3"] == nil {
		t.Fatal("Role 3 was not properly put into the map")
	}

	if roleMap.rolesByName["TEST ROLE 1"].ID != "0123456789" {
		t.Fatalf("Expected id for role 1: (0123456789) but received: (%s)", roleMap.rolesByName["TEST ROLE 1"].ID)
	}

	if roleMap.rolesByName["TEST ROLE 2"].ID != "0234567890" {
		t.Fatalf("Expected id for role 2: (0234567890) but received: (%s)", roleMap.rolesByName["TEST ROLE 1"].ID)
	}

	if roleMap.rolesByName["TEST ROLE 3"].ID != "0345678901" {
		t.Fatalf("Expected id for role 3: (0345678901) but received: (%s)", roleMap.rolesByName["TEST ROLE 1"].ID)
	}
}

func TestUpdateRolesWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return(nil, &mockError{"OUCH!"}),
	)

	err := roleMap.UpdateRoles()

	if err == nil || err.Error() != "OUCH!" {
		t.Fatalf("Received nil or the wrong string, expected (OUCH!) but received: (%s)", err.Error())
	}
}

func TestGetRoles(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return([]*discordgo.Role{
			{
				Name: "TEST ROLE 1",
				ID:   "0123456789",
			},
			{
				Name: "TEST ROLE 2",
				ID:   "0234567890",
			},
			{
				Name: "TEST ROLE 3",
				ID:   "0345678901",
			},
		}, nil),
	)

	err := roleMap.UpdateRoles()

	if err != nil {
		t.Fatalf("Received an error when none was expected: (%s)", err)
	}

	roles := roleMap.GetRoles()

	if len(roles) == 0 {
		t.Fatal("Expected more than zero roles")
	}

	if roles["TEST ROLE 1"] == nil {
		t.Fatal("Role 1 was not properly put into the map")
	}

	if roles["TEST ROLE 2"] == nil {
		t.Fatal("Role 2 was not properly put into the map")
	}

	if roles["TEST ROLE 3"] == nil {
		t.Fatal("Role 3 was not properly put into the map")
	}

	if roles["TEST ROLE 1"].ID != "0123456789" {
		t.Fatalf("Expected id for role 1: (0123456789) but received: (%s)", roles["TEST ROLE 1"].ID)
	}

	if roles["TEST ROLE 2"].ID != "0234567890" {
		t.Fatalf("Expected id for role 2: (0234567890) but received: (%s)", roles["TEST ROLE 1"].ID)
	}

	if roles["TEST ROLE 3"].ID != "0345678901" {
		t.Fatalf("Expected id for role 3: (0345678901) but received: (%s)", roles["TEST ROLE 1"].ID)
	}
}

func TestGetRoleId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return([]*discordgo.Role{
			{
				Name: "TEST ROLE 1",
				ID:   "0123456789",
			},
			{
				Name: "TEST ROLE 2",
				ID:   "0234567890",
			},
			{
				Name: "TEST ROLE 3",
				ID:   "0345678901",
			},
		}, nil),
	)

	err := roleMap.UpdateRoles()

	if err != nil {
		t.Fatalf("Received an error when none was expected: (%s)", err)
	}

	roleId := roleMap.GetRoleId("TEST ROLE 1")

	if len(roleId) == 0 {
		t.Fatal("Expected something as the role id but got 0 length string")
	}

	if roleId != "0123456789" {
		t.Fatalf("Expected role id: (%s) but received: (%s)", "0123456789", roleId)
	}
}

func TestGetRoleName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return([]*discordgo.Role{
			{
				Name: "TEST ROLE 1",
				ID:   "0123456789",
			},
			{
				Name: "TEST ROLE 2",
				ID:   "0234567890",
			},
			{
				Name: "TEST ROLE 3",
				ID:   "0345678901",
			},
		}, nil),
	)

	err := roleMap.UpdateRoles()

	if err != nil {
		t.Fatalf("Received an error when none was expected: (%s)", err)
	}

	roleName := roleMap.GetRoleName("0123456789")

	if roleName != "TEST ROLE 1" {
		t.Fatalf("Expected: (TEST ROLE 1) but recieved: (%s)", roleName)
	}
}

func TestGetRoleNameWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return([]*discordgo.Role{
			{
				Name: "TEST ROLE 1",
				ID:   "0123456789",
			},
			{
				Name: "TEST ROLE 2",
				ID:   "0234567890",
			},
			{
				Name: "TEST ROLE 3",
				ID:   "0345678901",
			},
		}, nil),
	)

	err := roleMap.UpdateRoles()

	if err != nil {
		t.Fatalf("Received an error when none was expected: (%s)", err)
	}

	roleName := roleMap.GetRoleName("01234567890")

	if roleName != "" {
		t.Fatalf("Expected: (\"\") but recieved: (%s)", roleName)
	}
}

func TestGetRoleIdForNonRole(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles("1234567890").Return([]*discordgo.Role{
			{
				Name: "TEST ROLE 1",
				ID:   "0123456789",
			},
			{
				Name: "TEST ROLE 2",
				ID:   "0234567890",
			},
			{
				Name: "TEST ROLE 3",
				ID:   "0345678901",
			},
		}, nil),
	)

	err := roleMap.UpdateRoles()

	if err != nil {
		t.Fatalf("Received an error when none was expected: (%s)", err)
	}

	roleId := roleMap.GetRoleId("DERP")

	if len(roleId) != 0 {
		t.Fatal("Received something when nothing was expected")
	}
}

func TestRoleMapImpl_GetRoleByName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	mockClient.EXPECT().GetAllRoles("1234567890").Times(1).Return([]*discordgo.Role{
		{
			ID:          "0123456789",
			Name:        "TEST ROLE 1",
			Color:       1,
			Hoist:       false,
			Managed:     false,
			Mentionable: false,
			Permissions: 1,
			Position:    1,
		},
		{
			ID:          "0234567890",
			Name:        "TEST ROLE 2",
			Color:       2,
			Hoist:       true,
			Managed:     true,
			Mentionable: true,
			Permissions: 2,
			Position:    2,
		},
		{
			ID:          "0345678901",
			Name:        "TEST ROLE 3",
			Color:       3,
			Hoist:       false,
			Managed:     true,
			Mentionable: false,
			Permissions: 3,
			Position:    3,
		},
	}, nil)

	err := roleMap.UpdateRoles()

	Convey("Given a populated RoleMap", t, func() {
		Convey("There is no error", func() {
			So(err, ShouldBeNil)
		})

		Convey("Searching for a role by name finds the right role", func() {
			role := roleMap.GetRoleByName("TEST ROLE 1")

			So(role, ShouldResemble, &discordgo.Role{
				ID:          "0123456789",
				Name:        "TEST ROLE 1",
				Color:       1,
				Hoist:       false,
				Managed:     false,
				Mentionable: false,
				Permissions: 1,
				Position:    1,
			})
		})
	})
}

func TestRoleMapImpl_GetRoleByName_NoRoleFound(t *testing.T) {

}

func TestRoleMapImpl_GetRoleById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := &roleMapImpl{guildID: "1234567890", client: mockClient}

	mockClient.EXPECT().GetAllRoles("1234567890").Times(1).Return([]*discordgo.Role{
		{
			ID:          "0123456789",
			Name:        "TEST ROLE 1",
			Color:       1,
			Hoist:       false,
			Managed:     false,
			Mentionable: false,
			Permissions: 1,
			Position:    1,
		},
		{
			ID:          "0234567890",
			Name:        "TEST ROLE 2",
			Color:       2,
			Hoist:       true,
			Managed:     true,
			Mentionable: true,
			Permissions: 2,
			Position:    2,
		},
		{
			ID:          "0345678901",
			Name:        "TEST ROLE 3",
			Color:       3,
			Hoist:       false,
			Managed:     true,
			Mentionable: false,
			Permissions: 3,
			Position:    3,
		},
	}, nil)

	err := roleMap.UpdateRoles()

	Convey("Given a populated RoleMap", t, func() {
		Convey("There is no error", func() {
			So(err, ShouldBeNil)
		})

		Convey("Searching for a role by name finds the right role", func() {
			role := roleMap.GetRoleById("0123456789")

			So(role, ShouldResemble, &discordgo.Role{
				ID:          "0123456789",
				Name:        "TEST ROLE 1",
				Color:       1,
				Hoist:       false,
				Managed:     false,
				Mentionable: false,
				Permissions: 1,
				Position:    1,
			})
		})
	})
}

func TestNewRoleMap(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := NewMockDiscordClient(mockCtrl)
	defer mockCtrl.Finish()

	roleMap := NewRoleMap("1234567890", mockClient).(*roleMapImpl)

	if roleMap.client != mockClient {
		t.Fatalf("Expected clients to match but they don't, original: (%+v) and result: (%+v)", mockClient, roleMap.client)
	}

	if roleMap.guildID != "1234567890" {
		t.Fatalf("Expected guild id's to match but they don't, origin: (1234567890) and result: (%s)", roleMap.guildID)
	}
}
