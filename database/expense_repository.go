package database

import (
	"context"
	"expense-tracker/models"
	"time"
)

func AddExpenseToDatabase(e models.Expense)error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	collection := DB.Database("expense_tracker").Collection("expenses")

	_, err := collection.InsertOne(ctx, e)

	if err != nil{
		return err
	}

	return nil
}