package myValidator

import (
	"expense-tracker/database"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidUserName(fl validator.FieldLevel)bool{
	userName, ok := fl.Field().Interface().(string)

	if !ok {return false}
	return database.CheckUserNameExist(userName)
}

func Future(fl validator.FieldLevel) bool {
	rawValue := fl.Field().Interface()
	var date time.Time

	switch v := rawValue.(type) {
    case time.Time:
        date = v 
    case string:
        parsedDate, err := time.Parse("2006-01-02", v)
        if err != nil {
            log.Printf("Failed to parse date: %v", err)
            return false
        }
        date = parsedDate
    default:
        log.Printf("Unsupported type: %T", rawValue)
        return false
    }
	return date.After(time.Now())
}