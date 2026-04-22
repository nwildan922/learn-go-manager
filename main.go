package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	handler "github.com/nwildan922/learn-go-manager/handler"
	redis "github.com/nwildan922/learn-go-manager/pkg/redis"
	"github.com/nwildan922/learn-go-manager/routes"
	"github.com/nwildan922/learn-go-manager/service"
)

func main() {

	loadEnv()
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
	log.Println(os.Getenv("APP_NAME"))
	log.Println(os.Getenv("APP_PORT"))

	port := os.Getenv("APP_PORT")

	e := echo.New()

	counterService := service.NewCounterService(redisClient)
	counterHandler := handler.NewCounterHandler(counterService)

	routes.RegisterRoutes(e, &routes.Router{
		Counter: counterHandler,
	})

	e.Logger.Fatal(e.Start(":" + port))

}

func loadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "develop"
	}

	file := ".env." + env

	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading %s file", file)
	}

	log.Println("Loaded environment:", env)
}
