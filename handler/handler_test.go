package handler

import (
	"errors"
	"github.com/abaeve/discord-gateway/discord"
	proto "github.com/abaeve/discord-gateway/proto"
	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestDiscordClient_GetAllMembers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("@me").Times(1).Return(&discordgo.User{ID: "U12345678"}, nil)
	mockClient.EXPECT().GetAllMembers("G123456", "", 1000).Times(1).Return(
		[]*discordgo.Member{
			{
				GuildID:  "G123456",
				Nick:     "nick 1",
				JoinedAt: "JA1234567",
				Deaf:     false,
				Mute:     false,
				User: &discordgo.User{
					ID:            "U123456",
					Username:      "username 1",
					Email:         "user1@test.com",
					Token:         "user1token",
					Discriminator: "user1discriminator",
					Avatar:        "user1avatar",
					Verified:      false,
					MFAEnabled:    false,
					Bot:           false,
				},
				Roles: []string{
					"R123456",
					"R234567",
					"R345678",
				},
			},
			{
				GuildID:  "G123456",
				Nick:     "nick 2",
				JoinedAt: "JA2345678",
				Deaf:     true,
				Mute:     true,
				User: &discordgo.User{
					ID:            "U234567",
					Username:      "username 2",
					Email:         "user2@test.com",
					Token:         "user2token",
					Discriminator: "user2discriminator",
					Avatar:        "user2avatar",
					Verified:      true,
					MFAEnabled:    true,
					Bot:           true,
				},
				Roles: []string{
					"R456789",
				},
			},
		}, nil,
	)

	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, err := NewDiscordGatewayHandler(mockClient, mockRoleMap)
	Convey("Client construction err is nil", t, func() {
		So(err, ShouldBeNil)
	})

	response := &proto.GetMembersResponse{}
	err = client.GetAllMembers(context.Background(), &proto.GetAllMembersRequest{GuildId: "G123456", After: "", NumberPerPage: 1000}, response)

	Convey("Given a response", t, func() {
		Convey("Response should be valid with no error", func() {
			So(err, ShouldBeNil)
			So(response, ShouldNotBeNil)
		})

		Convey("Should be 2 members", func() {
			So(len(response.Members), ShouldEqual, 2)
		})

		Convey("Members have expected attributes", func() {
			So(response.Members[0].GuildId, ShouldEqual, "G123456")
			So(response.Members[1].GuildId, ShouldEqual, "G123456")

			So(response.Members[0].Nick, ShouldEqual, "nick 1")
			So(response.Members[1].Nick, ShouldEqual, "nick 2")

			So(response.Members[0].JoinedAt, ShouldEqual, "JA1234567")
			So(response.Members[1].JoinedAt, ShouldEqual, "JA2345678")

			So(response.Members[0].Deaf, ShouldBeFalse)
			So(response.Members[1].Deaf, ShouldBeTrue)

			So(response.Members[0].Mute, ShouldBeFalse)
			So(response.Members[1].Mute, ShouldBeTrue)
		})

		Convey("Members users have expected attributes", func() {
			user1 := response.Members[0].User
			user2 := response.Members[1].User

			So(user1, ShouldNotBeNil)
			So(user2, ShouldNotBeNil)

			So(user1.Id, ShouldEqual, "U123456")
			So(user2.Id, ShouldEqual, "U234567")

			So(user1.Username, ShouldEqual, "username 1")
			So(user2.Username, ShouldEqual, "username 2")

			So(user1.Email, ShouldEqual, "user1@test.com")
			So(user2.Email, ShouldEqual, "user2@test.com")

			So(user1.Discriminator, ShouldEqual, "user1discriminator")
			So(user2.Discriminator, ShouldEqual, "user2discriminator")

			So(user1.Avatar, ShouldEqual, "user1avatar")
			So(user2.Avatar, ShouldEqual, "user2avatar")

			So(user1.Verified, ShouldBeFalse)
			So(user2.Verified, ShouldBeTrue)

			So(user1.MFAEnabled, ShouldBeFalse)
			So(user2.MFAEnabled, ShouldBeTrue)

			So(user1.Bot, ShouldBeFalse)
			So(user2.Bot, ShouldBeTrue)

			So(user1.Token, ShouldEqual, "user1token")
			So(user2.Token, ShouldEqual, "user2token")
		})

		Convey("Members roles have expected attributes", func() {
			user1Roles := response.Members[0].Roles
			user2Roles := response.Members[1].Roles

			So(user1Roles, ShouldNotBeNil)
			So(user2Roles, ShouldNotBeNil)

			So(len(user1Roles), ShouldEqual, 3)
			So(len(user2Roles), ShouldEqual, 1)

			So(user1Roles[0].Id, ShouldEqual, "R123456")
			So(user1Roles[1].Id, ShouldEqual, "R234567")
			So(user1Roles[2].Id, ShouldEqual, "R345678")
			So(user2Roles[0].Id, ShouldEqual, "R456789")
		})
	})
}

