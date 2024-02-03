package service

import (
	"context"
	"errors"
	"github.com/MamushevArup/discord-bot/internal/repo"
)

type User struct {
	user repo.InsertUser
}

func NewUser(user repo.InsertUser) *User {
	return &User{user: user}
}

var (
	emptyID = errors.New("id is empty")
)

func (u *User) CreateUser(ctx context.Context, id string) error {
	if id == "" {
		return emptyID
	}
	exists, err := u.user.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return u.user.AddUser(ctx, id)
}
