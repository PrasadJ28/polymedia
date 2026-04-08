package models_test

import (
	"testing"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/models"
	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	errResp := &models.ErrorResponse{
		Code:    404,
		Message: "Not Found",
	}

	assert.Equal(t, 404, errResp.Code)
	assert.Equal(t, "Not Found", errResp.Message)
}

func TestErrorResponse_Default(t *testing.T) {
	errResp := &models.ErrorResponse{}

	assert.Equal(t, 0, errResp.Code)
	assert.Equal(t, "", errResp.Message)
}

func TestErrorResponse_VariousCodes(t *testing.T) {
	tests := []struct {
		name    string
		code    int
		message string
	}{
		{"Bad Request", 400, "Invalid input"},
		{"Unauthorized", 401, "Authentication required"},
		{"Forbidden", 403, "Access denied"},
		{"Not Found", 404, "Resource not found"},
		{"Internal Server Error", 500, "Something went wrong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errResp := &models.ErrorResponse{
				Code:    tt.code,
				Message: tt.message,
			}
			assert.Equal(t, tt.code, errResp.Code)
			assert.Equal(t, tt.message, errResp.Message)
		})
	}
}
