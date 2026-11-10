package biz

import (
	"context"
	"kratos-layout/internal/data/repo"
)

type UserBiz struct {
	repo *repo.UserRepo
}

func NewUserBiz(repo *repo.UserRepo) *UserBiz {
	return &UserBiz{repo: repo}
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255;uniqueIndex"`
}

func (u *UserBiz) GetUserByEmail(ctx context.Context, email string) (string, error) {
	return u.repo.GetById(ctx, 5)
}
