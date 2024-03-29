package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/MamushevArup/discord-bot/internal/repo"
	"math/big"
	r "math/rand"
)

type Game struct {
	game repo.Gamer
}

func NewGame(game repo.Gamer) *Game {
	return &Game{
		game: game,
	}
}

func (g *Game) GenerateNumber(ctx context.Context, id string, up, low int) (int, error) {
	if id == "" {
		return 0, errors.New("id is empty")
	}
	// randomize using built-in packages
	seed, err := rand.Int(rand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		return 0, fmt.Errorf("error generating random seed: %v", err)
	}

	// Seed the random number generator
	r.Seed(seed.Int64())

	// Check if boundaries are valid
	if low >= up {
		return 0, fmt.Errorf("invalid boundaries: %d should be less than %d", low, up)
	}

	// Generate and return a random number within the specified range [low, up)
	random := r.Intn(up-low) + low
	// insert number to the storage
	err = g.game.InsertRandom(ctx, id, random)
	if err != nil {
		return 0, err
	}
	return random, nil
}
