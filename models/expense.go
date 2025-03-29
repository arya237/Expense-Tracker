package models

import()

type Expense struct{
	
	UserID any				`bson:"user_id" binding:"required,validUserName"`
	Title string			`bson:"title" binding:"required"`
	Category string			`bson:"category" binding:"required"`
	Date string				`bson:"date" binding:"required"`
	Status string 			`bson:"status" binding:"required"`
	DeadLine string			`bosn:"deadline" binding:"future"`
	Description string		`bson:"description"`
}