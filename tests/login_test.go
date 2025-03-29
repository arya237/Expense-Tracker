package tests

import(
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T){
	gin.SetMode(gin.TestMode)
	router := SetupRoutes()

	tests := []struct{
		name string
		payload map[string]any
		expectedCode int
	}{
		{
			name: "valid login",
			payload: map[string]any{
				"username": "arya237",
				"password": "12345",
			},

			expectedCode: http.StatusOK,
		},

		{
			name: "missing password",
			payload: map[string]any{
				"username":"arya237",
			},

			expectedCode: http.StatusBadRequest,
		},

		{
			name: "missing username",
			payload: map[string]any{
				"password": "12345",
			},

			expectedCode: http.StatusBadRequest,
		},

		{
			name: "wrong username",
			payload: map[string]any{
				"username": "arya231",
				"password": "12345",
			},

			expectedCode: http.StatusUnauthorized,
		},

		{
			name: "wrong password",
			payload: map[string]any{
				"username": "arya237",
				"password": "123456",
			},

			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests{

		t.Run(tt.name, func(t *testing.T) {

			payload, _ := json.Marshal(tt.payload)

			req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(payload))

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expectedCode, recorder.Code)
		})
	}


}