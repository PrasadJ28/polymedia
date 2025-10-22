package routes

import (
	"github.com/gin-gonic/gin"
)

type Registrar func(*gin.RouterGroup)

func Init() *gin.Engine {
	return CombineRoutes("", VideoRoutes)
}

func CombineRoutes(prefix string, regs ...Registrar) *gin.Engine {
	r := gin.Default()
	base := r.Group(prefix)

	for _, reg := range regs {
		reg(base)
	}
	return r
}