//Same test as above but validating that the bot user isn't pulled in.  I don't really need a new test but it's nice
//for documentation purposes ;)
func TestDiscordClient_GetAllMembers_ExceptBotMember(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("@me").Times(1).Return(&discordgo.User{ID: "U12345678"}, nil)
	mockClient.EXPECT().GetAllMembers("G123456", "", 1000).Times(1).Return(
		[]*discordgo.Member{
			{
				GuildID:  "G123456",
				Nick:     "nick 1",
				JoinedAt: "JA1234567",
				Deaf:     false,
				Mute:     false,
				User: &discordgo.User{
					ID:            "U123456",
					Username:      "username 1",
					Email:         "user1@test.com",
					Token:         "user1token",
					Discriminator: "user1discriminator",
					Avatar:        "user1avatar",
					Verified:      false,
					MFAEnabled:    false,
					Bot:           false,
				},
				Roles: []string{
					"R123456",
					"R234567",
					"R345678",
				},
			},
			{
				GuildID:  "G123456",
				Nick:     "nick 2",
				JoinedAt: "JA2345678",
				Deaf:     true,
				Mute:     true,
				User: &discordgo.User{
					ID:            "U234567",
					Username:      "username 2",
					Email:         "user2@test.com",
					Token:         "user2token",
					Discriminator: "user2discriminator",
					Avatar:        "user2avatar",
					Verified:      true,
					MFAEnabled:    true,
					Bot:           true,
				},
				Roles: []string{
					"R456789",
				},
			},
			{
				GuildID:  "G123457",
				Nick:     "BOT",
				JoinedAt: "JA1234568",
				Deaf:     false,
				Mute:     false,
				User: &discordgo.User{
					ID:            "U12345678",
					Username:      "username 1",
					Email:         "user1@test.com",
					Token:         "user1token",
					Discriminator: "user1discriminator",
					Avatar:        "user1avatar",
					Verified:      false,
					MFAEnabled:    false,
					Bot:           false,
				},
				Roles: []string{
					"R123456",
					"R234567",
					"R345678",
				},
			},
		}, nil,
	)
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, err := NewDiscordGatewayHandler(mockClient, mockRoleMap)
	Convey("Client construction err is nil", t, func() {
		So(err, ShouldBeNil)
	})

	response := &proto.GetMembersResponse{}
	err = client.GetAllMembers(context.Background(), &proto.GetAllMembersRequest{GuildId: "G123456", After: "", NumberPerPage: 1000}, response)

	Convey("User array does not contain bot user", t, func() {
		Convey("Response should be valid with no error", func() {
			So(err, ShouldBeNil)
			So(response, ShouldNotBeNil)
		})

		Convey("Should be 2 members", func() {
			So(len(response.Members), ShouldEqual, 2)
		})

		Convey("Members have expected attributes", func() {
			So(response.Members[0].GuildId, ShouldEqual, "G123456")
			So(response.Members[1].GuildId, ShouldEqual, "G123456")

			So(response.Members[0].Nick, ShouldEqual, "nick 1")
			So(response.Members[1].Nick, ShouldEqual, "nick 2")

			So(response.Members[0].JoinedAt, ShouldEqual, "JA1234567")
			So(response.Members[1].JoinedAt, ShouldEqual, "JA2345678")

			So(response.Members[0].Deaf, ShouldBeFalse)
			So(response.Members[1].Deaf, ShouldBeTrue)

			So(response.Members[0].Mute, ShouldBeFalse)
			So(response.Members[1].Mute, ShouldBeTrue)
		})

		Convey("Members users have expected attributes", func() {
			user1 := response.Members[0].User
			user2 := response.Members[1].User

			So(user1, ShouldNotBeNil)
			So(user2, ShouldNotBeNil)

			So(user1.Id, ShouldEqual, "U123456")
			So(user2.Id, ShouldEqual, "U234567")

			So(user1.Username, ShouldEqual, "username 1")
			So(user2.Username, ShouldEqual, "username 2")

			So(user1.Email, ShouldEqual, "user1@test.com")
			So(user2.Email, ShouldEqual, "user2@test.com")

			So(user1.Discriminator, ShouldEqual, "user1discriminator")
			So(user2.Discriminator, ShouldEqual, "user2discriminator")

			So(user1.Avatar, ShouldEqual, "user1avatar")
			So(user2.Avatar, ShouldEqual, "user2avatar")

			So(user1.Verified, ShouldBeFalse)
			So(user2.Verified, ShouldBeTrue)

			So(user1.MFAEnabled, ShouldBeFalse)
			So(user2.MFAEnabled, ShouldBeTrue)

			So(user1.Bot, ShouldBeFalse)
			So(user2.Bot, ShouldBeTrue)

			So(user1.Token, ShouldEqual, "user1token")
			So(user2.Token, ShouldEqual, "user2token")
		})

		Convey("Members roles have expected attributes", func() {
			user1Roles := response.Members[0].Roles
			user2Roles := response.Members[1].Roles

			So(user1Roles, ShouldNotBeNil)
			So(user2Roles, ShouldNotBeNil)

			So(len(user1Roles), ShouldEqual, 3)
			So(len(user2Roles), ShouldEqual, 1)

			So(user1Roles[0].Id, ShouldEqual, "R123456")
			So(user1Roles[1].Id, ShouldEqual, "R234567")
			So(user1Roles[2].Id, ShouldEqual, "R345678")
			So(user2Roles[0].Id, ShouldEqual, "R456789")
		})
	})
}

