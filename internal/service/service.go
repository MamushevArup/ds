package service

import (
	"context"
	"github.com/MamushevArup/discord-bot/internal/repo"
)

type Hello interface {
	CreateUser(ctx context.Context, id string) error
}

type Gamer interface {
	GenerateNumber(ctx context.Context, id string, up, low int) (int, error)
}

type Guess interface {
	MatchNumbers(ctx context.Context, id, number string) (string, error)
}

type Poll interface {
	CreatePoll(id, question string, options map[int]string) error
}

type Service struct {
	Game  Gamer
	Hello Hello
	Guess Guess
	Poll  Poll
}

func NewService(repos *repo.Repo) *Service {
	return &Service{
		Game:  NewGame(repos.Game),
		Hello: NewUser(repos.AddU),
		Guess: NewGuess(repos.Guess),
		Poll:  NewPoll(),
	}
}
