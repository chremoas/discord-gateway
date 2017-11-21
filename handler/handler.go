package handler

import (
	"errors"
	"fmt"
	"github.com/abaeve/discord-gateway/discord"
	proto "github.com/abaeve/discord-gateway/proto"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/net/context"
	"sort"
	"time"
)

type discordGatewayHandler struct {
	client  discord.DiscordClient
	roleMap discord.RoleMap

	lastRoleCall time.Time
}

//Provide a mechanism to provide roles in order of position as reported by Discord
type ByPosition []*proto.Role

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func (dgh *discordGatewayHandler) UpdateMember(ctx context.Context, request *proto.UpdateMemberRequest, response *proto.UpdateMemberResponse) error {
	var err error

	roles := make([]string, len(request.RoleIds))

	for idx, roleId := range request.RoleIds {
		roles[idx] = dgh.roleMap.GetRoleName(roleId)
	}

	switch request.Operation {
	case proto.MemberUpdateOperation_REMOVE_ROLES:
		var removeErr error

		for _, roleName := range roles {
			currentErr := dgh.client.RemoveMemberRole(request.GuildId, request.UserId, roleName)
			if currentErr != nil {
				removeErr = currentErr
			}

			err = removeErr
		}

		break
	case proto.MemberUpdateOperation_ADD_OR_UPDATE_ROLES:
		err = dgh.client.UpdateMember(request.GuildId, request.UserId, roles)
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

	members, err := dgh.client.GetAllMembers(request.GuildId, request.After, int(request.NumberPerPage))
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

		for _, roleName := range member.Roles {
			role := dgh.roleMap.GetRoleByName(roleName)

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
	role, err := dgh.client.CreateRole(request.GuildId)
	if err != nil {
		return err
	}

	editedRole, err := dgh.client.EditRole(request.GuildId, role.ID, request.Name, int(request.Color), int(request.Permissions), request.Hoist, request.Mentionable)
	if err != nil {
		deleteErr := dgh.client.DeleteRole(request.GuildId, role.ID)
		if deleteErr != nil {
			return errors.New(fmt.Sprintf("edit failure (%s), delete failure (%s)", err.Error(), deleteErr.Error()))
		}

		return err
	}

	//Now validate the edited role
	if !validateRole(request, editedRole) {
		err = dgh.client.DeleteRole(request.GuildId, role.ID)
		if err != nil {
			return errors.New(fmt.Sprintf("attempted to delete role due to invalid response but received error (%s)", err.Error()))
		}

		return errors.New("role create failed due to invalid response from discord")
	}

	response.RoleId = editedRole.ID

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
	if time.Now().Sub(dgh.lastRoleCall) >= time.Minute*5 {
		dgh.lastRoleCall = time.Now()
		return dgh.roleMap.UpdateRoles()
	}

	return nil
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

func NewDiscordGatewayHandler(client discord.DiscordClient, roleMap discord.RoleMap) (proto.DiscordGatewayHandler, error) {
	err := roleMap.UpdateRoles()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error calling roleMap.UpdateRoles (%s)", err.Error()))
	}
	return &discordGatewayHandler{client: client, roleMap: roleMap, lastRoleCall: time.Now()}, nil
}
