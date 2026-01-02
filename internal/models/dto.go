package models

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

type AddExpenseResponse struct {
	Message string `json:"message"`
}

type ListExpensesResponse struct {
	List []*Expense
}

type DeleteExpenseResponse struct {
	Message string `json:"message"`
}
