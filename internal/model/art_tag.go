package model

type ArticleTag struct {
	*Model
	ArticleId int32 `json:"article_id"`
	TagId 	  int32 `json:"tag_id"`
}

func (a ArticleTag) tableName() string  {
	return "blog_article_tag"
}
