package handler

import (
	"errors"
	"fmt"
	"github.com/chremoas/discord-gateway/discord"
	proto "github.com/chremoas/discord-gateway/proto"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/net/context"
	"sort"
	"time"
)

type discordGatewayHandler struct {
	discordServerId string

	client  discord.DiscordClient
	roleMap discord.RoleMap

	lastRoleCall time.Time
}

//Provide a mechanism to provide roles in order of position as reported by Discord
type ByPosition []*proto.Role

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func (dgh *discordGatewayHandler) SendMessage(ctx context.Context, request *proto.SendMessageRequest, response *proto.NilMessage) error {
	return dgh.client.SendMessage(request.ChannelId, request.Message)
}

func (dgh *discordGatewayHandler) UpdateMember(ctx context.Context, request *proto.UpdateMemberRequest, response *proto.UpdateMemberResponse) error {
	var err error

	roleIds := request.RoleIds

	switch request.Operation {
	case proto.MemberUpdateOperation_REMOVE_ROLES:
		var removeErr error

		for _, roleId := range roleIds {
			currentErr := dgh.client.RemoveMemberRole(dgh.discordServerId, request.UserId, roleId)
			if currentErr != nil {
				removeErr = currentErr
			}

			err = removeErr
		}

		break
	case proto.MemberUpdateOperation_ADD_OR_UPDATE_ROLES:
		err = dgh.client.UpdateMember(dgh.discordServerId, request.UserId, roleIds)
		break
	}

	if err == nil {
		response.Success = true
	}

	return err
}

func (dgh *discordGatewayHandler) GetAllMembers(ctx context.Context, request *proto.GetAllMembersRequest, response *proto.GetMembersResponse) error {
	err := dgh.updateRoles()
	if err != nil {
		return err
	}

	me, err := dgh.client.GetUser("@me")
	if err != nil {
		return err
	}

	fmt.Printf("Pre-GetAllMembers")
	members, err := dgh.client.GetAllMembers(dgh.discordServerId, request.After, int(request.NumberPerPage))
	fmt.Printf("Post-GetAllMembers")
	if err != nil {
		return err
	}

	for _, member := range members {
		if member.User.ID == me.ID {
			continue
		}
		protoMember := &proto.Member{
			GuildId: member.GuildID,
			User: &proto.User{
				Id:            member.User.ID,
				Token:         member.User.Token,
				Bot:           member.User.Bot,
				MFAEnabled:    member.User.MFAEnabled,
				Verified:      member.User.Verified,
				Avatar:        member.User.Avatar,
				Discriminator: member.User.Discriminator,
				Email:         member.User.Email,
				Username:      member.User.Username,
			},
			Mute:     member.Mute,
			Deaf:     member.Deaf,
			JoinedAt: member.JoinedAt,
			Nick:     member.Nick,
		}

		for _, roleId := range member.Roles {
			role := dgh.roleMap.GetRoleById(roleId)

			protoMember.Roles = append(protoMember.Roles, &proto.Role{
				Id:          role.ID,
				Name:        role.Name,
				Position:    int32(role.Position),
				Permissions: int32(role.Permissions),
				Mentionable: role.Mentionable,
				Managed:     role.Mentionable,
				Hoist:       role.Hoist,
				Color:       int32(role.Color),
			})
		}

		response.Members = append(response.Members, protoMember)
		sort.Sort(ByPosition(protoMember.Roles))
	}

	return nil
}

func (dgh *discordGatewayHandler) GetAllRoles(ctx context.Context, request *proto.GuildObjectRequest, response *proto.GetRoleResponse) error {
	err := dgh.updateRoles()
	if err != nil {
		return err
	}

	allRoles := dgh.roleMap.GetRoles()

	for key := range allRoles {
		response.Roles = append(response.Roles, &proto.Role{
			Id:          allRoles[key].ID,
			Name:        allRoles[key].Name,
			Managed:     allRoles[key].Managed,
			Mentionable: allRoles[key].Mentionable,
			Hoist:       allRoles[key].Hoist,
			Permissions: int32(allRoles[key].Permissions),
			Position:    int32(allRoles[key].Position),
			Color:       int32(allRoles[key].Color),
		})
	}

	sort.Sort(ByPosition(response.Roles))

	return nil
}

