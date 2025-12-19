package repository

import (
	"context"
	"errors"
	"expense-tracker/internal/db"
	"expense-tracker/internal/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpenseRepository interface {
	Save(expense *models.Expense, ctx context.Context) error
	GetAll(username string, ctx context.Context) ([]*models.Expense, error)
	GetByTime(startDate, endDate, username string, ctx context.Context) ([]*models.Expense, error)
	Delete(userID string, expenseID int, ctx context.Context) error
	Update(userID string, expense *models.Expense, ctx context.Context) error
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
	collection := r.db.Database("expense_tracker").Collection("expenses")

	_, err := collection.InsertOne(ctx, e)

	if err != nil {
		log.Printf("expense_repository.Save: Error inserting expense: %v", err)
		return ErrSaveExpense
	}

	return nil
}

func (r *expenseRepository) GetAll(username string, ctx context.Context) ([]*models.Expense, error) {
	collection := r.db.Database("expense_tracker").Collection("expenses")

	filter := bson.M{
		"user_id": username,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("expense_repository.GetAll: Error finding expenses of username %s error, %v ", username, err)
		return nil, fmt.Errorf("failed to retrive expense of username %s error, %v", username, err)
	}

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
		log.Printf("expense_repository: Error finding expenses of username %s error, %v ", username, err)
		return nil, fmt.Errorf("can't find expenses of username %s in database: %w", username, err)
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

func (r *expenseRepository) Update(username string, new *models.Expense, ctx context.Context) error {

	collection := db.DB.Database("expense_tracker").Collection("expenses")

	filter := bson.M{"user_id": username}
	update := bson.M{"$set": bson.M{"ID": new.ID, "title": new.Title, "description": new.Description}}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Printf("expense_repository.Update: Error updating expense %d of username %s with error %v:", new.ID, username, err)
		return fmt.Errorf("failed to updating expense %d of username %s details %w", new.ID, username, err)
	}

	return nil
}

func (r *expenseRepository) Delete(username string, ID int, ctx context.Context) error {
	collection := db.DB.Database("expense_tracker").Collection("expenses")

	filter := bson.M{"ID": ID, "user_id": username}

	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Printf("expense_repository: Error deleting expense %d of userID %s with error %v:", ID, username, err)
		return fmt.Errorf("failed to deleting expense %d of userID %s details %w", ID, username, err)
	}

	return nil
}

var (
	ErrNoExpense   = errors.New("there is no expense in database")
	ErrSaveExpense = errors.New("can't save expense")
)
