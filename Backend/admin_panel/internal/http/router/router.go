package router

import (
	"github.com/gin-gonic/gin"

	"Backend/admin_panel/internal/http/handlers"
	"Backend/admin_panel/internal/http/middleware"
)

func New(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("/:id", userHandler.Get)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
			users.GET("", userHandler.List)
		}
	}
	return router
}
