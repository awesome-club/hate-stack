package handler

import (
	"awesome/model"
	"awesome/view/comment"
	"awesome/service"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct{
	CommentService service.CommentService
}

func (h CommentHandler) CreateComment(ctx echo.Context) error {
	aid := ctx.FormValue("articleId")
	c := ctx.FormValue("comment")

	h.CommentService.Create(aid, c)

	return render(ctx, comment.Show(model.Comment{
		ArticleId: aid,
		Body:      c,
	}))
}
