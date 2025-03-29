package tests

import (
	"bytes"
	"encoding/json"
	"expense-tracker/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutes() *gin.Engine{
	router := gin.Default()
	routes.SetupRoutes(router)
	return router
}

func TestSignup(t *testing.T){
	gin.SetMode(gin.TestMode)

	router := SetupRoutes()

	tests := []struct{
		name string
		payload map[string]any
		expectedCode int
		expectedBody string
	}{
		{
			name: "valid signup",
			payload: map[string]any{
				"username":"arya230",
				"password":"12345",
			},

			expectedCode: http.StatusCreated,
			expectedBody: `{"message":"You registered successfuly!"}`,
		},

		{	
			name: "duplicate username",
			payload: map[string]any{
				"username": "arya237",
				"password": "1234",
			},

			expectedCode: http.StatusConflict,
			expectedBody: `{"error":"this username already exists"}`,
		},

		{
			name: "missing username",
			payload: map[string]any{
				"password" : "12345",
			},

			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error": "Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},

		{
			name: "missing password",
			payload: map[string]any{
				"username" : "arya237",
			},

			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error": "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},

	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			
			payload, _ := json.Marshal(tt.payload)

			req, _ := http.NewRequest(http.MethodPost, "/user/signup", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.JSONEq(t, tt.expectedBody, recorder.Body.String())
		})
	}
}

