package dto

import "time"

type BlogReq struct {
	Content string `binding:"required"`
}

type BlogResp struct {
	BlogId    uint64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
