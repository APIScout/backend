package mongodb

import (
	"context"
	"fmt"
	"log"

	"backend/app/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - used to connect to the mongodb database. It will return a mongodb client that can be used to perform
// queries on the database.
func Connect(config models.MongoConfig) *mongo.Client {
	ctx := context.TODO()
	uri := fmt.Sprintf(
		"%s://%s:%s@%s:%d/?directConnection=true&authSource=apis",
		config.Protocol, config.User, config.Password, config.Host, config.Port,
	)
	opts := options.Client().ApplyURI(uri)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}
	log.Print("Connected to MongoDB")

	return client
}
