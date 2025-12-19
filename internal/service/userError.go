package service

import "errors"

var (
	ErrorGetUser             = errors.New("can't get user by username")
	ErrUserNameAlreadyExists = errors.New("username already exists")
	ErrUserRegisterd         = errors.New("user is already registered")
	ErrUserNotFound          = errors.New("user not found")
	ErrPasswordNotMatch      = errors.New("username or password does not match")
	ErrNoExpense             = errors.New("user has no expense")
	ErrTokenGenerate         = errors.New("can't generate token")
)
