package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	dbconfig "github.com/sreehari212000/blog/dbConfig"
	"github.com/sreehari212000/blog/logs"
	"github.com/sreehari212000/blog/middlewares"
	"github.com/sreehari212000/blog/routes"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		logs.Error("error loading .env file!")
	}
	connectionString := os.Getenv("POSTGRESQL_CONNECTION_STRING")
	db, dbErr := dbconfig.InitDB(connectionString)
	if dbErr != nil {
		log.Fatal("Database Error: ", dbErr)
	}
	defer dbconfig.DB.Close()
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})
	routes.InitializeRoutes(app, db)
	port := os.Getenv("PORT")
	logs.Info(fmt.Sprintf("server started on port %v", port))
	app.Listen(fmt.Sprintf(":%v", port))
}
