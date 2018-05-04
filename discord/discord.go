package discord

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

// This is a very thin wrapper around the discordgo api for testability purposes
// and to easily keep track of what endpoints are being consumed
type DiscordClient interface {
	SendMessage(channelId, message string) error
	UpdateMember(guildID, userID string, roles []string) error
	RemoveMemberRole(guildID, userID, role string) error
	GetAllMembers(guildID, after string, limit int) ([]*discordgo.Member, error)
	GetAllRoles(guildID string) ([]*discordgo.Role, error)
	GetUser(userID string) (*discordgo.User, error)
	CreateRole(guildId string) (*discordgo.Role, error)
	DeleteRole(guildId, roleId string) error
	EditRole(guildId, roleId, name string, color, perm int, hoist, mention bool) (*discordgo.Role, error)
}

type client struct {
	session *discordgo.Session
	mutex   sync.Mutex
}

func (cl *client) SendMessage(channelId, message string) error {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	_, ok := cl.session.ChannelMessageSend(channelId, message)
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

func NewClient(token string) (DiscordClient, error) {
	session, err := discordgo.New("Bot " + token)
	var newClient client
	if err != nil {
		return nil, err
	}
	newClient = client{session: session}
	return &newClient, nil
}
