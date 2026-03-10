package routes

import (
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPublicEndpoints(router *gin.Engine, userHandlers *handlers.User) {
	router.GET("/users", userHandlers.GetAllUsers)
	router.GET("/users/:id", userHandlers.GetUser)
	router.POST("/users", userHandlers.CreateUser)
	router.PUT("/users/:id", userHandlers.UpdateUser)
	router.DELETE("/users/:id", userHandlers.DeleteUser)
}

func RegisterUploadEndpoints(router *gin.Engine, uploadHandler *handlers.UploadHandler) {
    upload := router.Group("/upload")
    {
        upload.POST("/start",    uploadHandler.StartUpload)
        upload.GET("/presign",   uploadHandler.PresignPart)
        upload.POST("/complete", uploadHandler.CompleteUpload)
    }
}
