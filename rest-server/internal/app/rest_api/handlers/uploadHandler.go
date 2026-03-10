package handlers

import (
    "net/http"
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/services"
)

type StartUploadRequest struct {
    Filename string `json:"filename" binding:"required"`
    Filesize int64  `json:"filesize" binding:"required"`
}

type StartUploadResponse struct {
    UploadId   string `json:"uploadId"`
    TotalParts int    `json:"totalParts"`
    PartSize   int64  `json:"partSize"`
}

type CompleteUploadRequest struct {
    UploadId string                    `json:"uploadId" binding:"required"`
    Key      string                    `json:"key"      binding:"required"`
    Parts    []services.CompletedPart  `json:"parts"    binding:"required"`
}

type UploadHandler struct {
    uploadService *services.Upload
}

func NewUploadHandler(uploadService *services.Upload) *UploadHandler {
    return &UploadHandler{uploadService: uploadService}
}

func (h *UploadHandler) StartUpload(ctx *gin.Context) {
    var req StartUploadRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    uploadId, totalParts, partSize, err := h.uploadService.StartUpload(
        req.Filename,
        req.Filesize,
    )
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to start upload",
        })
        return
    }

    ctx.JSON(http.StatusOK, StartUploadResponse{
        UploadId:   uploadId,
        TotalParts: totalParts,
        PartSize:   partSize,
    })
}

func (h *UploadHandler) PresignPart(ctx *gin.Context) {
    uploadId := ctx.Query("uploadId")
    key := ctx.Query("key")
    partNumber := ctx.Query("partNumber")

    if uploadId == "" || key == "" || partNumber == "" {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "uploadId, key and partNumber are required",
        })
        return
    }

    partNum, err := strconv.Atoi(partNumber)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "partNumber must be a valid integer",
        })
        return
    }

    presignedURL, err := h.uploadService.PresignPart(uploadId, key, partNum)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to generate upload URL",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "presignedURL": presignedURL,
    })
}

func (h *UploadHandler) CompleteUpload(ctx *gin.Context) {
    var req CompleteUploadRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    if len(req.Parts) == 0 {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "parts cannot be empty",
        })
        return
    }

    err := h.uploadService.CompleteUpload(req.UploadId, req.Key, req.Parts)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to complete upload",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Upload complete",
        "key":     req.Key,
    })
}