func TestDiscordClient_GetAllMembers_OneMemberHasNoRoles(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("@me").Times(1).Return(&discordgo.User{ID: "U12345678"}, nil)
	mockClient.EXPECT().GetAllMembers("G123456", "", 1000).Times(1).Return(
		[]*discordgo.Member{
			{
				GuildID:  "G123456",
				Nick:     "nick 1",
				JoinedAt: "JA1234567",
				Deaf:     false,
				Mute:     false,
				User: &discordgo.User{
					ID:            "U123456",
					Username:      "username 1",
					Email:         "user1@test.com",
					Token:         "user1token",
					Discriminator: "user1discriminator",
					Avatar:        "user1avatar",
					Verified:      false,
					MFAEnabled:    false,
					Bot:           false,
				},
				Roles: []string{
					"R123456",
					"R234567",
					"R345678",
				},
			},
			{
				GuildID:  "G123456",
				Nick:     "nick 2",
				JoinedAt: "JA2345678",
				Deaf:     true,
				Mute:     true,
				User: &discordgo.User{
					ID:            "U234567",
					Username:      "username 2",
					Email:         "user2@test.com",
					Token:         "user2token",
					Discriminator: "user2discriminator",
					Avatar:        "user2avatar",
					Verified:      true,
					MFAEnabled:    true,
					Bot:           true,
				},
				Roles: []string{},
			},
		}, nil,
	)

	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, err := NewDiscordGatewayHandler(mockClient, mockRoleMap)
	Convey("Client construction err is nil", t, func() {
		So(err, ShouldBeNil)
	})

	response := &proto.GetMembersResponse{}
	err = client.GetAllMembers(context.Background(), &proto.GetAllMembersRequest{GuildId: "G123456", After: "", NumberPerPage: 1000}, response)

	Convey("Given a response", t, func() {
		Convey("Response should be valid with no error", func() {
			So(err, ShouldBeNil)
			So(response, ShouldNotBeNil)
		})

		Convey("Should be 2 members", func() {
			So(len(response.Members), ShouldEqual, 2)
		})

		Convey("Members have expected attributes", func() {
			So(response.Members[0].GuildId, ShouldEqual, "G123456")
			So(response.Members[1].GuildId, ShouldEqual, "G123456")

			So(response.Members[0].Nick, ShouldEqual, "nick 1")
			So(response.Members[1].Nick, ShouldEqual, "nick 2")

			So(response.Members[0].JoinedAt, ShouldEqual, "JA1234567")
			So(response.Members[1].JoinedAt, ShouldEqual, "JA2345678")

			So(response.Members[0].Deaf, ShouldBeFalse)
			So(response.Members[1].Deaf, ShouldBeTrue)

			So(response.Members[0].Mute, ShouldBeFalse)
			So(response.Members[1].Mute, ShouldBeTrue)
		})

		Convey("Members users have expected attributes", func() {
			user1 := response.Members[0].User
			user2 := response.Members[1].User

			So(user1, ShouldNotBeNil)
			So(user2, ShouldNotBeNil)

			So(user1.Id, ShouldEqual, "U123456")
			So(user2.Id, ShouldEqual, "U234567")

			So(user1.Username, ShouldEqual, "username 1")
			So(user2.Username, ShouldEqual, "username 2")

			So(user1.Email, ShouldEqual, "user1@test.com")
			So(user2.Email, ShouldEqual, "user2@test.com")

			So(user1.Discriminator, ShouldEqual, "user1discriminator")
			So(user2.Discriminator, ShouldEqual, "user2discriminator")

			So(user1.Avatar, ShouldEqual, "user1avatar")
			So(user2.Avatar, ShouldEqual, "user2avatar")

			So(user1.Verified, ShouldBeFalse)
			So(user2.Verified, ShouldBeTrue)

			So(user1.MFAEnabled, ShouldBeFalse)
			So(user2.MFAEnabled, ShouldBeTrue)

			So(user1.Bot, ShouldBeFalse)
			So(user2.Bot, ShouldBeTrue)

			So(user1.Token, ShouldEqual, "user1token")
			So(user2.Token, ShouldEqual, "user2token")
		})

		Convey("Members roles have expected attributes", func() {
			user1Roles := response.Members[0].Roles
			user2Roles := response.Members[1].Roles

			So(user1Roles, ShouldNotBeNil)
			So(user2Roles, ShouldBeNil)

			So(len(user1Roles), ShouldEqual, 3)
			So(len(user2Roles), ShouldEqual, 0)

			So(user1Roles[0].Id, ShouldEqual, "R123456")
			So(user1Roles[1].Id, ShouldEqual, "R234567")
			So(user1Roles[2].Id, ShouldEqual, "R345678")
		})
	})
}

