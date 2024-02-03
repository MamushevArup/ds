package service

import (
	"context"
	"github.com/MamushevArup/discord-bot/internal/repo"
	"strconv"
)

type Gus struct {
	gus repo.Guess
}

func NewGuess(gus repo.Guess) *Gus {
	return &Gus{gus: gus}
}

func (g *Gus) MatchNumbers(ctx context.Context, id, number string) (string, error) {
	try, err := g.gus.Try(ctx, id)
	if err != nil {
		return "", err
	}
	v, err := strconv.Atoi(number)
	if err != nil {
		return "", err
	}
	if try > v {
		return "Too low try again!!!!", nil
	} else if try < v {
		return "Too high try again!!!", nil
	}
	return "Exact number found!!!", nil
}
