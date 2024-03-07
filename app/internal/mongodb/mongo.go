package mongodb

import (
	"context"
	"fmt"
	"log"

	"backend/app/internal/structs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config structs.Mongo) *mongo.Client {
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
