package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/handlers"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/models/dtos"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "abc"}}

	handler := handlers.NewUserHandler(nil)
	handler.GetUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User ID not valid", response["error"])
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUserHandler(nil)
	handler.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUser_ShortFirstName(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"first_name": "Jo", "last_name": "Doe", "email": "john@example.com", "phone_number": "1234567890"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUserHandler(nil)
	handler.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response, "errors")
}

func TestUpdateUser_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"first_name": "John", "last_name": "Doe", "email": "john@example.com", "phone_number": "1234567890"}`
	req := httptest.NewRequest(http.MethodPut, "/users/abc", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "abc"}}

	handler := handlers.NewUserHandler(nil)
	handler.UpdateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteUser_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodDelete, "/users/abc", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "abc"}}

	handler := handlers.NewUserHandler(nil)
	handler.DeleteUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"first_name": "John", "last_name": "Doe", "email": "invalid-email", "phone_number": "1234567890"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUserHandler(nil)
	handler.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUserRequest_ToEntity(t *testing.T) {
	dto := &dtos.CreateUserRequest{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	user := dto.ToUser()
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "1234567890", user.PhoneNumber)
}

func TestUpdateUserRequest_ToEntity(t *testing.T) {
	dto := &dtos.UpdateUserRequest{
		FirstName:   "Jane",
		LastName:    "Smith",
		Email:       "jane@example.com",
		PhoneNumber: "0987654321",
	}

	user := dto.ToUser()
	assert.Equal(t, "Jane", user.FirstName)
	assert.Equal(t, "Smith", user.LastName)
	assert.Equal(t, "jane@example.com", user.Email)
	assert.Equal(t, "0987654321", user.PhoneNumber)
}
