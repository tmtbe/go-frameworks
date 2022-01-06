package services

import (
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
