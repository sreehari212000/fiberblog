package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/sreehari212000/blog/handlers"
	"github.com/sreehari212000/blog/middlewares"
)

func InitializeRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")
	// initialize handlers
	userHandler := handlers.NewUserHandler(db)
	postHandler := handlers.NewPostHandler(db)
	commentHandler := handlers.NewCommentHandler(db)
	// user routes
	auth := api.Group("/auth")
	{
		auth.Post("/signup", userHandler.SignUp)
		auth.Post("/signin", userHandler.Login)
	}
	// posts and comments routes
	post := api.Group("/posts")
	{
		post.Get("/", postHandler.GetAllPosts)
		post.Post("/", middlewares.CheckAuth, postHandler.CreatePost)
		post.Get("/:id", postHandler.GetPostById)
		post.Delete("/:id", postHandler.DeletePost)
		post.Post("/:id/comments", middlewares.CheckAuth, commentHandler.AddComment)
		post.Get("/:id/comments", commentHandler.GetPostComments)
		post.Post("/:id/likes", middlewares.CheckAuth, postHandler.LikePost)
	}

}
