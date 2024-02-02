package repo

import (
	"context"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertCommand interface {
	Insert(ctx context.Context, command, description string) error
}

type Gamer interface {
	InsertRandom(ctx context.Context, id string, number int) error
}

type Repo struct {
	Insert InsertCommand
	Gamer  Gamer
}

func NewRepo(lg *logger.Logger, db *mongo.Database) *Repo {
	return &Repo{
		Insert: NewCommand(lg, db),
	}
}
