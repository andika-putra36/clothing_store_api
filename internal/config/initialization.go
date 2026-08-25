package config

import (
	"clothing_store_api/internal/auth"
	"clothing_store_api/internal/cart"
	"clothing_store_api/internal/product"
	"clothing_store_api/internal/transaction"
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
	productRepository := product.NewRepository(db)
	cartRepository := cart.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	//Declaring service
	authService := auth.NewService(authRepository)
	productService := product.NewService(productRepository)
	cartService := cart.NewService(cartRepository)
	transactionService := transaction.NewService(transactionRepository)

	//Declaring handler
	authHandler := auth.NewHandler(authService)
	productHandler := product.NewHandler(productService)
	cartHandler := cart.NewHandler(cartService)
	transactionHandler := transaction.NewHandler(transactionService)

	router := gin.Default()
	v1 := router.Group("api/v1")

	auth.RegisterRoutes(v1, *authHandler)
	product.RegisterRoutes(v1, *productHandler)
	cart.RegisterRoutes(v1, *cartHandler)
	transaction.RegisterRoutes(v1, *transactionHandler)
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
