package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func init(){

	var err error
	DB, err = ConnectToMongodb("mongodb://localhost:27017")  

	if err != nil{
		log.Fatal("failed to connect to mongodb: ", err.Error())
	}
}

func ConnectToMongodb(uri string) (*mongo.Client, error){
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOption)

	if err != nil{
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil{
		return nil, err
	}

	return client, nil
}