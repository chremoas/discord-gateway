package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chremoas/services-common/config"
	"go.uber.org/zap"
	"sync"
	"fmt"
	"strconv"
)

// This is a very thin wrapper around the discordgo api for testability purposes
// and to easily keep track of what endpoints are being consumed
type DiscordClient interface {
	SendMessage(channelId, message string) error
	ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed)
	UpdateMember(guildID, userID string, roles []string) error
	RemoveMemberRole(guildID, userID, role string) error
	GetAllMembers(guildID, after string, limit int) ([]*discordgo.Member, error)
	GetAllMembersAsSlice(guildID string) ([]*discordgo.Member, error)
	GetAllRoles(guildID string) ([]*discordgo.Role, error)
	GetUser(userID string) (*discordgo.User, error)
	CreateRole(guildId string) (*discordgo.Role, error)
	DeleteRole(guildId, roleId string) error
	EditRole(guildId, roleId, name string, color, perm int, hoist, mention bool) (*discordgo.Role, error)
}

type client struct {
	session *discordgo.Session
	mutex   sync.Mutex
	logger  *zap.Logger
}

func (cl *client) SendMessage(channelId, message string) error {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	_, ok := cl.session.ChannelMessageSend(channelId, message)
	return ok
}

func (cl *client) ChannelMessageSendEmbed(channelId string, embed *discordgo.MessageEmbed) error {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	_, ok := cl.session.ChannelMessageSendEmbed(channelId, embed)
	return ok
}

func (cl *client) UpdateMember(guildID, userID string, roles []string) error {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildMemberEdit(guildID, userID, roles)
}

func (cl *client) GetAllMembers(guildID, after string, limit int) ([]*discordgo.Member, error) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildMembers(guildID, after, limit)
}

func (cl *client) GetAllMembersAsSlice(guildID string) ([]*discordgo.Member, error) {
	sugar := cl.logger.Sugar()
	var numberPerPage = 1000
	var memberCount = 1
	var memberId = ""
	var output []*discordgo.Member

	for memberCount > 0 {
		members, err := cl.GetAllMembers(guildID, memberId, numberPerPage)
		if err != nil {
			msg := fmt.Sprintf("GetAllMembersAsSlice: GetAllMembers: %s", err.Error())
			sugar.Error(msg)
			return nil, err
		}

		for m := range members {
			user := members[m].User
			output = append(output, members[m])

			oldNum, _ := strconv.Atoi(user.ID)
			newNum, _ := strconv.Atoi(memberId)

			if oldNum > newNum {
				memberId = user.ID
			}
		}

		memberCount = len(members)
	}

	return output, nil
}

func (cl *client) GetAllRoles(guildID string) ([]*discordgo.Role, error) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildRoles(guildID)
}

func (cl *client) RemoveMemberRole(guildID, userID, role string) error {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildMemberRoleRemove(guildID, userID, role)
}

func (cl *client) GetUser(userID string) (*discordgo.User, error) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.User(userID)
}

func (cl *client) CreateRole(guildId string) (*discordgo.Role, error) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildRoleCreate(guildId)
}

func (cl *client) DeleteRole(guildId, roleId string) error {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildRoleDelete(guildId, roleId)
}

func (cl *client) EditRole(guildId, roleId, name string, color, perm int, hoist, mention bool) (*discordgo.Role, error) {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	return cl.session.GuildRoleEdit(guildId, roleId, name, color, hoist, perm, mention)
}

func NewClient(config *config.Configuration, logger *zap.Logger) (DiscordClient, error) {
	session, err := discordgo.New("Bot " + config.Bot.BotToken)

	var newClient client
	if err != nil {
		return nil, err
	}

	newClient = client{session: session, logger: logger}

	// Start up name resolution cache updater
	//go newClient.nameResolutionCacheUpdater(config)

	return &newClient, nil
}

//
// This is unused but I'm leaving it here for now.
//

//func (cl *client) nameResolutionCacheUpdater(config *config.Configuration) {
//	sugar := cl.logger.Sugar()
//	ticker := time.NewTicker(time.Minute * 1)
//
//	sugar.Info("Starting nameResolutionCacheUpdater")
//
//	for range ticker.C {
//		t := time.Now()
//		cl.cacheUpdaterPoll(config)
//		sugar.Infof("Poller finished [%s]", time.Since(t))
//	}
//}

//func (cl *client) cacheUpdaterPoll(config *config.Configuration) {
//	sugar := cl.logger.Sugar()
//	var numberPerPage = 1000
//	var memberCount = 1
//	var memberId = ""
//
//	addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
//	redisClient := redis.Init(addr, config.Redis.Password, 1, config.LookupService("srv", "perms"))
//
//	for memberCount > 0 {
//		members, err := cl.GetAllMembers(config.Bot.DiscordServerId, memberId, numberPerPage)
//		if err != nil {
//			msg := fmt.Sprintf("nameResolutionCacheUpdater: GetAllMembers: %s", err.Error())
//			sugar.Error(msg)
//			continue
//		}
//
//		sugar.Infof("members: %v", reflect.TypeOf(members))
//
//		for m := range members {
//			user := members[m].User
//
//			_, err = redisClient.Client.HMSet(user.ID, structs.Map(user)).Result()
//
//			if err != nil {
//				sugar.Error(err)
//			}
//
//			oldNum, _ := strconv.Atoi(user.ID)
//			newNum, _ := strconv.Atoi(memberId)
//
//			if oldNum > newNum {
//				memberId = user.ID
//			}
//		}
//
//		memberCount = len(members)
//	}
//}