package models

import()

type Expense struct{
	
	UserID any				`bson:"user_id"`
	Title string			`bson:"title"`
	Category string			`bson:"category"`
	Date string				`bson:"date"`
	Status string 			`bson:"status"`
	DeadLine string			`bosn:"deadline"`
	Description string		`bson:"description"`
}