package discord

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

type RoleMap interface {
	UpdateRoles() error
	GetRoles() map[string]*discordgo.Role
	GetRoleId(roleName string) string
	GetRoleName(roleId string) string
	GetRoleByName(roleName string) *discordgo.Role
	GetRoleById(roleId string) *discordgo.Role
}

type roleMapImpl struct {
	guildID     string
	client      DiscordClient
	rolesByName map[string]*discordgo.Role
	rolesById   map[string]*discordgo.Role
	mutex       sync.Mutex
}

func (rm *roleMapImpl) UpdateRoles() error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	roles, err := rm.client.GetAllRoles(rm.guildID)

	if err != nil {
		return err
	}

	rm.rolesByName = make(map[string]*discordgo.Role)
	rm.rolesById = make(map[string]*discordgo.Role)

	for _, role := range roles {
		rm.rolesByName[role.Name] = role
		rm.rolesById[role.ID] = role
	}

	return nil
}

func (rm *roleMapImpl) GetRoles() map[string]*discordgo.Role {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()
	return rm.rolesByName
}

func (rm *roleMapImpl) GetRoleId(roleName string) string {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	role := rm.rolesByName[roleName]

	if role == nil {
		return ""
	}

	return role.ID
}

func (rm *roleMapImpl) GetRoleName(roleId string) string {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	role := rm.rolesById[roleId]

	if role == nil {
		return ""
	}

	return role.Name
}

func (rm *roleMapImpl) GetRoleByName(roleName string) *discordgo.Role {
	return rm.rolesByName[roleName]
}

func (rm *roleMapImpl) GetRoleById(roleId string) *discordgo.Role {
	return rm.rolesById[roleId]
}

func NewRoleMap(guildID string, client DiscordClient) RoleMap {
	return &roleMapImpl{guildID: guildID, client: client}
}
