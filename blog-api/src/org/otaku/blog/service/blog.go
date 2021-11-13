package service

import (
	"org/otaku/blog/db"
	"org/otaku/blog/domain"
	"org/otaku/blog/dto"

	"github.com/jinzhu/copier"
)

type blogService struct{}

var BlogService blogService = blogService{}

func (s blogService) AddBlog(userId uint64, reqDTO *dto.BlogReq) (uint64, error) {
	blog := domain.Blog{
		UserId:  userId,
		Content: reqDTO.Content,
	}
	if err := db.GetDB().Create(&blog).Error; err != nil {
		return 0, err
	}
	return blog.BlogId, nil
}

func (s blogService) GetBlogList(userId uint64) ([]dto.BlogResp, error) {
	var blogs []domain.Blog
	if err := db.GetDB().Where("user_id = ?", userId).Find(&blogs).Error; err != nil {
		return nil, err
	}
	respList := make([]dto.BlogResp, len(blogs))
	copier.Copy(&respList, &blogs)
	return respList, nil
}
