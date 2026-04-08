package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/handlers"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadHandler_StartUpload_MissingFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"filename": ""}`
	req := httptest.NewRequest(http.MethodPost, "/upload/start", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.StartUpload(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadHandler_StartUpload_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/upload/start", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.StartUpload(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadHandler_PresignPart_MissingParams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/upload/presign", nil)
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.PresignPart(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "uploadId, key and partNumber are required", response["error"])
}

func TestUploadHandler_PresignPart_InvalidPartNumber(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/upload/presign?uploadId=test&key=video.mp4&partNumber=abc", nil)
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.PresignPart(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "partNumber must be a valid integer", response["error"])
}

func TestUploadHandler_CompleteUpload_EmptyParts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"uploadId": "test-id", "key": "video.mp4", "parts": []}`
	req := httptest.NewRequest(http.MethodPost, "/upload/complete", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.CompleteUpload(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "parts cannot be empty", response["error"])
}

func TestUploadHandler_CompleteUpload_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/upload/complete", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.CompleteUpload(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadHandler_CompleteUpload_MissingUploadId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"uploadId": "", "key": "video.mp4", "parts": [{"partNumber": 1, "etag": "abc"}]}`
	req := httptest.NewRequest(http.MethodPost, "/upload/complete", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.CompleteUpload(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadHandler_CompleteUpload_MissingKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"uploadId": "test-id", "key": "", "parts": [{"partNumber": 1, "etag": "abc"}]}`
	req := httptest.NewRequest(http.MethodPost, "/upload/complete", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	handler := handlers.NewUploadHandler(nil)
	handler.CompleteUpload(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestStartUploadRequest_Binding(t *testing.T) {
	req := handlers.StartUploadRequest{
		Filename: "video.mp4",
		Filesize: 1024000,
	}
	assert.Equal(t, "video.mp4", req.Filename)
	assert.Equal(t, int64(1024000), req.Filesize)
}

func TestStartUploadResponse(t *testing.T) {
	resp := handlers.StartUploadResponse{
		UploadId:   "upload-123",
		TotalParts: 10,
		PartSize:   10485760,
	}
	assert.Equal(t, "upload-123", resp.UploadId)
	assert.Equal(t, 10, resp.TotalParts)
	assert.Equal(t, int64(10485760), resp.PartSize)
}

func TestCompleteUploadRequest(t *testing.T) {
	parts := []services.CompletedPart{
		{PartNumber: 1, ETag: "etag1"},
		{PartNumber: 2, ETag: "etag2"},
	}
	req := handlers.CompleteUploadRequest{
		UploadId: "upload-123",
		Key:      "videos/video.mp4",
		Parts:    parts,
	}
	assert.Equal(t, "upload-123", req.UploadId)
	assert.Equal(t, "videos/video.mp4", req.Key)
	assert.Len(t, req.Parts, 2)
}
