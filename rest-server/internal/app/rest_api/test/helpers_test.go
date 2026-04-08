package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func CreateJSONRequest(method, path string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}
	req := httptest.NewRequest(method, path, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func ParseResponseBody(r *httptest.ResponseRecorder, v interface{}) error {
	return json.Unmarshal(r.Body.Bytes(), v)
}
