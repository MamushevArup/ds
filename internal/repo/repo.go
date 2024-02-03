package repo

import (
	"context"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Gamer interface {
	InsertRandom(ctx context.Context, id string, number int) error
}
type InsertUser interface {
	AddUser(ctx context.Context, id string) error
	UserExists(ctx context.Context, userID string) (bool, error)
}

type Guess interface {
	Try(ctx context.Context, id string) (int, error)
}

// Repo is struct of interfaces for easy management
type Repo struct {
	Game  Gamer
	AddU  InsertUser
	Guess Guess
}

func NewRepo(lg *logger.Logger, db *mongo.Database) *Repo {
	return &Repo{
		AddU:  NewUser(lg, db),
		Game:  NewGamer(lg, db),
		Guess: NewGus(lg, db),
	}
}
