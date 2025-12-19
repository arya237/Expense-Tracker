package repository

import (
	"errors"
	"expense-tracker/internal/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Save(u *models.User, ctx context.Context) error
	GetByUsername(username string, ctx context.Context) (*models.User, error)
	CheckUserNameExist(u string, ctx context.Context) bool
	Delete(username string, ctx context.Context) error
}

type userRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Save(u *models.User, ctx context.Context) error {

	collection := r.db.Database("expense_tracker").Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		return err
	}

	_, err = collection.InsertOne(ctx, u)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Print("repository.save: User already exists with username", u.Username)
			return ErrUserAlreadyExists
		}
		log.Printf("repository.save: fail to save user with username %s with error: %v", u.Username, err)
		return fmt.Errorf("failed to save user with id %s: %w", u.Username, err)
	}

	return nil
}

func (r *userRepository) GetByUsername(username string, ctx context.Context) (*models.User, error) {

	collection := r.db.Database("expense_tracker").Collection("users")

	filter := bson.M{"username": username}

	var user models.User

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Printf("repository.GetByUsername: User with username %s not found", username)
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
	}
	return &user, nil
}

func (r *userRepository) CheckUserNameExist(userName string, ctx context.Context) bool {

	collection := r.db.Database("expense_tracker").Collection("users")

	filter := bson.M{"username": userName}

	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)

	return err == nil
}

func (r *userRepository) Delete(username string, ctx context.Context) error {

	collection := r.db.Database("expense_tracker").Collection("users")

	filter := bson.M{"username": username}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)
