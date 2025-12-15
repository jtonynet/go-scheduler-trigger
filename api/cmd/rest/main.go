package main

import (
	"context"
	"log"

	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/http/router"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	appREST, err := bootstrap.NewREST(cfg)
	if err != nil {
		log.Fatal("cannot initiate app: ", err)
	}

	routes, err := router.NewGin(cfg.API, *appREST)
	if err != nil {
		log.Fatal("cannot initiate routes: ", err)
	}
	routes.HandleRequests(ctx)
}
