package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sreehari212000/blog/handlers"
	"github.com/sreehari212000/blog/middlewares"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/api")

	// user routes
	auth := api.Group("/auth")
	{
		auth.Post("/signup", handlers.SignUp)
		auth.Post("/signin", handlers.Login)
	}
	// posts and comments routes
	post := api.Group("/posts")
	{
		post.Get("/", handlers.GetAllPosts)
		post.Post("/", middlewares.CheckAuth, handlers.CreatePost)
		post.Get("/:id", handlers.GetPostById)
		post.Delete("/:id", handlers.DeletePost)
		post.Post("/:id/comments", middlewares.CheckAuth, handlers.AddComment)
		post.Get("/:id/comments", handlers.GetPostComments)
		post.Post("/:id/likes", middlewares.CheckAuth, handlers.LikePost)
	}

}
