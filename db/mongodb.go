package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func NewDB() *mongo.Client {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")

	DB, err := mongo.Connect(ctx, clientOption)

	if err != nil {
		log.Printf("failed to connect to mongodb: %v", err)
		return nil
	}

	err = DB.Ping(ctx, nil)

	if err != nil {
		log.Printf("failed to ping mongodb: %v", err)
		return nil
	}

	return DB
}
