package routes

import (
	"Rent-And-Buy-App/internal/api/handlers"
	"Rent-And-Buy-App/internal/api/middlewares"
	"Rent-And-Buy-App/pkg/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	authHandler *handlers.AuthHandler,
	carHandler *handlers.CarHandler,
	houseHandler *handlers.HouseHandler,
	userHandler *handlers.UserHandler,
	jwtMgr *auth.JWTManager) {
	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		protected := api.Group("/")
		protected.Use(middlewares.JWTAuth(jwtMgr))
		{
			user := protected.Group("/user")
			{
				user.GET("/", userHandler.GetAll)
				user.GET("/:id", userHandler.GetById)
				user.PUT("/:id", userHandler.UpdateUser)
				user.DELETE("/:id", userHandler.DeleteUser)
			}
			car := protected.Group("/car")
			{
				car.GET("/", carHandler.GetCarAll)
				car.GET("/:id", carHandler.GetCarById)
				car.GET("/plate/:plate", carHandler.GetCarByPlate)
				car.POST("/", carHandler.CreateCar)
				car.PUT("/:id", carHandler.UpdateCar)
				car.DELETE("/:id", carHandler.DeleteCar)
			}
			house := protected.Group("/house")
			{
				house.GET("/", houseHandler.GetHouseAll)
				house.GET("/:id", houseHandler.GetHouse)
				house.POST("/", houseHandler.CreateHouse)
				house.PUT("/:id", houseHandler.UpdateHouse)
				house.DELETE("/:id", houseHandler.DeleteHouse)
			}
		}
	}
}
