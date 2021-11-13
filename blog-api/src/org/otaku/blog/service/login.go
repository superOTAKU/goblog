package service

import (
	"errors"
	"fmt"
	"math/rand"
	"org/otaku/blog/db"
	"org/otaku/blog/domain"
	"org/otaku/blog/dto"
	"org/otaku/blog/util"
	"strconv"

	"gorm.io/gorm"
)

//定义单例的loginService
type loginService struct{}

var LoginService loginService = loginService{}

//为loginService定义方法

//登录
func (s loginService) Login(loginReq *dto.LoginReq) (*dto.AccessToken, error) {
	var user domain.User
	if err := db.GetDB().Where("email = ?", loginReq.Email).Find(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("email not exists")
	}
	passwordHash, err := util.Sha1(loginReq.Password, user.PasswordSolt)
	if err != nil {
		return nil, fmt.Errorf("check password error")
	}
	if passwordHash != user.PasswordHash {
		return nil, fmt.Errorf("password incorrect")
	}
	token, err := util.GenJWTToken(user.UserId)
	if err != nil {
		return nil, fmt.Errorf("gen jwt token error")
	}
	return &dto.AccessToken{Token: token}, nil
}

//注册
func (s loginService) Register(registerReq *dto.RegisterReq) error {
	var count int64
	db.GetDB().Model(&domain.User{}).Where("email = ?", registerReq.Email).Count(&count)
	if count > 0 {
		return fmt.Errorf("email exists")
	}
	user := &domain.User{
		Email:        registerReq.Email,
		PasswordSolt: randomSolt(),
	}
	user.PasswordHash, _ = util.Sha1(registerReq.Password, user.PasswordSolt)
	db.GetDB().Create(user)
	return nil
}

func randomSolt() string {
	solt := ""
	for i := 0; i < 4; i++ {
		solt += strconv.Itoa(rand.Intn(10))
	}
	return solt
}