//Same test as above but validating that the bot user isn't pulled in.  I don't really need a new test but it's nice
//for documentation purposes ;)
func TestDiscordClient_GetAllMembers_ExceptBotMember_ErrorOnGetUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("@me").Times(1).Return(nil, errors.New("dave I failed you on GetUser"))
	mockClient.EXPECT().GetAllMembers("G123456", "", 1000).Times(0)
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, err := NewDiscordGatewayHandler(mockClient, mockRoleMap)
	Convey("Client construction err is nil", t, func() {
		So(err, ShouldBeNil)
	})

	response := &proto.GetMembersResponse{}
	err = client.GetAllMembers(context.Background(), &proto.GetAllMembersRequest{GuildId: "G123456", After: "", NumberPerPage: 1000}, response)

	Convey("GetUser call errored out and the error is passed up", t, func() {
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you on GetUser")
		So(response, ShouldResemble, &proto.GetMembersResponse{})
	})
}

func TestDiscordClient_GetAllMembers_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("@me").Times(1).Return(&discordgo.User{ID: "U12345678"}, nil)
	mockClient.EXPECT().GetAllMembers(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("dave I failed you"))
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, err := NewDiscordGatewayHandler(mockClient, mockRoleMap)
	Convey("Client construction err is nil", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("Calling GetAllMembers returns the error from discord", t, func() {
		response := &proto.GetMembersResponse{}
		err := client.GetAllMembers(context.Background(), &proto.GetAllMembersRequest{GuildId: "G123456", After: "", NumberPerPage: 1000}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you")
	})
}

func TestDiscordClient_GetAllMembers_ErrorOnRolesUpdate_DuringClientConstruction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Return(nil, errors.New("dave I failed to update roles"))
	//</editor-fold>

	client, err := NewDiscordGatewayHandler(mockClient, mockRoleMap)
	Convey("Client construction err is not nil", t, func() {
		So(err, ShouldNotBeNil)
		So(client, ShouldBeNil)
	})
}

func TestDiscordClient_GetAllMembers_ErrorOnRolesUpdate_DuringGetAllMembersCall(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Return(nil, errors.New("dave I failed to update roles"))
	//</editor-fold>

	client := &discordGatewayHandler{client: mockClient, roleMap: mockRoleMap, lastRoleCall: time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)}

	Convey("Call to GetAllMembers produces error due to update roles issue", t, func() {
		response := &proto.GetMembersResponse{}
		err := client.GetAllMembers(context.Background(), &proto.GetAllMembersRequest{GuildId: "G123456", After: "", NumberPerPage: 1000}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed to update roles")
	})
}

