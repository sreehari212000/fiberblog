package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/sreehari212000/blog/dbConfig"
	"github.com/sreehari212000/blog/models"
	"github.com/sreehari212000/blog/utils"
)

type UserHandler struct {
	db *sql.DB
}

// DEPENDENCY INJECTION
func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func SignUp(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}
	hashedPassword, hashError := utils.HashPassword(user.Password)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "something went wrong")
	}
	user.Password = hashedPassword
	_, dbErr := dbconfig.InsertIntoDb(user.Name, user.Email, user.Password)
	if dbErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "could not insert into DB")
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "message": "user signed up successfully", "data": user})
}

func Login(c *fiber.Ctx) error {
	var user models.LoginRequestBody
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}
	if user.Email == "" || user.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}
	queryString := `SELECT user_id, email, password FROM users WHERE email = $1`
	var id int
	var email, password string
	err := dbconfig.DB.QueryRow(queryString, user.Email).Scan(&id, &email, &password)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	if user.Password != password {
		return fiber.NewError(401, "could not authenticate you")
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": "user logged in succesfully", "data": password})
}
