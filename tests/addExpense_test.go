package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddExpense(t *testing.T){
	gin.SetMode(gin.TestMode)
	router := SetupRoutes()

	tests := []struct{
		name string
		payload map[string]any
		expectedCode int
	}{
		{
			name: "valid add expense",
			payload: map[string]any{
				"UserID" : "arya230",
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusCreated,
		},

		{
			name: "invalid user name",
			payload: map[string]any{
				"UserID" : "arya234",
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "invalid dead line",
			payload: map[string]any{
				"UserID" : "arya234",
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"DeadLine": "2024-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing userID",
			payload: map[string]any{
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing title",
			payload: map[string]any{
				"UserID" : "arya230",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing category",
			payload: map[string]any{
				"UserID" : "arya230",
				"Title"  : "reading",		 
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing date",
			payload: map[string]any{
				"UserID" : "arya230",
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Status"  : "pending",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing status",
			payload: map[string]any{
				"UserID" : "arya230",
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"DeadLine": "2025-04-04",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing dead line",
			payload: map[string]any{
				"UserID" : "arya230",
				"Title"  : "reading",		
				"Category": "self improvment",  
				"Date" 	  : "2025-03-05",
				"Status"  : "pending",
				"Description" : "",
			},
			expectedCode: http.StatusBadRequest,
		},	
	}

	key := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDMzMDc5MzN9.VGYBJIQJaARS23ojVDK3TqUDB3ufy1203Azkr53stxsF_3sI6WdQzeE0EUkBxxLuh0m0i3NQAUh3fzxP-vGXCw"

	for _, tt := range tests{
		payload, _ := json.Marshal(tt.payload)

		req, _ := http.NewRequest(http.MethodPost,"/expense/addExpense", bytes.NewBuffer(payload))
		req.Header.Set("Authorization", key)
		
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, tt.expectedCode, recorder.Code)
	}
}