func TestDiscordClient_GetAllRoles(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	response := &proto.GetRoleResponse{}
	client.GetAllRoles(context.Background(), &proto.GuildObjectRequest{GuildId: "G123456"}, response)

	Convey("Roles have expected attributes", t, func() {
		roles := response.Roles

		So(roles, ShouldNotBeNil)
		So(len(roles), ShouldEqual, 4)

		So(roles[0].Id, ShouldEqual, "R123456")
		So(roles[1].Id, ShouldEqual, "R234567")
		So(roles[2].Id, ShouldEqual, "R345678")
		So(roles[3].Id, ShouldEqual, "R456789")

		So(roles[0].Name, ShouldEqual, "role 1")
		So(roles[1].Name, ShouldEqual, "role 2")
		So(roles[2].Name, ShouldEqual, "role 3")
		So(roles[3].Name, ShouldEqual, "role 4")

		So(roles[0].Position, ShouldEqual, 1)
		So(roles[1].Position, ShouldEqual, 2)
		So(roles[2].Position, ShouldEqual, 3)
		So(roles[3].Position, ShouldEqual, 4)

		So(roles[0].Permissions, ShouldEqual, 1)
		So(roles[1].Permissions, ShouldEqual, 2)
		So(roles[2].Permissions, ShouldEqual, 3)
		So(roles[3].Permissions, ShouldEqual, 4)

		So(roles[0].Mentionable, ShouldEqual, false)
		So(roles[1].Mentionable, ShouldEqual, true)
		So(roles[2].Mentionable, ShouldEqual, false)
		So(roles[3].Mentionable, ShouldEqual, true)

		So(roles[0].Managed, ShouldEqual, false)
		So(roles[1].Managed, ShouldEqual, true)
		So(roles[2].Managed, ShouldEqual, true)
		So(roles[3].Managed, ShouldEqual, false)

		So(roles[0].Hoist, ShouldEqual, false)
		So(roles[1].Hoist, ShouldEqual, true)
		So(roles[2].Hoist, ShouldEqual, false)
		So(roles[3].Hoist, ShouldEqual, true)

		So(roles[0].Color, ShouldEqual, 1)
		So(roles[1].Color, ShouldEqual, 2)
		So(roles[2].Color, ShouldEqual, 3)
		So(roles[3].Color, ShouldEqual, 4)
	})
}

func TestDiscordClient_GetAllRoles_CacheTest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    1,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    2,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    3,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    4,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	response := &proto.GetRoleResponse{}
	client.GetAllRoles(context.Background(), &proto.GuildObjectRequest{GuildId: "G123456"}, response)
	client.GetAllRoles(context.Background(), &proto.GuildObjectRequest{GuildId: "G123456"}, response)
}

func TestDiscordClient_GetAllRoles_ArrayError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	//</editor-fold>

	client := &discordGatewayHandler{client: mockClient, roleMap: mockRoleMap, lastRoleCall: time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)}

	Convey("Given Discords positioning starting at 0", t, func() {
		response := &proto.GetRoleResponse{}
		client.GetAllRoles(context.Background(), &proto.GuildObjectRequest{GuildId: "G123456"}, response)

		roles := response.Roles

		So(roles, ShouldNotBeNil)
		So(len(roles), ShouldEqual, 4)

		So(roles[0].Id, ShouldEqual, "R123456")
		So(roles[1].Id, ShouldEqual, "R234567")
		So(roles[2].Id, ShouldEqual, "R345678")
		So(roles[3].Id, ShouldEqual, "R456789")

		So(roles[0].Name, ShouldEqual, "role 1")
		So(roles[1].Name, ShouldEqual, "role 2")
		So(roles[2].Name, ShouldEqual, "role 3")
		So(roles[3].Name, ShouldEqual, "role 4")

		So(roles[0].Position, ShouldEqual, 0)
		So(roles[1].Position, ShouldEqual, 1)
		So(roles[2].Position, ShouldEqual, 2)
		So(roles[3].Position, ShouldEqual, 3)

		So(roles[0].Permissions, ShouldEqual, 1)
		So(roles[1].Permissions, ShouldEqual, 2)
		So(roles[2].Permissions, ShouldEqual, 3)
		So(roles[3].Permissions, ShouldEqual, 4)

		So(roles[0].Mentionable, ShouldEqual, false)
		So(roles[1].Mentionable, ShouldEqual, true)
		So(roles[2].Mentionable, ShouldEqual, false)
		So(roles[3].Mentionable, ShouldEqual, true)

		So(roles[0].Managed, ShouldEqual, false)
		So(roles[1].Managed, ShouldEqual, true)
		So(roles[2].Managed, ShouldEqual, true)
		So(roles[3].Managed, ShouldEqual, false)

		So(roles[0].Hoist, ShouldEqual, false)
		So(roles[1].Hoist, ShouldEqual, true)
		So(roles[2].Hoist, ShouldEqual, false)
		So(roles[3].Hoist, ShouldEqual, true)

		So(roles[0].Color, ShouldEqual, 1)
		So(roles[1].Color, ShouldEqual, 2)
		So(roles[2].Color, ShouldEqual, 3)
		So(roles[3].Color, ShouldEqual, 4)
	})
}

func TestDiscordClient_GetAllRoles_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Return(nil, errors.New("dave I failed you"))
	//</editor-fold>

	client := &discordGatewayHandler{client: mockClient, roleMap: mockRoleMap, lastRoleCall: time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)}

	Convey("Calling GetAllRoles returns the error from Discrod", t, func() {
		response := &proto.GetRoleResponse{}
		err := client.GetAllRoles(context.Background(), &proto.GuildObjectRequest{GuildId: "G123456"}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you")
	})
}

