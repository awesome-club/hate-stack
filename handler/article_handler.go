package handler

import (
	"awesome/service"
	"awesome/view/article"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleService service.ArticleService
}

func (h ArticleHandler) ShowArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	a := h.ArticleService.Read(id)
	return render(ctx, article.Show(a))
}

func (h ArticleHandler) CreateArticle(ctx echo.Context) error {
	t := ctx.FormValue("title")
	b := ctx.FormValue("body")
	a := h.ArticleService.Create(t, b)
	return render(ctx, article.Show(a))
}
