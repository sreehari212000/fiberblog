package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	dbconfig "github.com/sreehari212000/blog/dbConfig"
	"github.com/sreehari212000/blog/models"
)

func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}
	if post.Title == "" || post.Description == "" {
		return fiber.NewError(fiber.StatusBadRequest, "required field missing")
	}
	// todo: extract user details from c and store instead of the hardcoded value
	userId := 1
	post.Author = userId
	_, dbErr := dbconfig.DB.Exec("INSERT INTO posts (title, description, author_id) VALUES ($1, $2, $3) RETURNING post_id", post.Title, post.Description, post.Author)
	if dbErr != nil {
		fmt.Println(dbErr)
		return fiber.NewError(fiber.StatusInternalServerError, "error inserting data into DB")
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "message": "post created succesfully", "data": post})
}
func DeletePost(c *fiber.Ctx) error {
	deleteID := c.Params("id")
	queryString := `DELETE FROM posts WHERE post_id = $1`
	res, dbErr := dbconfig.DB.Exec(queryString, deleteID)
	if dbErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "coult not delete post")
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("post with id %s not found", deleteID))
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": "successfully deleted post", "id": deleteID})
}
func GetAllPosts(c *fiber.Ctx) error {
	queryString := `SELECT post_id, title, description, author_id FROM posts`
	rows, dbErr := dbconfig.DB.Query(queryString)
	if dbErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "could not fetch posts")
	}
	defer rows.Close()
	var posts []models.Post
	var post models.Post
	for rows.Next() {
		if err := rows.Scan(&post.POST_ID, &post.Title, &post.Description, &post.Author); err != nil {
			fmt.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError, "error fetching posts")
		}
		posts = append(posts, post)
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": "fetched posts succesfully", "data": posts})
}
func GetPostById(c *fiber.Ctx) error {
	postId := c.Params("id")
	queryString := `SELECT post_id, title, description, author_id FROM posts WHERE post_id = $1`
	var post models.Post
	err := dbconfig.DB.QueryRow(queryString, postId).Scan(&post.POST_ID, &post.Title, &post.Description, &post.Author)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not find post with that id")
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": fmt.Sprintf("fetched post with id %v", postId), "data": post})
}
