package mongorm

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		cancel()
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to Database")
	return client, ctx, nil
}
