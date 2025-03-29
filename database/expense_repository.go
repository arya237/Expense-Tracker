package database

import (
	"context"
	"errors"
	"expense-tracker/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func ListExpense(filterTime string, username string) ([]models.Expense, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	collection := DB.Database("expense_tracker").Collection("expenses")

	Time := time.Now()
	var startTime string

	switch filterTime{
		case "lastweek":{
			startTime = Time.AddDate(0, 0, -7).Format("2006-01-02")
		}

		case "lastmonth":{
			startTime = Time.AddDate(0, -1, 0).Format("2006-01-02")
		}

		case "last3month":{
			startTime = Time.AddDate(0, -3, 0).Format("2006-01-02")
		}

		default:{return nil, errors.New("this format is incorrect")}
	}

	filter := bson.M{
		"date": bson.M{"$gt":startTime},
		"user_id": username,
	}

	cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
	
	var list []models.Expense

	for cursor.Next(ctx){
		var expense models.Expense
		if err := cursor.Decode(&expense); err != nil{
			return nil, err
		}

		list = append(list, expense)
	}

	return list, nil
}

func ListExpenseWithDeadLine(date string, username string) ([]models.Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	collection := DB.Database("expense_tracker").Collection("expenses")

	filter := bson.M{
		"user_id": username,
		"deadline": date,
	}

	var list []models.Expense

	cursor, err := collection.Find(ctx, filter)

	if err != nil{
		return nil, err
	}

	for cursor.Next(ctx) {
		
		var expense models.Expense
		if err := cursor.Decode(&expense); err != nil{
			return nil, err
		}

		list = append(list, expense)
	}

	return list, nil
}

func UpdateStatus(ID string, status string)error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	collection := DB.Database("expense_tracker").Collection("expenses")

	objID, err := primitive.ObjectIDFromHex(ID)

	if err != nil{
		return err 
	}

	filter := bson.M{"_id" : objID}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil{
		return err
	}

	return nil
}

func DeleteExpenseFromDatabase(ID string)error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	collection := DB.Database("expense_tracker").Collection("expenses")

	objID, err := primitive.ObjectIDFromHex(ID)

	if err != nil{
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil{
		return err 
	}

	log.Print(result)

	return nil
}