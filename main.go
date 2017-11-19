package main

import (
	"fmt"
	"github.com/abaeve/discord-gateway/discord"
	"github.com/abaeve/discord-gateway/handler"
	proto "github.com/abaeve/discord-gateway/proto"
	"github.com/abaeve/services-common/config"
	"github.com/micro/go-micro"
)

var Version = "1.0.0"
var service micro.Service

func main() {
	service = config.NewService(Version, "discord-gateway", initialize)

	err := service.Run()
	if err != nil {
		fmt.Printf("Service stopped (%s)\n", err.Error())
	}
}

func initialize(config *config.Configuration) error {
	client, err := discord.NewClient(config.Chat.Discord.Token)
	if err != nil {
		return err
	}
	roleMap := discord.NewRoleMap(config.Bot.DiscordServerId, client)

	theHandler, err := handler.NewDiscordGatewayHandler(client, roleMap)
	if err != nil {
		return err
	}

	proto.RegisterDiscordGatewayHandler(service.Server(), theHandler)

	return nil
}
