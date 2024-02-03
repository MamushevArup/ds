package main

import (
	"context"
	"github.com/MamushevArup/discord-bot/internal/config"
	"github.com/MamushevArup/discord-bot/internal/handler"
	"github.com/MamushevArup/discord-bot/internal/repo"
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"github.com/MamushevArup/discord-bot/pkg/mongodb"
	"net/http"
)

func main() {
	// Init logger
	lg := logger.NewLogger()
	// Init config use clean env library
	cfg, err := config.NewConfig()
	if err != nil {
		lg.Fatalf("error with reading config %e", err)
	}
	// Init the package universal mongo client creator
	mg, err := mongodb.NewClient(context.Background(), cfg.Mongo.URL, cfg.Mongo.Database)
	if err != nil {
		lg.Fatalf("unable to connect to the storage %v", err)
	}

	// Init storage | repo layer
	storage := repo.NewRepo(lg, mg)
	srv := service.NewService(storage)
	hdl := handler.NewBot(srv)

	if err = http.ListenAndServe(":"+cfg.HTTP.Port, hdl.InitRoutes()); err != nil {
		return
	}
}
