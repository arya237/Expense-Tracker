package expense

import (
	"expense-tracker/internal/middleware"
	models2 "expense-tracker/internal/models"
	"expense-tracker/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type expenseHandler struct {
	userService service.UserService
}

func NewExpenseHandler(userService service.UserService) *expenseHandler {
	return &expenseHandler{userService}
}

func RegisterRoutes(group *gin.RouterGroup, expenseHandler *expenseHandler) {
	group.Use(middleware.Authorization)
	group.POST("/add", expenseHandler.AddExpense)
	group.DELETE("/delete/:id", expenseHandler.DeleteExpense)
	group.GET("/list", expenseHandler.ListExpense)
	group.GET("/list/date", expenseHandler.ListExpenseByDate)
}

func (h *expenseHandler) AddExpense(c *gin.Context) {

	var expense models2.Expense
	username, _ := c.Get("username")
	expense.UserID = username.(string)

	err := c.ShouldBindBodyWithJSON(&expense)

	if err != nil {
		log.Print("can't bind expense: ", err.Error())
		c.JSON(http.StatusBadRequest, models2.ErrorResponse{
			Error: "incorrect expense. please try again",
		})
		return
	}

	err = h.userService.AddExpense(&expense, c)

	if err != nil {
		log.Print("can't add expense error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models2.ErrorResponse{
			Error: "An unexpected error has occurred",
		})
		return
	}

	c.JSON(http.StatusCreated, models2.AddExpenseResponse{
		Message: "expense added successfully",
	})
}

func (h *expenseHandler) ListExpenseByDate(c *gin.Context) {

	date1 := c.Query("date1")
	date2 := c.Query("date2")
	username, _ := c.Get("username")
	usernameString := username.(string)

	list, err := h.userService.GetSpecificExpense(usernameString, date1, date2, c)

	if err != nil {
		log.Print("can't get list of expense: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *expenseHandler) ListExpense(c *gin.Context) {

	username, _ := c.Get("username")
	usernameString := username.(string)

	list, err := h.userService.GetAllExpense(usernameString, c)

	if err != nil {
		log.Print("can't get list of expenses: ", err.Error())
		c.JSON(http.StatusInternalServerError, models2.ErrorResponse{
			Error: "An unexpected error has occurred",
		})
		return
	}

	c.JSON(http.StatusOK, list)
}

//func (h *authHandler) UpdateExpenseStatus(c *gin.Context) {
//
//	expenseID := c.Query("id")
//	status := c.Query("status")
//
//	err := h.userService
//
//	if err != nil {
//		log.Print("can't update expense: ", err.Error())
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Your expense updated successfuly"})
//}

func (h *expenseHandler) DeleteExpense(c *gin.Context) {
	expenseID := c.Param("id")
	username, _ := c.Get("username")

	usernameString := username.(string)
	expenseIDint, _ := strconv.Atoi(expenseID)

	err := h.userService.DeleteExpense(usernameString, expenseIDint, c)

	if err != nil {
		log.Print("can't delete this expense: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your expense deleted successfuly"})
}
