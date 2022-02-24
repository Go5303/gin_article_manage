package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string  {
	return "blog_article"
}

/**
创建文章
*/
func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}