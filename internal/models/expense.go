package models

type Expense struct {
	UserID      string `bson:"user_id"`
	ID          int    `bson:"id" binding:"required"`
	Title       string `bson:"title" binding:"required"`
	Category    string `bson:"category" binding:"required"`
	Date        string `bson:"date" binding:"required"`
	Status      string `bson:"status" binding:"required"`
	Description string `bson:"description"`
}
