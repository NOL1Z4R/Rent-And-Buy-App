package main

import (
	"Rent-And-Buy-App/config"
	"Rent-And-Buy-App/internal/api/handlers"
	"Rent-And-Buy-App/internal/api/routes"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//to configure later
	cfg := config.LoadConfig()
	dsn := config.GetDsn(cfg)
	db, _ := config.ConnectToDB(dsn)

	err := db.AutoMigrate(&entity.User{}, &entity.Car{}, &entity.House{})
	if err != nil {
		log.Fatalf("Error migrating users: %v", err)
	}

	//TODO burdaki *lara bir bak
	userRepo := repository.NewUserRepository(db)
	carRepo := repository.NewCarRepository(db)
	jwtManager := auth.NewJwtManager(cfg.JwtSecret, cfg.JwtExpireHours)
	authSrv := service.NewAuthService(*userRepo, jwtManager)
	userSrv := service.NewUserService(*userRepo)
	carSrv := service.NewCarService(*carRepo)

	authHandler := handlers.NewAuthHandler(*authSrv)
	userHandler := handlers.NewUserHandler(*userSrv)
	carHandler := handlers.NewCarHandler(carSrv)

	route := gin.Default()
	routes.SetupRoutes(route, authHandler, carHandler, userHandler, jwtManager)

	route.Run()
}
