package handler

import (
	"awesome/service"
	"awesome/view/home"
	"awesome/view/write"
	"github.com/labstack/echo/v4"
)

type BasicHandler struct {
	ArticleService service.ArticleService
}

func (h BasicHandler) ShowHome(ctx echo.Context) error {
	articles := h.ArticleService.ReadAll()
	return render(ctx, home.ShowHome(articles))
}

func (h BasicHandler) ShowWrite(ctx echo.Context) error {
	return render(ctx, write.ShowWrite())
}
