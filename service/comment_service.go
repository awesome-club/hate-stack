package service

import (
	"fmt"
)

type CommentService struct {
}

func (s CommentService) Create(aid string, c string) error {
	_, err := DbCall(fmt.Sprintf("INSERT INTO comment(artcile_id, body) VALUES ('%s', '%s')", aid, c))
	return err
}
