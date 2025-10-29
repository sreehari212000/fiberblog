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
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file!")
	}
	connectionString := os.Getenv("POSTGRESQL_CONNECTION_STRING")
	_, dbErr := dbconfig.InitDB(connectionString)
	if dbErr != nil {
		log.Fatal("Database Error: ", dbErr)
	}
	defer dbconfig.DB.Close()
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	api := app.Group("/api")
	// user routes
	api.Post("/users/signup", handlers.SignUp)
	api.Post("/users/signin", handlers.Login)
	// posts routes
	api.Get("/posts", handlers.GetAllPosts)
	api.Post("/posts", handlers.CreatePost)
	api.Get("/posts/:id", handlers.GetPostById)
	api.Delete("/posts/:id", handlers.DeletePost)
	// api.Patch("/posts/:id", handlers)
	app.Listen(":3000")
}
