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
	"gorm.io/gorm"
	"log"
)

func main() {
	var db *gorm.DB
	var err error
	//to configure later
	cfg := config.LoadConfig()
	dsn := config.GetDsn(cfg)
	db, err = config.ConnectToDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Car{}, &entity.House{}, &entity.Post{}, &entity.Product{}, &entity.Sale{})
	if err != nil {
		log.Fatalf("Error migrating users: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	carRepo := repository.NewCarRepository(db)
	houseRepo := repository.NewHouseRepository(db)
	postRepo := repository.NewPostRepository(db)
	productRepo := repository.NewProductRepository(db)
	saleRepo := repository.NewSaleRepository(db)
	jwtManager := auth.NewJwtManager(cfg.JwtSecret, cfg.JwtExpireHours)

	authSrv := service.NewAuthService(userRepo, jwtManager)
	userSrv := service.NewUserService(userRepo)
	carSrv := service.NewCarService(carRepo)
	houseSrv := service.NewHouseService(houseRepo)
	productSrv := service.NewProductService(productRepo, carSrv, houseSrv)
	postSrv := service.NewPostService(postRepo, productSrv, userSrv)
	saleSrv := service.NewSaleService(saleRepo, userSrv, postSrv)

	authHandler := handlers.NewAuthHandler(*authSrv)
	userHandler := handlers.NewUserHandler(*userSrv)
	carHandler := handlers.NewCarHandler(carSrv)
	houseHandler := handlers.NewHouseHandler(houseSrv)
	postHandler := handlers.NewPostHandler(*postSrv)
	saleHandler := handlers.NewSaleHandler(saleSrv)

	route := gin.Default()
	routes.SetupRoutes(route, authHandler, carHandler, houseHandler, postHandler, saleHandler, userHandler, jwtManager)

	route.Use(gin.Recovery())
	route.Run()
}
