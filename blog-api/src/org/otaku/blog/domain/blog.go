package domain

import "time"

type Blog struct {
	BlogId   uint64 `gorm:"primaryKey"`
	UserId   uint64 `gorm:"index:user_id_idx"`
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
}
