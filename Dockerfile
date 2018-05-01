FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>

ADD discord-gateway-linux-amd64 discord-gateway
VOLUME /etc/chremoas

ENTRYPOINT ["/discord-gateway", "--configuration_file", "/etc/chremoas/chremoas.yaml"]