func (dgh *discordGatewayHandler) CreateRole(ctx context.Context, request *proto.CreateRoleRequest, response *proto.CreateRolesResponse) error {
	err := dgh.roleMap.UpdateRoles()
	if err != nil {
		return err
	}

	allRoles := dgh.roleMap.GetRoles()

	for key := range allRoles {
		if allRoles[key].Name == request.Name {
			return fmt.Errorf("The role '%s' already exists", allRoles[key].Name)
		}
	}

	role, err := dgh.client.CreateRole(dgh.discordServerId)
	if err != nil {
		return err
	}

	editedRole, err := dgh.client.EditRole(dgh.discordServerId, role.ID, request.Name, int(request.Color), int(request.Permissions), request.Hoist, request.Mentionable)
	if err != nil {
		deleteErr := dgh.client.DeleteRole(dgh.discordServerId, role.ID)
		if deleteErr != nil {
			return errors.New(fmt.Sprintf("edit failure (%s), delete failure (%s)", err.Error(), deleteErr.Error()))
		}

		return err
	}

	//Now validate the edited role
	if !validateRole(request, editedRole) {
		err = dgh.client.DeleteRole(dgh.discordServerId, role.ID)
		if err != nil {
			return errors.New(fmt.Sprintf("attempted to delete role due to invalid response but received error (%s)", err.Error()))
		}

		return errors.New("role create failed due to invalid response from discord")
	}

	response.RoleId = editedRole.ID

	return nil
}

func (dgh *discordGatewayHandler) DeleteRole(ctx context.Context, request *proto.DeleteRoleRequest, response *proto.DeleteRoleResponse) error {
	role := dgh.roleMap.GetRoleByName(request.Name)

	if role == nil {
		return fmt.Errorf("Role doesn't exist: %s\n", request.Name)
	}

	err := dgh.client.DeleteRole(dgh.discordServerId, role.ID)
	if err != nil {
		return err
	}

	return nil
}

func (dgh *discordGatewayHandler) GetUser(ctx context.Context, request *proto.GetUserRequest, response *proto.GetUserResponse) error {
	user, err := dgh.client.GetUser(request.UserId)
	if err != nil {
		return err
	}

	response.User = &proto.User{
		Id:            user.ID,
		Discriminator: user.Discriminator,
		Email:         user.Email,
		Username:      user.Username,
		Token:         user.Token,
		Bot:           user.Bot,
		MFAEnabled:    user.MFAEnabled,
		Verified:      user.Verified,
		Avatar:        user.Avatar,
	}

	return nil
}

func (dgh *discordGatewayHandler) updateRoles() error {
	// Going to not cache for now. Not sure we even need that or not. We'll address this later.
	//if time.Now().Sub(dgh.lastRoleCall) >= time.Minute*5 {
	//	dgh.lastRoleCall = time.Now()
		return dgh.roleMap.UpdateRoles()
	//}
	//
	//return nil
}

func validateRole(request *proto.CreateRoleRequest, role *discordgo.Role) bool {
	valid := true

	valid = valid && role.Permissions == int(request.Permissions)
	valid = valid && role.Hoist == request.Hoist
	valid = valid && role.Mentionable == request.Mentionable
	valid = valid && role.Color == int(request.Color)
	valid = valid && role.Name == request.Name

	return valid
}

func NewDiscordGatewayHandler(discordServerId string, client discord.DiscordClient, roleMap discord.RoleMap) (proto.DiscordGatewayHandler, error) {
	err := roleMap.UpdateRoles()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error calling roleMap.UpdateRoles (%s)", err.Error()))
	}
	return &discordGatewayHandler{discordServerId: discordServerId, client: client, roleMap: roleMap, lastRoleCall: time.Now()}, nil
}
