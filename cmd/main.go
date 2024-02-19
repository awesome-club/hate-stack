package main

import (
	"awesome/handler"
	"awesome/service"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	// Services
	articleService := service.ArticleService{}
	commentService := service.CommentService{}

	// Handlers
	basicHandler := handler.BasicHandler{
		ArticleService: articleService,
	}
	articleHandler := handler.ArticleHandler{
		ArticleService: articleService,
	}
	commentHandler := handler.CommentHandler{
		CommentService: commentService,
	}

	app.GET("/", basicHandler.ShowHome)
	app.GET("/write", basicHandler.ShowWrite)
	app.GET("/articles/:id", articleHandler.ShowArticle)
	app.POST("/articles", articleHandler.CreateArticle)

	app.POST("/comments", commentHandler.CreateComment)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app.Logger.Fatal(app.Start(port))
}
