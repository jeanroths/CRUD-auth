package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"github.com/jeanroths/CRUD-auth/src/configuration/logger"
	"os"
)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_USER_DB= "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context,) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Info("Connected to MongoDB")

	return client.Database(mongodb_database), nil
}