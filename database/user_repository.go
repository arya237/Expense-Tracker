package database

import (
	"errors"
	"expense-tracker/models"
	"expense-tracker/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func AddUserToDatabase(u models.User) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	collection := DB.Database("expense_tracker").Collection("users")

	indexModel := mongo.IndexModel{
		Keys: bson.M{"username":1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil{
		return err
	}

	_, err = collection.InsertOne(ctx, u)

	if err != nil{
		return err
	}

	return nil
}

func GetUserFromDatabase(u models.User) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	collection := DB.Database("expense_tracker").Collection("users")

	filter := bson.M{"username": u.Username}

	var user models.User

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil{
		log.Print(err.Error())
		return nil, err
	} else if err := utils.CompareHashedPassword(user.Password, u.Password); err != nil{
		return nil, errors.New("username or password is incorrect")
	}

	return &user, nil
}

func CheckUserNameExist(userName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	collection := DB.Database("expense_tracker").Collection("users")

	filter := bson.M{"username": userName}

	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)

	return err == nil 
}