func TestDiscordClient_GetUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("U123456").Times(1).Return(
		&discordgo.User{
			ID:            "U123456",
			Username:      "username1",
			Email:         "user@test.com",
			Discriminator: "userdescriminator",
			Avatar:        "useravatar",
			Verified:      false,
			MFAEnabled:    false,
			Bot:           false,
			Token:         "usertoken",
		}, nil,
	)
	//</editor-fold>

	client := &discordGatewayHandler{client: mockClient}

	Convey("Requesting a user gets the users details", t, func() {
		response := &proto.GetUserResponse{}
		err := client.GetUser(context.Background(), &proto.GetUserRequest{UserId: "U123456"}, response)

		So(err, ShouldBeNil)
		So(response.User, ShouldNotBeNil)

		So(response.User, ShouldResemble, &proto.User{
			Id:            "U123456",
			Username:      "username1",
			Email:         "user@test.com",
			Discriminator: "userdescriminator",
			Avatar:        "useravatar",
			Verified:      false,
			MFAEnabled:    false,
			Bot:           false,
			Token:         "usertoken",
		})
	})
}

func TestDiscordClient_GetUser_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetUser("U123456").Times(1).Return(nil, errors.New("dave I failed you"))
	//</editor-fold>

	client := &discordGatewayHandler{client: mockClient}

	Convey("Requesting a user gets the users details", t, func() {
		response := &proto.GetUserResponse{}
		err := client.GetUser(context.Background(), &proto.GetUserRequest{UserId: "U123456"}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you")
	})
}

func TestDiscordClient_RemoveMemberRoles(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
			[]*discordgo.Role{
				{
					ID:          "R123456",
					Name:        "role 1",
					Position:    0,
					Permissions: 1,
					Mentionable: false,
					Managed:     false,
					Hoist:       false,
					Color:       1,
				},
				{
					ID:          "R234567",
					Name:        "role 2",
					Position:    1,
					Permissions: 2,
					Mentionable: true,
					Managed:     true,
					Hoist:       true,
					Color:       2,
				},
				{
					ID:          "R345678",
					Name:        "role 3",
					Position:    2,
					Permissions: 3,
					Mentionable: false,
					Managed:     true,
					Hoist:       false,
					Color:       3,
				},
				{
					ID:          "R456789",
					Name:        "role 4",
					Position:    3,
					Permissions: 4,
					Mentionable: true,
					Managed:     false,
					Hoist:       true,
					Color:       4,
				},
			}, nil,
		),
	)
	mockClient.EXPECT().RemoveMemberRole("G123456", "U123456", "role 1").Times(1).Return(nil)
	mockClient.EXPECT().RemoveMemberRole("G123456", "U123456", "role 2").Times(1).Return(nil)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Remove roles (role 1 and role 2)", t, func() {
		response := &proto.UpdateMemberResponse{}
		err := client.UpdateMember(context.Background(),
			&proto.UpdateMemberRequest{
				GuildId:   "G123456",
				UserId:    "U123456",
				Operation: proto.MemberUpdateOperation_REMOVE_ROLES,
				RoleIds:   []string{"R123456", "R234567"},
			},
			response,
		)

		So(err, ShouldBeNil)
		So(response.Success, ShouldBeTrue)
	})
}

func TestDiscordClient_RemoveMemberRole_ErrorFirstRemove(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
			[]*discordgo.Role{
				{
					ID:          "R123456",
					Name:        "role 1",
					Position:    0,
					Permissions: 1,
					Mentionable: false,
					Managed:     false,
					Hoist:       false,
					Color:       1,
				},
				{
					ID:          "R234567",
					Name:        "role 2",
					Position:    1,
					Permissions: 2,
					Mentionable: true,
					Managed:     true,
					Hoist:       true,
					Color:       2,
				},
				{
					ID:          "R345678",
					Name:        "role 3",
					Position:    2,
					Permissions: 3,
					Mentionable: false,
					Managed:     true,
					Hoist:       false,
					Color:       3,
				},
				{
					ID:          "R456789",
					Name:        "role 4",
					Position:    3,
					Permissions: 4,
					Mentionable: true,
					Managed:     false,
					Hoist:       true,
					Color:       4,
				},
			}, nil,
		),
	)
	mockClient.EXPECT().RemoveMemberRole("G123456", "U123456", "role 1").Times(1).Return(errors.New("dave I failed you on the first try"))
	mockClient.EXPECT().RemoveMemberRole("G123456", "U123456", "role 2").Times(1).Return(nil)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Update a members roles with an error", t, func() {
		response := &proto.UpdateMemberResponse{}
		err := client.UpdateMember(context.Background(),
			&proto.UpdateMemberRequest{
				GuildId:   "G123456",
				UserId:    "U123456",
				Operation: proto.MemberUpdateOperation_REMOVE_ROLES,
				RoleIds:   []string{"R123456", "R234567"},
			},
			response,
		)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you on the first try")
	})
}

