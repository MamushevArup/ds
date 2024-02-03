package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/MamushevArup/discord-bot/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var userCollection = "users"

type User struct {
	lg *logger.Logger
	db *mongo.Collection
}

func NewUser(lg *logger.Logger, db *mongo.Database) *User {
	return &User{lg: lg, db: db.Collection(userCollection)}
}

func (u *User) AddUser(ctx context.Context, id string) error {

	// Construct the document
	document := map[string]interface{}{
		"_id":       id,
		"createdAt": time.Now(),
	}

	// Insert the document into the database
	_, err := u.db.InsertOne(ctx, document)
	if err != nil {
		u.lg.Errorf("error inserting user: %v", err)
		return err
	}

	return nil
}
func (u *User) UserExists(ctx context.Context, userID string) (bool, error) {
	// Create a filter based on the provided userID
	filter := bson.M{"_id": userID} // Assuming "_id" is the field for the user ID, update it as needed

	// Perform a find operation with the filter
	result := u.db.FindOne(ctx, filter)

	// Check if the user exists
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return false, nil // User does not exist
	} else if result.Err() != nil {
		return false, fmt.Errorf("error checking user existence: %v", result.Err())
	}

	return true, nil // User exists
}
