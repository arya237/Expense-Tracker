package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){

	hashedpassByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil{
		return "", err
	}

	return string(hashedpassByte), nil
}

func CompareHashedPassword(hashedPass string, password string) error {  
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}