package models

import()

type User struct{
	Username string		`bson:"username"`
	Password string		`bson:"password"`
	ID int				`bson:"_id"`
}

