package config

import (
	"clothing_store_api/internal/admin"
	"clothing_store_api/internal/auth"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeEverything() *gin.Engine {
	db := initializeDB()
	// Declaring repository
	authRepository := auth.NewRepository(db)
	adminRepository := admin.NewRepository(db)

	//Declaring service
	authService := auth.NewService(authRepository)
	adminService := admin.NewService(adminRepository)

	//Declaring handler
	authHandler := auth.NewHandler(authService)
	adminHandler := admin.NewHandler(adminService)

	router := gin.Default()
	v1 := router.Group("api/v1")

	auth.RegisterRoutes(v1, *authHandler)
	admin.RegisterRoutes(v1, *adminHandler)
	return router
}

func initializeDB() *gorm.DB {
	env := getEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.Host,
		env.User,
		env.Password,
		env.Name,
		env.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database is failed to connect")
	}

	return db
}

func getEnv() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting .env file")
	}

	return Env{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}
