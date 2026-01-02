package auth

import (
	"errors"
	"expense-tracker/internal/models"
	service2 "expense-tracker/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	userService service2.UserService
}

func NewAuthHandler(userService service2.UserService) *authHandler {
	return &authHandler{userService: userService}
}

func RegisterRoutes(group *gin.RouterGroup, loginHandler *authHandler) {
	group.POST("/login", loginHandler.Login)
	group.POST("/signup", loginHandler.Signup)
}

// @Summary     Signup a new user
// @Description Register a new user and generate token for it
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body models.SignupRequest true "Signup info"
// @Success     201 {object} models.SignupResponse
// @Failure     400 {object} models.ErrorResponse
// @Failure     409 {object} models.ErrorResponse
// @Router      /api/v1/auth/signup [post]
func (h *authHandler) Signup(c *gin.Context) {
	var req models.SignupRequest
	err := c.ShouldBindBodyWithJSON(&req)

	if err != nil {
		log.Print("can't parse user: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := h.userService.Signup(req.Username, req.Password, c)

	if err != nil {
		if errors.Is(err, service2.ErrUserNameAlreadyExists) {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		} else {

			log.Print("can't insert user in repository: ", err.Error())
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "An unexpected error has occurred. please try again.",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, models.SignupResponse{
		Message:  "you registered successfully!",
		Username: user.Username,
		Token:    user.Token,
	})
}

// AutoSave     godoc
// @Summary     login a user
// @Description login a user and generate token for it
// @Tags        Auth
// @Accept      json
// @Param   	request body models.LoginRequest true "Signup info"
// @Produce     json
// @Success     200 {object} models.LoginResponse
// @Failure     400 {object} models.ErrorResponse
// @Failure     401 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Router      /api/v1/auth/login [POST]
func (h *authHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindBodyWithJSON(&req)

	if err != nil {
		log.Print("error to parsing user: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Login(req.Username, req.Password, c)

	if err != nil {
		if errors.Is(err, service2.ErrUserNotFound) {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: "you are not registered yet.",
			})
			return
		} else if errors.Is(err, service2.ErrPasswordNotMatch) {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: "username and password do not match.",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "An unexpected error has occurred. please try again.",
			})
			return
		}
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Username: user.Username,
		Token:    user.Token,
	})
}
