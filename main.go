package main

import (
	"fmt"
	"github.com/chremoas/discord-gateway/discord"
	"github.com/chremoas/discord-gateway/handler"
	proto "github.com/chremoas/discord-gateway/proto"
	"github.com/chremoas/services-common/config"
	"github.com/micro/go-micro"
)

var Version = "1.0.0"
var service micro.Service
var name = "discord"

func main() {
	service = config.NewService(Version, "gateway", name, initialize)

	err := service.Run()
	if err != nil {
		fmt.Printf("Service stopped (%s)\n", err.Error())
	}
}

func initialize(config *config.Configuration) error {
	client, err := discord.NewClient(config.Bot.BotToken)
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
