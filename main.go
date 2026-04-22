package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/nwildan922/learn-go-manager/routes"
)

func main() {

	loadEnv()
	log.Println(os.Getenv("APP_NAME"))
	log.Println(os.Getenv("APP_PORT"))

	port := os.Getenv("APP_PORT")

	e := echo.New()
	routes.RegisterRoutes(e)

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
