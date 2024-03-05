package repository

import (
	"context"
	"mosong/webook/internal/domain"
	"mosong/webook/internal/repository/dao"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	err := ur.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
	return err
}
