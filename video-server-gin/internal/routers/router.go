package router

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/video-server-gin/internal/controllers"
	mid "github.com/video-server-gin/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	public := router.Group("/api")
	{
		public.POST("", ctrl.Register)

	}

	protected := router.Group("/api")
	protected.Use(mid.AuthMiddleWare())
	{
		protected.GET("", ctrl.CreateUser)
	}

	return router
}
