package main

import (
	"fmt"
	"github.com/chremoas/discord-gateway/discord"
	"github.com/chremoas/discord-gateway/handler"
	proto "github.com/chremoas/discord-gateway/proto"
	"github.com/chremoas/services-common/config"
	"github.com/micro/go-micro"
	"go.uber.org/zap"
)

var Version = "SET ME YOU KNOB"
var service micro.Service
var name = "discord"
var logger *zap.Logger

func main() {
	service = config.NewService(Version, "gateway", name, initialize)
	var err error

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Initialized logger")

	err = service.Run()
	if err != nil {
		fmt.Printf("Service stopped (%s)\n", err.Error())
	}
}

func initialize(config *config.Configuration) error {
	client, err := discord.NewClient(config, logger)
	if err != nil {
		return err
	}
	roleMap := discord.NewRoleMap(config.Bot.DiscordServerId, client)

	theHandler, err := handler.NewDiscordGatewayHandler(config.Bot.DiscordServerId, client, roleMap)
	if err != nil {
		return err
	}

	proto.RegisterDiscordGatewayHandler(service.Server(), theHandler)

	return nil
}
