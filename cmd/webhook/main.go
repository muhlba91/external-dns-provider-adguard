package main

import (
	"fmt"

	"github.com/muhlba91/external-dns-provider-adguard/cmd/webhook/init/configuration"
	"github.com/muhlba91/external-dns-provider-adguard/cmd/webhook/init/dnsprovider"
	"github.com/muhlba91/external-dns-provider-adguard/cmd/webhook/init/logging"
	"github.com/muhlba91/external-dns-provider-adguard/cmd/webhook/init/server"
	"github.com/muhlba91/external-dns-provider-adguard/pkg/webhook"
	log "github.com/sirupsen/logrus"
)

const banner = `
external-dns-provider-adguard
version: %s (%s)

`

var (
	Version = "local"
	Gitsha  = "?"
)

func main() {
	fmt.Printf(banner, Version, Gitsha)

	logging.Init()

	config := configuration.Init()
	provider, err := dnsprovider.Init(config)
	if err != nil {
		log.Fatalf("failed to initialize provider: %v", err)
	}

	srv := server.Init(config, webhook.New(provider))
	server.ShutdownGracefully(srv)
}
