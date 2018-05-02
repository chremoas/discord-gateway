FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>

ADD discord-gateway-linux-amd64 discord-gateway
ADD ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
VOLUME /etc/chremoas

ENTRYPOINT ["/discord-gateway", "--configuration_file", "/etc/chremoas/chremoas.yaml"]
