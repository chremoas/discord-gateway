package discord_gateway

//go:generate protoc --go_out=plugins=micro:. gateway.proto
//go:generate mockgen -package=discord_gateway -source=gateway.pb.go -destination=gateway_mocks.go DiscordGatewayClient
