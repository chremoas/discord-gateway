package main

import (
	"fmt"

	"github.com/chremoas/services-common/config"
	chremoasPrometheus "github.com/chremoas/services-common/prometheus"
	"github.com/micro/go-micro"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"

	"github.com/chremoas/discord-gateway/discord"
	"github.com/chremoas/discord-gateway/handler"
	proto "github.com/chremoas/discord-gateway/proto"
)

var (
	Version  = "SET ME YOU KNOB"
	service  micro.Service
	name     = "discord"
	logger   *zap.Logger
	apiCalls = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_gateway_api_calls",
		Help: "The total number of APIs called made to discord",
	})
)

func main() {
	var err error

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Initialized logger")

	go chremoasPrometheus.PrometheusExporter(logger)

	service = config.NewService(Version, "gateway", name, initialize)

	err = service.Run()
	if err != nil {
		fmt.Printf("Service stopped (%s)\n", err.Error())
	}
}

func initialize(config *config.Configuration) error {
	client, err := discord.New(config, logger)
	if err != nil {
		return err
	}
	roleMap := discord.NewRoleMap(config.Bot.DiscordServerId, client)

	theHandler, err := handler.New(config.Bot.DiscordServerId, client, roleMap, logger, apiCalls)
	if err != nil {
		return err
	}

	proto.RegisterDiscordGatewayHandler(service.Server(), theHandler)

	return nil
}