func TestDiscordClient_RemoveMemberRole_ErrorSecondRemove(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	gomock.InOrder(
		mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
			[]*discordgo.Role{
				{
					ID:          "R123456",
					Name:        "role 1",
					Position:    0,
					Permissions: 1,
					Mentionable: false,
					Managed:     false,
					Hoist:       false,
					Color:       1,
				},
				{
					ID:          "R234567",
					Name:        "role 2",
					Position:    1,
					Permissions: 2,
					Mentionable: true,
					Managed:     true,
					Hoist:       true,
					Color:       2,
				},
				{
					ID:          "R345678",
					Name:        "role 3",
					Position:    2,
					Permissions: 3,
					Mentionable: false,
					Managed:     true,
					Hoist:       false,
					Color:       3,
				},
				{
					ID:          "R456789",
					Name:        "role 4",
					Position:    3,
					Permissions: 4,
					Mentionable: true,
					Managed:     false,
					Hoist:       true,
					Color:       4,
				},
			}, nil,
		),
	)
	mockClient.EXPECT().RemoveMemberRole("G123456", "U123456", "role 1").Times(1).Return(nil)
	mockClient.EXPECT().RemoveMemberRole("G123456", "U123456", "role 2").Times(1).Return(errors.New("dave I failed you on the second try"))
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Update a members roles with an error", t, func() {
		response := &proto.UpdateMemberResponse{}
		err := client.UpdateMember(context.Background(),
			&proto.UpdateMemberRequest{
				GuildId:   "G123456",
				UserId:    "U123456",
				Operation: proto.MemberUpdateOperation_REMOVE_ROLES,
				RoleIds:   []string{"R123456", "R234567"},
			},
			response,
		)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you on the second try")
	})
}

func TestDiscordClient_UpdateMember(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().UpdateMember("G123456", "U123456", []string{"role 1", "role 2"}).Times(1).Return(nil)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Update a members roles from 3 to 2 roles", t, func() {
		response := &proto.UpdateMemberResponse{}
		err := client.UpdateMember(context.Background(),
			&proto.UpdateMemberRequest{
				GuildId:   "G123456",
				UserId:    "U123456",
				Operation: proto.MemberUpdateOperation_ADD_OR_UPDATE_ROLES,
				RoleIds:   []string{"R123456", "R234567"},
			},
			response,
		)

		So(err, ShouldBeNil)
		So(response.Success, ShouldBeTrue)
	})
}

func TestDiscordClient_UpdateMember_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations and results">
	mockClient.EXPECT().GetAllRoles(gomock.Any()).Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().UpdateMember("G123456", "U123456", []string{"role 1", "role 2"}).Times(1).Return(errors.New("dave I failed you"))
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Update a members roles with an error", t, func() {
		response := &proto.UpdateMemberResponse{}
		err := client.UpdateMember(context.Background(),
			&proto.UpdateMemberRequest{
				GuildId:   "G123456",
				UserId:    "U123456",
				Operation: proto.MemberUpdateOperation_ADD_OR_UPDATE_ROLES,
				RoleIds:   []string{"R123456", "R234567"},
			},
			response,
		)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you")
	})
}

func TestDiscordClient_CreateRole(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations an results>
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().CreateRole("G123456").Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().EditRole("G123456", "R567890", "role 5", 1, 1, true, true).Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "role 5",
			Managed:     false,
			Position:    1,
			Permissions: 1,
			Color:       1,
			Hoist:       true,
			Mentionable: true,
		}, nil,
	)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Given a client create a role with the given attributes", t, func() {
		response := &proto.CreateRolesResponse{}
		err := client.CreateRole(context.Background(), &proto.CreateRoleRequest{
			Mentionable: true,
			Hoist:       true,
			Color:       1,
			Permissions: 1,
			Name:        "role 5",
			GuildId:     "G123456",
		}, response)

		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)
		So(response.RoleId, ShouldEqual, "R567890")
	})
}

func TestDiscordClient_CreateRole_ResponseValidationFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations an results>
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().CreateRole("G123456").Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().EditRole("G123456", "R567890", "role 5", 1, 1, true, true).Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().DeleteRole("G123456", "R567890").Times(1).Return(nil)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Given a client fail role creation due to invalid response from discord", t, func() {
		response := &proto.CreateRolesResponse{}
		err := client.CreateRole(context.Background(), &proto.CreateRoleRequest{
			Mentionable: true,
			Hoist:       true,
			Color:       1,
			Permissions: 1,
			Name:        "role 5",
			GuildId:     "G123456",
		}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "role create failed due to invalid response from discord")
	})
}

