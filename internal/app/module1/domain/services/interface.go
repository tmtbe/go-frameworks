package services

import (
	"test/internal/app/module1/domain/repos"
	"time"
)

type UserDetailService interface {
	GetUserDetail(ID uint64) (*UserDetail, error)
}

type UserDetail struct {
	ID          uint64
	UserName    string
	Password    string
	Email       string
	Name        string
	Price       float32
	CreatedTime time.Time
}

func FromInfraDetailAndUser(
	detail *repos.DetailRecord,
	user *repos.UserRecord,
) *UserDetail {
	if detail == nil {
		return nil
	}
	if user == nil {
		return nil
	}
	return &UserDetail{
		ID:          detail.ID,
		UserName:    user.UserName,
		Password:    user.Password,
		Email:       user.Email,
		Name:        detail.Name,
		Price:       detail.Price,
		CreatedTime: detail.CreatedAt,
	}
}
