package repo

import (
	"context"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Command struct {
	mg *mongo.Database
	lg *logger.Logger
}

func NewCommand(lg *logger.Logger, mg *mongo.Database) *Command {
	return &Command{mg: mg, lg: lg}
}

func (c *Command) Insert(ctx context.Context, command, desc string) error {
	return nil
}
