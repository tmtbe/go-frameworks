package services

import (
	"go.uber.org/zap"
	"test/internal/app/module1/domain/exceptions"
	"test/internal/app/module1/domain/repos"
)

type UserDetailServiceImpl struct {
	logger           *zap.Logger
	detailRepository repos.DetailRepository
	userRepository   repos.UserRepository
}

func NewUserDetailServiceImpl(
	logger *zap.Logger,
	detailRepository repos.DetailRepository,
	userRepository repos.UserRepository,
) *UserDetailServiceImpl {
	u := &UserDetailServiceImpl{
		logger:           logger.With(zap.String("type", "UserDetailServiceImpl")),
		detailRepository: detailRepository,
		userRepository:   userRepository,
	}
	return u
}

func (s *UserDetailServiceImpl) GetUserDetail(ID uint64) (p *UserDetail, err error) {
	d := s.detailRepository.FindDetailById(ID)
	u := s.userRepository.FindUserById(ID)
	p = FromInfraDetailAndUser(d, u)
	if p == nil {
		err = exceptions.BusinessError("Can't find any user or detail")
	}
	return
}
