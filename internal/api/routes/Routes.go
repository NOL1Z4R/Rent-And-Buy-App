package routes

import (
	"Rent-And-Buy-App/internal/api/handlers"
	"Rent-And-Buy-App/internal/api/middlewares"
	"Rent-And-Buy-App/pkg/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, jwtMgr *auth.JWTManager) {
	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		protected := api.Group("/")
		protected.Use(middlewares.JWTAuth(jwtMgr))
		{
			protected.GET("/user", userHandler.GetAll)
			protected.GET("/user/:id", userHandler.GetById)
			protected.PUT("/user/:id", userHandler.UpdateUser)
			protected.DELETE("/user/:id", userHandler.DeleteUser)
		}
	}
}