func TestDiscordClient_CreateRole_ResponseValidationFailure_ErrorOnDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations an results>
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().CreateRole("G123456").Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().EditRole("G123456", "R567890", "role 5", 1, 1, true, true).Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().DeleteRole("G123456", "R567890").Times(1).Return(errors.New("dave I failed to delete after you validated"))
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Given a client fail role delete after invalid check", t, func() {
		response := &proto.CreateRolesResponse{}
		err := client.CreateRole(context.Background(), &proto.CreateRoleRequest{
			Mentionable: true,
			Hoist:       true,
			Color:       1,
			Permissions: 1,
			Name:        "role 5",
			GuildId:     "G123456",
		}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "attempted to delete role due to invalid response but received error (dave I failed to delete after you validated)")
	})
}

func TestDiscorClient_CreateRole_ErrorOnCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations an results>
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().CreateRole("G123456").Times(1).Return(nil, errors.New("dave I failed you on create"))
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Given a client fail to create a role", t, func() {
		response := &proto.CreateRolesResponse{}
		err := client.CreateRole(context.Background(), &proto.CreateRoleRequest{
			Mentionable: true,
			Hoist:       true,
			Color:       1,
			Permissions: 1,
			Name:        "role 5",
			GuildId:     "G123456",
		}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you on create")
	})
}

func TestDiscorClient_CreateRole_ErrorOnEdit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations an results>
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().CreateRole("G123456").Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().EditRole("G123456", "R567890", "role 5", 1, 1, true, true).Times(1).Return(nil, errors.New("dave I failed you on edit"))
	mockClient.EXPECT().DeleteRole("G123456", "R567890").Times(1).Return(nil)
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Given a client fail to edit a role", t, func() {
		response := &proto.CreateRolesResponse{}
		err := client.CreateRole(context.Background(), &proto.CreateRoleRequest{
			Mentionable: true,
			Hoist:       true,
			Color:       1,
			Permissions: 1,
			Name:        "role 5",
			GuildId:     "G123456",
		}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you on edit")
	})
}

func TestDiscorClient_CreateRole_ErrorOnEditAndDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockClient := discord.NewMockDiscordClient(mockCtrl)
	mockRoleMap := discord.NewRoleMap("G123456", mockClient)
	SetDefaultFailureMode(FailureHalts)
	defer mockCtrl.Finish()

	//<editor-fold desc="Mock Expectations an results>
	mockClient.EXPECT().GetAllRoles("G123456").Times(1).Return(
		[]*discordgo.Role{
			{
				ID:          "R123456",
				Name:        "role 1",
				Position:    0,
				Permissions: 1,
				Mentionable: false,
				Managed:     false,
				Hoist:       false,
				Color:       1,
			},
			{
				ID:          "R234567",
				Name:        "role 2",
				Position:    1,
				Permissions: 2,
				Mentionable: true,
				Managed:     true,
				Hoist:       true,
				Color:       2,
			},
			{
				ID:          "R345678",
				Name:        "role 3",
				Position:    2,
				Permissions: 3,
				Mentionable: false,
				Managed:     true,
				Hoist:       false,
				Color:       3,
			},
			{
				ID:          "R456789",
				Name:        "role 4",
				Position:    3,
				Permissions: 4,
				Mentionable: true,
				Managed:     false,
				Hoist:       true,
				Color:       4,
			},
		}, nil,
	)
	mockClient.EXPECT().CreateRole("G123456").Times(1).Return(
		&discordgo.Role{
			ID:          "R567890",
			Name:        "default role name",
			Managed:     false,
			Position:    5,
			Permissions: 5,
			Color:       5,
			Hoist:       false,
			Mentionable: false,
		}, nil,
	)
	mockClient.EXPECT().EditRole("G123456", "R567890", "role 5", 1, 1, true, true).Times(1).Return(nil, errors.New("dave I failed you on edit"))
	mockClient.EXPECT().DeleteRole("G123456", "R567890").Times(1).Return(errors.New("dave I failed you on delete"))
	//</editor-fold>

	client, _ := NewDiscordGatewayHandler(mockClient, mockRoleMap)

	Convey("Given a client fail to edit and delete a role", t, func() {
		response := &proto.CreateRolesResponse{}
		err := client.CreateRole(context.Background(), &proto.CreateRoleRequest{
			Mentionable: true,
			Hoist:       true,
			Color:       1,
			Permissions: 1,
			Name:        "role 5",
			GuildId:     "G123456",
		}, response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "edit failure (dave I failed you on edit), delete failure (dave I failed you on delete)")
	})
}

//TODO: Define and implement these
func TestDiscordClient_DeleteRole(t *testing.T) {

}

func TestDiscordClient_DeleteRole_Error(t *testing.T) {

}
