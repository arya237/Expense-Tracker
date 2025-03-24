package models

import()

type Expense struct{
	
	ID int					`bson:"_id"`
	UserID any				`bson:"user_id"`
	Title string			`bson:"title"`
	Category string			`bson:"category"`
	Date string				`bson:"date"`
	ExpirationDate string	`bosn:"expire"`
	Description string		`bson:"description"`
}