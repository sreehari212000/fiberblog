package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	dbconfig "github.com/sreehari212000/blog/dbConfig"
	"github.com/sreehari212000/blog/handlers"
	"github.com/sreehari212000/blog/middlewares"
	"github.com/sreehari212000/blog/routes"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file!")
	}
	connectionString := os.Getenv("POSTGRESQL_CONNECTION_STRING")
	newdb, dbErr := dbconfig.InitDB(connectionString)
	if dbErr != nil {
		log.Fatal("Database Error: ", dbErr)
	}
	defer dbconfig.DB.Close()
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	userHanlder := handlers.NewUserHandler(newdb) // checking Dependency Injection DI
	app.Get("/", userHanlder.Sample)              // DI
	routes.InitializeRoutes(app)
	app.Listen(":3000")
}
