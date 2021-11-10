package domain

import (
	"time"
)

type User struct {
	UserId       uint64 `gorm:"primaryKey"`
	Email        string `gorm:"index:email_idx,unique"`
	PasswordHash string //密码hash值
	PasswordSolt string //密码盐值
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
