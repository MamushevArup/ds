package main

import (
	"context"
	"github.com/MamushevArup/discord-bot/internal/config"
	"github.com/MamushevArup/discord-bot/internal/handler"
	"github.com/MamushevArup/discord-bot/internal/repo"
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"github.com/MamushevArup/discord-bot/pkg/mongodb"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	// Init logger
	lg := logger.NewLogger()
	// Init environment reader
	if err := godotenv.Load(); err != nil {
		lg.Fatalf("cannot load env %v", err)
	}
	// Init config use clean env library
	cfg, err := config.NewConfig()
	if err != nil {
		lg.Fatalf("error with reading config %e", err)
	}

	mgUrl := os.Getenv("MONGO_URL")
	// Init the package universal mongo client creator
	mg, err := mongodb.NewClient(context.Background(), mgUrl, cfg.Mongo.Database)
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
