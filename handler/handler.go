package handler

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/chremoas/discord-gateway/discord"
	proto "github.com/chremoas/discord-gateway/proto"
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

// There has got to be a better way to do this.
func (dgh *discordGatewayHandler) GetMessages(ctx context.Context, request *proto.GetMessagesRequest, response *proto.GetMessagesResponse) error {
	var mentions []*proto.User
	var attachments []*proto.MessageAttachment
	var embeds []*proto.MessageEmbed
	var reactions []*proto.MessageReactions
	var messageEmbedFields []*proto.MessageEmbedField

	fmt.Printf("Pre-ChannelMessages\n")
	messages, err := dgh.client.ChannelMessages(request.ChannelID, int(request.Limit), request.BeforeID, request.AfterID, request.AroundID)
	if (err != nil) {
		fmt.Printf("ChannelMessages error: %s\n", err.Error())
		return err
	}
	fmt.Printf("messages: %+v\n", messages)

	for m := range messages {
		mIt := messages[m]
		for men := range mIt.Mentions {
			menIt := mIt.Mentions[men]
			mentions = append(mentions, &proto.User{
				Id:            menIt.ID,
				Email:         menIt.Email,
				Username:      menIt.Username,
				Avatar:        menIt.Avatar,
				Discriminator: menIt.Discriminator,
				Token:         menIt.Token,
				Verified:      menIt.Verified,
				MFAEnabled:    menIt.MFAEnabled,
				Bot:           menIt.Bot,
			})
		}

		for att := range mIt.Attachments {
			attIt := mIt.Attachments[att]
			attachments = append(attachments, &proto.MessageAttachment{
				ID:       attIt.ID,
				URL:      attIt.URL,
				ProxyURL: attIt.ProxyURL,
				Filename: attIt.Filename,
				Width:    int64(attIt.Width),
				Height:   int64(attIt.Height),
				Size:     int64(attIt.Size),
			})
		}


		for emb := range mIt.Embeds {
			embIt := mIt.Embeds[emb]
			for mef := range embIt.Fields {
				mefIt := embIt.Fields[mef]
				messageEmbedFields = append(messageEmbedFields, &proto.MessageEmbedField{
					Name: mefIt.Name,
					Value: mefIt.Value,
					Inline: mefIt.Inline,
				})
			}

			embeds = append(embeds, &proto.MessageEmbed{
				URL: embIt.URL,
				Type: embIt.Type,
				Title: embIt.Title,
				Description: embIt.Description,
				Timestamp: embIt.Timestamp,
				Color: int64(embIt.Color),
				Footer: &proto.MessageEmbedFooter{
					Text: embIt.Footer.Text,
					IconURL: embIt.Footer.IconURL,
					ProxyIconURL: embIt.Footer.ProxyIconURL,
				},
				Image: &proto.MessageEmbedItem{
					URL: embIt.Image.URL,
					ProxyURL: embIt.Image.ProxyURL,
					Width: int64(embIt.Image.Width),
					Height: int64(embIt.Image.Height),
				},
				Thumbnail: &proto.MessageEmbedItem{
					URL: embIt.Thumbnail.URL,
					ProxyURL: embIt.Thumbnail.ProxyURL,
					Width: int64(embIt.Thumbnail.Width),
					Height: int64(embIt.Thumbnail.Height),
				},
				Video: &proto.MessageEmbedItem{
					URL: embIt.Video.URL,
					ProxyURL: embIt.Video.ProxyURL,
					Width: int64(embIt.Video.Width),
					Height: int64(embIt.Video.Height),
				},
				Provider: &proto.MessageEmbedProvider{
					URL: embIt.Provider.URL,
					Name: embIt.Provider.Name,

				},
				Author: &proto.MessageEmbedAuthor{
					URL: embIt.Author.URL,
					Name: embIt.Author.Name,
					IconURL: embIt.Author.IconURL,
					ProxyIconURL: embIt.Author.ProxyIconURL,

				},
				Fields: messageEmbedFields,
			})
		}

		for rea := range mIt.Reactions {
			reaIt := mIt.Reactions[rea]
			reactions = append(reactions, &proto.MessageReactions{
				Count: int64(reaIt.Count),
				Me: reaIt.Me,
				Emoji: &proto.Emoji{
					ID: reaIt.Emoji.ID,
					Name: reaIt.Emoji.Name,
					Roles: reaIt.Emoji.Roles,
					Managed: reaIt.Emoji.Managed,
					RequireColons: reaIt.Emoji.RequireColons,
					Animated: reaIt.Emoji.Animated,
				},
			})
		}

		message := &proto.Message{
			ID:              mIt.ID,
			ChannelID:       mIt.ChannelID,
			Content:         mIt.Content,
			Timestamp:       string(mIt.Timestamp),
			EditedTimestamp: string(mIt.EditedTimestamp),
			MentionRoles:    mIt.MentionRoles,
			Tts:             mIt.Tts,
			MentionEveryone: mIt.MentionEveryone,
			Author: &proto.User{
				Id:            mIt.Author.ID,
				Email:         mIt.Author.Email,
				Username:      mIt.Author.Username,
				Avatar:        mIt.Author.Avatar,
				Discriminator: mIt.Author.Discriminator,
				Token:         mIt.Author.Token,
				Verified:      mIt.Author.Verified,
				MFAEnabled:    mIt.Author.MFAEnabled,
				Bot:           mIt.Author.Bot,
			},
			Attachments: attachments,
			Embeds:      embeds,
			Mentions:    mentions,
			Reactions:   reactions,
			Type:        int64(mIt.Type),
		}

		response.Messages = append(response.Messages, message)
	}
	return err
}

func (dgh *discordGatewayHandler) BulkDeleteMessages(ctx context.Context, request *proto.BulkDeleteMessagesRequest, response *proto.NilMessage) error {
	return dgh.client.ChannelMessagesBulkDelete(request.ChannelID, request.Messages)
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

	fmt.Printf("Pre-GetAllMembers\n")
	members, err := dgh.client.GetAllMembers(dgh.discordServerId, request.After, int(request.NumberPerPage))
	fmt.Printf("Post-GetAllMembers\n")
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
				Nick:          member.Nick,
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

func (dgh *discordGatewayHandler) GetAllMembersAsSlice(ctx context.Context, request *proto.GetAllMembersRequest, response *proto.GetMembersResponse) error {
	err := dgh.updateRoles()
	if err != nil {
		return err
	}

	members, err := dgh.client.GetAllMembersAsSlice(dgh.discordServerId)
	if err != nil {
		return err
	}

	for _, member := range members {
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

func (dgh *discordGatewayHandler) EditRole(ctx context.Context, request *proto.EditRoleRequest, response *proto.EditRoleResponse) error {
	role := dgh.roleMap.GetRoleByName(request.Name)

	if role == nil {
		return fmt.Errorf("Role doesn't exist: %s\n", request.Name)
	}

	newRole, err := dgh.client.EditRole(
		dgh.discordServerId,
		role.ID,
		request.Name,
		int(request.Color),
		int(request.Perm),
		request.Hoist,
		request.Mention,
	)

	if err != nil {
		return err
	}

	fmt.Printf("newRole: %+v\n", newRole)

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
