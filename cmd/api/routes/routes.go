package routes

import (
	"github.com/gin-gonic/gin"
	"lightsaid.com/millionare/cmd/api/handlers"
)

// NewRoutes 返回路由
func NewRoutes(handler *handlers.APIHandler) *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", handler.Register)
		apiv1.POST("/login", handler.Login)
	}

	userv1 := apiv1.Group("/user")
	{
		userv1.POST("", handler.UpdateUser)
		userv1.POST("/avatar", handler.UpdateAvatar)
	}

	return r
}
