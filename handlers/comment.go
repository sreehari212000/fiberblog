package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sreehari212000/blog/models"
)

type CommentHandler struct {
	db *sql.DB
}

func NewCommentHandler(db *sql.DB) *CommentHandler {
	return &CommentHandler{db: db}
}

func (h *CommentHandler) AddComment(c *fiber.Ctx) error {
	var comment models.Comment
	userId := c.Locals("user_id")
	postID := c.Params("id")
	if err := c.BodyParser(&comment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}
	if comment.Text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please add comment")
	}
	queryString := `INSERT INTO comments(post_id, author_id, text) VALUES ($1, $2, $3)`
	_, err := h.db.Exec(queryString, postID, userId, comment.Text)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "error inserting data into DB")
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "message": "comment created succesfully", "data": comment})
}
func (h *CommentHandler) GetPostComments(c *fiber.Ctx) error {
	postId := c.Params("id")
	queryString := `SELECT comment_id, post_id, author_id, text FROM comments WHERE post_id = $1`
	rows, err := h.db.Query(queryString, postId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error fetching data from DB")
	}
	var comment models.Comment
	var comments []models.Comment
	for rows.Next() {
		if err := rows.Scan(&comment.Comment_id, &comment.Post_id, &comment.Author_id, &comment.Text); err != nil {
			fmt.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError, "error fetching comments")
		}
		comments = append(comments, comment)
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "message": "comments fetched successfullt", "data": comments})
}
