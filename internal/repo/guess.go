package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Gus struct {
	lg *logger.Logger
	db *mongo.Collection
}

func NewGus(lg *logger.Logger, db *mongo.Database) *Gus {
	return &Gus{lg: lg, db: db.Collection(userCollection)}
}

func (g *Gus) Try(ctx context.Context, id string) (int, error) {
	filter := bson.M{"_id": id}

	// Create a variable to store the result
	var result struct {
		Number int `bson:"number"`
	}

	// Perform a find operation with the filter
	err := g.db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, fmt.Errorf("document with ID %s not found", id)
		}
		return 0, fmt.Errorf("error fetching number: %v", err)
	}

	return result.Number, nil
}
