package routes

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/middlewares"
	"go-rust-drop/internal/api/routes/auth"
	"go-rust-drop/internal/api/routes/public"
)

type Route interface {
}

func RouteHandle(router *gin.Engine) {
	publicGroup := router.Group("/api/v1")
	public.Routes(publicGroup)

	authGroup := router.Group("/api/v1")
	authMiddleware := middlewares.AuthMiddleware{}

	authGroup.Use(authMiddleware.Authenticated)
	{
		auth.Routes(authGroup)
	}
}
