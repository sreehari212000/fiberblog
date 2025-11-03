package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sreehari212000/blog/models"
)

type PostHanlder struct {
	db *sql.DB
}

func NewPostHandler(db *sql.DB) *PostHanlder {
	return &PostHanlder{db: db}
}

func (h *PostHanlder) CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}
	if post.Title == "" || post.Description == "" {
		return fiber.NewError(fiber.StatusBadRequest, "required field missing")
	}
	userId := c.Locals("user_id").(string)
	post.Author, _ = strconv.Atoi(userId)
	_, dbErr := h.db.Exec("INSERT INTO posts (title, description, author_id) VALUES ($1, $2, $3) RETURNING post_id", post.Title, post.Description, post.Author)
	if dbErr != nil {
		fmt.Println(dbErr)
		return fiber.NewError(fiber.StatusInternalServerError, "error inserting data into DB")
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "message": "post created succesfully", "data": post})
}
func (h *PostHanlder) DeletePost(c *fiber.Ctx) error {
	deleteID := c.Params("id")
	queryString := `DELETE FROM posts WHERE post_id = $1`
	res, dbErr := h.db.Exec(queryString, deleteID)
	if dbErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "coult not delete post")
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("post with id %s not found", deleteID))
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": "successfully deleted post", "id": deleteID})
}
func (h *PostHanlder) GetAllPosts(c *fiber.Ctx) error {
	queryString := `SELECT post_id, title, description, author_id FROM posts`
	rows, dbErr := h.db.Query(queryString)
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
func (h *PostHanlder) GetPostById(c *fiber.Ctx) error {
	postId := c.Params("id")
	queryString := `SELECT post_id, title, description, author_id FROM posts WHERE post_id = $1`
	var post models.Post
	err := h.db.QueryRow(queryString, postId).Scan(&post.POST_ID, &post.Title, &post.Description, &post.Author)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not find post with that id")
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": fmt.Sprintf("fetched post with id %v", postId), "data": post})
}
func (h *PostHanlder) LikePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	userId := c.Locals("user_id")
	queryString := `INSERT INTO likes(user_id, post_id) VALUES($1, $2) RETURNING like_id`
	var likeid int
	err := h.db.QueryRow(queryString, userId, postId).Scan(&likeid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error adding like to the post")
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "message": "liked post successfully", "likeId": postId})
}
