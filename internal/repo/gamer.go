package repo

import (
	"context"
	"fmt"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GamerUp struct {
	lg *logger.Logger
	db *mongo.Collection
}

func NewGamer(lg *logger.Logger, db *mongo.Database) *GamerUp {
	return &GamerUp{lg: lg, db: db.Collection(userCollection)}
}

func (g *GamerUp) InsertRandom(ctx context.Context, id string, number int) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"number": number}}

	// Perform the update
	result, err := g.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	// Check if the document was updated
	if result.ModifiedCount == 0 {
		return fmt.Errorf("document with ID %s not found", id)
	}

	return nil
}
