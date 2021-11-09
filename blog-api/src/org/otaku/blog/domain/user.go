package domain

import (
	"time"
)

type User struct {
	UserId       uint64 `gorm:"primaryKey"`
	Email        string
	PasswordHash string //密码hash值
	PasswordSolt string //密码盐值
	CreateAt     time.Time
	UpdateAt     time.Time
}
