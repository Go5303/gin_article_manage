package dao

import (
	"blog-serice/internal/model"
)

func (d *Dao) CreateArticle(title string, desc string, content string, createdBy string) error {
	article := model.Article{
		Title: title,
		Desc: desc,
		Content: content,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return article.Create(d.engine)
}
