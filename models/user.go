package models

import()

type User struct{
	Username string		`bson:"username" binding:"required,min=3"`
	Password string		`bson:"password" binding:"required"`
}

