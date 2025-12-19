package service

import (
	"context"
	"errors"
	models2 "expense-tracker/internal/models"
	repository2 "expense-tracker/internal/repository"
	"expense-tracker/utils"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Signup(username, password string, ctx context.Context) (*models2.SignupResponse, error)
	Login(username, password string, ctx context.Context) (*models2.LoginResponse, error)
	Delete(username string, ctx context.Context) error
	GetAllExpense(username string, ctx context.Context) ([]*models2.Expense, error)
	GetSpecificExpense(username, startDate, endDate string, ctx context.Context) ([]*models2.Expense, error)
	AddExpense(expense *models2.Expense, ctx context.Context) error
	DeleteExpense(username string, expenseID int, ctx context.Context) error
}

type userService struct {
	userRepo    repository2.UserRepository
	expenseRepo repository2.ExpenseRepository
}

func NewUserService(userRepo repository2.UserRepository, expenseRepo repository2.ExpenseRepository) UserService {
	return &userService{
		userRepo:    userRepo,
		expenseRepo: expenseRepo,
	}
}

func (u *userService) Signup(username, password string, ctx context.Context) (*models2.SignupResponse, error) {
	exist := u.userRepo.CheckUserNameExist(username, ctx)
	if exist {
		return nil, ErrUserNameAlreadyExists
	}

	pass, err := utils.HashPassword(password)

	if err != nil {
		log.Printf("service.Signup: can't hashing password %s with error %v: ", password, err)
		return nil, fmt.Errorf("can't hashing password details %w", err)
	}

	claims := utils.CreateJwtClaims(username)
	token, err := utils.CreateToken(claims)

	if err != nil {
		log.Print("service.Signup: can't create jwt claims: ", err.Error())
		return nil, ErrTokenGenerate
	}

	user := &models2.User{
		Username: username,
		Password: pass,
	}

	err = u.userRepo.Save(user, ctx)
	if err != nil {
		log.Printf("service.Signup: fail to save user with username %s with error: %v", username, err)
		return nil, fmt.Errorf("failed to save user details %w", err)
	}

	return &models2.SignupResponse{Username: username, Token: token}, nil
}

func (u *userService) Login(username, password string, ctx context.Context) (*models2.LoginResponse, error) {
	user, err := u.userRepo.GetByUsername(username, ctx)
	if err != nil {
		log.Printf("repository.login: fail to get user by username %s with error: %v", username, err)
		if errors.Is(err, repository2.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by username details %w", err)
	}

	if err := utils.CompareHashedPassword(user.Password, password); err != nil {
		log.Printf("service.login: password is incorrect %s with error %v", password, err)
		return nil, ErrPasswordNotMatch
	}

	claims := utils.CreateJwtClaims(username)
	token, err := utils.CreateToken(claims)
	if err != nil {
		log.Print("service.Signup: can't create jwt claims: ", err.Error())
		return nil, ErrTokenGenerate
	}

	return &models2.LoginResponse{Token: token, Username: username}, nil
}

func (u *userService) Delete(username string, ctx context.Context) error {
	user, err := u.userRepo.GetByUsername(username, ctx)
	if err != nil {
		log.Printf("repository.delete: fail to get user by username %s with error: %v", username, err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to get user details %w", err)
	}

	err = u.userRepo.Delete(user.Username, ctx)
	if err != nil {
		log.Printf("repository.delete: fail to delete user with username %s with error: %v", username, err)
		return fmt.Errorf("failed to delete user details: %w", err)
	}

	return nil
}

func (u *userService) GetAllExpense(username string, ctx context.Context) ([]*models2.Expense, error) {

	_, err := u.userRepo.GetByUsername(username, ctx)
	if err != nil {
		log.Printf("service.GetAllExpense: failed to get from repository: user by username %s with error: %v", username, err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to retrive user with details: %w", err)
	}

	expenses, err := u.expenseRepo.GetAll(username, ctx)
	if err != nil {
		log.Printf("service.GetAllExpense: failed to get user %s expenses with details: %s", username, err)
		if errors.Is(err, repository2.ErrNoExpense) {
			return nil, ErrNoExpense
		}

		return nil, fmt.Errorf("failed to get user expenses : %w", err)
	}

	return expenses, nil
}

func (u *userService) GetSpecificExpense(username, startDate, endDate string, ctx context.Context) ([]*models2.Expense, error) {
	_, err := u.userRepo.GetByUsername(username, ctx)
	if err != nil {
		log.Printf("service.GetAllExpense: failed to get from repository: user by username %s with error: %v", username, err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to retrive user with details: %w", err)
	}

	expenses, err := u.expenseRepo.GetByTime(startDate, endDate, username, ctx)
	if err != nil {
		log.Printf("service.GetAllExpense: failed to get user %s expenses with details: %s", username, err)
		if errors.Is(err, repository2.ErrNoExpense) {
			return nil, ErrNoExpense
		}

		return nil, fmt.Errorf("failed to get user expenses : %w", err)
	}

	return expenses, nil
}

func (u *userService) AddExpense(expense *models2.Expense, ctx context.Context) error {
	_, err := u.userRepo.GetByUsername(expense.UserID, ctx)
	if err != nil {
		log.Printf("service.GetAllExpense: failed to get from repository: user by username %s with error: %v", expense.UserID, err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to retrive user with details: %w", err)
	}

	err = u.expenseRepo.Save(expense, ctx)
	if err != nil {
		log.Printf("can't save expense with error: %v", err)
		return fmt.Errorf("can't save expense with error: %w", err)
	}

	return nil
}

func (u *userService) DeleteExpense(username string, expenseID int, ctx context.Context) error {
	_, err := u.userRepo.GetByUsername(username, ctx)
	if err != nil {
		log.Printf("service.GetAllExpense: failed to get from repository: user by username %s with error: %v", username, err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to retrive user with details: %w", err)
	}

	//TODO check expense with specifiate id has exist

	err = u.expenseRepo.Delete(username, expenseID, ctx)
	if err != nil {
		log.Printf("can't delete expense with error: %v", err)
		return fmt.Errorf("can't delete expense with error: %w", err)
	}

	return nil
}
