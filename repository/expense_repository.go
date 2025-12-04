package repository

import (
	"context"
	"expense-tracker/db"
	"expense-tracker/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpenseRepository interface {
	Save(expense *models.Expense, ctx context.Context) error
	GetAll(username string, ctx context.Context) ([]*models.Expense, error)
	GetByTime(startDate, endDate, username string, ctx context.Context) ([]*models.Expense, error)
	Delete(ID string, ctx context.Context) error
	Update(id string, expense *models.Expense, ctx context.Context) error
}

type expenseRepository struct {
	db *mongo.Client
}

func NewExpenseRepository(db *mongo.Client) ExpenseRepository {
	return &expenseRepository{
		db: db,
	}
}

func (r *expenseRepository) Save(e *models.Expense, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Database("expense_tracker").Collection("expenses")

	_, err := collection.InsertOne(ctx, e)

	if err != nil {
		return err
	}

	return nil
}

func (r *expenseRepository) GetAll(username string, ctx context.Context) ([]*models.Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Database("expense_tracker").Collection("expenses")

	filter := bson.M{
		"user_id": username,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var list []*models.Expense

	for cursor.Next(ctx) {
		var expense *models.Expense
		if err := cursor.Decode(&expense); err != nil {
			return nil, err
		}

		list = append(list, expense)
	}

	return list, nil
}

func (r *expenseRepository) GetByTime(startDate, endDate, username string, ctx context.Context) ([]*models.Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Database("expense_tracker").Collection("expenses")

	filter := bson.M{
		"user_id": username,
		"date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	var list []*models.Expense

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {

		var expense *models.Expense
		if err := cursor.Decode(&expense); err != nil {
			return nil, err
		}

		list = append(list, expense)
	}

	return list, nil
}

func (r *expenseRepository) Update(ID string, new *models.Expense, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Database("expense_tracker").Collection("expenses")

	objID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"title": new.Title, "description": new.Description}}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (r *expenseRepository) Delete(ID string, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Database("expense_tracker").Collection("expenses")

	objID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	log.Print(result)

	return nil
}
