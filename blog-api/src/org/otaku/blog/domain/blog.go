package domain

import "time"

type Blog struct {
	BlogId    uint64 `gorm:"primaryKey"`
	UserId    uint64 `gorm:"index:user_id_idx"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
