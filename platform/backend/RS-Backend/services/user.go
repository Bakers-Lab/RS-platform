// user_service.go

package service

import (
	"RS-Backend/dal/db"
	"RS-Backend/models/dao"
	"RS-Backend/models/dto"
	"context"
)

type UserService interface {
	Register(ctx context.Context, user *dao.User) error
	LogIn(ctx context.Context,  email string, password string) (*dto.User, error)
}

type UserServiceImpl struct {
	accesser db.IUserAccesser
}

func NewUserService(accesser db.IUserAccesser) UserService {
	return &UserServiceImpl{accesser: accesser}
}

func (s *UserServiceImpl) Register(ctx context.Context, user *dao.User) error {
	return s.accesser.Register(ctx, user)
}
func (s *UserServiceImpl) LogIn(ctx context.Context,  email string, password string) (*dto.User, error) {
	return s.accesser.LogIn(ctx, email, password)
}