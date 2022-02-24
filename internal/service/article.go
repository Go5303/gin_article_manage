package service

type CreateArticleRequest struct {
	Title     string `form:"title" binding:"required,min=2,max=100"`
	Desc 	  string `form:"desc" binding:"required,min=2,max=100"`
	Content   string `form:"content" binding:"required"`
	CreatedBy string `form:"created_by" binding:"required"`
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CreatedBy)
}
