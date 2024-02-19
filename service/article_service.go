package service

import (
	"awesome/model"
	"fmt"
)

type ArticleService struct {
}

func (s ArticleService) Create(t string, b string) model.Article {
	resp, _ :=
		DbCall(fmt.Sprintf("INSERT INTO article(title, body) VALUES ('%s', '%s')", t, b))
	id := ExtractValue(resp, "last_insert_rowid")
	return model.Article{
		Id:    id,
		Title: t,
		Body:  b,
	}
}

func (s ArticleService) Read(id string) model.Article {
	data, _ := DbCall(fmt.Sprintf("SELECT * FROM article where id = %s", id))
	result, _ := ExtractRows(data, "rows")
	return model.Article{
		Id:    result[0][0].Value,
		Title: result[0][1].Value,
		Body:  result[0][2].Value,
	}
}

func (s ArticleService) ReadAll() []model.Article {
	data, _ := DbCall("SELECT * FROM article")
	result, _ := ExtractRows(data, "rows")

	articles := make([]model.Article, len(result))
	for i, entry := range result {
		articles[i] = model.Article{
			Id:    entry[0].Value,
			Title: entry[1].Value,
			Body:  entry[2].Value,
		}
	}

	return articles
}
