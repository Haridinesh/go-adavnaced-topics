package router

import (
	"blogpost/handler"
	"blogpost/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router() {

	router := fiber.New()
	router.Post("/signup", handler.ToSignUpUser)
	router.Post("/login", handler.LoginUser)
	router.Get("/logout", middleware.AutherizationMiddleware(), handler.LogoutUser)
	router.Put("/admin/updateuser/:id", middleware.AutherizationMiddleware(), handler.UpdateUserData)
	router.Delete("/admin/deleteuser/:id", middleware.AutherizationMiddleware(), handler.DeleteUserData)
	router.Post("/admin/blogpost", middleware.AutherizationMiddleware(), handler.CreateNewPost)
	router.Put("/admin/blogpost/:id", middleware.AutherizationMiddleware(), handler.UpdateBlogPost)
	router.Delete("/admin/blogpost/:id", middleware.AutherizationMiddleware(), handler.DeletePost)

	router.Post("/admin/categories", middleware.AutherizationMiddleware(), handler.CreatingNewCategory)
	router.Put("/categories/:id", middleware.AutherizationMiddleware(), handler.ToUpdateCategory)
	router.Delete("/categories/:id", middleware.AutherizationMiddleware(), handler.DeleteCategory)

	router.Post("/comment/:postid", middleware.AutherizationMiddleware(), handler.AddCommentToPost)
	router.Put("/comment/:id", middleware.AutherizationMiddleware(), handler.ToupdateComment)
	router.Delete("/comment/:id", handler.DeleteComment)
	router.Get("/comments/:id", handler.CommentsOnPostsById)

	router.Get("/admin/comments/:id", handler.CommentsOnPostsByUserid)
	router.Get("/blogpost/comments/:id", middleware.AutherizationMiddleware(), handler.GetCommentsInPosts)
	router.Get("/blogpost/category", middleware.AutherizationMiddleware(), handler.ToGetPostByCategory)
	router.Get("/blogpost", middleware.AutherizationMiddleware(), handler.GetAllBlogPosts)
	router.Get("/blogpost/archives", handler.GetPostsByArchives)
	router.Get("/admin/overview", middleware.AutherizationMiddleware(), handler.OverviewOfPosts)
	router.Listen(":8080")
}
