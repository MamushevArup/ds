package service

import (
	"context"
	"github.com/MamushevArup/discord-bot/internal/repo"
)

type Hello interface {
	CreateUser(ctx context.Context, id string) error
}

type Gamer interface {
	GenerateNumber(ctx context.Context, up, low int) (int, error)
}

type Guess interface {
}

type Service struct {
	Gamer Gamer
}

func NewService(repos *repo.Repo) *Service {
	return &Service{
		Gamer: NewGame(repos.Gamer),
	}
}
