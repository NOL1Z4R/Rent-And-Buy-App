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

	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("Error migrating users: %v", err)
	}

	//TODO burdaki *lara bir bak
	userRepo := repository.NewUserRepository(db)
	jwtManager := auth.NewJwtManager(cfg.JwtSecret, cfg.JwtExpireHours)
	authSrv := service.NewAuthService(*userRepo, jwtManager)
	userSrv := service.NewUserService(*userRepo)

	authHandler := handlers.NewAuthHandler(*authSrv)
	userHandler := handlers.NewUserHandler(*userSrv)

	route := gin.Default()
	routes.SetupRoutes(route, authHandler, userHandler, jwtManager)

	route.Run()
}
